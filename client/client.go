package client

import (
	api "github.com/instana/instana-go-client/api"
	"github.com/instana/instana-go-client/shared/rest"
)

// instanaAPI is the concrete implementation of the InstanaAPI interface.
// It uses lazy initialization for all API clients to improve performance.
type instanaAPI struct {
	restClient rest.RestClient

	// Lazy-initialized API clients
	apiTokens                     rest.RestResource[*api.APIToken]
	alertingChannels              rest.RestResource[*api.AlertingChannel]
	alertingConfigurations        rest.RestResource[*api.AlertingConfiguration]
	applicationAlertConfigs       rest.RestResource[*api.ApplicationAlertConfig]
	globalApplicationAlertConfigs rest.RestResource[*api.ApplicationAlertConfig]
	applicationConfigs            rest.RestResource[*api.ApplicationConfig]
	automationActions             rest.RestResource[*api.AutomationAction]
	automationPolicies            rest.RestResource[*api.AutomationPolicy]
	builtinEventSpecifications    rest.ReadOnlyRestResource[*api.BuiltinEventSpecification]
	customDashboards              rest.RestResource[*api.CustomDashboard]
	customEventSpecifications     rest.RestResource[*api.CustomEventSpecification]
	groups                        rest.RestResource[*api.Group]
	hostAgents                    rest.ReadOnlyRestResource[*api.HostAgent]
	infraAlertConfigs             rest.RestResource[*api.InfraAlertConfig]
	logAlertConfigs               rest.RestResource[*api.LogAlertConfig]
	maintenanceWindows            rest.RestResource[*api.MaintenanceWindow]
	mobileAlertConfigs            rest.RestResource[*api.MobileAlertConfig]
	roles                         rest.RestResource[*api.Role]
	sliConfigs                    rest.RestResource[*api.SliConfig]
	sloAlertConfigs               rest.RestResource[*api.SloAlertConfig]
	sloConfigs                    rest.RestResource[*api.SloConfig]
	sloCorrections                rest.RestResource[*api.SloCorrectionConfig]
	syntheticAlertConfigs         rest.RestResource[*api.SyntheticAlertConfig]
	syntheticLocations            rest.ReadOnlyRestResource[*api.SyntheticLocation]
	syntheticTests                rest.RestResource[*api.SyntheticTest]
	teams                         rest.RestResource[*api.Team]
	users                         rest.ReadOnlyRestResource[*api.User]
	websiteAlertConfigs           rest.RestResource[*api.WebsiteAlertConfig]
	websiteMonitoringConfigs      rest.RestResource[*api.WebsiteMonitoringConfig]
}

// NewInstanaAPI creates a new Instana API client with the provided REST client.
// All resource clients are lazily initialized on first access.
func NewInstanaAPI(restClient rest.RestClient) InstanaAPI {
	return &instanaAPI{
		restClient: restClient,
	}
}

// APITokens returns the API tokens client (lazy initialization)
func (c *instanaAPI) APITokens() rest.RestResource[*api.APIToken] {
	if c.apiTokens == nil {
		c.apiTokens = NewRestResource[*api.APIToken](
			c.restClient,
			api.APITokensResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			rest.NewGenericUnmarshaller[*api.APIToken](),
		)
	}
	return c.apiTokens
}

// AlertingChannels returns the alerting channels client (lazy initialization)
func (c *instanaAPI) AlertingChannels() rest.RestResource[*api.AlertingChannel] {
	if c.alertingChannels == nil {
		c.alertingChannels = NewRestResource[*api.AlertingChannel](
			c.restClient,
			api.AlertingchannelResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePUT,
			rest.NewGenericUnmarshaller[*api.AlertingChannel](),
		)
	}
	return c.alertingChannels
}

// AlertingConfigurations returns the alerting configurations client (lazy initialization)
func (c *instanaAPI) AlertingConfigurations() rest.RestResource[*api.AlertingConfiguration] {
	if c.alertingConfigurations == nil {
		c.alertingConfigurations = NewRestResource[*api.AlertingConfiguration](
			c.restClient,
			api.AlertsResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePUT,
			rest.NewCustomPayloadFieldsUnmarshallerAdapter(rest.NewGenericUnmarshaller[*api.AlertingConfiguration]()),
		)
	}
	return c.alertingConfigurations
}

