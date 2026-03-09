# ConfigBuilder Pattern - Default Values Analysis

## Question
What happens to configuration parameters when using the builder pattern with only partial configuration?

```go
config, err := instana.NewConfigBuilder().
    WithBaseURL("https://tenant.instana.io").
    WithAPIToken("your-token").
    WithMaxRetryAttempts(5).
    Build()
```

## Answer: All Unspecified Parameters Get Default Values ✅

## How It Works

### Step 1: Builder Initialization
```go
func NewConfigBuilder() *ConfigBuilder {
    return &ConfigBuilder{
        config: DefaultClientConfig(),  // ← Starts with ALL defaults
    }
}
```

**Key Point**: The builder is initialized with a **complete** default configuration, not an empty one.

### Step 2: Override Specific Values
```go
// Only these 3 values are changed from defaults:
WithBaseURL("https://tenant.instana.io")    // Overrides BaseURL
WithAPIToken("your-token")                   // Overrides APIToken
WithMaxRetryAttempts(5)                      // Overrides Retry.MaxAttempts
```

### Step 3: Build Returns Complete Config
```go
Build()  // Returns config with:
         // - Your 3 custom values
         // - ALL other values from DefaultClientConfig()
```

## Complete Configuration Result

When you use the example above, here's the **complete** configuration you get:

### ✅ Your Custom Values (3 parameters)
```go
BaseURL:                "https://tenant.instana.io"  // ← You set this
APIToken:               "your-token"                  // ← You set this
Retry.MaxAttempts:      5                            // ← You set this
```

### ✅ Default Values (All Other Parameters)

#### Timeouts (5 parameters)
```go
Timeout.Connection:     30 * time.Second   // Default
Timeout.Request:        60 * time.Second   // Default
Timeout.IdleConnection: 90 * time.Second   // Default
Timeout.ResponseHeader: 10 * time.Second   // Default
Timeout.TLSHandshake:   10 * time.Second   // Default
```

#### Retry Configuration (7 parameters - 1 custom, 6 default)
```go
Retry.MaxAttempts:            5                                    // ← Custom
Retry.InitialDelay:           1 * time.Second                     // Default
Retry.MaxDelay:               30 * time.Second                    // Default
Retry.BackoffMultiplier:      2.0                                 // Default
Retry.RetryableStatusCodes:   []int{408, 429, 500, 502, 503, 504} // Default
Retry.RetryOnTimeout:         true                                // Default
Retry.RetryOnConnectionError: true                                // Default
Retry.Jitter:                 true                                // Default
```

#### Rate Limiting (4 parameters)
```go
RateLimit.RequestsPerSecond: 100   // Default
RateLimit.BurstCapacity:     200   // Default
RateLimit.Enabled:           true  // Default
RateLimit.WaitForToken:      true  // Default
```

#### Connection Pool (6 parameters)
```go
ConnectionPool.MaxIdleConnections:        100              // Default
ConnectionPool.MaxConnectionsPerHost:     10               // Default
ConnectionPool.MaxIdleConnectionsPerHost: 10               // Default
ConnectionPool.KeepAliveDuration:         30 * time.Second // Default
ConnectionPool.DisableKeepAlives:         false            // Default
ConnectionPool.DisableCompression:        false            // Default
```

#### Headers (2 parameters)
```go
Headers.Custom:                make(map[string]string) // Default (empty map)
Headers.DisableDefaultHeaders: false                   // Default
```

#### Batch Configuration (4 parameters)
```go
Batch.Size:               100   // Default
Batch.ConcurrentRequests: 5     // Default
Batch.StopOnError:        false // Default
Batch.RetryFailedItems:   true  // Default
```

#### Other Settings (4 parameters)
```go
UserAgent:  "instana-go-client/v1.0.0" // Default
Debug:      false                       // Default
Logger:     nil                         // Default (will be set by NewClientWithConfig)
HTTPClient: nil                         // Default (will be created by NewClientWithConfig)
```

## Total Configuration Summary

| Category | Custom | Default | Total |
|----------|--------|---------|-------|
| **Basic** | 2 | 2 | 4 |
| **Timeouts** | 0 | 5 | 5 |
| **Retry** | 1 | 7 | 8 |
| **Rate Limit** | 0 | 4 | 4 |
| **Connection Pool** | 0 | 6 | 6 |
| **Headers** | 0 | 2 | 2 |
| **Batch** | 0 | 4 | 4 |
| **Other** | 0 | 4 | 4 |
| **TOTAL** | **3** | **34** | **37** |

