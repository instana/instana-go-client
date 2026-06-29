# Changelog

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
