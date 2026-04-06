package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
)

func TestSliConfigResourcePath(t *testing.T) {
	expected := "/api/settings/v2/sli"
	if SliConfigResourcePath != expected {
		t.Errorf("Expected SliConfigResourcePath to be %s, got %s", expected, SliConfigResourcePath)
	}
}

func TestSliConfigGetIDForResourcePath(t *testing.T) {
	id := "test-sli-id"
	config := SliConfig{ID: id}

	result := config.GetIDForResourcePath()

	if result != id {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", id, result)
	}
}

func TestSliConfigStructure(t *testing.T) {
	id := "sli-id"
	name := "Test SLI"

	config := SliConfig{
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

func TestMetricConfigurationStructure(t *testing.T) {
	name := "response.time"
	aggregation := "sum"
	threshold := 100.5

	metric := MetricConfiguration{
		Name:        name,
		Aggregation: aggregation,
		Threshold:   threshold,
	}

	if metric.Name != name {
		t.Errorf("Expected Name to be %s, got %s", name, metric.Name)
	}
	if metric.Aggregation != aggregation {
		t.Errorf("Expected Aggregation to be %s, got %s", aggregation, metric.Aggregation)
	}
	if metric.Threshold != threshold {
		t.Errorf("Expected Threshold to be %f, got %f", threshold, metric.Threshold)
	}
}

func TestSliEntityStructure(t *testing.T) {
	sliType := "availability"
	appID := "app-123"

	entity := SliEntity{
		Type:          sliType,
		ApplicationID: &appID,
	}

	if entity.Type != sliType {
		t.Errorf("Expected Type to be %s, got %s", sliType, entity.Type)
	}
	if *entity.ApplicationID != appID {
		t.Errorf("Expected ApplicationID to be %s, got %s", appID, *entity.ApplicationID)
	}
}
