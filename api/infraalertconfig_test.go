package api_test

import (
	"encoding/json"
	"testing"

	. "github.com/instana/instana-go-client/api"
	tagfilter "github.com/instana/instana-go-client/shared/tagfilter"
	"github.com/instana/instana-go-client/shared/types"
)

func TestInfraAlertConfigResourcePath(t *testing.T) {
	expected := "/api/events/settings/infra-alert-configs"
	if InfraAlertConfigResourcePath != expected {
		t.Errorf("Expected InfraAlertConfigResourcePath to be %s, got %s", expected, InfraAlertConfigResourcePath)
	}
}

func TestInfraAlertConfigGetIDForResourcePath(t *testing.T) {
	testID := "test-infra-alert-123"
	config := &InfraAlertConfig{
		ID:   testID,
		Name: "Test Infra Alert",
	}

	result := config.GetIDForResourcePath()
	if result != testID {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", testID, result)
	}
}

func TestInfraAlertConfigStructure(t *testing.T) {
	config := InfraAlertConfig{
		ID:          "infra-456",
		Name:        "CPU Alert",
		Description: "High CPU usage",
	}

	if config.ID != "infra-456" {
		t.Errorf("Expected ID 'infra-456', got %s", config.ID)
	}
	if config.Name != "CPU Alert" {
		t.Errorf("Expected Name 'CPU Alert', got %s", config.Name)
	}
}

func TestInfraAlertConfigCustomPayloadFields(t *testing.T) {
	config := &InfraAlertConfig{
		ID:   "test-id",
		Name: "Test Config",
	}

	// Test GetCustomerPayloadFields - initially empty slice, not nil
	fields := config.GetCustomerPayloadFields()
	if fields == nil {
		fields = []types.CustomPayloadField[any]{}
	}
	if len(fields) != 0 {
		t.Errorf("Expected 0 initial custom payload fields, got %d", len(fields))
	}

	// Test SetCustomerPayloadFields
	newFields := []types.CustomPayloadField[any]{
		{Key: "field1", Value: "value1"},
		{Key: "field2", Value: 123},
	}
	config.SetCustomerPayloadFields(newFields)

	retrievedFields := config.GetCustomerPayloadFields()
	if len(retrievedFields) != 2 {
		t.Errorf("Expected 2 custom payload fields, got %d", len(retrievedFields))
	}
	if retrievedFields[0].Key != "field1" {
		t.Errorf("Expected first field key 'field1', got %s", retrievedFields[0].Key)
	}
}

func TestInfraToStringSlice(t *testing.T) {
	typeval := InfraAlertEvaluationTypes{
		EvaluationTypePerEntity,
	}
	typeSet := typeval.ToStringSlice()
	if typeSet[0] != "PER_ENTITY" {
		t.Error("ToStringSlice not working correctly")
	}
}

func TestInfraAlertConfigMultiRuleAndJSON(t *testing.T) {
	thresholdVal := 80.0
	config := &InfraAlertConfig{
		ID:                  "infra-123",
		Name:                "Multi Rule Infra Alert",
		Description:         "Alert with multiple rules and metric filters",
		RuleLogicalOperator: "AND",
		EvaluationType:      EvaluationTypePerEntity,
		Triggering:          true,
		Rules: []types.RuleWithThreshold[InfraAlertRule]{
			{
				ThresholdOperator: types.ThresholdOperatorGreaterThan,
				Rule: InfraAlertRule{
					AlertType:              "system",
					MetricName:             "cpu.used",
					EntityType:             "host",
					Aggregation:            types.MeanAggregation,
					CrossSeriesAggregation: types.SumAggregation,
					Regex:                  false,
					MetricGroupBy:          []string{"host.name"},
					MetricTagFilterExpression: &tagfilter.TagFilter{
						Type:  tagfilter.TagFilterType,
						Name:  func() *string { s := "agent.tag"; return &s }(),
						Value: "production",
					},
				},
				Thresholds: map[types.AlertSeverity]types.ThresholdRule{
					types.CriticalSeverity: {
						Type:  "static",
						Value: &thresholdVal,
					},
				},
			},
		},
	}

	data, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("Failed to marshal InfraAlertConfig: %v", err)
	}

	var unmarshalled InfraAlertConfig
	if err := json.Unmarshal(data, &unmarshalled); err != nil {
		t.Fatalf("Failed to unmarshal InfraAlertConfig: %v", err)
	}

	if unmarshalled.RuleLogicalOperator != "AND" {
		t.Errorf("Expected RuleLogicalOperator 'AND', got %s", unmarshalled.RuleLogicalOperator)
	}
	if len(unmarshalled.Rules) != 1 {
		t.Fatalf("Expected 1 rule, got %d", len(unmarshalled.Rules))
	}
	rule := unmarshalled.Rules[0].Rule
	if len(rule.MetricGroupBy) != 1 || rule.MetricGroupBy[0] != "host.name" {
		t.Errorf("Expected MetricGroupBy ['host.name'], got %v", rule.MetricGroupBy)
	}
	if rule.MetricTagFilterExpression == nil || rule.MetricTagFilterExpression.Value != "production" {
		t.Errorf("Expected MetricTagFilterExpression with Value 'production', got %+v", rule.MetricTagFilterExpression)
	}
}
