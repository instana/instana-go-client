package api

import (
	"encoding/json"

	"github.com/instana/instana-go-client/shared/types"
)

// CustomDashboardsResourcePath is the path to the Custom Dashboards resource in the Instana API
const CustomDashboardsResourcePath = "/api/custom-dashboard"

type CustomDashboard struct {
	ID          string             `json:"id"`
	Title       string             `json:"title"`
	AccessRules []types.AccessRule `json:"accessRules"`
	Widgets     json.RawMessage    `json:"widgets"`
}

// GetIDForResourcePath implementation of the interface InstanaDataObject for CustomDashboard
func (a *CustomDashboard) GetIDForResourcePath() string {
	return a.ID
}
