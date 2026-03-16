# Instana Go Client - API Interface Design

## Main API Interface

The main entry point for applications using this client will be the `InstanaAPI` interface in the `instana` package.

## Package Structure

```
instana-go-client/
├── instana/                    # Main package - Application entry point
│   ├── api.go                 # Main InstanaAPI interface and implementation
│   ├── client.go              # Factory methods for creating API
│   └── doc.go                 # Package documentation
│
├── api/                        # Individual API packages (28 packages)
│   ├── applicationalertconfig/
│   ├── applicationconfig/
│   ├── apitoken/
│   └── ... (25 more packages)
│
├── client/                     # HTTP client
├── config/                     # Configuration
├── shared/                     # Shared utilities
└── ...
```

## Main API Interface (`instana/api.go`)

```go
package instana

import (
    "github.com/instana/instana-go-client/api/applicationalertconfig"
    "github.com/instana/instana-go-client/api/applicationconfig"
    "github.com/instana/instana-go-client/api/apitoken"
    "github.com/instana/instana-go-client/api/alertingchannel"
    "github.com/instana/instana-go-client/api/alertingconfig"
    "github.com/instana/instana-go-client/api/automationaction"
    "github.com/instana/instana-go-client/api/automationpolicy"
    "github.com/instana/instana-go-client/api/builtineventspec"
    "github.com/instana/instana-go-client/api/customdashboard"
    "github.com/instana/instana-go-client/api/customeventspec"
    "github.com/instana/instana-go-client/api/group"
    "github.com/instana/instana-go-client/api/hostagent"
    "github.com/instana/instana-go-client/api/infraalertconfig"
    "github.com/instana/instana-go-client/api/logalertconfig"
    "github.com/instana/instana-go-client/api/maintenancewindow"
    "github.com/instana/instana-go-client/api/mobilealertconfig"
    "github.com/instana/instana-go-client/api/role"
    "github.com/instana/instana-go-client/api/sliconfig"
    "github.com/instana/instana-go-client/api/sloalertconfig"
    "github.com/instana/instana-go-client/api/sloconfig"
    "github.com/instana/instana-go-client/api/slocorrection"
    "github.com/instana/instana-go-client/api/syntheticalertconfig"
    "github.com/instana/instana-go-client/api/syntheticlocation"
    "github.com/instana/instana-go-client/api/synthetictest"
    "github.com/instana/instana-go-client/api/team"
    "github.com/instana/instana-go-client/api/user"
    "github.com/instana/instana-go-client/api/websitealertconfig"
    "github.com/instana/instana-go-client/api/websitemonitoring"
    "github.com/instana/instana-go-client/client"
    "github.com/instana/instana-go-client/config"
)

// InstanaAPI is the main interface for interacting with Instana API
// It provides access to all API endpoint clients
type InstanaAPI interface {
    // Alert Configurations
    ApplicationAlertConfig() *applicationalertconfig.Client
    WebsiteAlertConfig() *websitealertconfig.Client
    MobileAlertConfig() *mobilealertconfig.Client
    InfraAlertConfig() *infraalertconfig.Client
    LogAlertConfig() *logalertconfig.Client
    SLOAlertConfig() *sloalertconfig.Client
    SyntheticAlertConfig() *syntheticalertconfig.Client
    
    // Application Monitoring
    ApplicationConfig() *applicationconfig.Client
    
    // Events
    CustomEventSpec() *customeventspec.Client
    BuiltinEventSpec() *builtineventspec.Client
    
    // Alerting
    AlertingChannel() *alertingchannel.Client
    AlertingConfig() *alertingconfig.Client
    
    // SLO/SLI
    SLIConfig() *sliconfig.Client
    SLOConfig() *sloconfig.Client
    SLOCorrection() *slocorrection.Client
    
    // Synthetic Monitoring
    SyntheticTest() *synthetictest.Client
    SyntheticLocation() *syntheticlocation.Client
    
    // Website Monitoring
    WebsiteMonitoring() *websitemonitoring.Client
    
    // RBAC
    Group() *group.Client
    Role() *role.Client
    Team() *team.Client
    User() *user.Client
    
    // Automation
    AutomationAction() *automationaction.Client
    AutomationPolicy() *automationpolicy.Client
    
    // Infrastructure
    HostAgent() *hostagent.Client
    
    // Dashboards
    CustomDashboard() *customdashboard.Client
    
    // Tokens
    APIToken() *apitoken.Client
    
    // Maintenance
    MaintenanceWindow() *maintenancewindow.Client
}

// instanaAPI is the concrete implementation of InstanaAPI
type instanaAPI struct {
    restClient *client.Client
    
    // Cached client instances
    applicationAlertConfig   *applicationalertconfig.Client
    websiteAlertConfig       *websitealertconfig.Client
    mobileAlertConfig        *mobilealertconfig.Client
    infraAlertConfig         *infraalertconfig.Client
    logAlertConfig           *logalertconfig.Client
    sloAlertConfig           *sloalertconfig.Client
    syntheticAlertConfig     *syntheticalertconfig.Client
    applicationConfig        *applicationconfig.Client
    customEventSpec          *customeventspec.Client
    builtinEventSpec         *builtineventspec.Client
    alertingChannel          *alertingchannel.Client
    alertingConfig           *alertingconfig.Client
    sliConfig                *sliconfig.Client
    sloConfig                *sloconfig.Client
    sloCorrection            *slocorrection.Client
    syntheticTest            *synthetictest.Client
    syntheticLocation        *syntheticlocation.Client
    websiteMonitoring        *websitemonitoring.Client
    group                    *group.Client
    role                     *role.Client
    team                     *team.Client
    user                     *user.Client
    automationAction         *automationaction.Client
    automationPolicy         *automationpolicy.Client
    hostAgent                *hostagent.Client
    customDashboard          *customdashboard.Client
    apiToken                 *apitoken.Client
    maintenanceWindow        *maintenancewindow.Client
}

// NewAPI creates a new InstanaAPI instance
func NewAPI(cfg *config.Config) (InstanaAPI, error) {
    // Create REST client
    restClient, err := client.NewClient(cfg)
    if err != nil {
        return nil, err
    }
    
    return &instanaAPI{
        restClient: restClient,
    }, nil
}

// Lazy initialization methods for each API client

func (api *instanaAPI) ApplicationAlertConfig() *applicationalertconfig.Client {
    if api.applicationAlertConfig == nil {
        api.applicationAlertConfig = applicationalertconfig.NewClient(api.restClient)
    }
    return api.applicationAlertConfig
}

func (api *instanaAPI) WebsiteAlertConfig() *websitealertconfig.Client {
    if api.websiteAlertConfig == nil {
        api.websiteAlertConfig = websitealertconfig.NewClient(api.restClient)
    }
    return api.websiteAlertConfig
}

func (api *instanaAPI) MobileAlertConfig() *mobilealertconfig.Client {
    if api.mobileAlertConfig == nil {
        api.mobileAlertConfig = mobilealertconfig.NewClient(api.restClient)
    }
    return api.mobileAlertConfig
}

func (api *instanaAPI) InfraAlertConfig() *infraalertconfig.Client {
    if api.infraAlertConfig == nil {
        api.infraAlertConfig = infraalertconfig.NewClient(api.restClient)
    }
    return api.infraAlertConfig
}

func (api *instanaAPI) LogAlertConfig() *logalertconfig.Client {
    if api.logAlertConfig == nil {
        api.logAlertConfig = logalertconfig.NewClient(api.restClient)
    }
    return api.logAlertConfig
}

func (api *instanaAPI) SLOAlertConfig() *sloalertconfig.Client {
    if api.sloAlertConfig == nil {
        api.sloAlertConfig = sloalertconfig.NewClient(api.restClient)
    }
    return api.sloAlertConfig
}

func (api *instanaAPI) SyntheticAlertConfig() *syntheticalertconfig.Client {
    if api.syntheticAlertConfig == nil {
        api.syntheticAlertConfig = syntheticalertconfig.NewClient(api.restClient)
    }
    return api.syntheticAlertConfig
}

func (api *instanaAPI) ApplicationConfig() *applicationconfig.Client {
    if api.applicationConfig == nil {
        api.applicationConfig = applicationconfig.NewClient(api.restClient)
    }
    return api.applicationConfig
}

func (api *instanaAPI) CustomEventSpec() *customeventspec.Client {
    if api.customEventSpec == nil {
        api.customEventSpec = customeventspec.NewClient(api.restClient)
    }
    return api.customEventSpec
}

func (api *instanaAPI) BuiltinEventSpec() *builtineventspec.Client {
    if api.builtinEventSpec == nil {
        api.builtinEventSpec = builtineventspec.NewClient(api.restClient)
    }
    return api.builtinEventSpec
}

func (api *instanaAPI) AlertingChannel() *alertingchannel.Client {
    if api.alertingChannel == nil {
        api.alertingChannel = alertingchannel.NewClient(api.restClient)
    }
    return api.alertingChannel
}

func (api *instanaAPI) AlertingConfig() *alertingconfig.Client {
    if api.alertingConfig == nil {
        api.alertingConfig = alertingconfig.NewClient(api.restClient)
    }
    return api.alertingConfig
}

func (api *instanaAPI) SLIConfig() *sliconfig.Client {
    if api.sliConfig == nil {
        api.sliConfig = sliconfig.NewClient(api.restClient)
    }
    return api.sliConfig
}

func (api *instanaAPI) SLOConfig() *sloconfig.Client {
    if api.sloConfig == nil {
        api.sloConfig = sloconfig.NewClient(api.restClient)
    }
    return api.sloConfig
}

func (api *instanaAPI) SLOCorrection() *slocorrection.Client {
    if api.sloCorrection == nil {
        api.sloCorrection = slocorrection.NewClient(api.restClient)
    }
    return api.sloCorrection
}

func (api *instanaAPI) SyntheticTest() *synthetictest.Client {
    if api.syntheticTest == nil {
        api.syntheticTest = synthetictest.NewClient(api.restClient)
    }
    return api.syntheticTest
}

func (api *instanaAPI) SyntheticLocation() *syntheticlocation.Client {
    if api.syntheticLocation == nil {
        api.syntheticLocation = syntheticlocation.NewClient(api.restClient)
    }
    return api.syntheticLocation
}

func (api *instanaAPI) WebsiteMonitoring() *websitemonitoring.Client {
    if api.websiteMonitoring == nil {
        api.websiteMonitoring = websitemonitoring.NewClient(api.restClient)
    }
    return api.websiteMonitoring
}

func (api *instanaAPI) Group() *group.Client {
    if api.group == nil {
        api.group = group.NewClient(api.restClient)
    }
    return api.group
}

func (api *instanaAPI) Role() *role.Client {
    if api.role == nil {
        api.role = role.NewClient(api.restClient)
    }
    return api.role
}

func (api *instanaAPI) Team() *team.Client {
    if api.team == nil {
        api.team = team.NewClient(api.restClient)
    }
    return api.team
}

func (api *instanaAPI) User() *user.Client {
    if api.user == nil {
        api.user = user.NewClient(api.restClient)
    }
    return api.user
}

func (api *instanaAPI) AutomationAction() *automationaction.Client {
    if api.automationAction == nil {
        api.automationAction = automationaction.NewClient(api.restClient)
    }
    return api.automationAction
}

func (api *instanaAPI) AutomationPolicy() *automationpolicy.Client {
    if api.automationPolicy == nil {
        api.automationPolicy = automationpolicy.NewClient(api.restClient)
    }
    return api.automationPolicy
}

func (api *instanaAPI) HostAgent() *hostagent.Client {
    if api.hostAgent == nil {
        api.hostAgent = hostagent.NewClient(api.restClient)
    }
    return api.hostAgent
}

func (api *instanaAPI) CustomDashboard() *customdashboard.Client {
    if api.customDashboard == nil {
        api.customDashboard = customdashboard.NewClient(api.restClient)
    }
    return api.customDashboard
}

func (api *instanaAPI) APIToken() *apitoken.Client {
    if api.apiToken == nil {
        api.apiToken = apitoken.NewClient(api.restClient)
    }
    return api.apiToken
}

func (api *instanaAPI) MaintenanceWindow() *maintenancewindow.Client {
    if api.maintenanceWindow == nil {
        api.maintenanceWindow = maintenancewindow.NewClient(api.restClient)
    }
    return api.maintenanceWindow
}
```

