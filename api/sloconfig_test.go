package api_test

import (
	"encoding/json"
	"testing"

	. "github.com/instana/instana-go-client/api"
	"github.com/instana/instana-go-client/shared/tagfilter"
	"github.com/stretchr/testify/require"
)

func TestSloConfigResourcePath(t *testing.T) {
	expected := "/api/settings/slo"
	if SloConfigResourcePath != expected {
		t.Errorf("Expected SloConfigResourcePath to be %s, got %s", expected, SloConfigResourcePath)
	}
}

func TestSloConfigGetIDForResourcePath(t *testing.T) {
	id := "test-slo-id"
	config := SloConfig{ID: id}

	result := config.GetIDForResourcePath()

	if result != id {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", id, result)
	}
}

func TestSloConfigStructure(t *testing.T) {
	id := "slo-id"
	name := "Test SLO"

	config := SloConfig{
		ID:   id,
		Name: name,
	}

	if config.ID != id {
		t.Errorf("Expected ID to be %s, got %s", id, config.ID)
	}
	if config.Name != name {
		t.Errorf("Expected Name to be %s, got %s", name, config.Name)
	}
}

func TestNewSloAgentJSONUnmarshaller(t *testing.T) {
	testData := &testObject{
		ID:   defaultObjectId,
		Name: defaultObjectName,
	}
	testObjects := []*testObject{testData, testData}

	serializedJSON, _ := json.Marshal(testObjects)

	sut := NewSloConfigJSONUnmarshaller(&testObject{})

	_, err := sut.Unmarshal(serializedJSON)

	require.Error(t, err)
}

func TestShouldSuccessfullyUnmarshalSloArrayOfObjects(t *testing.T) {
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

	sut := NewSloConfigJSONUnmarshaller(&testObject{})

	result, err := sut.UnmarshalArray(serializedJSON)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, len(testObjects), len(*result))
	require.Equal(t, testObjects[0].ID, (*result)[0].ID)
	require.Equal(t, testObjects[0].Name, (*result)[0].Name)
}

// ---- Blueprint constants ----

func TestSloBlueprintConstants(t *testing.T) {
	require.Equal(t, "latency", SloBlueprintLatency)
	require.Equal(t, "availability", SloBlueprintAvailability)
	require.Equal(t, "traffic", SloBlueprintTraffic)
	require.Equal(t, "saturation", SloBlueprintSaturation)
	require.Equal(t, "custom", SloBlueprintCustom)
	require.Equal(t, "advanced-custom", SloBlueprintAdvancedCustom)
}

// ---- SloMobileEntity ----

func TestSloMobileEntityJSONRoundTrip(t *testing.T) {
	raw := `{
		"type": "mobile",
		"mobileIds": ["app-1", "app-2"],
		"tagFilterExpression": {
			"type": "EXPRESSION",
			"logicalOperator": "AND",
			"elements": [],
			"entity": null,
			"name": null,
			"operator": null,
			"booleanValue": null,
			"numberValue": null,
			"stringValue": null,
			"key": null,
			"value": null
		}
	}`

	var entity SloMobileEntity
	require.NoError(t, json.Unmarshal([]byte(raw), &entity))

	require.Equal(t, "mobile", entity.Type)
	require.Equal(t, []string{"app-1", "app-2"}, entity.MobileIds)
	require.NotNil(t, entity.FilterExpression)
	require.Equal(t, tagfilter.TagFilterExpressionType, entity.FilterExpression.Type)

	marshaled, err := json.Marshal(entity)
	require.NoError(t, err)
	require.JSONEq(t, raw, string(marshaled))
}

func TestSloMobileEntityNilMobileIdsSerialiseAsNull(t *testing.T) {
	entity := SloMobileEntity{Type: "mobile"}
	data, err := json.Marshal(entity)
	require.NoError(t, err)
	require.Contains(t, string(data), `"mobileIds":null`)
}

func TestNewSloMobileEntityDefaultsEmptyAndExpression(t *testing.T) {
	entity := NewSloMobileEntity([]string{"app-1"}, nil)

	require.Equal(t, "mobile", entity.Type)
	require.Equal(t, []string{"app-1"}, entity.MobileIds)
	require.NotNil(t, entity.FilterExpression)

	data, err := json.Marshal(entity)
	require.NoError(t, err)
	require.Contains(t, string(data), `"tagFilterExpression"`)
	require.Contains(t, string(data), `"type":"EXPRESSION"`)
	require.Contains(t, string(data), `"logicalOperator":"AND"`)
}

func TestNewSloMobileEntityPreservesProvidedFilter(t *testing.T) {
	filter := tagfilter.NewLogicalAndTagFilter([]*tagfilter.TagFilter{})
	entity := NewSloMobileEntity([]string{"app-1"}, filter)

	require.Equal(t, filter, entity.FilterExpression)
}

// ---- SloEntityMetricScope ----

