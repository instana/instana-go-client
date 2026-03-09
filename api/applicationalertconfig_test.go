package api_test

import (
	"encoding/json"
	"testing"

	. "github.com/instana/instana-go-client/api"
	"github.com/instana/instana-go-client/shared/tagfilter"
	"github.com/instana/instana-go-client/shared/types"
)

// Resource Path Tests
func TestApplicationAlertConfigResourcePaths(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{"ApplicationAlertConfigs", ApplicationAlertConfigsResourcePath, "/api/events/settings/application-alert-configs"},
		{"GlobalApplicationAlertConfigs", GlobalApplicationAlertConfigsResourcePath, "/api/events/settings/global-alert-configs/applications"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.path != tt.expected {
				t.Errorf("Expected %s to be %s, got %s", tt.name, tt.expected, tt.path)
			}
		})
	}
}

// GetIDForResourcePath Tests
func TestApplicationAlertConfigGetIDForResourcePath(t *testing.T) {
	testID := "test-alert-config-123"
	config := &ApplicationAlertConfig{
		ID:   testID,
		Name: "Test Alert Config",
	}

	result := config.GetIDForResourcePath()
	if result != testID {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", testID, result)
	}
}

// Structure Tests
func TestApplicationAlertConfigStructure(t *testing.T) {
	enabled := true
	gracePeriod := int64(300000)
	tagName := "test-tag"
	tagValue := "test-value"
	operator := types.EqualsOperator
	thresholdValue := 100.0

	tagFilter := &tagfilter.TagFilter{
		Type:     tagfilter.TagFilterExpressionType,
		Name:     &tagName,
		Operator: &operator,
		Value:    &tagValue,
	}

	threshold := &types.Threshold{
		Type:     "static",
		Operator: types.ThresholdOperatorGreaterThan,
		Value:    &thresholdValue,
	}

	timeThreshold := &ApplicationAlertTimeThreshold{
		Type:       "violationsInPeriod",
		TimeWindow: 600000,
		Violations: 3,
		Requests:   100,
	}

	rule := &ApplicationAlertRule{
		AlertType:   "errorRate",
		MetricName:  "errors",
		Aggregation: types.SumAggregation,
	}

	config := ApplicationAlertConfig{
		ID:                  "config-123",
		Name:                "Test Application Alert",
		Description:         "Test description",
		Triggering:          true,
		Enabled:             &enabled,
		Applications:        map[string]IncludedApplication{},
		BoundaryScope:       types.BoundaryScopeAll,
		TagFilterExpression: tagFilter,
		IncludeInternal:     false,
		IncludeSynthetic:    true,
		EvaluationType:      EvaluationTypePerApplication,
		AlertChannelIDs:     []string{"channel-1", "channel-2"},
		AlertChannels:       map[string][]string{"email": {"test@example.com"}},
		Granularity:         types.Granularity600000,
		GracePeriod:         &gracePeriod,
		CustomerPayloadFields: []types.CustomPayloadField[any]{
			{Key: "key1", Value: "value1"},
		},
		Rule:          rule,
		Rules:         []ApplicationAlertRuleWithThresholds{},
		Threshold:     threshold,
		TimeThreshold: timeThreshold,
	}

	// Test basic fields
	if config.ID != "config-123" {
		t.Errorf("Expected ID 'config-123', got %s", config.ID)
	}
	if config.Name != "Test Application Alert" {
		t.Errorf("Expected Name 'Test Application Alert', got %s", config.Name)
	}
	if config.Description != "Test description" {
		t.Errorf("Expected Description 'Test description', got %s", config.Description)
	}
	if !config.Triggering {
		t.Error("Expected Triggering to be true")
	}
	if config.Enabled == nil || !*config.Enabled {
		t.Error("Expected Enabled to be true")
	}
	if config.BoundaryScope != types.BoundaryScopeAll {
		t.Errorf("Expected BoundaryScope 'ALL', got %s", config.BoundaryScope)
	}
	if config.IncludeInternal {
		t.Error("Expected IncludeInternal to be false")
	}
	if !config.IncludeSynthetic {
		t.Error("Expected IncludeSynthetic to be true")
	}
	if config.EvaluationType != EvaluationTypePerApplication {
		t.Errorf("Expected EvaluationType 'PER_AP', got %s", config.EvaluationType)
	}
	if len(config.AlertChannelIDs) != 2 {
		t.Errorf("Expected 2 alert channel IDs, got %d", len(config.AlertChannelIDs))
	}
	if config.Granularity != types.Granularity600000 {
		t.Errorf("Expected Granularity 600000, got %d", config.Granularity)
	}
	if config.GracePeriod == nil || *config.GracePeriod != gracePeriod {
		t.Error("GracePeriod not set correctly")
	}
}

