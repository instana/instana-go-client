package rest

import (
	"errors"
	"testing"
)

// Tests for NewReadOnlyRestResource
func TestNewReadOnlyRestResource(t *testing.T) {
	client := &mockRestClient{}
	unmarshaller := &mockUnmarshaller[*testDataObject]{}

	resource := NewReadOnlyRestResource("/api/test", unmarshaller, client)

	if resource == nil {
		t.Fatal("Expected non-nil resource")
	}

	rr, ok := resource.(*readOnlyRestResource[*testDataObject])
	if !ok {
		t.Fatal("Expected readOnlyRestResource type")
	}

	if rr.resourcePath != "/api/test" {
		t.Errorf("Expected path /api/test, got %s", rr.resourcePath)
	}
}

// Tests for GetAll
func TestReadOnlyRestResource_GetAll_Success(t *testing.T) {
	expectedData := []byte(`[{"id":"1","name":"test"}]`)
	expectedObjects := &[]*testDataObject{{ID: "1", Name: "test"}}

	client := &mockRestClient{
		getFunc: func(path string) ([]byte, error) {
			if path != "/api/test" {
				t.Errorf("Expected path /api/test, got %s", path)
			}
			return expectedData, nil
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{
		unmarshalArrayFunc: func(data []byte) (*[]*testDataObject, error) {
			return expectedObjects, nil
		},
	}

	resource := NewReadOnlyRestResource("/api/test", unmarshaller, client)

	result, err := resource.GetAll()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("Expected non-nil result")
	}
	if len(*result) != 1 {
		t.Errorf("Expected 1 object, got %d", len(*result))
	}
}

func TestReadOnlyRestResource_GetAll_ClientError(t *testing.T) {
	expectedError := errors.New("client error")

	client := &mockRestClient{
		getFunc: func(path string) ([]byte, error) {
			return nil, expectedError
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{}
	resource := NewReadOnlyRestResource("/api/test", unmarshaller, client)

	result, err := resource.GetAll()

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if result != nil {
		t.Error("Expected nil result on error")
	}
}

func TestReadOnlyRestResource_GetAll_UnmarshalError(t *testing.T) {
	expectedData := []byte(`[{"id":"1","name":"test"}]`)
	expectedError := errors.New("unmarshal error")

	client := &mockRestClient{
		getFunc: func(path string) ([]byte, error) {
			return expectedData, nil
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{
		unmarshalArrayFunc: func(data []byte) (*[]*testDataObject, error) {
			return nil, expectedError
		},
	}

	resource := NewReadOnlyRestResource("/api/test", unmarshaller, client)

	result, err := resource.GetAll()

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if result != nil {
		t.Error("Expected nil result on unmarshal error")
	}
}

// Tests for GetByQuery
func TestReadOnlyRestResource_GetByQuery_Success(t *testing.T) {
	expectedData := []byte(`[{"id":"1","name":"test"}]`)
	expectedObjects := &[]*testDataObject{{ID: "1", Name: "test"}}
	queryParams := map[string]string{"filter": "active"}

	client := &mockRestClient{
		getByQueryFunc: func(path string, params map[string]string) ([]byte, error) {
			if path != "/api/test" {
				t.Errorf("Expected path /api/test, got %s", path)
			}
			if params["filter"] != "active" {
				t.Errorf("Expected filter=active, got %v", params)
			}
			return expectedData, nil
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{
		unmarshalArrayFunc: func(data []byte) (*[]*testDataObject, error) {
			return expectedObjects, nil
		},
	}

	resource := NewReadOnlyRestResource("/api/test", unmarshaller, client)

	result, err := resource.GetByQuery(queryParams)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("Expected non-nil result")
	}
	if len(*result) != 1 {
		t.Errorf("Expected 1 object, got %d", len(*result))
	}
}

func TestReadOnlyRestResource_GetByQuery_ClientError(t *testing.T) {
	expectedError := errors.New("client error")
	queryParams := map[string]string{"filter": "active"}

	client := &mockRestClient{
		getByQueryFunc: func(path string, params map[string]string) ([]byte, error) {
			return nil, expectedError
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{}
	resource := NewReadOnlyRestResource("/api/test", unmarshaller, client)

	result, err := resource.GetByQuery(queryParams)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if result != nil {
		t.Error("Expected nil result on error")
	}
}

func TestReadOnlyRestResource_GetByQuery_UnmarshalError(t *testing.T) {
	expectedData := []byte(`[{"id":"1","name":"test"}]`)
	expectedError := errors.New("unmarshal error")
	queryParams := map[string]string{"filter": "active"}

	client := &mockRestClient{
		getByQueryFunc: func(path string, params map[string]string) ([]byte, error) {
			return expectedData, nil
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{
		unmarshalArrayFunc: func(data []byte) (*[]*testDataObject, error) {
			return nil, expectedError
		},
	}

	resource := NewReadOnlyRestResource("/api/test", unmarshaller, client)

	result, err := resource.GetByQuery(queryParams)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if result != nil {
		t.Error("Expected nil result on unmarshal error")
	}
}

// Tests for GetOne
func TestReadOnlyRestResource_GetOne_Success(t *testing.T) {
	expectedData := []byte(`{"id":"1","name":"test"}`)
	expectedObject := &testDataObject{ID: "1", Name: "test"}

	client := &mockRestClient{
		getOneFunc: func(id, path string) ([]byte, error) {
			if id != "1" {
				t.Errorf("Expected id 1, got %s", id)
			}
			if path != "/api/test" {
				t.Errorf("Expected path /api/test, got %s", path)
			}
			return expectedData, nil
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{
		unmarshalFunc: func(data []byte) (*testDataObject, error) {
			return expectedObject, nil
		},
	}

	resource := NewReadOnlyRestResource("/api/test", unmarshaller, client)

	result, err := resource.GetOne("1")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "1" {
		t.Errorf("Expected ID 1, got %s", result.ID)
	}
	if result.Name != "test" {
		t.Errorf("Expected name test, got %s", result.Name)
	}
}

func TestReadOnlyRestResource_GetOne_ClientError(t *testing.T) {
	expectedError := errors.New("client error")

	client := &mockRestClient{
		getOneFunc: func(id, path string) ([]byte, error) {
			return nil, expectedError
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{}
	resource := NewReadOnlyRestResource("/api/test", unmarshaller, client)

	result, err := resource.GetOne("1")

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	// Result should be zero value
	if result != nil && result.ID != "" {
		t.Error("Expected zero value result on error")
	}
}

func TestReadOnlyRestResource_GetOne_UnmarshalError(t *testing.T) {
	expectedData := []byte(`{"id":"1","name":"test"}`)
	expectedError := errors.New("unmarshal error")

	client := &mockRestClient{
		getOneFunc: func(id, path string) ([]byte, error) {
			return expectedData, nil
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{
		unmarshalFunc: func(data []byte) (*testDataObject, error) {
			return nil, expectedError
		},
	}

	resource := NewReadOnlyRestResource("/api/test", unmarshaller, client)

	_, err := resource.GetOne("1")

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

// Test with empty query params
func TestReadOnlyRestResource_GetByQuery_EmptyParams(t *testing.T) {
	expectedData := []byte(`[{"id":"1","name":"test"}]`)
	expectedObjects := &[]*testDataObject{{ID: "1", Name: "test"}}
	queryParams := map[string]string{}

	client := &mockRestClient{
		getByQueryFunc: func(path string, params map[string]string) ([]byte, error) {
			if len(params) != 0 {
				t.Error("Expected empty params")
			}
			return expectedData, nil
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{
		unmarshalArrayFunc: func(data []byte) (*[]*testDataObject, error) {
			return expectedObjects, nil
		},
	}

	resource := NewReadOnlyRestResource("/api/test", unmarshaller, client)

	result, err := resource.GetByQuery(queryParams)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("Expected non-nil result")
	}
}
