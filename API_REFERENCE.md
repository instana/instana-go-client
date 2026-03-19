# API Reference - Instana Go Client

Complete API reference for the Instana Go Client library.

## Table of Contents

- [Core Packages](#core-packages)
- [Client Creation](#client-creation)
- [Configuration](#configuration)
- [API Resources](#api-resources)
- [Error Handling](#error-handling)
- [Utilities](#utilities)

---

## Core Packages

### `instana`

Main package providing client creation and configuration.

**Key Functions:**

- [`NewInstanaAPI(apiToken, endpoint, skipTLS)`](instana/Instana-api.go:29) - Create basic API client
- [`NewInstanaAPIWithUserAgent(apiToken, endpoint, skipTLS, userAgent)`](instana/Instana-api.go:57) - Create client with custom user agent
- [`NewInstanaAPIWithConfig(config)`](instana/Instana-api.go:83) - Create client with full configuration
- [`NewClient(apiToken, host, skipTLS)`](instana/rest-client.go) - Create REST client (legacy)
- [`NewClientWithConfig(config)`](instana/rest-client.go) - Create REST client with configuration

### `client`

Package providing the main API interface and resource clients.

**Key Types:**

- [`InstanaAPI`](client/interface.go:11) - Main interface for all API operations
- Factory function: [`NewInstanaAPI(restClient)`](client/factory.go) - Create API client from REST client

### `config`

Package for client configuration management.

**Key Types:**

- [`ClientConfig`](config/config.go:12) - Main configuration structure
- [`TimeoutConfig`](config/config.go:54) - Timeout settings
- [`RetryConfig`](config/config.go:77) - Retry behavior settings
- [`RateLimitConfig`](config/rate_limiter.go) - Rate limiting settings
- [`ConfigBuilder`](config/config_builder.go) - Fluent configuration builder

**Key Functions:**

- [`DefaultClientConfig()`](config/config.go) - Get default configuration
- [`NewConfigBuilder()`](config/config_builder.go) - Create configuration builder
- [`LoadFromEnv()`](config/config_loader.go) - Load configuration from environment
- [`LoadFromJSON(path)`](config/config_loader.go) - Load configuration from JSON file
- [`LoadFromJSONWithEnvOverride(path)`](config/config_loader.go) - Load from JSON with env overrides

---

## Client Creation

### Basic Client Creation

```go
import "github.com/instana/instana-go-client/instana"

// Simple client
api := instana.NewInstanaAPI("api-token", "tenant.instana.io", false)

// With custom user agent
api := instana.NewInstanaAPIWithUserAgent(
    "api-token",
    "tenant.instana.io",
    false,
    "MyApp/1.0.0",
)
```

### Advanced Client Creation

```go
import (
    "github.com/instana/instana-go-client/instana"
    "time"
)

// Using builder pattern
config, err := instana.NewConfigBuilder().
    WithBaseURL("https://tenant.instana.io").
    WithAPIToken("api-token").
    WithConnectionTimeout(30 * time.Second).
    WithMaxRetryAttempts(5).
    WithRateLimitRequestsPerSecond(100).
    Build()

api, err := instana.NewInstanaAPIWithConfig(config)
```

### Configuration from Environment

```go
// Set environment variables first:
// export INSTANA_BASE_URL="https://tenant.instana.io"
// export INSTANA_API_TOKEN="your-token"

config, err := instana.LoadFromEnv()
api, err := instana.NewInstanaAPIWithConfig(config)
```

---

## Configuration

### Configuration Builder

The [`ConfigBuilder`](config/config_builder.go) provides a fluent interface for configuration:

#### Core Settings

```go
builder := instana.NewConfigBuilder().
    WithBaseURL("https://tenant.instana.io").
    WithAPIToken("token").
    WithUserAgent("MyApp/1.0").
    WithDebug(true)
```

#### Timeout Configuration

```go
builder.
    WithConnectionTimeout(30 * time.Second).
    WithRequestTimeout(60 * time.Second).
    WithIdleConnectionTimeout(90 * time.Second).
    WithResponseHeaderTimeout(10 * time.Second).
    WithTLSHandshakeTimeout(10 * time.Second)
```

#### Retry Configuration

```go
builder.
    WithMaxRetryAttempts(5).
    WithRetryInitialDelay(1 * time.Second).
    WithRetryMaxDelay(30 * time.Second).
    WithRetryBackoffMultiplier(2.0).
    WithRetryOnTimeout(true).
    WithRetryOnConnectionError(true).
    WithRetryJitter(true)
```

#### Rate Limiting

```go
builder.
    WithRateLimitEnabled(true).
    WithRateLimitRequestsPerSecond(100).
    WithRateLimitBurstCapacity(200).
    WithRateLimitWaitForToken(true)
```

#### Connection Pooling

```go
builder.
    WithMaxIdleConnections(100).
    WithMaxConnectionsPerHost(10).
    WithMaxIdleConnectionsPerHost(10).
    WithKeepAliveDuration(30 * time.Second).
    WithDisableKeepAlives(false).
    WithDisableCompression(false)
```

#### Custom Headers

```go
builder.
    WithCustomHeader("X-Request-ID", "unique-id").
    WithCustomHeaders(map[string]string{
        "X-Trace-ID": "trace-123",
        "X-Custom": "value",
    })
```

### Environment Variables

| Variable | Type | Default | Description |
|----------|------|---------|-------------|
| `INSTANA_BASE_URL` | string | - | Base URL of Instana API |
| `INSTANA_API_TOKEN` | string | - | API authentication token |
| `INSTANA_DEBUG` | bool | false | Enable debug logging |
| `INSTANA_CONNECTION_TIMEOUT` | duration | 30s | Connection timeout |
| `INSTANA_REQUEST_TIMEOUT` | duration | 60s | Request timeout |
| `INSTANA_IDLE_CONNECTION_TIMEOUT` | duration | 90s | Idle connection timeout |
| `INSTANA_MAX_RETRY_ATTEMPTS` | int | 3 | Maximum retry attempts |
| `INSTANA_RETRY_INITIAL_DELAY` | duration | 1s | Initial retry delay |
| `INSTANA_RETRY_MAX_DELAY` | duration | 30s | Maximum retry delay |
| `INSTANA_RETRY_BACKOFF_MULTIPLIER` | float | 2.0 | Backoff multiplier |
| `INSTANA_RATE_LIMIT_ENABLED` | bool | true | Enable rate limiting |
| `INSTANA_RATE_LIMIT_RPS` | int | 100 | Requests per second |
| `INSTANA_RATE_LIMIT_BURST` | int | 200 | Burst capacity |
| `INSTANA_MAX_IDLE_CONNECTIONS` | int | 100 | Max idle connections |
| `INSTANA_MAX_CONNECTIONS_PER_HOST` | int | 10 | Max connections per host |
| `INSTANA_KEEP_ALIVE_DURATION` | duration | 30s | Keep-alive duration |
| `INSTANA_BATCH_SIZE` | int | 100 | Batch operation size |
| `INSTANA_BATCH_CONCURRENT_REQUESTS` | int | 5 | Concurrent batch requests |

---

## API Resources

The [`InstanaAPI`](client/interface.go:11) interface provides access to all Instana resources:

### Monitoring & Observability

#### Application Monitoring

```go
// Application configurations
apps := api.ApplicationConfigs()
allApps, err := apps.GetAll()
app, err := apps.GetOne("app-id")
created, err := apps.Create(newApp)
updated, err := apps.Update(app)
err = apps.Delete(app)
```

#### Application Alerts

```go
// Application alert configurations
alerts := api.ApplicationAlertConfigs()
allAlerts, err := alerts.GetAll()

// Global application alerts
globalAlerts := api.GlobalApplicationAlertConfigs()
```

#### Infrastructure Monitoring

```go
// Infrastructure alert configurations
infraAlerts := api.InfraAlertConfigs()

// Host agents (read-only)
agents := api.HostAgents()
allAgents, err := agents.GetAll()
agent, err := agents.GetOne("agent-id")
```

### Synthetic Monitoring

```go
// Synthetic tests
tests := api.SyntheticTests()
allTests, err := tests.GetAll()
test, err := tests.GetOne("test-id")
created, err := tests.Create(newTest)

// Synthetic locations (read-only)
locations := api.SyntheticLocations()
allLocations, err := locations.GetAll()

// Synthetic alerts
syntheticAlerts := api.SyntheticAlertConfigs()
```

### Website Monitoring

```go
// Website monitoring configurations
websites := api.WebsiteMonitoringConfigs()
allSites, err := websites.GetAll()
site, err := websites.GetOne("site-id")

// Website alert configurations
websiteAlerts := api.WebsiteAlertConfigs()

// Mobile alert configurations
mobileAlerts := api.MobileAlertConfigs()
```

### SLO/SLI Management

```go
// SLO configurations
slos := api.SloConfigs()
allSlos, err := slos.GetAll()
slo, err := slos.Create(newSlo)

// SLI configurations
slis := api.SliConfigs()

// SLO alerts
sloAlerts := api.SloAlertConfigs()

// SLO corrections
corrections := api.SloCorrectionConfigs()
```

### Alerting & Notifications

```go
// Alerting channels
channels := api.AlertingChannels()
allChannels, err := channels.GetAll()
channel, err := channels.Create(newChannel)

// Alerting configurations
alertConfigs := api.AlertingConfigurations()

// Log alert configurations
logAlerts := api.LogAlertConfigs()
```

### Events & Dashboards

```go
// Custom event specifications
customEvents := api.CustomEventSpecifications()
allEvents, err := customEvents.GetAll()

// Built-in event specifications (read-only)
builtinEvents := api.BuiltinEventSpecifications()
allBuiltin, err := builtinEvents.GetAll()

// Custom dashboards
dashboards := api.CustomDashboards()
allDashboards, err := dashboards.GetAll()
dashboard, err := dashboards.Create(newDashboard)
```

### Automation

```go
// Automation actions
actions := api.AutomationActions()
allActions, err := actions.GetAll()
action, err := actions.Create(newAction)

// Automation policies
policies := api.AutomationPolicies()
allPolicies, err := policies.GetAll()
```

### Maintenance

```go
// Maintenance windows
windows := api.MaintenanceWindowConfigs()
allWindows, err := windows.GetAll()
window, err := windows.Create(newWindow)
updated, err := windows.Update(window)
err = windows.Delete(window)
```

### Access Control (RBAC)

```go
// API tokens
tokens := api.APITokens()
allTokens, err := tokens.GetAll()
token, err := tokens.Create(newToken)

// Users (read-only)
users := api.Users()
allUsers, err := users.GetAll()
user, err := users.GetOne("user-id")

// Groups
groups := api.Groups()
allGroups, err := groups.GetAll()
group, err := groups.Create(newGroup)

// Roles
roles := api.Roles()
allRoles, err := roles.GetAll()
role, err := roles.Create(newRole)

// Teams
teams := api.Teams()
allTeams, err := teams.GetAll()
team, err := teams.Create(newTeam)
```

### REST Resource Interface

All resources implement either [`RestResource`](shared/rest/resource.go:9) or [`ReadOnlyRestResource`](shared/rest/resource.go:20):

```go
// RestResource interface (full CRUD)
type RestResource[T InstanaDataObject] interface {
    GetAll() (*[]T, error)
    GetOne(id string) (T, error)
    Create(data T) (T, error)
    Update(data T) (T, error)
    Delete(data T) error
    DeleteByID(id string) error
}

// ReadOnlyRestResource interface (read-only)
type ReadOnlyRestResource[T InstanaDataObject] interface {
    GetAll() (*[]T, error)
    GetByQuery(queryParams map[string]string) (*[]T, error)
    GetOne(id string) (T, error)
}
```

---

## Error Handling

### Error Types

The library provides typed errors through [`InstanaError`](config/errors.go):

```go
type InstanaError struct {
    Type       ErrorType
    Message    string
    StatusCode int
    Err        error
}
```

#### Error Types

- `ErrorTypeNetwork` - Network connectivity errors (retryable)
- `ErrorTypeAPI` - API response errors (may be retryable)
- `ErrorTypeValidation` - Input validation errors (not retryable)
- `ErrorTypeAuthentication` - Authentication failures (not retryable)
- `ErrorTypeRateLimit` - Rate limit exceeded (retryable)
- `ErrorTypeTimeout` - Request timeout (retryable)
- `ErrorTypeSerialization` - JSON serialization errors (not retryable)

### Error Checking

```go
import "github.com/instana/instana-go-client/config"

data, err := client.Get("/api/endpoint")
if err != nil {
    // Check if error is retryable
    if config.IsRetryableError(err) {
        // Retry the operation
    }
    
    // Check if error is temporary
    if config.IsTemporaryError(err) {
        // Wait and retry
    }
    
    // Extract HTTP status code
    statusCode := config.ExtractStatusCode(err)
    
    // Type assertion for detailed information
    if instanaErr, ok := err.(*config.InstanaError); ok {
        fmt.Printf("Error Type: %s\n", instanaErr.Type)
        fmt.Printf("Status Code: %d\n", instanaErr.StatusCode)
        fmt.Printf("Retryable: %v\n", instanaErr.IsRetryable())
    }
}
```

### Common Error Scenarios

```go
// Handle authentication errors
if instanaErr, ok := err.(*config.InstanaError); ok {
    if instanaErr.Type == config.ErrorTypeAuthentication {
        // Invalid API token
    }
}

// Handle rate limiting
if config.IsRateLimitError(err) {
    // Wait before retrying
    time.Sleep(time.Second)
}

// Handle validation errors
if instanaErr, ok := err.(*config.InstanaError); ok {
    if instanaErr.Type == config.ErrorTypeValidation {
        // Fix input data
    }
}
```

---

## Utilities

### Retry Mechanism

The [`Retryer`](config/retry.go) provides automatic retry with exponential backoff:

```go
import (
    "context"
    "github.com/instana/instana-go-client/config"
)

retryConfig := config.DefaultRetryConfig()
retryer := config.NewRetryer(retryConfig, logger)

err := retryer.Do(context.Background(), func() error {
    // Your operation here
    return someAPICall()
})
```

### Rate Limiter

The [`RateLimiter`](config/rate_limiter.go) implements token bucket algorithm:

```go
import (
    "context"
    "github.com/instana/instana-go-client/config"
)

rateLimitConfig := config.RateLimitConfig{
    Enabled:           true,
    RequestsPerSecond: 50,
    BurstCapacity:     100,
    WaitForToken:      true,
}

rateLimiter := config.NewRateLimiter(rateLimitConfig, logger)
defer rateLimiter.Stop()

// Wait for rate limit token
if err := rateLimiter.Wait(context.Background()); err != nil {
    // Handle error
}
```

### Logging

The library provides a [`Logger`](config/logger.go) interface:

```go
// Use default logger
logger := config.NewDefaultLogger(config.ClientLogLevelInfo)

// Use no-op logger (disable logging)
logger := config.NewNoOpLogger()

// Custom logger implementation
type MyLogger struct{}

func (l *MyLogger) Debug(msg string, keysAndValues ...interface{}) {}
func (l *MyLogger) Info(msg string, keysAndValues ...interface{}) {}
func (l *MyLogger) Error(msg string, keysAndValues ...interface{}) {}
```

### Log Levels

```go
const (
    ClientLogLevelDebug ClientLogLevel = "debug"
    ClientLogLevelInfo  ClientLogLevel = "info"
    ClientLogLevelError ClientLogLevel = "error"
)
```

---

## Best Practices

### 1. Use Configuration Builder

```go
// Recommended
config, err := instana.NewConfigBuilder().
    WithBaseURL("https://tenant.instana.io").
    WithAPIToken("token").
    WithMaxRetryAttempts(5).
    Build()
```

### 2. Enable Retry and Rate Limiting

```go
config, err := instana.NewConfigBuilder().
    WithMaxRetryAttempts(5).
    WithRateLimitEnabled(true).
    WithRateLimitRequestsPerSecond(100).
    Build()
```

### 3. Handle Errors Properly

```go
data, err := api.ApplicationConfigs().GetAll()
if err != nil {
    if config.IsRetryableError(err) {
        // Implement retry logic
    } else {
        // Handle non-retryable error
    }
}
```

### 4. Use Context for Cancellation

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

// Pass context to operations that support it
```

### 5. Reuse Client Instances

```go
// Create once, reuse many times
api := instana.NewInstanaAPI("token", "tenant.instana.io", false)

// Use across your application
apps, _ := api.ApplicationConfigs().GetAll()
alerts, _ := api.ApplicationAlertConfigs().GetAll()
```

---

## See Also

- [README](README.md) - Project overview and features
- [Quick Start Guide](QUICK_START.md) - Getting started quickly
- [Usage Guide](USAGE_GUIDE.md) - Detailed usage patterns
- [Architecture](ARCHITECTURE.md) - System architecture
- [Examples](examples/) - Code examples