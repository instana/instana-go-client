package rest

import (
	"encoding/json"
	"fmt"
)

// GenericUnmarshaller is a generic JSON unmarshaller that works for any type.
// It eliminates the need for type-specific unmarshaller implementations.
type GenericUnmarshaller[T any] struct{}

// NewGenericUnmarshaller creates a new generic unmarshaller for the specified type.
// This is a convenience function that returns a pointer to GenericUnmarshaller.
//
// Example:
//
//	unmarshaller := rest.NewGenericUnmarshaller[*apitoken.APIToken]()
//	token, err := unmarshaller.Unmarshal(jsonData)
func NewGenericUnmarshaller[T any]() *GenericUnmarshaller[T] {
	return &GenericUnmarshaller[T]{}
}

// Unmarshal converts JSON bytes into the target type T.
// It returns an error if the JSON is invalid or cannot be unmarshalled into type T.
//
// Example:
//
//	var jsonData = []byte(`{"id": "123", "name": "test"}`)
//	token, err := unmarshaller.Unmarshal(jsonData)
func (u *GenericUnmarshaller[T]) Unmarshal(data []byte) (T, error) {
	var target T
	if err := json.Unmarshal(data, &target); err != nil {
		var zero T
		return zero, fmt.Errorf("failed to parse json; %s", err)
	}
	return target, nil
}

// UnmarshalArray converts JSON bytes into a slice of the target type T.
// It returns an error if the JSON is invalid or cannot be unmarshalled into []T.
//
// Example:
//
//	var jsonData = []byte(`[{"id": "123"}, {"id": "456"}]`)
//	tokens, err := unmarshaller.UnmarshalArray(jsonData)
func (u *GenericUnmarshaller[T]) UnmarshalArray(data []byte) (*[]T, error) {
	var target []T
	if err := json.Unmarshal(data, &target); err != nil {
		return &target, fmt.Errorf("failed to parse json; %s", err)
	}
	return &target, nil
}

// Made with Bob
