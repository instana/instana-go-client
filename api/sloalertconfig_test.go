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
