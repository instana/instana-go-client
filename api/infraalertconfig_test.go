package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
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
