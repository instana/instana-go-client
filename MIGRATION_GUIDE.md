# Migration Guide - Instana Go Client

**Version**: v1.0.0  
**Last Updated**: 2026-03-09

This guide helps you migrate from the embedded REST API client in the Terraform provider to the standalone Instana Go Client library.

---

## 📋 Table of Contents

1. [Overview](#overview)
2. [Breaking Changes](#breaking-changes)
3. [Migration Steps](#migration-steps)
4. [Code Examples](#code-examples)
5. [Configuration Changes](#configuration-changes)
6. [API Changes](#api-changes)
7. [Error Handling](#error-handling)
8. [Testing Changes](#testing-changes)
9. [Troubleshooting](#troubleshooting)

---

## Overview

### What Changed?

The Instana REST API client has been extracted from the Terraform provider into a standalone library with significant enhancements:

**Before** (Embedded in Terraform Provider):
- Basic REST client functionality
- Limited configuration options
- Simple error handling
- No retry mechanism
- No rate limiting

**After** (Standalone Library):
- Enhanced REST client with 37 configuration parameters
- Comprehensive error handling with typed errors
- Automatic retry with exponential backoff
- Built-in rate limiting
- Connection pooling
- Environment variable support
- JSON configuration file support
- Builder pattern for easy configuration

### Why Migrate?

✅ **Better Performance**: Connection pooling and rate limiting  
✅ **Reliability**: Automatic retries with exponential backoff  
✅ **Flexibility**: 37 configurable parameters  
✅ **Developer Experience**: Builder pattern, environment variables, JSON config  
✅ **Error Handling**: Typed errors with detailed information  
✅ **Testing**: Comprehensive test coverage (133+ tests)  
✅ **Maintenance**: Separate versioning and release cycle

---

## Breaking Changes

### 1. Import Path Change

**Before:**
```go
import "github.com/instana/terraform-provider-instana/internal/restapi"
```

**After:**
```go
import "github.com/instana/instana-go-client/instana"
```

### 2. Client Creation

**Before:**
```go
client := restapi.NewClient(apiToken, host, skipTlsVerification)
```

**After (Legacy Compatible):**
```go
client := instana.NewClient(apiToken, host, skipTlsVerification)
```

**After (Recommended):**
```go
config, err := instana.NewConfigBuilder().
    WithBaseURL("https://" + host).
    WithAPIToken(apiToken).
    Build()
if err != nil {
    // handle error
}
client, err := instana.NewClientWithConfig(config)
```

### 3. Method Signatures

Most method signatures remain the same, but error handling is enhanced:

**Before:**
```go
data, err := client.Get("/api/endpoint")
// Simple error checking
if err != nil {
    return err
}
```

**After:**
```go
data, err := client.Get("/api/endpoint")
// Enhanced error checking with typed errors
if err != nil {
    var instanaErr *instana.InstanaError
    if errors.As(err, &instanaErr) {
        // Handle specific error types
        switch instanaErr.Type {
        case instana.ErrorTypeAuthentication:
            // Handle auth error
        case instana.ErrorTypeRateLimit:
            // Handle rate limit
        }
    }
    return err
}
```

### 4. Configuration Structure

**Before:**
```go
// Limited configuration options
client := restapi.NewClient(apiToken, host, skipTls)
```

**After:**
```go
// Rich configuration with 37 parameters
config := instana.DefaultClientConfig()
config.BaseURL = "https://" + host
config.APIToken = apiToken
config.Timeout.Connection = 30 * time.Second
config.Timeout.Request = 60 * time.Second
config.Retry.MaxAttempts = 5
config.RateLimit.Enabled = true
config.RateLimit.RequestsPerSecond = 100
// ... and 30 more parameters
```

---

## Migration Steps

### Step 1: Update Dependencies

**Update go.mod:**
```bash
# Remove old dependency (if explicitly listed)
go mod edit -dropreplace github.com/instana/terraform-provider-instana/internal/restapi

# Add new dependency
go get github.com/instana/instana-go-client@latest
```

### Step 2: Update Imports

**Find and replace in all files:**
```bash
# Using sed (Linux/macOS)
find . -type f -name "*.go" -exec sed -i 's|github.com/instana/terraform-provider-instana/internal/restapi|github.com/instana/instana-go-client/instana|g' {} +

# Using PowerShell (Windows)
Get-ChildItem -Recurse -Filter *.go | ForEach-Object {
    (Get-Content $_.FullName) -replace 'github.com/instana/terraform-provider-instana/internal/restapi', 'github.com/instana/instana-go-client/instana' | Set-Content $_.FullName
}
```

### Step 3: Update Client Creation

**Option A: Minimal Changes (Legacy Compatible)**
```go
// Old code - still works!
client := instana.NewClient(apiToken, host, skipTlsVerification)
```

**Option B: Use New Configuration System (Recommended)**
```go
config, err := instana.NewConfigBuilder().
    WithBaseURL("https://" + host).
    WithAPIToken(apiToken).
    WithConnectionTimeout(30 * time.Second).
    WithMaxRetryAttempts(5).
    Build()

if err != nil {
    return nil, fmt.Errorf("failed to create config: %w", err)
}

client, err := instana.NewClientWithConfig(config)
if err != nil {
    return nil, fmt.Errorf("failed to create client: %w", err)
}
```

### Step 4: Update Error Handling

**Enhance error handling to use typed errors:**
```go
data, err := client.Get("/api/endpoint")
if err != nil {
    // Check for specific error types
    var instanaErr *instana.InstanaError
    if errors.As(err, &instanaErr) {
        switch instanaErr.Type {
        case instana.ErrorTypeAuthentication:
            log.Printf("Authentication failed: %s", instanaErr.Message)
        case instana.ErrorTypeRateLimit:
            log.Printf("Rate limit exceeded, retry after: %v", instanaErr.RetryAfter)
        case instana.ErrorTypeNotFound:
            log.Printf("Resource not found: %s", instanaErr.Message)
        default:
            log.Printf("API error: %s (status: %d)", instanaErr.Message, instanaErr.StatusCode)
        }
    }
    return err
}
```

### Step 5: Run Tests

```bash
# Run tests to verify migration
go test ./...

# Run with race detection
go test -race ./...

# Check coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

---

## Code Examples

### Example 1: Basic Migration

**Before:**
```go
package main

import (
    "log"
    "github.com/instana/terraform-provider-instana/internal/restapi"
)

func main() {
    client := restapi.NewClient("api-token", "tenant.instana.io", false)
    
    data, err := client.Get("/api/application-monitoring/applications")
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Data: %s", string(data))
}
```

**After (Minimal Changes):**
```go
package main

import (
    "log"
    "github.com/instana/instana-go-client/instana"
)

func main() {
    client := instana.NewClient("api-token", "tenant.instana.io", false)
    
    data, err := client.Get("/api/application-monitoring/applications")
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Data: %s", string(data))
}
```

**After (With Enhancements):**
```go
package main

import (
    "log"
    "time"
    "github.com/instana/instana-go-client/instana"
)

func main() {
    config, err := instana.NewConfigBuilder().
        WithBaseURL("https://tenant.instana.io").
        WithAPIToken("api-token").
        WithConnectionTimeout(30 * time.Second).
        WithMaxRetryAttempts(3).
        WithRateLimitEnabled(true).
        WithRateLimitRequestsPerSecond(100).
        Build()
    
    if err != nil {
        log.Fatal(err)
    }
    
    client, err := instana.NewClientWithConfig(config)
    if err != nil {
        log.Fatal(err)
    }
    
    data, err := client.Get("/api/application-monitoring/applications")
    if err != nil {
        var instanaErr *instana.InstanaError
        if errors.As(err, &instanaErr) {
            log.Printf("API Error: %s (Type: %v, Status: %d)", 
                instanaErr.Message, instanaErr.Type, instanaErr.StatusCode)
        }
        log.Fatal(err)
    }
    
    log.Printf("Data: %s", string(data))
}
```

### Example 2: Terraform Provider Integration

**Before (provider.go):**
```go
import (
    "github.com/instana/terraform-provider-instana/internal/restapi"
)

func (p *InstanaProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
    // ... config parsing ...
    
    client := restapi.NewClient(apiToken, endpoint, skipTls)
    
    resp.DataSourceData = client
    resp.ResourceData = client
}
```

**After (provider.go):**
```go
import (
    "github.com/instana/instana-go-client/instana"
)

func (p *InstanaProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
    // ... config parsing ...
    
    config, err := instana.NewConfigBuilder().
        WithBaseURL("https://" + endpoint).
        WithAPIToken(apiToken).
        WithConnectionTimeout(30 * time.Second).
        WithRequestTimeout(60 * time.Second).
        WithMaxRetryAttempts(3).
        WithRetryInitialDelay(1 * time.Second).
        WithRetryMaxDelay(30 * time.Second).
        WithRateLimitEnabled(true).
        WithRateLimitRequestsPerSecond(100).
        WithDebug(debugMode).
        Build()
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to Create Instana API Client Configuration",
            "An unexpected error occurred when creating the Instana API client configuration. "+
                "Error: "+err.Error(),
        )
        return
    }
    
    client, err := instana.NewClientWithConfig(config)
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to Create Instana API Client",
            "An unexpected error occurred when creating the Instana API client. "+
                "Error: "+err.Error(),
        )
        return
    }
    
    resp.DataSourceData = client
    resp.ResourceData = client
}
```

### Example 3: Resource Implementation

**Before:**
```go
func (r *ApplicationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    client := req.ProviderData.(restapi.RestClient)
    
    data, err := client.Post(app, "/api/application-monitoring/applications")
    if err != nil {
        resp.Diagnostics.AddError("API Error", err.Error())
        return
    }
    
    // ... process response ...
}
```

**After (with enhanced error handling):**
```go
func (r *ApplicationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    client := req.ProviderData.(instana.RestClient)
    
    data, err := client.Post(app, "/api/application-monitoring/applications")
    if err != nil {
        var instanaErr *instana.InstanaError
        if errors.As(err, &instanaErr) {
            switch instanaErr.Type {
            case instana.ErrorTypeValidation:
                resp.Diagnostics.AddError(
                    "Validation Error",
                    fmt.Sprintf("Invalid application configuration: %s", instanaErr.Message),
                )
            case instana.ErrorTypeAuthentication:
                resp.Diagnostics.AddError(
                    "Authentication Error",
                    "API token is invalid or expired. Please check your credentials.",
                )
            case instana.ErrorTypeRateLimit:
                resp.Diagnostics.AddError(
                    "Rate Limit Exceeded",
                    fmt.Sprintf("Too many requests. Retry after: %v", instanaErr.RetryAfter),
                )
            default:
                resp.Diagnostics.AddError(
                    "API Error",
                    fmt.Sprintf("Failed to create application: %s (Status: %d)", 
                        instanaErr.Message, instanaErr.StatusCode),
                )
            }
        } else {
            resp.Diagnostics.AddError("API Error", err.Error())
        }
        return
    }
    
    // ... process response ...
}
```

---

## Configuration Changes

### Environment Variables (New Feature)

You can now configure the client using environment variables:

```bash
# Core Configuration
export INSTANA_BASE_URL="https://tenant.instana.io"
export INSTANA_API_TOKEN="your-api-token"
export INSTANA_DEBUG="true"

# Timeouts
export INSTANA_CONNECTION_TIMEOUT="30s"
export INSTANA_REQUEST_TIMEOUT="60s"

# Retry Configuration
export INSTANA_MAX_RETRY_ATTEMPTS="5"
export INSTANA_RETRY_INITIAL_DELAY="1s"
export INSTANA_RETRY_MAX_DELAY="30s"

# Rate Limiting
export INSTANA_RATE_LIMIT_ENABLED="true"
export INSTANA_RATE_LIMIT_RPS="100"
```

**Load configuration from environment:**
```go
config, err := instana.LoadFromEnv()
if err != nil {
    log.Fatal(err)
}

client, err := instana.NewClientWithConfig(config)
```

### JSON Configuration (New Feature)

Create a `config.json` file:
```json
{
  "baseURL": "https://tenant.instana.io",
  "apiToken": "your-api-token",
  "timeout": {
    "connection": 30000000000,
    "request": 60000000000
  },
  "retry": {
    "maxAttempts": 5,
    "initialDelay": 1000000000,
    "maxDelay": 30000000000
  },
  "rateLimit": {
    "enabled": true,
    "requestsPerSecond": 100
  }
}
```

**Load configuration from JSON:**
```go
config, err := instana.LoadFromJSON("config.json")
if err != nil {
    log.Fatal(err)
}

client, err := instana.NewClientWithConfig(config)
```

---

## API Changes

### Method Signatures (Unchanged)

All existing methods maintain backward compatibility:

```go
// GET requests
Get(resourcePath string) ([]byte, error)
GetOne(id string, resourcePath string) ([]byte, error)
GetByQuery(resourcePath string, queryParams map[string]string) ([]byte, error)

// POST requests
Post(data InstanaDataObject, resourcePath string) ([]byte, error)
PostWithID(data InstanaDataObject, resourcePath string) ([]byte, error)
PostByQuery(resourcePath string, queryParams map[string]string) ([]byte, error)

// PUT requests
Put(data InstanaDataObject, resourcePath string) ([]byte, error)
PutByQuery(resourcePath string, id string, queryParams map[string]string) ([]byte, error)

// DELETE requests
Delete(resourceID string, resourceBasePath string) error
```

### New Error Types

```go
// Error types for better error handling
const (
    ErrorTypeUnknown        ErrorType = "unknown"
    ErrorTypeNetwork        ErrorType = "network"
    ErrorTypeTimeout        ErrorType = "timeout"
    ErrorTypeAuthentication ErrorType = "authentication"
    ErrorTypeAuthorization  ErrorType = "authorization"
    ErrorTypeValidation     ErrorType = "validation"
    ErrorTypeNotFound       ErrorType = "not_found"
    ErrorTypeRateLimit      ErrorType = "rate_limit"
    ErrorTypeServer         ErrorType = "server"
)

// Enhanced error structure
type InstanaError struct {
    Type       ErrorType
    Message    string
    StatusCode int
    RetryAfter time.Duration
    Err        error
}
```

---

## Error Handling

### Old Error Handling

**Before:**
```go
data, err := client.Get("/api/endpoint")
if err != nil {
    if err == restapi.ErrEntityNotFound {
        // Handle not found
    }
    return err
}
```

### New Error Handling

**After:**
```go
data, err := client.Get("/api/endpoint")
if err != nil {
    // Check for specific error types
    var instanaErr *instana.InstanaError
    if errors.As(err, &instanaErr) {
        switch instanaErr.Type {
        case instana.ErrorTypeNotFound:
            // Handle not found
            log.Printf("Resource not found: %s", instanaErr.Message)
        case instana.ErrorTypeAuthentication:
            // Handle authentication error
            log.Printf("Authentication failed: %s", instanaErr.Message)
        case instana.ErrorTypeRateLimit:
            // Handle rate limit
            log.Printf("Rate limit exceeded, retry after: %v", instanaErr.RetryAfter)
        case instana.ErrorTypeTimeout:
            // Handle timeout
            log.Printf("Request timed out: %s", instanaErr.Message)
        default:
            log.Printf("API error: %s (status: %d)", instanaErr.Message, instanaErr.StatusCode)
        }
    }
    
    // Check if error is retryable
    if instana.IsRetryableError(err) {
        log.Println("This error can be retried")
    }
    
    return err
}
```

---

## Testing Changes

### Mock Client

**Before:**
```go
// Create mock client
mockClient := &MockRestClient{}
mockClient.On("Get", "/api/endpoint").Return([]byte(`{"id":"123"}`), nil)
```

**After (same approach, different import):**
```go
import "github.com/instana/instana-go-client/instana"

// Create mock client
mockClient := &MockRestClient{}
mockClient.On("Get", "/api/endpoint").Return([]byte(`{"id":"123"}`), nil)
```

### Test Configuration

**New: Use test configuration:**
```go
func TestMyFunction(t *testing.T) {
    // Create test configuration
    config := instana.DefaultClientConfig()
    config.BaseURL = "http://localhost:8080"  // Test server
    config.APIToken = "test-token"
    config.Retry.MaxAttempts = 1  // Disable retries in tests
    
    client, err := instana.NewClientWithConfig(config)
    if err != nil {
        t.Fatal(err)
    }
    
    // Run tests...
}
```

---

## Troubleshooting

### Issue 1: Import Errors

**Error:**
```
cannot find package "github.com/instana/terraform-provider-instana/internal/restapi"
```

**Solution:**
```bash
# Update imports
find . -type f -name "*.go" -exec sed -i 's|github.com/instana/terraform-provider-instana/internal/restapi|github.com/instana/instana-go-client/instana|g' {} +

# Update dependencies
go mod tidy
```

### Issue 2: Type Mismatch

**Error:**
```
cannot use client (type *instana.restClientImpl) as type restapi.RestClient
```

**Solution:**
Update variable types:
```go
// Before
var client restapi.RestClient

// After
var client instana.RestClient
```

### Issue 3: Configuration Validation Errors

**Error:**
```
invalid client configuration: base URL is required
```

**Solution:**
Ensure all required fields are set:
```go
config := instana.DefaultClientConfig()
config.BaseURL = "https://tenant.instana.io"  // Required
config.APIToken = "your-token"                 // Required

// Validate before creating client
if err := config.Validate(); err != nil {
    log.Fatal(err)
}
```

### Issue 4: Timeout Issues

**Error:**
```
context deadline exceeded
```

**Solution:**
Increase timeouts:
```go
config, _ := instana.NewConfigBuilder().
    WithConnectionTimeout(60 * time.Second).
    WithRequestTimeout(120 * time.Second).
    Build()
```

### Issue 5: Rate Limiting

**Error:**
```
rate limit exceeded
```

**Solution:**
Enable rate limiting to prevent hitting API limits:
```go
config, _ := instana.NewConfigBuilder().
    WithRateLimitEnabled(true).
    WithRateLimitRequestsPerSecond(50).  // Adjust as needed
    Build()
```

---

## Checklist

Use this checklist to track your migration progress:

- [ ] Update `go.mod` with new dependency
- [ ] Update all import statements
- [ ] Update client creation code
- [ ] Enhance error handling with typed errors
- [ ] Add retry configuration (optional)
- [ ] Add rate limiting (optional)
- [ ] Update tests
- [ ] Run `go test ./...` to verify
- [ ] Run `go build ./...` to verify compilation
- [ ] Update documentation
- [ ] Test in development environment
- [ ] Deploy to staging
- [ ] Monitor for issues
- [ ] Deploy to production

---

## Support

- **Issues**: [GitHub Issues](https://github.com/instana/instana-go-client/issues)
- **Documentation**: [Full Documentation](./README.md)
- **Quick Start**: [Quick Start Guide](./QUICK_START.md)
- **Examples**: [Examples Directory](./examples/)

---

**Migration completed successfully? Let us know!**  
Open an issue or PR to share your experience and help improve this guide.