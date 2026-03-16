package models

// AccessType custom type for access type
type AccessType string

// AccessTypes custom type for a slice of AccessType
type AccessTypes []AccessType

// ToStringSlice Returns the corresponding string representations
func (types AccessTypes) ToStringSlice() []string {
	result := make([]string, len(types))
	for i, v := range types {
		result[i] = string(v)
	}
	return result
}

const (
	//AccessTypeRead constant value for the READ AccessType
	AccessTypeRead = AccessType("READ")
	//AccessTypeReadWrite constant value for the READ_WRITE AccessType
	AccessTypeReadWrite = AccessType("READ_WRITE")
)

// SupportedAccessTypes list of all supported AccessType
var SupportedAccessTypes = AccessTypes{AccessTypeRead, AccessTypeReadWrite}

// RelationType custom type for relation type
type RelationType string

// RelationTypes custom type for a slice of RelationType
type RelationTypes []RelationType

// ToStringSlice Returns the corresponding string representations
func (types RelationTypes) ToStringSlice() []string {
	result := make([]string, len(types))
	for i, v := range types {
		result[i] = string(v)
	}
	return result
}

const (
	//RelationTypeUser constant value for the USER RelationType
	RelationTypeUser = RelationType("USER")
	//RelationTypeApiToken constant value for the API_TOKEN RelationType
	RelationTypeApiToken = RelationType("API_TOKEN")
	//RelationTypeRole constant value for the ROLE RelationType
	RelationTypeRole = RelationType("ROLE")
	//RelationTypeTeam constant value for the TEAM RelationType
	RelationTypeTeam = RelationType("TEAM")
	//RelationTypeGlobal constant value for the GLOBAL RelationType
	RelationTypeGlobal = RelationType("GLOBAL")
)

// SupportedRelationTypes list of all supported RelationType
var SupportedRelationTypes = RelationTypes{RelationTypeUser, RelationTypeApiToken, RelationTypeRole, RelationTypeTeam, RelationTypeGlobal}

// AccessRule represents an access control rule
type AccessRule struct {
	AccessType   AccessType   `json:"accessType"`
	RelatedID    *string      `json:"relatedId"`
	RelationType RelationType `json:"relationType"`
}

// Made with Bob
