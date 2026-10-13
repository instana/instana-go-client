package api_test

import (
	"encoding/json"
	"testing"

	. "github.com/instana/instana-go-client/api"
	"github.com/instana/instana-go-client/shared/types"
)

func TestSyntheticAlertConfigsResourcePath(t *testing.T) {
	expected := "/api/events/settings/global-alert-configs/synthetics"
	if SyntheticAlertConfigsResourcePath != expected {
		t.Errorf("Expected SyntheticAlertConfigsResourcePath to be %s, got %s", expected, SyntheticAlertConfigsResourcePath)
	}
}

func TestSyntheticAlertConfigGetIDForResourcePath(t *testing.T) {
	id := "test-synthetic-alert-id"
	config := SyntheticAlertConfig{ID: id}

	result := config.GetIDForResourcePath()

	if result != id {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", id, result)
	}
}

func TestSyntheticAlertConfigStructure(t *testing.T) {
	id := "synthetic-alert-id"
	name := "Test Synthetic Alert"
	description := "Test synthetic description"
	severity := 5
	enabled := true
	includeOnDemand := true
	evaluationType := "PER_SYNTHETIC_TEST"
	affectedLocationsCount := 2

	config := SyntheticAlertConfig{
		ID:                     id,
		Name:                   name,
		Description:            description,
		Severity:               severity,
		Enabled:                enabled,
		IncludeOnDemand:        includeOnDemand,
		EvaluationType:         evaluationType,
		AffectedLocationsCount: affectedLocationsCount,
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
	if config.Enabled != enabled {
		t.Errorf("Expected Enabled to be %v, got %v", enabled, config.Enabled)
	}
	if config.IncludeOnDemand != includeOnDemand {
		t.Errorf("Expected IncludeOnDemand to be %v, got %v", includeOnDemand, config.IncludeOnDemand)
	}
	if config.EvaluationType != evaluationType {
		t.Errorf("Expected EvaluationType to be %s, got %s", evaluationType, config.EvaluationType)
	}
	if config.AffectedLocationsCount != affectedLocationsCount {
		t.Errorf("Expected AffectedLocationsCount to be %d, got %d", affectedLocationsCount, config.AffectedLocationsCount)
	}
}

func TestSyntheticAlertConfigCustomPayloadFields(t *testing.T) {
	config := &SyntheticAlertConfig{
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

func TestSyntheticAlertConfigJSONSerialization(t *testing.T) {
	jsonResponse := `{
		"name": "tf-syn-test",
		"description": "test desc",
		"syntheticTestIds": [
			"3tjCRCXPG2P0ZIeVKxVw"
		],
		"severity": 5,
		"tagFilterExpression": {
			"type": "EXPRESSION",
			"logicalOperator": "AND",
			"elements": []
		},
		"rule": {
			"alertType": "failure",
			"metricName": "status",
			"aggregation": "SUM"
		},
		"alertChannelIds": [
			"uMiJMd1wBDpEqnO3"
		],
		"gracePeriod": 60000,
		"timeThreshold": {
			"type": "violationsInSequence",
			"violationsCount": 3
		},
		"customPayloadFields": [
			{
				"type": "staticString",
				"key": "test",
				"value": "test"
			}
		],
		"includeOnDemand": true,
		"affectedLocationsCount": 2,
		"evaluationType": "PER_SYNTHETIC_TEST",
		"id": "4dFiU8y9TZWUHZf-EdWZiA",
		"enabled": false
	}`

	var config SyntheticAlertConfig
	if err := json.Unmarshal([]byte(jsonResponse), &config); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if config.ID != "4dFiU8y9TZWUHZf-EdWZiA" {
		t.Errorf("Expected ID '4dFiU8y9TZWUHZf-EdWZiA', got '%s'", config.ID)
	}
	if config.Name != "tf-syn-test" {
		t.Errorf("Expected Name 'tf-syn-test', got '%s'", config.Name)
	}
	if config.IncludeOnDemand != true {
		t.Errorf("Expected IncludeOnDemand true, got %v", config.IncludeOnDemand)
	}
	if config.EvaluationType != "PER_SYNTHETIC_TEST" {
		t.Errorf("Expected EvaluationType 'PER_SYNTHETIC_TEST', got '%s'", config.EvaluationType)
	}
	if config.AffectedLocationsCount != 2 {
		t.Errorf("Expected AffectedLocationsCount 2, got %d", config.AffectedLocationsCount)
	}
	if config.Enabled != false {
		t.Errorf("Expected Enabled false, got %v", config.Enabled)
	}

	marshalled, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	var roundtrip SyntheticAlertConfig
	if err := json.Unmarshal(marshalled, &roundtrip); err != nil {
		t.Fatalf("Failed to unmarshal roundtrip JSON: %v", err)
	}

	if roundtrip.IncludeOnDemand != true {
		t.Errorf("Roundtrip IncludeOnDemand mismatch: expected true, got %v", roundtrip.IncludeOnDemand)
	}
	if roundtrip.EvaluationType != "PER_SYNTHETIC_TEST" {
		t.Errorf("Roundtrip EvaluationType mismatch: expected 'PER_SYNTHETIC_TEST', got '%s'", roundtrip.EvaluationType)
	}
	if roundtrip.AffectedLocationsCount != 2 {
		t.Errorf("Roundtrip AffectedLocationsCount mismatch: expected 2, got %d", roundtrip.AffectedLocationsCount)
	}
	if roundtrip.Enabled != false {
		t.Errorf("Roundtrip Enabled mismatch: expected false, got %v", roundtrip.Enabled)
	}
}
