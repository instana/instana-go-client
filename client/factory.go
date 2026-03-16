package client

import "github.com/instana/instana-go-client/shared/rest"

// NewRestResource creates a REST resource client with the specified mode.
// It uses a generic unmarshaller to handle JSON conversion automatically.
// This function centralizes the client creation logic that was previously duplicated across all API packages.
//
// Parameters:
//   - restClient: The REST client to use for HTTP operations
//   - resourcePath: The API endpoint path (e.g., "/api/api-tokens")
//   - mode: The create/update behavior mode (PUT/POST combinations)
//
// Returns:
//   - A fully configured REST resource client for type T
//
// Example:
//
//	// Create a client for API tokens with POST for create, PUT for update
//	tokens := client.NewRestResource[*apitoken.APIToken](
//	    restClient,
//	    apitoken.ResourcePath,
//	    rest.DefaultRestResourceModeCreatePOSTUpdatePUT,
//	)
//	allTokens, err := tokens.GetAll()
func NewRestResource[T rest.InstanaDataObject](
	restClient rest.RestClient,
	resourcePath string,
	mode rest.DefaultRestResourceMode,
	unmarshaller rest.JSONUnmarshaller[T],
) rest.RestResource[T] {
	//unmarshaller := rest.NewGenericUnmarshaller[T](),

	switch mode {
	case rest.DefaultRestResourceModeCreateAndUpdatePUT:
		return rest.NewCreatePUTUpdatePUTRestResource(resourcePath, unmarshaller, restClient)
	case rest.DefaultRestResourceModeCreatePOSTUpdatePUT:
		return rest.NewCreatePOSTUpdatePUTRestResource(resourcePath, unmarshaller, restClient)
	case rest.DefaultRestResourceModeCreateAndUpdatePOST:
		return rest.NewCreatePOSTUpdatePOSTRestResource(resourcePath, unmarshaller, restClient)
	case rest.DefaultRestResourceModeCreatePOSTAndUpdateNotSupported:
		return rest.NewCreatePOSTUpdateNotSupportedRestResource(resourcePath, unmarshaller, restClient)
	case rest.DefaultRestResourceModeCreatePUTAndUpdateNotSupported:
		return rest.NewCreatePUTUpdateNotSupportedRestResource(resourcePath, unmarshaller, restClient)
	default:
		// Default to POST for create, PUT for update (most common pattern)
		return rest.NewCreatePOSTUpdatePUTRestResource(resourcePath, unmarshaller, restClient)
	}
}

// NewReadOnlyRestResource creates a read-only REST resource client.
// This is used for resources that don't support create, update, or delete operations.
//
// Parameters:
//   - restClient: The REST client to use for HTTP operations
//   - resourcePath: The API endpoint path (e.g., "/api/events/settings/built-in-events")
//
// Returns:
//   - A read-only REST resource client for type T
//
// Example:
//
//	// Create a read-only client for built-in event specifications
//	events := client.NewReadOnlyRestResource[*builtineventspec.BuiltinEventSpecification](
//	    restClient,
//	    builtineventspec.ResourcePath,
//	)
//	allEvents, err := events.GetAll()
func NewReadOnlyRestResource[T rest.InstanaDataObject](
	restClient rest.RestClient,
	resourcePath string,
	unmarshaller rest.JSONUnmarshaller[T],
) rest.ReadOnlyRestResource[T] {
	//unmarshaller := rest.NewGenericUnmarshaller[T]()
	return rest.NewReadOnlyRestResource(resourcePath, unmarshaller, restClient)
}

// Made with Bob
