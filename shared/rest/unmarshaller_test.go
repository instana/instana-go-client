package rest_test

import (
	"testing"

	"github.com/instana/instana-go-client/shared/rest"
	"github.com/stretchr/testify/require"
)

// TestModel is a simple test model for unmarshalling tests
type TestModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (t *TestModel) GetIDForResourcePath() string {
	return t.ID
}

func TestGenericUnmarshaller_Unmarshal_Success(t *testing.T) {
	unmarshaller := rest.NewGenericUnmarshaller[*TestModel]()
	jsonData := []byte(`{"id": "test-123", "name": "Test Name"}`)

	result, err := unmarshaller.Unmarshal(jsonData)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, "test-123", result.ID)
	require.Equal(t, "Test Name", result.Name)
}

func TestGenericUnmarshaller_Unmarshal_InvalidJSON(t *testing.T) {
	unmarshaller := rest.NewGenericUnmarshaller[*TestModel]()
	jsonData := []byte(`{"id": "test-123", "name": }`) // Invalid JSON

	result, err := unmarshaller.Unmarshal(jsonData)

	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to parse json")
	require.Nil(t, result)
}

func TestGenericUnmarshaller_Unmarshal_EmptyJSON(t *testing.T) {
	unmarshaller := rest.NewGenericUnmarshaller[*TestModel]()
	jsonData := []byte(`{}`)

	result, err := unmarshaller.Unmarshal(jsonData)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, "", result.ID)
	require.Equal(t, "", result.Name)
}

func TestGenericUnmarshaller_UnmarshalArray_Success(t *testing.T) {
	unmarshaller := rest.NewGenericUnmarshaller[*TestModel]()
	jsonData := []byte(`[
		{"id": "test-1", "name": "Name 1"},
		{"id": "test-2", "name": "Name 2"},
		{"id": "test-3", "name": "Name 3"}
	]`)

	result, err := unmarshaller.UnmarshalArray(jsonData)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.Len(t, *result, 3)
	require.Equal(t, "test-1", (*result)[0].ID)
	require.Equal(t, "Name 1", (*result)[0].Name)
	require.Equal(t, "test-2", (*result)[1].ID)
	require.Equal(t, "Name 2", (*result)[1].Name)
	require.Equal(t, "test-3", (*result)[2].ID)
	require.Equal(t, "Name 3", (*result)[2].Name)
}

func TestGenericUnmarshaller_UnmarshalArray_EmptyArray(t *testing.T) {
	unmarshaller := rest.NewGenericUnmarshaller[*TestModel]()
	jsonData := []byte(`[]`)

	result, err := unmarshaller.UnmarshalArray(jsonData)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.Len(t, *result, 0)
}

func TestGenericUnmarshaller_UnmarshalArray_InvalidJSON(t *testing.T) {
	unmarshaller := rest.NewGenericUnmarshaller[*TestModel]()
	jsonData := []byte(`[{"id": "test-1", "name": }]`) // Invalid JSON

	result, err := unmarshaller.UnmarshalArray(jsonData)

	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to parse json")
	require.NotNil(t, result) // Returns pointer to empty slice
	require.Len(t, *result, 0)
}

func TestGenericUnmarshaller_UnmarshalArray_NotAnArray(t *testing.T) {
	unmarshaller := rest.NewGenericUnmarshaller[*TestModel]()
	jsonData := []byte(`{"id": "test-1", "name": "Name 1"}`) // Object, not array

	_, err := unmarshaller.UnmarshalArray(jsonData)

	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to parse json")
}

// TestGenericUnmarshaller_WithDifferentTypes verifies the unmarshaller works with various types
func TestGenericUnmarshaller_WithDifferentTypes(t *testing.T) {
	t.Run("string type", func(t *testing.T) {
		unmarshaller := rest.NewGenericUnmarshaller[string]()
		jsonData := []byte(`"test string"`)

		result, err := unmarshaller.Unmarshal(jsonData)

		require.NoError(t, err)
		require.Equal(t, "test string", result)
	})

	t.Run("int type", func(t *testing.T) {
		unmarshaller := rest.NewGenericUnmarshaller[int]()
		jsonData := []byte(`42`)

		result, err := unmarshaller.Unmarshal(jsonData)

		require.NoError(t, err)
		require.Equal(t, 42, result)
	})

	t.Run("map type", func(t *testing.T) {
		unmarshaller := rest.NewGenericUnmarshaller[map[string]string]()
		jsonData := []byte(`{"key1": "value1", "key2": "value2"}`)

		result, err := unmarshaller.Unmarshal(jsonData)

		require.NoError(t, err)
		require.Len(t, result, 2)
		require.Equal(t, "value1", result["key1"])
		require.Equal(t, "value2", result["key2"])
	})
}
