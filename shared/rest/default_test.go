package rest

import (
	"errors"
	"testing"
)

// Mock implementations for testing
type mockRestClient struct {
	getFunc         func(string) ([]byte, error)
	getOneFunc      func(string, string) ([]byte, error)
	postFunc        func(InstanaDataObject, string) ([]byte, error)
	postWithIDFunc  func(InstanaDataObject, string) ([]byte, error)
	putFunc         func(InstanaDataObject, string) ([]byte, error)
	deleteFunc      func(string, string) error
	getByQueryFunc  func(string, map[string]string) ([]byte, error)
	postByQueryFunc func(string, map[string]string) ([]byte, error)
	putByQueryFunc  func(string, string, map[string]string) ([]byte, error)
}

func (m *mockRestClient) Get(path string) ([]byte, error) {
	if m.getFunc != nil {
		return m.getFunc(path)
	}
	return nil, nil
}

func (m *mockRestClient) GetOne(id, path string) ([]byte, error) {
	if m.getOneFunc != nil {
		return m.getOneFunc(id, path)
	}
	return nil, nil
}

func (m *mockRestClient) Post(data InstanaDataObject, path string) ([]byte, error) {
	if m.postFunc != nil {
		return m.postFunc(data, path)
	}
	return nil, nil
}

func (m *mockRestClient) PostWithID(data InstanaDataObject, path string) ([]byte, error) {
	if m.postWithIDFunc != nil {
		return m.postWithIDFunc(data, path)
	}
	return nil, nil
}

func (m *mockRestClient) Put(data InstanaDataObject, path string) ([]byte, error) {
	if m.putFunc != nil {
		return m.putFunc(data, path)
	}
	return nil, nil
}

func (m *mockRestClient) Delete(id, path string) error {
	if m.deleteFunc != nil {
		return m.deleteFunc(id, path)
	}
	return nil
}

func (m *mockRestClient) GetByQuery(path string, params map[string]string) ([]byte, error) {
	if m.getByQueryFunc != nil {
		return m.getByQueryFunc(path, params)
	}
	return nil, nil
}

func (m *mockRestClient) PostByQuery(path string, params map[string]string) ([]byte, error) {
	if m.postByQueryFunc != nil {
		return m.postByQueryFunc(path, params)
	}
	return nil, nil
}

func (m *mockRestClient) PutByQuery(path, is string, params map[string]string) ([]byte, error) {
	if m.putByQueryFunc != nil {
		return m.putByQueryFunc(path, is, params)
	}
	return nil, nil
}

type mockUnmarshaller[T any] struct {
	unmarshalFunc      func([]byte) (T, error)
	unmarshalArrayFunc func([]byte) (*[]T, error)
}

func (m *mockUnmarshaller[T]) Unmarshal(data []byte) (T, error) {
	if m.unmarshalFunc != nil {
		return m.unmarshalFunc(data)
	}
	var zero T
	return zero, nil
}

func (m *mockUnmarshaller[T]) UnmarshalArray(data []byte) (*[]T, error) {
	if m.unmarshalArrayFunc != nil {
		return m.unmarshalArrayFunc(data)
	}
	return nil, nil
}

// Test data object
type testDataObject struct {
	ID   string
	Name string
}

func (t *testDataObject) GetIDForResourcePath() string {
	return t.ID
}

// Tests for factory functions
func TestNewCreatePUTUpdatePUTRestResource(t *testing.T) {
	client := &mockRestClient{}
	unmarshaller := &mockUnmarshaller[*testDataObject]{}

	resource := NewCreatePUTUpdatePUTRestResource("/api/test", unmarshaller, client)

	if resource == nil {
		t.Fatal("Expected non-nil resource")
	}

	dr, ok := resource.(*defaultRestResource[*testDataObject])
	if !ok {
		t.Fatal("Expected defaultRestResource type")
	}

	if dr.mode != DefaultRestResourceModeCreateAndUpdatePUT {
		t.Errorf("Expected mode CREATE_PUT_UPDATE_PUT, got %s", dr.mode)
	}
}