## Usage Examples

### Example 1: Basic Usage

```go
package main

import (
    "context"
    "log"
    
    "github.com/instana/instana-go-client/config"
    "github.com/instana/instana-go-client/instana"
)

func main() {
    // Create configuration
    cfg, err := config.NewBuilder().
        WithBaseURL("https://tenant.instana.io").
        WithAPIToken("your-api-token").
        Build()
    if err != nil {
        log.Fatal(err)
    }
    
    // Create main API client
    api, err := instana.NewAPI(cfg)
    if err != nil {
        log.Fatal(err)
    }
    
    // Use specific API clients
    ctx := context.Background()
    
    // Get application alert config
    alertConfig, err := api.ApplicationAlertConfig().Get(ctx, "alert-id")
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Alert: %s", alertConfig.Name)
    
    // Get application config
    appConfig, err := api.ApplicationConfig().Get(ctx, "app-id")
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Application: %s", appConfig.Label)
}
```

### Example 2: Direct Package Usage (Alternative)

If you prefer to use API packages directly without the main facade:

```go
package main

import (
    "context"
    "log"
    
    "github.com/instana/instana-go-client/client"
    "github.com/instana/instana-go-client/config"
    "github.com/instana/instana-go-client/api/applicationalertconfig"
    "github.com/instana/instana-go-client/api/applicationconfig"
)

func main() {
    // Create configuration
    cfg, err := config.NewBuilder().
        WithBaseURL("https://tenant.instana.io").
        WithAPIToken("your-api-token").
        Build()
    if err != nil {
        log.Fatal(err)
    }
    
    // Create REST client
    restClient, err := client.NewClient(cfg)
    if err != nil {
        log.Fatal(err)
    }
    
    // Create specific API clients directly
    alertClient := applicationalertconfig.NewClient(restClient)
    appClient := applicationconfig.NewClient(restClient)
    
    // Use the clients
    ctx := context.Background()
    
    alertConfig, err := alertClient.Get(ctx, "alert-id")
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Alert: %s", alertConfig.Name)
    
    appConfig, err := appClient.Get(ctx, "app-id")
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Application: %s", appConfig.Label)
}
```