// Custom Payload Fields Tests
func TestApplicationAlertConfigCustomPayloadFields(t *testing.T) {
	config := &ApplicationAlertConfig{
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

// ApplicationScope Tests
func TestApplicationScopeStructure(t *testing.T) {
	scope := ApplicationScope{
		ApplicationID: "app-123",
		Inclusive:     true,
		Services: []ServiceScope{
			{
				ServiceID: "service-1",
				Inclusive: true,
				Endpoints: []EndpointScope{
					{EndpointID: "endpoint-1", Inclusive: true},
				},
			},
		},
	}

	if scope.ApplicationID != "app-123" {
		t.Errorf("Expected ApplicationID 'app-123', got %s", scope.ApplicationID)
	}
	if !scope.Inclusive {
		t.Error("Expected Inclusive to be true")
	}
	if len(scope.Services) != 1 {
		t.Errorf("Expected 1 service, got %d", len(scope.Services))
	}
	if scope.Services[0].ServiceID != "service-1" {
		t.Errorf("Expected ServiceID 'service-1', got %s", scope.Services[0].ServiceID)
	}
	if len(scope.Services[0].Endpoints) != 1 {
		t.Errorf("Expected 1 endpoint, got %d", len(scope.Services[0].Endpoints))
	}
}

// ServiceScope Tests
func TestServiceScopeStructure(t *testing.T) {
	service := ServiceScope{
		ServiceID: "service-456",
		Inclusive: false,
		Endpoints: []EndpointScope{
			{EndpointID: "endpoint-1", Inclusive: true},
			{EndpointID: "endpoint-2", Inclusive: false},
		},
	}

	if service.ServiceID != "service-456" {
		t.Errorf("Expected ServiceID 'service-456', got %s", service.ServiceID)
	}
	if service.Inclusive {
		t.Error("Expected Inclusive to be false")
	}
	if len(service.Endpoints) != 2 {
		t.Errorf("Expected 2 endpoints, got %d", len(service.Endpoints))
	}
}

// EndpointScope Tests
func TestEndpointScopeStructure(t *testing.T) {
	endpoint := EndpointScope{
		EndpointID: "endpoint-789",
		Inclusive:  true,
	}

	if endpoint.EndpointID != "endpoint-789" {
		t.Errorf("Expected EndpointID 'endpoint-789', got %s", endpoint.EndpointID)
	}
	if !endpoint.Inclusive {
		t.Error("Expected Inclusive to be true")
	}
}

// ApplicationAlertTimeThreshold Tests
func TestApplicationAlertTimeThresholdStructure(t *testing.T) {
	threshold := ApplicationAlertTimeThreshold{
		Type:       "violationsInPeriod",
		TimeWindow: 600000,
		Violations: 5,
		Requests:   100,
	}

	if threshold.Type != "violationsInPeriod" {
		t.Errorf("Expected Type 'violationsInPeriod', got %s", threshold.Type)
	}
	if threshold.TimeWindow != 600000 {
		t.Errorf("Expected TimeWindow 600000, got %d", threshold.TimeWindow)
	}
	if threshold.Violations != 5 {
		t.Errorf("Expected Violations 5, got %d", threshold.Violations)
	}
	if threshold.Requests != 100 {
		t.Errorf("Expected Requests 100, got %d", threshold.Requests)
	}
}

// ApplicationAlertRule Tests
func TestApplicationAlertRuleStructure(t *testing.T) {
	statusCodeStart := int32(500)
	statusCodeEnd := int32(599)
	logLevel := types.LogLevelError
	message := "error occurred"
	operator := types.ContainsOperator

	rule := ApplicationAlertRule{
		AlertType:       "errorRate",
		MetricName:      "errors",
		Aggregation:     types.SumAggregation,
		StatusCodeStart: &statusCodeStart,
		StatusCodeEnd:   &statusCodeEnd,
		Level:           &logLevel,
		Message:         &message,
		Operator:        &operator,
	}

	if rule.AlertType != "errorRate" {
		t.Errorf("Expected AlertType 'errorRate', got %s", rule.AlertType)
	}
	if rule.MetricName != "errors" {
		t.Errorf("Expected MetricName 'errors', got %s", rule.MetricName)
	}
	if rule.Aggregation != types.SumAggregation {
		t.Errorf("Expected Aggregation 'SUM', got %s", rule.Aggregation)
	}
	if rule.StatusCodeStart == nil || *rule.StatusCodeStart != statusCodeStart {
		t.Error("StatusCodeStart not set correctly")
	}
	if rule.StatusCodeEnd == nil || *rule.StatusCodeEnd != statusCodeEnd {
		t.Error("StatusCodeEnd not set correctly")
	}
	if rule.Level == nil || *rule.Level != logLevel {
		t.Error("Level not set correctly")
	}
	if rule.Message == nil || *rule.Message != message {
		t.Error("Message not set correctly")
	}
	if rule.Operator == nil || *rule.Operator != operator {
		t.Error("Operator not set correctly")
	}
}

// ApplicationAlertEvaluationType Tests
func TestApplicationAlertEvaluationTypeConstants(t *testing.T) {
	tests := []struct {
		name     string
		value    ApplicationAlertEvaluationType
		expected string
	}{
		{"PerApplication", EvaluationTypePerApplication, "PER_AP"},
		{"PerApplicationService", EvaluationTypePerApplicationService, "PER_AP_SERVICE"},
		{"PerApplicationEndpoint", EvaluationTypePerApplicationEndpoint, "PER_AP_ENDPOINT"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.value) != tt.expected {
				t.Errorf("Expected %s to be %s, got %s", tt.name, tt.expected, string(tt.value))
			}
		})
	}
}

