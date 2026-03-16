# Phase 10: Generic Client Factory & Unmarshaller Consolidation

## Overview
Proposal to consolidate repetitive client factory code and unmarshallers from individual API packages into a centralized, generic implementation in the `client` package. This will eliminate code duplication across 28 API packages and improve maintainability.

## Current State Analysis

### Current Pattern (Repeated 28 Times)

Each API package currently contains:

1. **client.go** - Factory function (14 lines)
```go
package apitoken

func NewClient(restClient rest.RestClient) rest.RestResource[*APIToken] {
    return rest.NewCreatePOSTUpdatePUTRestResource(
        ResourcePath,
        NewUnmarshaller(),
        restClient,
    )
}
```

2. **unmarshaller.go** - JSON unmarshaller (34 lines)
```go
package apitoken

type Unmarshaller struct{}

func NewUnmarshaller() *Unmarshaller {
    return &Unmarshaller{}
}

func (u *Unmarshaller) Unmarshal(data []byte) (*APIToken, error) {
    var target APIToken
    if err := json.Unmarshal(data, &target); err != nil {
        return &target, fmt.Errorf("failed to parse json; %s", err)
    }
    return &target, nil
}

func (u *Unmarshaller) UnmarshalArray(data []byte) (*[]*APIToken, error) {
    var target []*APIToken
    if err := json.Unmarshal(data, &target); err != nil {
        return &target, fmt.Errorf("failed to parse json; %s", err)
    }
    return &target, nil
}
```

### Problems with Current Approach

1. **Code Duplication**: 28 identical unmarshaller implementations (952 lines total)
2. **Maintenance Burden**: Any change requires updating 28 files
3. **Inconsistency Risk**: Easy to have slight variations between packages
4. **Boilerplate**: Each package has 48 lines of repetitive code
5. **Testing Overhead**: Need to test the same logic 28 times

## Proposed Solution

### 1. Generic Unmarshaller in `shared/rest`

Create a single generic unmarshaller that works for all types:

```go
// shared/rest/unmarshaller.go
package rest

import (
    "encoding/json"
    "fmt"
)

// GenericUnmarshaller is a generic JSON unmarshaller for any type
type GenericUnmarshaller[T any] struct{}

// NewGenericUnmarshaller creates a new generic unmarshaller
func NewGenericUnmarshaller[T any]() *GenericUnmarshaller[T] {
    return &GenericUnmarshaller[T]{}
}

// Unmarshal converts JSON bytes into the target type
func (u *GenericUnmarshaller[T]) Unmarshal(data []byte) (T, error) {
    var target T
    if err := json.Unmarshal(data, &target); err != nil {
        var zero T
        return zero, fmt.Errorf("failed to parse json; %s", err)
    }
    return target, nil
}

// UnmarshalArray converts JSON bytes into a slice of the target type
func (u *GenericUnmarshaller[T]) UnmarshalArray(data []byte) (*[]T, error) {
    var target []T
    if err := json.Unmarshal(data, &target); err != nil {
        return &target, fmt.Errorf("failed to parse json; %s", err)
    }
    return &target, nil
}
```

### 2. Centralized Client Factory in `client` Package

Move all client creation logic to the `client` package:

```go
// client/factory.go
package client

import (
    "github.com/instana/instana-go-client/shared/rest"
)

// ClientFactory provides methods to create REST resource clients
type ClientFactory struct {
    restClient rest.RestClient
}

// NewClientFactory creates a new client factory
func NewClientFactory(restClient rest.RestClient) *ClientFactory {
    return &ClientFactory{restClient: restClient}
}

// NewRestResource creates a REST resource with the specified mode
func (f *ClientFactory) NewRestResource[T rest.InstanaDataObject](
    resourcePath string,
    mode rest.DefaultRestResourceMode,
) rest.RestResource[T] {
    unmarshaller := rest.NewGenericUnmarshaller[T]()
    
    switch mode {
    case rest.DefaultRestResourceModeCreateAndUpdatePUT:
        return rest.NewCreatePUTUpdatePUTRestResource(resourcePath, unmarshaller, f.restClient)
    case rest.DefaultRestResourceModeCreatePOSTUpdatePUT:
        return rest.NewCreatePOSTUpdatePUTRestResource(resourcePath, unmarshaller, f.restClient)
    case rest.DefaultRestResourceModeCreateAndUpdatePOST:
        return rest.NewCreatePOSTUpdatePOSTRestResource(resourcePath, unmarshaller, f.restClient)
    case rest.DefaultRestResourceModeCreatePOSTAndUpdateNotSupported:
        return rest.NewCreatePOSTUpdateNotSupportedRestResource(resourcePath, unmarshaller, f.restClient)
    case rest.DefaultRestResourceModeCreatePUTAndUpdateNotSupported:
        return rest.NewCreatePUTUpdateNotSupportedRestResource(resourcePath, unmarshaller, f.restClient)
    default:
        // Default to POST for create, PUT for update
        return rest.NewCreatePOSTUpdatePUTRestResource(resourcePath, unmarshaller, f.restClient)
    }
}

// NewReadOnlyRestResource creates a read-only REST resource
func (f *ClientFactory) NewReadOnlyRestResource[T rest.InstanaDataObject](
    resourcePath string,
) rest.ReadOnlyRestResource[T] {
    unmarshaller := rest.NewGenericUnmarshaller[T]()
    return rest.NewReadOnlyRestResource(resourcePath, unmarshaller, f.restClient)
}
```

