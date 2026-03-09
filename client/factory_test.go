package client

import (
	"testing"

	"github.com/instana/instana-go-client/mocks"
	"github.com/instana/instana-go-client/shared/rest"
	"go.uber.org/mock/gomock"
)

// TestNewRestResourceWithCreateAndUpdatePUT tests factory with PUT for both create and update
func TestNewRestResourceWithCreateAndUpdatePUT(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	mockUnmarshaller := mocks.NewMockJSONUnmarshaller[*mocks.MockInstanaDataObject](ctrl)

	resource := NewRestResource[*mocks.MockInstanaDataObject](
		mockRestClient,
		"/api/test",
		rest.DefaultRestResourceModeCreateAndUpdatePUT,
		mockUnmarshaller,
	)

	if resource == nil {
		t.Fatal("Expected non-nil RestResource")
	}
}

// TestNewRestResourceWithCreatePOSTUpdatePUT tests factory with POST for create, PUT for update
func TestNewRestResourceWithCreatePOSTUpdatePUT(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	mockUnmarshaller := mocks.NewMockJSONUnmarshaller[*mocks.MockInstanaDataObject](ctrl)

	resource := NewRestResource[*mocks.MockInstanaDataObject](
		mockRestClient,
		"/api/test",
		rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
		mockUnmarshaller,
	)

	if resource == nil {
		t.Fatal("Expected non-nil RestResource")
	}
}

// TestNewRestResourceWithCreateAndUpdatePOST tests factory with POST for both create and update
func TestNewRestResourceWithCreateAndUpdatePOST(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	mockUnmarshaller := mocks.NewMockJSONUnmarshaller[*mocks.MockInstanaDataObject](ctrl)

	resource := NewRestResource[*mocks.MockInstanaDataObject](
		mockRestClient,
		"/api/test",
		rest.DefaultRestResourceModeCreateAndUpdatePOST,
		mockUnmarshaller,
	)

	if resource == nil {
		t.Fatal("Expected non-nil RestResource")
	}
}

// TestNewRestResourceWithCreatePOSTUpdateNotSupported tests factory with POST for create, no update
func TestNewRestResourceWithCreatePOSTUpdateNotSupported(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	mockUnmarshaller := mocks.NewMockJSONUnmarshaller[*mocks.MockInstanaDataObject](ctrl)

	resource := NewRestResource[*mocks.MockInstanaDataObject](
		mockRestClient,
		"/api/test",
		rest.DefaultRestResourceModeCreatePOSTAndUpdateNotSupported,
		mockUnmarshaller,
	)

	if resource == nil {
		t.Fatal("Expected non-nil RestResource")
	}
}

// TestNewRestResourceWithCreatePUTUpdateNotSupported tests factory with PUT for create, no update
func TestNewRestResourceWithCreatePUTUpdateNotSupported(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	mockUnmarshaller := mocks.NewMockJSONUnmarshaller[*mocks.MockInstanaDataObject](ctrl)

	resource := NewRestResource[*mocks.MockInstanaDataObject](
		mockRestClient,
		"/api/test",
		rest.DefaultRestResourceModeCreatePUTAndUpdateNotSupported,
		mockUnmarshaller,
	)

	if resource == nil {
		t.Fatal("Expected non-nil RestResource")
	}
}

// TestNewRestResourceWithDefaultMode tests factory with default/unknown mode
func TestNewRestResourceWithDefaultMode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	mockUnmarshaller := mocks.NewMockJSONUnmarshaller[*mocks.MockInstanaDataObject](ctrl)

	// Use an invalid mode value to trigger default case
	resource := NewRestResource[*mocks.MockInstanaDataObject](
		mockRestClient,
		"/api/test",
		rest.DefaultRestResourceMode("INVALID_MODE"), // Invalid mode
		mockUnmarshaller,
	)

	if resource == nil {
		t.Fatal("Expected non-nil RestResource with default mode")
	}
}

// TestNewRestResourceAllModes tests all supported modes
func TestNewRestResourceAllModes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	mockUnmarshaller := mocks.NewMockJSONUnmarshaller[*mocks.MockInstanaDataObject](ctrl)

	modes := []struct {
		name string
		mode rest.DefaultRestResourceMode
	}{
		{"CreateAndUpdatePUT", rest.DefaultRestResourceModeCreateAndUpdatePUT},
		{"CreatePOSTUpdatePUT", rest.DefaultRestResourceModeCreatePOSTUpdatePUT},
		{"CreateAndUpdatePOST", rest.DefaultRestResourceModeCreateAndUpdatePOST},
		{"CreatePOSTUpdateNotSupported", rest.DefaultRestResourceModeCreatePOSTAndUpdateNotSupported},
		{"CreatePUTUpdateNotSupported", rest.DefaultRestResourceModeCreatePUTAndUpdateNotSupported},
	}

	for _, tt := range modes {
		t.Run(tt.name, func(t *testing.T) {
			resource := NewRestResource[*mocks.MockInstanaDataObject](
				mockRestClient,
				"/api/test",
				tt.mode,
				mockUnmarshaller,
			)

			if resource == nil {
				t.Errorf("Expected non-nil RestResource for mode %s", tt.name)
			}
		})
	}
}

