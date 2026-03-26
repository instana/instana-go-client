package client

import (
	"testing"

	"github.com/instana/instana-go-client/mocks"
	"go.uber.org/mock/gomock"
)

// TestNewInstanaAPI tests the creation of a new InstanaAPI client
func TestNewInstanaAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)

	client := NewInstanaAPI(mockRestClient)

	if client == nil {
		t.Fatal("Expected non-nil InstanaAPI client")
	}

	// Verify it's the correct type
	if _, ok := client.(*instanaAPI); !ok {
		t.Fatal("Expected client to be of type *instanaAPI")
	}
}

// TestAPITokensLazyInitialization tests lazy initialization of APITokens client
func TestAPITokensLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	// Initially should be nil
	if client.apiTokens != nil {
		t.Error("Expected apiTokens to be nil before first access")
	}

	// First access should initialize
	tokens := client.APITokens()
	if tokens == nil {
		t.Fatal("Expected non-nil APITokens client")
	}

	// Second access should return same instance
	tokens2 := client.APITokens()
	if tokens != tokens2 {
		t.Error("Expected same instance on second access")
	}
}

// TestAlertingChannelsLazyInitialization tests lazy initialization of AlertingChannels client
func TestAlertingChannelsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	if client.alertingChannels != nil {
		t.Error("Expected alertingChannels to be nil before first access")
	}

	channels := client.AlertingChannels()
	if channels == nil {
		t.Fatal("Expected non-nil AlertingChannels client")
	}

	channels2 := client.AlertingChannels()
	if channels != channels2 {
		t.Error("Expected same instance on second access")
	}
}

// TestAlertingConfigurationsLazyInitialization tests lazy initialization of AlertingConfigurations client
func TestAlertingConfigurationsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	configs := client.AlertingConfigurations()
	if configs == nil {
		t.Fatal("Expected non-nil AlertingConfigurations client")
	}
}

// TestApplicationAlertConfigsLazyInitialization tests lazy initialization of ApplicationAlertConfigs client
func TestApplicationAlertConfigsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	configs := client.ApplicationAlertConfigs()
	if configs == nil {
		t.Fatal("Expected non-nil ApplicationAlertConfigs client")
	}
}

// TestGlobalApplicationAlertConfigsLazyInitialization tests lazy initialization of GlobalApplicationAlertConfigs client
func TestGlobalApplicationAlertConfigsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	configs := client.GlobalApplicationAlertConfigs()
	if configs == nil {
		t.Fatal("Expected non-nil GlobalApplicationAlertConfigs client")
	}
}

// TestApplicationConfigsLazyInitialization tests lazy initialization of ApplicationConfigs client
func TestApplicationConfigsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	configs := client.ApplicationConfigs()
	if configs == nil {
		t.Fatal("Expected non-nil ApplicationConfigs client")
	}
}

// TestAutomationActionsLazyInitialization tests lazy initialization of AutomationActions client
func TestAutomationActionsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	actions := client.AutomationActions()
	if actions == nil {
		t.Fatal("Expected non-nil AutomationActions client")
	}
}

// TestAutomationPoliciesLazyInitialization tests lazy initialization of AutomationPolicies client
func TestAutomationPoliciesLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	policies := client.AutomationPolicies()
	if policies == nil {
		t.Fatal("Expected non-nil AutomationPolicies client")
	}
}

// TestBuiltinEventSpecificationsLazyInitialization tests lazy initialization of BuiltinEventSpecifications client
func TestBuiltinEventSpecificationsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	specs := client.BuiltinEventSpecifications()
	if specs == nil {
		t.Fatal("Expected non-nil BuiltinEventSpecifications client")
	}
}

// TestCustomDashboardsLazyInitialization tests lazy initialization of CustomDashboards client
func TestCustomDashboardsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	dashboards := client.CustomDashboards()
	if dashboards == nil {
		t.Fatal("Expected non-nil CustomDashboards client")
	}
}

// TestCustomEventSpecificationsLazyInitialization tests lazy initialization of CustomEventSpecifications client
func TestCustomEventSpecificationsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	specs := client.CustomEventSpecifications()
	if specs == nil {
		t.Fatal("Expected non-nil CustomEventSpecifications client")
	}
}

