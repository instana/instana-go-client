# Instana Go Client Examples

This directory contains practical examples demonstrating how to use the Instana Go Client library.

## 📁 Examples

### 1. basic_usage.go
Demonstrates the fundamental usage patterns of the Instana Go Client:
- Basic client creation (legacy compatible)
- Using the builder pattern for configuration
- Loading configuration from environment variables
- Enhanced error handling with typed errors

**Run:**
```bash
# Set environment variables
export INSTANA_BASE_URL="https://your-tenant.instana.io"
export INSTANA_API_TOKEN="your-api-token"

# Run the example
go run basic_usage.go
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
    "github.com/instana/instana-go-client/instana"
)

func main() {
    // Create client
    client := instana.NewClient("api-token", "tenant.instana.io", false)
    
    // Make request
    data, err := client.Get("/api/application-monitoring/applications")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Retrieved %d bytes\n", len(data))
}
```

## 📚 More Resources

- [Quick Start Guide](../QUICK_START.md) - Comprehensive getting started guide
- [Migration Guide](../MIGRATION_GUIDE.md) - Migrating from Terraform provider
- [Configuration Guide](../DEFAULT_CONFIG_ANALYSIS.md) - Detailed configuration options
- [API Documentation](https://pkg.go.dev/github.com/instana/instana-go-client)

## 💡 Tips

1. **Use Builder Pattern**: For production code, use the builder pattern for better configuration control
2. **Enable Retry**: Configure retry mechanism for better reliability
3. **Rate Limiting**: Enable rate limiting to avoid hitting API limits
4. **Error Handling**: Use typed errors for better error handling
5. **Environment Variables**: Use environment variables for configuration in different environments

## 🆘 Need Help?

- [GitHub Issues](https://github.com/instana/instana-go-client/issues)
- [Documentation](../README.md)