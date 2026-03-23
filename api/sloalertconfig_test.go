package api

import (
	"testing"
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

// Made with Bob
