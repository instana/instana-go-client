package api

import "github.com/instana/instana-go-client/shared/rest"

const (
	// SyntheticCredentialResourcePath is the base path for create/update/delete operations
	SyntheticCredentialResourcePath = "/api/synthetics/settings/credentials"

	// SyntheticCredentialAssociationsResourcePath is the path used for read operations;
	// it returns the full credential object including scope associations.
	SyntheticCredentialAssociationsResourcePath = "/api/synthetics/settings/credentials/associations"
)

// SyntheticCredential represents the REST resource for a Synthetic Credential at Instana
type SyntheticCredential struct {
	// CredentialName is the unique name of the credential (used as the resource identifier)
	CredentialName string `json:"credentialName"`
	// CredentialValue is the secret value of the credential (write-only; not returned by the read endpoints)
	CredentialValue string `json:"credentialValue,omitempty"`
	// Applications is the list of application IDs that the credential is scoped to
	Applications []string `json:"applications,omitempty"`
	// MobileApps is the list of mobile app IDs that the credential is scoped to
	MobileApps []string `json:"mobileApps,omitempty"`
	// Websites is the list of website IDs that the credential is scoped to
	Websites []string `json:"websites,omitempty"`
	// RbacTags controls the RBAC visibility of the credential
	RbacTags []RbacTag `json:"rbacTags,omitempty"`
}

// GetIDForResourcePath returns the credential name as the unique resource path identifier
func (s *SyntheticCredential) GetIDForResourcePath() string {
	return s.CredentialName
}

// NewSyntheticCredentialRestResource creates a custom REST resource for synthetic credentials.
// The create and update endpoints return an empty body, so after each mutation the implementation
// fetches the full object via the associations endpoint.
func NewSyntheticCredentialRestResource(unmarshaller rest.JSONUnmarshaller[*SyntheticCredential], client rest.RestClient) rest.RestResource[*SyntheticCredential] {
	return &syntheticCredentialRestResource{
		resourcePath:             SyntheticCredentialResourcePath,
		associationsResourcePath: SyntheticCredentialAssociationsResourcePath,
		unmarshaller:             unmarshaller,
		client:                   client,
	}
}

type syntheticCredentialRestResource struct {
	resourcePath             string
	associationsResourcePath string
	unmarshaller             rest.JSONUnmarshaller[*SyntheticCredential]
	client                   rest.RestClient
}

// GetAll returns names only (the base GET endpoint); use GetOne for the full object.
func (r *syntheticCredentialRestResource) GetAll() (*[]*SyntheticCredential, error) {
	data, err := r.client.Get(r.resourcePath)
	if err != nil {
		return nil, err
	}
	return r.unmarshaller.UnmarshalArray(data)
}

// GetOne fetches the full credential (with associations) via the associations sub-path.
func (r *syntheticCredentialRestResource) GetOne(name string) (*SyntheticCredential, error) {
	data, err := r.client.GetOne(name, r.associationsResourcePath)
	if err != nil {
		return nil, err
	}
	return r.unmarshal(data)
}

// Create posts the new credential and then reads it back via GetOne.
// The POST endpoint returns an empty body so we cannot parse the response directly.
func (r *syntheticCredentialRestResource) Create(data *SyntheticCredential) (*SyntheticCredential, error) {
	_, err := r.client.Post(data, r.resourcePath)
	if err != nil {
		return nil, err
	}
	return r.GetOne(data.CredentialName)
}

// Update sends a PUT for the credential and then reads it back via GetOne.
// The PUT endpoint returns an empty body so we cannot parse the response directly.
func (r *syntheticCredentialRestResource) Update(data *SyntheticCredential) (*SyntheticCredential, error) {
	_, err := r.client.Put(data, r.resourcePath)
	if err != nil {
		return nil, err
	}
	return r.GetOne(data.CredentialName)
}

func (r *syntheticCredentialRestResource) Delete(data *SyntheticCredential) error {
	return r.DeleteByID(data.CredentialName)
}

func (r *syntheticCredentialRestResource) DeleteByID(id string) error {
	return r.client.Delete(id, r.resourcePath)
}

func (r *syntheticCredentialRestResource) unmarshal(data []byte) (*SyntheticCredential, error) {
	return r.unmarshaller.Unmarshal(data)
}
