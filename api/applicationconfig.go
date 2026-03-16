package api

import (
	"github.com/instana/instana-go-client/shared/tagfilter"
	"github.com/instana/instana-go-client/shared/types"
)

const (
	//ApplicationConfigsResourcePath path to application config resource of Instana RESTful API
	ApplicationConfigsResourcePath = "/api/application-monitoring/settings/application"
)

// ApplicationConfigResource represents the REST resource of application perspective configuration at Instana
type ApplicationConfigResource interface {
	GetOne(id string) (ApplicationConfig, error)
	Upsert(rule ApplicationConfig) (ApplicationConfig, error)
	Delete(rule ApplicationConfig) error
	DeleteByID(applicationID string) error
}

// ApplicationConfig is the representation of a application perspective configuration in Instana
type ApplicationConfig struct {
	ID                  string                 `json:"id"`
	Label               string                 `json:"label"`
	TagFilterExpression *tagfilter.TagFilter   `json:"tagFilterExpression"`
	Scope               ApplicationConfigScope `json:"scope"`
	BoundaryScope       types.BoundaryScope    `json:"boundaryScope"`
	AccessRules         []types.AccessRule     `json:"accessRules"`
}

// GetIDForResourcePath implementation of the interface InstanaDataObject
func (a *ApplicationConfig) GetIDForResourcePath() string {
	return a.ID
}

// ApplicationConfigScope type definition of the application config scope of the Instana Web REST API
type ApplicationConfigScope string

// ApplicationConfigScopes type definition of slice of ApplicationConfigScope
type ApplicationConfigScopes []ApplicationConfigScope

// ToStringSlice returns a slice containing the string representations of the given scopes
func (scopes ApplicationConfigScopes) ToStringSlice() []string {
	result := make([]string, len(scopes))
	for i, s := range scopes {
		result[i] = string(s)
	}
	return result
}

const (
	//ApplicationConfigScopeIncludeNoDownstream constant for the scope INCLUDE_NO_DOWNSTREAM
	ApplicationConfigScopeIncludeNoDownstream = ApplicationConfigScope("INCLUDE_NO_DOWNSTREAM")
	//ApplicationConfigScopeIncludeImmediateDownstreamDatabaseAndMessaging constant for the scope INCLUDE_IMMEDIATE_DOWNSTREAM_DATABASE_AND_MESSAGING
	ApplicationConfigScopeIncludeImmediateDownstreamDatabaseAndMessaging = ApplicationConfigScope("INCLUDE_IMMEDIATE_DOWNSTREAM_DATABASE_AND_MESSAGING")
	//ApplicationConfigScopeIncludeAllDownstream constant for the scope INCLUDE_ALL_DOWNSTREAM
	ApplicationConfigScopeIncludeAllDownstream = ApplicationConfigScope("INCLUDE_ALL_DOWNSTREAM")
)

// SupportedApplicationConfigScopes supported ApplicationConfigScopes of the Instana Web REST API
var SupportedApplicationConfigScopes = ApplicationConfigScopes{
	ApplicationConfigScopeIncludeNoDownstream,
	ApplicationConfigScopeIncludeImmediateDownstreamDatabaseAndMessaging,
	ApplicationConfigScopeIncludeAllDownstream,
}

// IncludedEndpoint custom type to include of a specific endpoint in an alert config
type IncludedEndpoint struct {
	EndpointID string `json:"endpointId"`
	Inclusive  bool   `json:"inclusive"`
}

// IncludedService custom type to include of a specific service in an alert config
type IncludedService struct {
	ServiceID string `json:"serviceId"`
	Inclusive bool   `json:"inclusive"`

	Endpoints map[string]IncludedEndpoint `json:"endpoints"`
}

// IncludedApplication custom type to include specific applications in an alert config
type IncludedApplication struct {
	ApplicationID string `json:"applicationId"`
	Inclusive     bool   `json:"inclusive"`

	Services map[string]IncludedService `json:"services"`
}
