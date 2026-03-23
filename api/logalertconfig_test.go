package api

import (
	"testing"
)

func TestLogAlertConfigResourcePath(t *testing.T) {
	expected := "/api/events/settings/global-alert-configs/logs"
	if LogAlertConfigResourcePath != expected {
		t.Errorf("Expected LogAlertConfigResourcePath to be %s, got %s", expected, LogAlertConfigResourcePath)
	}
}

func TestLogAlertConfigGetIDForResourcePath(t *testing.T) {
	id := "test-id-123"
	config := LogAlertConfig{ID: id}

	result := config.GetIDForResourcePath()

	if result != id {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", id, result)
	}
}

func TestLogAlertConfigStructure(t *testing.T) {
	id := "alert-id"
	name := "Test Log Alert"
	description := "Test description"

	config := LogAlertConfig{
		ID:          id,
		Name:        name,
		Description: description,
	}

	if config.ID != id {
		t.Errorf("Expected ID to be %s, got %s", id, config.ID)
	}
	if config.Name != name {
		t.Errorf("Expected Name to be %s, got %s", name, config.Name)
	}
	if config.Description != description {
		t.Errorf("Expected Description to be %s, got %s", description, config.Description)
	}
}

func TestLogAlertRuleStructure(t *testing.T) {
	alertType := "log-alert"
	metricName := "log.count"

	rule := LogAlertRule{
		AlertType:  alertType,
		MetricName: metricName,
	}

	if rule.AlertType != alertType {
		t.Errorf("Expected AlertType to be %s, got %s", alertType, rule.AlertType)
	}
	if rule.MetricName != metricName {
		t.Errorf("Expected MetricName to be %s, got %s", metricName, rule.MetricName)
	}
}

func TestLogTimeThresholdStructure(t *testing.T) {
	thresholdType := "static"
	timeWindow := int64(60000)

	threshold := LogTimeThreshold{
		Type:       thresholdType,
		TimeWindow: timeWindow,
	}

	if threshold.Type != thresholdType {
		t.Errorf("Expected Type to be %s, got %s", thresholdType, threshold.Type)
	}
	if threshold.TimeWindow != timeWindow {
		t.Errorf("Expected TimeWindow to be %d, got %d", timeWindow, threshold.TimeWindow)
	}
}

func TestGroupByTagStructure(t *testing.T) {
	tagName := "service"
	key := "name"

	tag := GroupByTag{
		TagName: tagName,
		Key:     key,
	}

	if tag.TagName != tagName {
		t.Errorf("Expected TagName to be %s, got %s", tagName, tag.TagName)
	}
	if tag.Key != key {
		t.Errorf("Expected Key to be %s, got %s", key, tag.Key)
	}
}

func TestLogLevelsToStringSlice(t *testing.T) {
	levels := LogLevels{LogLevelWarning, LogLevelError, LogLevelAny}
	result := levels.ToStringSlice()

	if len(result) != 3 {
		t.Errorf("Expected 3 log levels, got %d", len(result))
	}
	if result[0] != "WARN" {
		t.Errorf("Expected first level to be WARN, got %s", result[0])
	}
	if result[1] != "ERROR" {
		t.Errorf("Expected second level to be ERROR, got %s", result[1])
	}
	if result[2] != "ANY" {
		t.Errorf("Expected third level to be ANY, got %s", result[2])
	}
}

// Made with Bob
