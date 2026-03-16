# Instana Go Client - Refactoring Implementation Guide

## Status: Phase 1 & 2 Complete ✅

This guide provides the complete implementation details for the package refactoring.

---

## Completed Work

### Phase 1: Shared Packages (3 packages, 12 files, 856 lines)

#### 1. `shared/types/` - Common Type Definitions
- ✅ `severity.go` - Severity levels (Warning, Critical)
- ✅ `granularity.go` - Time granularity constants
- ✅ `operator.go` - Logical and expression operators
- ✅ `aggregation.go` - Aggregation types
- ✅ `threshold.go` - Threshold definitions
- ✅ `doc.go` - Package documentation

#### 2. `shared/tagfilter/` - Tag Filter Logic
- ✅ `filter.go` - Tag filter expressions
- ✅ `doc.go` - Package documentation

#### 3. `shared/rest/` - REST Resource Abstractions
- ✅ `resource.go` - Core interfaces (RestResource, ReadOnlyRestResource, RestClient, JSONUnmarshaller)
- ✅ `readonly.go` - Read-only REST resource implementation
- ✅ `default.go` - Full CRUD REST resource with 5 factory functions
- ✅ `doc.go` - Package documentation

### Phase 2: API Package Template

#### `api/apitoken/` - Template for All API Packages
- ✅ `model.go` - Data model with GetIDForResourcePath() method
- ✅ `constants.go` - ResourcePath constant
- ✅ `unmarshaller.go` - JSON unmarshaller implementation
- ✅ `client.go` - NewClient() factory function
- ✅ `doc.go` - Package documentation with examples

---

## Phase 3: Remaining 27 API Packages

### Template Pattern (5 files per package)

Each API package follows this exact structure:

```
api/{packagename}/
├── model.go          # Data structures
├── constants.go      # Resource path
├── unmarshaller.go   # JSON handling
├── client.go         # Factory function
└── doc.go           # Documentation
```

### API Packages to Create

| # | Package Name | Source File | REST Method | Notes |
|---|--------------|-------------|-------------|-------|
| 1 | ✅ apitoken | api-tokens-api.go | POST/PUT | Template created |
| 2 | alertingchannel | alerting-channels-api.go | POST/PUT | |
| 3 | alertingconfig | alerts-api.go | POST/PUT | |
| 4 | applicationalertconfig | application-alert-config.go | POST/PUT | |
| 5 | applicationconfig | application-configs-api.go | POST/PUT | |
| 6 | automationaction | automation-action-api.go | POST/PUT | |
| 7 | automationpolicy | automation-policy-api.go | POST/PUT | |
| 8 | builtineventspec | builtin-event-specification-api.go | Read-only | |
| 9 | customdashboard | custom-dashboard.go | POST/PUT | |
| 10 | customeventspec | custom-event-specficiations-api.go | POST/PUT | |
| 11 | group | groups-api.go | POST/PUT | RBAC |
| 12 | hostagent | host-agent-api.go | Read-only | |
| 13 | infraalertconfig | infra-alert-config-api.go | POST/PUT | |
| 14 | logalertconfig | log-alert-config-api.go | POST/PUT | |
| 15 | maintenancewindow | maintenance-window-api.go | POST/PUT | |
| 16 | mobilealertconfig | mobile-alert-config.go | POST/PUT | |
| 17 | role | roles-api.go | POST/PUT | RBAC |
| 18 | sliconfig | sli-config-api.go | POST/PUT | |
| 19 | sloalertconfig | slo-alert-config-api.go | POST/PUT | |
| 20 | sloconfig | slo-config-api.go | POST/PUT | |
| 21 | slocorrection | slo-correction-config-api.go | POST/PUT | |
| 22 | syntheticalertconfig | synthetic-alert-config-api.go | POST/PUT | |
| 23 | syntheticlocation | synthetic-location-api.go | Read-only | |
| 24 | synthetictest | synthetic-test-api.go | POST/PUT | |
| 25 | team | teams-api.go | POST/PUT | RBAC |
| 26 | user | user-api.go | Read-only | RBAC |
| 27 | websitealertconfig | website-alert-config.go | POST/PUT | |
| 28 | websitemonitoring | website-monitoring-config-api.go | POST/PUT | |

---

## Implementation Steps for Each Package

### Step 1: Identify Source Files

