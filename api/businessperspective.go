package api

import (
	"github.com/instana/instana-go-client/shared/tagfilter"
)

// BusinessPerspectiveResourcePath path to business perspective resource of Instana RESTful API
const BusinessPerspectiveResourcePath = "/api/business-monitoring/business-perspectives"

// BusinessPerspectiveConfig represents the REST resource of business perspective configuration at Instana
type BusinessPerspectiveConfig struct {
	ID                  string               `json:"id,omitempty"`
	Name                string               `json:"name"`
	Description         string               `json:"description,omitempty"`
	TagFilterExpression *tagfilter.TagFilter `json:"tagFilterExpression,omitempty"`
	RbacTags            []RbacTag            `json:"rbacTags,omitempty"`
}

// GetIDForResourcePath implementation of the interface InstanaDataObject
func (b *BusinessPerspectiveConfig) GetIDForResourcePath() string {
	return b.ID
}
