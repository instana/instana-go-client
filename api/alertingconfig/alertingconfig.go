package alertingconfig

import (
	"github.com/instana/instana-go-client/shared/types"
)

// AlertsResourcePath path to Alerts resource of Instana RESTful API
const AlertsResourcePath = "/api/events/settings" + "/alerts"

// EventFilteringConfiguration type definiton of an EventFilteringConfiguration of a AlertingConfiguration of the Instana ReST AOI
type EventFilteringConfiguration struct {
	Query                     *string          `json:"query"`
	RuleIDs                   []string         `json:"ruleIds"`
	EventTypes                []AlertEventType `json:"eventTypes"`
	ApplicationAlertConfigIds []string         `json:"applicationAlertConfigIds"`
}

// AlertingConfiguration type definition of an Alerting Configuration in Instana REST API
type AlertingConfiguration struct {
	ID                          string                          `json:"id"`
	AlertName                   string                          `json:"alertName"`
	IntegrationIDs              []string                        `json:"integrationIds"`
	EventFilteringConfiguration EventFilteringConfiguration     `json:"eventFilteringConfiguration"`
	CustomerPayloadFields       []types.CustomPayloadField[any] `json:"customPayloadFields"`
}

// GetIDForResourcePath implementation of the interface InstanaDataObject
func (c *AlertingConfiguration) GetIDForResourcePath() string {
	return c.ID
}

// GetCustomerPayloadFields implementation of the interface customPayloadFieldsAwareInstanaDataObject
func (a *AlertingConfiguration) GetCustomerPayloadFields() []types.CustomPayloadField[any] {
	return a.CustomerPayloadFields
}

// SetCustomerPayloadFields implementation of the interface customPayloadFieldsAwareInstanaDataObject
func (a *AlertingConfiguration) SetCustomerPayloadFields(fields []types.CustomPayloadField[any]) {
	a.CustomerPayloadFields = fields
}

// AlertEventType type definition of EventTypes of an Instana Alert
type AlertEventType string

const (
	//IncidentAlertEventType constant value for alert event type incident
	IncidentAlertEventType = AlertEventType("incident")
	//CriticalAlertEventType constant value for alert event type critical
	CriticalAlertEventType = AlertEventType("critical")
	//WarningAlertEventType constant value for alert event type warning
	WarningAlertEventType = AlertEventType("warning")
	//ChangeAlertEventType constant value for alert event type change
	ChangeAlertEventType = AlertEventType("change")
	//OnlineAlertEventType constant value for alert event type online
	OnlineAlertEventType = AlertEventType("online")
	//OfflineAlertEventType constant value for alert event type offline
	OfflineAlertEventType = AlertEventType("offline")
	//NoneAlertEventType constant value for alert event type none
	NoneAlertEventType = AlertEventType("none")
	//AgentMonitoringIssueEventType constant value for alert event type none
	AgentMonitoringIssueEventType = AlertEventType("agent_monitoring_issue")
)

// SupportedAlertEventTypes list of supported alert event types of Instana API
var SupportedAlertEventTypes = []AlertEventType{
	IncidentAlertEventType,
	CriticalAlertEventType,
	WarningAlertEventType,
	ChangeAlertEventType,
	OnlineAlertEventType,
	OfflineAlertEventType,
	NoneAlertEventType,
	AgentMonitoringIssueEventType,
}