func TestNewCreatePOSTUpdatePUTRestResource(t *testing.T) {
	client := &mockRestClient{}
	unmarshaller := &mockUnmarshaller[*testDataObject]{}

	resource := NewCreatePOSTUpdatePUTRestResource("/api/test", unmarshaller, client)

	if resource == nil {
		t.Fatal("Expected non-nil resource")
	}

	dr := resource.(*defaultRestResource[*testDataObject])
	if dr.mode != DefaultRestResourceModeCreatePOSTUpdatePUT {
		t.Errorf("Expected mode CREATE_POST_UPDATE_PUT, got %s", dr.mode)
	}
}

func TestNewCreatePOSTUpdatePOSTRestResource(t *testing.T) {
	client := &mockRestClient{}
	unmarshaller := &mockUnmarshaller[*testDataObject]{}

	resource := NewCreatePOSTUpdatePOSTRestResource("/api/test", unmarshaller, client)

	if resource == nil {
		t.Fatal("Expected non-nil resource")
	}

	dr := resource.(*defaultRestResource[*testDataObject])
	if dr.mode != DefaultRestResourceModeCreateAndUpdatePOST {
		t.Errorf("Expected mode CREATE_POST_UPDATE_POST, got %s", dr.mode)
	}
}

func TestNewCreatePOSTUpdateNotSupportedRestResource(t *testing.T) {
	client := &mockRestClient{}
	unmarshaller := &mockUnmarshaller[*testDataObject]{}

	resource := NewCreatePOSTUpdateNotSupportedRestResource("/api/test", unmarshaller, client)

	if resource == nil {
		t.Fatal("Expected non-nil resource")
	}

	dr := resource.(*defaultRestResource[*testDataObject])
	if dr.mode != DefaultRestResourceModeCreatePOSTAndUpdateNotSupported {
		t.Errorf("Expected mode CREATE_POST_UPDATE_NOT_SUPPORTED, got %s", dr.mode)
	}
}

func TestNewCreatePUTUpdateNotSupportedRestResource(t *testing.T) {
	client := &mockRestClient{}
	unmarshaller := &mockUnmarshaller[*testDataObject]{}

	resource := NewCreatePUTUpdateNotSupportedRestResource("/api/test", unmarshaller, client)

	if resource == nil {
		t.Fatal("Expected non-nil resource")
	}

	dr := resource.(*defaultRestResource[*testDataObject])
	if dr.mode != DefaultRestResourceModeCreatePUTAndUpdateNotSupported {
		t.Errorf("Expected mode CREATE_PUT_UPDATE_NOT_SUPPORTED, got %s", dr.mode)
	}
}

