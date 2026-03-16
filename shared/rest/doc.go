// Package rest provides generic REST resource interfaces and implementations for the Instana API client.
//
// This package contains the core abstractions for interacting with REST resources:
//   - RestResource: Full CRUD operations interface
//   - ReadOnlyRestResource: Read-only operations interface
//   - RestClient: HTTP client interface for making API requests
//   - JSONUnmarshaller: Interface for JSON deserialization
//
// The package provides factory functions for creating REST resources with different
// HTTP method combinations (PUT/POST for create/update operations).
//
// Example usage:
//
//	// Create a REST resource that uses POST for create and PUT for update
//	resource := rest.NewCreatePOSTUpdatePUTRestResource(
//	    "/api/v1/resources",
//	    unmarshaller,
//	    client,
//	)
//
//	// Use the resource
//	item, err := resource.GetOne("resource-id")
//	if err != nil {
//	    // handle error
//	}
package rest

// Made with Bob