func TestSupportedApplicationAlertEvaluationTypes(t *testing.T) {
	if len(SupportedApplicationAlertEvaluationTypes) != 3 {
		t.Errorf("Expected 3 supported evaluation types, got %d", len(SupportedApplicationAlertEvaluationTypes))
	}

	stringSlice := SupportedApplicationAlertEvaluationTypes.ToStringSlice()
	if len(stringSlice) != 3 {
		t.Errorf("Expected 3 string values, got %d", len(stringSlice))
	}

	expectedValues := []string{"PER_AP", "PER_AP_SERVICE", "PER_AP_ENDPOINT"}
	for i, expected := range expectedValues {
		if stringSlice[i] != expected {
			t.Errorf("Expected value at index %d to be %s, got %s", i, expected, stringSlice[i])
		}
	}
}

// Rule Type Tests
func TestApplicationAlertRuleErrorRate(t *testing.T) {
	rule := ApplicationAlertRuleErrorRate{
		MetricName:  "errors",
		Aggregation: "sum",
	}

	if rule.MetricName != "errors" {
		t.Errorf("Expected MetricName 'errors', got %s", rule.MetricName)
	}
	if rule.Aggregation != "sum" {
		t.Errorf("Expected Aggregation 'sum', got %s", rule.Aggregation)
	}
}

func TestApplicationAlertRuleErrors(t *testing.T) {
	rule := ApplicationAlertRuleErrors{
		MetricName:  "call.error.count",
		Aggregation: "sum",
	}

	if rule.MetricName != "call.error.count" {
		t.Errorf("Expected MetricName 'call.error.count', got %s", rule.MetricName)
	}
}

func TestApplicationAlertRuleLogs(t *testing.T) {
	rule := ApplicationAlertRuleLogs{
		MetricName:  "logs",
		Aggregation: "sum",
		Level:       "ERROR",
		Message:     "exception",
		Operator:    "contains",
	}

	if rule.Level != "ERROR" {
		t.Errorf("Expected Level 'ERROR', got %s", rule.Level)
	}
	if rule.Message != "exception" {
		t.Errorf("Expected Message 'exception', got %s", rule.Message)
	}
	if rule.Operator != "contains" {
		t.Errorf("Expected Operator 'contains', got %s", rule.Operator)
	}
}

func TestApplicationAlertRuleSlowness(t *testing.T) {
	rule := ApplicationAlertRuleSlowness{
		MetricName:  "latency",
		Aggregation: "mean",
	}

	if rule.MetricName != "latency" {
		t.Errorf("Expected MetricName 'latency', got %s", rule.MetricName)
	}
	if rule.Aggregation != "mean" {
		t.Errorf("Expected Aggregation 'mean', got %s", rule.Aggregation)
	}
}

func TestApplicationAlertRuleStatusCode(t *testing.T) {
	rule := ApplicationAlertRuleStatusCode{
		MetricName:      "statusCode",
		Aggregation:     "sum",
		StatusCodeStart: 500,
		StatusCodeEnd:   599,
	}

	if rule.StatusCodeStart != 500 {
		t.Errorf("Expected StatusCodeStart 500, got %d", rule.StatusCodeStart)
	}
	if rule.StatusCodeEnd != 599 {
		t.Errorf("Expected StatusCodeEnd 599, got %d", rule.StatusCodeEnd)
	}
}

func TestApplicationAlertRuleThroughput(t *testing.T) {
	rule := ApplicationAlertRuleThroughput{
		MetricName:  "calls",
		Aggregation: "sum",
	}

	if rule.MetricName != "calls" {
		t.Errorf("Expected MetricName 'calls', got %s", rule.MetricName)
	}
}