// ApplicationAlertConfigs returns the application alert configurations client (lazy initialization)
func (c *instanaAPI) ApplicationAlertConfigs() rest.RestResource[*api.ApplicationAlertConfig] {
	if c.applicationAlertConfigs == nil {
		c.applicationAlertConfigs = NewRestResource[*api.ApplicationAlertConfig](
			c.restClient,
			api.ApplicationAlertConfigsResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePOST,
			rest.NewCustomPayloadFieldsUnmarshallerAdapter(rest.NewGenericUnmarshaller[*api.ApplicationAlertConfig]()),
		)
	}
	return c.applicationAlertConfigs
}

// GlobalApplicationAlertConfigs returns the global application alert configurations client (lazy initialization)
func (c *instanaAPI) GlobalApplicationAlertConfigs() rest.RestResource[*api.ApplicationAlertConfig] {
	if c.globalApplicationAlertConfigs == nil {
		c.globalApplicationAlertConfigs = NewRestResource[*api.ApplicationAlertConfig](
			c.restClient,
			api.GlobalApplicationAlertConfigsResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePOST,
			rest.NewCustomPayloadFieldsUnmarshallerAdapter(rest.NewGenericUnmarshaller[*api.ApplicationAlertConfig]()),
		)
	}
	return c.globalApplicationAlertConfigs
}

// ApplicationConfigs returns the application configurations client (lazy initialization)
func (c *instanaAPI) ApplicationConfigs() rest.RestResource[*api.ApplicationConfig] {
	if c.applicationConfigs == nil {
		c.applicationConfigs = NewRestResource[*api.ApplicationConfig](
			c.restClient,
			api.ApplicationConfigsResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			rest.NewGenericUnmarshaller[*api.ApplicationConfig](),
		)
	}
	return c.applicationConfigs
}

// AutomationActions returns the automation actions client (lazy initialization)
func (c *instanaAPI) AutomationActions() rest.RestResource[*api.AutomationAction] {
	if c.automationActions == nil {
		c.automationActions = NewRestResource[*api.AutomationAction](
			c.restClient,
			api.AutomationActionResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			rest.NewGenericUnmarshaller[*api.AutomationAction](),
		)
	}
	return c.automationActions
}

// AutomationPolicies returns the automation policies client (lazy initialization)
func (c *instanaAPI) AutomationPolicies() rest.RestResource[*api.AutomationPolicy] {
	if c.automationPolicies == nil {
		c.automationPolicies = NewRestResource[*api.AutomationPolicy](
			c.restClient,
			api.AutomationPolicyResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			rest.NewGenericUnmarshaller[*api.AutomationPolicy](),
		)
	}
	return c.automationPolicies
}

// BuiltinEventSpecifications returns the builtin event specifications client (lazy initialization)
func (c *instanaAPI) BuiltinEventSpecifications() rest.ReadOnlyRestResource[*api.BuiltinEventSpecification] {
	if c.builtinEventSpecifications == nil {
		c.builtinEventSpecifications = NewReadOnlyRestResource[*api.BuiltinEventSpecification](
			c.restClient,
			api.BuiltinEventSpecificationResourcePath,
			rest.NewGenericUnmarshaller[*api.BuiltinEventSpecification](),
		)
	}
	return c.builtinEventSpecifications
}

// CustomDashboards returns the custom dashboards client (lazy initialization)
func (c *instanaAPI) CustomDashboards() rest.RestResource[*api.CustomDashboard] {
	if c.customDashboards == nil {
		c.customDashboards = NewRestResource[*api.CustomDashboard](
			c.restClient,
			api.CustomDashboardsResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			rest.NewGenericUnmarshaller[*api.CustomDashboard](),
		)
	}
	return c.customDashboards
}

// CustomEventSpecifications returns the custom event specifications client (lazy initialization)
func (c *instanaAPI) CustomEventSpecifications() rest.RestResource[*api.CustomEventSpecification] {
	if c.customEventSpecifications == nil {
		c.customEventSpecifications = NewRestResource[*api.CustomEventSpecification](
			c.restClient,
			api.CustomeventspecResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePUT,
			rest.NewGenericUnmarshaller[*api.CustomEventSpecification](),
		)
	}
	return c.customEventSpecifications
}

// Groups returns the groups client (lazy initialization)
func (c *instanaAPI) Groups() rest.RestResource[*api.Group] {
	if c.groups == nil {
		c.groups = NewRestResource[*api.Group](
			c.restClient,
			api.GroupResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			rest.NewGenericUnmarshaller[*api.Group](),
		)
	}
	return c.groups
}

