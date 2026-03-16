# Instana Go Client - Initialization Methods

## Current vs New Initialization

### Current Structure (instana/Instana-api.go)

```go
// Current - Simple initialization
func NewInstanaAPI(apiToken string, endpoint string, skipTlsVerification bool) InstanaAPI

// Current - With user agent
func NewInstanaAPIWithUserAgent(apiToken string, endpoint string, skipTlsVerification bool, userAgent string) InstanaAPI
```

### New Structure (instana/client.go)

The new structure will have **multiple initialization methods** in `instana/client.go`:

```go
package instana

import (
    "crypto/tls"
    "net/http"
    
    "github.com/instana/instana-go-client/client"
    "github.com/instana/instana-go-client/config"
)

// NewInstanaAPI creates a new instance of the Instana API (backward compatible)
// This maintains the same signature as the current implementation
func NewInstanaAPI(apiToken string, endpoint string, skipTlsVerification bool) InstanaAPI {
    // Create basic configuration
    cfg := config.DefaultClientConfig()
    cfg.APIToken = apiToken
    cfg.BaseURL = "https://" + endpoint
    
    // Handle TLS verification
    if skipTlsVerification {
        cfg.TLS.InsecureSkipVerify = true
    }
    
    // Create REST client
    restClient, err := client.NewClient(cfg)
    if err != nil {
        // Fallback to basic client
        cfg.Logger.Error("Failed to create client", "error", err)
        return nil
    }
    
    return &instanaAPI{
        restClient: restClient,
    }
}

// NewInstanaAPIWithUserAgent creates a new instance with custom user agent (backward compatible)
// The userAgent parameter should include the client name and version (e.g., "Terraform/1.2.3")
func NewInstanaAPIWithUserAgent(apiToken string, endpoint string, skipTlsVerification bool, userAgent string) InstanaAPI {
    // Create configuration with user agent
    cfg := config.DefaultClientConfig()
    cfg.APIToken = apiToken
    cfg.BaseURL = "https://" + endpoint
    cfg.UserAgent = userAgent
    cfg.Logger = config.NewDefaultLogger(config.ClientLogLevelInfo)
    
    // Handle TLS verification
    if skipTlsVerification {
        cfg.TLS.InsecureSkipVerify = true
    }
    
    // Create REST client
    restClient, err := client.NewClient(cfg)
    if err != nil {
        cfg.Logger.Error("Failed to create client with config", "error", err)
        // Fallback to basic initialization
        return NewInstanaAPI(apiToken, endpoint, skipTlsVerification)
    }
    
    return &instanaAPI{
        restClient: restClient,
    }
}

// NewAPI creates a new Instana API instance with full configuration (new method)
// This is the recommended method for new code
func NewAPI(cfg *config.Config) (InstanaAPI, error) {
    // Validate configuration
    if err := cfg.Validate(); err != nil {
        return nil, err
    }
    
    // Create REST client
    restClient, err := client.NewClient(cfg)
    if err != nil {
        return nil, err
    }
    
    return &instanaAPI{
        restClient: restClient,
    }, nil
}

// NewAPIWithConfig is an alias for NewAPI (for clarity)
func NewAPIWithConfig(cfg *config.Config) (InstanaAPI, error) {
    return NewAPI(cfg)
}
```

## File Location

All initialization methods will be in: **`instana/client.go`**

```
instana-go-client/
└── instana/
    ├── api.go          # InstanaAPI interface definition
    ├── client.go       # Initialization methods (NewInstanaAPI, NewAPI, etc.)
    └── doc.go         # Package documentation
```

## Usage Examples

### Example 1: Current Method (Backward Compatible)

```go
package main

import (
    "github.com/instana/instana-go-client/instana"
)

func main() {
    // Same as current implementation - no changes needed
    api := instana.NewInstanaAPI(
        "your-api-token",
        "tenant.instana.io",
        false, // skipTlsVerification
    )
    
    // Use the API
    alertClient := api.ApplicationAlertConfig()
    // ...
}
```

### Example 2: With User Agent (Backward Compatible)

```go
package main

import (
    "github.com/instana/instana-go-client/instana"
)

func main() {
    // Same as current implementation - no changes needed
    api := instana.NewInstanaAPIWithUserAgent(
        "your-api-token",
        "tenant.instana.io",
        false,              // skipTlsVerification
        "Terraform/1.2.3",  // userAgent
    )
    
    // Use the API
    alertClient := api.ApplicationAlertConfig()
    // ...
}
```

### Example 3: New Method with Full Configuration (Recommended)

```go
package main

import (
    "time"
    
    "github.com/instana/instana-go-client/config"
    "github.com/instana/instana-go-client/instana"
)

func main() {
    // Create configuration using builder
    cfg, err := config.NewBuilder().
        WithBaseURL("https://tenant.instana.io").
        WithAPIToken("your-api-token").
        WithUserAgent("Terraform/1.2.3").
        WithConnectionTimeout(30 * time.Second).
        WithMaxRetryAttempts(3).
        WithRateLimitRequestsPerSecond(100).
        Build()
    if err != nil {
        panic(err)
    }
    
    // Create API with full configuration
    api, err := instana.NewAPI(cfg)
    if err != nil {
        panic(err)
    }
    
    // Use the API
    alertClient := api.ApplicationAlertConfig()
    // ...
}
```

### Example 4: Environment Variables

