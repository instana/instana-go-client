# Instana Go Client Library

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://github.com/instana/instana-go-client)

A comprehensive, production-ready Go client library for the [Instana](https://www.instana.com/) API. This library provides a clean, idiomatic Go interface for interacting with Instana's monitoring and observability platform.


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
### 1. Simple Client Creation 

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/instana/instana-go-client/instana"
)

func main() {
    // Create client with basic configuration
    client := instana.NewClient(
        "your-api-token",
        "your-tenant.instana.io",
        false, // skipTlsVerification
    )
    
    // Use the client...
    data, err := client.Get("/api/application-monitoring/applications")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Response: %s\n", string(data))
}
```

#### 2. Builder Pattern (Recommended)

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

api, err := instana.NewInstanaAPIWithConfig(cfg)
```

#### 3. Using Default Configuration

```go
cfg := config.DefaultClientConfig()
cfg.BaseURL = "https://tenant-unit.instana.io"
cfg.APIToken = "your-api-token"
cfg.Retry.MaxAttempts = 5

api, err := instana.NewInstanaAPIWithConfig(cfg)
```


## Examples

See the [examples](examples/) directory for complete examples:

- [Basic Usage](examples/basic_usage/) - Simple client initialization and configuration
- More examples coming soon...

## Documentation

- [Implementation Plan](IMPLEMENTATION_PLAN.md) - Detailed implementation roadmap

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


## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Support

- **Issues**: [GitHub Issues](https://github.com/instana/instana-go-client/issues)



## Changelog

See [CHANGELOG.md](CHANGELOG.md) for version history and changes.

---
