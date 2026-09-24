package api

import (
	"encoding/json"
	"fmt"

	"github.com/instana/instana-go-client/shared/rest"
)

// IPFilteringResourcePath is the path to the IP filtering resource of the Instana RESTful API.
const IPFilteringResourcePath = "/api/settings/ip-filtering"

// IPFilteringVerifyResourcePath is the path to verify and promote IP filtering to permanently active.
const IPFilteringVerifyResourcePath = "/api/settings/ip-filtering/verify"

// IPFilteringRule represents a single traffic rule enforcing allow/block based on incoming IP address.
type IPFilteringRule struct {
	// Target of this rule, can be either an IP address or an IP address range.
	Target string `json:"target"`
	// Block indicates if the traffic should be blocked (true) or allowed (false).
	Block bool `json:"block"`
}

// IPFiltering is the representation of the tenant unit IP filtering settings in Instana.
// The API does not assign an ID to this resource — it is a singleton per tenant unit.
type IPFiltering struct {
	// Active indicates if the configuration is currently enforced (permanent or temporary).
	Active bool `json:"active,omitempty"`
	// DenyAll is the fallback configuration if unmatched requests should be denied.
	DenyAll bool `json:"denyAll"`
	// Enabled indicates if the configuration should be enforced.
	Enabled bool `json:"enabled"`
	// LastChangedAt is the timestamp when the configuration was last changed.
	LastChangedAt *int64 `json:"lastChangedAt,omitempty"`
	// LastVerifiedAt is the timestamp when the configuration was last verified.
	LastVerifiedAt *int64 `json:"lastVerifiedAt,omitempty"`
	// Rules contains the actual traffic rules evaluated in the given order.
	Rules []IPFilteringRule `json:"rules"`
	// SupportAccessEnabled is an optional flag for support access functionality.
	SupportAccessEnabled bool `json:"supportAccessEnabled"`
}

// IPFilteringRestResource is the singleton REST resource client interface for IP filtering,
// extending SingletonRestResource with the Verify operation.
type IPFilteringRestResource interface {
	rest.SingletonRestResource[*IPFiltering]
	// Verify promotes the IP filtering configuration from temporary to permanently active.
	Verify() (*IPFiltering, error)
}

// NewIPFilteringRestResource creates a new singleton REST resource client for IP filtering.
func NewIPFilteringRestResource(client rest.RestClient) IPFilteringRestResource {
	return &ipFilteringRestResource{
		resourcePath: IPFilteringResourcePath,
		verifyPath:   IPFilteringVerifyResourcePath,
		client:       client,
	}
}

type ipFilteringRestResource struct {
	resourcePath string
	verifyPath   string
	client       rest.RestClient
}

// Get retrieves the current IP filtering configuration via HTTP GET.
func (r *ipFilteringRestResource) Get() (*IPFiltering, error) {
	data, err := r.client.Get(r.resourcePath)
	if err != nil {
		return nil, err
	}
	return r.unmarshal(data)
}

// Upsert creates or updates the IP filtering configuration via HTTP PUT.
// The API always replaces the entire resource; partial updates are not supported.
func (r *ipFilteringRestResource) Upsert(filtering *IPFiltering) (*IPFiltering, error) {
	body, err := json.Marshal(filtering)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal IP filtering: %s", err)
	}
	data, err := r.client.Put(&ipFilteringDataObject{payload: body, path: r.resourcePath}, r.resourcePath)
	if err != nil {
		return nil, err
	}
	return r.unmarshal(data)
}

// Verify promotes the IP filtering configuration from temporary to permanently active via HTTP PATCH.
func (r *ipFilteringRestResource) Verify() (*IPFiltering, error) {
	verifyPayload, err := json.Marshal(map[string]bool{"verify": true})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal IP filtering verify payload: %s", err)
	}
	data, err := r.client.Patch(&ipFilteringDataObject{payload: verifyPayload, path: r.verifyPath}, r.verifyPath)
	if err != nil {
		return nil, err
	}
	return r.unmarshal(data)
}

// Delete removes the IP filtering configuration.
func (r *ipFilteringRestResource) Delete() error {
	return r.client.Delete("", r.resourcePath)
}

func (r *ipFilteringRestResource) unmarshal(data []byte) (*IPFiltering, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("empty response received for IP filtering")
	}
	var f IPFiltering
	if err := json.Unmarshal(data, &f); err != nil {
		return nil, fmt.Errorf("failed to parse IP filtering response: %s", err)
	}
	return &f, nil
}

// ipFilteringDataObject is a thin adapter that satisfies rest.InstanaDataObject so that
// the shared RestClient.Put method can be used. The ID is intentionally empty because the
// IP filtering endpoint has no per-resource ID in its path.
type ipFilteringDataObject struct {
	payload []byte
	path    string
}

// GetIDForResourcePath returns an empty string because IP filtering lives at the fixed path
// /api/settings/ip-filtering with no trailing ID segment.
func (o *ipFilteringDataObject) GetIDForResourcePath() string {
	return ""
}

// MarshalJSON delegates serialisation to the pre-marshalled payload.
func (o *ipFilteringDataObject) MarshalJSON() ([]byte, error) {
	return o.payload, nil
}