// HostAgents returns the host agents client (lazy initialization)
func (c *instanaAPI) HostAgents() rest.ReadOnlyRestResource[*api.HostAgent] {
	if c.hostAgents == nil {
		c.hostAgents = NewReadOnlyRestResource[*api.HostAgent](
			c.restClient,
			api.HostAgentResourcePath,
			api.NewHostAgentJSONUnmarshaller(&api.HostAgent{}),
		)
	}
	return c.hostAgents
}

// InfraAlertConfigs returns the infrastructure alert configurations client (lazy initialization)
func (c *instanaAPI) InfraAlertConfigs() rest.RestResource[*api.InfraAlertConfig] {
	if c.infraAlertConfigs == nil {
		c.infraAlertConfigs = NewRestResource[*api.InfraAlertConfig](
			c.restClient,
			api.ResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePOST,
			rest.NewCustomPayloadFieldsUnmarshallerAdapter(rest.NewGenericUnmarshaller[*api.InfraAlertConfig]()),
		)
	}
	return c.infraAlertConfigs
}

// LogAlertConfigs returns the log alert configurations client (lazy initialization)
func (c *instanaAPI) LogAlertConfigs() rest.RestResource[*api.LogAlertConfig] {
	if c.logAlertConfigs == nil {
		c.logAlertConfigs = NewRestResource[*api.LogAlertConfig](
			c.restClient,
			api.LogAlertConfigResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePOST,
			rest.NewCustomPayloadFieldsUnmarshallerAdapter(rest.NewGenericUnmarshaller[*api.LogAlertConfig]()),
		)
	}
	return c.logAlertConfigs
}

// MaintenanceWindowConfigs returns the maintenance window configurations client (lazy initialization)
func (c *instanaAPI) MaintenanceWindowConfigs() rest.RestResource[*api.MaintenanceWindow] {
	if c.maintenanceWindows == nil {
		c.maintenanceWindows = NewRestResource[*api.MaintenanceWindow](
			c.restClient,
			api.MaintenanceWindowConfigResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePUT,
			rest.NewGenericUnmarshaller[*api.MaintenanceWindow](),
		)
	}
	return c.maintenanceWindows
}

// MobileAlertConfigs returns the mobile alert configurations client (lazy initialization)
func (c *instanaAPI) MobileAlertConfigs() rest.RestResource[*api.MobileAlertConfig] {
	if c.mobileAlertConfigs == nil {
		c.mobileAlertConfigs = NewRestResource[*api.MobileAlertConfig](
			c.restClient,
			api.MobileAlertConfigResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePOST,
			rest.NewCustomPayloadFieldsUnmarshallerAdapter(rest.NewGenericUnmarshaller[*api.MobileAlertConfig]()),
		)
	}
	return c.mobileAlertConfigs
}

// Roles returns the roles client (lazy initialization)
func (c *instanaAPI) Roles() rest.RestResource[*api.Role] {
	if c.roles == nil {
		c.roles = NewRestResource[*api.Role](
			c.restClient,
			api.ResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			rest.NewGenericUnmarshaller[*api.Role](),
		)
	}
	return c.roles
}

// SliConfigs returns the SLI configurations client (lazy initialization)
func (c *instanaAPI) SliConfigs() rest.RestResource[*api.SliConfig] {
	if c.sliConfigs == nil {
		c.sliConfigs = NewRestResource[*api.SliConfig](
			c.restClient,
			api.SliConfigResourcePath,
			rest.DefaultRestResourceModeCreatePOSTAndUpdateNotSupported,
			rest.NewGenericUnmarshaller[*api.SliConfig](),
		)
	}
	return c.sliConfigs
}

// SloAlertConfigs returns the SLO alert configurations client (lazy initialization)
func (c *instanaAPI) SloAlertConfigs() rest.RestResource[*api.SloAlertConfig] {
	if c.sloAlertConfigs == nil {
		c.sloAlertConfigs = NewRestResource[*api.SloAlertConfig](
			c.restClient,
			api.SloAlertConfigResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePOST,
			rest.NewCustomPayloadFieldsUnmarshallerAdapter(rest.NewGenericUnmarshaller[*api.SloAlertConfig]()),
		)
	}
	return c.sloAlertConfigs
}

