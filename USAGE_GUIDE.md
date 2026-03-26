# Instana Go Client - Usage Guide

## Overview

This guide demonstrates how to use the Instana Go Client library effectively. The configuration system is fully implemented and integrated with the REST client.

---

## Basic Usage

### Simple Client Creation

The simplest way to create a client (legacy compatible):

```go
package main

import (
    "github.com/instana/instana-go-client/instana"
)

func main() {
    // Simple client creation
    client := instana.NewClient(
        "your-api-token",
        "tenant.instana.io",
        false, // skipTlsVerification
    )
    
    // Use the client for API calls
    data, err := client.Get("/api/application-monitoring/applications")
    if err != nil {
        // Handle error
    }
}
```

---

## Advanced Usage with Configuration

### Method 1: Using Builder Pattern (Recommended)

```go
package main

import (
    "log"
    "time"
    
    "github.com/instana/instana-go-client/config"
    "github.com/instana/instana-go-client/instana"
)

func main() {
    // Create configuration using builder pattern
    cfg, err := config.NewConfigBuilder().
        WithBaseURL("https://tenant-unit.instana.io").
        WithAPIToken("your-api-token").
        WithConnectionTimeout(45 * time.Second).
        WithRequestTimeout(90 * time.Second).
        WithMaxRetryAttempts(5).
        WithRetryInitialDelay(2 * time.Second).
        WithRateLimitRequestsPerSecond(50).
        WithCustomHeader("X-Request-ID", "unique-id").
        WithDebug(true).
        Build()
    
    if err != nil {
        log.Fatalf("Failed to build config: %v", err)
    }
    
    // Create API client with configuration
    api, err := instana.NewInstanaAPIWithConfig(cfg)
    if err != nil {
        log.Fatalf("Failed to create API: %v", err)
    }
    
    // Use the API - now with retry, rate limiting, etc.
    tokens, err := api.APITokens().GetAll()
    if err != nil {
        log.Fatalf("API call failed: %v", err)
    }
    
    log.Printf("Retrieved %d tokens", len(*tokens))
}
```

### Method 2: Using Default Configuration

```go
package main

import (
    "log"
    
    "github.com/instana/instana-go-client/config"
    "github.com/instana/instana-go-client/instana"
)

func main() {
    // Start with default configuration
    cfg := config.DefaultClientConfig()
    cfg.BaseURL = "https://tenant-unit.instana.io"
    cfg.APIToken = "your-api-token"
    
    // Customize as needed
    cfg.Retry.MaxAttempts = 5
    cfg.RateLimit.RequestsPerSecond = 50
    
    // Validate configuration
    if err := cfg.Validate(); err != nil {
        log.Fatalf("Invalid config: %v", err)
    }
    
    // Create API client
    api, err := instana.NewInstanaAPIWithConfig(cfg)
    if err != nil {
        log.Fatalf("Failed to create API: %v", err)
    }
    
    // Use the API
    apps, err := api.ApplicationConfigs().GetAll()
    if err != nil {
        log.Fatalf("Failed to get applications: %v", err)
    }
    
    log.Printf("Retrieved %d applications", len(*apps))
}
```

---

## Configuration Options

### Timeouts

```go
cfg, err := config.NewConfigBuilder().
    WithConnectionTimeout(30 * time.Second).
    WithRequestTimeout(60 * time.Second).
    WithIdleConnectionTimeout(90 * time.Second).
    WithResponseHeaderTimeout(10 * time.Second).
    WithTLSHandshakeTimeout(10 * time.Second).
    Build()
```

### Retry Configuration

```go
cfg, err := config.NewConfigBuilder().
    WithMaxRetryAttempts(5).
    WithRetryInitialDelay(1 * time.Second).
    WithRetryMaxDelay(30 * time.Second).
    WithRetryBackoffMultiplier(2.0).
    WithRetryOnTimeout(true).
    WithRetryOnConnectionError(true).
    WithRetryJitter(true).
    Build()
```

### Rate Limiting

```go
cfg, err := config.NewConfigBuilder().
    WithRateLimitEnabled(true).
    WithRateLimitRequestsPerSecond(100).
    WithRateLimitBurstCapacity(200).
    WithRateLimitWaitForToken(true).
    Build()
```

### Connection Pooling

```go
cfg, err := config.NewConfigBuilder().
    WithMaxIdleConnections(100).
    WithMaxConnectionsPerHost(10).
    WithMaxIdleConnectionsPerHost(10).
    WithKeepAliveDuration(30 * time.Second).
    WithDisableKeepAlives(false).
    WithDisableCompression(false).
    Build()
```

### Custom Headers

```go
cfg, err := config.NewConfigBuilder().
    WithCustomHeader("X-Request-ID", "unique-id").
    WithCustomHeader("X-Trace-ID", "trace-123").
    WithCustomHeaders(map[string]string{
        "X-Custom-1": "value1",
        "X-Custom-2": "value2",
    }).
    Build()
```

### Logging

```go
// Use default logger with log level
logger := config.NewDefaultLogger(config.ClientLogLevelInfo)

// Or use no-op logger to disable logging
logger := config.NewNoOpLogger()

cfg, err := config.NewConfigBuilder().
    WithLogger(logger).
    WithDebug(true).
    Build()
```

