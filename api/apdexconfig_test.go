package api_test

import (
	"encoding/json"
	"testing"

	"github.com/instana/instana-go-client/api"
	"github.com/instana/instana-go-client/shared/tagfilter"
)

func TestApdexEntityConstants(t *testing.T) {
	t.Run("should have correct type constants", func(t *testing.T) {
		if api.ApdexTypeApplication != "application" {
			t.Errorf("Expected ApdexTypeApplication to be 'application', got '%s'", api.ApdexTypeApplication)
		}
		if api.ApdexTypeWebsite != "website" {
			t.Errorf("Expected ApdexTypeWebsite to be 'website', got '%s'", api.ApdexTypeWebsite)
		}
	})

	t.Run("should have correct boundary scope constants", func(t *testing.T) {
		if api.BoundaryScopeAll != "ALL" {
			t.Errorf("Expected BoundaryScopeAll to be 'ALL', got '%s'", api.BoundaryScopeAll)
		}
		if api.BoundaryScopeInbound != "INBOUND" {
			t.Errorf("Expected BoundaryScopeInbound to be 'INBOUND', got '%s'", api.BoundaryScopeInbound)
		}
	})

	t.Run("should have correct beacon type constants", func(t *testing.T) {
		if api.BeaconTypeHTTPRequest != "httpRequest" {
			t.Errorf("Expected BeaconTypeHTTPRequest to be 'httpRequest', got '%s'", api.BeaconTypeHTTPRequest)
		}
		if api.BeaconTypePageLoad != "pageLoad" {
			t.Errorf("Expected BeaconTypePageLoad to be 'pageLoad', got '%s'", api.BeaconTypePageLoad)
		}
		if api.BeaconTypeCustom != "custom" {
			t.Errorf("Expected BeaconTypeCustom to be 'custom', got '%s'", api.BeaconTypeCustom)
		}
	})
}

func TestNewApplicationApdexEntity(t *testing.T) {
	t.Run("should create valid application entity", func(t *testing.T) {
		entity, err := api.NewApplicationApdexEntity(
			"app-123",
			500,
			api.BoundaryScopeAll,
			true,
			false,
			nil,
		)

		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}
		if entity.Type != api.ApdexTypeApplication {
			t.Errorf("Expected type 'application', got '%s'", entity.Type)
		}
		if entity.EntityID != "app-123" {
			t.Errorf("Expected entityId 'app-123', got '%s'", entity.EntityID)
		}
		if entity.Threshold != 500 {
			t.Errorf("Expected threshold 500, got %d", entity.Threshold)
		}
		if entity.BoundaryScope == nil || *entity.BoundaryScope != api.BoundaryScopeAll {
			t.Error("Expected boundaryScope to be 'ALL'")
		}
		if entity.IncludeInternal == nil || !*entity.IncludeInternal {
			t.Error("Expected includeInternal to be true")
		}
		if entity.IncludeSynthetic == nil || *entity.IncludeSynthetic {
			t.Error("Expected includeSynthetic to be false")
		}
	})

	t.Run("should fail with empty entityId", func(t *testing.T) {
		_, err := api.NewApplicationApdexEntity(
			"",
			500,
			api.BoundaryScopeAll,
			true,
			false,
			nil,
		)

		if err == nil {
			t.Error("Expected error for empty entityId, got nil")
		}
	})

	t.Run("should fail with threshold less than 1", func(t *testing.T) {
		_, err := api.NewApplicationApdexEntity(
			"app-123",
			0,
			api.BoundaryScopeAll,
			true,
			false,
			nil,
		)

		if err == nil {
			t.Error("Expected error for threshold < 1, got nil")
		}
	})

	t.Run("should fail with invalid boundary scope", func(t *testing.T) {
		_, err := api.NewApplicationApdexEntity(
			"app-123",
			500,
			"INVALID",
			true,
			false,
			nil,
		)

		if err == nil {
			t.Error("Expected error for invalid boundary scope, got nil")
		}
	})
}

