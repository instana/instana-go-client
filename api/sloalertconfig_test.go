package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
	"github.com/instana/instana-go-client/shared/types"
)

func TestSloAlertConfigResourcePath(t *testing.T) {
	expected := "/api/events/settings/global-alert-configs/service-levels"
	if SloAlertConfigResourcePath != expected {
		t.Errorf("Expected SloAlertConfigResourcePath to be %s, got %s", expected, SloAlertConfigResourcePath)
	}
}

func TestSloAlertConfigGetIDForResourcePath(t *testing.T) {
	id := "test-slo-alert-id"
	config := SloAlertConfig{ID: id}

	result := config.GetIDForResourcePath()

	if result != id {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", id, result)
	}
}

func TestSloAlertConfigStructure(t *testing.T) {
	id := "slo-alert-id"
	name := "Test SLO Alert"
	description := "Test SLO description"
	severity := 5

	config := SloAlertConfig{
		ID:          id,
		Name:        name,
		Description: description,
		Severity:    severity,
	}

	if config.ID != id {
		t.Errorf("Expected ID to be %s, got %s", id, config.ID)
	}
	if config.Name != name {
		t.Errorf("Expected Name to be %s, got %s", name, config.Name)
	}
	if config.Description != description {
		t.Errorf("Expected Description to be %s, got %s", description, config.Description)
	}
	if config.Severity != severity {
		t.Errorf("Expected Severity to be %d, got %d", severity, config.Severity)
	}
}

func TestSloAlertConfigCustomPayloadFields(t *testing.T) {
	config := &SloAlertConfig{
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

func TestSloAlertRuleAlertTypeApdexConstant(t *testing.T) {
	if SloAlertRuleAlertTypeApdex != "APDEX" {
		t.Errorf("Expected SloAlertRuleAlertTypeApdex to be \"APDEX\", got %s", SloAlertRuleAlertTypeApdex)
	}
}

func TestSloAlertMetricScoreConstant(t *testing.T) {
	if SloAlertMetricScore != "SCORE" {
		t.Errorf("Expected SloAlertMetricScore to be \"SCORE\", got %s", SloAlertMetricScore)
	}
}

func TestSloAlertConfigWithApdexIds(t *testing.T) {
	config := SloAlertConfig{
		ID:          "apdex-alert-id",
		Name:        "Test Apdex Alert",
		Description: "Test Apdex description",
		Severity:    10,
		Rule: SloAlertRule{
			AlertType: SloAlertRuleAlertTypeApdex,
			Metric:    SloAlertMetricScore,
		},
		ApdexIds:        []string{"apdex-id-1", "apdex-id-2"},
		SloIds:          []string{},
		AlertChannelIds: []string{"channel-1"},
	}

	if config.Rule.AlertType != SloAlertRuleAlertTypeApdex {
		t.Errorf("Expected Rule.AlertType to be %s, got %s", SloAlertRuleAlertTypeApdex, config.Rule.AlertType)
	}
	if config.Rule.Metric != SloAlertMetricScore {
		t.Errorf("Expected Rule.Metric to be %s, got %s", SloAlertMetricScore, config.Rule.Metric)
	}
	if len(config.ApdexIds) != 2 {
		t.Errorf("Expected 2 ApdexIds, got %d", len(config.ApdexIds))
	}
	if config.ApdexIds[0] != "apdex-id-1" {
		t.Errorf("Expected first ApdexId to be \"apdex-id-1\", got %s", config.ApdexIds[0])
	}
	if len(config.SloIds) != 0 {
		t.Errorf("Expected SloIds to be empty for Apdex alert, got %d entries", len(config.SloIds))
	}
}

func TestSloAlertConfigBackwardsCompatibility(t *testing.T) {
	config := SloAlertConfig{
		ID:          "slo-alert-id",
		Name:        "Test SLO Alert",
		Description: "Test SLO description",
		Severity:    5,
		Rule: SloAlertRule{
			AlertType: "SERVICE_LEVELS_OBJECTIVE",
			Metric:    "STATUS",
		},
		SloIds:   []string{"slo-id-1"},
		ApdexIds: []string{},
	}

	if config.Rule.AlertType != "SERVICE_LEVELS_OBJECTIVE" {
		t.Errorf("Expected Rule.AlertType to be SERVICE_LEVELS_OBJECTIVE, got %s", config.Rule.AlertType)
	}
	if len(config.SloIds) != 1 || config.SloIds[0] != "slo-id-1" {
		t.Errorf("Expected SloIds to contain slo-id-1, got %v", config.SloIds)
	}
	if len(config.ApdexIds) != 0 {
		t.Errorf("Expected ApdexIds to be empty for SLO alert, got %d entries", len(config.ApdexIds))
	}
}
