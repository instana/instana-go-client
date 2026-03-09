package testutils

import (
	"testing"
)

// TestMockInstanaAPI_AllMethods tests that all methods return nil as expected
func TestMockInstanaAPI_AllMethods(t *testing.T) {
	mock := &MockInstanaAPI{}

	// Test all methods return nil
	tests := []struct {
		name   string
		result interface{}
	}{
		{"CustomEventSpecifications", mock.CustomEventSpecifications()},
		{"BuiltinEventSpecifications", mock.BuiltinEventSpecifications()},
		{"APITokens", mock.APITokens()},
		{"ApplicationConfigs", mock.ApplicationConfigs()},
		{"ApplicationAlertConfigs", mock.ApplicationAlertConfigs()},
		{"GlobalApplicationAlertConfigs", mock.GlobalApplicationAlertConfigs()},
		{"AlertingChannels", mock.AlertingChannels()},
		{"AlertingConfigurations", mock.AlertingConfigurations()},
		{"SliConfigs", mock.SliConfigs()},
		{"SloConfigs", mock.SloConfigs()},
		{"SloAlertConfig", mock.SloAlertConfig()},
		{"SloCorrectionConfig", mock.SloCorrectionConfig()},
		{"WebsiteMonitoringConfig", mock.WebsiteMonitoringConfig()},
		{"WebsiteAlertConfig", mock.WebsiteAlertConfig()},
		{"InfraAlertConfig", mock.InfraAlertConfig()},
		{"MobileAlertConfig", mock.MobileAlertConfig()},
		{"MaintenanceWindows", mock.MaintenanceWindows()},
		{"Teams", mock.Teams()},
		{"Groups", mock.Groups()},
		{"Roles", mock.Roles()},
		{"CustomDashboards", mock.CustomDashboards()},
		{"SyntheticTest", mock.SyntheticTest()},
		{"SyntheticLocation", mock.SyntheticLocation()},
		{"SyntheticAlertConfigs", mock.SyntheticAlertConfigs()},
		{"AutomationActions", mock.AutomationActions()},
		{"AutomationPolicies", mock.AutomationPolicies()},
		{"HostAgents", mock.HostAgents()},
		{"Users", mock.Users()},
		{"LogAlertConfig", mock.LogAlertConfig()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.result != nil {
				t.Errorf("%s() should return nil, got %v", tt.name, tt.result)
			}
		})
	}
}

// TestMockInstanaAPI_CanBeEmbedded tests that MockInstanaAPI can be embedded
func TestMockInstanaAPI_CanBeEmbedded(t *testing.T) {
	// Create a custom mock that embeds MockInstanaAPI
	type CustomMock struct {
		MockInstanaAPI
	}

	custom := &CustomMock{}

	// Should still return nil for all methods
	if custom.APITokens() != nil {
		t.Error("Embedded mock should return nil")
	}
	if custom.Teams() != nil {
		t.Error("Embedded mock should return nil")
	}
}
