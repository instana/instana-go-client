package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
	model "github.com/instana/instana-go-client/shared/types"
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

func TestGetCustomerPayloadFields(t *testing.T) {
	id := "website-alert-id"
	name := "Test Website Alert"
	description := "Test website description"
	websiteID := "website-123"
	triggering := true

	// Create custom payload fields
	customPayloadFields := []model.CustomPayloadField[any]{
		{
			Type:  model.StaticStringCustomPayloadType,
			Key:   "testKey1",
			Value: "testValue1",
		},
		{
			Type:  model.StaticStringCustomPayloadType,
			Key:   "testKey2",
			Value: "testValue2",
		},
	}

	config := WebsiteAlertConfig{
		ID:                    id,
		Name:                  name,
		Description:           description,
		WebsiteID:             websiteID,
		Triggering:            triggering,
		CustomerPayloadFields: customPayloadFields,
	}

	response := config.GetCustomerPayloadFields()
	if response == nil {
		t.Errorf("Expected custom payload fields. Got nil")
	}
	if len(response) != 2 {
		t.Errorf("Expected 2 custom payload fields, got %d", len(response))
	}
	if response[0].Key != "testKey1" {
		t.Errorf("Expected first field key to be 'testKey1', got %s", response[0].Key)
	}
	if response[0].Value != "testValue1" {
		t.Errorf("Expected first field value to be 'testValue1', got %v", response[0].Value)
	}
}

func TestSetCustomerPayloadFields(t *testing.T) {
	config := WebsiteAlertConfig{
		ID:   "test-id",
		Name: "Test Config",
	}

	// Create custom payload fields
	customPayloadFields := []model.CustomPayloadField[any]{
		{
			Type:  model.StaticStringCustomPayloadType,
			Key:   "newKey",
			Value: "newValue",
		},
	}

	config.SetCustomerPayloadFields(customPayloadFields)

	result := config.GetCustomerPayloadFields()
	if result == nil {
		t.Errorf("Expected custom payload fields after setting. Got nil")
	}
	if len(result) != 1 {
		t.Errorf("Expected 1 custom payload field, got %d", len(result))
	}
	if result[0].Key != "newKey" {
		t.Errorf("Expected field key to be 'newKey', got %s", result[0].Key)
	}
}

func TestWebsiteImpactMeasurementMethodsToStringSlice(t *testing.T) {
	methods := WebsiteImpactMeasurementMethods{
		model.WebsiteImpactMeasurementMethodAggregated,
		model.WebsiteImpactMeasurementMethodPerWindow,
	}

	result := methods.ToStringSlice()

	if len(result) != 2 {
		t.Errorf("Expected 2 methods in string slice, got %d", len(result))
	}
	if result[0] != string(model.WebsiteImpactMeasurementMethodAggregated) {
		t.Errorf("Expected first method to be '%s', got '%s'", model.WebsiteImpactMeasurementMethodAggregated, result[0])
	}
	if result[1] != string(model.WebsiteImpactMeasurementMethodPerWindow) {
		t.Errorf("Expected second method to be '%s', got '%s'", model.WebsiteImpactMeasurementMethodPerWindow, result[1])
	}
}

func TestWebsiteImpactMeasurementMethodsToStringSliceEmpty(t *testing.T) {
	methods := WebsiteImpactMeasurementMethods{}

	result := methods.ToStringSlice()

	if len(result) != 0 {
		t.Errorf("Expected empty string slice, got %d elements", len(result))
	}
}