// SloConfigs returns the SLO configurations client (lazy initialization)
func (c *instanaAPI) SloConfigs() rest.RestResource[*api.SloConfig] {
	if c.sloConfigs == nil {
		c.sloConfigs = NewRestResource[*api.SloConfig](
			c.restClient,
			api.SloConfigResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			api.NewSloConfigJSONUnmarshaller[*api.SloConfig](&api.SloConfig{}),
		)
	}
	return c.sloConfigs
}

// SloCorrectionConfigs returns the SLO correction configurations client (lazy initialization)
func (c *instanaAPI) SloCorrectionConfigs() rest.RestResource[*api.SloCorrectionConfig] {
	if c.sloCorrections == nil {
		c.sloCorrections = NewRestResource[*api.SloCorrectionConfig](
			c.restClient,
			api.SloCorrectionConfigResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			api.NewSloCorrectionConfigJSONUnmarshaller[*api.SloCorrectionConfig](&api.SloCorrectionConfig{}),
		)
	}
	return c.sloCorrections
}

// SyntheticAlertConfigs returns the synthetic alert configurations client (lazy initialization)
func (c *instanaAPI) SyntheticAlertConfigs() rest.RestResource[*api.SyntheticAlertConfig] {
	if c.syntheticAlertConfigs == nil {
		c.syntheticAlertConfigs = NewRestResource[*api.SyntheticAlertConfig](
			c.restClient,
			api.SyntheticAlertConfigsResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePOST,
			rest.NewCustomPayloadFieldsUnmarshallerAdapter(rest.NewGenericUnmarshaller[*api.SyntheticAlertConfig]()),
		)
	}
	return c.syntheticAlertConfigs
}

// SyntheticLocations returns the synthetic locations client (lazy initialization)
func (c *instanaAPI) SyntheticLocations() rest.ReadOnlyRestResource[*api.SyntheticLocation] {
	if c.syntheticLocations == nil {
		c.syntheticLocations = NewReadOnlyRestResource[*api.SyntheticLocation](
			c.restClient,
			api.ResourcePath,
			rest.NewGenericUnmarshaller[*api.SyntheticLocation](),
		)
	}
	return c.syntheticLocations
}

// SyntheticTests returns the synthetic tests client (lazy initialization)
func (c *instanaAPI) SyntheticTests() rest.RestResource[*api.SyntheticTest] {
	if c.syntheticTests == nil {
		c.syntheticTests = NewRestResource[*api.SyntheticTest](
			c.restClient,
			api.ResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			rest.NewGenericUnmarshaller[*api.SyntheticTest](),
		)
	}
	return c.syntheticTests
}

// Teams returns the teams client (lazy initialization)
func (c *instanaAPI) Teams() rest.RestResource[*api.Team] {
	if c.teams == nil {
		c.teams = NewRestResource[*api.Team](
			c.restClient,
			api.ResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			rest.NewGenericUnmarshaller[*api.Team](),
		)
	}
	return c.teams
}

// Users returns the users client (lazy initialization)
func (c *instanaAPI) Users() rest.ReadOnlyRestResource[*api.User] {
	if c.users == nil {
		c.users = NewReadOnlyRestResource[*api.User](
			c.restClient,
			api.ResourcePath,
			rest.NewGenericUnmarshaller[*api.User](),
		)
	}
	return c.users
}

// WebsiteAlertConfigs returns the website alert configurations client (lazy initialization)
func (c *instanaAPI) WebsiteAlertConfigs() rest.RestResource[*api.WebsiteAlertConfig] {
	if c.websiteAlertConfigs == nil {
		c.websiteAlertConfigs = NewRestResource[*api.WebsiteAlertConfig](
			c.restClient,
			api.ResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePOST,
			rest.NewCustomPayloadFieldsUnmarshallerAdapter(rest.NewGenericUnmarshaller[*api.WebsiteAlertConfig]()),
		)
	}
	return c.websiteAlertConfigs
}

// WebsiteMonitoringConfigs returns the website monitoring configurations client (lazy initialization)
func (c *instanaAPI) WebsiteMonitoringConfigs() rest.RestResource[*api.WebsiteMonitoringConfig] {
	if c.websiteMonitoringConfigs == nil {
		c.websiteMonitoringConfigs = api.NewWebsiteMonitoringConfigRestResource(
			rest.NewGenericUnmarshaller[*api.WebsiteMonitoringConfig](),
			c.restClient,
		)
	}
	return c.websiteMonitoringConfigs
}