---

## Error Handling

### Basic Error Handling

```go
data, err := client.Get("/api/endpoint")
if err != nil {
    log.Printf("Error: %v\n", err)
}
```

### Advanced Error Handling with Typed Errors

```go
import "github.com/instana/instana-go-client/config"

data, err := client.Get("/api/endpoint")
if err != nil {
    // Check if it's an Instana error
    if instanaErr, ok := err.(*config.InstanaError); ok {
        switch instanaErr.Type {
        case config.ErrorTypeAuthentication:
            log.Println("Authentication failed")
        case config.ErrorTypeRateLimit:
            log.Println("Rate limit exceeded")
        case config.ErrorTypeValidation:
            log.Println("Validation error")
        case config.ErrorTypeNetwork:
            log.Println("Network error")
        case config.ErrorTypeTimeout:
            log.Println("Request timeout")
        }
        
        log.Printf("Status Code: %d\n", instanaErr.StatusCode)
        log.Printf("Message: %s\n", instanaErr.Message)
        log.Printf("Retryable: %v\n", instanaErr.IsRetryable())
    }
    
    // Check if error is retryable
    if config.IsRetryableError(err) {
        log.Println("This error can be retried")
    }
    
    // Check if error is temporary
    if config.IsTemporaryError(err) {
        log.Println("This is a temporary error")
    }
    
    // Extract status code
    statusCode := config.ExtractStatusCode(err)
    log.Printf("HTTP Status Code: %d\n", statusCode)
}
```

---

## Using the API Client

### Accessing Resources

```go
// Create API client
api, err := instana.NewInstanaAPIWithConfig(cfg)
if err != nil {
    log.Fatal(err)
}

// API Tokens
tokens, err := api.APITokens().GetAll()
token, err := api.APITokens().GetOne("token-id")
created, err := api.APITokens().Create(newToken)
updated, err := api.APITokens().Update(token)
err = api.APITokens().Delete(token)

// Application Configurations
apps, err := api.ApplicationConfigs().GetAll()
app, err := api.ApplicationConfigs().GetOne("app-id")
created, err := api.ApplicationConfigs().Create(newApp)

// Alert Configurations
alerts, err := api.ApplicationAlertConfigs().GetAll()
alert, err := api.ApplicationAlertConfigs().GetOne("alert-id")

// And many more resources...
```

---

## Production-Ready Configuration

```go
cfg, err := config.NewConfigBuilder().
    // Core settings
    WithBaseURL("https://tenant-unit.instana.io").
    WithAPIToken("your-api-token").
    WithUserAgent("MyApp/1.0.0").
    
    // Timeouts
    WithConnectionTimeout(10 * time.Second).
    WithRequestTimeout(30 * time.Second).
    WithIdleConnectionTimeout(90 * time.Second).
    
    // Retry with exponential backoff
    WithMaxRetryAttempts(5).
    WithRetryInitialDelay(1 * time.Second).
    WithRetryMaxDelay(30 * time.Second).
    WithRetryBackoffMultiplier(2.0).
    WithRetryJitter(true).
    WithRetryOnTimeout(true).
    WithRetryOnConnectionError(true).
    
    // Rate limiting
    WithRateLimitEnabled(true).
    WithRateLimitRequestsPerSecond(10).
    WithRateLimitBurstCapacity(20).
    
    // Connection pooling
    WithMaxIdleConnections(100).
    WithMaxConnectionsPerHost(10).
    WithKeepAliveDuration(90 * time.Second).
    
    // Custom headers for tracking
    WithCustomHeader("X-Application", "my-app").
    WithCustomHeader("X-Environment", "production").
    
    // Logging
    WithDebug(false).
    Build()

if err != nil {
    log.Fatalf("Failed to build config: %v", err)
}

api, err := instana.NewInstanaAPIWithConfig(cfg)
if err != nil {
    log.Fatalf("Failed to create API: %v", err)
}
```

---

## Migration from Legacy Client

### Before (Legacy)

```go
client := instana.NewClient("token", "tenant.instana.io", false)
data, err := client.Get("/api/endpoint")
```

### After (With Configuration)

```go
cfg, _ := config.NewConfigBuilder().
    WithBaseURL("https://tenant.instana.io").
    WithAPIToken("token").
    WithMaxRetryAttempts(5).
    WithRateLimitEnabled(true).
    Build()

client, _ := instana.NewClientWithConfig(cfg)
data, err := client.Get("/api/endpoint")
```

---

## Best Practices

1. **Use Builder Pattern**: For production code, use the builder pattern for better configuration control
2. **Enable Retry**: Configure retry mechanism for better reliability
3. **Rate Limiting**: Enable rate limiting to avoid hitting API limits
4. **Error Handling**: Use typed errors for better error handling
5. **Connection Pooling**: Configure connection pooling for better performance
6. **Custom Headers**: Add custom headers for request tracking and debugging
7. **Logging**: Use appropriate log levels for different environments

---

## See Also

- [README](README.md) - Project overview and features
- [Quick Start Guide](QUICK_START.md) - Getting started quickly
- [API Reference](API_REFERENCE.md) - Complete API documentation
- [Architecture](ARCHITECTURE.md) - System architecture
- [Examples](examples/) - Code examples