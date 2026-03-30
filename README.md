# Instana Go Client Library

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://github.com/instana/instana-go-client)

A comprehensive, production-ready Go client library for the [Instana](https://www.instana.com/) API. This library provides a clean, idiomatic Go interface for interacting with Instana's monitoring and observability platform.

## Features

### 🚀 Production-Ready
- **Comprehensive Configuration**: Flexible configuration via builder pattern
- **Automatic Retry**: Exponential backoff with jitter for transient failures
- **Rate Limiting**: Token bucket algorithm to prevent API throttling
- **Connection Pooling**: Optimized HTTP connection management
- **Typed Errors**: Rich error types for better error handling
- **Structured Logging**: Built-in logging with sensitive data redaction

### 🎯 Developer-Friendly
- **Fluent API**: Intuitive builder pattern for configuration
- **Sensible Defaults**: Works out of the box with minimal configuration
- **Context Support**: All operations respect context cancellation
- **Thread-Safe**: Safe for concurrent use
- **Zero Dependencies**: Minimal external dependencies

### 📦 Complete API Coverage
- Application Monitoring Configuration
- Custom Events and Metrics
- Infrastructure Monitoring
- Synthetic Monitoring
- Alert Configurations (Application, Website, Mobile, Smart Alerts)
- SLO/SLI Management
- User and Group Management
- API Tokens Management
- Maintenance Windows
- Custom Dashboards

## Installation

```bash
go get github.com/instana/instana-go-client
```

## Quick Start

### Basic Usage

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
        WithConnectionTimeout(30 * time.Second).
        WithMaxRetryAttempts(3).
        Build()
    
    if err != nil {
        log.Fatalf("Failed to create config: %v", err)
    }
    
    // Create API client
    api, err := instana.NewInstanaAPIWithConfig(cfg)
    if err != nil {
        log.Fatalf("Failed to create API: %v", err)
    }
    
    // Use the API
    tokens, err := api.APITokens().GetAll()
    if err != nil {
        log.Fatalf("Failed to get tokens: %v", err)
    }
    
    log.Printf("Retrieved %d tokens", len(*tokens))
}
```

### Configuration Methods

#### 1. Builder Pattern (Recommended)

```go
config, err := instana.NewConfigBuilder().
    WithBaseURL("https://tenant-unit.instana.io").
    WithAPIToken("your-api-token").
    WithConnectionTimeout(45 * time.Second).
    WithRequestTimeout(90 * time.Second).
    WithMaxRetryAttempts(5).
    WithRetryInitialDelay(2 * time.Second).
    WithRateLimitRequestsPerSecond(50).
    WithCustomHeader("X-Custom-Header", "value").
    WithDebug(true).
    Build()
```

#### 2. Using Default Configuration

```go
cfg := config.DefaultClientConfig()
cfg.BaseURL = "https://tenant-unit.instana.io"
cfg.APIToken = "your-api-token"
cfg.Retry.MaxAttempts = 5

api, err := instana.NewInstanaAPIWithConfig(cfg)
```

## Configuration Options

### Timeouts

```go
cfg, err := config.NewConfigBuilder().
    WithConnectionTimeout(30 * time.Second).      // Default: 30s
    WithRequestTimeout(60 * time.Second).         // Default: 60s
    WithIdleConnectionTimeout(90 * time.Second).  // Default: 90s
    WithResponseHeaderTimeout(10 * time.Second).  // Default: 10s
    WithTLSHandshakeTimeout(10 * time.Second).    // Default: 10s
    Build()
```

### Retry Configuration

```go
cfg, err := config.NewConfigBuilder().
    WithMaxRetryAttempts(5).                      // Default: 3
    WithRetryInitialDelay(2 * time.Second).       // Default: 1s
    WithRetryMaxDelay(60 * time.Second).          // Default: 30s
    WithRetryBackoffMultiplier(2.5).              // Default: 2.0
    WithRetryOnTimeout(true).                     // Default: true
    WithRetryOnConnectionError(true).             // Default: true
    WithRetryJitter(true).                        // Default: true
    Build()
```

### Rate Limiting

```go
cfg, err := config.NewConfigBuilder().
    WithRateLimitEnabled(true).                   // Default: true
    WithRateLimitRequestsPerSecond(50).           // Default: 100
    WithRateLimitBurstCapacity(100).              // Default: 200
    WithRateLimitWaitForToken(true).              // Default: true
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

### Connection Pooling

