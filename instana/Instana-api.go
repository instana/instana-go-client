package instana

import (
	"crypto/tls"
	"net/http"

	"github.com/instana/instana-go-client/client"
	"github.com/instana/instana-go-client/config"
)

const (
	//InstanaAPIBasePath path to Instana RESTful API
	InstanaAPIBasePath = "/api"
	//EventsBasePath path to Events resource of Instana RESTful API
	EventsBasePath = InstanaAPIBasePath + "/events"
	//settingsPathElement path element to settings
	settingsPathElement = "/settings"
	//EventSettingsBasePath path to Event Settings resource of Instana RESTful API
	EventSettingsBasePath = EventsBasePath + settingsPathElement
	//SettingsBasePath path to Event Settings resource of Instana RESTful API
	SettingsBasePath = InstanaAPIBasePath + settingsPathElement
	//RBACSettingsBasePath path to Role Based Access Control Settings resources of Instana RESTful API
	RBACSettingsBasePath = SettingsBasePath + "/rbac"
	//WebsiteMonitoringResourcePath path to website monitoring
	WebsiteMonitoringResourcePath = InstanaAPIBasePath + "/website-monitoring"
	//SyntheticSettingsBasePath path to synthetic monitoring
	SyntheticSettingsBasePath = InstanaAPIBasePath + "/synthetics" + settingsPathElement
	//SyntheticTestResourcePath path to synthetic monitoring tests
	SyntheticTestResourcePath = SyntheticSettingsBasePath + "/tests"
	//SyntheticLocationResourcePath path to synthetic monitoring tests
	SyntheticLocationResourcePath = SyntheticSettingsBasePath + "/locations"
	// AutomationBasePath path to Automation resources of Instana RESTful API
	AutomationBasePath = InstanaAPIBasePath + "/automation"
	// HostAgentResourcePath path to host agent resources of Instana RESTful API
	HostAgentResourcePath = InstanaAPIBasePath + "/host-agent"
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
//
// Deprecated: This function is maintained for backward compatibility.
// New code should use the client package directly via client.NewInstanaAPI().
func NewInstanaAPI(apiToken string, endpoint string, skipTlsVerification bool) client.InstanaAPI {
	restClient := NewClient(apiToken, endpoint, skipTlsVerification)
	return client.NewInstanaAPI(restClient)
}

// NewInstanaAPIWithUserAgent creates a new instance of the Instana API client with a custom user agent.
// This is useful for identifying the client application in API logs and metrics.
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
//
// Deprecated: This function is maintained for backward compatibility.
// New code should use the client package directly via client.NewInstanaAPI().
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
//
// Deprecated: This function is maintained for backward compatibility.
// New code should use the client package directly via client.NewInstanaAPI().
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
