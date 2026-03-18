package api

import (
	"github.com/instana/instana-go-client/shared/rest"
	model "github.com/instana/instana-go-client/shared/types"
)

// WebsiteMonitoringConfigResourcePath path to website monitoring config resource of Instana RESTful API
const WebsiteMonitoringConfigResourcePath = "/api/website-monitoring/config"

// WebsiteMonitoringConfig data structure of a Website Monitoring Configuration of the Instana API
type WebsiteMonitoringConfig struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	AppName string `json:"appName"`
}

// GetIDForResourcePath implemention of the interface InstanaDataObject
func (r *WebsiteMonitoringConfig) GetIDForResourcePath() string {
	return r.ID
}

type websiteMonitoringConfigRestResource struct {
	resourcePath string
	unmarshaller rest.JSONUnmarshaller[*WebsiteMonitoringConfig]
	client       rest.RestClient
}

func (r *websiteMonitoringConfigRestResource) GetAll() (*[]*WebsiteMonitoringConfig, error) {
	data, err := r.client.Get(r.resourcePath)
	if err != nil {
		return nil, err
	}
	objects, err := r.unmarshaller.UnmarshalArray(data)
	if err != nil {
		return nil, err
	}
	return objects, nil
}

func (r *websiteMonitoringConfigRestResource) GetOne(id string) (*WebsiteMonitoringConfig, error) {
	data, err := r.client.GetOne(id, r.resourcePath)
	if err != nil {
		return nil, err
	}
	return r.validateResponseAndConvertToStruct(data)
}

func (r *websiteMonitoringConfigRestResource) Create(data *WebsiteMonitoringConfig) (*WebsiteMonitoringConfig, error) {
	response, err := r.client.PostByQuery(r.resourcePath, map[string]string{"name": data.Name})
	if err != nil {
		return data, err
	}
	return r.validateResponseAndConvertToStruct(response)
}

func (r *websiteMonitoringConfigRestResource) Update(data *WebsiteMonitoringConfig) (*WebsiteMonitoringConfig, error) {
	response, err := r.client.PutByQuery(r.resourcePath, data.GetIDForResourcePath(), map[string]string{"name": data.Name})
	if err != nil {
		return data, err
	}
	return r.validateResponseAndConvertToStruct(response)
}

func (r *websiteMonitoringConfigRestResource) validateResponseAndConvertToStruct(data []byte) (*WebsiteMonitoringConfig, error) {
	dataObject, err := r.unmarshaller.Unmarshal(data)
	if err != nil {
		return nil, err
	}
	return dataObject, nil
}

func (r *websiteMonitoringConfigRestResource) Delete(data *WebsiteMonitoringConfig) error {
	return r.DeleteByID(data.GetIDForResourcePath())
}

func (r *websiteMonitoringConfigRestResource) DeleteByID(id string) error {
	return r.client.Delete(id, r.resourcePath)
}

// WebsiteTimeThreshold struct representing the API model of a website time threshold
type WebsiteTimeThreshold struct {
	Type                    string                                `json:"type"`
	TimeWindow              *int64                                `json:"timeWindow"`
	Violations              *int32                                `json:"violations"`
	ImpactMeasurementMethod *model.WebsiteImpactMeasurementMethod `json:"impactMeasurementMethod"`
	UserPercentage          *float64                              `json:"userPercentage"`
	Users                   *int32                                `json:"users"`
}
