package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
)

func TestRolesResourcePath(t *testing.T) {
	expected := "/api/settings/rbac/roles"
	if RolesResourcePath != expected {
		t.Errorf("Expected RolesResourcePath to be %s, got %s", expected, RolesResourcePath)
	}
}

func TestRoleGetIDForResourcePath(t *testing.T) {
	testID := "role-id-123"
	role := &Role{
		ID:   testID,
		Name: "Test Role",
	}

	result := role.GetIDForResourcePath()
	if result != testID {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", testID, result)
	}
}

func TestRoleStructure(t *testing.T) {
	email1 := "user1@example.com"
	name1 := "User One"
	email2 := "user2@example.com"
	name2 := "User Two"

	role := Role{
		ID:   "role-123",
		Name: "Administrator",
		Members: []APIMember{
			{
				UserID: "user-1",
				Email:  &email1,
				Name:   &name1,
			},
			{
				UserID: "user-2",
				Email:  &email2,
				Name:   &name2,
			},
		},
		Permissions: []string{
			"CAN_CONFIGURE_APPLICATIONS",
			"CAN_CONFIGURE_USERS",
			"CAN_VIEW_LOGS",
		},
	}

	if role.ID != "role-123" {
		t.Errorf("Expected ID 'role-123', got %s", role.ID)
	}
	if role.Name != "Administrator" {
		t.Errorf("Expected Name 'Administrator', got %s", role.Name)
	}
	if len(role.Members) != 2 {
		t.Errorf("Expected 2 members, got %d", len(role.Members))
	}
	if len(role.Permissions) != 3 {
		t.Errorf("Expected 3 permissions, got %d", len(role.Permissions))
	}
}

func TestAPIMemberStructure(t *testing.T) {
	email := "test@example.com"
	name := "Test User"

	member := APIMember{
		UserID: "user-123",
		Email:  &email,
		Name:   &name,
	}

	if member.UserID != "user-123" {
		t.Errorf("Expected UserID 'user-123', got %s", member.UserID)
	}
	if member.Email == nil || *member.Email != email {
		t.Error("Email not set correctly")
	}
	if member.Name == nil || *member.Name != name {
		t.Error("Name not set correctly")
	}
}

func TestRoleWithEmptyCollections(t *testing.T) {
	role := Role{
		ID:          "role-456",
		Name:        "Empty Role",
		Members:     []APIMember{},
		Permissions: []string{},
	}

	if len(role.Members) != 0 {
		t.Errorf("Expected 0 members, got %d", len(role.Members))
	}
	if len(role.Permissions) != 0 {
		t.Errorf("Expected 0 permissions, got %d", len(role.Permissions))
	}
}