func TestSloEntityMetricScopeJSONRoundTrip(t *testing.T) {
	// tagFilterExpression is always present; null when no filter is set
	raw := `{"type":"httpRequests","tagFilterExpression":null}`

	var scope SloEntityMetricScope
	require.NoError(t, json.Unmarshal([]byte(raw), &scope))

	require.Equal(t, "httpRequests", scope.Type)
	require.Nil(t, scope.FilterExpression)

	marshaled, err := json.Marshal(scope)
	require.NoError(t, err)
	require.JSONEq(t, raw, string(marshaled))
}

// ---- SloEntityMetric ----

func TestSloEntityMetricJSONRoundTrip(t *testing.T) {
	raw := `{"name":"duration","scope":{"type":"httpRequests","tagFilterExpression":null}}`

	var metric SloEntityMetric
	require.NoError(t, json.Unmarshal([]byte(raw), &metric))

	require.Equal(t, "duration", metric.Name)
	require.NotNil(t, metric.Scope)
	require.Equal(t, "httpRequests", metric.Scope.Type)

	marshaled, err := json.Marshal(metric)
	require.NoError(t, err)
	require.JSONEq(t, raw, string(marshaled))
}

func TestSloEntityMetricNilScopeSerialiseAsNull(t *testing.T) {
	metric := SloEntityMetric{Name: "duration"}
	data, err := json.Marshal(metric)
	require.NoError(t, err)
	// scope is always present in the payload, serialised as null when not set
	require.Contains(t, string(data), `"scope":null`)
}

// ---- SloAdvancedFilter ----

func TestSloAdvancedFilterJSONRoundTrip(t *testing.T) {
	raw := `{"aggregation":"MEAN","threshold":500.0,"operator":"<","metric":{"name":"duration","scope":{"type":"httpRequests","tagFilterExpression":null}}}`

	var filter SloAdvancedFilter
	require.NoError(t, json.Unmarshal([]byte(raw), &filter))

	require.Equal(t, "MEAN", filter.Aggregation)
	require.Equal(t, 500.0, filter.Threshold)
	require.Equal(t, "<", filter.Operator)
	require.NotNil(t, filter.Metric)
	require.Equal(t, "duration", filter.Metric.Name)

	marshaled, err := json.Marshal(filter)
	require.NoError(t, err)
	require.JSONEq(t, raw, string(marshaled))
}

// ---- SloAdvancedCustomIndicator ----

func TestSloAdvancedCustomIndicatorJSONRoundTrip(t *testing.T) {
	raw := `{
		"blueprint":"advanced-custom",
		"type":"eventBased",
		"goodEvents":{"aggregation":"MEAN","threshold":500.0,"operator":"<","metric":{"name":"duration","scope":{"type":"httpRequests","tagFilterExpression":null}}},
		"badEvents":{"aggregation":"MEAN","threshold":500.0,"operator":">=","metric":{"name":"duration","scope":{"type":"httpRequests","tagFilterExpression":null}}}
	}`

	var indicator SloAdvancedCustomIndicator
	require.NoError(t, json.Unmarshal([]byte(raw), &indicator))

	require.Equal(t, SloBlueprintAdvancedCustom, indicator.Blueprint)
	require.Equal(t, "eventBased", indicator.Type)
	require.NotNil(t, indicator.GoodEvents)
	require.Equal(t, "<", indicator.GoodEvents.Operator)
	require.NotNil(t, indicator.BadEvents)
	require.Equal(t, ">=", indicator.BadEvents.Operator)

	marshaled, err := json.Marshal(indicator)
	require.NoError(t, err)
	require.JSONEq(t, raw, string(marshaled))
}

func TestSloAdvancedCustomIndicatorNilBadEventsSerialiseAsNull(t *testing.T) {
	indicator := SloAdvancedCustomIndicator{
		Blueprint: SloBlueprintAdvancedCustom,
		Type:      "eventBased",
		GoodEvents: &SloAdvancedFilter{
			Aggregation: "MEAN",
			Threshold:   500.0,
			Operator:    "<",
		},
	}
	data, err := json.Marshal(indicator)
	require.NoError(t, err)
	// badEvents is always present in the payload, serialised as null when not set
	require.Contains(t, string(data), `"badEvents":null`)
}

// ---- SloSaturationIndicator ----

func TestSloSaturationIndicatorJSONRoundTrip(t *testing.T) {
	metricName := "system.cpu.user"
	operator := ">="
	// metric is always present in the payload, serialised as null when not set
	raw := `{"blueprint":"saturation","type":"timeBased","metricName":"system.cpu.user","threshold":80.0,"aggregation":"MEAN","operator":">=","metric":null}`

	var indicator SloSaturationIndicator
	require.NoError(t, json.Unmarshal([]byte(raw), &indicator))

	require.Equal(t, SloBlueprintSaturation, indicator.Blueprint)
	require.Equal(t, "timeBased", indicator.Type)
	require.Equal(t, &metricName, indicator.MetricName)
	require.Equal(t, 80.0, indicator.Threshold)
	require.Equal(t, "MEAN", indicator.Aggregation)
	require.Equal(t, &operator, indicator.Operator)
	require.Nil(t, indicator.Metric)

	marshaled, err := json.Marshal(indicator)
	require.NoError(t, err)
	require.JSONEq(t, raw, string(marshaled))
}

