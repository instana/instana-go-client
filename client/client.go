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

// instanaAPI is the concrete implementation of the InstanaAPI interface.
// It uses lazy initialization for all API clients to improve performance.
type instanaAPI struct {
	restClient rest.RestClient

	// Lazy-initialized API clients
	apiTokens                     rest.RestResource[*apitoken.APIToken]
	alertingChannels              rest.RestResource[*alertingchannel.AlertingChannel]
	alertingConfigurations        rest.RestResource[*alertingconfig.AlertingConfiguration]
	applicationAlertConfigs       rest.RestResource[*applicationalertconfig.ApplicationAlertConfig]
	globalApplicationAlertConfigs rest.RestResource[*applicationalertconfig.ApplicationAlertConfig]
	applicationConfigs            rest.RestResource[*applicationconfig.ApplicationConfig]
	automationActions             rest.RestResource[*automationaction.AutomationAction]
	automationPolicies            rest.RestResource[*automationpolicy.AutomationPolicy]
	builtinEventSpecifications    rest.ReadOnlyRestResource[*builtineventspec.BuiltinEventSpecification]
	customDashboards              rest.RestResource[*customdashboard.CustomDashboard]
	customEventSpecifications     rest.RestResource[*customeventspec.CustomEventSpecification]
	groups                        rest.RestResource[*group.Group]
	hostAgents                    rest.ReadOnlyRestResource[*hostagent.HostAgent]
	infraAlertConfigs             rest.RestResource[*infraalertconfig.InfraAlertConfig]
	logAlertConfigs               rest.RestResource[*logalertconfig.LogAlertConfig]
	maintenanceWindows            rest.RestResource[*maintenancewindow.MaintenanceWindow]
	mobileAlertConfigs            rest.RestResource[*mobilealertconfig.MobileAlertConfig]
	roles                         rest.RestResource[*role.Role]
	sliConfigs                    rest.RestResource[*sliconfig.SliConfig]
	sloAlertConfigs               rest.RestResource[*sloalertconfig.SloAlertConfig]
	sloConfigs                    rest.RestResource[*sloconfig.SloConfig]
	sloCorrections                rest.RestResource[*slocorrection.SloCorrectionConfig]
	syntheticAlertConfigs         rest.RestResource[*syntheticalertconfig.SyntheticAlertConfig]
	syntheticLocations            rest.ReadOnlyRestResource[*syntheticlocation.SyntheticLocation]
	syntheticTests                rest.RestResource[*synthetictest.SyntheticTest]
	teams                         rest.RestResource[*team.Team]
	users                         rest.ReadOnlyRestResource[*user.User]
	websiteAlertConfigs           rest.RestResource[*websitealertconfig.WebsiteAlertConfig]
	websiteMonitoringConfigs      rest.RestResource[*websitemonitoring.WebsiteMonitoringConfig]
}

// NewInstanaAPI creates a new Instana API client with the provided REST client.
// All resource clients are lazily initialized on first access.
func NewInstanaAPI(restClient rest.RestClient) InstanaAPI {
	return &instanaAPI{
		restClient: restClient,
	}
}

// APITokens returns the API tokens client (lazy initialization)
func (api *instanaAPI) APITokens() rest.RestResource[*apitoken.APIToken] {
	if api.apiTokens == nil {
		api.apiTokens = NewRestResource[*apitoken.APIToken](
			api.restClient,
			apitoken.APITokensResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			rest.NewGenericUnmarshaller[*apitoken.APIToken](),
		)
	}
	return api.apiTokens
}

// AlertingChannels returns the alerting channels client (lazy initialization)
func (api *instanaAPI) AlertingChannels() rest.RestResource[*alertingchannel.AlertingChannel] {
	if api.alertingChannels == nil {
		api.alertingChannels = NewRestResource[*alertingchannel.AlertingChannel](
			api.restClient,
			alertingchannel.AlertingchannelResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePUT,
			rest.NewGenericUnmarshaller[*alertingchannel.AlertingChannel](),
		)
	}
	return api.alertingChannels
}

// AlertingConfigurations returns the alerting configurations client (lazy initialization)
func (api *instanaAPI) AlertingConfigurations() rest.RestResource[*alertingconfig.AlertingConfiguration] {
	if api.alertingConfigurations == nil {
		api.alertingConfigurations = NewRestResource[*alertingconfig.AlertingConfiguration](
			api.restClient,
			alertingconfig.AlertsResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePUT,
			rest.NewGenericUnmarshaller[*alertingconfig.AlertingConfiguration](),
		)
	}
	return api.alertingConfigurations
}

