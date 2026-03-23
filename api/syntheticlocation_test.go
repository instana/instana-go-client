package api

import (
	"testing"
)

func TestSyntheticLocationResourcePath(t *testing.T) {
	expected := "/api/synthetics/settings/locations"
	if SyntheticLocationResourcePath != expected {
		t.Errorf("Expected SyntheticLocationResourcePath to be %s, got %s", expected, SyntheticLocationResourcePath)
	}
}

func TestSyntheticLocationGetIDForResourcePath(t *testing.T) {
	id := "test-location-id"
	location := SyntheticLocation{ID: id}

	result := location.GetIDForResourcePath()

	if result != id {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", id, result)
	}
}

func TestSyntheticLocationStructure(t *testing.T) {
	id := "location-id"
	label := "Test Location"
	description := "Test location description"
	locationType := "PoP"

	location := SyntheticLocation{
		ID:           id,
		Label:        label,
		Description:  description,
		LocationType: locationType,
	}

	if location.ID != id {
		t.Errorf("Expected ID to be %s, got %s", id, location.ID)
	}
	if location.Label != label {
		t.Errorf("Expected Label to be %s, got %s", label, location.Label)
	}
	if location.Description != description {
		t.Errorf("Expected Description to be %s, got %s", description, location.Description)
	}
	if location.LocationType != locationType {
		t.Errorf("Expected LocationType to be %s, got %s", locationType, location.LocationType)
	}
}

// Made with Bob