For each API, locate these files in `instana/`:
- `{name}-api.go` - Contains API client method and resource path
- `{name}.go` - Contains data model(s)
- `{name}_test.go` - Contains tests (if exists)

### Step 2: Create Package Directory

```bash
mkdir -p api/{packagename}
```

### Step 3: Create model.go

Extract the data model from the source file:

```go
package {packagename}

// {ModelName} is the representation of {description}
type {ModelName} struct {
    // Copy all fields from original struct
}

// GetIDForResourcePath implementation of the interface InstanaDataObject
func (r *{ModelName}) GetIDForResourcePath() string {
    return r.{IDField}
}
```

### Step 4: Create constants.go

Extract the resource path constant:

```go
package {packagename}

// ResourcePath is the path to {description} resource of Instana RESTful API
const ResourcePath = "{actual-path}"
```

### Step 5: Create unmarshaller.go

```go
package {packagename}

import (
    "encoding/json"
    "fmt"
)

// NewUnmarshaller creates a new instance of a JSONUnmarshaller for {ModelName}
func NewUnmarshaller() *Unmarshaller {
    return &Unmarshaller{}
}

// Unmarshaller is a JSONUnmarshaller implementation for {ModelName}
type Unmarshaller struct{}

// Unmarshal converts JSON bytes into a {ModelName}
func (u *Unmarshaller) Unmarshal(data []byte) (*{ModelName}, error) {
    var target {ModelName}
    if err := json.Unmarshal(data, &target); err != nil {
        return &target, fmt.Errorf("failed to parse json; %s", err)
    }
    return &target, nil
}

// UnmarshalArray converts JSON bytes into a slice of {ModelName}
func (u *Unmarshaller) UnmarshalArray(data []byte) (*[]*{ModelName}, error) {
    var target []*{ModelName}
    if err := json.Unmarshal(data, &target); err != nil {
        return &target, fmt.Errorf("failed to parse json; %s", err)
    }
    return &target, nil
}
```

### Step 6: Create client.go

Determine the REST method from the original API implementation:

**For POST/PUT resources:**
```go
package {packagename}

import "github.com/instana/instana-go-client/shared/rest"

// NewClient creates a new API client for {description}
func NewClient(restClient rest.RestClient) rest.RestResource[*{ModelName}] {
    return rest.NewCreatePOSTUpdatePUTRestResource(
        ResourcePath,
        NewUnmarshaller(),
        restClient,
    )
}
```

**For Read-only resources:**
```go
package {packagename}

import "github.com/instana/instana-go-client/shared/rest"

// NewClient creates a new read-only API client for {description}
func NewClient(restClient rest.RestClient) rest.ReadOnlyRestResource[*{ModelName}] {
    return rest.NewReadOnlyRestResource(
        ResourcePath,
        NewUnmarshaller(),
        restClient,
    )
}
```

### Step 7: Create doc.go

```go
// Package {packagename} provides the API client for managing {description}.
//
// {Brief description of what this API manages}
//
// Example usage:
//
//  // Create a new client
//  client := {packagename}.NewClient(restClient)
//
//  // Get all items
//  items, err := client.GetAll()
//  if err != nil {
//      // handle error
//  }
//
//  // Get a specific item
//  item, err := client.GetOne("item-id")
//  if err != nil {
//      // handle error
//  }
//
//  // Create a new item
//  newItem := &{packagename}.{ModelName}{
//      // Set fields
//  }
//  created, err := client.Create(newItem)
//  if err != nil {
//      // handle error
//  }
//
//  // Update an existing item
//  item.{Field} = "new value"
//  updated, err := client.Update(item)
//  if err != nil {
//      // handle error
//  }
//
//  // Delete an item
//  err = client.DeleteByID("item-id")
//  if err != nil {
//      // handle error
//  }
package {packagename}
```

---

## Phase 4: Client Package

Create `client/` package with the main InstanaAPI interface:

### client/interface.go

```go
package client

import (
    "github.com/instana/instana-go-client/shared/rest"
    "github.com/instana/instana-go-client/api/apitoken"
    "github.com/instana/instana-go-client/api/alertingchannel"
    // ... import all 28 API packages
)

// InstanaAPI is the main interface for interacting with the Instana API
type InstanaAPI interface {
    APITokens() rest.RestResource[*apitoken.APIToken]
    AlertingChannels() rest.RestResource[*alertingchannel.AlertingChannel]
    // ... methods for all 28 APIs
}
```