### 3. Updated `client/client.go` Implementation

Simplify the client implementation using the factory:

```go
// client/client.go
package client

import (
    "github.com/instana/instana-go-client/api/apitoken"
    "github.com/instana/instana-go-client/api/alertingchannel"
    // ... other imports
    "github.com/instana/instana-go-client/shared/rest"
)

type instanaAPI struct {
    restClient rest.RestClient
    factory    *ClientFactory
    
    // Lazy-initialized API clients
    apiTokens rest.RestResource[*apitoken.APIToken]
    // ... other fields
}

func NewInstanaAPI(restClient rest.RestClient) InstanaAPI {
    return &instanaAPI{
        restClient: restClient,
        factory:    NewClientFactory(restClient),
    }
}

func (api *instanaAPI) APITokens() rest.RestResource[*apitoken.APIToken] {
    if api.apiTokens == nil {
        api.apiTokens = api.factory.NewRestResource[*apitoken.APIToken](
            apitoken.ResourcePath,
            rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
        )
    }
    return api.apiTokens
}

// Similar pattern for all other resources...
```

### 4. Simplified API Packages

Each API package would only need:

```
api/apitoken/
├── constants.go    # ResourcePath and other constants
├── model.go        # APIToken struct with GetIDForResourcePath()
└── doc.go          # Package documentation
```

**No more client.go or unmarshaller.go files needed!**

## Benefits

### 1. **Massive Code Reduction**
- **Before**: 28 packages × 48 lines = 1,344 lines of boilerplate
- **After**: ~150 lines in shared/rest + client/factory.go
- **Savings**: ~1,200 lines of code eliminated

### 2. **Single Source of Truth**
- One unmarshaller implementation to maintain
- One client factory pattern
- Changes apply to all resources automatically

### 3. **Type Safety**
- Go generics ensure compile-time type safety
- No runtime reflection needed
- Full IDE autocomplete support

### 4. **Easier Testing**
- Test generic unmarshaller once
- Test factory once
- No need to test 28 identical implementations

### 5. **Better Maintainability**
- Clear separation: API packages define models, client package handles wiring
- Easier to add new resources (just add model + constants)
- Consistent behavior across all resources

### 6. **Flexibility**
- Easy to add custom unmarshallers for special cases
- Can override factory behavior when needed
- Backward compatible with existing code

## Migration Strategy

### Phase 10.1: Create Generic Infrastructure
1. Add `GenericUnmarshaller` to `shared/rest/unmarshaller.go`
2. Add `ClientFactory` to `client/factory.go`
3. Add tests for both
4. Verify compilation

### Phase 10.2: Update Client Package
1. Update `client/client.go` to use factory
2. Remove individual `NewClient()` calls from API packages
3. Inline resource creation in lazy initialization
4. Run tests to ensure no regressions

### Phase 10.3: Clean Up API Packages
1. Remove `client.go` from all 28 API packages
2. Remove `unmarshaller.go` from all 28 API packages
3. Keep only `constants.go`, `model.go`, and `doc.go`
4. Update any package documentation

### Phase 10.4: Verification
1. Run full test suite
2. Verify all 28 resources still work
3. Check for any breaking changes
4. Update documentation

## File Changes Summary

