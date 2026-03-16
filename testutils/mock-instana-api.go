package testutils

import (
	"github.com/instana/instana-go-client/api/alertingchannel"
	"github.com/instana/instana-go-client/api/alertingconfig"
	"github.com/instana/instana-go-client/api/apitoken"
	"github.com/instana/instana-go-client/api/applicationalertconfig"
	"github.com/instana/instana-go-client/api/applicationconfig"
	"github.com/instana/instana-go-client/api/automationaction"
	"github.com/instana/instana-go-client/api/automationpolicy"
	"github.com/instana/instana-go-client/api/builtineventspec"
	"github.com/instana/instana-go-client/api/customdashboard"
	"github.com/instana/instana-go-client/api/customeventspec"
	"github.com/instana/instana-go-client/api/group"
	"github.com/instana/instana-go-client/api/hostagent"
	"github.com/instana/instana-go-client/api/infraalertconfig"
	"github.com/instana/instana-go-client/api/logalertconfig"
	"github.com/instana/instana-go-client/api/maintenancewindow"
	"github.com/instana/instana-go-client/api/mobilealertconfig"
	"github.com/instana/instana-go-client/api/role"
	"github.com/instana/instana-go-client/api/sliconfig"
	"github.com/instana/instana-go-client/api/sloalertconfig"
	"github.com/instana/instana-go-client/api/sloconfig"
	"github.com/instana/instana-go-client/api/slocorrection"
	"github.com/instana/instana-go-client/api/syntheticalertconfig"
	"github.com/instana/instana-go-client/api/syntheticlocation"
	"github.com/instana/instana-go-client/api/synthetictest"
	"github.com/instana/instana-go-client/api/team"
	"github.com/instana/instana-go-client/api/user"
	"github.com/instana/instana-go-client/api/websitealertconfig"
	"github.com/instana/instana-go-client/api/websitemonitoring"
	"github.com/instana/instana-go-client/shared/rest"
)

// MockInstanaAPI is a mock implementation of the InstanaAPI interface for testing purposes.
// It returns nil for all methods by default. Tests can override specific methods by embedding
// this struct and providing custom implementations for the methods they need.
type MockInstanaAPI struct{}

// CustomEventSpecifications mock implementation
func (m *MockInstanaAPI) CustomEventSpecifications() rest.RestResource[*customeventspec.CustomEventSpecification] {
	return nil
}

// BuiltinEventSpecifications mock implementation
func (m *MockInstanaAPI) BuiltinEventSpecifications() rest.ReadOnlyRestResource[*builtineventspec.BuiltinEventSpecification] {
	return nil
}

// APITokens mock implementation
func (m *MockInstanaAPI) APITokens() rest.RestResource[*apitoken.APIToken] {
	return nil
}

// ApplicationConfigs mock implementation
func (m *MockInstanaAPI) ApplicationConfigs() rest.RestResource[*applicationconfig.ApplicationConfig] {
	return nil
}

// ApplicationAlertConfigs mock implementation
func (m *MockInstanaAPI) ApplicationAlertConfigs() rest.RestResource[*applicationalertconfig.ApplicationAlertConfig] {
	return nil
}

// GlobalApplicationAlertConfigs mock implementation
func (m *MockInstanaAPI) GlobalApplicationAlertConfigs() rest.RestResource[*applicationalertconfig.ApplicationAlertConfig] {
	return nil
}

// AlertingChannels mock implementation
func (m *MockInstanaAPI) AlertingChannels() rest.RestResource[*alertingchannel.AlertingChannel] {
	return nil
}

// AlertingConfigurations mock implementation
func (m *MockInstanaAPI) AlertingConfigurations() rest.RestResource[*alertingconfig.AlertingConfiguration] {
	return nil
}

// SliConfigs mock implementation
func (m *MockInstanaAPI) SliConfigs() rest.RestResource[*sliconfig.SliConfig] {
	return nil
}

// SloConfigs mock implementation
func (m *MockInstanaAPI) SloConfigs() rest.RestResource[*sloconfig.SloConfig] {
	return nil
}

// SloAlertConfig mock implementation
func (m *MockInstanaAPI) SloAlertConfig() rest.RestResource[*sloalertconfig.SloAlertConfig] {
	return nil
}

// SloCorrectionConfig mock implementation
func (m *MockInstanaAPI) SloCorrectionConfig() rest.RestResource[*slocorrection.SloCorrectionConfig] {
	return nil
}

// WebsiteMonitoringConfig mock implementation
func (m *MockInstanaAPI) WebsiteMonitoringConfig() rest.RestResource[*websitemonitoring.WebsiteMonitoringConfig] {
	return nil
}

// WebsiteAlertConfig mock implementation
func (m *MockInstanaAPI) WebsiteAlertConfig() rest.RestResource[*websitealertconfig.WebsiteAlertConfig] {
	return nil
}

// InfraAlertConfig mock implementation
func (m *MockInstanaAPI) InfraAlertConfig() rest.RestResource[*infraalertconfig.InfraAlertConfig] {
	return nil
}

// MobileAlertConfig mock implementation
func (m *MockInstanaAPI) MobileAlertConfig() rest.RestResource[*mobilealertconfig.MobileAlertConfig] {
	return nil
}

// MaintenanceWindows mock implementation
func (m *MockInstanaAPI) MaintenanceWindows() rest.RestResource[*maintenancewindow.MaintenanceWindow] {
	return nil
}

// Teams mock implementation
func (m *MockInstanaAPI) Teams() rest.RestResource[*team.Team] {
	return nil
}

// Groups mock implementation
func (m *MockInstanaAPI) Groups() rest.RestResource[*group.Group] {
	return nil
}

// Roles mock implementation
func (m *MockInstanaAPI) Roles() rest.RestResource[*role.Role] {
	return nil
}

// CustomDashboards mock implementation
func (m *MockInstanaAPI) CustomDashboards() rest.RestResource[*customdashboard.CustomDashboard] {
	return nil
}

// SyntheticTest mock implementation
func (m *MockInstanaAPI) SyntheticTest() rest.RestResource[*synthetictest.SyntheticTest] {
	return nil
}

// SyntheticLocation mock implementation
func (m *MockInstanaAPI) SyntheticLocation() rest.ReadOnlyRestResource[*syntheticlocation.SyntheticLocation] {
	return nil
}

// SyntheticAlertConfigs mock implementation
func (m *MockInstanaAPI) SyntheticAlertConfigs() rest.RestResource[*syntheticalertconfig.SyntheticAlertConfig] {
	return nil
}

// AutomationActions mock implementation
func (m *MockInstanaAPI) AutomationActions() rest.RestResource[*automationaction.AutomationAction] {
	return nil
}

// AutomationPolicies mock implementation
func (m *MockInstanaAPI) AutomationPolicies() rest.RestResource[*automationpolicy.AutomationPolicy] {
	return nil
}

// HostAgents mock implementation
func (m *MockInstanaAPI) HostAgents() rest.ReadOnlyRestResource[*hostagent.HostAgent] {
	return nil
}

// Users mock implementation
func (m *MockInstanaAPI) Users() rest.ReadOnlyRestResource[*user.User] {
	return nil
}

// LogAlertConfig mock implementation
func (m *MockInstanaAPI) LogAlertConfig() rest.RestResource[*logalertconfig.LogAlertConfig] {
	return nil
}
