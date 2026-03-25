package api_test

import (
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

	config := SyntheticAlertConfig{
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
