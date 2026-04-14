package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
	"github.com/instana/instana-go-client/shared/types"
)

func TestAlertsResourcePath(t *testing.T) {
	expected := "/api/events/settings/alerts"
	if AlertsResourcePath != expected {
		t.Errorf("Expected AlertsResourcePath to be %s, got %s", expected, AlertsResourcePath)
	}
}

func TestAlertEventTypeConstants(t *testing.T) {
	tests := []struct {
		name     string
		value    AlertEventType
		expected string
	}{
		{"IncidentAlertEventType", IncidentAlertEventType, "incident"},
		{"CriticalAlertEventType", CriticalAlertEventType, "critical"},
		{"WarningAlertEventType", WarningAlertEventType, "warning"},
		{"ChangeAlertEventType", ChangeAlertEventType, "change"},
		{"OnlineAlertEventType", OnlineAlertEventType, "online"},
		{"OfflineAlertEventType", OfflineAlertEventType, "offline"},
		{"NoneAlertEventType", NoneAlertEventType, "none"},
		{"AgentMonitoringIssueEventType", AgentMonitoringIssueEventType, "agent_monitoring_issue"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.value) != tt.expected {
				t.Errorf("Expected %s to be %s, got %s", tt.name, tt.expected, string(tt.value))
			}
		})
	}
}

func TestSupportedAlertEventTypes(t *testing.T) {
	expected := 8
	if len(SupportedAlertEventTypes) != expected {
		t.Errorf("Expected %d supported alert event types, got %d", expected, len(SupportedAlertEventTypes))
	}

	// Verify all expected types are present
	expectedTypes := map[AlertEventType]bool{
		IncidentAlertEventType:        true,
		CriticalAlertEventType:        true,
		WarningAlertEventType:         true,
		ChangeAlertEventType:          true,
		OnlineAlertEventType:          true,
		OfflineAlertEventType:         true,
		NoneAlertEventType:            true,
		AgentMonitoringIssueEventType: true,
	}

	for _, eventType := range SupportedAlertEventTypes {
		if !expectedTypes[eventType] {
			t.Errorf("Unexpected event type in SupportedAlertEventTypes: %s", eventType)
		}
	}
}

func TestAlertingConfigurationGetIDForResourcePath(t *testing.T) {
	testID := "test-alert-config-123"
	config := &AlertingConfiguration{
		ID:        testID,
		AlertName: "Test Alert",
	}

	result := config.GetIDForResourcePath()
	if result != testID {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", testID, result)
	}
}

func TestAlertingConfigurationCustomPayloadFields(t *testing.T) {
	fields := []types.CustomPayloadField[any]{
		{Key: "field1", Value: "value1"},
		{Key: "field2", Value: "value2"},
	}

	config := &AlertingConfiguration{
		ID:                    "test-id",
		AlertName:             "Test Alert",
		CustomerPayloadFields: fields,
	}

	// Test GetCustomerPayloadFields
	retrievedFields := config.GetCustomerPayloadFields()
	if len(retrievedFields) != len(fields) {
		t.Errorf("Expected %d custom payload fields, got %d", len(fields), len(retrievedFields))
	}

	// Test SetCustomerPayloadFields
	newFields := []types.CustomPayloadField[any]{
		{Key: "newField", Value: "newValue"},
	}
	config.SetCustomerPayloadFields(newFields)

	if len(config.CustomerPayloadFields) != 1 {
		t.Errorf("Expected 1 custom payload field after set, got %d", len(config.CustomerPayloadFields))
	}
	if config.CustomerPayloadFields[0].Key != "newField" {
		t.Errorf("Expected custom payload field key 'newField', got %s", config.CustomerPayloadFields[0].Key)
	}
}

func TestEventFilteringConfiguration(t *testing.T) {
	query := "test query"
	config := EventFilteringConfiguration{
		Query:                     &query,
		RuleIDs:                   []string{"rule1", "rule2"},
		EventTypes:                []AlertEventType{IncidentAlertEventType, CriticalAlertEventType},
		ApplicationAlertConfigIds: []string{"app1", "app2"},
	}

	if config.Query == nil || *config.Query != query {
		t.Error("Query not set correctly")
	}
	if len(config.RuleIDs) != 2 {
		t.Errorf("Expected 2 rule IDs, got %d", len(config.RuleIDs))
	}
	if len(config.EventTypes) != 2 {
		t.Errorf("Expected 2 event types, got %d", len(config.EventTypes))
	}
	if len(config.ApplicationAlertConfigIds) != 2 {
		t.Errorf("Expected 2 application alert config IDs, got %d", len(config.ApplicationAlertConfigIds))
	}
}

func TestAlertingConfigurationStructure(t *testing.T) {
	query := "entity.type:jvm"
	config := AlertingConfiguration{
		ID:             "alert-123",
		AlertName:      "Test Alert Configuration",
		IntegrationIDs: []string{"integration1", "integration2"},
		EventFilteringConfiguration: EventFilteringConfiguration{
			Query:      &query,
			RuleIDs:    []string{"rule1"},
			EventTypes: []AlertEventType{IncidentAlertEventType},
		},
		CustomerPayloadFields: []types.CustomPayloadField[any]{
			{Key: "severity", Value: "high"},
		},
	}

	if config.ID != "alert-123" {
		t.Errorf("Expected ID 'alert-123', got %s", config.ID)
	}
	if config.AlertName != "Test Alert Configuration" {
		t.Errorf("Expected AlertName 'Test Alert Configuration', got %s", config.AlertName)
	}
	if len(config.IntegrationIDs) != 2 {
		t.Errorf("Expected 2 integration IDs, got %d", len(config.IntegrationIDs))
	}
	if config.EventFilteringConfiguration.Query == nil {
		t.Error("EventFilteringConfiguration.Query should not be nil")
	}
}