// ApplicationAlertConfigs returns the application alert configurations client (lazy initialization)
func (api *instanaAPI) ApplicationAlertConfigs() rest.RestResource[*applicationalertconfig.ApplicationAlertConfig] {
	if api.applicationAlertConfigs == nil {
		api.applicationAlertConfigs = NewRestResource[*applicationalertconfig.ApplicationAlertConfig](
			api.restClient,
			applicationalertconfig.ApplicationAlertConfigsResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePOST,
			rest.NewGenericUnmarshaller[*applicationalertconfig.ApplicationAlertConfig](),
		)
	}
	return api.applicationAlertConfigs
}

// GlobalApplicationAlertConfigs returns the global application alert configurations client (lazy initialization)
func (api *instanaAPI) GlobalApplicationAlertConfigs() rest.RestResource[*applicationalertconfig.ApplicationAlertConfig] {
	if api.globalApplicationAlertConfigs == nil {
		api.globalApplicationAlertConfigs = NewRestResource[*applicationalertconfig.ApplicationAlertConfig](
			api.restClient,
			applicationalertconfig.GlobalApplicationAlertConfigsResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePOST,
			rest.NewGenericUnmarshaller[*applicationalertconfig.ApplicationAlertConfig](),
		)
	}
	return api.globalApplicationAlertConfigs
}

// ApplicationConfigs returns the application configurations client (lazy initialization)
func (api *instanaAPI) ApplicationConfigs() rest.RestResource[*applicationconfig.ApplicationConfig] {
	if api.applicationConfigs == nil {
		api.applicationConfigs = NewRestResource[*applicationconfig.ApplicationConfig](
			api.restClient,
			applicationconfig.ApplicationConfigsResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			rest.NewGenericUnmarshaller[*applicationconfig.ApplicationConfig](),
		)
	}
	return api.applicationConfigs
}

// AutomationActions returns the automation actions client (lazy initialization)
func (api *instanaAPI) AutomationActions() rest.RestResource[*automationaction.AutomationAction] {
	if api.automationActions == nil {
		api.automationActions = NewRestResource[*automationaction.AutomationAction](
			api.restClient,
			automationaction.AutomationActionResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			rest.NewGenericUnmarshaller[*automationaction.AutomationAction](),
		)
	}
	return api.automationActions
}

// AutomationPolicies returns the automation policies client (lazy initialization)
func (api *instanaAPI) AutomationPolicies() rest.RestResource[*automationpolicy.AutomationPolicy] {
	if api.automationPolicies == nil {
		api.automationPolicies = NewRestResource[*automationpolicy.AutomationPolicy](
			api.restClient,
			automationpolicy.AutomationPolicyResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			rest.NewGenericUnmarshaller[*automationpolicy.AutomationPolicy](),
		)
	}
	return api.automationPolicies
}

// BuiltinEventSpecifications returns the builtin event specifications client (lazy initialization)
func (api *instanaAPI) BuiltinEventSpecifications() rest.ReadOnlyRestResource[*builtineventspec.BuiltinEventSpecification] {
	if api.builtinEventSpecifications == nil {
		api.builtinEventSpecifications = NewReadOnlyRestResource[*builtineventspec.BuiltinEventSpecification](
			api.restClient,
			builtineventspec.BuiltinEventSpecificationResourcePath,
			rest.NewGenericUnmarshaller[*builtineventspec.BuiltinEventSpecification](),
		)
	}
	return api.builtinEventSpecifications
}

// CustomDashboards returns the custom dashboards client (lazy initialization)
func (api *instanaAPI) CustomDashboards() rest.RestResource[*customdashboard.CustomDashboard] {
	if api.customDashboards == nil {
		api.customDashboards = NewRestResource[*customdashboard.CustomDashboard](
			api.restClient,
			customdashboard.CustomDashboardsResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			rest.NewGenericUnmarshaller[*customdashboard.CustomDashboard](),
		)
	}
	return api.customDashboards
}

// CustomEventSpecifications returns the custom event specifications client (lazy initialization)
func (api *instanaAPI) CustomEventSpecifications() rest.RestResource[*customeventspec.CustomEventSpecification] {
	if api.customEventSpecifications == nil {
		api.customEventSpecifications = NewRestResource[*customeventspec.CustomEventSpecification](
			api.restClient,
			customeventspec.CustomeventspecResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePUT,
			rest.NewGenericUnmarshaller[*customeventspec.CustomEventSpecification](),
		)
	}
	return api.customEventSpecifications
}

// Groups returns the groups client (lazy initialization)
func (api *instanaAPI) Groups() rest.RestResource[*group.Group] {
	if api.groups == nil {
		api.groups = NewRestResource[*group.Group](
			api.restClient,
			group.GroupResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			rest.NewGenericUnmarshaller[*group.Group](),
		)
	}
	return api.groups
}

