package instana

import (
	"crypto/tls"
	"net/http"

	"github.com/instana/instana-go-client/client"
	"github.com/instana/instana-go-client/config"
)

// NewInstanaAPI creates a new instance of the Instana API client with basic configuration.
// This is the primary entry point for creating an Instana API client.
//
// Parameters:
//   - apiToken: The API token for authentication
//   - endpoint: The Instana endpoint (e.g., "tenant-unit.instana.io")
//   - skipTlsVerification: Whether to skip TLS certificate verification (use with caution)
//
// Returns:
//   - client.InstanaAPI: The API client interface for accessing all Instana resources
//
// Example:
//
//	api := instana.NewInstanaAPI("your-api-token", "tenant-unit.instana.io", false)
//	tokens, err := api.APITokens().GetAll()
func NewInstanaAPI(apiToken string, endpoint string, skipTlsVerification bool) client.InstanaAPI {
	restClient := NewClient(apiToken, endpoint, skipTlsVerification)
	return client.NewInstanaAPI(restClient)
}

// NewInstanaAPIWithUserAgent creates a new instance of the Instana API client with a custom user agent.
//
// Parameters:
//   - apiToken: The API token for authentication
//   - endpoint: The Instana endpoint (e.g., "tenant-unit.instana.io")
//   - skipTlsVerification: Whether to skip TLS certificate verification (use with caution)
//   - userAgent: Custom user agent string (e.g., "Terraform/1.2.3")
//
// Returns:
//   - client.InstanaAPI: The API client interface for accessing all Instana resources
//
// Example:
//
//	api := instana.NewInstanaAPIWithUserAgent(
//	    "your-api-token",
//	    "tenant-unit.instana.io",
//	    false,
//	    "MyApp/1.0.0",
//	)
func NewInstanaAPIWithUserAgent(apiToken string, endpoint string, skipTlsVerification bool, userAgent string) client.InstanaAPI {
	cfg := config.DefaultClientConfig()
	cfg.APIToken = apiToken
	cfg.BaseURL = "https://" + endpoint
	cfg.UserAgent = userAgent
	cfg.Logger = config.NewDefaultLogger(config.ClientLogLevelInfo)

	// Create HTTP client with TLS configuration
	httpClient := createHTTPClient(cfg)
	if skipTlsVerification {
		httpClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec
		}
	}
	cfg.HTTPClient = httpClient

	restClient, err := NewClientWithConfig(cfg)
	if err != nil {
		// Fall back to basic client if config fails
		cfg.Logger.Error("Failed to create client with config, using basic client", "error", err)
		restClient = NewClient(apiToken, endpoint, skipTlsVerification)
	}

	return client.NewInstanaAPI(restClient)
}

// NewInstanaAPIWithConfig creates a new instance of the Instana API client with full configuration control.
// This provides the most flexibility for advanced use cases.
//
// Parameters:
//   - config: Complete client configuration
//
// Returns:
//   - client.InstanaAPI: The API client interface for accessing all Instana resources
//   - error: Any error that occurred during client creation
//
// Example:
//
//	config := instana.DefaultClientConfig()
//	config.APIToken = "your-api-token"
//	config.BaseURL = "https://tenant-unit.instana.io"
//	config.Retry.MaxAttempts = 5
//	config.RateLimit.Enabled = true
//
//	api, err := instana.NewInstanaAPIWithConfig(config)
//	if err != nil {
//	    log.Fatal(err)
//	}
func NewInstanaAPIWithConfig(cfg *config.ClientConfig) (client.InstanaAPI, error) {
	restClient, err := NewClientWithConfig(cfg)
	if err != nil {
		return nil, err
	}

	return client.NewInstanaAPI(restClient), nil
}

// createHTTPClient creates an HTTP client with connection pooling configuration.
// This is an internal helper function used by the initialization methods.
func createHTTPClient(cfg *config.ClientConfig) *http.Client {
	transport := &http.Transport{
		MaxIdleConns:        cfg.ConnectionPool.MaxIdleConnections,
		MaxIdleConnsPerHost: cfg.ConnectionPool.MaxConnectionsPerHost,
		IdleConnTimeout:     cfg.Timeout.IdleConnection,
		DisableKeepAlives:   false,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   cfg.Timeout.Request,
	}
}
