package api_test

import (
	"encoding/json"
	"testing"

	. "github.com/instana/instana-go-client/api"
	"github.com/stretchr/testify/require"
)

func TestSloConfigResourcePath(t *testing.T) {
	expected := "/api/settings/slo"
	if SloConfigResourcePath != expected {
		t.Errorf("Expected SloConfigResourcePath to be %s, got %s", expected, SloConfigResourcePath)
	}
}

func TestSloConfigGetIDForResourcePath(t *testing.T) {
	id := "test-slo-id"
	config := SloConfig{ID: id}

	result := config.GetIDForResourcePath()

	if result != id {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", id, result)
	}
}

func TestSloConfigStructure(t *testing.T) {
	id := "slo-id"
	name := "Test SLO"

	config := SloConfig{
		ID:   id,
		Name: name,
	}

	if config.ID != id {
		t.Errorf("Expected ID to be %s, got %s", id, config.ID)
	}
	if config.Name != name {
		t.Errorf("Expected Name to be %s, got %s", name, config.Name)
	}
}

func TestNewSloAgentJSONUnmarshaller(t *testing.T) {
	testData := &testObject{
		ID:   defaultObjectId,
		Name: defaultObjectName,
	}
	testObjects := []*testObject{testData, testData}

	serializedJSON, _ := json.Marshal(testObjects)

	sut := NewSloConfigJSONUnmarshaller(&testObject{})

	_, err := sut.Unmarshal(serializedJSON)

	require.Error(t, err)
}

func TestShouldSuccessfullyUnmarshalSloArrayOfObjects(t *testing.T) {
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

	sut := NewSloConfigJSONUnmarshaller(&testObject{})

	result, err := sut.UnmarshalArray(serializedJSON)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, len(testObjects), len(*result))
	require.Equal(t, testObjects[0].ID, (*result)[0].ID)
	require.Equal(t, testObjects[0].Name, (*result)[0].Name)
}
