# Instana Go Client - Usage Guide

## Current Status

The configuration system is **fully implemented** but **not yet integrated** with the REST client. This guide shows:
1. **Current Usage** - How to use the existing client (as-is)
2. **Future Usage** - How to use the new configuration system (after integration)

---

## Current Usage (Existing Client)

### How It Works Now

The existing `NewClient()` function uses simple parameters:

```go
package main

import (
    "github.com/instana/instana-go-client/instana"
)

func main() {
    // Current way - simple parameters
    client := instana.NewClient(
        "your-api-token",           // API token
        "https://tenant.instana.io", // Host URL
        false,                       // Skip TLS verification
    )
    
    // Use the client for API calls
    // data, err := client.Get("/api/application-monitoring/applications")
}
```

### Limitations of Current Approach
- ❌ Hardcoded timeout (30 seconds)
- ❌ Hardcoded throttle rate (5 requests/second)
- ❌ No retry logic
- ❌ No rate limiting configuration
- ❌ No custom headers support
- ❌ No connection pooling configuration
- ❌ Basic error messages (not typed)

---

## Future Usage (After Integration)

### How It Will Work

Once we integrate the configuration system (Phase 3 completion), you'll be able to use the new `NewClientWithConfig()` function:

```go
package main

import (
    "log"
    "time"
    
    "github.com/instana/instana-go-client/instana"
)

func main() {
    // Method 1: Using Builder Pattern (Recommended)
    config, err := instana.NewConfigBuilder().
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
    
    // Create client with configuration
    client, err := instana.NewClientWithConfig(config)
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }
    
    // Use the client - now with retry, rate limiting, etc.
    data, err := client.Get("/api/application-monitoring/applications")
    if err != nil {
        log.Fatalf("API call failed: %v", err)
    }
}
```

### Method 2: Environment Variables

```bash
# Set environment variables
export INSTANA_BASE_URL=https://tenant-unit.instana.io
export INSTANA_API_TOKEN=your-api-token
export INSTANA_CONNECTION_TIMEOUT=30s
export INSTANA_MAX_RETRY_ATTEMPTS=3
export INSTANA_RATE_LIMIT_RPS=100
```

```go
package main

import (
    "log"
    
    "github.com/instana/instana-go-client/instana"
)

func main() {
    // Load configuration from environment
    config, err := instana.LoadFromEnv()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }
    
    // Create client
    client, err := instana.NewClientWithConfig(config)
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }
    
    // Use the client
    data, err := client.Get("/api/application-monitoring/applications")
}
```

### Method 3: JSON Configuration File

**config.json:**
```json
{
  "baseURL": "https://tenant-unit.instana.io",
  "apiToken": "your-api-token",
  "timeout": {
    "connection": "30s",
    "request": "60s",
    "idleConnection": "90s"
  },
  "retry": {
    "maxAttempts": 3,
    "initialDelay": "1s",
    "maxDelay": "30s",
    "backoffMultiplier": 2.0,
    "retryableStatusCodes": [408, 429, 500, 502, 503, 504],
    "retryOnTimeout": true,
    "retryOnConnectionError": true,
    "jitter": true
  },
  "rateLimit": {
    "enabled": true,
    "requestsPerSecond": 100,
    "burstCapacity": 200,
    "waitForToken": true
  },
  "batch": {
    "size": 100,
    "concurrentRequests": 5
  },
  "connectionPool": {
    "maxIdleConnections": 100,
    "maxConnectionsPerHost": 10,
    "keepAliveDuration": "30s"
  },
  "userAgent": "my-app/1.0.0",
  "debug": false
}
```

```go
package main

import (
    "log"
    
    "github.com/instana/instana-go-client/instana"
)

func main() {
    // Load configuration from JSON file
    config, err := instana.LoadFromJSON("config.json")
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }
    
    // Create client
    client, err := instana.NewClientWithConfig(config)
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }
    
    // Use the client
    data, err := client.Get("/api/application-monitoring/applications")
}
```

### Method 4: Hybrid (JSON + Environment Override)

```go
package main

import (
    "log"
    
    "github.com/instana/instana-go-client/instana"
)

func main() {
    // Load from JSON, override with environment variables
    config, err := instana.LoadFromJSONWithEnvOverride("config.json")
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }
    
    // Create client
    client, err := instana.NewClientWithConfig(config)
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }
    
    // Use the client
    data, err := client.Get("/api/application-monitoring/applications")
}
```

