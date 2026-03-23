package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
)

func TestAutomationActionResourcePath(t *testing.T) {
	expected := "/api/automation/actions"
	if AutomationActionResourcePath != expected {
		t.Errorf("Expected AutomationActionResourcePath to be %s, got %s", expected, AutomationActionResourcePath)
	}
}

func TestAutomationActionGetIDForResourcePath(t *testing.T) {
	testID := "test-action-123"
	action := &AutomationAction{
		ID:   testID,
		Name: "Test Action",
	}

	result := action.GetIDForResourcePath()
	if result != testID {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", testID, result)
	}
}

func TestAutomationActionStructure(t *testing.T) {
	action := AutomationAction{
		ID:          "action-123",
		Name:        "SSH Script Action",
		Description: "Execute SSH script",
		Type:        "script",
		Tags:        []string{"production", "deployment"},
		Timeout:     300,
		Fields: []Field{
			{Name: "script_ssh", Description: "script content", Value: "#!/bin/bash\necho 'test'"},
		},
		InputParameters: []Parameter{
			{Name: "host", Label: "Host", Type: "string", Required: true},
		},
	}

	if action.ID != "action-123" {
		t.Errorf("Expected ID 'action-123', got %s", action.ID)
	}
	if action.Name != "SSH Script Action" {
		t.Errorf("Expected Name 'SSH Script Action', got %s", action.Name)
	}
	if action.Description != "Execute SSH script" {
		t.Errorf("Expected Description 'Execute SSH script', got %s", action.Description)
	}
	if action.Type != "script" {
		t.Errorf("Expected Type 'script', got %s", action.Type)
	}
	if action.Timeout != 300 {
		t.Errorf("Expected Timeout 300, got %d", action.Timeout)
	}
	if len(action.Fields) != 1 {
		t.Errorf("Expected 1 field, got %d", len(action.Fields))
	}
	if len(action.InputParameters) != 1 {
		t.Errorf("Expected 1 input parameter, got %d", len(action.InputParameters))
	}
}

func TestParameterStructure(t *testing.T) {
	param := Parameter{
		Name:        "hostname",
		Label:       "Host Name",
		Description: "Target hostname",
		Type:        "string",
		Value:       "server.example.com",
		Required:    true,
		Hidden:      false,
		Secured:     false,
		ValueType:   "text",
	}

	if param.Name != "hostname" {
		t.Errorf("Expected Name 'hostname', got %s", param.Name)
	}
	if param.Label != "Host Name" {
		t.Errorf("Expected Label 'Host Name', got %s", param.Label)
	}
	if !param.Required {
		t.Error("Expected Required to be true")
	}
	if param.Hidden {
		t.Error("Expected Hidden to be false")
	}
	if param.Secured {
		t.Error("Expected Secured to be false")
	}
}

func TestFieldStructure(t *testing.T) {
	field := Field{
		Name:        "script_ssh",
		Description: "script content",
		Encoding:    "base64",
		Value:       "IyEvYmluL2Jhc2g=",
		Secured:     false,
	}

	if field.Name != "script_ssh" {
		t.Errorf("Expected Name 'script_ssh', got %s", field.Name)
	}
	if field.Description != "script content" {
		t.Errorf("Expected Description 'script content', got %s", field.Description)
	}
	if field.Encoding != "base64" {
		t.Errorf("Expected Encoding 'base64', got %s", field.Encoding)
	}
	if field.Secured {
		t.Error("Expected Secured to be false")
	}
}

func TestFieldConstants(t *testing.T) {
	tests := []struct {
		name        string
		fieldName   string
		description string
	}{
		{"Subtype", SubtypeFieldName, SubtypeFieldDescription},
		{"ScriptSsh", ScriptSshFieldName, ScriptSshFieldDescription},
		{"Timeout", TimeoutFieldName, TimeoutFieldDescription},
		{"HttpHost", HttpHostFieldName, HttpHostFieldDescription},
		{"HttpBody", HttpBodyFieldName, HttpBodyFieldDescription},
		{"HttpMethod", HttpMethodFieldName, HttpMethodFieldDescription},
		{"HttpHeader", HttpHeaderFieldName, HttpHeaderFieldDescription},
		{"HttpIgnoreCertErrors", HttpIgnoreCertErrorsFieldName, HttpIgnoreCertErrorsFieldDescription},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fieldName == "" {
				t.Errorf("Expected %s field name to be non-empty", tt.name)
			}
			if tt.description == "" {
				t.Errorf("Expected %s field description to be non-empty", tt.name)
			}
		})
	}
}