## Key Design Decisions

### 1. Main API Interface (`InstanaAPI`)
- **Single entry point** for all API operations
- **Lazy initialization** - API clients created only when first accessed
- **Clean interface** - Easy to discover available APIs
- **Type-safe** - Returns specific client types

### 2. No Backward Compatibility Layer
- **Clean slate** - No legacy code to maintain
- **Modern design** - Use latest Go best practices
- **Simple structure** - No compatibility shims

### 3. Two Usage Patterns

**Pattern A: Via Main API (Recommended for most users)**
```go
api, _ := instana.NewAPI(cfg)
alertClient := api.ApplicationAlertConfig()
```
- ✅ Convenient - Single entry point
- ✅ Discoverable - IDE autocomplete shows all APIs
- ✅ Consistent - Same pattern for all APIs

**Pattern B: Direct Package Import (For advanced users)**
```go
restClient, _ := client.NewClient(cfg)
alertClient := applicationalertconfig.NewClient(restClient)
```
- ✅ Minimal imports - Only import what you need
- ✅ Explicit - Clear dependencies
- ✅ Flexible - Direct control

### 4. Client Lifecycle
- **Stateless clients** - No internal state beyond REST client
- **Reusable** - Can be called multiple times
- **Thread-safe** - Safe for concurrent use
- **Context-aware** - All operations accept context

## Summary

The main application interface is:

1. **`instana.InstanaAPI`** - Main interface with methods for each API
2. **`instana.NewAPI(config)`** - Factory to create API instance
3. **Lazy-loaded clients** - Each API client created on first access
4. **No backward compatibility** - Clean, modern design
5. **Two usage patterns** - Via main API or direct package import

This design provides a clean, intuitive interface for applications while maintaining the granular, self-contained package structure you requested.