package api

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/instana/instana-go-client/shared/rest"
	"github.com/instana/instana-go-client/shared/tagfilter"
)

const (
	// SloConfigResourcePath path to sli config resource of Instana RESTful API
	SloConfigResourcePath = "/api/settings/slo"

	// Blueprint type name constants — match BlueprintType enum in the backend
	SloBlueprintLatency        = "latency"
	SloBlueprintAvailability   = "availability"
	SloBlueprintTraffic        = "traffic"
	SloBlueprintSaturation     = "saturation"
	SloBlueprintCustom         = "custom"
	SloBlueprintAdvancedCustom = "advanced-custom"
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
	Type                          string               `json:"type"`
	ApplicationID                 *string              `json:"applicationId"`
	ServiceID                     *string              `json:"serviceId"`
	EndpointID                    *string              `json:"endpointId"`
	BoundaryScope                 *string              `json:"boundaryScope"`
	IncludeSynthetic              *bool                `json:"includeSynthetic"`
	IncludeInternal               *bool                `json:"includeInternal"`
	FilterExpression              *tagfilter.TagFilter `json:"tagFilterExpression"`
	WebsiteId                     *string              `json:"websiteId"`
	BeaconType                    *string              `json:"beaconType"`
	SyntheticTestIDs              []interface{}        `json:"syntheticTestIds"`
	IncludeUnscheduledTestResults *bool                `json:"includeUnscheduledTestResults"`
	InfraType                     *string              `json:"infraType"`
	MobileIds                     []string             `json:"mobileIds"`
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

type SloMobileEntity struct {
	Type             string               `json:"type"`
	MobileIds        []string             `json:"mobileIds"`
	FilterExpression *tagfilter.TagFilter `json:"tagFilterExpression"`
}

// NewSloMobileEntity creates a SloMobileEntity, defaulting FilterExpression to an empty
// AND expression when nil — matching the backend SloEntity constructor behaviour.
func NewSloMobileEntity(mobileIds []string, filterExpression *tagfilter.TagFilter) SloMobileEntity {
	if filterExpression == nil {
		filterExpression = tagfilter.NewLogicalAndTagFilter([]*tagfilter.TagFilter{})
	}
	return SloMobileEntity{
		Type:             "mobile",
		MobileIds:        mobileIds,
		FilterExpression: filterExpression,
	}
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
	MetricName                *string              `json:"metricName"`
	GoodEventFilterExpression *tagfilter.TagFilter `json:"goodEventsFilter"`
	BadEventFilterExpression  *tagfilter.TagFilter `json:"badEventsFilter"`
	Metric                    *SloEntityMetric     `json:"metric"`
	// GoodEvents and BadEvents are populated when Blueprint == SloBlueprintAdvancedCustom
	GoodEvents *SloAdvancedFilter `json:"goodEvents"`
	BadEvents  *SloAdvancedFilter `json:"badEvents"`
}

// Blueprints
type SloTimeBasedLatencyIndicator struct {
	Blueprint   string           `json:"blueprint"`
	Type        string           `json:"type"`
	Threshold   float64          `json:"threshold"`
	Aggregation string           `json:"aggregation"`
	Operator    *string          `json:"operator"`
	Metric      *SloEntityMetric `json:"metric"`
}

type SloTimeBasedAvailabilityIndicator struct {
	Blueprint   string           `json:"blueprint"`
	Type        string           `json:"type"`
	Threshold   float64          `json:"threshold"`
	Aggregation string           `json:"aggregation"`
	Operator    *string          `json:"operator"`
	Metric      *SloEntityMetric `json:"metric"`
}

type SloTrafficIndicator struct {
	Blueprint   string           `json:"blueprint"`
	TrafficType string           `json:"trafficType"`
	Threshold   float64          `json:"threshold"`
	Aggregation string           `json:"aggregation"`
	Operator    *string          `json:"operator"`
	Metric      *SloEntityMetric `json:"metric"`
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

// SloSaturationIndicator represents the saturation blueprint indicator
// (maps to SaturationBlueprintIndicator). MetricName identifies the infra/custom metric;
// Metric carries the entity-metric scope used for mobile/advanced scenarios.
type SloSaturationIndicator struct {
	Blueprint   string           `json:"blueprint"`
	Type        string           `json:"type"`
	MetricName  *string          `json:"metricName"`
	Threshold   float64          `json:"threshold"`
	Aggregation string           `json:"aggregation"`
	Operator    *string          `json:"operator"`
	Metric      *SloEntityMetric `json:"metric"`
}

// SloEntityMetricScope represents the scope of an entity metric (maps to EntityMetricScope).
// For a Mobile App SLO, Type is the beacon type (e.g. "httpRequests", "crashes").
type SloEntityMetricScope struct {
	Type             string               `json:"type"`
	FilterExpression *tagfilter.TagFilter `json:"tagFilterExpression"`
}

// SloEntityMetric represents the metric definition for an indicator (maps to EntityMetric).
// Name is the metric name and Scope narrows which beacons/signals are measured.
type SloEntityMetric struct {
	Name  string                `json:"name"`
	Scope *SloEntityMetricScope `json:"scope"`
}

// SloAdvancedFilter represents one side (good or bad) of an advanced-custom indicator
// (maps to AdvancedFilter). The Metric field identifies the signal; Aggregation, Threshold
// and Operator define the pass/fail condition.
type SloAdvancedFilter struct {
	Aggregation string           `json:"aggregation"`
	Threshold   float64          `json:"threshold"`
	Operator    string           `json:"operator"`
	Metric      *SloEntityMetric `json:"metric"`
}

// SloAdvancedCustomIndicator represents the advanced-custom blueprint indicator
// (maps to AdvancedCustomBlueprintIndicator). GoodEvents and BadEvents each carry
// their own metric, aggregation, threshold and operator.
type SloAdvancedCustomIndicator struct {
	Blueprint  string             `json:"blueprint"`
	Type       string             `json:"type"`
	GoodEvents *SloAdvancedFilter `json:"goodEvents"`
	BadEvents  *SloAdvancedFilter `json:"badEvents"`
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
