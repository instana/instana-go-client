package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
)

func TestCustomDashboardResourcePath(t *testing.T) {
	expected := "/api/custom-dashboard"
	if CustomDashboardsResourcePath != expected {
		t.Errorf("Expected CustomDashboardsResourcePath to be %s, got %s", expected, CustomDashboardsResourcePath)
	}
}

func TestCustomDashboardGetIDForResourcePath(t *testing.T) {
	testID := "test-dashboard-123"
	dashboard := &CustomDashboard{
		ID:    testID,
		Title: "Test Dashboard",
	}

	result := dashboard.GetIDForResourcePath()
	if result != testID {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", testID, result)
	}
}

func TestCustomDashboardStructure(t *testing.T) {
	dashboard := CustomDashboard{
		ID:    "dashboard-456",
		Title: "Production Dashboard",
	}

	if dashboard.ID != "dashboard-456" {
		t.Errorf("Expected ID 'dashboard-456', got %s", dashboard.ID)
	}
	if dashboard.Title != "Production Dashboard" {
		t.Errorf("Expected Title 'Production Dashboard', got %s", dashboard.Title)
	}
}