### client/client.go

```go
package client

import (
    "github.com/instana/instana-go-client/shared/rest"
    "github.com/instana/instana-go-client/api/apitoken"
    // ... import all API packages
)

type instanaAPI struct {
    restClient rest.RestClient
    
    // Lazy-initialized API clients
    apiTokens rest.RestResource[*apitoken.APIToken]
    // ... fields for all 28 APIs
}

// NewInstanaAPI creates a new Instana API client
func NewInstanaAPI(restClient rest.RestClient) InstanaAPI {
    return &instanaAPI{
        restClient: restClient,
    }
}

// APITokens returns the API tokens client (lazy initialization)
func (api *instanaAPI) APITokens() rest.RestResource[*apitoken.APIToken] {
    if api.apiTokens == nil {
        api.apiTokens = apitoken.NewClient(api.restClient)
    }
    return api.apiTokens
}

// ... implement all 28 API methods with lazy initialization
```

---

## Phase 5: Update instana Package

The `instana/` package becomes minimal, containing ONLY:

### instana/client.go (NEW - minimal initialization)

```go
package instana

import (
    "github.com/instana/instana-go-client/client"
    "github.com/instana/instana-go-client/shared/rest"
)

// NewInstanaAPI creates a new Instana API client with default configuration
func NewInstanaAPI(apiToken string, host string) (client.InstanaAPI, error) {
    config := DefaultClientConfig()
    config.APIToken = apiToken
    config.Host = host
    
    restClient, err := NewRestClient(config)
    if err != nil {
        return nil, err
    }
    
    return client.NewInstanaAPI(restClient), nil
}

// NewInstanaAPIWithConfig creates a new Instana API client with custom configuration
func NewInstanaAPIWithConfig(config *ClientConfig) (client.InstanaAPI, error) {
    if err := config.Validate(); err != nil {
        return nil, err
    }
    
    restClient, err := NewRestClient(config)
    if err != nil {
        return nil, err
    }
    
    return client.NewInstanaAPI(restClient), nil
}
```

### Keep in instana/ package:
- `config.go` - Client configuration
- `config_builder.go` - Configuration builder
- `config_loader.go` - Load config from env/file
- `config_validator.go` - Configuration validation
- `rest-client.go` - HTTP client implementation
- `errors.go` - Error types
- All test files for the above

### Remove from instana/ package:
- All `*-api.go` files (migrated to `api/` packages)
- All model files (migrated to `api/` packages)
- All API-specific constants (migrated to `api/` packages)

---

## Phase 6: Update Imports

Update all import statements across the codebase:

### Before:
```go
import "github.com/instana/instana-go-client/instana"

token := &instana.APIToken{...}
client := instana.NewInstanaAPI(...)
tokens := client.APITokens()
```

### After:
```go
import (
    "github.com/instana/instana-go-client/instana"
    "github.com/instana/instana-go-client/api/apitoken"
)

token := &apitoken.APIToken{...}
client, _ := instana.NewInstanaAPI(...)
tokens := client.APITokens()
```

---

## Verification Checklist

For each API package, verify:

- [ ] Model struct matches original exactly
- [ ] GetIDForResourcePath() returns correct field
- [ ] ResourcePath constant matches original
- [ ] Unmarshaller handles both single and array
- [ ] Client uses correct REST method (POST/PUT or Read-only)
- [ ] Documentation is clear and includes examples
- [ ] No logic changes - pure refactoring only

---

## Benefits of New Structure

1. **Discoverability**: Find all code for an API in one place
2. **Maintainability**: Changes isolated to single package
3. **Testability**: Easy to test each API independently
4. **Clarity**: Clear separation of concerns
5. **Scalability**: Easy to add new APIs
6. **Documentation**: Each package self-documented
7. **Type Safety**: Strong typing with generics
8. **Performance**: Lazy initialization of API clients

---

## Next Steps

1. Create remaining 27 API packages following the template
2. Create the `client/` package with InstanaAPI interface
3. Update `instana/` package to minimal initialization
4. Update all imports across codebase
5. Run tests and fix any issues
6. Update documentation and examples

---

## Notes

- This is **pure refactoring** - no logic changes
- All existing functionality preserved
- Backward compatibility NOT maintained (new project)
- Follow Go naming conventions strictly
- Keep packages focused and cohesive