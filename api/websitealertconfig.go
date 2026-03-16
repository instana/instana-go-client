package api

import (
	"github.com/instana/instana-go-client/shared/tagfilter"
	"github.com/instana/instana-go-client/shared/types"
)

// WebsiteAlertConfigResourcePath path to website alert config resource of Instana RESTful API
const WebsiteAlertConfigResourcePath = "/api/events/settings/website-alert-configs"

// WebsiteAlertConfig is the representation of an website alert configuration in Instana
type WebsiteAlertConfig struct {
	ID                    string                           `json:"id"`
	Name                  string                           `json:"name"`
	Description           string                           `json:"description"`
	Severity              *int                             `json:"severity"`
	Triggering            bool                             `json:"triggering"`
	Enabled               *bool                            `json:"enabled,omitempty"`
	WebsiteID             string                           `json:"websiteId"`
	TagFilterExpression   *tagfilter.TagFilter             `json:"tagFilterExpression"`
	AlertChannelIDs       []string                         `json:"alertChannelIds"`
	Granularity           types.Granularity                `json:"granularity"`
	CustomerPayloadFields []types.CustomPayloadField[any]  `json:"customPayloadFields"`
	Rule                  *WebsiteAlertRule                `json:"rule"`
	Threshold             *types.Threshold                 `json:"threshold"`
	TimeThreshold         types.WebsiteTimeThreshold       `json:"timeThreshold"`
	Rules                 []WebsiteAlertRuleWithThresholds `json:"rules"`
}

type WebsiteAlertRuleWithThresholds struct {
	Rule              *WebsiteAlertRule                           `json:"rule"`
	ThresholdOperator string                                      `json:"thresholdOperator"`
	Thresholds        map[types.AlertSeverity]types.ThresholdRule `json:"thresholds"`
}

// GetIDForResourcePath implementation of the interface InstanaDataObject
func (r *WebsiteAlertConfig) GetIDForResourcePath() string {
	return r.ID
}

// GetCustomerPayloadFields implementation of the interface customPayloadFieldsAwareInstanaDataObject
func (a *WebsiteAlertConfig) GetCustomerPayloadFields() []types.CustomPayloadField[any] {
	return a.CustomerPayloadFields
}

// SetCustomerPayloadFields implementation of the interface customPayloadFieldsAwareInstanaDataObject
func (a *WebsiteAlertConfig) SetCustomerPayloadFields(fields []types.CustomPayloadField[any]) {
	a.CustomerPayloadFields = fields
}

// WebsiteAlertRule struct representing the API model of a website alert rule
type WebsiteAlertRule struct {
	AlertType   string                    `json:"alertType"`
	MetricName  string                    `json:"metricName"`
	Aggregation *types.Aggregation        `json:"aggregation"`
	Operator    *types.ExpressionOperator `json:"operator"`
	Value       *string                   `json:"value"`
}

// WebsiteImpactMeasurementMethod custom type for impact measurement method of website alert rules
type WebsiteImpactMeasurementMethod = types.WebsiteImpactMeasurementMethod

// WebsiteImpactMeasurementMethods custom type for a slice of WebsiteImpactMeasurementMethod
type WebsiteImpactMeasurementMethods []types.WebsiteImpactMeasurementMethod

// ToStringSlice Returns the corresponding string representations
func (methods WebsiteImpactMeasurementMethods) ToStringSlice() []string {
	result := make([]string, len(methods))
	for i, v := range methods {
		result[i] = string(v)
	}
	return result
}

const (
	//types.WebsiteImpactMeasurementMethodAggregated constant value for the website impact measurement method aggregated
	WebsiteImpactMeasurementMethodAggregated = types.WebsiteImpactMeasurementMethod("AGGREGATED")
	//types.WebsiteImpactMeasurementMethodPerWindow constant value for the website impact measurement method per_window
	WebsiteImpactMeasurementMethodPerWindow = WebsiteImpactMeasurementMethod("PER_WINDOW")
)

// SupportedWebsiteImpactMeasurementMethods list of all supported WebsiteImpactMeasurementMethod
var SupportedWebsiteImpactMeasurementMethods = WebsiteImpactMeasurementMethods{types.WebsiteImpactMeasurementMethodAggregated, types.WebsiteImpactMeasurementMethodPerWindow}
