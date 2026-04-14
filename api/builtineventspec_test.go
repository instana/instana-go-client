package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
)

func TestBuiltinEventSpecificationResourcePath(t *testing.T) {
	expected := "/api/events/settings/event-specifications/built-in"
	if BuiltinEventSpecificationResourcePath != expected {
		t.Errorf("Expected BuiltinEventSpecificationResourcePath to be %s, got %s", expected, BuiltinEventSpecificationResourcePath)
	}
}

func TestBuiltinEventSpecificationGetIDForResourcePath(t *testing.T) {
	testID := "test-spec-123"
	description := "Test description"
	spec := &BuiltinEventSpecification{
		ID:            testID,
		ShortPluginID: "plugin-1",
		Name:          "Test Spec",
		Description:   &description,
		Severity:      5,
		Triggering:    true,
		Enabled:       true,
	}

	result := spec.GetIDForResourcePath()
	if result != testID {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", testID, result)
	}
}

func TestBuiltinEventSpecificationStructure(t *testing.T) {
	description := "High CPU usage detected"
	spec := BuiltinEventSpecification{
		ID:            "spec-456",
		ShortPluginID: "cpu",
		Name:          "CPU Alert",
		Description:   &description,
		Severity:      10,
		Triggering:    true,
		Enabled:       false,
	}

	if spec.ID != "spec-456" {
		t.Errorf("Expected ID 'spec-456', got %s", spec.ID)
	}
	if spec.ShortPluginID != "cpu" {
		t.Errorf("Expected ShortPluginID 'cpu', got %s", spec.ShortPluginID)
	}
	if spec.Name != "CPU Alert" {
		t.Errorf("Expected Name 'CPU Alert', got %s", spec.Name)
	}
	if spec.Description == nil || *spec.Description != description {
		t.Error("Description not set correctly")
	}
	if spec.Severity != 10 {
		t.Errorf("Expected Severity 10, got %d", spec.Severity)
	}
	if !spec.Triggering {
		t.Error("Expected Triggering to be true")
	}
	if spec.Enabled {
		t.Error("Expected Enabled to be false")
	}
}
