# Changelog

## [v1.3.0](https://github.com/instana/instana-go-client/releases/tag/v1.3.0) - 2026-08-17

### Added
- **Synthetic Credentials**: Added support for the Synthetic Credentials API (`/api/synthetics/settings/credentials`)
  - `SyntheticCredential` struct with `CredentialName`, `CredentialValue`, `Applications`, `MobileApps`, `Websites`, and `RbacTags` fields
  - Custom REST resource implementation that re-fetches the full object via `/associations` after create/update, working around the empty-body API responses from `POST`/`PUT`
  - `GetOne` reads from the `/associations` sub-path to return the full credential including scope associations
  - `SyntheticCredentials()` accessor added to the `InstanaAPI` interface and client with lazy initialisation


## [v1.2.0](https://github.com/instana/instana-go-client/releases/tag/v1.2.0) - 2026-08-17

### Added
- **Session Settings**: Added support for the Session Settings API (`/api/settings/session`)
  - `SessionSettings` struct with `TokenLifeTimeInMillis` and `IdleTimeInMillis` fields
  - Singleton REST resource client with get/update operations
  - Server-enforced constraints: token lifetime 10 min–7 days, idle time 10 min–8 hours


## [v1.1.4](https://github.com/instana/instana-go-client/releases/tag/v1.1.4) - 2026-07-28

### Added
- **Group Mapping (RBAC)**: Added support for the Group Mapping API (`/api/settings/rbac/mappings`)
  - `GroupMapping` struct mapping IdP (LDAP, OIDC, SAML) attribute key/value pairs to Instana groups
  - Optional `TeamID` field for team-scoped mappings


## [v1.1.3](https://github.com/instana/instana-go-client/releases/tag/v1.1.3) - 2026-07-13

### Added
- **Mobile SLO Support**: Extended SLO configuration with mobile entity support
  - `SloMobileEntity` struct with mobile app IDs and tag filter expression
  - `NewSloMobileEntity` constructor defaulting filter to an empty AND expression
  - `MobileIds` field added to `SloEntity`
  - Blueprint type constants: `latency`, `availability`, `traffic`, `saturation`, `custom`, `advanced-custom`
  - `SloIndicator` extended with `Metric`, `GoodEvents`, and `BadEvents` fields for advanced custom blueprints


## [v1.1.2](https://github.com/instana/instana-go-client/releases/tag/v1.1.2) - 2026-07-07

### Added
- **RBAC Tags**: Added `RbacTags` field to `AlertingChannel`, `CustomDashboard`, and `SyntheticTest` API structs
  - Enables team/RBAC tag assignment for these resources in terraform-provider-instana
  - `SyntheticTest.RbacTags` type updated from `[]ApiTag` to `[]RbacTag` to align with the `RbacTag` struct


## [v1.1.1](https://github.com/instana/instana-go-client/releases/tag/v1.1.1) - 2026-07-01

### Added
- **Grace Period**: Added `grace_period` field support to `InfraAlertConfig` and `WebsiteAlertConfig`
- **Apdex Smart Alerts**: Extended `SloAlertConfig` with Apdex smart alert support
  - `SloAlertRuleAlertTypeApdex` constant (`"APDEX"`)
  - `SloAlertMetricScore` constant (`"SCORE"`)
  - `ApdexIds` field on `SloAlertConfig` for linking Apdex configurations to SLO alerts


## [v1.1.0](https://github.com/instana/instana-go-client/releases/tag/v1.1.0) - 2026-06-29

### Added
- **Apdex V2 Configuration Support**: Complete implementation of Apdex configuration management
  - Support for Application Apdex configurations with boundary scope options
  - Support for Website Apdex configurations with beacon types


## [v1.0.2](https://github.com/instana/instana-go-client/releases/tag/v1.0.2) - 2026-06-12

### Added
- **API Token Permissions**: Added 7 missing permission fields to API Token resource
  - `CanCollectNetTraceLogs` - Permission to collect network trace logs
  - `CanConfigureCustomEntities` - Permission to configure custom entities
  - `CanConfigureWebsiteConversions` - Permission to configure website conversions
  - `CanConfigureIPFiltering` - Permission to configure IP filtering
  - `CanConfigureLlmModelPrice` - Permission to configure LLM model pricing
  - `CanConfigurePersonallyIdentifiableInformationMasking` - Permission to configure PII masking
  - `CanDownloadAgentConfiguration` - Permission to download agent configuration


## [v1.0.1](https://github.com/instana/instana-go-client/releases/tag/v1.0.1) - 2026-06-08

### Added
- Added missing permissions for Roles
- Enhanced SLO resource handling


## [v1.0.0](https://github.com/instana/instana-go-client/releases/tag/v1.0.0) - 2026-04-14

### Added

#### API Resource Support
- **Alerting Channels**: Email, Slack, PagerDuty, Webhook, and more
- **Alerting Configurations**: Alert rule management
- **API Tokens**: Token lifecycle management
- **Application Alerts**: Application-level alert configurations
- **Application Configs**: Application monitoring configuration
- **Automation Actions**: Automated response actions
- **Automation Policies**: Policy-based automation rules
- **Built-in Event Specs**: Pre-defined event specifications
- **Custom Dashboards**: Dashboard creation and management
- **Custom Event Specs**: Custom event definitions
- **Groups**: RBAC group management
- **Host Agents**: Agent discovery and management
- **Infrastructure Alerts**: Infrastructure monitoring alerts
- **Log Alerts**: Log-based alerting
- **Maintenance Windows**: Scheduled maintenance configuration
- **Mobile Alerts**: Mobile app alert configurations
- **Mobile App Configs**: Mobile application monitoring
- **Roles**: RBAC role management
- **SLI Configs**: Service Level Indicator configurations
- **SLO Alerts**: Service Level Objective alerting
- **SLO Configs**: SLO definition and management
- **SLO Corrections**: SLO correction windows
- **Synthetic Alerts**: Synthetic monitoring alerts
- **Synthetic Locations**: Synthetic test location management
- **Synthetic Tests**: Synthetic test configuration
- **Teams**: Team management for RBAC
- **Users**: User account management
- **Website Alerts**: Website monitoring alerts
- **Website Monitoring**: Website monitoring configuration
