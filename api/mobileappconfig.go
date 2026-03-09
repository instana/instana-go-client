package api

import "github.com/instana/instana-go-client/shared/rest"

// MobileAppConfigResourcePath path to mobile app monitoring config resource of Instana RESTful API
const MobileAppConfigResourcePath = "/api/mobile-app-monitoring/config"

// MobileAppConfig data structure of a Mobile App Configuration of the Instana API
type MobileAppConfig struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetIDForResourcePath implementation of the interface InstanaDataObject
func (r *MobileAppConfig) GetIDForResourcePath() string {
	return r.ID
}

// NewMobileAppConfigRestResource creates a new REST resource for the mobile app config
func NewMobileAppConfigRestResource(unmarshaller rest.JSONUnmarshaller[*MobileAppConfig], client rest.RestClient) rest.RestResource[*MobileAppConfig] {
	return &mobileAppConfigRestResource{
		resourcePath: MobileAppConfigResourcePath,
		unmarshaller: unmarshaller,
		client:       client,
	}
}

type mobileAppConfigRestResource struct {
	resourcePath string
	unmarshaller rest.JSONUnmarshaller[*MobileAppConfig]
	client       rest.RestClient
}

func (r *mobileAppConfigRestResource) GetAll() (*[]*MobileAppConfig, error) {
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

func (r *mobileAppConfigRestResource) GetOne(id string) (*MobileAppConfig, error) {
	data, err := r.client.GetOne(id, r.resourcePath)
	if err != nil {
		return nil, err
	}
	return r.validateResponseAndConvertToStruct(data)
}

func (r *mobileAppConfigRestResource) Create(data *MobileAppConfig) (*MobileAppConfig, error) {

	response, err := r.client.PostByQuery(r.resourcePath, map[string]string{"name": data.Name})
	if err != nil {
		return data, err
	}
	return r.validateResponseAndConvertToStruct(response)
}

func (r *mobileAppConfigRestResource) Update(data *MobileAppConfig) (*MobileAppConfig, error) {

	response, err := r.client.PutByQuery(r.resourcePath, data.GetIDForResourcePath(), map[string]string{"name": data.Name})
	if err != nil {
		return data, err
	}
	return r.validateResponseAndConvertToStruct(response)
}

func (r *mobileAppConfigRestResource) validateResponseAndConvertToStruct(data []byte) (*MobileAppConfig, error) {
	dataObject, err := r.unmarshaller.Unmarshal(data)
	if err != nil {
		return nil, err
	}
	return dataObject, nil
}

func (r *mobileAppConfigRestResource) Delete(data *MobileAppConfig) error {
	return r.DeleteByID(data.GetIDForResourcePath())
}

func (r *mobileAppConfigRestResource) DeleteByID(id string) error {
	return r.client.Delete(id, r.resourcePath)
}
