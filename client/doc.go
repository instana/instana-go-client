// Package client provides the main interface for interacting with the Instana API.
//
// This package contains the InstanaAPI interface which serves as the primary entry point
// for all API operations. It provides access to all 28 API endpoints through dedicated
// client methods with lazy initialization for optimal performance.
//
// Example usage:
//
//	import (
//	    "github.com/instana/instana-go-client/client"
//	    "github.com/instana/instana-go-client/instana"
//	)
//
//	// Create REST client with configuration
//	config := instana.DefaultClientConfig()
//	config.APIToken = "your-api-token"
//	config.Host = "https://your-tenant.instana.io"
//
//	restClient, err := instana.NewRestClient(config)
//	if err != nil {
//	    // handle error
//	}
//
//	// Create Instana API client
//	api := client.NewInstanaAPI(restClient)
//
//	// Use API clients (lazy initialization)
//	tokens, err := api.APITokens().GetAll()
//	if err != nil {
//	    // handle error
//	}
//
//	// Create a new API token
//	newToken := &apitoken.APIToken{
//	    Name: "My Token",
//	    CanConfigureApplications: true,
//	}
//	created, err := api.APITokens().Create(newToken)
//	if err != nil {
//	    // handle error
//	}
//
// The client package uses lazy initialization, meaning API clients are only created
// when first accessed. This improves performance by avoiding unnecessary initialization
// of unused API clients.
//
// All API methods return either rest.RestResource or rest.ReadOnlyRestResource interfaces,
// providing a consistent interface for CRUD operations across all endpoints.
package client
