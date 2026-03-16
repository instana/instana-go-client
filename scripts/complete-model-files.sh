#!/bin/bash

# Script to complete model.go files by copying struct definitions from instana/ package
# This automates the migration of data models to the new package structure

set -e

# Color output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${GREEN}=== Completing Model Files ===${NC}"
echo ""

# Base directories
BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
INSTANA_DIR="${BASE_DIR}/instana"
API_DIR="${BASE_DIR}/api"

# Function to extract struct definition from instana file and update api model.go
complete_model() {
    local package_name=$1
    local model_name=$2
    local source_file=$3
    
    local model_file="${API_DIR}/${package_name}/model.go"
    local source_path="${INSTANA_DIR}/${source_file}"
    
    if [ ! -f "$source_path" ]; then
        echo -e "${RED}✗ Source file not found: ${source_file}${NC}"
        return 1
    fi
    
    echo -e "${BLUE}Processing: ${package_name}/${model_name}${NC}"
    
    # Extract the struct definition from source file
    # This uses awk to extract from "type ModelName struct {" to the closing "}"
    struct_def=$(awk "/^type ${model_name} struct \{/,/^}/" "$source_path")
    
    if [ -z "$struct_def" ]; then
        echo -e "${RED}✗ Could not find struct ${model_name} in ${source_file}${NC}"
        return 1
    fi
    
    # Extract GetIDForResourcePath method
    getid_method=$(awk "/^func \(r \*${model_name}\) GetIDForResourcePath\(\)/,/^}/" "$source_path")
    
    # Create the complete model.go file
    cat > "$model_file" <<EOF
package ${package_name}

${struct_def}

${getid_method}
EOF
    
    echo -e "${GREEN}✓ Completed ${package_name}/model.go${NC}"
}

# Complete all model files
echo -e "${YELLOW}Completing model files...${NC}"
echo ""

# Already complete
echo -e "${GREEN}✓ apitoken/model.go (already complete)${NC}"

# Complete remaining 27 packages
complete_model "alertingchannel" "AlertingChannel" "alerting-channels-api.go"
complete_model "alertingconfig" "AlertingConfiguration" "alerts-api.go"
complete_model "applicationalertconfig" "ApplicationAlertConfig" "application-alert-config.go"
complete_model "applicationconfig" "ApplicationConfig" "application-configs-api.go"
complete_model "automationaction" "AutomationAction" "automation-action-api.go"
complete_model "automationpolicy" "AutomationPolicy" "automation-policy-api.go"
complete_model "builtineventspec" "BuiltinEventSpecification" "builtin-event-specification-api.go"
complete_model "customdashboard" "CustomDashboard" "custom-dashboard.go"
complete_model "customeventspec" "CustomEventSpecification" "custom-event-specficiations-api.go"
complete_model "group" "Group" "groups-api.go"
complete_model "hostagent" "HostAgent" "host-agents-api.go"
complete_model "infraalertconfig" "InfraAlertConfig" "infra-alert-configs.go"
complete_model "logalertconfig" "LogAlertConfig" "log-alert-config-api.go"
complete_model "maintenancewindow" "MaintenanceWindowConfig" "maintenance-window-config-api.go"
complete_model "mobilealertconfig" "MobileAlertConfig" "mobile-alert-config.go"
complete_model "role" "Role" "roles-api.go"
complete_model "sliconfig" "SliConfig" "sli-config-api.go"
complete_model "sloalertconfig" "SloAlertConfig" "slo-alert-config-api.go"
complete_model "sloconfig" "SloConfig" "slo-config-api.go"
complete_model "slocorrection" "SloCorrectionConfig" "slo-correction-config-api.go"
complete_model "syntheticalertconfig" "SyntheticAlertConfig" "synthetic-alert-config.go"
complete_model "syntheticlocation" "SyntheticLocation" "synthetic-location.go"
complete_model "synthetictest" "SyntheticTest" "synthetic-test.go"
complete_model "team" "Team" "teams-api.go"
complete_model "user" "User" "users-api.go"
complete_model "websitealertconfig" "WebsiteAlertConfig" "website-alert-config.go"
complete_model "websitemonitoring" "WebsiteMonitoringConfig" "website-monitoring-config-api.go"

echo ""
echo -e "${GREEN}=== Model Files Completion Complete ===${NC}"
echo ""
echo -e "${YELLOW}NEXT STEPS:${NC}"
echo "1. Review the generated model.go files"
echo "2. Check for any missing type dependencies"
echo "3. Run 'go build ./api/...' to verify compilation"
echo ""

# Made with Bob
