# API Package Recovery Mapping

This document maps Terraform provider files to instana-go-client API packages.

## Package Mappings

| API Package | Source Files (internal/restapi/) | Notes |
|-------------|----------------------------------|-------|
| alertingchannel | alerting-channels-api.go, alerting-channel-type.go | ✅ RECOVERED |
| alertingconfig | alerts-api.go, alert-event-type.go | Pending |
| apitoken | api-tokens-api.go | Pending |
| applicationalertconfig | application-alert-config.go, application-alert-config-types.go, application-alert-rule.go, application-alert-evaluation-type.go | Pending |
| applicationconfig | application-configs-api.go, application-config-scope.go, included-application.go | Pending |
| automationaction | automation-action-api.go | Pending |
| automationpolicy | automation-policy-api.go | Pending |
| builtineventspec | builtin-event-specification-api.go | Pending |
| customdashboard | custom-dashboard.go | Pending |
| customeventspec | custom-event-specficiations-api.go | Pending |
| group | groups-api.go | Pending |
| hostagent | host-agents-api.go, host-agent-json-unmarshaller.go | Pending |
| infraalertconfig | infra-alert-configs.go, infra-alert-rule.go, infra-alert-evaluation-type.go, infra-time-threshold.go | Pending |
| logalertconfig | log-alert-config-api.go, log-level.go | Pending |
| maintenancewindow | maintenance-window-config-api.go | Pending |
| mobilealertconfig | mobile-alert-config.go | Pending |
| role | roles-api.go | Pending |
| sliconfig | sli-config-api.go | Pending |
| sloalertconfig | slo-alert-config-api.go | Pending |
| sloconfig | slo-config-api.go | Pending |
| slocorrection | slo-correction-config-api.go | Pending |
| syntheticalertconfig | synthetic-alert-config.go | Pending |
| syntheticlocation | synthetic-location.go | Pending |
| synthetictest | synthetic-test.go, synthetic-test-rest-resource.go | Pending |
| team | teams-api.go | Pending |
| user | users-api.go | Pending |
| websitealertconfig | website-alert-config.go, website-alert-rule.go, website-impact-measurement-method.go | Pending |
| websitemonitoring | website-monitoring-config-api.go, website-monitoring-config-rest-resource.go, website-time-threshold.go | Pending |

## Shared Types (already in shared/types/)

These types are already extracted to shared packages:
- access-rule.go, access-type.go → shared/types/
- aggregation.go → shared/types/
- boundary-scope.go → shared/types/
- custom-payload-field.go → shared/types/
- granularity.go → shared/types/
- operator.go → shared/types/
- relation-type.go → shared/types/
- rule-with-threshold.go → shared/types/
- severity.go → shared/types/
- tag-filter.go → shared/tagfilter/
- threshold.go → shared/types/

