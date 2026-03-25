package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
)

func TestCustomEventSpecResourcePath(t *testing.T) {
	expected := "/api/events/settings/event-specifications/custom"
	if CustomeventspecResourcePath != expected {
		t.Errorf("Expected CustomeventspecResourcePath to be %s, got %s", expected, CustomeventspecResourcePath)
	}
}

func TestCustomEventSpecRuleTypeConstants(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"SystemRuleType", SystemRuleType, "system"},
		{"ThresholdRuleType", ThresholdRuleType, "threshold"},
		{"EntityVerificationRuleType", EntityVerificationRuleType, "entity_verification"},
		{"EntityCountRuleType", EntityCountRuleType, "entity_count"},
		{"EntityCountVerificationRuleType", EntityCountVerificationRuleType, "entity_count_verification"},
		{"HostAvailabilityRuleType", HostAvailabilityRuleType, "host_availability"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.expected {
				t.Errorf("Expected %s to be %s, got %s", tt.name, tt.expected, tt.value)
			}
		})
	}
}

func TestGetIDForResourcePath(t *testing.T) {
	config := CustomEventSpecification{
		ID: "12345",
	}
	if config.GetIDForResourcePath() != "12345" {
		t.Errorf("Expected %s to be %s", config.ID, "12345")
	}

}
