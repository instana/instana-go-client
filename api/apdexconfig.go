package api

import (
	"encoding/json"
	"fmt"

	"github.com/instana/instana-go-client/shared/rest"
	"github.com/instana/instana-go-client/shared/tagfilter"
)

const (
	// ApdexConfigResourcePath path to apdex config resource of Instana RESTful API
	ApdexConfigResourcePath = "/api/settings/apdex/v2"

	// ApdexTypeApplication represents application apdex entity type
	ApdexTypeApplication = "application"
	// ApdexTypeWebsite represents website apdex entity type
	ApdexTypeWebsite = "website"

	// BoundaryScopeAll represents ALL boundary scope for application entities
	BoundaryScopeAll = "ALL"
	// BoundaryScopeInbound represents INBOUND boundary scope for application entities
	BoundaryScopeInbound = "INBOUND"

	// BeaconTypeHTTPRequest represents httpRequest beacon type for website entities
	BeaconTypeHTTPRequest = "httpRequest"
	// BeaconTypePageLoad represents pageLoad beacon type for website entities
	BeaconTypePageLoad = "pageLoad"
	// BeaconTypeCustom represents custom beacon type for website entities
	BeaconTypeCustom = "custom"
)

// ApdexConfig represents the REST resource of apdex configuration at Instana
type ApdexConfig struct {
	ID          string      `json:"id"`
	ApdexName   string      `json:"apdexName"`
	ApdexEntity ApdexEntity `json:"apdexEntity"`
	Tags        []string    `json:"tags"`
	CreatedAt   int64       `json:"createdAt,omitempty"`
	LastUpdated int64       `json:"lastUpdated,omitempty"`
	RbacTags    []RbacTag   `json:"rbacTags,omitempty"`
}

// ApdexEntity represents the polymorphic apdex entity with discriminator field "apdexType"
// This structure supports both application and website entity types following the backend design
type ApdexEntity struct {
	Type      string               `json:"apdexType"`
	EntityID  string               `json:"entityId"`
	Threshold int                  `json:"threshold"`
	TagFilter *tagfilter.TagFilter `json:"tagFilterExpression,omitempty"`

	// Application-specific fields (only valid when Type == "application")
	BoundaryScope    *string `json:"boundaryScope,omitempty"`
	IncludeInternal  *bool   `json:"includeInternal,omitempty"`
	IncludeSynthetic *bool   `json:"includeSynthetic,omitempty"`

	// Website-specific fields (only valid when Type == "website")
	BeaconType *string `json:"beaconType,omitempty"`
}

// GetIDForResourcePath implementation of the interface InstanaDataObject
func (a *ApdexConfig) GetIDForResourcePath() string {
	return a.ID
}

// apdexConfigArrayResponse represents the paginated response structure from the API
type apdexConfigArrayResponse[T any] struct {
	Items     []T `json:"items"`
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	TotalHits int `json:"totalHits"`
}

// NewApdexConfigJSONUnmarshaller creates a new instance of a generic JSONUnmarshaller for Apdex configs.
func NewApdexConfigJSONUnmarshaller[T rest.InstanaDataObject](objectType T) rest.JSONUnmarshaller[T] {
	return &apdexConfigJSONUnmarshaller[T]{
		objectType: objectType,
	}
}

type apdexConfigJSONUnmarshaller[T any] struct {
	objectType T
}

// Unmarshal unmarshals JSON data into the target object (for Get method).
func (u *apdexConfigJSONUnmarshaller[T]) Unmarshal(data []byte) (T, error) {
	// Create a new instance to avoid shared state issues
	var target T
	if err := json.Unmarshal(data, &target); err != nil {
		var zero T
		return zero, fmt.Errorf("failed to parse json: %w", err)
	}
	return target, nil
}

// UnmarshalArray unmarshals JSON array data into a slice of target objects (for GetAll method).
func (u *apdexConfigJSONUnmarshaller[T]) UnmarshalArray(data []byte) (*[]T, error) {
	var response apdexConfigArrayResponse[T]
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("failed to parse json: %w", err)
	}
	return &response.Items, nil
}