func TestNewWebsiteApdexEntity(t *testing.T) {
	t.Run("should create valid website entity with httpRequest beacon", func(t *testing.T) {
		entity, err := api.NewWebsiteApdexEntity(
			"website-456",
			1000,
			api.BeaconTypeHTTPRequest,
			nil,
		)

		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}
		if entity.Type != api.ApdexTypeWebsite {
			t.Errorf("Expected type 'website', got '%s'", entity.Type)
		}
		if entity.EntityID != "website-456" {
			t.Errorf("Expected entityId 'website-456', got '%s'", entity.EntityID)
		}
		if entity.Threshold != 1000 {
			t.Errorf("Expected threshold 1000, got %d", entity.Threshold)
		}
		if entity.BeaconType == nil || *entity.BeaconType != api.BeaconTypeHTTPRequest {
			t.Error("Expected beaconType to be 'httpRequest'")
		}
	})

	t.Run("should create valid website entity with pageLoad beacon", func(t *testing.T) {
		entity, err := api.NewWebsiteApdexEntity(
			"website-456",
			1000,
			api.BeaconTypePageLoad,
			nil,
		)

		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}
		if entity.BeaconType == nil || *entity.BeaconType != api.BeaconTypePageLoad {
			t.Error("Expected beaconType to be 'pageLoad'")
		}
	})

	t.Run("should create valid website entity with custom beacon", func(t *testing.T) {
		entity, err := api.NewWebsiteApdexEntity(
			"website-456",
			1000,
			api.BeaconTypeCustom,
			nil,
		)

		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}
		if entity.BeaconType == nil || *entity.BeaconType != api.BeaconTypeCustom {
			t.Error("Expected beaconType to be 'custom'")
		}
	})

	t.Run("should fail with invalid beacon type", func(t *testing.T) {
		_, err := api.NewWebsiteApdexEntity(
			"website-456",
			1000,
			"invalidBeacon",
			nil,
		)

		if err == nil {
			t.Error("Expected error for invalid beacon type, got nil")
		}
	})

	t.Run("should fail with empty entityId", func(t *testing.T) {
		_, err := api.NewWebsiteApdexEntity(
			"",
			1000,
			api.BeaconTypeHTTPRequest,
			nil,
		)

		if err == nil {
			t.Error("Expected error for empty entityId, got nil")
		}
	})
}

func TestApdexEntityValidation(t *testing.T) {
	t.Run("should validate application entity correctly", func(t *testing.T) {
		boundaryScope := api.BoundaryScopeAll
		includeInternal := true
		includeSynthetic := false

		entity := &api.ApdexEntity{
			Type:             api.ApdexTypeApplication,
			EntityID:         "app-123",
			Threshold:        500,
			BoundaryScope:    &boundaryScope,
			IncludeInternal:  &includeInternal,
			IncludeSynthetic: &includeSynthetic,
		}

		if err := entity.Validate(); err != nil {
			t.Errorf("Expected valid entity, got error: %v", err)
		}
	})

	t.Run("should validate website entity correctly", func(t *testing.T) {
		beaconType := api.BeaconTypePageLoad

		entity := &api.ApdexEntity{
			Type:       api.ApdexTypeWebsite,
			EntityID:   "website-456",
			Threshold:  1000,
			BeaconType: &beaconType,
		}

		if err := entity.Validate(); err != nil {
			t.Errorf("Expected valid entity, got error: %v", err)
		}
	})

	t.Run("should fail validation for application entity with website fields", func(t *testing.T) {
		boundaryScope := api.BoundaryScopeAll
		includeInternal := true
		includeSynthetic := false
		beaconType := api.BeaconTypePageLoad

		entity := &api.ApdexEntity{
			Type:             api.ApdexTypeApplication,
			EntityID:         "app-123",
			Threshold:        500,
			BoundaryScope:    &boundaryScope,
			IncludeInternal:  &includeInternal,
			IncludeSynthetic: &includeSynthetic,
			BeaconType:       &beaconType, // Should not be set for application
		}

		if err := entity.Validate(); err == nil {
			t.Error("Expected validation error for application entity with beaconType, got nil")
		}
	})

	t.Run("should fail validation for website entity with application fields", func(t *testing.T) {
		beaconType := api.BeaconTypePageLoad
		boundaryScope := api.BoundaryScopeAll

		entity := &api.ApdexEntity{
			Type:          api.ApdexTypeWebsite,
			EntityID:      "website-456",
			Threshold:     1000,
			BeaconType:    &beaconType,
			BoundaryScope: &boundaryScope, // Should not be set for website
		}

		if err := entity.Validate(); err == nil {
			t.Error("Expected validation error for website entity with boundaryScope, got nil")
		}
	})

	t.Run("should fail validation for invalid type", func(t *testing.T) {
		entity := &api.ApdexEntity{
			Type:      "invalid",
			EntityID:  "test-123",
			Threshold: 500,
		}

		if err := entity.Validate(); err == nil {
			t.Error("Expected validation error for invalid type, got nil")
		}
	})
}

func TestApdexEntityTypeCheckers(t *testing.T) {
	t.Run("IsApplicationEntity should return true for application type", func(t *testing.T) {
		entity := &api.ApdexEntity{Type: api.ApdexTypeApplication}
		if !entity.IsApplicationEntity() {
			t.Error("Expected IsApplicationEntity to return true")
		}
		if entity.IsWebsiteEntity() {
			t.Error("Expected IsWebsiteEntity to return false")
		}
	})

	t.Run("IsWebsiteEntity should return true for website type", func(t *testing.T) {
		entity := &api.ApdexEntity{Type: api.ApdexTypeWebsite}
		if !entity.IsWebsiteEntity() {
			t.Error("Expected IsWebsiteEntity to return true")
		}
		if entity.IsApplicationEntity() {
			t.Error("Expected IsApplicationEntity to return false")
		}
	})
}

