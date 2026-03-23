package api

import (
	"testing"
)

func TestMobileAlertConfigResourcePath(t *testing.T) {
	expected := "/api/events/settings/mobile-app-alert-configs"
	if MobileAlertConfigResourcePath != expected {
		t.Errorf("Expected MobileAlertConfigResourcePath to be %s, got %s", expected, MobileAlertConfigResourcePath)
	}
}

func TestMobileAlertConfigGetIDForResourcePath(t *testing.T) {
	id := "test-mobile-id"
	config := MobileAlertConfig{ID: id}

	result := config.GetIDForResourcePath()

	if result != id {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", id, result)
	}
}

func TestMobileAlertConfigStructure(t *testing.T) {
	id := "mobile-alert-id"
	name := "Test Mobile Alert"
	description := "Test mobile description"
	mobileAppID := "app-123"
	triggering := true

	config := MobileAlertConfig{
		ID:          id,
		Name:        name,
		Description: description,
		MobileAppID: mobileAppID,
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
	if config.MobileAppID != mobileAppID {
		t.Errorf("Expected MobileAppID to be %s, got %s", mobileAppID, config.MobileAppID)
	}
	if config.Triggering != triggering {
		t.Errorf("Expected Triggering to be %v, got %v", triggering, config.Triggering)
	}
}

// Made with Bob
