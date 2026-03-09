package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
	"github.com/instana/instana-go-client/shared/types"
)

func TestApplicationConfigResourcePath(t *testing.T) {
	expected := "/api/application-monitoring/settings/application"
	if ApplicationConfigsResourcePath != expected {
		t.Errorf("Expected ApplicationConfigsResourcePath to be %s, got %s", expected, ApplicationConfigsResourcePath)
	}
}

func TestApplicationConfigGetIDForResourcePath(t *testing.T) {
	testID := "test-app-config-123"
	config := &ApplicationConfig{
		ID:    testID,
		Label: "Test Application",
	}

	result := config.GetIDForResourcePath()
	if result != testID {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", testID, result)
	}
}

func TestApplicationConfigStructure(t *testing.T) {
	config := ApplicationConfig{
		ID:            "app-config-123",
		Label:         "Production Application",
		Scope:         ApplicationConfigScopeIncludeAllDownstream,
		BoundaryScope: types.BoundaryScopeAll,
		AccessRules:   []types.AccessRule{},
	}

	if config.ID != "app-config-123" {
		t.Errorf("Expected ID 'app-config-123', got %s", config.ID)
	}
	if config.Label != "Production Application" {
		t.Errorf("Expected Label 'Production Application', got %s", config.Label)
	}
	if config.Scope != ApplicationConfigScopeIncludeAllDownstream {
		t.Errorf("Expected Scope 'INCLUDE_ALL_DOWNSTREAM', got %s", config.Scope)
	}
	if config.BoundaryScope != types.BoundaryScopeAll {
		t.Errorf("Expected BoundaryScope 'ALL', got %s", config.BoundaryScope)
	}
}

func TestApplicationConfigScopeConstants(t *testing.T) {
	tests := []struct {
		name     string
		value    ApplicationConfigScope
		expected string
	}{
		{"IncludeNoDownstream", ApplicationConfigScopeIncludeNoDownstream, "INCLUDE_NO_DOWNSTREAM"},
		{"IncludeImmediateDownstream", ApplicationConfigScopeIncludeImmediateDownstreamDatabaseAndMessaging, "INCLUDE_IMMEDIATE_DOWNSTREAM_DATABASE_AND_MESSAGING"},
		{"IncludeAllDownstream", ApplicationConfigScopeIncludeAllDownstream, "INCLUDE_ALL_DOWNSTREAM"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.value) != tt.expected {
				t.Errorf("Expected %s to be %s, got %s", tt.name, tt.expected, string(tt.value))
			}
		})
	}
}

func TestSupportedApplicationConfigScopes(t *testing.T) {
	if len(SupportedApplicationConfigScopes) != 3 {
		t.Errorf("Expected 3 supported scopes, got %d", len(SupportedApplicationConfigScopes))
	}

	stringSlice := SupportedApplicationConfigScopes.ToStringSlice()
	if len(stringSlice) != 3 {
		t.Errorf("Expected 3 string values, got %d", len(stringSlice))
	}
}

func TestIncludedEndpointStructure(t *testing.T) {
	endpoint := IncludedEndpoint{
		EndpointID: "endpoint-123",
		Inclusive:  true,
	}

	if endpoint.EndpointID != "endpoint-123" {
		t.Errorf("Expected EndpointID 'endpoint-123', got %s", endpoint.EndpointID)
	}
	if !endpoint.Inclusive {
		t.Error("Expected Inclusive to be true")
	}
}

func TestIncludedServiceStructure(t *testing.T) {
	service := IncludedService{
		ServiceID: "service-456",
		Inclusive: false,
		Endpoints: map[string]IncludedEndpoint{
			"endpoint-1": {EndpointID: "endpoint-1", Inclusive: true},
		},
	}

	if service.ServiceID != "service-456" {
		t.Errorf("Expected ServiceID 'service-456', got %s", service.ServiceID)
	}
	if service.Inclusive {
		t.Error("Expected Inclusive to be false")
	}
	if len(service.Endpoints) != 1 {
		t.Errorf("Expected 1 endpoint, got %d", len(service.Endpoints))
	}
}

func TestIncludedApplicationStructure(t *testing.T) {
	app := IncludedApplication{
		ApplicationID: "app-789",
		Inclusive:     true,
		Services: map[string]IncludedService{
			"service-1": {ServiceID: "service-1", Inclusive: true},
		},
	}

	if app.ApplicationID != "app-789" {
		t.Errorf("Expected ApplicationID 'app-789', got %s", app.ApplicationID)
	}
	if !app.Inclusive {
		t.Error("Expected Inclusive to be true")
	}
	if len(app.Services) != 1 {
		t.Errorf("Expected 1 service, got %d", len(app.Services))
	}
}