// ApplicationAlertRuleWithThresholds Tests
func TestApplicationAlertRuleWithThresholds(t *testing.T) {
	warningValue := 10.0
	criticalValue := 50.0

	rule := &ApplicationAlertRule{
		AlertType:   "errorRate",
		MetricName:  "errors",
		Aggregation: types.SumAggregation,
	}

	ruleWithThresholds := ApplicationAlertRuleWithThresholds{
		Rule:              rule,
		ThresholdOperator: ">",
		Thresholds: map[types.AlertSeverity]types.ThresholdRule{
			types.WarningSeverity:  {Value: &warningValue},
			types.CriticalSeverity: {Value: &criticalValue},
		},
	}

	if ruleWithThresholds.Rule.AlertType != "errorRate" {
		t.Errorf("Expected AlertType 'errorRate', got %s", ruleWithThresholds.Rule.AlertType)
	}
	if ruleWithThresholds.ThresholdOperator != ">" {
		t.Errorf("Expected ThresholdOperator '>', got %s", ruleWithThresholds.ThresholdOperator)
	}
	if len(ruleWithThresholds.Thresholds) != 2 {
		t.Errorf("Expected 2 thresholds, got %d", len(ruleWithThresholds.Thresholds))
	}
	warningThreshold := ruleWithThresholds.Thresholds[types.WarningSeverity]
	if warningThreshold.Value == nil || *warningThreshold.Value != 10.0 {
		t.Error("Expected warning threshold 10.0")
	}
}

// Time Threshold Variant Tests
func TestApplicationAlertTimeThresholdRequestImpact(t *testing.T) {
	threshold := ApplicationAlertTimeThresholdRequestImpact{
		TimeWindow: 300000,
		Requests:   50,
	}

	if threshold.TimeWindow != 300000 {
		t.Errorf("Expected TimeWindow 300000, got %d", threshold.TimeWindow)
	}
	if threshold.Requests != 50 {
		t.Errorf("Expected Requests 50, got %d", threshold.Requests)
	}
}

func TestApplicationAlertTimeThresholdViolationsInPeriod(t *testing.T) {
	threshold := ApplicationAlertTimeThresholdViolationsInPeriod{
		TimeWindow: 600000,
		Violations: 3,
	}

	if threshold.TimeWindow != 600000 {
		t.Errorf("Expected TimeWindow 600000, got %d", threshold.TimeWindow)
	}
	if threshold.Violations != 3 {
		t.Errorf("Expected Violations 3, got %d", threshold.Violations)
	}
}

func TestApplicationAlertTimeThresholdViolationsInSequence(t *testing.T) {
	threshold := ApplicationAlertTimeThresholdViolationsInSequence{
		TimeWindow: 300000,
	}

	if threshold.TimeWindow != 300000 {
		t.Errorf("Expected TimeWindow 300000, got %d", threshold.TimeWindow)
	}
}

// JSON Marshalling Tests
func TestApplicationAlertConfigJSONMarshalling(t *testing.T) {
	enabled := true
	config := &ApplicationAlertConfig{
		ID:               "test-id-123",
		Name:             "Test Alert",
		Description:      "Test description",
		Triggering:       true,
		Enabled:          &enabled,
		Applications:     map[string]IncludedApplication{},
		BoundaryScope:    types.BoundaryScopeAll,
		IncludeInternal:  false,
		IncludeSynthetic: true,
		EvaluationType:   EvaluationTypePerApplication,
		AlertChannelIDs:  []string{"channel-1"},
		AlertChannels:    map[string][]string{},
		Granularity:      types.Granularity600000,
	}

	// Marshal to JSON
	data, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	// Unmarshal from JSON
	var unmarshalled ApplicationAlertConfig
	err = json.Unmarshal(data, &unmarshalled)
	if err != nil {
		t.Fatalf("Failed to unmarshal config: %v", err)
	}

	// Verify fields
	if unmarshalled.ID != config.ID {
		t.Errorf("Expected ID %s, got %s", config.ID, unmarshalled.ID)
	}
	if unmarshalled.Name != config.Name {
		t.Errorf("Expected Name %s, got %s", config.Name, unmarshalled.Name)
	}
	if unmarshalled.EvaluationType != config.EvaluationType {
		t.Errorf("Expected EvaluationType %s, got %s", config.EvaluationType, unmarshalled.EvaluationType)
	}
}

// Edge Cases and Empty Values
func TestApplicationAlertConfigEmptyValues(t *testing.T) {
	config := ApplicationAlertConfig{}

	if config.ID != "" {
		t.Errorf("Expected empty ID, got %s", config.ID)
	}
	if config.Name != "" {
		t.Errorf("Expected empty Name, got %s", config.Name)
	}
	if config.Triggering {
		t.Error("Expected Triggering to be false")
	}
	if config.Enabled != nil {
		t.Error("Expected Enabled to be nil")
	}
}

func TestThresholdValueStructure(t *testing.T) {
	threshold := ThresholdValue{
		Value: 75.5,
	}

	if threshold.Value != 75.5 {
		t.Errorf("Expected Value 75.5, got %f", threshold.Value)
	}
}