// Validate validates the ApdexEntity structure based on its type
func (e *ApdexEntity) Validate() error {
	// Validate common required fields
	if e.EntityID == "" {
		return fmt.Errorf("entityId is required")
	}
	if e.Threshold < 1 {
		return fmt.Errorf("threshold must be at least 1, got %d", e.Threshold)
	}

	// Validate type-specific fields
	switch e.Type {
	case ApdexTypeApplication:
		return e.validateApplicationEntity()
	case ApdexTypeWebsite:
		return e.validateWebsiteEntity()
	default:
		return fmt.Errorf("invalid apdexType: %s, must be either '%s' or '%s'", e.Type, ApdexTypeApplication, ApdexTypeWebsite)
	}
}

// validateApplicationEntity validates application-specific fields
func (e *ApdexEntity) validateApplicationEntity() error {
	if e.BoundaryScope == nil {
		return fmt.Errorf("boundaryScope is required for application entities")
	}
	if *e.BoundaryScope != BoundaryScopeAll && *e.BoundaryScope != BoundaryScopeInbound {
		return fmt.Errorf("boundaryScope must be '%s' or '%s', got '%s'", BoundaryScopeAll, BoundaryScopeInbound, *e.BoundaryScope)
	}
	if e.IncludeInternal == nil {
		return fmt.Errorf("includeInternal is required for application entities")
	}
	if e.IncludeSynthetic == nil {
		return fmt.Errorf("includeSynthetic is required for application entities")
	}
	// Website-specific fields should not be set
	if e.BeaconType != nil {
		return fmt.Errorf("beaconType should not be set for application entities")
	}
	return nil
}

// validateWebsiteEntity validates website-specific fields
func (e *ApdexEntity) validateWebsiteEntity() error {
	if e.BeaconType == nil {
		return fmt.Errorf("beaconType is required for website entities")
	}
	validBeacons := []string{BeaconTypeHTTPRequest, BeaconTypePageLoad, BeaconTypeCustom}
	isValid := false
	for _, valid := range validBeacons {
		if *e.BeaconType == valid {
			isValid = true
			break
		}
	}
	if !isValid {
		return fmt.Errorf("beaconType must be one of [%s, %s, %s], got '%s'", BeaconTypeHTTPRequest, BeaconTypePageLoad, BeaconTypeCustom, *e.BeaconType)
	}
	// Application-specific fields should not be set
	if e.BoundaryScope != nil {
		return fmt.Errorf("boundaryScope should not be set for website entities")
	}
	if e.IncludeInternal != nil {
		return fmt.Errorf("includeInternal should not be set for website entities")
	}
	if e.IncludeSynthetic != nil {
		return fmt.Errorf("includeSynthetic should not be set for website entities")
	}
	return nil
}

// NewApplicationApdexEntity creates a new application apdex entity with validation
func NewApplicationApdexEntity(entityID string, threshold int, boundaryScope string, includeInternal, includeSynthetic bool, tagFilter *tagfilter.TagFilter) (*ApdexEntity, error) {
	entity := &ApdexEntity{
		Type:             ApdexTypeApplication,
		EntityID:         entityID,
		Threshold:        threshold,
		TagFilter:        tagFilter,
		BoundaryScope:    &boundaryScope,
		IncludeInternal:  &includeInternal,
		IncludeSynthetic: &includeSynthetic,
	}
	if err := entity.Validate(); err != nil {
		return nil, fmt.Errorf("invalid application apdex entity: %w", err)
	}
	return entity, nil
}

// NewWebsiteApdexEntity creates a new website apdex entity with validation
func NewWebsiteApdexEntity(entityID string, threshold int, beaconType string, tagFilter *tagfilter.TagFilter) (*ApdexEntity, error) {
	entity := &ApdexEntity{
		Type:       ApdexTypeWebsite,
		EntityID:   entityID,
		Threshold:  threshold,
		TagFilter:  tagFilter,
		BeaconType: &beaconType,
	}
	if err := entity.Validate(); err != nil {
		return nil, fmt.Errorf("invalid website apdex entity: %w", err)
	}
	return entity, nil
}

// IsApplicationEntity returns true if this is an application entity
func (e *ApdexEntity) IsApplicationEntity() bool {
	return e.Type == ApdexTypeApplication
}

// IsWebsiteEntity returns true if this is a website entity
func (e *ApdexEntity) IsWebsiteEntity() bool {
	return e.Type == ApdexTypeWebsite
}
