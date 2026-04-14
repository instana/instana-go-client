package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
)

func TestGroupResourcePath(t *testing.T) {
	expected := "/api/settings/rbac/groups"
	if GroupsResourcePath != expected {
		t.Errorf("Expected GroupsResourcePath to be %s, got %s", expected, GroupsResourcePath)
	}
}

func TestGroupGetIDForResourcePath(t *testing.T) {
	testID := "test-group-123"
	group := &Group{
		ID:   testID,
		Name: "Test Group",
	}

	result := group.GetIDForResourcePath()
	if result != testID {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", testID, result)
	}
}

func TestGroupStructure(t *testing.T) {
	email := "user@example.com"
	name := "John Doe"

	group := Group{
		ID:   "group-456",
		Name: "Developers",
		Members: []APIMember{
			{UserID: "user-1", Email: &email, Name: &name},
		},
		PermissionSet: APIPermissionSetWithRoles{
			Permissions: []InstanaPermission{PermissionCanConfigureApplications},
		},
	}

	if group.ID != "group-456" {
		t.Errorf("Expected ID 'group-456', got %s", group.ID)
	}
	if group.Name != "Developers" {
		t.Errorf("Expected Name 'Developers', got %s", group.Name)
	}
	if len(group.Members) != 1 {
		t.Errorf("Expected 1 member, got %d", len(group.Members))
	}
	if len(group.PermissionSet.Permissions) != 1 {
		t.Errorf("Expected 1 permission, got %d", len(group.PermissionSet.Permissions))
	}
}

func TestGroupAPIMemberStructure(t *testing.T) {
	email := "test@example.com"
	name := "Test User"

	member := APIMember{
		UserID: "user-789",
		Email:  &email,
		Name:   &name,
	}

	if member.UserID != "user-789" {
		t.Errorf("Expected UserID 'user-789', got %s", member.UserID)
	}
	if member.Email == nil || *member.Email != email {
		t.Error("Email not set correctly")
	}
	if member.Name == nil || *member.Name != name {
		t.Error("Name not set correctly")
	}
}

func TestScopeBindingStructure(t *testing.T) {
	roleID := "role-123"

	binding := ScopeBinding{
		ScopeID:     "scope-456",
		ScopeRoleID: &roleID,
	}

	if binding.ScopeID != "scope-456" {
		t.Errorf("Expected ScopeID 'scope-456', got %s", binding.ScopeID)
	}
	if binding.ScopeRoleID == nil || *binding.ScopeRoleID != roleID {
		t.Error("ScopeRoleID not set correctly")
	}
}

func TestAPIPermissionSetWithRolesIsEmpty(t *testing.T) {
	// Test empty permission set
	emptySet := APIPermissionSetWithRoles{}
	if !emptySet.IsEmpty() {
		t.Error("Expected IsEmpty to return true for empty permission set")
	}

	// Test non-empty permission set
	nonEmptySet := APIPermissionSetWithRoles{
		Permissions: []InstanaPermission{PermissionCanConfigureApplications},
	}
	if nonEmptySet.IsEmpty() {
		t.Error("Expected IsEmpty to return false for non-empty permission set")
	}
}

func TestToStringSlice(t *testing.T) {
	SupportedInstanaPermissions := InstanaPermissions{
		PermissionCanConfigureApplications,
	}
	persmissionSet := SupportedInstanaPermissions.ToStringSlice()
	if persmissionSet[0] != "CAN_CONFIGURE_APPLICATIONS" {
		t.Error("ToStringSlice not working correctly")
	}
}
