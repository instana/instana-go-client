#!/bin/bash
# Add ResourcePath constants to all API packages

# alertingchannel - already has it
# alertingconfig
sed -i '' '3a\
\
// ResourcePath is the path to the Alerting Configurations resource in the Instana API\
const ResourcePath = "/api/events/settings/alerts"\
' api/alertingconfig/alertingconfig.go

# apitoken
sed -i '' '3a\
\
// ResourcePath is the path to the API Tokens resource in the Instana API\
const ResourcePath = "/api/settings/api-tokens"\
' api/apitoken/apitoken.go

# applicationalertconfig
sed -i '' '8a\
\
// ResourcePath is the path to the Application Alert Configurations resource in the Instana API\
const ResourcePath = "/api/events/settings/application-alert-configs"\
' api/applicationalertconfig/applicationalertconfig.go

# applicationconfig
sed -i '' '8a\
\
// ResourcePath is the path to the Application Configurations resource in the Instana API\
const ResourcePath = "/api/application-monitoring/settings/application"\
' api/applicationconfig/applicationconfig.go

# automationaction
sed -i '' '3a\
\
// ResourcePath is the path to the Automation Actions resource in the Instana API\
const ResourcePath = "/api/automation/actions"\
' api/automationaction/automationaction.go

# automationpolicy
sed -i '' '8a\
\
// ResourcePath is the path to the Automation Policies resource in the Instana API\
const ResourcePath = "/api/automation/policies"\
' api/automationpolicy/automationpolicy.go

# builtineventspec
sed -i '' '3a\
\
// ResourcePath is the path to the Builtin Event Specifications resource in the Instana API\
const ResourcePath = "/api/events/settings/event-specifications/built-in"\
' api/builtineventspec/builtineventspec.go

# customdashboard
sed -i '' '3a\
\
// ResourcePath is the path to the Custom Dashboards resource in the Instana API\
const ResourcePath = "/api/custom-dashboard"\
' api/customdashboard/customdashboard.go

# customeventspec
sed -i '' '8a\
\
// ResourcePath is the path to the Custom Event Specifications resource in the Instana API\
const ResourcePath = "/api/events/settings/event-specifications/custom"\
' api/customeventspec/customeventspec.go

# group
sed -i '' '3a\
\
// ResourcePath is the path to the Groups resource in the Instana API\
const ResourcePath = "/api/settings/rbac/groups"\
' api/group/group.go

# hostagent
sed -i '' '5a\
\
// ResourcePath is the path to the Host Agents resource in the Instana API\
const ResourcePath = "/api/host-agent"\
' api/hostagent/hostagent.go

# infraalertconfig
sed -i '' '8a\
\
// ResourcePath is the path to the Infrastructure Alert Configurations resource in the Instana API\
const ResourcePath = "/api/events/settings/infra-alert-configs"\
' api/infraalertconfig/infraalertconfig.go

# logalertconfig
sed -i '' '8a\
\
// ResourcePath is the path to the Log Alert Configurations resource in the Instana API\
const ResourcePath = "/api/events/settings/log-alert-configs"\
' api/logalertconfig/logalertconfig.go

# maintenancewindow
sed -i '' '5a\
\
// ResourcePath is the path to the Maintenance Windows resource in the Instana API\
const ResourcePath = "/api/settings/maintenance-windows"\
' api/maintenancewindow/maintenancewindow.go

# mobilealertconfig
sed -i '' '3a\
\
// ResourcePath is the path to the Mobile Alert Configurations resource in the Instana API\
const ResourcePath = "/api/events/settings/mobile-alert-configs"\
' api/mobilealertconfig/mobilealertconfig.go

# role
sed -i '' '5a\
\
// ResourcePath is the path to the Roles resource in the Instana API\
const ResourcePath = "/api/settings/rbac/roles"\
' api/role/role.go

# sliconfig
sed -i '' '5a\
\
// ResourcePath is the path to the SLI Configurations resource in the Instana API\
const ResourcePath = "/api/settings/sli"\
' api/sliconfig/sliconfig.go

# sloalertconfig
sed -i '' '3a\
\
// ResourcePath is the path to the SLO Alert Configurations resource in the Instana API\
const ResourcePath = "/api/events/settings/slo-alert-configs"\
' api/sloalertconfig/sloalertconfig.go

# sloconfig
sed -i '' '3a\
\
// ResourcePath is the path to the SLO Configurations resource in the Instana API\
const ResourcePath = "/api/settings/slo"\
' api/sloconfig/sloconfig.go

# slocorrection
sed -i '' '3a\
\
// ResourcePath is the path to the SLO Correction Configurations resource in the Instana API\
const ResourcePath = "/api/settings/slo/correction"\
' api/slocorrection/slocorrection.go

# syntheticalertconfig
sed -i '' '3a\
\
// ResourcePath is the path to the Synthetic Alert Configurations resource in the Instana API\
const ResourcePath = "/api/events/settings/synthetic-alert-configs"\
' api/syntheticalertconfig/syntheticalertconfig.go

# syntheticlocation
sed -i '' '3a\
\
// ResourcePath is the path to the Synthetic Locations resource in the Instana API\
const ResourcePath = "/api/synthetics/settings/locations"\
' api/syntheticlocation/syntheticlocation.go

# synthetictest
sed -i '' '3a\
\
// ResourcePath is the path to the Synthetic Tests resource in the Instana API\
const ResourcePath = "/api/synthetics/settings/tests"\
' api/synthetictest/synthetictest.go

# team
sed -i '' '5a\
\
// ResourcePath is the path to the Teams resource in the Instana API\
const ResourcePath = "/api/settings/rbac/teams"\
' api/team/team.go

# user
sed -i '' '5a\
\
// ResourcePath is the path to the Users resource in the Instana API\
const ResourcePath = "/api/settings/users"\
' api/user/user.go

# websitealertconfig
sed -i '' '8a\
\
// ResourcePath is the path to the Website Alert Configurations resource in the Instana API\
const ResourcePath = "/api/events/settings/website-alert-configs"\
' api/websitealertconfig/websitealertconfig.go

# websitemonitoring
sed -i '' '3a\
\
// ResourcePath is the path to the Website Monitoring Configurations resource in the Instana API\
const ResourcePath = "/api/website-monitoring"\
' api/websitemonitoring/websitemonitoring.go

echo "✅ ResourcePath constants added to all packages"
