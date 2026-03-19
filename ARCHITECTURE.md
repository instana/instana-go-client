# Architecture - Instana Go Client

This document describes the architecture, design patterns, and internal structure of the Instana Go Client library.

## Table of Contents

- [Overview](#overview)
- [Package Structure](#package-structure)
- [Design Patterns](#design-patterns)
- [Core Components](#core-components)
- [Data Flow](#data-flow)
- [Configuration System](#configuration-system)
- [Error Handling Strategy](#error-handling-strategy)
- [Performance Optimizations](#performance-optimizations)

---

## Overview

The Instana Go Client is designed as a production-ready, type-safe client library for the Instana API. It follows Go best practices and provides a clean, idiomatic interface for interacting with Instana's monitoring platform.

### Design Goals

1. **Type Safety** - Leverage Go's type system for compile-time safety
2. **Simplicity** - Provide intuitive APIs that are easy to use
3. **Flexibility** - Support multiple configuration methods
4. **Reliability** - Built-in retry, rate limiting, and error handling
5. **Performance** - Efficient connection pooling and resource management
6. **Testability** - Interfaces and dependency injection for easy testing

---

## Package Structure

```
instana-go-client/
├── instana/              # Main package - client creation and REST operations
│   ├── Instana-api.go    # High-level API client factory functions
│   └── rest-client.go    # Low-level REST client implementation
├── client/               # API interface and resource clients
│   ├── interface.go      # InstanaAPI interface definition
│   ├── factory.go        # Client factory implementation
│   └── client.go         # Main client implementation with lazy loading
├── config/               # Configuration management
│   ├── config.go         # Configuration structures
│   ├── config_builder.go # Fluent builder pattern
│   ├── config_loader.go  # Environment and JSON loading
│   ├── config_validator.go # Configuration validation
│   ├── retry.go          # Retry mechanism
│   ├── rate_limiter.go   # Rate limiting
│   ├── errors.go         # Error types
│   └── logger.go         # Logging interface
├── api/                  # API resource data models
│   ├── apitoken.go       # API token models
│   ├── applicationconfig.go # Application configuration models
│   ├── alertingchannel.go # Alerting channel models
│   └── ...               # Other resource models
├── shared/               # Shared utilities and interfaces
│   ├── rest/             # REST resource interfaces
│   │   ├── resource.go   # RestResource interface
│   │   ├── default.go    # Default REST implementation
│   │   └── readonly.go   # Read-only REST implementation
│   ├── types/            # Shared type definitions
│   └── tagfilter/        # Tag filtering utilities
├── mocks/                # Mock implementations for testing
├── testutils/            # Testing utilities
├── utils/                # Generic utilities
└── examples/             # Usage examples
```

### Package Responsibilities

| Package | Responsibility |
|---------|---------------|
| `instana` | Entry point, client creation, REST operations |
| `client` | High-level API interface, resource management |
| `config` | Configuration, validation, retry, rate limiting |
| `api` | Data models for all API resources |
| `shared/rest` | REST resource interfaces and implementations |
| `shared/types` | Common type definitions |
| `mocks` | Generated mocks for testing |
| `testutils` | Test helpers and utilities |

---

## Design Patterns

### 1. Builder Pattern

Used for flexible configuration construction:

```go
// config/config_builder.go
type ConfigBuilder struct {
    config *ClientConfig
}

func (b *ConfigBuilder) WithBaseURL(url string) *ConfigBuilder {
    b.config.BaseURL = url
    return b
}

func (b *ConfigBuilder) Build() (*ClientConfig, error) {
    return b.config, b.validate()
}
```

**Benefits:**
- Fluent, chainable API
- Immutable configuration after build
- Validation at build time
- Optional parameters with defaults

### 2. Factory Pattern

Used for creating API clients and resources:

```go
// client/factory.go
func NewInstanaAPI(restClient rest.RestClient) InstanaAPI {
    return &instanaAPIImpl{
        restClient: restClient,
    }
}
```

**Benefits:**
- Encapsulates object creation
- Allows for different implementations
- Simplifies client initialization

### 3. Lazy Initialization

Used for API resource clients:

```go
// client/client.go
func (c *instanaAPIImpl) APITokens() rest.RestResource[*api.APIToken] {
    c.once.apiTokens.Do(func() {
        c.apiTokensClient = createAPITokenClient(c.restClient)
    })
    return c.apiTokensClient
}
```

**Benefits:**
- Only creates clients when needed
- Reduces memory footprint
- Improves startup time
- Thread-safe initialization

### 4. Strategy Pattern

Used for retry and rate limiting:

```go
// config/retry.go
type Retryer struct {
    config RetryConfig
    shouldRetry func(error) bool
}

func (r *Retryer) Do(ctx context.Context, fn func() error) error {
    // Implements retry strategy
}
```

**Benefits:**
- Pluggable retry strategies
- Configurable behavior
- Easy to test and extend

### 5. Interface Segregation

Separate interfaces for different capabilities:

```go
// shared/rest/resource.go
type RestResource[T InstanaDataObject] interface {
    GetAll() (*[]T, error)
    GetOne(id string) (T, error)
    Create(data T) (T, error)
    Update(data T) (T, error)
    Delete(data T) error
    DeleteByID(id string) error
}

type ReadOnlyRestResource[T InstanaDataObject] interface {
    GetAll() (*[]T, error)
    GetByQuery(queryParams map[string]string) (*[]T, error)
    GetOne(id string) (T, error)
}
```

**Benefits:**
- Clients only depend on what they need
- Clear API contracts
- Easier to mock and test

---

## Core Components

### 1. REST Client

**Location:** [`instana/rest-client.go`](instana/rest-client.go)

The REST client handles low-level HTTP operations:

```
┌─────────────────────────────────────┐
│         REST Client                 │
├─────────────────────────────────────┤
│ - HTTP Client                       │
│ - Rate Limiter                      │
│ - Retryer                           │
│ - Logger                            │
├─────────────────────────────────────┤
│ + Get(path)                         │
│ + Post(data, path)                  │
│ + Put(data, path)                   │
│ + Delete(id, path)                  │
└─────────────────────────────────────┘
```

**Responsibilities:**
- Execute HTTP requests
- Apply rate limiting
- Handle retries
- Manage authentication
- Log requests/responses

### 2. API Client

**Location:** [`client/client.go`](client/client.go)

The API client provides high-level resource access:

```
┌─────────────────────────────────────┐
│         InstanaAPI                  │
├─────────────────────────────────────┤
│ - REST Client                       │
│ - Resource Clients (lazy)           │
├─────────────────────────────────────┤
│ + APITokens()                       │
│ + ApplicationConfigs()              │
│ + AlertingChannels()                │
│ + ... (28 resources)                │
└─────────────────────────────────────┘
```

**Responsibilities:**
- Manage resource clients
- Lazy initialization
- Provide unified interface
- Handle resource lifecycle

### 3. Configuration System

**Location:** [`config/`](config/)

Multi-layered configuration system:

```
┌─────────────────────────────────────┐
│     Configuration Sources           │
├─────────────────────────────────────┤
│  Builder → Env Vars → JSON → Code  │
└─────────────────────────────────────┘
           ↓
┌─────────────────────────────────────┐
│     Configuration Validator         │
├─────────────────────────────────────┤
│  - Validate URLs                    │
│  - Validate tokens                  │
│  - Validate timeouts                │
│  - Apply defaults                   │
└─────────────────────────────────────┘
           ↓
┌─────────────────────────────────────┐
│     ClientConfig                    │
├─────────────────────────────────────┤
│  - Core settings                    │
│  - Timeout config                   │
│  - Retry config                     │
│  - Rate limit config                │
│  - Connection pool config           │
└─────────────────────────────────────┘
```

### 4. Retry Mechanism

**Location:** [`config/retry.go`](config/retry.go)

Exponential backoff with jitter:

```
Attempt 1: Immediate
Attempt 2: 1s + jitter
Attempt 3: 2s + jitter
Attempt 4: 4s + jitter
Attempt 5: 8s + jitter (capped at maxDelay)
```

**Algorithm:**
```
delay = min(initialDelay * (backoffMultiplier ^ attempt), maxDelay)
jitter = random(0, delay * 0.3)
actualDelay = delay + jitter
```

### 5. Rate Limiter

**Location:** [`config/rate_limiter.go`](config/rate_limiter.go)

Token bucket algorithm:

```
┌─────────────────────────────────────┐
│        Token Bucket                 │
├─────────────────────────────────────┤
│  Capacity: 200 tokens               │
│  Refill Rate: 100 tokens/sec        │
├─────────────────────────────────────┤
│  Request → Take Token → Proceed     │
│  No Token → Wait or Reject          │
└─────────────────────────────────────┘
```

**Features:**
- Configurable rate and burst
- Context-aware waiting
- Graceful shutdown
- Thread-safe

---

## Data Flow

### Request Flow

```
User Code
    ↓
InstanaAPI Interface
    ↓
Resource Client (e.g., APITokens)
    ↓
REST Client
    ↓
Rate Limiter (wait for token)
    ↓
HTTP Request
    ↓
Retry Logic (on failure)
    ↓
Response Processing
    ↓
JSON Unmarshalling
    ↓
Return to User
```

### Detailed Request Flow

```
┌──────────────────────────────────────────────────────────┐
│ 1. User calls api.APITokens().GetAll()                   │
└──────────────────────────────────────────────────────────┘
                        ↓
┌──────────────────────────────────────────────────────────┐
│ 2. Lazy initialization of APITokens client               │
│    - Creates client if not exists (thread-safe)          │
└──────────────────────────────────────────────────────────┘
                        ↓
┌──────────────────────────────────────────────────────────┐
│ 3. Resource client calls restClient.Get(path)            │
└──────────────────────────────────────────────────────────┘
                        ↓
┌──────────────────────────────────────────────────────────┐
│ 4. Rate limiter checks token availability                │
│    - Wait if no tokens available                         │
│    - Respect context cancellation                        │
└──────────────────────────────────────────────────────────┘
                        ↓
┌──────────────────────────────────────────────────────────┐
│ 5. Execute HTTP request                                  │
│    - Add authentication header                           │
│    - Add custom headers                                  │
│    - Set timeouts                                        │
└──────────────────────────────────────────────────────────┘
                        ↓
┌──────────────────────────────────────────────────────────┐
│ 6. Handle response                                       │
│    - Check status code                                   │
│    - Read response body                                  │
│    - Log request/response (if debug)                     │
└──────────────────────────────────────────────────────────┘
                        ↓
┌──────────────────────────────────────────────────────────┐
│ 7. Error handling                                        │
│    - Create typed error                                  │
│    - Check if retryable                                  │
│    - Apply retry logic if needed                         │
└──────────────────────────────────────────────────────────┘
                        ↓
┌──────────────────────────────────────────────────────────┐
│ 8. Unmarshal JSON to Go struct                           │
│    - Type-safe unmarshalling                             │
│    - Validate data structure                             │
└──────────────────────────────────────────────────────────┘
                        ↓
┌──────────────────────────────────────────────────────────┐
│ 9. Return result to user                                 │
└──────────────────────────────────────────────────────────┘
```

---

## Configuration System

### Configuration Hierarchy

```
1. Code Defaults (lowest priority)
   ↓
2. JSON Configuration File
   ↓
3. Environment Variables
   ↓
4. Builder Pattern (highest priority)
```

### Configuration Loading

```go
// 1. Start with defaults
config := DefaultClientConfig()

// 2. Load from JSON (optional)
if jsonConfig, err := LoadFromJSON("config.json"); err == nil {
    config = mergeConfigs(config, jsonConfig)
}

// 3. Override with environment variables
if envConfig, err := LoadFromEnv(); err == nil {
    config = mergeConfigs(config, envConfig)
}

// 4. Apply builder overrides
config = builder.applyOverrides(config)

// 5. Validate final configuration
if err := ValidateConfig(config); err != nil {
    return nil, err
}
```

### Configuration Validation

```
┌─────────────────────────────────────┐
│   Configuration Validator           │
├─────────────────────────────────────┤
│ ✓ BaseURL format and scheme         │
│ ✓ APIToken non-empty                │
│ ✓ Timeout values > 0                │
│ ✓ Retry attempts > 0                │
│ ✓ Rate limit values valid           │
│ ✓ Connection pool settings          │
└─────────────────────────────────────┘
```

---

## Error Handling Strategy

### Error Type Hierarchy

```
error (interface)
    ↓
InstanaError (struct)
    ├── Type: ErrorType
    ├── Message: string
    ├── StatusCode: int
    └── Err: error (wrapped)
```

### Error Classification

```go
type ErrorType string

const (
    ErrorTypeNetwork         ErrorType = "network"         // Retryable
    ErrorTypeAPI            ErrorType = "api"             // Maybe retryable
    ErrorTypeValidation     ErrorType = "validation"      // Not retryable
    ErrorTypeAuthentication ErrorType = "authentication"  // Not retryable
    ErrorTypeRateLimit      ErrorType = "rate_limit"      // Retryable
    ErrorTypeTimeout        ErrorType = "timeout"         // Retryable
    ErrorTypeSerialization  ErrorType = "serialization"   // Not retryable
)
```

### Error Handling Flow

```
Error Occurs
    ↓
Classify Error Type
    ↓
Is Retryable? ──No──→ Return Error
    ↓ Yes
Check Retry Attempts
    ↓
Attempts Left? ──No──→ Return Error
    ↓ Yes
Calculate Backoff Delay
    ↓
Wait (with jitter)
    ↓
Retry Request
```

### Error Wrapping

```go
// Wrap errors to preserve context
return &InstanaError{
    Type:       ErrorTypeNetwork,
    Message:    "failed to connect",
    StatusCode: 0,
    Err:        originalError, // Wrapped error
}

// Unwrap for error checking
if errors.Is(err, context.Canceled) {
    // Handle cancellation
}
```

---

## Performance Optimizations

### 1. Connection Pooling

```go
// HTTP Transport configuration
transport := &http.Transport{
    MaxIdleConns:        100,  // Total idle connections
    MaxIdleConnsPerHost: 10,   // Per-host idle connections
    MaxConnsPerHost:     10,   // Per-host total connections
    IdleConnTimeout:     90 * time.Second,
    DisableKeepAlives:   false, // Reuse connections
}
```

**Benefits:**
- Reuses TCP connections
- Reduces connection overhead
- Improves throughput
- Lower latency

### 2. Lazy Initialization

```go
// Resource clients created only when accessed
func (c *instanaAPIImpl) APITokens() rest.RestResource[*api.APIToken] {
    c.once.apiTokens.Do(func() {
        c.apiTokensClient = createAPITokenClient(c.restClient)
    })
    return c.apiTokensClient
}
```

**Benefits:**
- Faster startup time
- Lower memory usage
- Only pay for what you use
- Thread-safe initialization

### 3. Rate Limiting

```go
// Token bucket prevents API throttling
rateLimiter := NewRateLimiter(RateLimitConfig{
    RequestsPerSecond: 100,
    BurstCapacity:     200,
})
```

**Benefits:**
- Prevents API throttling
- Smooth request distribution
- Allows burst traffic
- Protects backend

### 4. Efficient JSON Processing

```go
// Streaming JSON unmarshalling
decoder := json.NewDecoder(response.Body)
var result []APIToken
if err := decoder.Decode(&result); err != nil {
    return nil, err
}
```

**Benefits:**
- Lower memory allocation
- Faster processing
- Handles large responses
- Streaming support

### 5. Context Propagation

```go
// Respect context cancellation
func (r *Retryer) Do(ctx context.Context, fn func() error) error {
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
        return fn()
    }
}
```

**Benefits:**
- Graceful cancellation
- Timeout enforcement
- Resource cleanup
- Better control flow

---

## Thread Safety

### Concurrent Access

All public APIs are thread-safe:

1. **Lazy Initialization** - Uses `sync.Once`
2. **Rate Limiter** - Uses `sync.Mutex`
3. **HTTP Client** - Thread-safe by design
4. **Configuration** - Immutable after creation

### Example: Safe Concurrent Usage

```go
api := instana.NewInstanaAPI("token", "tenant.instana.io", false)

// Safe to call from multiple goroutines
var wg sync.WaitGroup
for i := 0; i < 10; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        tokens, _ := api.APITokens().GetAll()
        // Process tokens
    }()
}
wg.Wait()
```

---

## Testing Strategy

### 1. Unit Tests

- Test individual components in isolation
- Mock dependencies using interfaces
- Focus on business logic

### 2. Integration Tests

- Test component interactions
- Use test HTTP server
- Verify end-to-end flows

### 3. Mock Generation

```bash
# Generate mocks using mockgen
go generate ./...
```

### 4. Test Utilities

**Location:** [`testutils/`](testutils/)

- Mock HTTP server
- Test fixtures
- Helper functions

---

## Extension Points

### 1. Custom Logger

Implement the [`Logger`](config/logger.go) interface:

```go
type CustomLogger struct{}

func (l *CustomLogger) Debug(msg string, keysAndValues ...interface{}) {
    // Custom implementation
}
```

### 2. Custom HTTP Client

Provide custom `http.Client`:

```go
config.HTTPClient = &http.Client{
    Transport: customTransport,
    Timeout:   customTimeout,
}
```

### 3. Custom Retry Strategy

Implement custom retry logic:

```go
retryer := NewRetryer(RetryConfig{
    ShouldRetry: func(err error) bool {
        // Custom retry logic
    },
})
```

---

## Future Enhancements

### Planned Features

1. **Circuit Breaker** - Prevent cascading failures
2. **Request Tracing** - Distributed tracing support
3. **Metrics Collection** - Built-in metrics
4. **Caching Layer** - Response caching
5. **Batch Operations** - Efficient bulk operations
6. **Webhook Support** - Event notifications

### Extensibility

The architecture supports future enhancements through:

- Interface-based design
- Pluggable components
- Configuration flexibility
- Minimal breaking changes

---

## See Also

- [API Reference](API_REFERENCE.md) - Complete API documentation
- [README](README.md) - Project overview
- [Quick Start](QUICK_START.md) - Getting started guide
- [Contributing](CONTRIBUTING.md) - Contribution guidelines