// TestGroupsLazyInitialization tests lazy initialization of Groups client
func TestGroupsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	groups := client.Groups()
	if groups == nil {
		t.Fatal("Expected non-nil Groups client")
	}
}

// TestHostAgentsLazyInitialization tests lazy initialization of HostAgents client
func TestHostAgentsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	agents := client.HostAgents()
	if agents == nil {
		t.Fatal("Expected non-nil HostAgents client")
	}
}

// TestInfraAlertConfigsLazyInitialization tests lazy initialization of InfraAlertConfigs client
func TestInfraAlertConfigsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	configs := client.InfraAlertConfigs()
	if configs == nil {
		t.Fatal("Expected non-nil InfraAlertConfigs client")
	}
}

// TestLogAlertConfigsLazyInitialization tests lazy initialization of LogAlertConfigs client
func TestLogAlertConfigsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	configs := client.LogAlertConfigs()
	if configs == nil {
		t.Fatal("Expected non-nil LogAlertConfigs client")
	}
}

// TestMaintenanceWindowConfigsLazyInitialization tests lazy initialization of MaintenanceWindowConfigs client
func TestMaintenanceWindowConfigsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	configs := client.MaintenanceWindowConfigs()
	if configs == nil {
		t.Fatal("Expected non-nil MaintenanceWindowConfigs client")
	}
}

// TestMobileAlertConfigsLazyInitialization tests lazy initialization of MobileAlertConfigs client
func TestMobileAlertConfigsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	configs := client.MobileAlertConfigs()
	if configs == nil {
		t.Fatal("Expected non-nil MobileAlertConfigs client")
	}
}

// TestRolesLazyInitialization tests lazy initialization of Roles client
func TestRolesLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	roles := client.Roles()
	if roles == nil {
		t.Fatal("Expected non-nil Roles client")
	}
}

// TestSliConfigsLazyInitialization tests lazy initialization of SliConfigs client
func TestSliConfigsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	configs := client.SliConfigs()
	if configs == nil {
		t.Fatal("Expected non-nil SliConfigs client")
	}
}

// TestSloAlertConfigsLazyInitialization tests lazy initialization of SloAlertConfigs client
func TestSloAlertConfigsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	configs := client.SloAlertConfigs()
	if configs == nil {
		t.Fatal("Expected non-nil SloAlertConfigs client")
	}
}

// TestSloConfigsLazyInitialization tests lazy initialization of SloConfigs client
func TestSloConfigsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	configs := client.SloConfigs()
	if configs == nil {
		t.Fatal("Expected non-nil SloConfigs client")
	}
}

// TestSloCorrectionConfigsLazyInitialization tests lazy initialization of SloCorrectionConfigs client
func TestSloCorrectionConfigsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	configs := client.SloCorrectionConfigs()
	if configs == nil {
		t.Fatal("Expected non-nil SloCorrectionConfigs client")
	}
}

// TestSyntheticAlertConfigsLazyInitialization tests lazy initialization of SyntheticAlertConfigs client
func TestSyntheticAlertConfigsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	configs := client.SyntheticAlertConfigs()
	if configs == nil {
		t.Fatal("Expected non-nil SyntheticAlertConfigs client")
	}
}

// TestSyntheticLocationsLazyInitialization tests lazy initialization of SyntheticLocations client
func TestSyntheticLocationsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	locations := client.SyntheticLocations()
	if locations == nil {
		t.Fatal("Expected non-nil SyntheticLocations client")
	}
}

// TestSyntheticTestsLazyInitialization tests lazy initialization of SyntheticTests client
func TestSyntheticTestsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	tests := client.SyntheticTests()
	if tests == nil {
		t.Fatal("Expected non-nil SyntheticTests client")
	}
}

// TestTeamsLazyInitialization tests lazy initialization of Teams client
func TestTeamsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	teams := client.Teams()
	if teams == nil {
		t.Fatal("Expected non-nil Teams client")
	}
}

// TestUsersLazyInitialization tests lazy initialization of Users client
func TestUsersLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	users := client.Users()
	if users == nil {
		t.Fatal("Expected non-nil Users client")
	}
}

