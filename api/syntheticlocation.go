package api

// ResourcePath is the path to the Synthetic Locations resource in the Instana API
const SyntheticLocationResourcePath = "/api/synthetics/settings/locations"

type SyntheticLocation struct {
	ID           string `json:"id"`
	Label        string `json:"label"`
	Description  string `json:"description"`
	LocationType string `json:"locationType"`
}

// GetIDForResourcePath implementation of the interface InstanaDataObject for SyntheticLocation
func (s *SyntheticLocation) GetIDForResourcePath() string {
	return s.ID
}
