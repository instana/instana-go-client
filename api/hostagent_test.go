package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
)

func TestHostAgentResourcePath(t *testing.T) {
	expected := "/api/host-agent"
	if HostAgentResourcePath != expected {
		t.Errorf("Expected HostAgentResourcePath to be %s, got %s", expected, HostAgentResourcePath)
	}
}

func TestHostAgentGetIDForResourcePath(t *testing.T) {
	testID := "test-snapshot-123"
	agent := &HostAgent{
		SnapshotID: testID,
		Label:      "Test Agent",
		Host:       "server1",
	}

	result := agent.GetIDForResourcePath()
	if result != testID {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", testID, result)
	}
}

func TestHostAgentStructure(t *testing.T) {
	agent := HostAgent{
		SnapshotID: "snapshot-456",
		Label:      "Production Agent",
		Host:       "prod-server-1",
		Plugin:     "host",
		Tags:       []string{"production", "critical"},
	}

	if agent.SnapshotID != "snapshot-456" {
		t.Errorf("Expected SnapshotID 'snapshot-456', got %s", agent.SnapshotID)
	}
	if agent.Label != "Production Agent" {
		t.Errorf("Expected Label 'Production Agent', got %s", agent.Label)
	}
	if agent.Host != "prod-server-1" {
		t.Errorf("Expected Host 'prod-server-1', got %s", agent.Host)
	}
	if len(agent.Tags) != 2 {
		t.Errorf("Expected 2 tags, got %d", len(agent.Tags))
	}
}
