package api

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/instana/instana-go-client/shared/rest"
	"github.com/instana/instana-go-client/shared/tagfilter"
)

const (
	//SloConfigResourcePath path to sli config resource of Instana RESTful API
	SloConfigResourcePath = "/api/settings/slo"
)

// SloConfig represents the REST resource of slo configuration at Instana
type SloConfig struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Target      float64       `json:"target"`
	Tags        []string      `json:"tags"`
	Entity      SloEntity     `json:"entity"`
	Indicator   SloIndicator  `json:"indicator"`
	TimeWindow  SloTimeWindow `json:"timeWindow"`
	RbacTags    []RbacTag     `json:"rbacTags,omitempty"`
	CreatedDate int64         `json:"createdDate,omitempty"`
	LastUpdated int64         `json:"lastUpdated,omitempty"`
}

// RbacTag represents a RBAC tag in the SLO configuration
type RbacTag struct {
	DisplayName string `json:"displayName"`
	ID          string `json:"id"`
}

type SloEntity struct {
	Type             string               `json:"type"`
	ApplicationID    *string              `json:"applicationId"`
	ServiceID        *string              `json:"serviceId"`
	EndpointID       *string              `json:"endpointId"`
	BoundaryScope    *string              `json:"boundaryScope"`
	IncludeSynthetic *bool                `json:"includeSynthetic"`
	IncludeInternal  *bool                `json:"includeInternal"`
	FilterExpression *tagfilter.TagFilter `json:"tagFilterExpression"`
	WebsiteId        *string              `json:"websiteId"`
	BeaconType       *string              `json:"beaconType"`
	SyntheticTestIDs []interface{}        `json:"syntheticTestIds"`
	InfraType        *string              `json:"infraType"`
}

// SloEntity represents the nested object sli entity of the sli config REST resource at Instana
type SloApplicationEntity struct {
	Type             string               `json:"type"`
	ApplicationID    *string              `json:"applicationId"`
	ServiceID        *string              `json:"serviceId"`
	EndpointID       *string              `json:"endpointId"`
	BoundaryScope    *string              `json:"boundaryScope"`
	IncludeSynthetic *bool                `json:"includeSynthetic"`
	IncludeInternal  *bool                `json:"includeInternal"`
	FilterExpression *tagfilter.TagFilter `json:"tagFilterExpression"`
}

type SloWebsiteEntity struct {
	Type             string               `json:"type"`
	WebsiteId        *string              `json:"websiteId"`
	BeaconType       *string              `json:"beaconType"`
	FilterExpression *tagfilter.TagFilter `json:"tagFilterExpression"`
}

type SloInfraEntity struct {
	Type      string `json:"type"`
	InfraType string `json:"infraType"`
}

type SloSyntheticEntity struct {
	Type             string               `json:"type"`
	SyntheticTestIDs []interface{}        `json:"syntheticTestIds"`
	FilterExpression *tagfilter.TagFilter `json:"tagFilterExpression"`
}

type SloIndicator struct {
	Blueprint                 string               `json:"blueprint"`
	Type                      string               `json:"type"`
	Threshold                 float64              `json:"threshold"`
	Aggregation               *string              `json:"aggregation"`
	Operator                  *string              `json:"operator"`
	TrafficType               *string              `json:"trafficType"`
	MetricName                *string              `json:"metricName,omitempty"`
	GoodEventFilterExpression *tagfilter.TagFilter `json:"goodEventsFilter"`
	BadEventFilterExpression  *tagfilter.TagFilter `json:"badEventsFilter"`
}

// Blueprints
type SloTimeBasedLatencyIndicator struct {
	Blueprint   string  `json:"blueprint"`
	Type        string  `json:"type"`
	Threshold   float64 `json:"threshold"`
	Aggregation string  `json:"aggregation"`
}

type SloTimeBasedAvailabilityIndicator struct {
	Blueprint   string  `json:"blueprint"`
	Type        string  `json:"type"`
	Threshold   float64 `json:"threshold"`
	Aggregation string  `json:"aggregation"`
}

type SloTrafficIndicator struct {
	Blueprint   string  `json:"blueprint"`
	TrafficType string  `json:"trafficType"`
	Threshold   float64 `json:"threshold"`
	Aggregation string  `json:"aggregation"`
}

type SloEventBasedLatencyIndicator struct {
	Blueprint string  `json:"blueprint"`
	Type      string  `json:"type"`
	Threshold float64 `json:"threshold"`
}

type SloEventBasedAvailabilityIndicator struct {
	Blueprint string `json:"blueprint"`
	Type      string `json:"type"`
}

type SloCustomIndicator struct {
	Type                      string               `json:"type"`
	Blueprint                 string               `json:"blueprint"`
	GoodEventFilterExpression *tagfilter.TagFilter `json:"goodEventsFilter"`
	BadEventFilterExpression  *tagfilter.TagFilter `json:"badEventsFilter"`
}

type SloTimeWindow struct {
	Type         string  `json:"type"`
	Duration     int     `json:"duration"`
	DurationUnit string  `json:"durationUnit"`
	Timezone     string  `json:"timezone,omitempty"`
	StartTime    float64 `json:"startTimestamp"`
}

// time windows
type SloRollingTimeWindow struct {
	Type         string `json:"type"`
	Duration     int    `json:"duration"`
	DurationUnit string `json:"durationUnit"`
	Timezone     string `json:"timezone,omitempty"`
}

type SloFixedTimeWindow struct {
	Type         string  `json:"type"`
	Duration     int     `json:"duration"`
	DurationUnit string  `json:"durationUnit"`
	Timezone     string  `json:"timezone,omitempty"`
	StartTime    float64 `json:"startTimestamp"`
}

// GetIDForResourcePath implementation of the interface InstanaDataObject
func (s *SloConfig) GetIDForResourcePath() string {
	fmt.Fprintln(os.Stderr, ">> GetIDForResourcePath: "+s.ID)
	return s.ID
}

// sloConfigArrayResponse represents the paginated response structure from the API
type sloConfigArrayResponse[T any] struct {
	Items     []T `json:"items"`
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	TotalHits int `json:"totalHits"`
}

// NewSloConfigJSONUnmarshaller creates a new instance of a generic JSONUnmarshaller for SLO configs.
func NewSloConfigJSONUnmarshaller[T rest.InstanaDataObject](objectType T) rest.JSONUnmarshaller[T] {
	return &sloConfigJSONUnmarshaller[T]{
		objectType: objectType,
	}
}

type sloConfigJSONUnmarshaller[T any] struct {
	objectType T
}

// Unmarshal unmarshals JSON data into the target object (for Get method).
func (u *sloConfigJSONUnmarshaller[T]) Unmarshal(data []byte) (T, error) {
	// Create a new instance to avoid shared state issues
	var target T
	if err := json.Unmarshal(data, &target); err != nil {
		var zero T
		return zero, fmt.Errorf("failed to parse json: %w", err)
	}
	return target, nil
}

// UnmarshalArray unmarshals JSON array data into a slice of target objects (for GetAll method).
func (u *sloConfigJSONUnmarshaller[T]) UnmarshalArray(data []byte) (*[]T, error) {
	var response sloConfigArrayResponse[T]
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("failed to parse json: %w", err)
	}
	return &response.Items, nil
}
