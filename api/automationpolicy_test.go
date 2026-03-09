package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
)

func TestAutomationPolicyResourcePath(t *testing.T) {
	expected := "/api/automation/policies"
	if AutomationPolicyResourcePath != expected {
		t.Errorf("Expected AutomationPolicyResourcePath to be %s, got %s", expected, AutomationPolicyResourcePath)
	}
}

func TestAutomationPolicyGetIDForResourcePath(t *testing.T) {
	testID := "test-policy-123"
	policy := &AutomationPolicy{
		ID:   testID,
		Name: "Test Policy",
	}

	result := policy.GetIDForResourcePath()
	if result != testID {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", testID, result)
	}
}

func TestAutomationPolicyStructure(t *testing.T) {
	policy := AutomationPolicy{
		ID:          "policy-123",
		Name:        "Alert Response Policy",
		Description: "Automated response to alerts",
		Tags:        []string{"production", "critical"},
		Trigger: Trigger{
			Id:          "trigger-1",
			Type:        "alert",
			Name:        "Alert Trigger",
			Description: "Triggered on alert",
		},
		TypeConfigurations: []TypeConfiguration{
			{
				Name: "config-1",
				Condition: &Condition{
					Query: "entity.type:host",
				},
				Runnable: Runnable{
					Id:   "runnable-1",
					Type: "action",
				},
			},
		},
	}

	if policy.ID != "policy-123" {
		t.Errorf("Expected ID 'policy-123', got %s", policy.ID)
	}
	if policy.Name != "Alert Response Policy" {
		t.Errorf("Expected Name 'Alert Response Policy', got %s", policy.Name)
	}
	if policy.Description != "Automated response to alerts" {
		t.Errorf("Expected Description 'Automated response to alerts', got %s", policy.Description)
	}
	if len(policy.TypeConfigurations) != 1 {
		t.Errorf("Expected 1 type configuration, got %d", len(policy.TypeConfigurations))
	}
}

func TestTriggerStructure(t *testing.T) {
	trigger := Trigger{
		Id:          "trigger-456",
		Type:        "scheduled",
		Name:        "Daily Backup",
		Description: "Run daily backup",
	}

	if trigger.Id != "trigger-456" {
		t.Errorf("Expected Id 'trigger-456', got %s", trigger.Id)
	}
	if trigger.Type != "scheduled" {
		t.Errorf("Expected Type 'scheduled', got %s", trigger.Type)
	}
	if trigger.Name != "Daily Backup" {
		t.Errorf("Expected Name 'Daily Backup', got %s", trigger.Name)
	}
}

func TestTypeConfigurationStructure(t *testing.T) {
	config := TypeConfiguration{
		Name: "host-config",
		Condition: &Condition{
			Query: "entity.type:host AND entity.zone:us-east",
		},
		Runnable: Runnable{
			Id:   "runnable-789",
			Type: "script",
		},
	}

	if config.Name != "host-config" {
		t.Errorf("Expected Name 'host-config', got %s", config.Name)
	}
	if config.Condition == nil {
		t.Error("Expected Condition to be non-nil")
	}
	if config.Condition.Query != "entity.type:host AND entity.zone:us-east" {
		t.Errorf("Expected Query 'entity.type:host AND entity.zone:us-east', got %s", config.Condition.Query)
	}
}

func TestConditionStructure(t *testing.T) {
	condition := Condition{
		Query: "entity.type:service",
	}

	if condition.Query != "entity.type:service" {
		t.Errorf("Expected Query 'entity.type:service', got %s", condition.Query)
	}
}

func TestRunnableStructure(t *testing.T) {
	runnable := Runnable{
		Id:   "runnable-abc",
		Type: "webhook",
		RunConfiguration: RunConfiguration{
			Actions: []AutomationActionPolicy{},
		},
	}

	if runnable.Id != "runnable-abc" {
		t.Errorf("Expected Id 'runnable-abc', got %s", runnable.Id)
	}
	if runnable.Type != "webhook" {
		t.Errorf("Expected Type 'webhook', got %s", runnable.Type)
	}
}

func TestAutomationActionPolicyStructure(t *testing.T) {
	actionPolicy := AutomationActionPolicy{
		Action: AutomationAction{
			ID:   "action-1",
			Name: "Restart Service",
		},
		AgentId: "agent-123",
	}

	if actionPolicy.Action.ID != "action-1" {
		t.Errorf("Expected Action ID 'action-1', got %s", actionPolicy.Action.ID)
	}
	if actionPolicy.AgentId != "agent-123" {
		t.Errorf("Expected AgentId 'agent-123', got %s", actionPolicy.AgentId)
	}
}

func TestActionStructure(t *testing.T) {
	action := Action{
		Id: "action-xyz",
	}

	if action.Id != "action-xyz" {
		t.Errorf("Expected Id 'action-xyz', got %s", action.Id)
	}
}

func TestInputParameterValueStructure(t *testing.T) {
	param := InputParameterValue{
		Name:  "hostname",
		Value: "server.example.com",
	}

	if param.Name != "hostname" {
		t.Errorf("Expected Name 'hostname', got %s", param.Name)
	}
	if param.Value != "server.example.com" {
		t.Errorf("Expected Value 'server.example.com', got %s", param.Value)
	}
}