func TestSloSaturationIndicatorWithEntityMetric(t *testing.T) {
	raw := `{"blueprint":"saturation","type":"timeBased","threshold":80.0,"aggregation":"MEAN","metric":{"name":"crashRate","scope":{"type":"crashes"}}}`

	var indicator SloSaturationIndicator
	require.NoError(t, json.Unmarshal([]byte(raw), &indicator))

	require.Equal(t, SloBlueprintSaturation, indicator.Blueprint)
	require.NotNil(t, indicator.Metric)
	require.Equal(t, "crashRate", indicator.Metric.Name)
	require.Equal(t, "crashes", indicator.Metric.Scope.Type)
}

// ---- SloTrafficIndicator with new fields ----

func TestSloTrafficIndicatorWithOperatorAndMetric(t *testing.T) {
	operator := ">="
	raw := `{"blueprint":"traffic","trafficType":"CALLS_TOTAL","threshold":1000.0,"aggregation":"SUM","operator":">=","metric":{"name":"requestCount","scope":{"type":"httpRequests","tagFilterExpression":null}}}`

	var indicator SloTrafficIndicator
	require.NoError(t, json.Unmarshal([]byte(raw), &indicator))

	require.Equal(t, SloBlueprintTraffic, indicator.Blueprint)
	require.Equal(t, "CALLS_TOTAL", indicator.TrafficType)
	require.Equal(t, 1000.0, indicator.Threshold)
	require.Equal(t, &operator, indicator.Operator)
	require.NotNil(t, indicator.Metric)
	require.Equal(t, "requestCount", indicator.Metric.Name)

	marshaled, err := json.Marshal(indicator)
	require.NoError(t, err)
	require.JSONEq(t, raw, string(marshaled))
}

// ---- SloIndicator flat struct — Metric, GoodEvents, BadEvents ----

func TestSloIndicatorWithMetricField(t *testing.T) {
	raw := `{"blueprint":"latency","type":"timeBased","threshold":500.0,"aggregation":"MEAN","operator":"<","metric":{"name":"duration","scope":{"type":"httpRequests"}}}`

	var indicator SloIndicator
	require.NoError(t, json.Unmarshal([]byte(raw), &indicator))

	require.Equal(t, SloBlueprintLatency, indicator.Blueprint)
	require.NotNil(t, indicator.Metric)
	require.Equal(t, "duration", indicator.Metric.Name)
	require.Equal(t, "httpRequests", indicator.Metric.Scope.Type)
	require.Nil(t, indicator.GoodEvents)
	require.Nil(t, indicator.BadEvents)
}

func TestSloIndicatorWithAdvancedCustomFields(t *testing.T) {
	raw := `{
		"blueprint":"advanced-custom",
		"type":"eventBased",
		"threshold":0,
		"goodEvents":{"aggregation":"MEAN","threshold":500.0,"operator":"<","metric":{"name":"duration","scope":{"type":"httpRequests"}}},
		"badEvents":{"aggregation":"MEAN","threshold":500.0,"operator":">=","metric":{"name":"duration","scope":{"type":"httpRequests"}}}
	}`

	var indicator SloIndicator
	require.NoError(t, json.Unmarshal([]byte(raw), &indicator))

	require.Equal(t, SloBlueprintAdvancedCustom, indicator.Blueprint)
	require.NotNil(t, indicator.GoodEvents)
	require.Equal(t, "duration", indicator.GoodEvents.Metric.Name)
	require.NotNil(t, indicator.BadEvents)
	require.Equal(t, ">=", indicator.BadEvents.Operator)
}

// ---- SloConfig end-to-end with mobile entity ----

func TestSloConfigWithMobileEntityJSONRoundTrip(t *testing.T) {
	raw := `{
		"id":"SLO-mobile-1",
		"name":"Mobile SLO",
		"target":0.99,
		"tags":["mobile"],
		"entity":{"type":"mobile","mobileIds":["app-1"]},
		"indicator":{"blueprint":"latency","type":"timeBased","threshold":500.0,"aggregation":"MEAN","operator":"<","metric":{"name":"duration","scope":{"type":"httpRequests"}}},
		"timeWindow":{"type":"rolling","duration":7,"durationUnit":"day","startTimestamp":0}
	}`

	var config SloConfig
	require.NoError(t, json.Unmarshal([]byte(raw), &config))

	require.Equal(t, "SLO-mobile-1", config.ID)
	require.Equal(t, "mobile", config.Entity.Type)
	require.Equal(t, []string{"app-1"}, config.Entity.MobileIds)
	require.Equal(t, SloBlueprintLatency, config.Indicator.Blueprint)
	require.NotNil(t, config.Indicator.Metric)
	require.Equal(t, "duration", config.Indicator.Metric.Name)
}
