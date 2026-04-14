package api

import (
	"encoding/json"
	"fmt"

	"github.com/instana/instana-go-client/shared/rest"
)

// ResourcePath is the path to the Host Agents resource in the Instana API
const HostAgentResourcePath = "/api/host-agent"

type HostAgent struct {
	SnapshotID string   `json:"snapshotId"`
	Label      string   `json:"label"`
	Host       string   `json:"host"`
	Plugin     string   `json:"plugin"`
	Tags       []string `json:"tags"`
}

// GetIDForResourcePath implemention of the interface InstanaDataObject
func (spec *HostAgent) GetIDForResourcePath() string {
	return spec.SnapshotID
}

// NewHostAgentJSONUnmarshaller creates a new instance of a generic JSONUnmarshaller.
func NewHostAgentJSONUnmarshaller[T rest.InstanaDataObject](objectType T) rest.JSONUnmarshaller[T] {
	arrayType := make(map[string][]T)
	arrayType["items"] = []T{}

	return &hostAgentJSONUnmarshaller[T]{
		objectType: objectType,
		arrayType:  &arrayType,
	}
}

type hostAgentJSONUnmarshaller[T any] struct {
	objectType T
	arrayType  *map[string][]T
}

// UnmarshalJSON unmarshals JSON data into the target object.
func (u *hostAgentJSONUnmarshaller[T]) Unmarshal(data []byte) (T, error) {
	// Create a new instance to avoid shared state issues
	var target T
	if err := json.Unmarshal(data, &target); err != nil {
		var zero T
		return zero, fmt.Errorf("failed to parse json: %w", err)
	}
	return target, nil
}

// UnmarshalJSONArray unmarshals JSON array data into a slice of target objects.
func (u *hostAgentJSONUnmarshaller[T]) UnmarshalArray(data []byte) (*[]T, error) {
	// Create a new map instance to avoid shared state issues
	target := make(map[string][]T)
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, fmt.Errorf("failed to parse json: %w", err)
	}
	hostAgents := target["items"]
	return &hostAgents, nil
}
