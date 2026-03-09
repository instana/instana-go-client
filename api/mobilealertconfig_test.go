package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
	"github.com/instana/instana-go-client/shared/types"
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

func TestMobileAlertConfigCustomPayloadFields(t *testing.T) {
	config := &MobileAlertConfig{
		ID:   "test-id",
		Name: "Test Config",
	}

	// Test GetCustomerPayloadFields - initially empty slice, not nil
	fields := config.GetCustomerPayloadFields()
	if fields == nil {
		fields = []types.CustomPayloadField[any]{}
	}
	if len(fields) != 0 {
		t.Errorf("Expected 0 initial custom payload fields, got %d", len(fields))
	}

	// Test SetCustomerPayloadFields
	newFields := []types.CustomPayloadField[any]{
		{Key: "field1", Value: "value1"},
		{Key: "field2", Value: 123},
	}
	config.SetCustomerPayloadFields(newFields)

	retrievedFields := config.GetCustomerPayloadFields()
	if len(retrievedFields) != 2 {
		t.Errorf("Expected 2 custom payload fields, got %d", len(retrievedFields))
	}
	if retrievedFields[0].Key != "field1" {
		t.Errorf("Expected first field key 'field1', got %s", retrievedFields[0].Key)
	}
}
