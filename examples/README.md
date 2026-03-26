# Instana Go Client Examples

This directory contains practical examples demonstrating how to use the Instana Go Client library.

## 📁 Examples

### 1. Basic Usage
**Location:** [`basic_usage/`](basic_usage/)

Demonstrates the fundamental usage patterns of the Instana Go Client:
- Basic client creation (legacy compatible)
- Using the builder pattern for configuration
- Enhanced error handling with typed errors

**Run:**
```bash
cd basic_usage
go run main.go
```

### 2. Advanced Usage
**Location:** [`advanced_usage/`](advanced_usage/)

Demonstrates advanced features and patterns:
- Comprehensive configuration with all options
- Advanced error handling and classification
- Retry mechanism with exponential backoff
- Rate limiting with token bucket algorithm
- Custom logger implementation

**Run:**
```bash
cd advanced_usage
go run main.go
```

### 3. REST Client Usage
**Location:** [`rest_client_usage/`](rest_client_usage/)

Demonstrates low-level REST client operations:
- Direct REST API calls
- Custom HTTP operations
- Request/response handling
- Production-ready configuration

**Run:**
```bash
cd rest_client_usage
go run main.go
```

## 🚀 Quick Start

### Prerequisites
```bash
go get github.com/instana/instana-go-client
```

### Basic Example
```go
package main

import (
    "fmt"
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
        log.Fatal(err)
    }
    
    // Create API client
    api, err := instana.NewInstanaAPIWithConfig(cfg)
    if err != nil {
        log.Fatal(err)
    }
    
    // Use the API
    tokens, err := api.APITokens().GetAll()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Retrieved %d tokens\n", len(*tokens))
}
```

## 📚 More Resources

- [Quick Start Guide](../QUICK_START.md) - Comprehensive getting started guide
- [Architecture Guide](../ARCHITECTURE.md) - System architecture and design
- [API Reference](../API_REFERENCE.md) - Complete API documentation
- [Usage Guide](../USAGE_GUIDE.md) - Detailed usage patterns
- [API Documentation](https://pkg.go.dev/github.com/instana/instana-go-client)

## 💡 Tips

1. **Use Builder Pattern**: For production code, use the builder pattern for better configuration control
2. **Enable Retry**: Configure retry mechanism for better reliability
3. **Rate Limiting**: Enable rate limiting to avoid hitting API limits
4. **Error Handling**: Use typed errors for better error handling
5. **Connection Pooling**: Configure connection pooling for better performance

## 🆘 Need Help?

- [GitHub Issues](https://github.com/instana/instana-go-client/issues)
- [Documentation](../README.md)