// TestWebsiteAlertConfigsLazyInitialization tests lazy initialization of WebsiteAlertConfigs client
func TestWebsiteAlertConfigsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	configs := client.WebsiteAlertConfigs()
	if configs == nil {
		t.Fatal("Expected non-nil WebsiteAlertConfigs client")
	}
}

// TestWebsiteMonitoringConfigsLazyInitialization tests lazy initialization of WebsiteMonitoringConfigs client
func TestWebsiteMonitoringConfigsLazyInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	configs := client.WebsiteMonitoringConfigs()
	if configs == nil {
		t.Fatal("Expected non-nil WebsiteMonitoringConfigs client")
	}
}

// TestAllClientsReturnNonNil tests that all client methods return non-nil values
func TestAllClientsReturnNonNil(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient)

	tests := []struct {
		name   string
		getter func() interface{}
	}{
		{"APITokens", func() interface{} { return client.APITokens() }},
		{"AlertingChannels", func() interface{} { return client.AlertingChannels() }},
		{"AlertingConfigurations", func() interface{} { return client.AlertingConfigurations() }},
		{"ApplicationAlertConfigs", func() interface{} { return client.ApplicationAlertConfigs() }},
		{"GlobalApplicationAlertConfigs", func() interface{} { return client.GlobalApplicationAlertConfigs() }},
		{"ApplicationConfigs", func() interface{} { return client.ApplicationConfigs() }},
		{"AutomationActions", func() interface{} { return client.AutomationActions() }},
		{"AutomationPolicies", func() interface{} { return client.AutomationPolicies() }},
		{"BuiltinEventSpecifications", func() interface{} { return client.BuiltinEventSpecifications() }},
		{"CustomDashboards", func() interface{} { return client.CustomDashboards() }},
		{"CustomEventSpecifications", func() interface{} { return client.CustomEventSpecifications() }},
		{"Groups", func() interface{} { return client.Groups() }},
		{"HostAgents", func() interface{} { return client.HostAgents() }},
		{"InfraAlertConfigs", func() interface{} { return client.InfraAlertConfigs() }},
		{"LogAlertConfigs", func() interface{} { return client.LogAlertConfigs() }},
		{"MaintenanceWindowConfigs", func() interface{} { return client.MaintenanceWindowConfigs() }},
		{"MobileAlertConfigs", func() interface{} { return client.MobileAlertConfigs() }},
		{"Roles", func() interface{} { return client.Roles() }},
		{"SliConfigs", func() interface{} { return client.SliConfigs() }},
		{"SloAlertConfigs", func() interface{} { return client.SloAlertConfigs() }},
		{"SloConfigs", func() interface{} { return client.SloConfigs() }},
		{"SloCorrectionConfigs", func() interface{} { return client.SloCorrectionConfigs() }},
		{"SyntheticAlertConfigs", func() interface{} { return client.SyntheticAlertConfigs() }},
		{"SyntheticLocations", func() interface{} { return client.SyntheticLocations() }},
		{"SyntheticTests", func() interface{} { return client.SyntheticTests() }},
		{"Teams", func() interface{} { return client.Teams() }},
		{"Users", func() interface{} { return client.Users() }},
		{"WebsiteAlertConfigs", func() interface{} { return client.WebsiteAlertConfigs() }},
		{"WebsiteMonitoringConfigs", func() interface{} { return client.WebsiteMonitoringConfigs() }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.getter()
			if result == nil {
				t.Errorf("%s returned nil", tt.name)
			}
		})
	}
}

// TestLazyInitializationCaching verifies that lazy initialization caches instances
func TestLazyInitializationCaching(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	client := NewInstanaAPI(mockRestClient).(*instanaAPI)

	// Test a few representative clients
	token1 := client.APITokens()
	token2 := client.APITokens()
	if token1 != token2 {
		t.Error("APITokens should return cached instance")
	}

	channel1 := client.AlertingChannels()
	channel2 := client.AlertingChannels()
	if channel1 != channel2 {
		t.Error("AlertingChannels should return cached instance")
	}

	group1 := client.Groups()
	group2 := client.Groups()
	if group1 != group2 {
		t.Error("Groups should return cached instance")
	}
}

// Made with Bob
