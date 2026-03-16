# Instana Go Client - Package Refactoring Plan

## Executive Summary

This document outlines a comprehensive refactoring plan for the `instana-go-client` repository to transform it from its current monolithic structure into a well-organized, maintainable, and idiomatic Go client library. The refactoring focuses on organizing the `/api` directory to mirror Instana's REST API structure, with each subdirectory corresponding to a specific API endpoint domain.

**Key Goals:**
- Organize code by API domain for better discoverability
- Separate concerns: API logic, models, and shared utilities
- Follow Go standard project layout conventions
- Maintain backward compatibility where feasible (new project, no legacy users)
- Optimize for testability and long-term maintenance

---

## Table of Contents

1. [Current State Analysis](#1-current-state-analysis)
2. [Proposed Package Structure](#2-proposed-package-structure)
3. [Package Organization Principles](#3-package-organization-principles)
4. [Detailed Package Breakdown](#4-detailed-package-breakdown)
5. [Migration Strategy](#5-migration-strategy)
6. [Implementation Roadmap](#6-implementation-roadmap)
7. [Breaking Changes & Compatibility](#7-breaking-changes--compatibility)
8. [Testing Strategy](#8-testing-strategy)
9. [Documentation Requirements](#9-documentation-requirements)
10. [Success Criteria](#10-success-criteria)

---

## 1. Current State Analysis

### 1.1 Current Directory Structure

```
instana-go-client/
├── instana/              # Monolithic package with all API implementations
│   ├── *-api.go         # API resource implementations
│   ├── *.go             # Models, types, and utilities mixed together
│   └── *_test.go        # Tests
├── api/                 # Empty directories (placeholders)
├── models/              # Empty directories (placeholders)
├── client/              # HTTP client implementation
├── config/              # Configuration management
├── internal/            # Internal utilities
├── utils/               # Public utilities
├── testutils/           # Test helpers
└── examples/            # Usage examples
```

### 1.2 Problems with Current Structure

1. **Monolithic `instana/` Package**: All API implementations, models, and types are in a single package with 100+ files
2. **Poor Discoverability**: Hard to find specific API implementations
3. **Tight Coupling**: Models, API logic, and utilities are mixed
4. **Difficult Testing**: Hard to mock and test individual components
5. **Namespace Pollution**: All types exported from single package
6. **Scalability Issues**: Adding new APIs becomes increasingly difficult

### 1.3 Identified API Domains

Based on Terraform provider resources and current code analysis:

**Core Domains:**
- Alerts (Application, Website, Mobile, Infrastructure, Log, SLO, Synthetic)
- Applications (Configs, Monitoring)
- Events (Custom, Builtin)
- SLO/SLI (Service Level Objectives/Indicators)
- Synthetic (Tests, Locations, Alerts)
- Website Monitoring (Configs, Alerts)
- RBAC (Groups, Roles, Teams, Users)
- Automation (Actions, Policies)
- Infrastructure (Host Agents, Alerts)
- Dashboards (Custom Dashboards)
- Tokens (API Tokens)
- Maintenance (Maintenance Windows)

---

## 2. Proposed Package Structure

### 2.1 High-Level Structure

```
instana-go-client/
├── client/                    # Core HTTP client (existing, enhanced)
│   ├── client.go             # Main REST client implementation
│   ├── retry.go              # Retry logic
│   ├── rate_limiter.go       # Rate limiting
│   └── client_test.go
│
├── config/                    # Configuration (existing, enhanced)
│   ├── config.go             # Configuration structures
│   ├── builder.go            # Builder pattern
│   ├── loader.go             # Environment/file loaders
│   └── validator.go          # Validation logic
│
├── api/                       # API implementations by domain
│   ├── alerts/               # Alert configurations
│   │   ├── application.go    # Application alert configs
│   │   ├── website.go        # Website alert configs
│   │   ├── mobile.go         # Mobile alert configs
│   │   ├── infrastructure.go # Infrastructure alert configs
│   │   ├── log.go            # Log alert configs
│   │   ├── slo.go            # SLO alert configs
│   │   ├── synthetic.go      # Synthetic alert configs
│   │   ├── client.go         # Alert API client interface
│   │   └── *_test.go         # Tests
│   │
│   ├── applications/         # Application monitoring
│   │   ├── config.go         # Application configs
│   │   ├── client.go         # Application API client
│   │   └── *_test.go
│   │
│   ├── events/               # Event specifications
│   │   ├── custom.go         # Custom event specs
│   │   ├── builtin.go        # Builtin event specs
│   │   ├── client.go         # Events API client
│   │   └── *_test.go
│   │
│   ├── slo/                  # SLO/SLI management
│   │   ├── slo.go            # SLO configs
│   │   ├── sli.go            # SLI configs
│   │   ├── correction.go     # SLO corrections
│   │   ├── client.go         # SLO API client
│   │   └── *_test.go
│   │
│   ├── synthetic/            # Synthetic monitoring
│   │   ├── test.go           # Synthetic tests
│   │   ├── location.go       # Synthetic locations
│   │   ├── client.go         # Synthetic API client
│   │   └── *_test.go
│   │
│   ├── website/              # Website monitoring
│   │   ├── monitoring.go     # Website monitoring configs
│   │   ├── client.go         # Website API client
│   │   └── *_test.go
│   │
│   ├── rbac/                 # Role-based access control
│   │   ├── group.go          # Groups
│   │   ├── role.go           # Roles
│   │   ├── team.go           # Teams
│   │   ├── user.go           # Users
│   │   ├── client.go         # RBAC API client
│   │   └── *_test.go
│   │
│   ├── automation/           # Automation
│   │   ├── action.go         # Automation actions
│   │   ├── policy.go         # Automation policies
│   │   ├── client.go         # Automation API client
│   │   └── *_test.go
│   │
│   ├── infrastructure/       # Infrastructure monitoring
│   │   ├── host_agent.go     # Host agents
│   │   ├── client.go         # Infrastructure API client
│   │   └── *_test.go
│   │
│   ├── dashboards/           # Dashboards
│   │   ├── custom.go         # Custom dashboards
│   │   ├── client.go         # Dashboard API client
│   │   └── *_test.go
│   │
│   ├── tokens/               # API tokens
│   │   ├── token.go          # API token management
│   │   ├── client.go         # Token API client
│   │   └── *_test.go
│   │
│   ├── maintenance/          # Maintenance windows
│   │   ├── window.go         # Maintenance window configs
│   │   ├── client.go         # Maintenance API client
│   │   └── *_test.go
│   │
│   └── channels/             # Alerting channels
│       ├── channel.go        # Alerting channel configs
│       ├── client.go         # Channel API client
│       └── *_test.go
│
├── models/                    # Shared data models
│   ├── common/               # Common models used across APIs
│   │   ├── threshold.go      # Threshold definitions
│   │   ├── time_window.go    # Time window configurations
│   │   ├── severity.go       # Severity levels
│   │   ├── granularity.go    # Granularity types
│   │   ├── boundary_scope.go # Boundary scope
│   │   ├── custom_payload.go # Custom payload fields
│   │   └── *_test.go
│   │
│   ├── filters/              # Filter models
│   │   ├── tag_filter.go     # Tag filter expressions
│   │   ├── entity_filter.go  # Entity filters
│   │   └── *_test.go
│   │
│   ├── rules/                # Rule models
│   │   ├── alert_rule.go     # Alert rule definitions
│   │   ├── threshold_rule.go # Threshold rules
│   │   └── *_test.go
│   │
│   └── enums/                # Enumeration types
│       ├── alert_type.go     # Alert types
│       ├── channel_type.go   # Channel types
│       ├── operator.go       # Comparison operators
│       ├── aggregation.go    # Aggregation types
│       └── *_test.go
│
├── internal/                  # Internal packages (not exported)
│   ├── rest/                 # REST resource implementations
│   │   ├── resource.go       # Generic REST resource
│   │   ├── readonly.go       # Read-only resource
│   │   └── *_test.go
│   │
│   ├── json/                 # JSON unmarshalling utilities
│   │   ├── unmarshaller.go   # Custom unmarshallers
│   │   └── *_test.go
│   │
│   └── validation/           # Internal validation
│       ├── validator.go      # Validation utilities
│       └── *_test.go
│
├── errors/                    # Error types and handling
│   ├── errors.go             # Error definitions
│   ├── types.go              # Error type constants
│   └── errors_test.go
│
├── instana/                   # Main package (facade/compatibility)
│   ├── api.go                # Main API interface
│   ├── client.go             # Client factory methods
│   ├── backward_compat.go    # Backward compatibility aliases
│   └── *_test.go
│
├── utils/                     # Public utilities
│   ├── string.go             # String utilities
│   ├── int.go                # Integer utilities
│   ├── slice.go              # Slice utilities
│   └── *_test.go
│
├── testutils/                 # Test utilities
│   ├── mock_server.go        # Mock HTTP server
│   ├── fixtures.go           # Test fixtures
│   └── *_test.go
│
├── examples/                  # Usage examples
│   ├── basic_usage/
│   ├── alerts/
│   ├── applications/
│   └── ...
│
├── mocks/                     # Generated mocks
│   └── *.go
│
├── docs/                      # Additional documentation
│   ├── architecture.md
│   ├── api_reference.md
│   └── migration_guide.md
│
├── go.mod
├── go.sum
├── README.md
├── CHANGELOG.md
├── LICENSE
└── Makefile
```

---

## 3. Package Organization Principles

### 3.1 Design Principles

1. **Domain-Driven Organization**: Each API domain gets its own package
2. **Single Responsibility**: Each package handles one API domain
3. **Clear Boundaries**: Well-defined interfaces between packages
4. **Minimal Dependencies**: Packages depend on interfaces, not implementations
5. **Testability**: Easy to mock and test in isolation
6. **Discoverability**: Intuitive package names matching API structure

### 3.2 Naming Conventions

- **Package Names**: Lowercase, singular, descriptive (e.g., `alert`, `application`)
- **File Names**: Lowercase with underscores (e.g., `application_alert.go`)
- **Type Names**: PascalCase, descriptive (e.g., `ApplicationAlertConfig`)
- **Interface Names**: End with `-er` or describe capability (e.g., `AlertClient`, `ConfigManager`)

### 3.3 Package Responsibilities

#### API Packages (`api/*`)
- Implement API-specific HTTP operations (CRUD)
- Handle API-specific request/response transformations
- Contain API-specific constants (endpoints, defaults)
- Define API client interfaces
- Include API-specific tests

#### Model Packages (`models/*`)
- Define data structures shared across multiple APIs
- Implement JSON marshaling/unmarshaling
- Contain validation logic for models
- Include model-specific tests

#### Internal Packages (`internal/*`)
- Implement generic REST resource patterns
- Provide internal utilities not meant for public use
- Handle low-level HTTP operations

#### Main Package (`instana/`)
- Provide main API facade
- Factory methods for creating clients
- Backward compatibility layer
- High-level documentation

---

## 4. Detailed Package Breakdown

### 4.1 Alert APIs (`api/alerts/`)

**Purpose**: Manage all alert configuration types across different monitoring domains

**Files**:
- `application.go` - Application alert configurations
- `website.go` - Website alert configurations  
- `mobile.go` - Mobile app alert configurations
- `infrastructure.go` - Infrastructure alert configurations
- `log.go` - Log alert configurations
- `slo.go` - SLO alert configurations
- `synthetic.go` - Synthetic test alert configurations
- `client.go` - Alert API client interface and implementation
- `constants.go` - Alert-specific constants
- `*_test.go` - Tests for each alert type

**Key Types**:
```go
// Client interface
type Client interface {
    // Application alerts
    GetApplicationAlert(ctx context.Context, id string) (*ApplicationAlertConfig, error)
    CreateApplicationAlert(ctx context.Context, config *ApplicationAlertConfig) (*ApplicationAlertConfig, error)
    UpdateApplicationAlert(ctx context.Context, config *ApplicationAlertConfig) (*ApplicationAlertConfig, error)
    DeleteApplicationAlert(ctx context.Context, id string) error
    
    // Website alerts
    GetWebsiteAlert(ctx context.Context, id string) (*WebsiteAlertConfig, error)
    // ... similar methods for other alert types
}

// Implementation
type client struct {
    restClient *client.Client
    basePath   string
}
```

**API Endpoints**:
- `/api/events/settings/application-alert-configs`
- `/api/events/settings/website-alert-configs`
- `/api/events/settings/mobile-app-alert-configs`
- `/api/events/settings/global-alert-configs/infrastructure`
- `/api/events/settings/global-alert-configs/logs`
- `/api/slo/config/alert`
- `/api/synthetics/settings/alerts`

### 4.2 Application APIs (`api/applications/`)

**Purpose**: Manage application monitoring configurations

**Files**:
- `config.go` - Application configuration CRUD operations
- `client.go` - Application API client interface
- `constants.go` - Application-specific constants
- `*_test.go` - Tests

**Key Types**:
```go
type Client interface {
    GetConfig(ctx context.Context, id string) (*Config, error)
    CreateConfig(ctx context.Context, config *Config) (*Config, error)
    UpdateConfig(ctx context.Context, config *Config) (*Config, error)
    DeleteConfig(ctx context.Context, id string) error
    ListConfigs(ctx context.Context) ([]*Config, error)
}

type Config struct {
    ID                  string
    Label               string
    TagFilterExpression *filters.TagFilter
    Scope               Scope
    BoundaryScope       common.BoundaryScope
    AccessRules         []AccessRule
}
```

**API Endpoints**:
- `/api/application-monitoring/settings/application`

### 4.3 Events APIs (`api/events/`)

**Purpose**: Manage custom and builtin event specifications

**Files**:
- `custom.go` - Custom event specifications
- `builtin.go` - Builtin event specifications (read-only)
- `client.go` - Events API client interface
- `constants.go` - Event-specific constants
- `*_test.go` - Tests

**Key Types**:
```go
type Client interface {
    // Custom events
    GetCustomEvent(ctx context.Context, id string) (*CustomEventSpec, error)
    CreateCustomEvent(ctx context.Context, spec *CustomEventSpec) (*CustomEventSpec, error)
    UpdateCustomEvent(ctx context.Context, spec *CustomEventSpec) (*CustomEventSpec, error)
    DeleteCustomEvent(ctx context.Context, id string) error
    
    // Builtin events (read-only)
    GetBuiltinEvent(ctx context.Context, id string) (*BuiltinEventSpec, error)
    ListBuiltinEvents(ctx context.Context) ([]*BuiltinEventSpec, error)
}
```

**API Endpoints**:
- `/api/events/settings/event-specifications/custom`
- `/api/events/settings/event-specifications/built-in`

### 4.4 SLO APIs (`api/slo/`)

**Purpose**: Manage Service Level Objectives and Indicators

**Files**:
- `slo.go` - SLO configurations
- `sli.go` - SLI configurations
- `correction.go` - SLO correction configurations
- `client.go` - SLO API client interface
- `constants.go` - SLO-specific constants
- `*_test.go` - Tests

**Key Types**:
```go
type Client interface {
    // SLO operations
    GetSLO(ctx context.Context, id string) (*SLOConfig, error)
    CreateSLO(ctx context.Context, config *SLOConfig) (*SLOConfig, error)
    UpdateSLO(ctx context.Context, config *SLOConfig) (*SLOConfig, error)
    DeleteSLO(ctx context.Context, id string) error
    
    // SLI operations
    GetSLI(ctx context.Context, id string) (*SLIConfig, error)
    CreateSLI(ctx context.Context, config *SLIConfig) (*SLIConfig, error)
    UpdateSLI(ctx context.Context, config *SLIConfig) (*SLIConfig, error)
    DeleteSLI(ctx context.Context, id string) error
    
    // Correction operations
    GetCorrection(ctx context.Context, id string) (*CorrectionConfig, error)
    CreateCorrection(ctx context.Context, config *CorrectionConfig) (*CorrectionConfig, error)
    UpdateCorrection(ctx context.Context, config *CorrectionConfig) (*CorrectionConfig, error)
    DeleteCorrection(ctx context.Context, id string) error
}
```

**API Endpoints**:
- `/api/slo/config`
- `/api/application-monitoring/settings/sli/config`
- `/api/slo/config/correction`

### 4.5 Synthetic APIs (`api/synthetic/`)

**Purpose**: Manage synthetic monitoring tests and locations

**Files**:
- `test.go` - Synthetic test configurations
- `location.go` - Synthetic locations (read-only)
- `client.go` - Synthetic API client interface
- `constants.go` - Synthetic-specific constants
- `*_test.go` - Tests

**API Endpoints**:
- `/api/synthetics/settings/tests`
- `/api/synthetics/settings/locations`

### 4.6 Website APIs (`api/website/`)

**Purpose**: Manage website monitoring configurations

**Files**:
- `monitoring.go` - Website monitoring configurations
- `client.go` - Website API client interface
- `constants.go` - Website-specific constants
- `*_test.go` - Tests

**API Endpoints**:
- `/api/website-monitoring/config`

### 4.7 RBAC APIs (`api/rbac/`)

**Purpose**: Manage role-based access control (groups, roles, teams, users)

**Files**:
- `group.go` - Group management
- `role.go` - Role management
- `team.go` - Team management
- `user.go` - User management (read-only)
- `client.go` - RBAC API client interface
- `constants.go` - RBAC-specific constants
- `*_test.go` - Tests

**API Endpoints**:
- `/api/settings/rbac/groups`
- `/api/settings/rbac/roles`
- `/api/settings/rbac/teams`
- `/api/settings/users`

### 4.8 Automation APIs (`api/automation/`)

**Purpose**: Manage automation actions and policies

**Files**:
- `action.go` - Automation actions
- `policy.go` - Automation policies
- `client.go` - Automation API client interface
- `constants.go` - Automation-specific constants
- `*_test.go` - Tests

**API Endpoints**:
- `/api/automation/actions`
- `/api/automation/policies`

### 4.9 Infrastructure APIs (`api/infrastructure/`)

**Purpose**: Manage infrastructure monitoring

**Files**:
- `host_agent.go` - Host agent information (read-only)
- `client.go` - Infrastructure API client interface
- `constants.go` - Infrastructure-specific constants
- `*_test.go` - Tests

**API Endpoints**:
- `/api/host-agent`

### 4.10 Dashboard APIs (`api/dashboards/`)

**Purpose**: Manage custom dashboards

**Files**:
- `custom.go` - Custom dashboard configurations
- `client.go` - Dashboard API client interface
- `constants.go` - Dashboard-specific constants
- `*_test.go` - Tests

**API Endpoints**:
- `/api/custom-dashboard`

### 4.11 Token APIs (`api/tokens/`)

**Purpose**: Manage API tokens

**Files**:
- `token.go` - API token management
- `client.go` - Token API client interface
- `constants.go` - Token-specific constants
- `*_test.go` - Tests

**API Endpoints**:
- `/api/settings/api-tokens`

### 4.12 Maintenance APIs (`api/maintenance/`)

**Purpose**: Manage maintenance windows

**Files**:
- `window.go` - Maintenance window configurations
- `client.go` - Maintenance API client interface
- `constants.go` - Maintenance-specific constants
- `*_test.go` - Tests

**API Endpoints**:
- `/api/settings/maintenance-windows`

### 4.13 Channel APIs (`api/channels/`)

**Purpose**: Manage alerting channels

**Files**:
- `channel.go` - Alerting channel configurations
- `client.go` - Channel API client interface
- `constants.go` - Channel-specific constants
- `*_test.go` - Tests

**API Endpoints**:
- `/api/events/settings/alertingChannels`

---

## 5. Migration Strategy

### 5.1 Migration Approach

**Strategy**: Incremental migration with backward compatibility layer

**Phases**:
1. **Phase 1**: Create new package structure (empty)
2. **Phase 2**: Move shared models to `models/` packages
3. **Phase 3**: Migrate API implementations to `api/` packages
4. **Phase 4**: Create backward compatibility layer in `instana/`
5. **Phase 5**: Update tests and examples
6. **Phase 6**: Update documentation
7. **Phase 7**: Deprecate old structure (if needed)

### 5.2 Backward Compatibility

Since this is a new project with no external users yet, we can be more aggressive with breaking changes. However, we should still provide:

1. **Type Aliases**: Create aliases in `instana/` package for commonly used types
2. **Factory Functions**: Provide convenience functions in main package
3. **Migration Guide**: Document how to update code for new structure

**Example Compatibility Layer**:
```go
// instana/backward_compat.go
package instana

import (
    "github.com/instana/instana-go-client/api/alerts"
    "github.com/instana/instana-go-client/api/applications"
)

// Type aliases for backward compatibility
type ApplicationAlertConfig = alerts.ApplicationAlertConfig
type ApplicationConfig = applications.Config

// Factory functions
func NewAlertClient(restClient *client.Client) alerts.Client {
    return alerts.NewClient(restClient)
}
```

### 5.3 File Migration Mapping

| Current Location | New Location | Notes |
|-----------------|--------------|-------|
| `instana/alerting-channels-api.go` | `api/channels/channel.go` | Split models to `models/` |
| `instana/application-alert-config.go` | `api/alerts/application.go` | Move rules to `models/rules/` |
| `instana/application-configs-api.go` | `api/applications/config.go` | Move scope to `models/common/` |
| `instana/custom-event-specficiations-api.go` | `api/events/custom.go` | Keep event-specific logic |
| `instana/builtin-event-specification-api.go` | `api/events/builtin.go` | Read-only operations |
| `instana/slo-config-api.go` | `api/slo/slo.go` | SLO operations |
| `instana/sli-config-api.go` | `api/slo/sli.go` | SLI operations |
| `instana/synthetic-test.go` | `api/synthetic/test.go` | Synthetic tests |
| `instana/website-monitoring-config-api.go` | `api/website/monitoring.go` | Website monitoring |
| `instana/groups-api.go` | `api/rbac/group.go` | RBAC groups |
| `instana/roles-api.go` | `api/rbac/role.go` | RBAC roles |
| `instana/teams-api.go` | `api/rbac/team.go` | RBAC teams |
| `instana/users-api.go` | `api/rbac/user.go` | RBAC users |
| `instana/automation-action-api.go` | `api/automation/action.go` | Automation actions |
| `instana/automation-policy-api.go` | `api/automation/policy.go` | Automation policies |
| `instana/host-agents-api.go` | `api/infrastructure/host_agent.go` | Infrastructure |
| `instana/custom-dashboard.go` | `api/dashboards/custom.go` | Dashboards |
| `instana/api-tokens-api.go` | `api/tokens/token.go` | API tokens |
| `instana/maintenance-window-config-api.go` | `api/maintenance/window.go` | Maintenance windows |
| `instana/threshold.go` | `models/common/threshold.go` | Shared model |
| `instana/tag-filter.go` | `models/filters/tag_filter.go` | Shared filter |
| `instana/severity.go` | `models/common/severity.go` | Shared enum |
| `instana/granularity.go` | `models/common/granularity.go` | Shared enum |
| `instana/operator.go` | `models/enums/operator.go` | Shared enum |
| `instana/aggregation.go` | `models/enums/aggregation.go` | Shared enum |

---

## 6. Implementation Roadmap

### Phase 1: Foundation (Week 1)
**Goal**: Set up new package structure and shared models

**Tasks**:
- [ ] Create new directory structure
- [ ] Move shared models to `models/common/`
- [ ] Move filter models to `models/filters/`
- [ ] Move enum types to `models/enums/`
- [ ] Move rule models to `models/rules/`
- [ ] Update imports in existing code
- [ ] Run tests to ensure nothing breaks

**Deliverables**:
- New `models/` package structure
- All shared models migrated
- Tests passing

### Phase 2: Alert APIs (Week 2)
**Goal**: Migrate all alert-related APIs

**Tasks**:
- [ ] Create `api/alerts/` package structure
- [ ] Migrate application alert API
- [ ] Migrate website alert API
- [ ] Migrate mobile alert API
- [ ] Migrate infrastructure alert API
- [ ] Migrate log alert API
- [ ] Migrate SLO alert API
- [ ] Migrate synthetic alert API
- [ ] Create unified alert client interface
- [ ] Write tests for each alert type

**Deliverables**:
- Complete `api/alerts/` package
- All alert APIs migrated
- Tests passing

### Phase 3: Core APIs (Week 3)
**Goal**: Migrate application, event, and SLO APIs

**Tasks**:
- [ ] Create `api/applications/` package
- [ ] Migrate application config API
- [ ] Create `api/events/` package
- [ ] Migrate custom event API
- [ ] Migrate builtin event API
- [ ] Create `api/slo/` package
- [ ] Migrate SLO config API
- [ ] Migrate SLI config API
- [ ] Migrate SLO correction API
- [ ] Write tests for each API

**Deliverables**:
- `api/applications/` package complete
- `api/events/` package complete
- `api/slo/` package complete
- Tests passing

### Phase 4: Monitoring APIs (Week 4)
**Goal**: Migrate synthetic and website monitoring APIs

**Tasks**:
- [ ] Create `api/synthetic/` package
- [ ] Migrate synthetic test API
- [ ] Migrate synthetic location API
- [ ] Create `api/website/` package
- [ ] Migrate website monitoring API
- [ ] Write tests for each API

**Deliverables**:
- `api/synthetic/` package complete
- `api/website/` package complete
- Tests passing

### Phase 5: Management APIs (Week 5)
**Goal**: Migrate RBAC, automation, and infrastructure APIs

**Tasks**:
- [ ] Create `api/rbac/` package
- [ ] Migrate group API
- [ ] Migrate role API
- [ ] Migrate team API
- [ ] Migrate user API
- [ ] Create `api/automation/` package
- [ ] Migrate automation action API
- [ ] Migrate automation policy API
- [ ] Create `api/infrastructure/` package
- [ ] Migrate host agent API
- [ ] Write tests for each API

**Deliverables**:
- `api/rbac/` package complete
- `api/automation/` package complete
- `api/infrastructure/` package complete
- Tests passing

### Phase 6: Remaining APIs (Week 6)
**Goal**: Migrate dashboard, token, maintenance, and channel APIs

**Tasks**:
- [ ] Create `api/dashboards/` package
- [ ] Migrate custom dashboard API
- [ ] Create `api/tokens/` package
- [ ] Migrate API token API
- [ ] Create `api/maintenance/` package
- [ ] Migrate maintenance window API
- [ ] Create `api/channels/` package
- [ ] Migrate alerting channel API
- [ ] Write tests for each API

**Deliverables**:
- All remaining API packages complete
- Tests passing

### Phase 7: Main Package & Compatibility (Week 7)
**Goal**: Create main package facade and backward compatibility

**Tasks**:
- [ ] Design main `instana/` package API
- [ ] Create factory functions
- [ ] Create type aliases for backward compatibility
- [ ] Update `InstanaAPI` interface
- [ ] Implement main API client
- [ ] Write integration tests
- [ ] Update examples

**Deliverables**:
- Main `instana/` package complete
- Backward compatibility layer
- Integration tests passing
- Examples updated

### Phase 8: Documentation & Polish (Week 8)
**Goal**: Complete documentation and final polish

**Tasks**:
- [ ] Write package documentation
- [ ] Create migration guide
- [ ] Update README
- [ ] Create architecture documentation
- [ ] Add code examples for each API
- [ ] Review and refactor code
- [ ] Run linters and fix issues
- [ ] Update CHANGELOG

**Deliverables**:
- Complete documentation
- Migration guide
- Clean, linted code
- Ready for release

---

## 7. Breaking Changes & Compatibility

### 7.1 Breaking Changes

Since this is a new project with no external users, we can make breaking changes freely. However, we should document them:

**Import Path Changes**:
```go
// Old
import "github.com/instana/instana-go-client/instana"

// New - specific imports
import (
    "github.com/instana/instana-go-client/api/alerts"
    "github.com/instana/instana-go-client/api/applications"
    "github.com/instana/instana-go-client/models/common"
)

// Or use main package for convenience
import "github.com/instana/instana-go-client/instana"
```

**Type Location Changes**:
```go
// Old
config := &instana.ApplicationAlertConfig{}

// New - direct import
config := &alerts.ApplicationAlertConfig{}

// Or via main package alias
config := &instana.ApplicationAlertConfig{}
```

**Client Creation Changes**:
```go
// Old
api := instana.NewInstanaAPI(token, endpoint, skipTLS)
alertConfigs := api.ApplicationAlertConfigs()

// New - domain-specific clients
alertClient := alerts.NewClient(restClient)
config, err := alertClient.GetApplicationAlert(ctx, id)

// Or via main package
api := instana.NewAPI(config)
config, err := api.Alerts().GetApplicationAlert(ctx, id)
```

### 7.2 Compatibility Strategy

**Option 1: Full Backward Compatibility** (Recommended for initial release)
- Keep all existing types in `instana/` package
- Create aliases pointing to new locations
- Provide factory functions in main package
- Mark old locations as deprecated

**Option 2: Clean Break** (For v2.0.0)
- Remove old structure completely
- Provide comprehensive migration guide
- Offer migration tool/script
- Clear communication about breaking changes

**Recommendation**: Start with Option 1, transition to Option 2 for v2.0.0

---

## 8. Testing Strategy

### 8.1 Test Organization

**Unit Tests**:
- Colocated with implementation (`*_test.go`)
- Test each API operation independently
- Mock HTTP client for isolation
- Test error handling and edge cases

**Integration Tests**:
- Separate `test/integration/` directory
- Test against mock Instana API server
- Test complete workflows
- Test backward compatibility

**Example Test Structure**:
```go
// api/alerts/application_test.go
package alerts_test

import (
    "context"
    "testing"
    
    "github.com/instana/instana-go-client/api/alerts"
    "github.com/instana/instana-go-client/testutils"
)

func TestGetApplicationAlert(t *testing.T) {
    // Setup mock server
    server := testutils.NewMockServer()
    defer server.Close()
    
    // Create client
    client := alerts.NewClient(server.Client())
    
    // Test
    config, err := client.GetApplicationAlert(context.Background(), "test-id")
    
    // Assertions
    require.NoError(t, err)
    require.NotNil(t, config)
    require.Equal(t, "test-id", config.ID)
}
```

### 8.2 Test Coverage Goals

- **Unit Test Coverage**: > 80%
- **Integration Test Coverage**: > 60%
- **Critical Path Coverage**: 100%

### 8.3 Test Utilities

**Mock Server** (`testutils/mock_server.go`):
```go
type MockServer struct {
    server *httptest.Server
    routes map[string]http.HandlerFunc
}

func NewMockServer() *MockServer
func (m *MockServer) AddRoute(method, path string, handler http.HandlerFunc)
func (m *MockServer) Client() *client.Client
func (m *MockServer) Close()
```

**Test Fixtures** (`testutils/fixtures.go`):
```go
func NewApplicationAlertConfig() *alerts.ApplicationAlertConfig
func NewWebsiteAlertConfig() *alerts.WebsiteAlertConfig
// ... fixtures for each type
```

---

## 9. Documentation Requirements

### 9.1 Package Documentation

Each package must have:
- Package-level documentation (`doc.go`)
- Usage examples
- API reference
- Common patterns

**Example**:
```go
// Package alerts provides clients for managing Instana alert configurations.
//
// This package supports multiple alert types:
//   - Application alerts
//   - Website alerts
//   - Mobile app alerts
//   - Infrastructure alerts
//   - Log alerts
//   - SLO alerts
//   - Synthetic alerts
//
// Basic usage:
//
//   client := alerts.NewClient(restClient)
//   config, err := client.GetApplicationAlert(ctx, "alert-id")
//   if err != nil {
//       log.Fatal(err)
//   }
//
// For more examples, see the examples directory.
package alerts
```

### 9.2 Migration Guide

Create `docs/migration_guide.md` with:
- Overview of changes
- Import path updates
- Type location changes
- Code examples (before/after)
- Common migration patterns
- Troubleshooting

### 9.3 Architecture Documentation

Create `docs/architecture.md` with:
- Package structure overview
- Design decisions and rationale
- Dependency graph
- Extension points
- Best practices

### 9.4 API Reference

Generate API documentation using `godoc`:
```bash
godoc -http=:6060
```

Ensure all exported types, functions, and methods have documentation.

---

## 10. Success Criteria

### 10.1 Technical Criteria

- [ ] All API implementations migrated to new structure
- [ ] All tests passing (unit + integration)
- [ ] Test coverage > 80%
- [ ] No linter errors
- [ ] All packages properly documented
- [ ] Examples updated and working
- [ ] Backward compatibility maintained (if applicable)

### 10.2 Quality Criteria

- [ ] Code follows Go best practices
- [ ] Clear separation of concerns
- [ ] Minimal coupling between packages
- [ ] Consistent naming conventions
- [ ] Comprehensive error handling
- [ ] Proper context usage throughout

### 10.3 Documentation Criteria

- [ ] README updated
- [ ] Migration guide complete
- [ ] Architecture documentation complete
- [ ] API reference generated
- [ ] Examples for each API domain
- [ ] CHANGELOG updated

### 10.4 User Experience Criteria

- [ ] Intuitive package structure
- [ ] Easy to discover APIs
- [ ] Clear error messages
- [ ] Helpful examples
- [ ] Smooth migration path

---

## Appendix A: Package Dependency Graph

```
┌─────────────────────────────────────────────────────────────┐
│                         instana/                             │
│                    (Main Package Facade)                     │
└────────────────────────┬────────────────────────────────────┘
                         │
         ┌───────────────┼───────────────┐
         │               │               │
         ▼               ▼               ▼
    ┌────────┐      ┌────────┐     ┌─────────┐
    │ client │      │ config │     │ errors  │
    └────┬───┘      └────────┘     └─────────┘
         │
         │
    ┌────┴──────────────────────────────────────────┐
    │                                                │
    ▼                                                ▼
┌────────────────────────────────┐      ┌──────────────────┐
│         api/* packages         │      │  models/*        │
│  (alerts, applications, etc.)  │◄─────┤  (common, etc.)  │
└────────────────────────────────┘      └──────────────────┘
    │
    │
    ▼
┌────────────────────────────────┐
│      internal/* packages       │
│   (rest, json, validation)     │
└────────────────────────────────┘
```

**Dependency Rules**:
1. `api/*` packages depend on `models/*`, `client`, `errors`
2. `models/*` packages have no dependencies (except standard library)
3. `internal/*` packages depend on `client`, `errors`
4. `instana/` depends on all `api/*` packages
5. No circular dependencies allowed

---

## Appendix B: Example Code Patterns

### Pattern 1: Creating an API Client

```go
package main

import (
    "context"
    "log"
    
    "github.com/instana/instana-go-client/instana"
)

func main() {
    // Option 1: Use main package facade
    config, _ := instana.NewConfigBuilder().
        WithBaseURL("https://tenant.instana.io").
        WithAPIToken("your-token").
        Build()
    
    api, err := instana.NewAPI(config)
    if err != nil {
        log.Fatal(err)
    }
    
    // Access domain-specific clients
    alertClient := api.Alerts()
    appClient := api.Applications()
    
    // Option 2: Create domain client directly
    restClient, _ := client.NewClient(config)
    alertClient := alerts.NewClient(restClient)
}
```

### Pattern 2: CRUD Operations

```go
// Create
config := &alerts.ApplicationAlertConfig{
    Name:        "High Error Rate",
    Description: "Alert on high error rate",
    // ... other fields
}
created, err := alertClient.CreateApplicationAlert(ctx, config)

// Read
config, err := alertClient.GetApplicationAlert(ctx, "alert-id")

// Update
config.Name = "Updated Name"
updated, err := alertClient.UpdateApplicationAlert(ctx, config)

// Delete
err := alertClient.DeleteApplicationAlert(ctx, "alert-id")
```

### Pattern 3: Error Handling

```go
config, err := alertClient.GetApplicationAlert(ctx, "alert-id")
if err != nil {
    // Check error type
    if errors.IsRetryableError(err) {
        // Retry logic
    }
    
    // Extract status code
    if statusCode := errors.ExtractStatusCode(err); statusCode == 404 {
        // Handle not found
    }
    
    // Type assertion for detailed info
    if instanaErr, ok := err.(*errors.InstanaError); ok {
        log.Printf("Error type: %s, Status: %d", 
            instanaErr.Type, instanaErr.StatusCode)
    }
}
```

---

## Appendix C: Frequently Asked Questions

**Q: Why separate `api/` and `models/` packages?**
A: This separation follows the principle of separation of concerns. Models are pure data structures that can be shared across multiple APIs, while API packages contain the logic for interacting with specific endpoints.

**Q: Why use `internal/` for REST resource implementations?**
A: The `internal/` package prevents external packages from depending on implementation details. It allows us to change internal implementations without breaking the public API.

**Q: How do I add a new API endpoint?**
A: 
1. Determine which domain it belongs to (or create new domain)
2. Add methods to the appropriate client interface
3. Implement the methods in the client
4. Add tests
5. Update documentation

**Q: Should I use the main `instana` package or import specific API packages?**
A: For most use cases, use the main `instana` package for convenience. Import specific API packages when you need fine-grained control or want to minimize dependencies.

**Q: How do I handle backward compatibility?**
A: Use type aliases and factory functions in the main `instana` package to maintain compatibility with existing code while encouraging migration to the new structure.

---

## Conclusion

This refactoring plan provides a comprehensive roadmap for transforming the instana-go-client into a well-organized, maintainable, and idiomatic Go library. The proposed structure:

✅ **Improves Discoverability**: Clear package names matching API domains
✅ **Enhances Maintainability**: Separation of concerns and single responsibility
✅ **Supports Scalability**: Easy to add new APIs and features
✅ **Follows Go Best Practices**: Standard project layout and conventions
✅ **Optimizes for Testing**: Clear boundaries and mockable interfaces
✅ **Provides Flexibility**: Multiple ways to use the library

The incremental migration strategy ensures we can make progress while maintaining stability, and the comprehensive testing and documentation requirements ensure quality throughout the process.

---

**Document Version**: 1.0
**Last Updated**: 2026-03-11
**Status**: Draft - Awaiting Approval