**Result**: You specify 3 parameters, get 37 complete parameters (3 custom + 34 defaults)

## Visual Representation

```
NewConfigBuilder()
    ↓
[Complete Default Config]
    ├─ BaseURL: ""
    ├─ APIToken: ""
    ├─ Timeout.Connection: 30s
    ├─ Timeout.Request: 60s
    ├─ Retry.MaxAttempts: 3        ← Will be overridden
    ├─ Retry.InitialDelay: 1s
    ├─ RateLimit.Enabled: true
    ├─ ... (all 37 parameters)
    ↓
WithBaseURL("https://tenant.instana.io")
    ↓
[Config with BaseURL updated]
    ├─ BaseURL: "https://tenant.instana.io"  ← Changed
    ├─ APIToken: ""
    ├─ Retry.MaxAttempts: 3
    ├─ ... (all others unchanged)
    ↓
WithAPIToken("your-token")
    ↓
[Config with APIToken updated]
    ├─ BaseURL: "https://tenant.instana.io"
    ├─ APIToken: "your-token"                ← Changed
    ├─ Retry.MaxAttempts: 3
    ├─ ... (all others unchanged)
    ↓
WithMaxRetryAttempts(5)
    ↓
[Config with MaxRetryAttempts updated]
    ├─ BaseURL: "https://tenant.instana.io"
    ├─ APIToken: "your-token"
    ├─ Retry.MaxAttempts: 5                  ← Changed
    ├─ Retry.InitialDelay: 1s                ← Still default
    ├─ RateLimit.Enabled: true               ← Still default
    ├─ ... (all others still default)
    ↓
Build()
    ↓
[Complete, Valid Configuration]
    ✅ 3 custom values
    ✅ 34 default values
    ✅ Ready to use!
```

## Code Examples

### Example 1: Minimal Configuration
```go
// Only set required fields
config, err := instana.NewConfigBuilder().
    WithBaseURL("https://tenant.instana.io").
    WithAPIToken("your-token").
    Build()

// Result: You get a complete config with:
// - Your BaseURL and APIToken
// - 60s request timeout (default)
// - 3 retry attempts (default)
// - 100 req/s rate limit (default)
// - Connection pooling (default)
// - All other defaults
```

### Example 2: Partial Customization
```go
// Customize only what you need
config, err := instana.NewConfigBuilder().
    WithBaseURL("https://tenant.instana.io").
    WithAPIToken("your-token").
    WithMaxRetryAttempts(5).              // Custom
    WithRequestTimeout(120 * time.Second). // Custom
    Build()

// Result: You get:
// - Your 4 custom values
// - 33 default values
// - Complete, working configuration
```

### Example 3: Verify Defaults Are Applied
```go
config, err := instana.NewConfigBuilder().
    WithBaseURL("https://tenant.instana.io").
    WithAPIToken("your-token").
    Build()

// Check that defaults are present:
fmt.Printf("Request Timeout: %v\n", config.Timeout.Request)
// Output: Request Timeout: 1m0s (default)

fmt.Printf("Max Retry Attempts: %d\n", config.Retry.MaxAttempts)
// Output: Max Retry Attempts: 3 (default)

fmt.Printf("Rate Limit Enabled: %v\n", config.RateLimit.Enabled)
// Output: Rate Limit Enabled: true (default)

fmt.Printf("Max Idle Connections: %d\n", config.ConnectionPool.MaxIdleConnections)
// Output: Max Idle Connections: 100 (default)
```

## Comparison: Different Configuration Methods

### Method 1: Builder with Minimal Config
```go
config, _ := instana.NewConfigBuilder().
    WithBaseURL("https://tenant.instana.io").
    WithAPIToken("your-token").
    Build()
// Result: 2 custom + 35 defaults = 37 total parameters
```

### Method 2: Builder with Partial Config
```go
config, _ := instana.NewConfigBuilder().
    WithBaseURL("https://tenant.instana.io").
    WithAPIToken("your-token").
    WithMaxRetryAttempts(5).
    WithRequestTimeout(120 * time.Second).
    WithRateLimitRequestsPerSecond(50).
    Build()
// Result: 5 custom + 32 defaults = 37 total parameters
```

### Method 3: Builder with Full Config
```go
config, _ := instana.NewConfigBuilder().
    WithBaseURL("https://tenant.instana.io").
    WithAPIToken("your-token").
    WithConnectionTimeout(10 * time.Second).
    WithRequestTimeout(120 * time.Second).
    WithMaxRetryAttempts(5).
    WithRetryInitialDelay(2 * time.Second).
    WithRateLimitRequestsPerSecond(50).
    WithMaxIdleConnections(200).
    // ... all 37 parameters
    Build()
// Result: 37 custom + 0 defaults = 37 total parameters
```

