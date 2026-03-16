#!/bin/bash

# Script to generate API packages following the template pattern
# This automates the creation of 27 remaining API packages

set -e

# Color output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}=== Instana Go Client - API Package Generator ===${NC}"
echo ""

# Base directory
BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
API_DIR="${BASE_DIR}/api"

# Function to create a basic API package structure
create_api_package() {
    local package_name=$1
    local model_name=$2
    local resource_path=$3
    local rest_method=$4  # "POST_PUT" or "READ_ONLY"
    local description=$5
    
    local pkg_dir="${API_DIR}/${package_name}"
    
    echo -e "${YELLOW}Creating package: ${package_name}${NC}"
    
    # Create directory
    mkdir -p "${pkg_dir}"
    
    # Create constants.go
    cat > "${pkg_dir}/constants.go" <<EOF
package ${package_name}

// ResourcePath is the path to ${description} resource of Instana RESTful API
const ResourcePath = "${resource_path}"
EOF
    
    # Create model.go (placeholder - needs manual completion)
    cat > "${pkg_dir}/model.go" <<EOF
package ${package_name}

// ${model_name} is the representation of ${description}
type ${model_name} struct {
    // TODO: Copy fields from instana/${package_name}*.go
    ID string \`json:"id"\`
}

// GetIDForResourcePath implementation of the interface InstanaDataObject
func (r *${model_name}) GetIDForResourcePath() string {
    return r.ID
}
EOF
    
    # Create unmarshaller.go
    cat > "${pkg_dir}/unmarshaller.go" <<EOF
package ${package_name}

import (
    "encoding/json"
    "fmt"
)

// NewUnmarshaller creates a new instance of a JSONUnmarshaller for ${model_name}
func NewUnmarshaller() *Unmarshaller {
    return &Unmarshaller{}
}

// Unmarshaller is a JSONUnmarshaller implementation for ${model_name}
type Unmarshaller struct{}

// Unmarshal converts JSON bytes into a ${model_name}
func (u *Unmarshaller) Unmarshal(data []byte) (*${model_name}, error) {
    var target ${model_name}
    if err := json.Unmarshal(data, &target); err != nil {
        return &target, fmt.Errorf("failed to parse json; %s", err)
    }
    return &target, nil
}

// UnmarshalArray converts JSON bytes into a slice of ${model_name}
func (u *Unmarshaller) UnmarshalArray(data []byte) (*[]*${model_name}, error) {
    var target []*${model_name}
    if err := json.Unmarshal(data, &target); err != nil {
        return &target, fmt.Errorf("failed to parse json; %s", err)
    }
    return &target, nil
}
EOF
    
    # Create client.go based on REST method
    if [ "$rest_method" = "READ_ONLY" ]; then
        cat > "${pkg_dir}/client.go" <<EOF
package ${package_name}

import "github.com/instana/instana-go-client/shared/rest"

// NewClient creates a new read-only API client for ${description}
func NewClient(restClient rest.RestClient) rest.ReadOnlyRestResource[*${model_name}] {
    return rest.NewReadOnlyRestResource(
        ResourcePath,
        NewUnmarshaller(),
        restClient,
    )
}
EOF
    elif [ "$rest_method" = "PUT_PUT" ]; then
        cat > "${pkg_dir}/client.go" <<EOF
package ${package_name}

import "github.com/instana/instana-go-client/shared/rest"

// NewClient creates a new API client for ${description}
func NewClient(restClient rest.RestClient) rest.RestResource[*${model_name}] {
    return rest.NewCreatePUTUpdatePUTRestResource(
        ResourcePath,
        NewUnmarshaller(),
        restClient,
    )
}
EOF
    elif [ "$rest_method" = "POST_POST" ]; then
        cat > "${pkg_dir}/client.go" <<EOF
package ${package_name}

import "github.com/instana/instana-go-client/shared/rest"

// NewClient creates a new API client for ${description}
func NewClient(restClient rest.RestClient) rest.RestResource[*${model_name}] {
    return rest.NewCreatePOSTUpdatePOSTRestResource(
        ResourcePath,
        NewUnmarshaller(),
        restClient,
    )
}
EOF
    elif [ "$rest_method" = "POST_NO_UPDATE" ]; then
        cat > "${pkg_dir}/client.go" <<EOF
package ${package_name}

import "github.com/instana/instana-go-client/shared/rest"

// NewClient creates a new API client for ${description}
func NewClient(restClient rest.RestClient) rest.RestResource[*${model_name}] {
    return rest.NewCreatePOSTUpdateNotSupportedRestResource(
        ResourcePath,
        NewUnmarshaller(),
        restClient,
    )
}
EOF
    else  # POST_PUT (default)
        cat > "${pkg_dir}/client.go" <<EOF
package ${package_name}

import "github.com/instana/instana-go-client/shared/rest"

// NewClient creates a new API client for ${description}
func NewClient(restClient rest.RestClient) rest.RestResource[*${model_name}] {
    return rest.NewCreatePOSTUpdatePUTRestResource(
        ResourcePath,
        NewUnmarshaller(),
        restClient,
    )
}
EOF
    fi
    
    # Create doc.go
    cat > "${pkg_dir}/doc.go" <<EOF
// Package ${package_name} provides the API client for managing ${description}.
//
// This package handles all operations related to ${description} in Instana.
//
// Example usage:
//
//  // Create a new client
//  client := ${package_name}.NewClient(restClient)
//
//  // Get all items
//  items, err := client.GetAll()
//  if err != nil {
//      // handle error
//  }
//
//  // Get a specific item
//  item, err := client.GetOne("item-id")
//  if err != nil {
//      // handle error
//  }
package ${package_name}
EOF
    
    echo -e "${GREEN}✓ Created ${package_name}${NC}"
    echo "  ${pkg_dir}/"
    echo "  ├── constants.go"
    echo "  ├── model.go (TODO: Complete model fields)"
    echo "  ├── unmarshaller.go"
    echo "  ├── client.go"
    echo "  └── doc.go"
    echo ""
}

# Generate all API packages
echo -e "${YELLOW}Generating API packages...${NC}"
echo ""

# Package definitions: name, ModelName, resource_path, rest_method, description
# REST methods verified from Instana-api.go - see API_REST_METHODS_MAPPING.md

# PUT/PUT (Upsert)
create_api_package "alertingchannel" "AlertingChannel" "/api/events/settings/alertingChannels" "PUT_PUT" "Instana alerting channels"
create_api_package "alertingconfig" "AlertingConfiguration" "/api/events/settings/alerts" "PUT_PUT" "Instana alerting configurations"
create_api_package "customeventspec" "CustomEventSpecification" "/api/events/settings/event-specifications/custom" "PUT_PUT" "custom event specifications"
create_api_package "maintenancewindow" "MaintenanceWindow" "/api/events/settings/maintenance-windows" "PUT_PUT" "maintenance windows"

# POST/PUT
create_api_package "applicationconfig" "ApplicationConfig" "/api/application-monitoring/settings/application" "POST_PUT" "application configurations"
create_api_package "automationaction" "AutomationAction" "/api/automation/actions" "POST_PUT" "automation actions"
create_api_package "automationpolicy" "AutomationPolicy" "/api/automation/policies" "POST_PUT" "automation policies"
create_api_package "customdashboard" "CustomDashboard" "/api/custom-dashboard" "POST_PUT" "custom dashboards"
create_api_package "group" "Group" "/api/settings/groups" "POST_PUT" "RBAC groups"
create_api_package "role" "Role" "/api/settings/roles" "POST_PUT" "RBAC roles"
create_api_package "sloconfig" "SLOConfig" "/api/slo/config" "POST_PUT" "SLO configurations"
create_api_package "slocorrection" "SLOCorrection" "/api/slo/config/correction" "POST_PUT" "SLO correction configurations"
create_api_package "synthetictest" "SyntheticTest" "/api/synthetics/settings/tests" "POST_PUT" "synthetic tests"
create_api_package "team" "Team" "/api/settings/teams" "POST_PUT" "RBAC teams"
create_api_package "websitemonitoring" "WebsiteMonitoringConfig" "/api/website-monitoring/config" "POST_PUT" "website monitoring configurations"

# POST/POST
create_api_package "applicationalertconfig" "ApplicationAlertConfig" "/api/application-monitoring/settings/alerts" "POST_POST" "application alert configurations"
create_api_package "infraalertconfig" "InfrastructureAlertConfig" "/api/events/settings/alerts" "POST_POST" "infrastructure alert configurations"
create_api_package "logalertconfig" "LogAlertConfig" "/api/events/settings/alerts" "POST_POST" "log alert configurations"
create_api_package "mobilealertconfig" "MobileAlertConfig" "/api/mobile-app-monitoring/settings/alerts" "POST_POST" "mobile app alert configurations"
create_api_package "sloalertconfig" "SLOAlertConfig" "/api/events/settings/alerts" "POST_POST" "SLO alert configurations"
create_api_package "syntheticalertconfig" "SyntheticAlertConfig" "/api/events/settings/alerts" "POST_POST" "synthetic alert configurations"
create_api_package "websitealertconfig" "WebsiteAlertConfig" "/api/website-monitoring/settings/alerts" "POST_POST" "website alert configurations"

# POST/NO-UPDATE
create_api_package "sliconfig" "SLIConfig" "/api/sli/config" "POST_NO_UPDATE" "SLI configurations"

# READ-ONLY
create_api_package "builtineventspec" "BuiltinEventSpecification" "/api/events/settings/event-specifications/built-in" "READ_ONLY" "built-in event specifications"
create_api_package "hostagent" "HostAgent" "/api/host-agent" "READ_ONLY" "host agents"
create_api_package "syntheticlocation" "SyntheticLocation" "/api/synthetics/settings/locations" "READ_ONLY" "synthetic test locations"
create_api_package "user" "User" "/api/settings/users" "READ_ONLY" "users"

echo ""
echo -e "${GREEN}=== Package Generation Complete ===${NC}"
echo ""
echo -e "${YELLOW}NEXT STEPS:${NC}"
echo "1. For each package in api/, complete the model.go file by copying fields from instana/*-api.go"
echo "2. Verify the ResourcePath constants match the original paths"
echo "3. Check if REST method (POST_PUT vs READ_ONLY) is correct"
echo "4. Run 'go build ./...' to verify compilation"
echo ""
echo -e "${YELLOW}TODO: Complete model fields for these packages:${NC}"
ls -1 "${API_DIR}" | grep -v "apitoken" | while read pkg; do
    echo "  - api/${pkg}/model.go"
done
echo ""

# Made with Bob
