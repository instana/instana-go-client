# Quick Start Guide - Instana Go Client

**Version**: v0.9.0 (Pre-release)  
**Target**: v1.0.0  
**Last Updated**: 2026-03-09

---

## 📦 Installation

```bash
go get github.com/instana/instana-go-client
```

---

## 🚀 Basic Usage

### 1. Simple Client Creation (Legacy API)

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

### 2. Using Configuration Builder (Recommended)

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
        WithBaseURL("https://your-tenant.instana.io").
        WithAPIToken("your-api-token").
        WithConnectionTimeout(30 * time.Second).
        WithRequestTimeout(60 * time.Second).
        WithMaxRetryAttempts(3).
        WithRateLimitEnabled(true).
        WithRateLimitRequestsPerSecond(100).
        WithDebug(true).
        Build()
    
    if err != nil {
        log.Fatal(err)
    }
    
    // Create API client with configuration
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

### 3. Using Default Configuration

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/instana/instana-go-client/config"
    "github.com/instana/instana-go-client/instana"
)

func main() {
    // Start with default configuration
    cfg := config.DefaultClientConfig()
    cfg.BaseURL = "https://your-tenant.instana.io"
    cfg.APIToken = "your-api-token"
    
    // Customize as needed
    cfg.Retry.MaxAttempts = 5
    cfg.RateLimit.RequestsPerSecond = 50
    
    // Validate configuration
    if err := cfg.Validate(); err != nil {
        log.Fatal(err)
    }
    
    // Create API client
    api, err := instana.NewInstanaAPIWithConfig(cfg)
    if err != nil {
        log.Fatal(err)
    }
    
    // Use the API
    apps, err := api.ApplicationConfigs().GetAll()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Retrieved %d applications\n", len(*apps))
}
```

---

## 🔧 REST Client API

### GET Requests

```go
// Get all resources
data, err := client.Get("/api/application-monitoring/applications")

// Get single resource by ID
data, err := client.GetOne("app-id-123", "/api/application-monitoring/applications")

// Get with query parameters
queryParams := map[string]string{
    "name": "MyApp",
    "page": "1",
}
data, err := client.GetByQuery("/api/application-monitoring/applications", queryParams)
```

### POST Requests

```go
// Create a new resource
type Application struct {
    Label string `json:"label"`
    // ... other fields
}

app := Application{Label: "My Application"}
data, err := client.Post(app, "/api/application-monitoring/applications")

// POST with query parameters
data, err := client.PostByQuery("/api/endpoint", queryParams)

// POST with ID in response
data, err := client.PostWithID(app, "/api/application-monitoring/applications")
```

### PUT Requests

```go
// Update a resource
app := Application{Label: "Updated Application"}
data, err := client.Put(app, "/api/application-monitoring/applications/app-id-123")

// PUT with query parameters
data, err := client.PutByQuery("/api/endpoint", "resource-id", queryParams)
```

### DELETE Requests

```go
// Delete a resource
err := client.Delete("app-id-123", "/api/application-monitoring/applications")
```

---

## ⚙️ Configuration Options

### Timeout Configuration

```go
config, _ := instana.NewConfigBuilder().
    WithConnectionTimeout(30 * time.Second).
    WithRequestTimeout(60 * time.Second).
    WithIdleConnectionTimeout(90 * time.Second).
    Build()
```

### Retry Configuration

```go
config, _ := instana.NewConfigBuilder().
    WithMaxRetryAttempts(5).
    WithRetryInitialDelay(1 * time.Second).
    WithRetryMaxDelay(30 * time.Second).
    WithRetryBackoffMultiplier(2.0).
    WithRetryOnTimeout(true).
    WithRetryOnConnectionError(true).
    Build()
```

### Rate Limiting

```go
config, _ := instana.NewConfigBuilder().
    WithRateLimitEnabled(true).
    WithRateLimitRequestsPerSecond(100).
    WithRateLimitBurstCapacity(200).
    Build()
```

### Connection Pooling

```go
config, _ := instana.NewConfigBuilder().
    WithMaxIdleConnections(100).
    WithMaxConnectionsPerHost(50).
    WithKeepAliveDuration(30 * time.Second).
    Build()
```

### Debug Mode

```go
config, _ := instana.NewConfigBuilder().
    WithDebug(true).
    Build()