### Method 4: Default Config Only
```go
config := instana.DefaultClientConfig()
config.BaseURL = "https://tenant.instana.io"
config.APIToken = "your-token"
// Result: 2 custom + 35 defaults = 37 total parameters
```

**All methods produce a complete, valid configuration!**

## Benefits of This Design

### 1. ✅ Safe Defaults
You never get an incomplete or invalid configuration. Every parameter has a sensible default value.

### 2. ✅ Minimal Configuration
You only need to specify what's different from defaults:
```go
// Minimum viable config (2 parameters)
config, _ := instana.NewConfigBuilder().
    WithBaseURL("https://tenant.instana.io").
    WithAPIToken("your-token").
    Build()
```

### 3. ✅ Progressive Enhancement
Start simple, add complexity as needed:
```go
// Start simple
config, _ := instana.NewConfigBuilder().
    WithBaseURL(url).
    WithAPIToken(token).
    Build()

// Later, add more customization
config, _ := instana.NewConfigBuilder().
    WithBaseURL(url).
    WithAPIToken(token).
    WithMaxRetryAttempts(5).        // Add retry config
    WithRequestTimeout(120*time.Second). // Add timeout config
    Build()
```

### 4. ✅ No Surprises
Unspecified parameters don't become zero/nil/empty. They get production-ready defaults:
```go
// You don't specify rate limit
config, _ := instana.NewConfigBuilder().
    WithBaseURL(url).
    WithAPIToken(token).
    Build()

// But you still get rate limiting!
// config.RateLimit.Enabled = true (default)
// config.RateLimit.RequestsPerSecond = 100 (default)
```

### 5. ✅ Validation Included
The `Build()` method validates the complete configuration:
```go
config, err := instana.NewConfigBuilder().
    WithBaseURL("").  // Invalid: empty URL
    WithAPIToken("token").
    Build()

if err != nil {
    // err contains validation errors
    fmt.Println(err) // "BaseURL is required"
}
```

## Common Patterns

### Pattern 1: Environment-Specific Overrides
```go
// Base config with defaults
builder := instana.NewConfigBuilder().
    WithBaseURL(os.Getenv("INSTANA_URL")).
    WithAPIToken(os.Getenv("INSTANA_TOKEN"))

// Production: Add stricter settings
if env == "production" {
    builder.WithMaxRetryAttempts(5).
           WithRequestTimeout(120 * time.Second)
}

// Development: Add debug logging
if env == "development" {
    builder.WithDebug(true)
}

config, _ := builder.Build()
```

### Pattern 2: Incremental Configuration
```go
// Start with defaults
builder := instana.NewConfigBuilder().
    WithBaseURL(url).
    WithAPIToken(token)

// Add features as needed
if needsHighThroughput {
    builder.WithRateLimitRequestsPerSecond(200).
           WithMaxIdleConnections(200)
}

if needsReliability {
    builder.WithMaxRetryAttempts(5).
           WithRetryMaxDelay(60 * time.Second)
}

config, _ := builder.Build()
```

### Pattern 3: Template Configuration
```go
// Create a template with common settings
template := instana.NewConfigBuilder().
    WithMaxRetryAttempts(5).
    WithRequestTimeout(120 * time.Second).
    WithRateLimitRequestsPerSecond(50)

// Create specific configs from template
prodConfig, _ := instana.NewConfigBuilderFromConfig(template.GetConfig()).
    WithBaseURL("https://prod.instana.io").
    WithAPIToken(prodToken).
    Build()

devConfig, _ := instana.NewConfigBuilderFromConfig(template.GetConfig()).
    WithBaseURL("https://dev.instana.io").
    WithAPIToken(devToken).
    WithDebug(true).
    Build()
```

## Summary

### Question
```go
config, err := instana.NewConfigBuilder().
    WithBaseURL("https://tenant.instana.io").
    WithAPIToken("your-token").
    WithMaxRetryAttempts(5).
    Build()
```
**What happens to the rest of the config parameters?**

### Answer
✅ **All unspecified parameters automatically get production-ready default values**

- You specify: 3 parameters
- You get: 37 complete parameters (3 custom + 34 defaults)
- Result: Complete, valid, production-ready configuration

### Key Takeaway
The builder pattern ensures you **always** get a complete configuration. You only specify what's different from defaults, and everything else is filled in with sensible, production-ready values. No surprises, no missing configuration, no invalid states! 🎯