### Files to Create
- `shared/rest/unmarshaller.go` (~40 lines)
- `shared/rest/unmarshaller_test.go` (~100 lines)
- `client/factory.go` (~80 lines)
- `client/factory_test.go` (~150 lines)

### Files to Modify
- `client/client.go` - Update all lazy initialization methods
- `client/interface.go` - No changes needed (interface stays same)

### Files to Delete (56 files total)
- `api/*/client.go` (28 files)
- `api/*/unmarshaller.go` (28 files)

### Files to Keep
- `api/*/constants.go` (28 files)
- `api/*/model.go` (28 files)
- `api/*/doc.go` (28 files)

## Example: Before vs After

### Before (apitoken package)
```
api/apitoken/
├── client.go          # 14 lines - DELETE
├── constants.go       # 8 lines - KEEP
├── doc.go            # 5 lines - KEEP
├── model.go          # 45 lines - KEEP
└── unmarshaller.go   # 34 lines - DELETE
```

### After (apitoken package)
```
api/apitoken/
├── constants.go      # 8 lines - KEEP
├── doc.go           # 5 lines - KEEP
└── model.go         # 45 lines - KEEP
```

**Result**: 48 lines removed per package × 28 packages = 1,344 lines eliminated!

## Backward Compatibility

### For External Consumers
✅ **No breaking changes** - The `InstanaAPI` interface remains unchanged:
```go
api := instana.NewInstanaAPI(token, endpoint, false)
tokens, err := api.APITokens().GetAll()  // Still works!
```

### For Internal Code
✅ **No changes needed** - All Terraform provider code continues to work as-is

### For Direct Package Users
⚠️ **Minor breaking change** - If anyone imports individual API packages directly:
```go
// Before (will break)
import "github.com/instana/instana-go-client/api/apitoken"
client := apitoken.NewClient(restClient)

// After (use client package instead)
import "github.com/instana/instana-go-client/client"
api := client.NewInstanaAPI(restClient)
tokens := api.APITokens()
```

**Mitigation**: Since this is a new project not yet used by anyone, this is acceptable.

## Performance Considerations

### Generic Unmarshaller Performance
- **No performance impact**: Generics are compile-time, no runtime overhead
- **Same JSON parsing**: Uses standard `encoding/json` package
- **Memory efficiency**: No additional allocations compared to current approach

### Factory Pattern Performance
- **Negligible overhead**: Factory just calls existing constructors
- **Lazy initialization**: Resources created only when needed (already implemented)
- **No reflection**: All type information resolved at compile time

## Testing Strategy

### Unit Tests
```go
// Test generic unmarshaller
func TestGenericUnmarshaller_Unmarshal(t *testing.T)
func TestGenericUnmarshaller_UnmarshalArray(t *testing.T)
func TestGenericUnmarshaller_InvalidJSON(t *testing.T)

// Test factory
func TestClientFactory_NewRestResource(t *testing.T)
func TestClientFactory_AllModes(t *testing.T)
func TestClientFactory_ReadOnlyResource(t *testing.T)
```

### Integration Tests
- Verify all 28 resources still work with real API calls
- Test CRUD operations for each resource type
- Ensure error handling works correctly

## Risks & Mitigation

### Risk 1: Generic Type Constraints
**Risk**: Some types might not work with generic unmarshaller
**Mitigation**: Keep ability to provide custom unmarshallers for special cases

### Risk 2: Breaking Changes
**Risk**: External users might depend on individual package clients
**Mitigation**: Document migration path, provide deprecation warnings

### Risk 3: Testing Coverage
**Risk**: Removing 28 test files might reduce coverage
**Mitigation**: Ensure generic tests cover all edge cases, add integration tests

## Recommendation

✅ **Strongly Recommended** - This refactoring provides significant benefits:
- Eliminates 1,200+ lines of boilerplate code
- Improves maintainability dramatically
- No performance impact
- Minimal risk (new project, not yet in production use)
- Aligns with Go best practices (DRY principle, generics usage)

## Next Steps

1. **Get approval** for this refactoring approach
2. **Implement Phase 10.1**: Create generic infrastructure
3. **Implement Phase 10.2**: Update client package
4. **Implement Phase 10.3**: Clean up API packages
5. **Implement Phase 10.4**: Verification and documentation

---
*Proposal Date: 2026-03-13*
*Estimated Effort: 4-6 hours*
*Risk Level: Low*
*Impact: High (positive)*