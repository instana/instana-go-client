package api

import (
	"testing"
)

func TestWebsiteAlertConfigResourcePath(t *testing.T) {
	expected := "/api/events/settings/website-alert-configs"
	if WebsiteAlertConfigResourcePath != expected {
		t.Errorf("Expected WebsiteAlertConfigResourcePath to be %s, got %s", expected, WebsiteAlertConfigResourcePath)
	}
}

func TestWebsiteAlertConfigGetIDForResourcePath(t *testing.T) {
	id := "test-website-alert-id"
	config := WebsiteAlertConfig{ID: id}

	result := config.GetIDForResourcePath()

	if result != id {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", id, result)
	}
}

func TestWebsiteAlertConfigStructure(t *testing.T) {
	id := "website-alert-id"
	name := "Test Website Alert"
	description := "Test website description"
	websiteID := "website-123"
	triggering := true

	config := WebsiteAlertConfig{
		ID:          id,
		Name:        name,
		Description: description,
		WebsiteID:   websiteID,
		Triggering:  triggering,
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
	if config.WebsiteID != websiteID {
		t.Errorf("Expected WebsiteID to be %s, got %s", websiteID, config.WebsiteID)
	}
	if config.Triggering != triggering {
		t.Errorf("Expected Triggering to be %v, got %v", triggering, config.Triggering)
	}
}

// Made with Bob
