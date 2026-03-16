package user
// UsersResourcePath path to User resource of Instana RESTful API
const UsersResourcePath = "/api/settings" + "/users"


// ResourcePath is the path to the Users resource in the Instana API
const ResourcePath = "/api/settings/users"

// User data structure for the Instana API model for users
type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	FullName     string `json:"fullName"`
	LastLoggedIn *int64 `json:"lastLoggedIn,omitempty"`
	GroupCount   *int   `json:"groupCount,omitempty"`
	TfaEnabled   *bool  `json:"tfaEnabled,omitempty"`
}

// GetIDForResourcePath implementation of the interface InstanaDataObject
func (u *User) GetIDForResourcePath() string {
	return u.ID
}
