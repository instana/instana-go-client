package api

import (
	"encoding/json"
	"fmt"

	"github.com/instana/instana-go-client/shared/rest"
)

const (
	//SloCorrectionConfigResourcePath path to slo correction config resource of Instana RESTful API
	SloCorrectionConfigResourcePath = "/api/settings/correction"
)

// SloCorrectionConfig represents the REST resource of SLO Correction Configuration at Instana
type SloCorrectionConfig struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Active      bool       `json:"active"`
	Scheduling  Scheduling `json:"scheduling"`
	SloIds      []string   `json:"sloIds"`
	Tags        []string   `json:"tags"`
	CreatedDate int64      `json:"createdDate,omitempty"`
	LastUpdated int64      `json:"lastUpdated,omitempty"`
}

type DurationUnit string

const (
	DurationUnitMinute DurationUnit = "MINUTE"
	DurationUnitHour   DurationUnit = "HOUR"
	DurationUnitDay    DurationUnit = "DAY"
)

type Scheduling struct {
	StartTime     int64        `json:"startTime"` // Unix timestamp in milliseconds
	Duration      int          `json:"duration"`
	DurationUnit  DurationUnit `json:"durationUnit"`
	RecurrentRule string       `json:"recurrentRule,omitempty"`
	Recurrent     bool         `json:"recurrent"`
}

// GetIDForResourcePath implementation of the interface InstanaDataObject
func (s *SloCorrectionConfig) GetIDForResourcePath() string {
	return s.ID
}

// sloCorrectionConfigArrayResponse represents the paginated response structure from the API
type sloCorrectionConfigArrayResponse[T any] struct {
	Items     []T `json:"items"`
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	TotalHits int `json:"totalHits"`
}

// NewSloCorrectionConfigJSONUnmarshaller creates a new instance of a generic JSONUnmarshaller for SLO correction configs.
func NewSloCorrectionConfigJSONUnmarshaller[T rest.InstanaDataObject](objectType T) rest.JSONUnmarshaller[T] {
	return &sloCorrectionConfigJSONUnmarshaller[T]{
		objectType: objectType,
	}
}

type sloCorrectionConfigJSONUnmarshaller[T any] struct {
	objectType T
}

// Unmarshal unmarshals JSON data into the target object (for Get method).
func (u *sloCorrectionConfigJSONUnmarshaller[T]) Unmarshal(data []byte) (T, error) {
	target := u.objectType
	if err := json.Unmarshal(data, &target); err != nil {
		return target, fmt.Errorf("failed to parse json: %w", err)
	}
	return target, nil
}

// UnmarshalArray unmarshals JSON array data into a slice of target objects (for GetAll method).
func (u *sloCorrectionConfigJSONUnmarshaller[T]) UnmarshalArray(data []byte) (*[]T, error) {
	var response sloCorrectionConfigArrayResponse[T]
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("failed to parse json: %w", err)
	}
	return &response.Items, nil
}
