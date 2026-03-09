# User Agent / Version Tracking

## Overview

The Instana Go Client Library now supports custom User-Agent headers for tracking client versions. This replaces the previous approach of reading version information from CHANGELOG.md files.

## Problem Statement

**Previous Approach**:
- Terraform provider read version from CHANGELOG.md at runtime
- Required file system access to read CHANGELOG.md
- Coupled the go-client to terraform-specific file structure
- Not suitable for standalone go-client library

**New Approach**:
- Client accepts User-Agent as a configuration parameter
- Terraform provider passes its version as User-Agent
- Go-client is decoupled from terraform-specific code
- Flexible for any client to identify itself

## Implementation

### 1. Go Client Library

#### New Function: `NewInstanaAPIWithUserAgent()`

```go
// NewInstanaAPIWithUserAgent creates a new instance of the instana API with a custom user agent
// The userAgent parameter should include the client name and version (e.g., "Terraform/1.2.3")
func NewInstanaAPIWithUserAgent(apiToken string, endpoint string, skipTlsVerification bool, userAgent string) InstanaAPI
```

**Parameters**:
- `apiToken`: Instana API token for authentication
- `endpoint`: Instana endpoint (e.g., "tenant-unit.instana.io")
- `skipTlsVerification`: Whether to skip TLS verification
- `userAgent`: Custom user agent string (e.g., "Terraform/1.2.3")

**Behavior**:
1. Creates a `ClientConfig` with all default values
2. Sets the custom `UserAgent` from the parameter
3. Creates HTTP client with TLS configuration if needed
4. Returns fully configured `InstanaAPI` instance

**Example**:
```go
api := instana.NewInstanaAPIWithUserAgent(
    "api-token",
    "tenant-unit.instana.io",
    false,
    "Terraform/1.2.3",
)
```

#### Backward Compatibility

The original `NewInstanaAPI()` function remains unchanged:

```go
// NewInstanaAPI creates a new instance of the instana API
func NewInstanaAPI(apiToken string, endpoint string, skipTlsVerification bool) InstanaAPI
```

**Default User-Agent**: `"instana-go-client/v1.0.0"`

### 2. Terraform Provider Integration

#### Provider Configuration

The terraform provider has access to its version through the `InstanaProvider.version` field:

```go
type InstanaProvider struct {
    version string  // Set to provider version on release
}
```

#### Updated Client Creation

**File**: `internal/provider/provider.go`

**Before**:
```go
instanaAPI := instana.NewInstanaAPI(apiToken, endpoint, skipTlsVerify)
```

**After**:
```go
// Pass the provider version as user agent for tracking
userAgent := fmt.Sprintf("Terraform/%s", p.version)
instanaAPI := instana.NewInstanaAPIWithUserAgent(apiToken, endpoint, skipTlsVerify, userAgent)
```

## User-Agent Format

### Recommended Format

```
<ClientName>/<Version>
```

### Examples

| Client | User-Agent | Description |
|--------|-----------|-------------|
| Terraform Provider | `Terraform/1.2.3` | Terraform provider version 1.2.3 |
| Go Application | `MyApp/2.0.0` | Custom Go application |
| CLI Tool | `InstanaCLI/0.5.0` | Command-line tool |
| SDK | `InstanaSDK/3.1.0` | SDK wrapper |

### HTTP Header

The User-Agent is sent in every HTTP request:

```http
GET /api/applications HTTP/1.1
Host: tenant-unit.instana.io
Accept: application/json
Authorization: apiToken <token>
User-Agent: Terraform/1.2.3
```

## Benefits

### 1. Better Tracking
- Instana backend can track which clients are making requests
- Identify terraform provider versions in use
- Monitor adoption of new versions
- Debug version-specific issues

### 2. Decoupling
- Go-client no longer depends on terraform file structure
- No file system access required
- Cleaner separation of concerns
- Easier to maintain

### 3. Flexibility
- Any client can identify itself
- Custom applications can set their own User-Agent
- Multiple clients can use the same go-client library
- Version information comes from the source (provider/app)

### 4. Backward Compatibility
- Existing code using `NewInstanaAPI()` continues to work
- Default User-Agent is set automatically
- No breaking changes

## Migration Guide

### For Terraform Provider Users

**No action required!** The terraform provider automatically passes its version.

### For Go Client Library Users

#### Option 1: Use Default User-Agent (Existing Code)

```go
// Uses default: "instana-go-client/v1.0.0"
api := instana.NewInstanaAPI(apiToken, endpoint, false)
```

#### Option 2: Set Custom User-Agent (New Code)