func TestApdexConfigJSONSerialization(t *testing.T) {
	t.Run("should serialize application apdex config correctly", func(t *testing.T) {
		boundaryScope := api.BoundaryScopeAll
		includeInternal := true
		includeSynthetic := false

		config := &api.ApdexConfig{
			ID:        "apdex-123",
			ApdexName: "My Application Apdex",
			ApdexEntity: api.ApdexEntity{
				Type:             api.ApdexTypeApplication,
				EntityID:         "app-123",
				Threshold:        500,
				BoundaryScope:    &boundaryScope,
				IncludeInternal:  &includeInternal,
				IncludeSynthetic: &includeSynthetic,
			},
			Tags:      []string{"production", "critical"},
			CreatedAt: 1234567890000,
		}

		jsonData, err := json.Marshal(config)
		if err != nil {
			t.Fatalf("Failed to marshal config: %v", err)
		}

		var unmarshaled api.ApdexConfig
		if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
			t.Fatalf("Failed to unmarshal config: %v", err)
		}

		if unmarshaled.ID != config.ID {
			t.Errorf("Expected ID '%s', got '%s'", config.ID, unmarshaled.ID)
		}
		if unmarshaled.ApdexEntity.Type != api.ApdexTypeApplication {
			t.Errorf("Expected type 'application', got '%s'", unmarshaled.ApdexEntity.Type)
		}
		if unmarshaled.ApdexEntity.EntityID != "app-123" {
			t.Errorf("Expected entityId 'app-123', got '%s'", unmarshaled.ApdexEntity.EntityID)
		}
	})

	t.Run("should serialize website apdex config correctly", func(t *testing.T) {
		beaconType := api.BeaconTypePageLoad

		config := &api.ApdexConfig{
			ID:        "apdex-456",
			ApdexName: "My Website Apdex",
			ApdexEntity: api.ApdexEntity{
				Type:       api.ApdexTypeWebsite,
				EntityID:   "website-456",
				Threshold:  1000,
				BeaconType: &beaconType,
			},
			Tags:      []string{"staging"},
			CreatedAt: 1234567890000,
		}

		jsonData, err := json.Marshal(config)
		if err != nil {
			t.Fatalf("Failed to marshal config: %v", err)
		}

		var unmarshaled api.ApdexConfig
		if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
			t.Fatalf("Failed to unmarshal config: %v", err)
		}

		if unmarshaled.ApdexEntity.Type != api.ApdexTypeWebsite {
			t.Errorf("Expected type 'website', got '%s'", unmarshaled.ApdexEntity.Type)
		}
		if unmarshaled.ApdexEntity.BeaconType == nil || *unmarshaled.ApdexEntity.BeaconType != api.BeaconTypePageLoad {
			t.Error("Expected beaconType to be 'pageLoad'")
		}
	})
}

func TestApdexConfigWithTagFilter(t *testing.T) {
	t.Run("should handle tag filter expression", func(t *testing.T) {
		tagFilter := tagfilter.NewStringTagFilter(
			tagfilter.TagFilterEntityNotApplicable,
			"call.type",
			"EQUALS",
			"HTTP",
		)

		entity, err := api.NewApplicationApdexEntity(
			"app-123",
			500,
			api.BoundaryScopeAll,
			true,
			false,
			tagFilter,
		)

		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}
		if entity.TagFilter == nil {
			t.Error("Expected tag filter to be set")
		}
		if entity.TagFilter.Name == nil || *entity.TagFilter.Name != "call.type" {
			t.Error("Expected tag filter name to be 'call.type'")
		}
	})
}

func TestApdexConfigGetIDForResourcePath(t *testing.T) {
	t.Run("should return correct ID", func(t *testing.T) {
		config := &api.ApdexConfig{
			ID:        "test-id-123",
			ApdexName: "Test Apdex",
		}

		id := config.GetIDForResourcePath()
		if id != "test-id-123" {
			t.Errorf("Expected ID 'test-id-123', got '%s'", id)
		}
	})
}

func TestApdexConfigResourcePath(t *testing.T) {
	t.Run("should have correct resource path", func(t *testing.T) {
		expectedPath := "/api/settings/apdex/v2"
		if api.ApdexConfigResourcePath != expectedPath {
			t.Errorf("Expected resource path '%s', got '%s'", expectedPath, api.ApdexConfigResourcePath)
		}
	})
}
