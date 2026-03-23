package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
)

func TestMaintenanceWindowResourcePath(t *testing.T) {
	expected := "/api/settings/v2/maintenance"
	if MaintenanceWindowConfigResourcePath != expected {
		t.Errorf("Expected MaintenanceWindowConfigResourcePath to be %s, got %s", expected, MaintenanceWindowConfigResourcePath)
	}
}

func TestMaintenanceWindowGetIDForResourcePath(t *testing.T) {
	testID := "test-mw-123"
	mw := &MaintenanceWindow{
		ID:   testID,
		Name: "Test Maintenance",
	}

	result := mw.GetIDForResourcePath()
	if result != testID {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", testID, result)
	}
}

func TestMaintenanceWindowStructure(t *testing.T) {
	mw := MaintenanceWindow{
		ID:    "mw-456",
		Name:  "Weekly Maintenance",
		Query: "entity.type:host",
	}

	if mw.ID != "mw-456" {
		t.Errorf("Expected ID 'mw-456', got %s", mw.ID)
	}
	if mw.Name != "Weekly Maintenance" {
		t.Errorf("Expected Name 'Weekly Maintenance', got %s", mw.Name)
	}
	if mw.Query != "entity.type:host" {
		t.Errorf("Expected Query 'entity.type:host', got %s", mw.Query)
	}
}
