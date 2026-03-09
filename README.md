# Instana Go Client Library

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://github.com/instana/instana-go-client)

A comprehensive, production-ready Go client library for the [Instana](https://www.instana.com/) API. This library provides a clean, idiomatic Go interface for interacting with Instana's monitoring and observability platform.

## Features

### 🚀 Production-Ready
- **Comprehensive Configuration**: Flexible configuration via builder pattern, environment variables, or JSON files
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
    
    "github.com/instana/instana-go-client/instana"
)

func main() {
    // Create configuration using builder pattern
    config, err := instana.NewConfigBuilder().
        WithBaseURL("https://tenant-unit.instana.io").
        WithAPIToken("your-api-token").
        WithConnectionTimeout(30 * time.Second).
        WithMaxRetryAttempts(3).
        Build()
    
    if err != nil {
        log.Fatalf("Failed to create config: %v", err)
    }
    
    // Use the configuration
    log.Printf("Client configured for: %s", config.BaseURL)
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

#### 2. Environment Variables

```bash
export INSTANA_BASE_URL=https://tenant-unit.instana.io
export INSTANA_API_TOKEN=your-api-token
export INSTANA_CONNECTION_TIMEOUT=30s
export INSTANA_MAX_RETRY_ATTEMPTS=3
export INSTANA_RATE_LIMIT_RPS=100
```

```go
config, err := instana.LoadFromEnv()
```

#### 3. JSON Configuration File

```json
{
  "baseURL": "https://tenant-unit.instana.io",
  "apiToken": "your-api-token",
  "timeout": {
    "connection": "30s",
    "request": "60s"
  },
  "retry": {
    "maxAttempts": 3,
    "initialDelay": "1s",
    "maxDelay": "30s"
  }
}
```

```go
config, err := instana.LoadFromJSON("config.json")
```

#### 4. Hybrid Approach (JSON + Environment Override)

```go
config, err := instana.LoadFromJSONWithEnvOverride("config.json")
```

## Configuration Options

### Timeouts

```go
config := instana.NewConfigBuilder().
    WithConnectionTimeout(30 * time.Second).      // Default: 30s
    WithRequestTimeout(60 * time.Second).         // Default: 60s
    WithIdleConnectionTimeout(90 * time.Second).  // Default: 90s
    WithResponseHeaderTimeout(10 * time.Second).  // Default: 10s
    WithTLSHandshakeTimeout(10 * time.Second).    // Default: 10s
    Build()
```

### Retry Configuration

```go
config := instana.NewConfigBuilder().
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
config := instana.NewConfigBuilder().
    WithRateLimitEnabled(true).                   // Default: true
    WithRateLimitRequestsPerSecond(50).           // Default: 100
    WithRateLimitBurstCapacity(100).              // Default: 200
    WithRateLimitWaitForToken(true).              // Default: true
    Build()
```

### Custom Headers

```go
config := instana.NewConfigBuilder().
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
config := instana.NewConfigBuilder().
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
config := instana.NewConfigBuilder().
    WithBatchSize(100).                           // Default: 100
    WithBatchConcurrentRequests(5).               // Default: 5
    WithBatchStopOnError(false).                  // Default: false
    WithBatchRetryFailedItems(true).              // Default: true
    Build()
```

### Logging

```go
// Use default logger with log level
logger := instana.NewDefaultLogger(instana.ClientLogLevelInfo)

// Or use no-op logger to disable logging
logger := instana.NewNoOpLogger()

config := instana.NewConfigBuilder().
    WithLogger(logger).
    WithDebug(true).
    Build()
```

## Error Handling

The library provides typed errors for better error handling:

```go
import "github.com/instana/instana-go-client/instana"

// Check if error is retryable
if instana.IsRetryableError(err) {
    // Retry the operation
}

// Check if error is temporary
if instana.IsTemporaryError(err) {
    // Wait and retry
}

// Extract HTTP status code
statusCode := instana.ExtractStatusCode(err)

// Type assertion for detailed error information
if instanaErr, ok := err.(*instana.InstanaError); ok {
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
    "github.com/instana/instana-go-client/instana"
)

// Use the built-in retryer
retryConfig := instana.DefaultRetryConfig()
retryer := instana.NewRetryer(retryConfig, logger)

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
    "github.com/instana/instana-go-client/instana"
)

rateLimitConfig := instana.RateLimitConfig{
    Enabled:           true,
    RequestsPerSecond: 50,
    BurstCapacity:     100,
    WaitForToken:      true,
}

rateLimiter := instana.NewRateLimiter(rateLimitConfig, logger)
defer rateLimiter.Stop()

// Wait for rate limit token before making request
if err := rateLimiter.Wait(context.Background()); err != nil {
    log.Printf("Rate limit error: %v", err)
}
```

## Environment Variables

The library supports the following environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `INSTANA_BASE_URL` | Base URL of Instana API | - |
| `INSTANA_API_TOKEN` | API token for authentication | - |
| `INSTANA_DEBUG` | Enable debug logging | false |
| `INSTANA_CONNECTION_TIMEOUT` | Connection timeout | 30s |
| `INSTANA_REQUEST_TIMEOUT` | Request timeout | 60s |
| `INSTANA_IDLE_CONNECTION_TIMEOUT` | Idle connection timeout | 90s |
| `INSTANA_MAX_RETRY_ATTEMPTS` | Maximum retry attempts | 3 |
| `INSTANA_RETRY_INITIAL_DELAY` | Initial retry delay | 1s |
| `INSTANA_RETRY_MAX_DELAY` | Maximum retry delay | 30s |
| `INSTANA_RETRY_BACKOFF_MULTIPLIER` | Backoff multiplier | 2.0 |
| `INSTANA_BATCH_SIZE` | Batch operation size | 100 |
| `INSTANA_BATCH_CONCURRENT_REQUESTS` | Concurrent batch requests | 5 |
| `INSTANA_RATE_LIMIT_ENABLED` | Enable rate limiting | true |
| `INSTANA_RATE_LIMIT_RPS` | Requests per second | 100 |
| `INSTANA_RATE_LIMIT_BURST` | Burst capacity | 200 |
| `INSTANA_MAX_IDLE_CONNECTIONS` | Max idle connections | 100 |
| `INSTANA_MAX_CONNECTIONS_PER_HOST` | Max connections per host | 10 |
| `INSTANA_KEEP_ALIVE_DURATION` | Keep-alive duration | 30s |

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