---

## Benefits of New Configuration System

### 1. Automatic Retry
```go
// Automatically retries on transient failures
data, err := client.Get("/api/applications")
// If it fails with 503, it will automatically retry with exponential backoff
```

### 2. Rate Limiting
```go
// Automatically respects rate limits
for i := 0; i < 1000; i++ {
    data, err := client.Get("/api/applications")
    // Rate limiter ensures we don't exceed configured RPS
}
```

### 3. Typed Errors
```go
data, err := client.Get("/api/applications")
if err != nil {
    // Check if error is retryable
    if instana.IsRetryableError(err) {
        // Handle retryable error
    }
    
    // Extract status code
    statusCode := instana.ExtractStatusCode(err)
    
    // Type assertion for detailed info
    if instanaErr, ok := err.(*instana.InstanaError); ok {
        log.Printf("Error Type: %s", instanaErr.Type)
        log.Printf("Retryable: %v", instanaErr.IsRetryable())
    }
}
```

### 4. Custom Headers
```go
config := instana.NewConfigBuilder().
    WithBaseURL("https://tenant.instana.io").
    WithAPIToken("token").
    WithCustomHeader("X-Request-ID", "unique-id").
    WithCustomHeader("X-Trace-ID", "trace-123").
    Build()

client, _ := instana.NewClientWithConfig(config)
// All requests will include custom headers
```

### 5. Connection Pooling
```go
config := instana.NewConfigBuilder().
    WithBaseURL("https://tenant.instana.io").
    WithAPIToken("token").
    WithMaxIdleConnections(100).
    WithMaxConnectionsPerHost(10).
    WithKeepAliveDuration(30 * time.Second).
    Build()

client, _ := instana.NewClientWithConfig(config)
// HTTP connections are pooled and reused efficiently
```

---

## Migration Path

### Step 1: Current Code (No Changes Needed)
```go
// This will continue to work
client := instana.NewClient("token", "https://tenant.instana.io", false)
```

### Step 2: After Integration (Opt-in to New Features)
```go
// Use new configuration for enhanced features
config := instana.NewConfigBuilder().
    WithBaseURL("https://tenant.instana.io").
    WithAPIToken("token").
    WithMaxRetryAttempts(5).
    Build()

client, _ := instana.NewClientWithConfig(config)
```

### Step 3: Deprecation (Future)
```go
// Old NewClient() will be marked as deprecated
// Recommended to migrate to NewClientWithConfig()
```

---

## Implementation Status

### ✅ Completed
- Configuration structures (`ClientConfig`, `TimeoutConfig`, etc.)
- Configuration validation
- Builder pattern (`NewConfigBuilder()`)
- Environment variable loading (`LoadFromEnv()`)
- JSON file loading (`LoadFromJSON()`)
- Retry mechanism (`Retryer`)
- Rate limiter (`RateLimiter`)
- Typed errors (`InstanaError`)
- Logging infrastructure (`Logger`)

### 🔄 In Progress (Phase 3 - 40% Remaining)
- Refactor `NewClient()` to `NewClientWithConfig()`
- Integrate retry mechanism into REST client
- Integrate rate limiter into REST client
- Add custom headers support to HTTP requests
- Configure HTTP transport with connection pooling
- Update error handling to use typed errors

### ⏳ Pending
- Comprehensive testing
- CI/CD pipeline
- Additional examples
- Migration guide

---

## Timeline

- **Current**: Configuration system ready, not yet integrated
- **Week 4**: Complete REST client integration
- **Week 5-6**: Testing and validation
- **Week 7+**: CI/CD, documentation, release

---

## Questions?

For more information:
- See [IMPLEMENTATION_PLAN.md](IMPLEMENTATION_PLAN.md) for detailed roadmap
- See [PROGRESS_SUMMARY.md](PROGRESS_SUMMARY.md) for current status
- See [README.md](README.md) for complete documentation
- See [examples/](examples/) for working code examples

---

## Summary

**Current State:**
- ✅ Configuration system is fully implemented and tested
- ✅ All configuration methods work (builder, env, JSON)
- ✅ Retry, rate limiting, and error handling are ready
- ❌ Not yet integrated with REST client

**Next Step:**
- Refactor REST client to accept `ClientConfig`
- This will enable all the new features

**Your Action:**
- Continue using `NewClient()` for now
- Watch for `NewClientWithConfig()` in the next update
- Start planning your configuration (JSON file or env vars)