```go
package main

import (
    "github.com/instana/instana-go-client/config"
    "github.com/instana/instana-go-client/instana"
)

func main() {
    // Load configuration from environment variables
    // INSTANA_BASE_URL, INSTANA_API_TOKEN, etc.
    cfg, err := config.LoadFromEnv()
    if err != nil {
        panic(err)
    }
    
    // Create API
    api, err := instana.NewAPI(cfg)
    if err != nil {
        panic(err)
    }
    
    // Use the API
    alertClient := api.ApplicationAlertConfig()
    // ...
}
```

## Complete Implementation

### instana/client.go (Full File)

```go
package instana

import (
    "github.com/instana/instana-go-client/client"
    "github.com/instana/instana-go-client/config"
)

// NewInstanaAPI creates a new instance of the Instana API
// This method maintains backward compatibility with the existing API
//
// Parameters:
//   - apiToken: Your Instana API token
//   - endpoint: Instana endpoint (e.g., "tenant.instana.io")
//   - skipTlsVerification: Whether to skip TLS certificate verification
//
// Example:
//   api := instana.NewInstanaAPI("token", "tenant.instana.io", false)
func NewInstanaAPI(apiToken string, endpoint string, skipTlsVerification bool) InstanaAPI {
    cfg := config.DefaultClientConfig()
    cfg.APIToken = apiToken
    cfg.BaseURL = "https://" + endpoint
    
    if skipTlsVerification {
        cfg.TLS.InsecureSkipVerify = true
    }
    
    restClient, err := client.NewClient(cfg)
    if err != nil {
        cfg.Logger.Error("Failed to create client", "error", err)
        return nil
    }
    
    return &instanaAPI{
        restClient: restClient,
    }
}

// NewInstanaAPIWithUserAgent creates a new instance of the Instana API with a custom user agent
// This method maintains backward compatibility with the existing API
//
// The userAgent parameter should include the client name and version (e.g., "Terraform/1.2.3")
//
// Parameters:
//   - apiToken: Your Instana API token
//   - endpoint: Instana endpoint (e.g., "tenant.instana.io")
//   - skipTlsVerification: Whether to skip TLS certificate verification
//   - userAgent: Custom user agent string
//
// Example:
//   api := instana.NewInstanaAPIWithUserAgent("token", "tenant.instana.io", false, "MyApp/1.0")
func NewInstanaAPIWithUserAgent(apiToken string, endpoint string, skipTlsVerification bool, userAgent string) InstanaAPI {
    cfg := config.DefaultClientConfig()
    cfg.APIToken = apiToken
    cfg.BaseURL = "https://" + endpoint
    cfg.UserAgent = userAgent
    cfg.Logger = config.NewDefaultLogger(config.ClientLogLevelInfo)
    
    if skipTlsVerification {
        cfg.TLS.InsecureSkipVerify = true
    }
    
    restClient, err := client.NewClient(cfg)
    if err != nil {
        cfg.Logger.Error("Failed to create client with config", "error", err)
        return NewInstanaAPI(apiToken, endpoint, skipTlsVerification)
    }
    
    return &instanaAPI{
        restClient: restClient,
    }
}

// NewAPI creates a new Instana API instance with full configuration
// This is the recommended method for new code as it provides access to all configuration options
//
// Parameters:
//   - cfg: Configuration object created via config.NewBuilder() or config.LoadFromEnv()
//
// Returns:
//   - InstanaAPI: The API instance
//   - error: Any error that occurred during initialization
//
// Example:
//   cfg, _ := config.NewBuilder().
//       WithBaseURL("https://tenant.instana.io").
//       WithAPIToken("token").
//       Build()
//   api, err := instana.NewAPI(cfg)
func NewAPI(cfg *config.Config) (InstanaAPI, error) {
    if err := cfg.Validate(); err != nil {
        return nil, err
    }
    
    restClient, err := client.NewClient(cfg)
    if err != nil {
        return nil, err
    }
    
    return &instanaAPI{
        restClient: restClient,
    }, nil
}

// NewAPIWithConfig is an alias for NewAPI for clarity
func NewAPIWithConfig(cfg *config.Config) (InstanaAPI, error) {
    return NewAPI(cfg)
}
```

## Migration Path

### For Existing Terraform Provider Code

**No changes required!** The existing initialization methods are maintained:

```go
// Current code - works as-is
api := instana.NewInstanaAPI(token, endpoint, skipTLS)

// Current code with user agent - works as-is
api := instana.NewInstanaAPIWithUserAgent(token, endpoint, skipTLS, userAgent)
```

### For New Code

Use the new configuration-based approach:

```go
// New recommended approach
cfg, _ := config.NewBuilder().
    WithBaseURL("https://tenant.instana.io").
    WithAPIToken("token").
    WithUserAgent("MyApp/1.0").
    Build()

api, err := instana.NewAPI(cfg)
```

## Summary

| Method | Location | Purpose | Status |
|--------|----------|---------|--------|
| `NewInstanaAPI()` | `instana/client.go` | Simple initialization | ✅ Backward compatible |
| `NewInstanaAPIWithUserAgent()` | `instana/client.go` | With user agent | ✅ Backward compatible |
| `NewAPI()` | `instana/client.go` | Full configuration | ✅ New recommended method |
| `NewAPIWithConfig()` | `instana/client.go` | Alias for NewAPI | ✅ New method |

**Key Points:**
1. ✅ All methods in `instana/client.go`
2. ✅ Existing methods maintained for backward compatibility
3. ✅ New methods provide more flexibility
4. ✅ No breaking changes for Terraform provider
5. ✅ Clear migration path for new features