```go
cfg, err := config.NewConfigBuilder().
    WithMaxIdleConnections(100).                  // Default: 100
    WithMaxConnectionsPerHost(10).                // Default: 10
    WithMaxIdleConnectionsPerHost(10).            // Default: 10
    WithKeepAliveDuration(30 * time.Second).      // Default: 30s
    WithDisableKeepAlives(false).                 // Default: false
    WithDisableCompression(false).                // Default: false
    Build()
```

### Batch Operations

```go
cfg, err := config.NewConfigBuilder().
    WithBatchSize(100).                           // Default: 100
    WithBatchConcurrentRequests(5).               // Default: 5
    WithBatchStopOnError(false).                  // Default: false
    WithBatchRetryFailedItems(true).              // Default: true
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

## Error Handling

The library provides typed errors for better error handling:

```go
import "github.com/instana/instana-go-client/config"

// Check if error is retryable
if config.IsRetryableError(err) {
    // Retry the operation
}

// Check if error is temporary
if config.IsTemporaryError(err) {
    // Wait and retry
}

// Extract HTTP status code
statusCode := config.ExtractStatusCode(err)

// Type assertion for detailed error information
if instanaErr, ok := err.(*config.InstanaError); ok {
    fmt.Printf("Error Type: %s\n", instanaErr.Type)
    fmt.Printf("Status Code: %d\n", instanaErr.StatusCode)
    fmt.Printf("Retryable: %v\n", instanaErr.IsRetryable())
}
```

### Error Types

- `ErrorTypeNetwork` - Network connectivity errors (retryable)
- `ErrorTypeAPI` - API response errors (may be retryable)
- `ErrorTypeValidation` - Input validation errors (not retryable)
- `ErrorTypeAuthentication` - Authentication failures (not retryable)
- `ErrorTypeRateLimit` - Rate limit exceeded (retryable)
- `ErrorTypeTimeout` - Request timeout (retryable)
- `ErrorTypeSerialization` - JSON serialization errors (not retryable)

## Retry Mechanism

The library includes automatic retry with exponential backoff:

```go
import (
    "context"
    "github.com/instana/instana-go-client/config"
)

// Use the built-in retryer
retryConfig := config.DefaultRetryConfig()
retryer := config.NewRetryer(retryConfig, logger)

err := retryer.Do(context.Background(), func() error {
    // Your operation here
    return someAPICall()
})
```

### Retry Features

- **Exponential Backoff**: Delay increases exponentially (default multiplier: 2.0)
- **Jitter**: Random jitter (up to 30%) prevents thundering herd
- **Configurable Conditions**: Retry based on error type or HTTP status code
- **Context Support**: Respects context cancellation
- **Detailed Logging**: Logs each retry attempt with delay information

## Rate Limiting

Built-in rate limiting using token bucket algorithm:

```go
import (
    "context"
    "github.com/instana/instana-go-client/config"
)

rateLimitConfig := config.RateLimitConfig{
    Enabled:           true,
    RequestsPerSecond: 50,
    BurstCapacity:     100,
    WaitForToken:      true,
}

rateLimiter := config.NewRateLimiter(rateLimitConfig, logger)
defer rateLimiter.Stop()

// Wait for rate limit token before making request
if err := rateLimiter.Wait(context.Background()); err != nil {
    log.Printf("Rate limit error: %v", err)
}
```

## Examples

See the [examples](examples/) directory for complete examples:

- [Basic Usage](examples/basic_usage/) - Simple client initialization and configuration
- More examples coming soon...

## Documentation

- [Implementation Plan](IMPLEMENTATION_PLAN.md) - Detailed implementation roadmap
- [Progress Summary](PROGRESS_SUMMARY.md) - Current progress and status
- [API Documentation](https://pkg.go.dev/github.com/instana/instana-go-client) - Full API reference

## Development

### Building

```bash
go build ./...
```

### Testing

```bash
go test ./...
```

### Linting

```bash
golangci-lint run
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Support

- **Issues**: [GitHub Issues](https://github.com/instana/instana-go-client/issues)
- **Documentation**: [pkg.go.dev](https://pkg.go.dev/github.com/instana/instana-go-client)
- **Instana Support**: [support.instana.com](https://support.instana.com)

## Roadmap

- [x] Phase 1: Package migration from terraform provider
- [x] Phase 2: Configuration system with validation
- [x] Phase 3: Error handling, retry, and rate limiting (60% complete)
- [ ] Phase 4: Comprehensive testing
- [ ] Phase 5: CI/CD pipeline
- [ ] Phase 6: Complete documentation and examples
- [ ] Phase 7: v1.0.0 release

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for version history and changes.

---

**Note**: This library is currently under active development. The API may change before the v1.0.0 release.
