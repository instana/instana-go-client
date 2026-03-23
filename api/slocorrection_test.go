package api

import (
	"testing"
)

func TestSloCorrectionConfigResourcePath(t *testing.T) {
	expected := "/api/settings/correction"
	if SloCorrectionConfigResourcePath != expected {
		t.Errorf("Expected SloCorrectionConfigResourcePath to be %s, got %s", expected, SloCorrectionConfigResourcePath)
	}
}

func TestSloCorrectionConfigGetIDForResourcePath(t *testing.T) {
	id := "test-correction-id"
	config := SloCorrectionConfig{ID: id}

	result := config.GetIDForResourcePath()

	if result != id {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", id, result)
	}
}

func TestSloCorrectionConfigStructure(t *testing.T) {
	id := "correction-id"
	name := "Test SLO Correction"
	description := "Test correction description"
	active := true

	config := SloCorrectionConfig{
		ID:          id,
		Name:        name,
		Description: description,
		Active:      active,
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
	if config.Active != active {
		t.Errorf("Expected Active to be %v, got %v", active, config.Active)
	}
}

// Made with Bob
