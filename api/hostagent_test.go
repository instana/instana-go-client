package api_test

import (
	"encoding/json"
	"testing"

	. "github.com/instana/instana-go-client/api"
	"github.com/stretchr/testify/require"
)

const (
	defaultObjectId   = "object-id"
	defaultObjectName = "object-name"
)

type testObject struct {
	ID   string
	Name string
}

func (t *testObject) GetIDForResourcePath() string {
	return t.ID
}

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

func TestNewHostAgentJSONUnmarshaller(t *testing.T) {
	testData := &testObject{
		ID:   defaultObjectId,
		Name: defaultObjectName,
	}
	testObjects := []*testObject{testData, testData}

	serializedJSON, _ := json.Marshal(testObjects)

	sut := NewHostAgentJSONUnmarshaller(&testObject{})

	_, err := sut.Unmarshal(serializedJSON)

	require.Error(t, err)
}

func TestShouldSuccessfullyUnmarshalArrayOfObjects(t *testing.T) {
	testData := &testObject{
		ID:   defaultObjectId,
		Name: defaultObjectName,
	}
	testObjects := []*testObject{testData, testData}

	// The UnmarshalArray expects JSON with "items" key
	wrappedData := map[string][]*testObject{
		"items": testObjects,
	}
	serializedJSON, _ := json.Marshal(wrappedData)

	sut := NewHostAgentJSONUnmarshaller(&testObject{})

	result, err := sut.UnmarshalArray(serializedJSON)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, len(testObjects), len(*result))
	require.Equal(t, testObjects[0].ID, (*result)[0].ID)
	require.Equal(t, testObjects[0].Name, (*result)[0].Name)
}
