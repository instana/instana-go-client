package infraalertconfig

import (
	"github.com/instana/instana-go-client/shared/tagfilter"
	"github.com/instana/instana-go-client/shared/types"
)

// ResourcePath is the path to the Infrastructure Alert Configurations resource in the Instana API
const ResourcePath = "/api/events/settings/infra-alert-configs"

const InfraAlertConfigResourcePath = "/api/events/settings" + "/infra-alert-configs"

type InfraAlertConfig struct {
	ID                    string                                    `json:"id"`
	Name                  string                                    `json:"name"`
	Description           string                                    `json:"description"`
	TagFilterExpression   *tagfilter.TagFilter                      `json:"tagFilterExpression"`
	GroupBy               []string                                  `json:"groupBy"`
	Granularity           types.Granularity                         `json:"granularity"`
	TimeThreshold         *InfraTimeThreshold                       `json:"timeThreshold"`
	CustomerPayloadFields []types.CustomPayloadField[any]           `json:"customPayloadFields"`
	Rules                 []types.RuleWithThreshold[InfraAlertRule] `json:"rules"`
	AlertChannels         map[types.AlertSeverity][]string          `json:"alertChannels"`
	EvaluationType        InfraAlertEvaluationType                  `json:"evaluationType"`
	Triggering            bool                                      `json:"triggering"`
}

func (config *InfraAlertConfig) GetIDForResourcePath() string {
	return config.ID
}

func (config *InfraAlertConfig) GetCustomerPayloadFields() []types.CustomPayloadField[any] {
	return config.CustomerPayloadFields
}

func (config *InfraAlertConfig) SetCustomerPayloadFields(fields []types.CustomPayloadField[any]) {
	config.CustomerPayloadFields = fields
}

type InfraAlertRule struct {
	AlertType              string            `json:"alertType"`
	MetricName             string            `json:"metricName"`
	EntityType             string            `json:"entityType"`
	Aggregation            types.Aggregation `json:"aggregation"`
	CrossSeriesAggregation types.Aggregation `json:"crossSeriesAggregation"`
	Regex                  bool              `json:"regex"`
}

// InfraAlertEvaluationType custom type representing the infrastructure alert evaluation type from the Instana API
type InfraAlertEvaluationType string

// InfraAlertEvaluationTypes custom type representing a slice of InfraAlertEvaluationType
type InfraAlertEvaluationTypes []InfraAlertEvaluationType

// ToStringSlice returns the corresponding string representations
func (types InfraAlertEvaluationTypes) ToStringSlice() []string {
	result := make([]string, len(types))
	for i, v := range types {
		result[i] = string(v)
	}
	return result
}

const (
	// EvaluationTypePerEntity constant value for InfraAlertEvaluationType PER_ENTITY
	EvaluationTypePerEntity = InfraAlertEvaluationType("PER_ENTITY")
	// EvaluationTypeCustom constant value for InfraAlertEvaluationType CUSTOM
	EvaluationTypeCustom = InfraAlertEvaluationType("CUSTOM")
)

// SupportedInfraAlertEvaluationTypes list of all supported InfraAlertEvaluationTypes
var SupportedInfraAlertEvaluationTypes = InfraAlertEvaluationTypes{
	EvaluationTypePerEntity,
	EvaluationTypeCustom,
}

type InfraTimeThreshold struct {
	Type       string `json:"type"`
	TimeWindow int64  `json:"timeWindow"`
}
