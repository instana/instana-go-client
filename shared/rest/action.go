package rest

import "fmt"

// ActionMethod is the HTTP method used to invoke a resource action.
type ActionMethod string

const (
	// ActionMethodPut invokes the action using HTTP PUT
	ActionMethodPut = ActionMethod("PUT")
	// ActionMethodPost invokes the action using HTTP POST
	ActionMethodPost = ActionMethod("POST")
	// ActionMethodPatch invokes the action using HTTP PATCH
	ActionMethodPatch = ActionMethod("PATCH")
)

const (
	// ActionEnable is the action sub-path used to enable a resource ({basePath}/{id}/enable)
	ActionEnable = "enable"
	// ActionDisable is the action sub-path used to disable a resource ({basePath}/{id}/disable)
	ActionDisable = "disable"
)

// ActionRestResource invokes sub-endpoints of the form {basePath}/{id}/{action}
// (e.g. /api/events/settings/infra-alert-configs/{id}/enable) that are not part of the
// regular create/update payload.
type ActionRestResource interface {
	// Execute calls {basePath}/{id}/{action} using the given HTTP method without a payload.
	Execute(method ActionMethod, id string, action string) error
}

// NewActionRestResource creates an ActionRestResource for the given base resource path.
func NewActionRestResource(client RestClient, basePath string) ActionRestResource {
	return &actionRestResource{basePath: basePath, client: client}
}

type actionRestResource struct {
	basePath string
	client   RestClient
}

// Execute calls {basePath}/{id}/{action} using the given HTTP method without a payload.
func (r *actionRestResource) Execute(method ActionMethod, id string, action string) error {
	obj := &actionDataObject{idPath: id + "/" + action}
	var err error
	switch method {
	case ActionMethodPut:
		_, err = r.client.Put(obj, r.basePath)
	case ActionMethodPost:
		_, err = r.client.PostWithID(obj, r.basePath)
	case ActionMethodPatch:
		_, err = r.client.Patch(obj, r.basePath)
	default:
		return fmt.Errorf("unsupported action method %q", method)
	}
	return err
}

// actionDataObject is a thin adapter that satisfies InstanaDataObject so that the shared
// RestClient methods can be used to call {basePath}/{id}/{action}.
type actionDataObject struct {
	idPath string
}

// GetIDForResourcePath returns "{id}/{action}" which the RestClient appends to the base path.
func (o *actionDataObject) GetIDForResourcePath() string {
	return o.idPath
}

// MarshalJSON sends a JSON null body, which is valid for the nullable request bodies of action endpoints.
func (o *actionDataObject) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}
