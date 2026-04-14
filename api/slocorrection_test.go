package api_test

import (
	"encoding/json"
	"testing"

	. "github.com/instana/instana-go-client/api"
	"github.com/stretchr/testify/require"
)

func TestSloCorrectionConfigResourcePath(t *testing.T) {
	expected := "/api/settings/correction"
	if SloCorrectionConfigResourcePath != expected {
		t.Errorf("Expected SloCorrectionConfigResourcePath to be %s, got %s", expected, SloCorrectionConfigResourcePath)
	}
}

func TestSloCorrectionConfigGetIDForResourcePath(t *testing.T) {
	id := "test-correction-id"
	config := SloCorrectionConfig{ID: id}

	result := config.GetIDForResourcePath()

	if result != id {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", id, result)
	}
}

func TestSloCorrectionConfigStructure(t *testing.T) {
	id := "correction-id"
	name := "Test SLO Correction"
	description := "Test correction description"
	active := true

	config := SloCorrectionConfig{
		ID:          id,
		Name:        name,
		Description: description,
		Active:      active,
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
	if config.Active != active {
		t.Errorf("Expected Active to be %v, got %v", active, config.Active)
	}
}

func TestNewSloCorrectionAgentJSONUnmarshaller(t *testing.T) {
	testData := &testObject{
		ID:   defaultObjectId,
		Name: defaultObjectName,
	}
	testObjects := []*testObject{testData, testData}

	serializedJSON, _ := json.Marshal(testObjects)

	sut := NewSloCorrectionConfigJSONUnmarshaller(&testObject{})

	_, err := sut.Unmarshal(serializedJSON)

	require.Error(t, err)
}

func TestShouldSuccessfullyUnmarshalSloCorrectionArrayOfObjects(t *testing.T) {
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

	sut := NewSloCorrectionConfigJSONUnmarshaller(&testObject{})

	result, err := sut.UnmarshalArray(serializedJSON)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, len(testObjects), len(*result))
	require.Equal(t, testObjects[0].ID, (*result)[0].ID)
	require.Equal(t, testObjects[0].Name, (*result)[0].Name)
}
