package api

import (
	"encoding/json"
	"fmt"

	"github.com/instana/instana-go-client/shared/rest"
)

// SessionSettingsResourcePath is the path to the session settings resource of the Instana RESTful API.
const SessionSettingsResourcePath = "/api/settings/session"

// defaultTokenLifeTimeInMillis is the server default for token lifetime (7 days).
const defaultTokenLifeTimeInMillis int64 = 604800000

// defaultIdleTimeInMillis is the server default for idle timeout (8 hours).
const defaultIdleTimeInMillis int64 = 28800000

// SessionSettings is the representation of the tenant unit session settings in Instana.
// The API does not assign an ID to this resource — it is a singleton per tenant unit.
//
// Constraints (enforced by the server):
//   - TokenLifeTimeInMillis: 600_000 ms (10 min) … 604_800_000 ms (7 days)
//   - IdleTimeInMillis:      600_000 ms (10 min) … 28_800_000 ms (8 hours)
type SessionSettings struct {
	// TokenLifeTimeInMillis is the maximum lifetime of an authentication token in milliseconds.
	// Valid range: 600000 (10 min) to 604800000 (7 days).
	TokenLifeTimeInMillis int64 `json:"tokenLifeTimeInMillis"`
	// IdleTimeInMillis is the idle timeout before a session expires in milliseconds.
	// Valid range: 600000 (10 min) to 28800000 (8 hours).
	IdleTimeInMillis int64 `json:"idleTimeInMillis"`
}

// NewSessionSettingsRestResource creates a new singleton REST resource client for session settings.
func NewSessionSettingsRestResource(client rest.RestClient) rest.SingletonRestResource[*SessionSettings] {
	return &sessionSettingsRestResource{
		resourcePath: SessionSettingsResourcePath,
		client:       client,
	}
}

type sessionSettingsRestResource struct {
	resourcePath string
	client       rest.RestClient
}

// Get retrieves the current session settings via HTTP GET.
func (r *sessionSettingsRestResource) Get() (*SessionSettings, error) {
	data, err := r.client.Get(r.resourcePath)
	if err != nil {
		return nil, err
	}
	return r.unmarshal(data)
}

// Upsert creates or updates the session settings via HTTP PUT.
// The API always replaces the entire resource; partial updates are not supported.
func (r *sessionSettingsRestResource) Upsert(settings *SessionSettings) (*SessionSettings, error) {
	body, err := json.Marshal(settings)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal session settings: %s", err)
	}
	data, err := r.client.Put(&sessionSettingsDataObject{payload: body, path: r.resourcePath}, r.resourcePath)
	if err != nil {
		return settings, err
	}
	return r.unmarshal(data)
}

// Delete removes the session settings, reverting the tenant unit to default values.
func (r *sessionSettingsRestResource) Delete() error {
	return r.client.Delete("", r.resourcePath)
}

func (r *sessionSettingsRestResource) unmarshal(data []byte) (*SessionSettings, error) {
	// The API returns 200 with an empty body when the tenant has not configured
	// custom session settings yet. Return the server defaults in that case.
	if len(data) == 0 {
		return &SessionSettings{
			TokenLifeTimeInMillis: defaultTokenLifeTimeInMillis,
			IdleTimeInMillis:      defaultIdleTimeInMillis,
		}, nil
	}
	var s SessionSettings
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, fmt.Errorf("failed to parse session settings response: %s", err)
	}
	return &s, nil
}

// sessionSettingsDataObject is a thin adapter that satisfies rest.InstanaDataObject so that
// the shared RestClient.Put method can be used. The ID is intentionally empty because the
// session settings endpoint has no per-resource ID in its path.
type sessionSettingsDataObject struct {
	payload []byte
	path    string
}

// GetIDForResourcePath returns an empty string because session settings live at the fixed path
// /api/settings/session with no trailing ID segment.
func (o *sessionSettingsDataObject) GetIDForResourcePath() string {
	return ""
}

// MarshalJSON delegates serialisation to the pre-marshalled payload.
func (o *sessionSettingsDataObject) MarshalJSON() ([]byte, error) {
	return o.payload, nil
}