```

---

## 🔍 Error Handling

### Basic Error Handling

```go
data, err := client.Get("/api/endpoint")
if err != nil {
    if err == instana.ErrEntityNotFound {
        log.Println("Resource not found")
    } else {
        log.Printf("Error: %v\n", err)
    }
}
```

### Advanced Error Handling (New Error System)

```go
data, err := client.Get("/api/endpoint")
if err != nil {
    // Check if it's an Instana error
    var instanaErr *instana.InstanaError
    if errors.As(err, &instanaErr) {
        switch instanaErr.Type {
        case instana.ErrorTypeAuthentication:
            log.Println("Authentication failed")
        case instana.ErrorTypeRateLimit:
            log.Println("Rate limit exceeded")
        case instana.ErrorTypeValidation:
            log.Println("Validation error")
        case instana.ErrorTypeServer:
            log.Println("Server error")
        }
        
        log.Printf("Status Code: %d\n", instanaErr.StatusCode)
        log.Printf("Message: %s\n", instanaErr.Message)
    }
    
    // Check if error is retryable
    if instana.IsRetryableError(err) {
        log.Println("This error can be retried")
    }
}
```

---

## 🔧 Configuration Defaults

The library provides sensible defaults for all configuration options:

```go
cfg := config.DefaultClientConfig()
// Defaults:
// - Connection Timeout: 30s
// - Request Timeout: 60s
// - Idle Connection Timeout: 90s
// - Max Retry Attempts: 3
// - Retry Initial Delay: 1s
// - Retry Max Delay: 30s
// - Rate Limit: 100 req/s
// - Burst Capacity: 200
// - Max Idle Connections: 100
// - Max Connections Per Host: 10
// - Keep-Alive Duration: 30s
// - Batch Size: 100
// - Batch Concurrent Requests: 5
```

You can override any of these using the builder pattern or by modifying the config directly.

---

## 📝 Complete Examples

### Example 1: List All Applications

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    
    "github.com/instana/instana-go-client/instana"
)

type Application struct {
    ID    string `json:"id"`
    Label string `json:"label"`
}

func main() {
    config, _ := instana.NewConfigBuilder().
        WithBaseURL("https://your-tenant.instana.io").
        WithAPIToken("your-api-token").
        Build()
    
    client, _ := instana.NewClientWithConfig(config)
    
    data, err := client.Get("/api/application-monitoring/applications")
    if err != nil {
        log.Fatal(err)
    }
    
    var apps []Application
    if err := json.Unmarshal(data, &apps); err != nil {
        log.Fatal(err)
    }
    
    for _, app := range apps {
        fmt.Printf("ID: %s, Label: %s\n", app.ID, app.Label)
    }
}
```

### Example 2: Create Application with Retry

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
    
    "github.com/instana/instana-go-client/instana"
)

type Application struct {
    Label            string `json:"label"`
    MatchSpecification struct {
        Type string `json:"type"`
    } `json:"matchSpecification"`
}

func main() {
    config, _ := instana.NewConfigBuilder().
        WithBaseURL("https://your-tenant.instana.io").
        WithAPIToken("your-api-token").
        WithMaxRetryAttempts(5).
        WithRetryInitialDelay(1 * time.Second).
        Build()
    
    client, _ := instana.NewClientWithConfig(config)
    
    app := Application{
        Label: "My New Application",
    }
    app.MatchSpecification.Type = "BINARY"
    
    data, err := client.Post(app, "/api/application-monitoring/applications")
    if err != nil {
        log.Fatal(err)
    }
    
    var createdApp Application
    if err := json.Unmarshal(data, &createdApp); err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Created application: %s\n", createdApp.Label)
}
```

### Example 3: Rate-Limited Batch Processing

```go
package main

import (
    "fmt"
    "log"
    "time"
    
    "github.com/instana/instana-go-client/instana"
)

func main() {
    config, _ := instana.NewConfigBuilder().
        WithBaseURL("https://your-tenant.instana.io").
        WithAPIToken("your-api-token").
        WithRateLimitEnabled(true).
        WithRateLimitRequestsPerSecond(10).
        Build()
    
    client, _ := instana.NewClientWithConfig(config)
    
    appIDs := []string{"app1", "app2", "app3", /* ... */}
    
    for _, appID := range appIDs {
        // Rate limiter automatically enforces limits
        data, err := client.GetOne(appID, "/api/application-monitoring/applications")
        if err != nil {
            log.Printf("Error fetching app %s: %v\n", appID, err)
            continue
        }
        fmt.Printf("Fetched app %s: %d bytes\n", appID, len(data))
    }
}
```

### Example 4: Production-Ready Configuration

```go
package main

