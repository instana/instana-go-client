package api

import (
	"testing"
)

func TestSloConfigResourcePath(t *testing.T) {
	expected := "/api/settings/slo"
	if SloConfigResourcePath != expected {
		t.Errorf("Expected SloConfigResourcePath to be %s, got %s", expected, SloConfigResourcePath)
	}
}

func TestSloConfigGetIDForResourcePath(t *testing.T) {
	id := "test-slo-id"
	config := SloConfig{ID: id}

	result := config.GetIDForResourcePath()

	if result != id {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", id, result)
	}
}

func TestSloConfigStructure(t *testing.T) {
	id := "slo-id"
	name := "Test SLO"

	config := SloConfig{
		ID:   id,
		Name: name,
	}

	if config.ID != id {
		t.Errorf("Expected ID to be %s, got %s", id, config.ID)
	}
	if config.Name != name {
		t.Errorf("Expected Name to be %s, got %s", name, config.Name)
	}
}

// Made with Bob