// Tests for GetAll
func TestDefaultRestResource_GetAll_Success(t *testing.T) {
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

	resource := NewCreatePOSTUpdatePUTRestResource("/api/test", unmarshaller, client)

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

func TestDefaultRestResource_GetAll_ClientError(t *testing.T) {
	expectedError := errors.New("client error")

	client := &mockRestClient{
		getFunc: func(path string) ([]byte, error) {
			return nil, expectedError
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{}
	resource := NewCreatePOSTUpdatePUTRestResource("/api/test", unmarshaller, client)

	result, err := resource.GetAll()

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if result != nil {
		t.Error("Expected nil result on error")
	}
}

// Tests for GetOne
func TestDefaultRestResource_GetOne_Success(t *testing.T) {
	expectedData := []byte(`{"id":"1","name":"test"}`)
	expectedObject := &testDataObject{ID: "1", Name: "test"}

	client := &mockRestClient{
		getOneFunc: func(id, path string) ([]byte, error) {
			if id != "1" {
				t.Errorf("Expected id 1, got %s", id)
			}
			return expectedData, nil
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{
		unmarshalFunc: func(data []byte) (*testDataObject, error) {
			return expectedObject, nil
		},
	}

	resource := NewCreatePOSTUpdatePUTRestResource("/api/test", unmarshaller, client)

	result, err := resource.GetOne("1")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "1" {
		t.Errorf("Expected ID 1, got %s", result.ID)
	}
}

// Tests for Create
func TestDefaultRestResource_Create_WithPUT(t *testing.T) {
	inputData := &testDataObject{ID: "1", Name: "test"}
	responseData := []byte(`{"id":"1","name":"test"}`)

	client := &mockRestClient{
		putFunc: func(data InstanaDataObject, path string) ([]byte, error) {
			return responseData, nil
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{
		unmarshalFunc: func(data []byte) (*testDataObject, error) {
			return inputData, nil
		},
	}

	resource := NewCreatePUTUpdatePUTRestResource("/api/test", unmarshaller, client)

	result, err := resource.Create(inputData)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "1" {
		t.Errorf("Expected ID 1, got %s", result.ID)
	}
}

func TestDefaultRestResource_Create_WithPOST(t *testing.T) {
	inputData := &testDataObject{ID: "1", Name: "test"}
	responseData := []byte(`{"id":"1","name":"test"}`)

	client := &mockRestClient{
		postFunc: func(data InstanaDataObject, path string) ([]byte, error) {
			return responseData, nil
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{
		unmarshalFunc: func(data []byte) (*testDataObject, error) {
			return inputData, nil
		},
	}

	resource := NewCreatePOSTUpdatePUTRestResource("/api/test", unmarshaller, client)

	result, err := resource.Create(inputData)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "1" {
		t.Errorf("Expected ID 1, got %s", result.ID)
	}
}

// Tests for Update
func TestDefaultRestResource_Update_WithPUT(t *testing.T) {
	inputData := &testDataObject{ID: "1", Name: "updated"}
	responseData := []byte(`{"id":"1","name":"updated"}`)

	client := &mockRestClient{
		putFunc: func(data InstanaDataObject, path string) ([]byte, error) {
			return responseData, nil
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{
		unmarshalFunc: func(data []byte) (*testDataObject, error) {
			return inputData, nil
		},
	}

	resource := NewCreatePOSTUpdatePUTRestResource("/api/test", unmarshaller, client)

	result, err := resource.Update(inputData)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Name != "updated" {
		t.Errorf("Expected name updated, got %s", result.Name)
	}
}

func TestDefaultRestResource_Update_WithPOST(t *testing.T) {
	inputData := &testDataObject{ID: "1", Name: "updated"}
	responseData := []byte(`{"id":"1","name":"updated"}`)

	client := &mockRestClient{
		postWithIDFunc: func(data InstanaDataObject, path string) ([]byte, error) {
			return responseData, nil
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{
		unmarshalFunc: func(data []byte) (*testDataObject, error) {
			return inputData, nil
		},
	}

	resource := NewCreatePOSTUpdatePOSTRestResource("/api/test", unmarshaller, client)

	result, err := resource.Update(inputData)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Name != "updated" {
		t.Errorf("Expected name updated, got %s", result.Name)
	}
}

func TestDefaultRestResource_Update_NoContentResponse(t *testing.T) {
	inputData := &testDataObject{ID: "1", Name: "test"}
	responseData := []byte(`{}`)

	client := &mockRestClient{
		postWithIDFunc: func(data InstanaDataObject, path string) ([]byte, error) {
			return responseData, nil
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{}
	resource := NewCreatePOSTUpdatePOSTRestResource("/api/test", unmarshaller, client)

	result, err := resource.Update(inputData)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "1" {
		t.Errorf("Expected original data returned")
	}
}

func TestDefaultRestResource_Update_NotSupported(t *testing.T) {
	inputData := &testDataObject{ID: "1", Name: "test"}

	client := &mockRestClient{}
	unmarshaller := &mockUnmarshaller[*testDataObject]{
		unmarshalFunc: func(data []byte) (*testDataObject, error) {
			return &testDataObject{}, nil
		},
	}

	resource := NewCreatePOSTUpdateNotSupportedRestResource("/api/test", unmarshaller, client)

	_, err := resource.Update(inputData)

	if err == nil {
		t.Fatal("Expected error for unsupported update")
	}
}

// Tests for Delete
func TestDefaultRestResource_Delete_Success(t *testing.T) {
	inputData := &testDataObject{ID: "1", Name: "test"}

	client := &mockRestClient{
		deleteFunc: func(id, path string) error {
			if id != "1" {
				t.Errorf("Expected id 1, got %s", id)
			}
			return nil
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{}
	resource := NewCreatePOSTUpdatePUTRestResource("/api/test", unmarshaller, client)

	err := resource.Delete(inputData)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestDefaultRestResource_DeleteByID_Success(t *testing.T) {
	client := &mockRestClient{
		deleteFunc: func(id, path string) error {
			if id != "1" {
				t.Errorf("Expected id 1, got %s", id)
			}
			return nil
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{}
	resource := NewCreatePOSTUpdatePUTRestResource("/api/test", unmarshaller, client)

	err := resource.DeleteByID("1")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestDefaultRestResource_DeleteByID_Error(t *testing.T) {
	expectedError := errors.New("delete error")

	client := &mockRestClient{
		deleteFunc: func(id, path string) error {
			return expectedError
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{}
	resource := NewCreatePOSTUpdatePUTRestResource("/api/test", unmarshaller, client)

	err := resource.DeleteByID("1")

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

// Additional error path tests
func TestDefaultRestResource_GetAll_UnmarshalError(t *testing.T) {
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

	resource := NewCreatePOSTUpdatePUTRestResource("/api/test", unmarshaller, client)

	result, err := resource.GetAll()

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if result != nil {
		t.Error("Expected nil result on unmarshal error")
	}
}

func TestDefaultRestResource_GetOne_ClientError(t *testing.T) {
	expectedError := errors.New("client error")

	client := &mockRestClient{
		getOneFunc: func(id, path string) ([]byte, error) {
			return nil, expectedError
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{}
	resource := NewCreatePOSTUpdatePUTRestResource("/api/test", unmarshaller, client)

	result, err := resource.GetOne("1")

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if result != nil && result.ID != "" {
		t.Error("Expected zero value on error")
	}
}

func TestDefaultRestResource_GetOne_UnmarshalError(t *testing.T) {
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

	resource := NewCreatePOSTUpdatePUTRestResource("/api/test", unmarshaller, client)

	_, err := resource.GetOne("1")

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestDefaultRestResource_Create_ClientError(t *testing.T) {
	inputData := &testDataObject{ID: "1", Name: "test"}
	expectedError := errors.New("client error")

	client := &mockRestClient{
		postFunc: func(data InstanaDataObject, path string) ([]byte, error) {
			return nil, expectedError
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{}
	resource := NewCreatePOSTUpdatePUTRestResource("/api/test", unmarshaller, client)

	_, err := resource.Create(inputData)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestDefaultRestResource_Create_UnmarshalError(t *testing.T) {
	inputData := &testDataObject{ID: "1", Name: "test"}
	responseData := []byte(`{"id":"1","name":"test"}`)
	expectedError := errors.New("unmarshal error")

	client := &mockRestClient{
		postFunc: func(data InstanaDataObject, path string) ([]byte, error) {
			return responseData, nil
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{
		unmarshalFunc: func(data []byte) (*testDataObject, error) {
			return nil, expectedError
		},
	}

	resource := NewCreatePOSTUpdatePUTRestResource("/api/test", unmarshaller, client)

	_, err := resource.Create(inputData)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestDefaultRestResource_Update_ClientError(t *testing.T) {
	inputData := &testDataObject{ID: "1", Name: "updated"}
	expectedError := errors.New("client error")

	client := &mockRestClient{
		putFunc: func(data InstanaDataObject, path string) ([]byte, error) {
			return nil, expectedError
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{}
	resource := NewCreatePOSTUpdatePUTRestResource("/api/test", unmarshaller, client)

	_, err := resource.Update(inputData)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestDefaultRestResource_Update_UnmarshalError(t *testing.T) {
	inputData := &testDataObject{ID: "1", Name: "updated"}
	responseData := []byte(`{"id":"1","name":"updated"}`)
	expectedError := errors.New("unmarshal error")

	client := &mockRestClient{
		putFunc: func(data InstanaDataObject, path string) ([]byte, error) {
			return responseData, nil
		},
	}

	unmarshaller := &mockUnmarshaller[*testDataObject]{
		unmarshalFunc: func(data []byte) (*testDataObject, error) {
			return nil, expectedError
		},
	}

	resource := NewCreatePOSTUpdatePUTRestResource("/api/test", unmarshaller, client)

	_, err := resource.Update(inputData)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestDefaultRestResource_Update_NotSupported_UnmarshalError(t *testing.T) {
	inputData := &testDataObject{ID: "1", Name: "test"}
	expectedError := errors.New("unmarshal error")

	client := &mockRestClient{}
	unmarshaller := &mockUnmarshaller[*testDataObject]{
		unmarshalFunc: func(data []byte) (*testDataObject, error) {
			return nil, expectedError
		},
	}

	resource := NewCreatePUTUpdateNotSupportedRestResource("/api/test", unmarshaller, client)

	_, err := resource.Update(inputData)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if err.Error() != "update is not supported for /api/test; unmarshal error" {
		t.Errorf("Unexpected error message: %v", err)
	}
}