// HostAgents returns the host agents client (lazy initialization)
func (api *instanaAPI) HostAgents() rest.ReadOnlyRestResource[*hostagent.HostAgent] {
	if api.hostAgents == nil {
		api.hostAgents = NewReadOnlyRestResource[*hostagent.HostAgent](
			api.restClient,
			hostagent.HostAgentResourcePath,
			hostagent.NewHostAgentJSONUnmarshaller(&hostagent.HostAgent{}),
		)
	}
	return api.hostAgents
}

// InfraAlertConfigs returns the infrastructure alert configurations client (lazy initialization)
func (api *instanaAPI) InfraAlertConfigs() rest.RestResource[*infraalertconfig.InfraAlertConfig] {
	if api.infraAlertConfigs == nil {
		api.infraAlertConfigs = NewRestResource[*infraalertconfig.InfraAlertConfig](
			api.restClient,
			infraalertconfig.ResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePOST,
			rest.NewGenericUnmarshaller[*infraalertconfig.InfraAlertConfig](),
		)
	}
	return api.infraAlertConfigs
}

// LogAlertConfigs returns the log alert configurations client (lazy initialization)
func (api *instanaAPI) LogAlertConfigs() rest.RestResource[*logalertconfig.LogAlertConfig] {
	if api.logAlertConfigs == nil {
		api.logAlertConfigs = NewRestResource[*logalertconfig.LogAlertConfig](
			api.restClient,
			logalertconfig.LogAlertConfigResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePOST,
			rest.NewGenericUnmarshaller[*logalertconfig.LogAlertConfig](),
		)
	}
	return api.logAlertConfigs
}

// MaintenanceWindowConfigs returns the maintenance window configurations client (lazy initialization)
func (api *instanaAPI) MaintenanceWindowConfigs() rest.RestResource[*maintenancewindow.MaintenanceWindow] {
	if api.maintenanceWindows == nil {
		api.maintenanceWindows = NewRestResource[*maintenancewindow.MaintenanceWindow](
			api.restClient,
			maintenancewindow.MaintenanceWindowConfigResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePUT,
			rest.NewGenericUnmarshaller[*maintenancewindow.MaintenanceWindow](),
		)
	}
	return api.maintenanceWindows
}

// MobileAlertConfigs returns the mobile alert configurations client (lazy initialization)
func (api *instanaAPI) MobileAlertConfigs() rest.RestResource[*mobilealertconfig.MobileAlertConfig] {
	if api.mobileAlertConfigs == nil {
		api.mobileAlertConfigs = NewRestResource[*mobilealertconfig.MobileAlertConfig](
			api.restClient,
			mobilealertconfig.MobileAlertConfigResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePOST,
			rest.NewGenericUnmarshaller[*mobilealertconfig.MobileAlertConfig](),
		)
	}
	return api.mobileAlertConfigs
}

// Roles returns the roles client (lazy initialization)
func (api *instanaAPI) Roles() rest.RestResource[*role.Role] {
	if api.roles == nil {
		api.roles = NewRestResource[*role.Role](
			api.restClient,
			role.ResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			rest.NewGenericUnmarshaller[*role.Role](),
		)
	}
	return api.roles
}

// SliConfigs returns the SLI configurations client (lazy initialization)
func (api *instanaAPI) SliConfigs() rest.RestResource[*sliconfig.SliConfig] {
	if api.sliConfigs == nil {
		api.sliConfigs = NewRestResource[*sliconfig.SliConfig](
			api.restClient,
			sliconfig.SliConfigResourcePath,
			rest.DefaultRestResourceModeCreatePOSTAndUpdateNotSupported,
			rest.NewGenericUnmarshaller[*sliconfig.SliConfig](),
		)
	}
	return api.sliConfigs
}

// SloAlertConfigs returns the SLO alert configurations client (lazy initialization)
func (api *instanaAPI) SloAlertConfigs() rest.RestResource[*sloalertconfig.SloAlertConfig] {
	if api.sloAlertConfigs == nil {
		api.sloAlertConfigs = NewRestResource[*sloalertconfig.SloAlertConfig](
			api.restClient,
			sloalertconfig.SloAlertConfigResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePOST,
			rest.NewGenericUnmarshaller[*sloalertconfig.SloAlertConfig](),
		)
	}
	return api.sloAlertConfigs
}

