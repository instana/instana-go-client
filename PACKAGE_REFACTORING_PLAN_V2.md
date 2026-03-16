# Instana Go Client - Package Refactoring Plan v2

## Executive Summary

This document outlines a comprehensive refactoring plan for the `instana-go-client` repository with a **granular, self-contained package structure**. Each API endpoint gets its own package containing all related code: API logic, models, constants, and tests. This approach maximizes maintainability, discoverability, and independence.

**Key Principles:**
- **One package per API endpoint** - Each API resource is completely self-contained
- **Colocated code** - Models, constants, and tests live with the API implementation
- **Minimal shared dependencies** - Only truly common utilities are shared
- **Clear naming** - Package names directly match API resources
- **Easy navigation** - Find everything related to an API in one place

---

## Table of Contents

1. [Current State Analysis](#1-current-state-analysis)
2. [Proposed Package Structure](#2-proposed-package-structure)
3. [Package Organization Principles](#3-package-organization-principles)
4. [Detailed Package Breakdown](#4-detailed-package-breakdown)
5. [Shared Packages](#5-shared-packages)
6. [Migration Strategy](#6-migration-strategy)
7. [Implementation Roadmap](#7-implementation-roadmap)
8. [Success Criteria](#8-success-criteria)

---

## 1. Current State Analysis

### 1.1 Problems with Current Structure

1. **Monolithic `instana/` Package**: 100+ files in single package
2. **Mixed Concerns**: API logic, models, and utilities intermingled
3. **Poor Discoverability**: Hard to find specific API implementations
4. **Difficult Testing**: Hard to test components in isolation
5. **Tight Coupling**: Changes ripple across unrelated code

### 1.2 Identified API Resources

Based on Terraform provider and current code:

- Application Alert Config
- Application Config
- API Token
- Alerting Channel
- Alerting Configuration
- Automation Action
- Automation Policy
- Builtin Event Specification
- Custom Dashboard
- Custom Event Specification
- Group (RBAC)
- Host Agent
- Infrastructure Alert Config
- Log Alert Config
- Maintenance Window Config
- Mobile Alert Config
- Role (RBAC)
- SLI Config
- SLO Alert Config
- SLO Config
- SLO Correction Config
- Synthetic Alert Config
- Synthetic Location
- Synthetic Test
- Team (RBAC)
- User (RBAC)
- Website Alert Config
- Website Monitoring Config

---

## 2. Proposed Package Structure

### 2.1 High-Level Structure

```
instana-go-client/
├── api/                                    # All API endpoint packages
│   ├── applicationalertconfig/            # Application alert configurations
│   │   ├── client.go                      # API client implementation
│   │   ├── models.go                      # Data models specific to this API
│   │   ├── constants.go                   # API-specific constants
│   │   ├── client_test.go                 # Tests
│   │   ├── models_test.go
│   │   └── doc.go                         # Package documentation
│   │
│   ├── applicationconfig/                 # Application configurations
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   ├── models_test.go
│   │   └── doc.go
│   │
│   ├── apitoken/                          # API token management
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── alertingchannel/                   # Alerting channels
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── alertingconfig/                    # Alerting configurations
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── automationaction/                  # Automation actions
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── automationpolicy/                  # Automation policies
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── builtineventspec/                  # Builtin event specifications
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── customdashboard/                   # Custom dashboards
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── customeventspec/                   # Custom event specifications
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── group/                             # RBAC groups
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── hostagent/                         # Host agents
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── infraalertconfig/                  # Infrastructure alert configs
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── logalertconfig/                    # Log alert configurations
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── maintenancewindow/                 # Maintenance windows
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── mobilealertconfig/                 # Mobile alert configurations
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── role/                              # RBAC roles
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── sliconfig/                         # SLI configurations
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── sloalertconfig/                    # SLO alert configurations
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── sloconfig/                         # SLO configurations
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── slocorrection/                     # SLO corrections
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── syntheticalertconfig/              # Synthetic alert configs
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── syntheticlocation/                 # Synthetic locations
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── synthetictest/                     # Synthetic tests
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── team/                              # RBAC teams
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── user/                              # RBAC users
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   ├── websitealertconfig/                # Website alert configurations
│   │   ├── client.go
│   │   ├── models.go
│   │   ├── constants.go
│   │   ├── client_test.go
│   │   └── doc.go
│   │
│   └── websitemonitoring/                 # Website monitoring configs
│       ├── client.go
│       ├── models.go
│       ├── constants.go
│       ├── client_test.go
│       └── doc.go
│
├── shared/                                 # Truly shared utilities
│   ├── tagfilter/                         # Tag filter (used by many APIs)
│   │   ├── filter.go
│   │   ├── parser.go
│   │   ├── filter_test.go
│   │   └── doc.go
│   │
│   ├── types/                             # Common types used across APIs
│   │   ├── severity.go                    # Alert severity levels
│   │   ├── granularity.go                 # Time granularity
│   │   ├── operator.go                    # Comparison operators
│   │   ├── aggregation.go                 # Aggregation types
│   │   ├── threshold.go                   # Threshold definitions
│   │   ├── types_test.go
│   │   └── doc.go
│   │
│   └── rest/                              # REST resource patterns
│       ├── resource.go                    # Generic REST resource
│       ├── readonly.go                    # Read-only resource
│       ├── resource_test.go
│       └── doc.go
│
├── client/                                 # HTTP client (existing)
│   ├── client.go
│   ├── retry.go
│   ├── rate_limiter.go
│   └── client_test.go
│
├── config/                                 # Configuration (existing)
│   ├── config.go
│   ├── builder.go
│   ├── loader.go
│   └── validator.go
│
├── errors/                                 # Error handling
│   ├── errors.go
│   └── errors_test.go
│
├── instana/                                # Main package (facade)
│   ├── api.go                             # Main API interface
│   ├── client.go                          # Client factory
│   ├── backward_compat.go                 # Compatibility layer
│   └── doc.go
│
├── utils/                                  # Public utilities
│   ├── string.go
│   ├── int.go
│   └── slice.go
│
├── testutils/                              # Test utilities
│   ├── mock_server.go
│   └── fixtures.go
│
├── examples/                               # Usage examples
│   ├── basic/
│   ├── alerts/
│   └── ...
│
├── docs/                                   # Documentation
│   ├── architecture.md
│   └── migration_guide.md
│
├── go.mod
├── go.sum
├── README.md
└── LICENSE
```

---

## 3. Package Organization Principles

### 3.1 Self-Contained Packages

Each API package is **completely self-contained**:

```
api/applicationconfig/
├── client.go          # API client with CRUD operations
├── models.go          # All data models for this API
├── constants.go       # API endpoints, defaults, enums
├── client_test.go     # Client tests
├── models_test.go     # Model tests
└── doc.go            # Package documentation
```

**Benefits**:
- ✅ Everything related to one API in one place
- ✅ Easy to find and modify
- ✅ Can be developed/tested independently
- ✅ Clear ownership and responsibility
- ✅ No confusion about where code belongs

### 3.2 Package Naming Convention

Package names should be:
- **Lowercase** - Go convention
- **Descriptive** - Match the API resource name
- **Singular** - `applicationconfig` not `applicationconfigs`
- **No underscores** - `applicationalertconfig` not `application_alert_config`

### 3.3 File Organization Within Package

**Standard files in each package**:

1. **`client.go`** - API client implementation
   ```go
   type Client struct {
       restClient *client.Client
       basePath   string
   }
   
   func NewClient(restClient *client.Client) *Client
   func (c *Client) Get(ctx context.Context, id string) (*Config, error)
   func (c *Client) Create(ctx context.Context, config *Config) (*Config, error)
   func (c *Client) Update(ctx context.Context, config *Config) (*Config, error)
   func (c *Client) Delete(ctx context.Context, id string) error
   ```

2. **`models.go`** - Data structures
   ```go
   type Config struct {
       ID          string
       Name        string
       Description string
       // ... other fields
   }
   
   type Rule struct {
       // ... rule fields
   }
   
   // Helper methods
   func (c *Config) Validate() error
   func (c *Config) GetID() string
   ```

3. **`constants.go`** - Constants and enums
   ```go
   const (
       BasePath = "/api/application-monitoring/settings/application"
       
       DefaultTimeout = 30 * time.Second
   )
   
   type Scope string
   
   const (
       ScopeIncludeNoDownstream Scope = "INCLUDE_NO_DOWNSTREAM"
       ScopeIncludeImmediateDownstream Scope = "INCLUDE_IMMEDIATE_DOWNSTREAM_DATABASE_AND_MESSAGING"
       ScopeIncludeAllDownstream Scope = "INCLUDE_ALL_DOWNSTREAM"
   )
   ```

4. **`client_test.go`** - Client tests
5. **`models_test.go`** - Model tests
6. **`doc.go`** - Package documentation

### 3.4 What Goes in Shared Packages?

**Only truly common code that is used by 3+ API packages:**

- **`shared/tagfilter/`** - Tag filter expressions (used by almost all APIs)
- **`shared/types/`** - Common types like Severity, Granularity, Operator
- **`shared/rest/`** - Generic REST resource patterns

**Rule of thumb**: If it's only used by 1-2 packages, keep it in those packages. Only extract to shared when it's genuinely reused.

---

## 4. Detailed Package Breakdown

### 4.1 Example: Application Alert Config Package

**Package**: `api/applicationalertconfig/`

**Files**:

#### `client.go`
```go
package applicationalertconfig

import (
    "context"
    "fmt"
    
    "github.com/instana/instana-go-client/client"
)

// Client provides operations for application alert configurations
type Client struct {
    restClient *client.Client
    basePath   string
}

// NewClient creates a new application alert config client
func NewClient(restClient *client.Client) *Client {
    return &Client{
        restClient: restClient,
        basePath:   BasePath,
    }
}

// Get retrieves an application alert configuration by ID
func (c *Client) Get(ctx context.Context, id string) (*Config, error) {
    var config Config
    err := c.restClient.Get(ctx, fmt.Sprintf("%s/%s", c.basePath, id), &config)
    return &config, err
}

// Create creates a new application alert configuration
func (c *Client) Create(ctx context.Context, config *Config) (*Config, error) {
    var result Config
    err := c.restClient.Post(ctx, c.basePath, config, &result)
    return &result, err
}

// Update updates an existing application alert configuration
func (c *Client) Update(ctx context.Context, config *Config) (*Config, error) {
    var result Config
    path := fmt.Sprintf("%s/%s", c.basePath, config.ID)
    err := c.restClient.Put(ctx, path, config, &result)
    return &result, err
}

// Delete deletes an application alert configuration
func (c *Client) Delete(ctx context.Context, id string) error {
    path := fmt.Sprintf("%s/%s", c.basePath, id)
    return c.restClient.Delete(ctx, path)
}

// List retrieves all application alert configurations
func (c *Client) List(ctx context.Context) ([]*Config, error) {
    var configs []*Config
    err := c.restClient.Get(ctx, c.basePath, &configs)
    return configs, err
}
```

#### `models.go`
```go
package applicationalertconfig

import (
    "github.com/instana/instana-go-client/shared/tagfilter"
    "github.com/instana/instana-go-client/shared/types"
)

// Config represents an application alert configuration
type Config struct {
    ID                  string                  `json:"id,omitempty"`
    Name                string                  `json:"name"`
    Description         string                  `json:"description,omitempty"`
    Enabled             bool                    `json:"enabled"`
    Triggering          bool                    `json:"triggering"`
    ApplicationID       string                  `json:"applicationId"`
    BoundaryScope       BoundaryScope           `json:"boundaryScope"`
    TagFilterExpression *tagfilter.Expression   `json:"tagFilterExpression,omitempty"`
    AlertChannels       map[types.Severity][]string `json:"alertChannels,omitempty"`
    Granularity         types.Granularity       `json:"granularity"`
    Rules               []RuleWithThreshold     `json:"rules"`
    TimeThreshold       *TimeThreshold          `json:"timeThreshold,omitempty"`
    CustomPayloadFields []CustomPayloadField    `json:"customPayloadFields,omitempty"`
}

// BoundaryScope defines the scope of the alert
type BoundaryScope string

const (
    BoundaryScopeDefault    BoundaryScope = "DEFAULT"
    BoundaryScopeInbound    BoundaryScope = "INBOUND"
    BoundaryScopeAll        BoundaryScope = "ALL"
)

// RuleWithThreshold represents an alert rule with thresholds
type RuleWithThreshold struct {
    Rule              *Rule                           `json:"rule"`
    ThresholdOperator types.Operator                  `json:"thresholdOperator"`
    Thresholds        map[types.Severity]Threshold    `json:"thresholds"`
}

// Rule represents an alert rule
type Rule struct {
    AlertType       AlertType           `json:"alertType"`
    MetricName      string              `json:"metricName,omitempty"`
    Aggregation     types.Aggregation   `json:"aggregation,omitempty"`
    MetricPattern   *MetricPattern      `json:"metricPattern,omitempty"`
}

// AlertType represents the type of alert
type AlertType string

const (
    AlertTypeApplicationError     AlertType = "applicationError"
    AlertTypeSlowness            AlertType = "slowness"
    AlertTypeErrorRate           AlertType = "errorRate"
    AlertTypeHTTPStatusCode      AlertType = "httpStatusCode"
    AlertTypeThroughput          AlertType = "throughput"
)

// Threshold represents a threshold value
type Threshold struct {
    Value    float64 `json:"value"`
    Operator types.Operator `json:"operator"`
}

// TimeThreshold represents time-based threshold configuration
type TimeThreshold struct {
    Type       TimeThresholdType `json:"type"`
    TimeWindow int64            `json:"timeWindow,omitempty"`
    Violations int32            `json:"violations,omitempty"`
}

// TimeThresholdType represents the type of time threshold
type TimeThresholdType string

const (
    TimeThresholdTypeImmediate TimeThresholdType = "IMMEDIATE"
    TimeThresholdTypeViolations TimeThresholdType = "VIOLATIONS_IN_SEQUENCE"
    TimeThresholdTypeRolling   TimeThresholdType = "VIOLATIONS_IN_PERIOD"
)

// MetricPattern represents a metric pattern for matching
type MetricPattern struct {
    Prefix   string `json:"prefix,omitempty"`
    Postfix  string `json:"postfix,omitempty"`
    Operator types.Operator `json:"operator"`
}

// CustomPayloadField represents a custom payload field
type CustomPayloadField struct {
    Type  string      `json:"type"`
    Key   string      `json:"key"`
    Value interface{} `json:"value"`
}

// Validate validates the configuration
func (c *Config) Validate() error {
    if c.Name == "" {
        return fmt.Errorf("name is required")
    }
    if c.ApplicationID == "" {
        return fmt.Errorf("applicationId is required")
    }
    if len(c.Rules) == 0 {
        return fmt.Errorf("at least one rule is required")
    }
    return nil
}

// GetID returns the configuration ID
func (c *Config) GetID() string {
    return c.ID
}
```

#### `constants.go`
```go
package applicationalertconfig

const (
    // BasePath is the API endpoint for application alert configurations
    BasePath = "/api/events/settings/application-alert-configs"
    
    // DefaultGranularity is the default time granularity
    DefaultGranularity = 600000 // 10 minutes in milliseconds
)
```

#### `doc.go`
```go
// Package applicationalertconfig provides a client for managing Instana application alert configurations.
//
// Application alert configurations define alerts that trigger based on application performance metrics
// such as error rates, response times, and throughput.
//
// Basic usage:
//
//   restClient, _ := client.NewClient(config)
//   alertClient := applicationalertconfig.NewClient(restClient)
//   
//   // Get an alert configuration
//   config, err := alertClient.Get(ctx, "alert-id")
//   if err != nil {
//       log.Fatal(err)
//   }
//   
//   // Create a new alert configuration
//   newConfig := &applicationalertconfig.Config{
//       Name:          "High Error Rate",
//       ApplicationID: "app-123",
//       Rules: []applicationalertconfig.RuleWithThreshold{
//           // ... rules
//       },
//   }
//   created, err := alertClient.Create(ctx, newConfig)
//
// For more examples, see the examples directory.
package applicationalertconfig
```

### 4.2 Package List with Descriptions

| Package | Description | API Endpoint |
|---------|-------------|--------------|
| `applicationalertconfig` | Application alert configurations | `/api/events/settings/application-alert-configs` |
| `applicationconfig` | Application monitoring configurations | `/api/application-monitoring/settings/application` |
| `apitoken` | API token management | `/api/settings/api-tokens` |
| `alertingchannel` | Alerting channel configurations | `/api/events/settings/alertingChannels` |
| `alertingconfig` | Alerting configurations | `/api/events/settings/alerting-configurations` |
| `automationaction` | Automation actions | `/api/automation/actions` |
| `automationpolicy` | Automation policies | `/api/automation/policies` |
| `builtineventspec` | Builtin event specifications (read-only) | `/api/events/settings/event-specifications/built-in` |
| `customdashboard` | Custom dashboard configurations | `/api/custom-dashboard` |
| `customeventspec` | Custom event specifications | `/api/events/settings/event-specifications/custom` |
| `group` | RBAC group management | `/api/settings/rbac/groups` |
| `hostagent` | Host agent information (read-only) | `/api/host-agent` |
| `infraalertconfig` | Infrastructure alert configurations | `/api/events/settings/global-alert-configs/infrastructure` |
| `logalertconfig` | Log alert configurations | `/api/events/settings/global-alert-configs/logs` |
| `maintenancewindow` | Maintenance window configurations | `/api/settings/maintenance-windows` |
| `mobilealertconfig` | Mobile app alert configurations | `/api/events/settings/mobile-app-alert-configs` |
| `role` | RBAC role management | `/api/settings/rbac/roles` |
| `sliconfig` | SLI configurations | `/api/application-monitoring/settings/sli/config` |
| `sloalertconfig` | SLO alert configurations | `/api/slo/config/alert` |
| `sloconfig` | SLO configurations | `/api/slo/config` |
| `slocorrection` | SLO correction configurations | `/api/slo/config/correction` |
| `syntheticalertconfig` | Synthetic test alert configurations | `/api/synthetics/settings/alerts` |
| `syntheticlocation` | Synthetic locations (read-only) | `/api/synthetics/settings/locations` |
| `synthetictest` | Synthetic test configurations | `/api/synthetics/settings/tests` |
| `team` | RBAC team management | `/api/settings/rbac/teams` |
| `user` | RBAC user management (read-only) | `/api/settings/users` |
| `websitealertconfig` | Website alert configurations | `/api/events/settings/website-alert-configs` |
| `websitemonitoring` | Website monitoring configurations | `/api/website-monitoring/config` |

---

## 5. Shared Packages

### 5.1 `shared/tagfilter/`

**Purpose**: Tag filter expressions used across many APIs

**Files**:
- `filter.go` - Tag filter data structures
- `parser.go` - Tag filter expression parser
- `filter_test.go` - Tests
- `doc.go` - Documentation

**Why shared?**: Used by 20+ API packages for filtering resources

### 5.2 `shared/types/`

**Purpose**: Common types used across multiple APIs

**Files**:
- `severity.go` - Alert severity levels (Warning, Critical)
- `granularity.go` - Time granularity (seconds, minutes)
- `operator.go` - Comparison operators (>, <, ==, etc.)
- `aggregation.go` - Aggregation types (sum, avg, min, max)
- `threshold.go` - Threshold definitions
- `types_test.go` - Tests
- `doc.go` - Documentation

**Why shared?**: These types are used consistently across all alert configurations

### 5.3 `shared/rest/`

**Purpose**: Generic REST resource patterns

**Files**:
- `resource.go` - Generic CRUD resource interface
- `readonly.go` - Read-only resource interface
- `resource_test.go` - Tests
- `doc.go` - Documentation

**Why shared?**: Provides common patterns for implementing REST clients

---

## 6. Migration Strategy

### 6.1 Migration Approach

**Strategy**: Incremental migration, one package at a time

**Process for each API package**:
1. Create new package directory under `api/`
2. Create `client.go` with API operations
3. Move models from `instana/*.go` to `models.go`
4. Extract constants to `constants.go`
5. Move tests to `*_test.go`
6. Add package documentation in `doc.go`
7. Update imports in other packages
8. Run tests to verify

### 6.2 Migration Order

**Phase 1: Shared Packages** (Week 1)
- Create `shared/tagfilter/`
- Create `shared/types/`
- Create `shared/rest/`

**Phase 2: Simple APIs** (Week 2)
- `apitoken`
- `hostagent`
- `user`
- `syntheticlocation`
- `builtineventspec`

**Phase 3: RBAC APIs** (Week 3)
- `group`
- `role`
- `team`

**Phase 4: Event APIs** (Week 4)
- `customeventspec`
- `alertingchannel`
- `alertingconfig`

**Phase 5: Application APIs** (Week 5)
- `applicationconfig`
- `applicationalertconfig`

**Phase 6: Alert APIs** (Week 6)
- `websitealertconfig`
- `mobilealertconfig`
- `infraalertconfig`
- `logalertconfig`
- `syntheticalertconfig`

**Phase 7: SLO APIs** (Week 7)
- `sliconfig`
- `sloconfig`
- `sloalertconfig`
- `slocorrection`

**Phase 8: Monitoring APIs** (Week 8)
- `websitemonitoring`
- `synthetictest`
- `customdashboard`

**Phase 9: Automation & Maintenance** (Week 9)
- `automationaction`
- `automationpolicy`
- `maintenancewindow`

**Phase 10: Main Package & Cleanup** (Week 10)
- Update `instana/` main package
- Create backward compatibility layer
- Update examples
- Update documentation

### 6.3 File Migration Mapping

| Current File | New Location | Notes |
|-------------|--------------|-------|
| `instana/application-alert-config.go` | `api/applicationalertconfig/models.go` | Split client to `client.go` |
| `instana/application-alert-config-types.go` | `api/applicationalertconfig/models.go` | Merge with models |
| `instana/application-configs-api.go` | `api/applicationconfig/client.go` + `models.go` | Split implementation |
| `instana/api-tokens-api.go` | `api/apitoken/client.go` + `models.go` | Split implementation |
| `instana/alerting-channels-api.go` | `api/alertingchannel/client.go` + `models.go` | Split implementation |
| `instana/alerts-api.go` | `api/alertingconfig/client.go` + `models.go` | Split implementation |
| `instana/automation-action-api.go` | `api/automationaction/client.go` + `models.go` | Split implementation |
| `instana/automation-policy-api.go` | `api/automationpolicy/client.go` + `models.go` | Split implementation |
| `instana/builtin-event-specification-api.go` | `api/builtineventspec/client.go` + `models.go` | Split implementation |
| `instana/custom-dashboard.go` | `api/customdashboard/client.go` + `models.go` | Split implementation |
| `instana/custom-event-specficiations-api.go` | `api/customeventspec/client.go` + `models.go` | Split implementation |
| `instana/groups-api.go` | `api/group/client.go` + `models.go` | Split implementation |
| `instana/host-agents-api.go` | `api/hostagent/client.go` + `models.go` | Split implementation |
| `instana/infra-alert-configs.go` | `api/infraalertconfig/client.go` + `models.go` | Split implementation |
| `instana/log-alert-config-api.go` | `api/logalertconfig/client.go` + `models.go` | Split implementation |
| `instana/maintenance-window-config-api.go` | `api/maintenancewindow/client.go` + `models.go` | Split implementation |
| `instana/mobile-alert-config.go` | `api/mobilealertconfig/client.go` + `models.go` | Split implementation |
| `instana/roles-api.go` | `api/role/client.go` + `models.go` | Split implementation |
| `instana/sli-config-api.go` | `api/sliconfig/client.go` + `models.go` | Split implementation |
| `instana/slo-alert-config-api.go` | `api/sloalertconfig/client.go` + `models.go` | Split implementation |
| `instana/slo-config-api.go` | `api/sloconfig/client.go` + `models.go` | Split implementation |
| `instana/slo-correction-config-api.go` | `api/slocorrection/client.go` + `models.go` | Split implementation |
| `instana/synthetic-alert-config.go` | `api/syntheticalertconfig/client.go` + `models.go` | Split implementation |
| `instana/synthetic-location.go` | `api/syntheticlocation/client.go` + `models.go` | Split implementation |
| `instana/synthetic-test.go` | `api/synthetictest/client.go` + `models.go` | Split implementation |
| `instana/teams-api.go` | `api/team/client.go` + `models.go` | Split implementation |
| `instana/users-api.go` | `api/user/client.go` + `models.go` | Split implementation |
| `instana/website-alert-config.go` | `api/websitealertconfig/client.go` + `models.go` | Split implementation |
| `instana/website-monitoring-config-api.go` | `api/websitemonitoring/client.go` + `models.go` | Split implementation |
| `instana/tag-filter.go` | `shared/tagfilter/filter.go` | Move to shared |
| `instana/severity.go` | `shared/types/severity.go` | Move to shared |
| `instana/granularity.go` | `shared/types/granularity.go` | Move to shared |
| `instana/operator.go` | `shared/types/operator.go` | Move to shared |
| `instana/aggregation.go` | `shared/types/aggregation.go` | Move to shared |
| `instana/threshold.go` | `shared/types/threshold.go` | Move to shared |

---

## 7. Implementation Roadmap

### Week 1: Foundation
- [ ] Create `shared/tagfilter/` package
- [ ] Create `shared/types/` package
- [ ] Create `shared/rest/` package
- [ ] Write tests for shared packages
- [ ] Update documentation

### Week 2: Simple APIs
- [ ] Migrate `apitoken` package
- [ ] Migrate `hostagent` package
- [ ] Migrate `user` package
- [ ] Migrate `syntheticlocation` package
- [ ] Migrate `builtineventspec` package

### Week 3: RBAC APIs
- [ ] Migrate `group` package
- [ ] Migrate `role` package
- [ ] Migrate `team` package

### Week 4: Event APIs
- [ ] Migrate `customeventspec` package
- [ ] Migrate `alertingchannel` package
- [ ] Migrate `alertingconfig` package

### Week 5: Application APIs
- [ ] Migrate `applicationconfig` package
- [ ] Migrate `applicationalertconfig` package

### Week 6: Alert APIs (Part 1)
- [ ] Migrate `websitealertconfig` package
- [ ] Migrate `mobilealertconfig` package
- [ ] Migrate `infraalertconfig` package

### Week 7: Alert APIs (Part 2) & SLO
- [ ] Migrate `logalertconfig` package
- [ ] Migrate `syntheticalertconfig` package
- [ ] Migrate `sliconfig` package
- [ ] Migrate `sloconfig` package

### Week 8: SLO & Monitoring
- [ ] Migrate `sloalertconfig` package
- [ ] Migrate `slocorrection` package
- [ ] Migrate `websitemonitoring` package
- [ ] Migrate `synthetictest` package

### Week 9: Remaining APIs
- [ ] Migrate `customdashboard` package
- [ ] Migrate `automationaction` package
- [ ] Migrate `automationpolicy` package
- [ ] Migrate `maintenancewindow` package

### Week 10: Integration & Documentation
- [ ] Update `instana/` main package
- [ ] Create backward compatibility layer
- [ ] Update all examples
- [ ] Write migration guide
- [ ] Update README and documentation
- [ ] Final testing and cleanup

---

## 8. Success Criteria

### 8.1 Technical Criteria

- [ ] All 28 API packages created and functional
- [ ] All tests passing (unit + integration)
- [ ] Test coverage > 80% per package
- [ ] No linter errors
- [ ] All packages properly documented
- [ ] Examples updated and working

### 8.2 Quality Criteria

- [ ] Each package is self-contained
- [ ] Clear separation of concerns
- [ ] Consistent file organization across packages
- [ ] Minimal shared dependencies
- [ ] Proper error handling throughout
- [ ] Context support in all API calls

### 8.3 Documentation Criteria

- [ ] Each package has `doc.go` with examples
- [ ] README updated with new structure
- [ ] Migration guide complete
- [ ] Architecture documentation updated
- [ ] CHANGELOG updated

### 8.4 User Experience Criteria

- [ ] Easy to find specific API packages
- [ ] Clear package naming
- [ ] Intuitive API design
- [ ] Helpful error messages
- [ ] Comprehensive examples

---

## Appendix A: Example Usage

### Using a Specific API Package

```go
package main

import (
    "context"
    "log"
    
    "github.com/instana/instana-go-client/client"
    "github.com/instana/instana-go-client/config"
    "github.com/instana/instana-go-client/api/applicationalertconfig"
)

func main() {
    // Create configuration
    cfg, _ := config.NewBuilder().
        WithBaseURL("https://tenant.instana.io").
        WithAPIToken("your-token").
        Build()
    
    // Create REST client
    restClient, _ := client.NewClient(cfg)
    
    // Create API-specific client
    alertClient := applicationalertconfig.NewClient(restClient)
    
    // Use the client
    ctx := context.Background()
    config, err := alertClient.Get(ctx, "alert-id")
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Alert: %s", config.Name)
}
```

### Using Main Package Facade

```go
package main

import (
    "context"
    "log"
    
    "github.com/instana/instana-go-client/instana"
)

func main() {
    // Create API client using main package
    api, _ := instana.NewAPI(instana.Config{
        BaseURL:  "https://tenant.instana.io",
        APIToken: "your-token",
    })
    
    // Access specific API clients
    alertClient := api.ApplicationAlertConfig()
    appClient := api.ApplicationConfig()
    
    // Use the clients
    ctx := context.Background()
    config, err := alertClient.Get(ctx, "alert-id")
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Alert: %s", config.Name)
}
```

---

## Appendix B: Package Template

Use this template when creating new API packages:

```
api/newpackage/
├── client.go          # API client implementation
├── models.go          # Data models
├── constants.go       # Constants and enums
├── client_test.go     # Client tests
├── models_test.go     # Model tests
└── doc.go            # Package documentation
```

**`client.go` template**:
```go
package newpackage

import (
    "context"
    "fmt"
    
    "github.com/instana/instana-go-client/client"
)

type Client struct {
    restClient *client.Client
    basePath   string
}

func NewClient(restClient *client.Client) *Client {
    return &Client{
        restClient: restClient,
        basePath:   BasePath,
    }
}

func (c *Client) Get(ctx context.Context, id string) (*Config, error) {
    // Implementation
}

func (c *Client) Create(ctx context.Context, config *Config) (*Config, error) {
    // Implementation
}

func (c *Client) Update(ctx context.Context, config *Config) (*Config, error) {
    // Implementation
}

func (c *Client) Delete(ctx context.Context, id string) error {
    // Implementation
}
```

---

## Conclusion

This refactoring plan provides a **granular, self-contained package structure** where:

✅ **Each API gets its own package** - Complete independence
✅ **Everything colocated** - Models, constants, tests in same package
✅ **Minimal shared code** - Only truly common utilities are shared
✅ **Easy to navigate** - Find everything related to an API in one place
✅ **Simple to maintain** - Changes are isolated to single packages
✅ **Clear ownership** - Each package has clear responsibility

The structure is optimized for:
- **Discoverability** - Easy to find what you need
- **Maintainability** - Changes don't ripple across packages
- **Testability** - Each package can be tested independently
- **Scalability** - Easy to add new API packages
- **Developer Experience** - Intuitive and predictable structure

---

**Document Version**: 2.0
**Last Updated**: 2026-03-11
**Status**: Draft - Awaiting Approval