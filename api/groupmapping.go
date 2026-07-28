package api

// GroupMappingResourcePath path to Group Mapping resource of Instana RESTful API
const GroupMappingResourcePath = "/api/settings/rbac/mappings"

// GroupMapping data structure for the Instana API model for group mappings.
// A group mapping maps an IdP (LDAP, OIDC, SAML) attribute key/value pair to an
// Instana group so that users whose IdP token contains that pair are automatically
// assigned to the corresponding group on login.
// TeamID is optional: when provided, the mapping is also scoped to that team.
type GroupMapping struct {
	ID      string  `json:"id"`
	Key     string  `json:"key"`
	Value   string  `json:"value"`
	GroupID string  `json:"groupId"`
	TeamID  *string `json:"teamId,omitempty"`
}

// GetIDForResourcePath implementation of the InstanaDataObject interface.
func (m *GroupMapping) GetIDForResourcePath() string {
	return m.ID
}
