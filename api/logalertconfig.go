package api

import (
	"github.com/instana/instana-go-client/shared/tagfilter"
	"github.com/instana/instana-go-client/shared/types"
)

const (
	// LogAlertConfigResourcePath path to Log Alert Config resource of Instana RESTful API
	LogAlertConfigResourcePath = "/api/events/settings/global-alert-configs/logs"
)

// LogAlertRule represents the rule configuration for log alerts
type LogAlertRule struct {
	AlertType   string            `json:"alertType"`
	MetricName  string            `json:"metricName"`
	Aggregation types.Aggregation `json:"aggregation,omitempty"`
}

// LogTimeThreshold represents the time threshold configuration for log alerts
type LogTimeThreshold struct {
	Type       string `json:"type"`
	TimeWindow int64  `json:"timeWindow"`
}

// GroupByTag represents a tag used for grouping in log alerts
type GroupByTag struct {
	TagName string `json:"tagName"`
	Key     string `json:"key,omitempty"`
}

// LogAlertConfig represents the Instana API model for log alert configurations
type LogAlertConfig struct {
	ID                    string                                  `json:"id,omitempty"`
	Name                  string                                  `json:"name"`
	Description           string                                  `json:"description"`
	TagFilterExpression   *tagfilter.TagFilter                    `json:"tagFilterExpression"`
	AlertChannels         map[types.AlertSeverity][]string        `json:"alertChannels,omitempty"`
	Granularity           types.Granularity                       `json:"granularity"`
	TimeThreshold         *LogTimeThreshold                       `json:"timeThreshold"`
	GracePeriod           int64                                   `json:"gracePeriod,omitempty"`
	CustomerPayloadFields []types.CustomPayloadField[any]         `json:"customPayloadFields,omitempty"`
	Rules                 []types.RuleWithThreshold[LogAlertRule] `json:"rules"`
	GroupBy               []GroupByTag                            `json:"groupBy,omitempty"`
}

// GetIDForResourcePath implementation of the InstanaDataObject interface
func (r *LogAlertConfig) GetIDForResourcePath() string {
	return r.ID
}

// GetCustomerPayloadFields implementation of the customPayloadFieldsAwareInstanaDataObject interface
func (r *LogAlertConfig) GetCustomerPayloadFields() []types.CustomPayloadField[any] {
	return r.CustomerPayloadFields
}

// SetCustomerPayloadFields implementation of the customPayloadFieldsAwareInstanaDataObject interface
func (r *LogAlertConfig) SetCustomerPayloadFields(fields []types.CustomPayloadField[any]) {
	r.CustomerPayloadFields = fields
}

// LogLevel custom type for log level
type LogLevel = types.LogLevel

// LogLevels custom type for a slice of LogLevel
type LogLevels []types.LogLevel

// ToStringSlice Returns the corresponding string representations
func (levels LogLevels) ToStringSlice() []string {
	result := make([]string, len(levels))
	for i, v := range levels {
		result[i] = string(v)
	}
	return result
}

const (
	//LogLevelWarning constant value for the warning log level
	LogLevelWarning = types.LogLevel("WARN")
	//types.LogLevelError constant value for the error log level
	LogLevelError = LogLevel("ERROR")
	//types.LogLevelAny constant value for the any log level
	LogLevelAny = LogLevel("ANY")
)

// SupportedLogLevels list of all supported LogLevel
var SupportedLogLevels = LogLevels{LogLevelWarning, types.LogLevelError, types.LogLevelAny}
