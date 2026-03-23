package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
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