import (
    "log"
    "time"
    
    "github.com/instana/instana-go-client/config"
    "github.com/instana/instana-go-client/instana"
)

func main() {
    // Create production-ready configuration
    cfg, err := config.NewConfigBuilder().
        WithBaseURL("https://your-tenant.instana.io").
        WithAPIToken("your-api-token").
        WithConnectionTimeout(10 * time.Second).
        WithRequestTimeout(30 * time.Second).
        WithMaxRetryAttempts(5).
        WithRetryInitialDelay(1 * time.Second).
        WithRetryMaxDelay(30 * time.Second).
        WithRateLimitEnabled(true).
        WithRateLimitRequestsPerSecond(10).
        WithRateLimitBurstCapacity(20).
        WithMaxIdleConnections(100).
        WithMaxConnectionsPerHost(10).
        WithCustomHeader("X-Application", "my-app").
        WithDebug(false).
        Build()
    
    if err != nil {
        log.Fatal(err)
    }
    
    api, err := instana.NewInstanaAPIWithConfig(cfg)
    if err != nil {
        log.Fatal(err)
    }
    
    // Use API...
    tokens, err := api.APITokens().GetAll()
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Retrieved %d tokens\n", len(*tokens))
}
```

---

## 🧪 Testing

### Mock Client for Testing

```go
package mypackage_test

import (
    "testing"
    
    "github.com/instana/instana-go-client/instana"
)

func TestMyFunction(t *testing.T) {
    // Create test configuration
    config := instana.DefaultClientConfig()
    config.BaseURL = "http://localhost:8080"  // Test server
    config.APIToken = "test-token"
    
    client, _ := instana.NewClientWithConfig(config)
    
    // Test your code...
}
```

---

## 🔧 Advanced Features

### Custom Logger

```go
type MyLogger struct{}

func (l *MyLogger) Debug(msg string, keysAndValues ...interface{}) {
    // Custom debug logging
}

func (l *MyLogger) Info(msg string, keysAndValues ...interface{}) {
    // Custom info logging
}

func (l *MyLogger) Error(msg string, keysAndValues ...interface{}) {
    // Custom error logging
}

config := instana.DefaultClientConfig()
config.Logger = &MyLogger{}
```

### Custom HTTP Client

```go
import (
    "crypto/tls"
    "net/http"
    "time"
)

httpClient := &http.Client{
    Timeout: 60 * time.Second,
    Transport: &http.Transport{
        TLSClientConfig: &tls.Config{
            InsecureSkipVerify: true,
        },
    },
}

config := instana.DefaultClientConfig()
config.HTTPClient = httpClient
```

---

## 📚 Additional Resources

- [Configuration Guide](./DEFAULT_CONFIG_ANALYSIS.md) - Detailed configuration options
- [Builder Pattern Guide](./BUILDER_PATTERN_ANALYSIS.md) - Builder pattern usage
- [Testing Summary](./TESTING_SUMMARY.md) - Test coverage and examples
- [Project Status](./PROJECT_STATUS.md) - Current project status
- [Implementation Plan](./IMPLEMENTATION_PLAN.md) - Development roadmap

---

## 🆘 Troubleshooting

### Common Issues

**Issue**: Connection timeout
```go
// Solution: Increase timeout
config, _ := instana.NewConfigBuilder().
    WithConnectionTimeout(60 * time.Second).
    Build()
```

**Issue**: Rate limit exceeded
```go
// Solution: Enable rate limiting
config, _ := instana.NewConfigBuilder().
    WithRateLimitEnabled(true).
    WithRateLimitRequestsPerSecond(50).
    Build()
```

**Issue**: Retry exhausted
```go
// Solution: Increase retry attempts or delays
config, _ := instana.NewConfigBuilder().
    WithMaxRetryAttempts(10).
    WithRetryMaxDelay(60 * time.Second).
    Build()
```

**Issue**: TLS verification errors
```go
// Solution: Skip TLS verification (not recommended for production)
client := instana.NewClient("api-token", "tenant.instana.io", true)
```

---

## 📞 Support

- **Issues**: [GitHub Issues](https://github.com/instana/instana-go-client/issues)
- **Documentation**: [Full Documentation](./README.md)
- **Examples**: [Examples Directory](./examples/)

---

**Made with ❤️ by the Instana Team**