// TestNewReadOnlyRestResource tests read-only resource factory
func TestNewReadOnlyRestResource(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	mockUnmarshaller := mocks.NewMockJSONUnmarshaller[*mocks.MockInstanaDataObject](ctrl)

	resource := NewReadOnlyRestResource[*mocks.MockInstanaDataObject](
		mockRestClient,
		"/api/test",
		mockUnmarshaller,
	)

	if resource == nil {
		t.Fatal("Expected non-nil ReadOnlyRestResource")
	}
}

// TestNewReadOnlyRestResourceWithDifferentPaths tests read-only resource with various paths
func TestNewReadOnlyRestResourceWithDifferentPaths(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	mockUnmarshaller := mocks.NewMockJSONUnmarshaller[*mocks.MockInstanaDataObject](ctrl)

	paths := []string{
		"/api/events/settings/built-in-events",
		"/api/infrastructure-monitoring/snapshots",
		"/api/settings/users",
	}

	for _, path := range paths {
		t.Run(path, func(t *testing.T) {
			resource := NewReadOnlyRestResource[*mocks.MockInstanaDataObject](
				mockRestClient,
				path,
				mockUnmarshaller,
			)

			if resource == nil {
				t.Errorf("Expected non-nil ReadOnlyRestResource for path %s", path)
			}
		})
	}
}

// TestFactoryWithNilRestClient tests factory behavior with nil rest client
func TestFactoryWithNilRestClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUnmarshaller := mocks.NewMockJSONUnmarshaller[*mocks.MockInstanaDataObject](ctrl)

	// This should not panic, but create a resource (behavior depends on implementation)
	resource := NewRestResource[*mocks.MockInstanaDataObject](
		nil,
		"/api/test",
		rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
		mockUnmarshaller,
	)

	if resource == nil {
		t.Fatal("Expected non-nil RestResource even with nil client")
	}
}

// TestFactoryWithEmptyPath tests factory with empty resource path
func TestFactoryWithEmptyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	mockUnmarshaller := mocks.NewMockJSONUnmarshaller[*mocks.MockInstanaDataObject](ctrl)

	resource := NewRestResource[*mocks.MockInstanaDataObject](
		mockRestClient,
		"",
		rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
		mockUnmarshaller,
	)

	if resource == nil {
		t.Fatal("Expected non-nil RestResource even with empty path")
	}
}

// TestReadOnlyFactoryWithNilRestClient tests read-only factory with nil rest client
func TestReadOnlyFactoryWithNilRestClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUnmarshaller := mocks.NewMockJSONUnmarshaller[*mocks.MockInstanaDataObject](ctrl)

	resource := NewReadOnlyRestResource[*mocks.MockInstanaDataObject](
		nil,
		"/api/test",
		mockUnmarshaller,
	)

	if resource == nil {
		t.Fatal("Expected non-nil ReadOnlyRestResource even with nil client")
	}
}

// TestFactoryReturnsCorrectType tests that factory returns correct resource type
func TestFactoryReturnsCorrectType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	mockUnmarshaller := mocks.NewMockJSONUnmarshaller[*mocks.MockInstanaDataObject](ctrl)

	resource := NewRestResource[*mocks.MockInstanaDataObject](
		mockRestClient,
		"/api/test",
		rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
		mockUnmarshaller,
	)

	// Verify it implements RestResource interface
	var _ rest.RestResource[*mocks.MockInstanaDataObject] = resource
}

// TestReadOnlyFactoryReturnsCorrectType tests that read-only factory returns correct type
func TestReadOnlyFactoryReturnsCorrectType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRestClient := mocks.NewMockRestClient(ctrl)
	mockUnmarshaller := mocks.NewMockJSONUnmarshaller[*mocks.MockInstanaDataObject](ctrl)

	resource := NewReadOnlyRestResource[*mocks.MockInstanaDataObject](
		mockRestClient,
		"/api/test",
		mockUnmarshaller,
	)

	// Verify it implements ReadOnlyRestResource interface
	var _ rest.ReadOnlyRestResource[*mocks.MockInstanaDataObject] = resource
}
