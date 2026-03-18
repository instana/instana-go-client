package client

import (
	"github.com/instana/instana-go-client/api"
	"github.com/instana/instana-go-client/shared/rest"
)

// InstanaAPI is the main interface for interacting with the Instana API.
// It provides access to all API endpoints through dedicated client methods.
// Each method returns a REST resource client for the corresponding API.
type InstanaAPI interface {
	// APITokens returns the API tokens client for managing API authentication tokens
	APITokens() rest.RestResource[*api.APIToken]

	// AlertingChannels returns the alerting channels client for managing notification channels
	AlertingChannels() rest.RestResource[*api.AlertingChannel]

	// AlertingConfigurations returns the alerting configurations client for managing alert rules
	AlertingConfigurations() rest.RestResource[*api.AlertingConfiguration]

	// ApplicationAlertConfigs returns the application alert configurations client
	ApplicationAlertConfigs() rest.RestResource[*api.ApplicationAlertConfig]

	// GlobalApplicationAlertConfigs returns the global application alert configurations client
	GlobalApplicationAlertConfigs() rest.RestResource[*api.ApplicationAlertConfig]

	// ApplicationConfigs returns the application configurations client
	ApplicationConfigs() rest.RestResource[*api.ApplicationConfig]

	// AutomationActions returns the automation actions client
	AutomationActions() rest.RestResource[*api.AutomationAction]

	// AutomationPolicies returns the automation policies client
	AutomationPolicies() rest.RestResource[*api.AutomationPolicy]

	// BuiltinEventSpecifications returns the read-only builtin event specifications client
	BuiltinEventSpecifications() rest.ReadOnlyRestResource[*api.BuiltinEventSpecification]

	// CustomDashboards returns the custom dashboards client
	CustomDashboards() rest.RestResource[*api.CustomDashboard]

	// CustomEventSpecifications returns the custom event specifications client
	CustomEventSpecifications() rest.RestResource[*api.CustomEventSpecification]

	// Groups returns the RBAC groups client
	Groups() rest.RestResource[*api.Group]

	// HostAgents returns the read-only host agents client
	HostAgents() rest.ReadOnlyRestResource[*api.HostAgent]

	// InfraAlertConfigs returns the infrastructure alert configurations client
	InfraAlertConfigs() rest.RestResource[*api.InfraAlertConfig]

	// LogAlertConfigs returns the log alert configurations client
	LogAlertConfigs() rest.RestResource[*api.LogAlertConfig]

	// MaintenanceWindowConfigs returns the maintenance window configurations client
	MaintenanceWindowConfigs() rest.RestResource[*api.MaintenanceWindow]

	// MobileAlertConfigs returns the mobile app alert configurations client
	MobileAlertConfigs() rest.RestResource[*api.MobileAlertConfig]

	// Roles returns the RBAC roles client
	Roles() rest.RestResource[*api.Role]

	// SliConfigs returns the SLI configurations client (create only, no update)
	SliConfigs() rest.RestResource[*api.SliConfig]

	// SloAlertConfigs returns the SLO alert configurations client
	SloAlertConfigs() rest.RestResource[*api.SloAlertConfig]

	// SloConfigs returns the SLO configurations client
	SloConfigs() rest.RestResource[*api.SloConfig]

	// SloCorrectionConfigs returns the SLO correction configurations client
	SloCorrectionConfigs() rest.RestResource[*api.SloCorrectionConfig]

	// SyntheticAlertConfigs returns the synthetic alert configurations client
	SyntheticAlertConfigs() rest.RestResource[*api.SyntheticAlertConfig]

	// SyntheticLocations returns the read-only synthetic test locations client
	SyntheticLocations() rest.ReadOnlyRestResource[*api.SyntheticLocation]

	// SyntheticTests returns the synthetic tests client
	SyntheticTests() rest.RestResource[*api.SyntheticTest]

	// Teams returns the RBAC teams client
	Teams() rest.RestResource[*api.Team]

	// Users returns the read-only users client
	Users() rest.ReadOnlyRestResource[*api.User]

	// WebsiteAlertConfigs returns the website alert configurations client
	WebsiteAlertConfigs() rest.RestResource[*api.WebsiteAlertConfig]

	// WebsiteMonitoringConfigs returns the website monitoring configurations client
	WebsiteMonitoringConfigs() rest.RestResource[*api.WebsiteMonitoringConfig]
}
