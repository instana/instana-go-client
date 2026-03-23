package api

import (
	"testing"
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

// Made with Bob
