package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
)

func TestUsersResourcePath(t *testing.T) {
	expected := "/api/settings/users"
	if UsersResourcePath != expected {
		t.Errorf("Expected UsersResourcePath to be %s, got %s", expected, UsersResourcePath)
	}
}

func TestUserGetIDForResourcePath(t *testing.T) {
	testID := "user-id-123"
	user := &User{
		ID:       testID,
		Email:    "test@example.com",
		FullName: "Test User",
	}

	result := user.GetIDForResourcePath()
	if result != testID {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", testID, result)
	}
}

func TestUserStructure(t *testing.T) {
	lastLoggedIn := int64(1234567890)
	groupCount := 3
	tfaEnabled := true

	user := User{
		ID:           "user-123",
		Email:        "john.doe@example.com",
		FullName:     "John Doe",
		LastLoggedIn: &lastLoggedIn,
		GroupCount:   &groupCount,
		TfaEnabled:   &tfaEnabled,
	}

	if user.ID != "user-123" {
		t.Errorf("Expected ID 'user-123', got %s", user.ID)
	}
	if user.Email != "john.doe@example.com" {
		t.Errorf("Expected Email 'john.doe@example.com', got %s", user.Email)
	}
	if user.FullName != "John Doe" {
		t.Errorf("Expected FullName 'John Doe', got %s", user.FullName)
	}
	if user.LastLoggedIn == nil || *user.LastLoggedIn != lastLoggedIn {
		t.Error("LastLoggedIn not set correctly")
	}
	if user.GroupCount == nil || *user.GroupCount != groupCount {
		t.Error("GroupCount not set correctly")
	}
	if user.TfaEnabled == nil || *user.TfaEnabled != tfaEnabled {
		t.Error("TfaEnabled not set correctly")
	}
}

func TestUserWithNilOptionalFields(t *testing.T) {
	user := User{
		ID:           "user-456",
		Email:        "jane.doe@example.com",
		FullName:     "Jane Doe",
		LastLoggedIn: nil,
		GroupCount:   nil,
		TfaEnabled:   nil,
	}

	if user.LastLoggedIn != nil {
		t.Error("Expected LastLoggedIn to be nil")
	}
	if user.GroupCount != nil {
		t.Error("Expected GroupCount to be nil")
	}
	if user.TfaEnabled != nil {
		t.Error("Expected TfaEnabled to be nil")
	}
}
