package alertingchannel

// ResourcePath is the path to the Alerting Channels resource in the Instana API
const AlertingchannelResourcePath = "/api/events/settings/alertingChannels"

// AlertingChannelType represents the type of alerting channel
type AlertingChannelType string

// Alerting channel type constants
const (
	// EmailChannelType constant value for alerting channel type EMAIL
	EmailChannelType = AlertingChannelType("EMAIL")
	// GoogleChatChannelType constant value for alerting channel type GOOGLE_CHAT
	GoogleChatChannelType = AlertingChannelType("GOOGLE_CHAT")
	// Office365ChannelType constant value for alerting channel type OFFICE_365
	Office365ChannelType = AlertingChannelType("OFFICE_365")
	// OpsGenieChannelType constant value for alerting channel type OPS_GENIE
	OpsGenieChannelType = AlertingChannelType("OPS_GENIE")
	// PagerDutyChannelType constant value for alerting channel type PAGER_DUTY
	PagerDutyChannelType = AlertingChannelType("PAGER_DUTY")
	// SlackChannelType constant value for alerting channel type SLACK
	SlackChannelType = AlertingChannelType("SLACK")
	// SplunkChannelType constant value for alerting channel type SPLUNK
	SplunkChannelType = AlertingChannelType("SPLUNK")
	// VictorOpsChannelType constant value for alerting channel type VICTOR_OPS
	VictorOpsChannelType = AlertingChannelType("VICTOR_OPS")
	// WebhookChannelType constant value for alerting channel type WEB_HOOK
	WebhookChannelType = AlertingChannelType("WEB_HOOK")
	// ServiceNowChannelType constant value for alerting channel type SERVICE_NOW_WEBHOOK
	ServiceNowChannelType = AlertingChannelType("SERVICE_NOW_WEBHOOK")
	// ServiceNowApplicationChannelType constant value for alerting channel type SERVICE_NOW_APPLICATION
	ServiceNowApplicationChannelType = AlertingChannelType("SERVICE_NOW_APPLICATION")
	// PrometheusWebhookChannelType constant value for alerting channel type PROMETHEUS_WEBHOOK
	PrometheusWebhookChannelType = AlertingChannelType("PROMETHEUS_WEBHOOK")
	// WebexTeamsWebhookChannelType constant value for alerting channel type WEBEX_TEAMS_WEBHOOK
	WebexTeamsWebhookChannelType = AlertingChannelType("WEBEX_TEAMS_WEBHOOK")
	// WatsonAIOpsWebhookChannelType constant value for alerting channel type WATSON_AIOPS_WEBHOOK
	WatsonAIOpsWebhookChannelType = AlertingChannelType("WATSON_AIOPS_WEBHOOK")
	// SlackAppChannelType constant value for alerting channel type BIDIRECTIONAL_SLACK
	SlackAppChannelType = AlertingChannelType("BIDIRECTIONAL_SLACK")
	// MsTeamsAppChannelType constant value for alerting channel type BIDIRECTIONAL_MS_TEAMS
	MsTeamsAppChannelType = AlertingChannelType("BIDIRECTIONAL_MS_TEAMS")
)

// AlertingChannel represents an alerting channel in Instana
type AlertingChannel struct {
	ID                    string              `json:"id"`
	Name                  string              `json:"name"`
	Kind                  AlertingChannelType `json:"kind"`
	Emails                []string            `json:"emails"`
	WebhookURL            *string             `json:"webhookUrl"`
	APIKey                *string             `json:"apiKey"`
	Tags                  *string             `json:"tags"`
	Region                *string             `json:"region"`
	RoutingKey            *string             `json:"routingKey"`
	ServiceIntegrationKey *string             `json:"serviceIntegrationKey"`
	IconURL               *string             `json:"iconUrl"`
	Channel               *string             `json:"channel"`
	URL                   *string             `json:"url"`
	Token                 *string             `json:"token"`
	WebhookURLs           []string            `json:"webhookUrls"`
	Headers               []string            `json:"headers"`
	// ServiceNow fields
	ServiceNowURL      *string `json:"serviceNowUrl"`
	Username           *string `json:"username"`
	Password           *string `json:"password"`
	AutoCloseIncidents *bool   `json:"autoCloseIncidents"`
	// ServiceNow Enhanced (ITSM) fields
	Tenant                         *string `json:"tenant"`
	Unit                           *string `json:"unit"`
	InstanaURL                     *string `json:"instanaUrl"`
	EnableSendInstanaNotes         *bool   `json:"enableSendInstanaNotes"`
	EnableSendServiceNowActivities *bool   `json:"enableSendServiceNowActivities"`
	EnableSendServiceNowWorkNotes  *bool   `json:"enableSendServiceNowWorkNotes"`
	ManuallyClosedIncidents        *bool   `json:"manuallyClosedIncidents"`
	ResolutionOfIncident           *bool   `json:"resolutionOfIncident"`
	SnowStatusOnCloseEvent         *int    `json:"snowStatusOnCloseEvent"`
	// Prometheus Webhook fields
	Receiver *string `json:"receiver"`
	// Bidirectional Slack App fields
	AppID          *string `json:"appId"`
	TeamID         *string `json:"teamId"`
	TeamName       *string `json:"teamName"`
	ChannelID      *string `json:"channelId"`
	ChannelName    *string `json:"channelName"`
	EmojiRendering *bool   `json:"emojiRendering"`
	// Bidirectional MS Teams App fields
	APITokenID *string `json:"apiTokenId"`
	ServiceURL *string `json:"serviceUrl"`
	TenantID   *string `json:"tenantId"`
	TenantName *string `json:"tenantName"`
}

// GetIDForResourcePath returns the ID to be used for the resource path
func (a *AlertingChannel) GetIDForResourcePath() string {
	return a.ID
}

// Made with Bob
