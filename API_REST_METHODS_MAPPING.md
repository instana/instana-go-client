# API REST Methods Mapping

This document maps each API to its exact REST resource type from the original implementation.

## REST Resource Types

1. **NewCreatePOSTUpdatePUTRestResource** - POST for create, PUT for update
2. **NewCreatePOSTUpdatePOSTRestResource** - POST for both create and update
3. **NewCreatePUTUpdatePUTRestResource** - PUT for both create and update (upsert)
4. **NewCreatePOSTUpdateNotSupportedRestResource** - POST for create, update not supported
5. **NewReadOnlyRestResource** - Read-only operations

## API Mappings (from Instana-api.go)

| API Name | Package Name | REST Method | Factory Function |
|----------|--------------|-------------|------------------|
| APITokens | apitoken | POST/PUT | NewCreatePOSTUpdatePUTRestResource |
| AlertingChannels | alertingchannel | PUT/PUT | NewCreatePUTUpdatePUTRestResource |
| AlertingConfigurations | alertingconfig | PUT/PUT | NewCreatePUTUpdatePUTRestResource |
| ApplicationConfigs | applicationconfig | POST/PUT | NewCreatePOSTUpdatePUTRestResource |
| ApplicationAlertConfigs | applicationalertconfig | POST/POST | NewCreatePOSTUpdatePOSTRestResource |
| GlobalApplicationAlertConfigs | applicationalertconfig | POST/POST | NewCreatePOSTUpdatePOSTRestResource |
| AutomationActions | automationaction | POST/PUT | NewCreatePOSTUpdatePUTRestResource |
| AutomationPolicies | automationpolicy | POST/PUT | NewCreatePOSTUpdatePUTRestResource |
| BuiltinEventSpecifications | builtineventspec | READ-ONLY | NewReadOnlyRestResource |
| CustomDashboards | customdashboard | POST/PUT | NewCreatePOSTUpdatePUTRestResource |
| CustomEventSpecifications | customeventspec | PUT/PUT | NewCreatePUTUpdatePUTRestResource |
| Groups | group | POST/PUT | NewCreatePOSTUpdatePUTRestResource |
| HostAgents | hostagent | READ-ONLY | NewReadOnlyRestResource |
| InfraAlertConfig | infraalertconfig | POST/POST | NewCreatePOSTUpdatePOSTRestResource |
| LogAlertConfig | logalertconfig | POST/POST | NewCreatePOSTUpdatePOSTRestResource |
| MaintenanceWindowConfigs | maintenancewindow | PUT/PUT | NewCreatePUTUpdatePUTRestResource |
| MobileAlertConfig | mobilealertconfig | POST/POST | NewCreatePOSTUpdatePOSTRestResource |
| Roles | role | POST/PUT | NewCreatePOSTUpdatePUTRestResource |
| SliConfigs | sliconfig | POST/NO-UPDATE | NewCreatePOSTUpdateNotSupportedRestResource |
| SloConfigs | sloconfig | POST/PUT | NewCreatePOSTUpdatePUTRestResource |
| SloAlertConfig | sloalertconfig | POST/POST | NewCreatePOSTUpdatePOSTRestResource |
| SloCorrectionConfig | slocorrection | POST/PUT | NewCreatePOSTUpdatePUTRestResource |
| SyntheticAlertConfigs | syntheticalertconfig | POST/POST | NewCreatePOSTUpdatePOSTRestResource |
| SyntheticLocations | syntheticlocation | READ-ONLY | NewReadOnlyRestResource |
| SyntheticTests | synthetictest | POST/PUT | NewCreatePOSTUpdatePUTRestResource |
| Teams | team | POST/PUT | NewCreatePOSTUpdatePUTRestResource |
| Users | user | READ-ONLY | NewReadOnlyRestResource |
| WebsiteAlertConfig | websitealertconfig | POST/POST | NewCreatePOSTUpdatePOSTRestResource |
| WebsiteMonitoringConfigs | websitemonitoring | POST/PUT | NewCreatePOSTUpdatePUTRestResource |

## Summary by REST Method

### PUT/PUT (Upsert - 4 APIs)
- alertingchannel
- alertingconfig
- customeventspec
- maintenancewindow

### POST/PUT (14 APIs)
- apitoken ✅ (template created)
- applicationconfig
- automationaction
- automationpolicy
- customdashboard
- group
- role
- sloconfig
- slocorrection
- synthetictest
- team
- websitemonitoring

### POST/POST (8 APIs)
- applicationalertconfig
- infraalertconfig
- logalertconfig
- mobilealertconfig
- sloalertconfig
- syntheticalertconfig
- websitealertconfig

### POST/NO-UPDATE (1 API)
- sliconfig

### READ-ONLY (4 APIs)
- builtineventspec
- hostagent
- syntheticlocation
- user

## Notes

- Some APIs use CustomPayloadFieldsUnmarshallerAdapter wrapper
- GlobalApplicationAlertConfigs shares model with ApplicationAlertConfigs
- All alert configs (except maintenance window) use POST/POST pattern