// SloConfigs returns the SLO configurations client (lazy initialization)
func (api *instanaAPI) SloConfigs() rest.RestResource[*sloconfig.SloConfig] {
	if api.sloConfigs == nil {
		api.sloConfigs = NewRestResource[*sloconfig.SloConfig](
			api.restClient,
			sloconfig.SloConfigResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			rest.NewGenericUnmarshaller[*sloconfig.SloConfig](),
		)
	}
	return api.sloConfigs
}

// SloCorrectionConfigs returns the SLO correction configurations client (lazy initialization)
func (api *instanaAPI) SloCorrectionConfigs() rest.RestResource[*slocorrection.SloCorrectionConfig] {
	if api.sloCorrections == nil {
		api.sloCorrections = NewRestResource[*slocorrection.SloCorrectionConfig](
			api.restClient,
			slocorrection.SloCorrectionConfigResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			rest.NewGenericUnmarshaller[*slocorrection.SloCorrectionConfig](),
		)
	}
	return api.sloCorrections
}

// SyntheticAlertConfigs returns the synthetic alert configurations client (lazy initialization)
func (api *instanaAPI) SyntheticAlertConfigs() rest.RestResource[*syntheticalertconfig.SyntheticAlertConfig] {
	if api.syntheticAlertConfigs == nil {
		api.syntheticAlertConfigs = NewRestResource[*syntheticalertconfig.SyntheticAlertConfig](
			api.restClient,
			syntheticalertconfig.SyntheticAlertConfigsResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePOST,
			rest.NewGenericUnmarshaller[*syntheticalertconfig.SyntheticAlertConfig](),
		)
	}
	return api.syntheticAlertConfigs
}

// SyntheticLocations returns the synthetic locations client (lazy initialization)
func (api *instanaAPI) SyntheticLocations() rest.ReadOnlyRestResource[*syntheticlocation.SyntheticLocation] {
	if api.syntheticLocations == nil {
		api.syntheticLocations = NewReadOnlyRestResource[*syntheticlocation.SyntheticLocation](
			api.restClient,
			syntheticlocation.ResourcePath,
			rest.NewGenericUnmarshaller[*syntheticlocation.SyntheticLocation](),
		)
	}
	return api.syntheticLocations
}

// SyntheticTests returns the synthetic tests client (lazy initialization)
func (api *instanaAPI) SyntheticTests() rest.RestResource[*synthetictest.SyntheticTest] {
	if api.syntheticTests == nil {
		api.syntheticTests = NewRestResource[*synthetictest.SyntheticTest](
			api.restClient,
			synthetictest.ResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			rest.NewGenericUnmarshaller[*synthetictest.SyntheticTest](),
		)
	}
	return api.syntheticTests
}

// Teams returns the teams client (lazy initialization)
func (api *instanaAPI) Teams() rest.RestResource[*team.Team] {
	if api.teams == nil {
		api.teams = NewRestResource[*team.Team](
			api.restClient,
			team.ResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			rest.NewGenericUnmarshaller[*team.Team](),
		)
	}
	return api.teams
}

// Users returns the users client (lazy initialization)
func (api *instanaAPI) Users() rest.ReadOnlyRestResource[*user.User] {
	if api.users == nil {
		api.users = NewReadOnlyRestResource[*user.User](
			api.restClient,
			user.ResourcePath,
			rest.NewGenericUnmarshaller[*user.User](),
		)
	}
	return api.users
}

// WebsiteAlertConfigs returns the website alert configurations client (lazy initialization)
func (api *instanaAPI) WebsiteAlertConfigs() rest.RestResource[*websitealertconfig.WebsiteAlertConfig] {
	if api.websiteAlertConfigs == nil {
		api.websiteAlertConfigs = NewRestResource[*websitealertconfig.WebsiteAlertConfig](
			api.restClient,
			websitealertconfig.ResourcePath,
			rest.DefaultRestResourceModeCreateAndUpdatePOST,
			rest.NewGenericUnmarshaller[*websitealertconfig.WebsiteAlertConfig](),
		)
	}
	return api.websiteAlertConfigs
}

// WebsiteMonitoringConfigs returns the website monitoring configurations client (lazy initialization)
func (api *instanaAPI) WebsiteMonitoringConfigs() rest.RestResource[*websitemonitoring.WebsiteMonitoringConfig] {
	if api.websiteMonitoringConfigs == nil {
		api.websiteMonitoringConfigs = NewRestResource[*websitemonitoring.WebsiteMonitoringConfig](
			api.restClient,
			websitemonitoring.WebsiteMonitoringConfigResourcePath,
			rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
			rest.NewGenericUnmarshaller[*websitemonitoring.WebsiteMonitoringConfig](),
		)
	}
	return api.websiteMonitoringConfigs
}

// Made with Bob
