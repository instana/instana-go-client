package api

// ReleaseResourcePath path to the releases resource of the Instana RESTful API
const ReleaseResourcePath = "/api/releases"

// Release represents the request/update body for a release in the Instana API
type Release struct {
	Name         string                     `json:"name"`
	Start        int64                      `json:"start"`
	Applications []*ReleaseApplicationScope `json:"applications,omitempty"`
	Services     []*ReleaseServiceScope     `json:"services,omitempty"`
}

// ReleaseWithMetadata represents a release as returned by the Instana API (includes server-set fields)
type ReleaseWithMetadata struct {
	ID           string                     `json:"id"`
	Name         string                     `json:"name"`
	Start        int64                      `json:"start"`
	LastUpdated  int64                      `json:"lastUpdated"`
	Applications []*ReleaseApplicationScope `json:"applications,omitempty"`
	Services     []*ReleaseServiceScope     `json:"services,omitempty"`
}

// ReleaseApplicationScope represents an application perspective where a release can be viewed
type ReleaseApplicationScope struct {
	Name string `json:"name"`
}

// ReleaseServiceScope represents a service where a release can be viewed
type ReleaseServiceScope struct {
	Name     string                  `json:"name"`
	ScopedTo *ReleaseServiceScopedTo `json:"scopedTo,omitempty"`
}

// ReleaseServiceScopedTo restricts a service scope to specific application perspectives (1–10 entries)
type ReleaseServiceScopedTo struct {
	Applications []*ReleaseApplicationScope `json:"applications"`
}

// GetIDForResourcePath implementation of the interface InstanaDataObject
func (r *ReleaseWithMetadata) GetIDForResourcePath() string {
	return r.ID
}
