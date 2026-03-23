package api

import (
	"testing"
)

func TestSyntheticTestResourcePath(t *testing.T) {
	expected := "/api/synthetics/settings/tests"
	if SyntheticTestResourcePath != expected {
		t.Errorf("Expected SyntheticTestResourcePath to be %s, got %s", expected, SyntheticTestResourcePath)
	}
}

func TestSyntheticTestGetIDForResourcePath(t *testing.T) {
	id := "test-synthetic-id"
	test := SyntheticTest{ID: id}

	result := test.GetIDForResourcePath()

	if result != id {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", id, result)
	}
}

func TestSyntheticTestStructure(t *testing.T) {
	id := "synthetic-test-id"
	label := "Test Synthetic"
	active := true

	test := SyntheticTest{
		ID:     id,
		Label:  label,
		Active: active,
	}

	if test.ID != id {
		t.Errorf("Expected ID to be %s, got %s", id, test.ID)
	}
	if test.Label != label {
		t.Errorf("Expected Label to be %s, got %s", label, test.Label)
	}
	if test.Active != active {
		t.Errorf("Expected Active to be %v, got %v", active, test.Active)
	}
}

func TestApiTagStructure(t *testing.T) {
	name := "environment"
	value := "production"

	tag := ApiTag{
		Name:  name,
		Value: value,
	}

	if tag.Name != name {
		t.Errorf("Expected Name to be %s, got %s", name, tag.Name)
	}
	if tag.Value != value {
		t.Errorf("Expected Value to be %s, got %s", value, tag.Value)
	}
}

func TestDNSFilterQueryTimeStructure(t *testing.T) {
	key := "queryTime"
	operator := ">"
	value := int64(100)

	filter := DNSFilterQueryTime{
		Key:      key,
		Operator: operator,
		Value:    value,
	}

	if filter.Key != key {
		t.Errorf("Expected Key to be %s, got %s", key, filter.Key)
	}
	if filter.Operator != operator {
		t.Errorf("Expected Operator to be %s, got %s", operator, filter.Operator)
	}
	if filter.Value != value {
		t.Errorf("Expected Value to be %d, got %d", value, filter.Value)
	}
}

// Made with Bob
