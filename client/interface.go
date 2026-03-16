package client

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

// InstanaAPI is the main interface for interacting with the Instana API.
// It provides access to all API endpoints through dedicated client methods.
// Each method returns a REST resource client for the corresponding API.
type InstanaAPI interface {
	// APITokens returns the API tokens client for managing API authentication tokens
	APITokens() rest.RestResource[*apitoken.APIToken]

	// AlertingChannels returns the alerting channels client for managing notification channels
	AlertingChannels() rest.RestResource[*alertingchannel.AlertingChannel]

	// AlertingConfigurations returns the alerting configurations client for managing alert rules
	AlertingConfigurations() rest.RestResource[*alertingconfig.AlertingConfiguration]

	// ApplicationAlertConfigs returns the application alert configurations client
	ApplicationAlertConfigs() rest.RestResource[*applicationalertconfig.ApplicationAlertConfig]

	// GlobalApplicationAlertConfigs returns the global application alert configurations client
	GlobalApplicationAlertConfigs() rest.RestResource[*applicationalertconfig.ApplicationAlertConfig]

	// ApplicationConfigs returns the application configurations client
	ApplicationConfigs() rest.RestResource[*applicationconfig.ApplicationConfig]

	// AutomationActions returns the automation actions client
	AutomationActions() rest.RestResource[*automationaction.AutomationAction]

	// AutomationPolicies returns the automation policies client
	AutomationPolicies() rest.RestResource[*automationpolicy.AutomationPolicy]

	// BuiltinEventSpecifications returns the read-only builtin event specifications client
	BuiltinEventSpecifications() rest.ReadOnlyRestResource[*builtineventspec.BuiltinEventSpecification]

	// CustomDashboards returns the custom dashboards client
	CustomDashboards() rest.RestResource[*customdashboard.CustomDashboard]

	// CustomEventSpecifications returns the custom event specifications client
	CustomEventSpecifications() rest.RestResource[*customeventspec.CustomEventSpecification]

	// Groups returns the RBAC groups client
	Groups() rest.RestResource[*group.Group]

	// HostAgents returns the read-only host agents client
	HostAgents() rest.ReadOnlyRestResource[*hostagent.HostAgent]

	// InfraAlertConfigs returns the infrastructure alert configurations client
	InfraAlertConfigs() rest.RestResource[*infraalertconfig.InfraAlertConfig]

	// LogAlertConfigs returns the log alert configurations client
	LogAlertConfigs() rest.RestResource[*logalertconfig.LogAlertConfig]

	// MaintenanceWindowConfigs returns the maintenance window configurations client
	MaintenanceWindowConfigs() rest.RestResource[*maintenancewindow.MaintenanceWindow]

	// MobileAlertConfigs returns the mobile app alert configurations client
	MobileAlertConfigs() rest.RestResource[*mobilealertconfig.MobileAlertConfig]

	// Roles returns the RBAC roles client
	Roles() rest.RestResource[*role.Role]

	// SliConfigs returns the SLI configurations client (create only, no update)
	SliConfigs() rest.RestResource[*sliconfig.SliConfig]

	// SloAlertConfigs returns the SLO alert configurations client
	SloAlertConfigs() rest.RestResource[*sloalertconfig.SloAlertConfig]

	// SloConfigs returns the SLO configurations client
	SloConfigs() rest.RestResource[*sloconfig.SloConfig]

	// SloCorrectionConfigs returns the SLO correction configurations client
	SloCorrectionConfigs() rest.RestResource[*slocorrection.SloCorrectionConfig]

	// SyntheticAlertConfigs returns the synthetic alert configurations client
	SyntheticAlertConfigs() rest.RestResource[*syntheticalertconfig.SyntheticAlertConfig]

	// SyntheticLocations returns the read-only synthetic test locations client
	SyntheticLocations() rest.ReadOnlyRestResource[*syntheticlocation.SyntheticLocation]

	// SyntheticTests returns the synthetic tests client
	SyntheticTests() rest.RestResource[*synthetictest.SyntheticTest]

	// Teams returns the RBAC teams client
	Teams() rest.RestResource[*team.Team]

	// Users returns the read-only users client
	Users() rest.ReadOnlyRestResource[*user.User]

	// WebsiteAlertConfigs returns the website alert configurations client
	WebsiteAlertConfigs() rest.RestResource[*websitealertconfig.WebsiteAlertConfig]

	// WebsiteMonitoringConfigs returns the website monitoring configurations client
	WebsiteMonitoringConfigs() rest.RestResource[*websitemonitoring.WebsiteMonitoringConfig]
}

// Made with Bob
