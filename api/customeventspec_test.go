package api_test

import (
	"encoding/json"
	"testing"

	. "github.com/instana/instana-go-client/api"
	tagfilter "github.com/instana/instana-go-client/shared/tagfilter"
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

func TestCustomEventSpecificationStructureAndJSON(t *testing.T) {
	query := "entity.type:jvm"
	desc := "JVM CPU spike"
	expirationTime := 300
	offlineDuration := 60
	closeAfter := 120

	tagName := "agent.tag"
	tagFilter := &tagfilter.TagFilter{
		Type:  tagfilter.TagFilterType,
		Name:  &tagName,
		Value: "prod",
	}

	spec := CustomEventSpecification{
		ID:                  "spec-123",
		Name:                "High CPU Usage",
		EntityType:          "jvm",
		Query:               &query,
		Triggering:          true,
		Description:         &desc,
		ExpirationTime:      &expirationTime,
		Enabled:             true,
		RuleLogicalOperator: "AND",
		Rules: []RuleSpecification{
			{
				DType:           ThresholdRuleType,
				Severity:        1,
				OfflineDuration: &offlineDuration,
				CloseAfter:      &closeAfter,
				TagFilter:       tagFilter,
			},
		},
		TransientEventEnabled:    true,
		TransientEventThreshold:  60000,
		TransientEventAlertMuted: true,
	}

	data, err := json.Marshal(spec)
	if err != nil {
		t.Fatalf("Failed to marshal CustomEventSpecification: %v", err)
	}

	var unmarshalled CustomEventSpecification
	if err := json.Unmarshal(data, &unmarshalled); err != nil {
		t.Fatalf("Failed to unmarshal CustomEventSpecification: %v", err)
	}

	if unmarshalled.ID != spec.ID {
		t.Errorf("Expected ID %s, got %s", spec.ID, unmarshalled.ID)
	}
	if unmarshalled.TransientEventEnabled != true {
		t.Errorf("Expected TransientEventEnabled to be true, got %v", unmarshalled.TransientEventEnabled)
	}
	if unmarshalled.TransientEventThreshold != 60000 {
		t.Errorf("Expected TransientEventThreshold to be 60000, got %d", unmarshalled.TransientEventThreshold)
	}
	if unmarshalled.TransientEventAlertMuted != true {
		t.Errorf("Expected TransientEventAlertMuted to be true, got %v", unmarshalled.TransientEventAlertMuted)
	}
	if len(unmarshalled.Rules) != 1 {
		t.Fatalf("Expected 1 rule, got %d", len(unmarshalled.Rules))
	}
	if unmarshalled.Rules[0].TagFilter == nil || unmarshalled.Rules[0].TagFilter.Name == nil || *unmarshalled.Rules[0].TagFilter.Name != "agent.tag" {
		t.Errorf("Expected TagFilter with Name 'agent.tag', got %+v", unmarshalled.Rules[0].TagFilter)
	}
}
