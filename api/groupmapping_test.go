package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
)

func TestGroupMappingResourcePath(t *testing.T) {
	expected := "/api/settings/rbac/mappings"
	if GroupMappingResourcePath != expected {
		t.Errorf("Expected GroupMappingResourcePath to be %s, got %s", expected, GroupMappingResourcePath)
	}
}

func TestGroupMappingGetIDForResourcePath(t *testing.T) {
	testID := "mappingId"
	mapping := &GroupMapping{
		ID:      testID,
		Key:     "roles",
		Value:   "analyst",
		GroupID: "-3",
	}

	result := mapping.GetIDForResourcePath()
	if result != testID {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", testID, result)
	}
}

func TestGroupMappingStructure(t *testing.T) {
	teamID := "team-1"
	mapping := GroupMapping{
		ID:      "mappingId",
		Key:     "roles",
		Value:   "analyst",
		GroupID: "-3",
		TeamID:  &teamID,
	}

	if mapping.ID != "mappingId" {
		t.Errorf("Expected ID 'mappingId', got %s", mapping.ID)
	}
	if mapping.Key != "roles" {
		t.Errorf("Expected Key 'roles', got %s", mapping.Key)
	}
	if mapping.Value != "analyst" {
		t.Errorf("Expected Value 'analyst', got %s", mapping.Value)
	}
	if mapping.GroupID != "-3" {
		t.Errorf("Expected GroupID '-3', got %s", mapping.GroupID)
	}
	if mapping.TeamID == nil || *mapping.TeamID != "team-1" {
		t.Errorf("Expected TeamID 'team-1', got %v", mapping.TeamID)
	}
}

func TestGroupMappingTeamIDOptional(t *testing.T) {
	// TeamID is optional; omitting it should leave the field nil.
	mapping := GroupMapping{
		ID:      "mappingId",
		Key:     "department",
		Value:   "engineering",
		GroupID: "group-42",
	}

	if mapping.TeamID != nil {
		t.Errorf("Expected TeamID to be nil when not set, got %v", mapping.TeamID)
	}
}

func TestGroupMappingEmptyID(t *testing.T) {
	// The API ignores the id on create; the provider should still handle an empty id.
	mapping := &GroupMapping{
		Key:     "department",
		Value:   "engineering",
		GroupID: "group-42",
	}

	if mapping.GetIDForResourcePath() != "" {
		t.Errorf("Expected empty string for GetIDForResourcePath when ID is not set, got %s", mapping.GetIDForResourcePath())
	}
}