```go
// Set your application name and version
api := instana.NewInstanaAPIWithUserAgent(
    apiToken,
    endpoint,
    false,
    "MyApp/1.0.0",  // Your custom user agent
)
```

#### Option 3: Use Configuration Builder

```go
config, _ := instana.NewConfigBuilder().
    WithBaseURL("https://tenant-unit.instana.io").
    WithAPIToken(apiToken).
    WithUserAgent("MyApp/1.0.0").  // Custom user agent
    Build()

client, _ := instana.NewClientWithConfig(config)
```

## Implementation Details

### Code Flow

```
Terraform Provider (version: "1.2.3")
    ↓
    Calls: NewInstanaAPIWithUserAgent(..., "Terraform/1.2.3")
    ↓
Go Client Library
    ↓
    Creates ClientConfig with UserAgent = "Terraform/1.2.3"
    ↓
    Creates RestClient with config
    ↓
    Every HTTP request includes: User-Agent: Terraform/1.2.3
    ↓
Instana Backend
    ↓
    Logs/tracks requests by User-Agent
```

### Configuration Priority

When creating a client, the User-Agent is set in this order:

1. **Explicit parameter** in `NewInstanaAPIWithUserAgent()`
2. **Configuration** via `ClientConfig.UserAgent`
3. **Default** value: `"instana-go-client/v1.0.0"`

### Example: Multiple Clients

```go
// Terraform provider
terraformAPI := instana.NewInstanaAPIWithUserAgent(
    token, endpoint, false, "Terraform/1.2.3",
)

// Custom application
appAPI := instana.NewInstanaAPIWithUserAgent(
    token, endpoint, false, "MyApp/2.0.0",
)

// CLI tool
cliAPI := instana.NewInstanaAPIWithUserAgent(
    token, endpoint, false, "InstanaCLI/0.5.0",
)

// Each client sends its own User-Agent in requests
```

## Testing

### Verify User-Agent is Sent

You can verify the User-Agent is being sent by:

1. **Enable Debug Logging**:
```go
config := instana.DefaultClientConfig()
config.Debug = true
config.Logger = instana.NewDefaultLogger(instana.ClientLogLevelDebug)
```

2. **Check HTTP Logs**:
The debug logs will show the User-Agent header in requests.

3. **Network Inspection**:
Use tools like Wireshark or browser dev tools to inspect HTTP headers.

### Unit Test Example

```go
func TestUserAgentIsSet(t *testing.T) {
    api := instana.NewInstanaAPIWithUserAgent(
        "test-token",
        "test.instana.io",
        false,
        "TestClient/1.0.0",
    )
    
    // Verify API is created
    if api == nil {
        t.Error("Expected API to be created")
    }
    
    // Make a request and verify User-Agent header
    // (requires mock HTTP server)
}
```

## Comparison: Old vs New

### Old Approach (Removed)

```go
// Read version from CHANGELOG.md
file, err := os.Open(basepath + "/CHANGELOG.md")
scanner := bufio.NewScanner(file)
for !strings.Contains(scanner.Text(), "##") {
    scanner.Scan()
}
terraformProviderVersion = scanner.Text()
// Parse and format version...
```

**Problems**:
- ❌ File system dependency
- ❌ Runtime file reading
- ❌ Coupled to terraform structure
- ❌ Error-prone parsing
- ❌ Not suitable for standalone library

### New Approach (Current)

```go
// Provider passes version directly
userAgent := fmt.Sprintf("Terraform/%s", p.version)
api := instana.NewInstanaAPIWithUserAgent(apiToken, endpoint, skipTlsVerify, userAgent)
```

**Benefits**:
- ✅ No file system access
- ✅ Version from source (provider)
- ✅ Decoupled from file structure
- ✅ Simple and reliable
- ✅ Works for any client

## Summary

| Aspect | Old Approach | New Approach |
|--------|-------------|--------------|
| **Version Source** | CHANGELOG.md file | Parameter/Config |
| **File Access** | Required | Not required |
| **Coupling** | High (terraform-specific) | Low (generic) |
| **Flexibility** | Limited to terraform | Any client |
| **Reliability** | File parsing errors | Direct parameter |
| **Performance** | File I/O overhead | No overhead |
| **Maintainability** | Complex | Simple |

## Conclusion

The new User-Agent tracking approach provides:
- ✅ Better tracking and monitoring capabilities
- ✅ Clean separation between go-client and terraform provider
- ✅ Flexibility for any client to identify itself
- ✅ 100% backward compatibility
- ✅ Simpler, more maintainable code

The terraform provider now passes its version directly to the go-client, which includes it in the User-Agent header of all HTTP requests to the Instana backend.