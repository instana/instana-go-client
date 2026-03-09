package client

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/instana/instana-go-client/config"
	"github.com/instana/instana-go-client/shared/rest"
	resty "gopkg.in/resty.v1"
)

// ErrEntityNotFound error message which is returned when the entity cannot be found at the server
var ErrEntityNotFound = errors.New("failed to get resource from Instana API. 404 - Resource not found")

const contentTypeHeader = "Content-Type"
const encodingApplicationJSON = "application/json; charset=utf-8"

// RestClient interface to access REST resources of the Instana API
type RestClient = rest.RestClient

// InstanaDataObject is a marker interface for any data object provided by any resource of the Instana REST API
type InstanaDataObject = rest.InstanaDataObject

type apiRequest struct {
	method          string
	url             string
	request         resty.Request
	responseChannel chan *apiResponse
	ctx             context.Context
}

type apiResponse struct {
	data []byte
	err  error
}

// NewClient creates a new instance of the Instana REST API client with default configuration
// Use NewClientWithConfig for more control over client behavior
func NewClient(apiToken string, host string, skipTlsVerification bool) RestClient {
	cfg := config.DefaultClientConfig()
	cfg.APIToken = apiToken
	cfg.BaseURL = fmt.Sprintf("https://%s", host)

	// Create HTTP client with TLS configuration
	httpClient := &http.Client{
		Timeout: cfg.Timeout.Request,
	}

	if skipTlsVerification {
		httpClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec
		}
	}

	cfg.HTTPClient = httpClient

	// Use default logger (standard log package)
	cfg.Logger = config.NewDefaultLogger(config.ClientLogLevelInfo)

	client, err := NewClientWithConfig(cfg)
	if err != nil {
		// This should never happen with default config, but handle it gracefully
		cfg.Logger.Error("Failed to create client with config", "error", err)
		// Fall back to basic client without advanced features
		return newBasicClient(apiToken, host, skipTlsVerification)
	}

	return client
}

// NewClientWithConfig creates a new instance of the Instana REST API client with custom configuration
func NewClientWithConfig(cfg *config.ClientConfig) (RestClient, error) {
	// Validate configuration
	if err := cfg.Validate(); err != nil {
		return nil, config.NewValidationError("invalid client configuration", err)
	}

	// Use default logger if not provided
	logger := cfg.Logger
	if logger == nil {
		logger = config.NewDefaultLogger(config.ClientLogLevelInfo)
	}

	// Create HTTP client if not provided
	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = createHTTPClient(cfg)
	}

	// Create resty client
	restyClient := resty.NewWithClient(httpClient)
	restyClient.SetTimeout(cfg.Timeout.Request)

	// Create rate limiter if enabled
	var rateLimiter *config.RateLimiter
	if cfg.RateLimit.Enabled {
		rateLimiter = config.NewRateLimiter(cfg.RateLimit, logger)
	}

	// Create retryer
	retryer := config.NewRetryer(cfg.Retry, logger)

	// Create throttle channel for backward compatibility
	throttledRequests := make(chan *apiRequest, 1000)

	client := &restClientImpl{
		config:            cfg,
		restyClient:       restyClient,
		logger:            logger,
		rateLimiter:       rateLimiter,
		retryer:           retryer,
		throttledRequests: throttledRequests,
		throttleRate:      time.Second / time.Duration(cfg.RateLimit.RequestsPerSecond),
	}

	// Start throttle processor for backward compatibility
	go client.processThrottledRequests()

	logger.Info("Instana REST client initialized",
		"base_url", cfg.BaseURL,
		"rate_limit_enabled", cfg.RateLimit.Enabled,
		"max_retry_attempts", cfg.Retry.MaxAttempts,
	)

	return client, nil
}

// newBasicClient creates a basic client without advanced features (fallback)
func newBasicClient(apiToken string, host string, skipTlsVerification bool) RestClient {
	restyClient := resty.New()
	if skipTlsVerification {
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) //nolint:gosec
	}

	throttleRate := time.Second / 5 // 5 write requests per second
	throttledRequests := make(chan *apiRequest, 1000)

	cfg := &config.ClientConfig{
		BaseURL:   fmt.Sprintf("https://%s", host),
		APIToken:  apiToken,
		UserAgent: "Instana-Go-Client/1.0.0",
	}

	client := &restClientImpl{
		config:            cfg,
		restyClient:       restyClient,
		logger:            config.NewNoOpLogger(),
		throttledRequests: throttledRequests,
		throttleRate:      throttleRate,
	}

	go client.processThrottledRequests()
	return client
}

type restClientImpl struct {
	config            *config.ClientConfig
	restyClient       *resty.Client
	logger            config.Logger
	rateLimiter       *config.RateLimiter
	retryer           *config.Retryer
	throttledRequests chan *apiRequest
	throttleRate      time.Duration
}

var emptyResponse = make([]byte, 0)

// Get request data via HTTP GET for the given resourcePath
func (client *restClientImpl) Get(resourcePath string) ([]byte, error) {
	url := client.buildURL(resourcePath)
	req := client.createRequest()
	return client.executeRequestWithRetry(context.Background(), resty.MethodGet, url, req)
}

// GetByQuery request data via HTTP GET for the given resourcePath and query parameters
func (client *restClientImpl) GetByQuery(resourcePath string, queryParams map[string]string) ([]byte, error) {
	url := client.buildURL(resourcePath)
	req := client.createRequest()
	client.appendQueryParameters(req, queryParams)
	return client.executeRequestWithRetry(context.Background(), resty.MethodGet, url, req)
}

// GetOne request the resource with the given ID
func (client *restClientImpl) GetOne(id string, resourcePath string) ([]byte, error) {
	url := client.buildResourceURL(resourcePath, id)
	req := client.createRequest()
	return client.executeRequestWithRetry(context.Background(), resty.MethodGet, url, req)
}

// Post executes a HTTP POST request to create or update the given resource
func (client *restClientImpl) Post(data InstanaDataObject, resourcePath string) ([]byte, error) {
	url := client.buildURL(resourcePath)
	req := client.createRequest().SetHeader(contentTypeHeader, encodingApplicationJSON).SetBody(data)
	return client.executeRequestWithThrottling(resty.MethodPost, url, req)
}

// PostWithID executes a HTTP POST request to create or update the given resource using the ID from the InstanaDataObject in the resource path
func (client *restClientImpl) PostWithID(data InstanaDataObject, resourcePath string) ([]byte, error) {
	url := client.buildResourceURL(resourcePath, data.GetIDForResourcePath())
	req := client.createRequest().SetHeader(contentTypeHeader, encodingApplicationJSON).SetBody(data)
	return client.executeRequestWithThrottling(resty.MethodPost, url, req)
}

// Put executes a HTTP PUT request to create or update the given resource
func (client *restClientImpl) Put(data InstanaDataObject, resourcePath string) ([]byte, error) {
	url := client.buildResourceURL(resourcePath, data.GetIDForResourcePath())
	req := client.createRequest().SetHeader(contentTypeHeader, encodingApplicationJSON).SetBody(data)
	return client.executeRequestWithThrottling(resty.MethodPut, url, req)
}

// Delete executes a HTTP DELETE request to delete the resource with the given ID
func (client *restClientImpl) Delete(resourceID string, resourceBasePath string) error {
	url := client.buildResourceURL(resourceBasePath, resourceID)
	req := client.createRequest()
	_, err := client.executeRequestWithThrottling(resty.MethodDelete, url, req)
	return err
}

// PostByQuery executes a HTTP POST request to create the resource by providing the data as query parameters
func (client *restClientImpl) PostByQuery(resourcePath string, queryParams map[string]string) ([]byte, error) {
	url := client.buildURL(resourcePath)
	req := client.createRequest()
	client.appendQueryParameters(req, queryParams)
	return client.executeRequestWithRetry(context.Background(), resty.MethodPost, url, req)
}

// PutByQuery executes a HTTP PUT request to update the resource with the given ID by providing the data as query parameters
func (client *restClientImpl) PutByQuery(resourcePath string, id string, queryParams map[string]string) ([]byte, error) {
	url := client.buildResourceURL(resourcePath, id)
	req := client.createRequest()
	client.appendQueryParameters(req, queryParams)
	return client.executeRequestWithRetry(context.Background(), resty.MethodPut, url, req)
}

// createRequest creates a new request with authentication and custom headers
func (client *restClientImpl) createRequest() *resty.Request {
	req := client.restyClient.R().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", fmt.Sprintf("apiToken %s", client.config.APIToken))

	// Add user agent
	userAgent := client.config.UserAgent
	if userAgent == "" {
		userAgent = "Instana-Go-Client"
	}
	req.SetHeader("User-Agent", userAgent)

	// Add custom headers
	for key, value := range client.config.Headers.Custom {
		req.SetHeader(key, value)
	}

	return req
}

// executeRequestWithRetry executes a request with retry logic
func (client *restClientImpl) executeRequestWithRetry(ctx context.Context, method string, url string, req *resty.Request) ([]byte, error) {
	// Apply rate limiting if enabled
	if client.rateLimiter != nil {
		if err := client.rateLimiter.Wait(ctx); err != nil {
			return emptyResponse, config.WrapError(err, "rate limit wait failed")
		}
	}

	// Execute with retry if retryer is available
	if client.retryer != nil {
		var result []byte
		var execErr error

		retryErr := client.retryer.Do(ctx, func() error {
			result, execErr = client.executeRequest(method, url, req)
			return execErr
		})

		if retryErr != nil {
			return emptyResponse, retryErr
		}
		return result, nil
	}

	// Execute without retry
	return client.executeRequest(method, url, req)
}

// executeRequestWithThrottling executes a request with throttling (for backward compatibility)
func (client *restClientImpl) executeRequestWithThrottling(method string, url string, req *resty.Request) ([]byte, error) {
	responseChannel := make(chan *apiResponse)
	ctx, cancel := context.WithTimeout(context.Background(), client.config.Timeout.Request)
	defer close(responseChannel)
	defer cancel()

	client.throttledRequests <- &apiRequest{
		method:          method,
		url:             url,
		request:         *req,
		ctx:             ctx,
		responseChannel: responseChannel,
	}

	select {
	case r := <-responseChannel:
		return r.data, r.err
	case <-ctx.Done():
		return nil, config.TimeoutError("API request timed out", ctx.Err())
	}
}

// processThrottledRequests processes throttled requests (for backward compatibility)
func (client *restClientImpl) processThrottledRequests() {
	throttle := time.NewTicker(client.throttleRate).C
	for req := range client.throttledRequests {
		<-throttle
		go client.handleThrottledAPIRequest(req)
	}
}

// handleThrottledAPIRequest handles a single throttled API request
func (client *restClientImpl) handleThrottledAPIRequest(req *apiRequest) {
	data, err := client.executeRequestWithRetry(req.ctx, req.method, req.url, &req.request)
	responseMessage := &apiResponse{
		data: data,
		err:  err,
	}
	select {
	case <-req.ctx.Done():
		return
	default:
		req.responseChannel <- responseMessage
	}
}

// executeRequest executes the actual HTTP request
func (client *restClientImpl) executeRequest(method string, url string, req *resty.Request) ([]byte, error) {
	client.logger.Debug("Executing HTTP request",
		"method", method,
		"url", client.redactURL(url),
	)

	resp, err := req.Execute(method, url)
	if err != nil {
		if resp == nil {
			return emptyResponse, config.NetworkError(
				fmt.Sprintf("failed to send HTTP %s request to Instana API", method),
				err,
			)
		}
		return emptyResponse, config.APIError(
			resp.StatusCode(),
			fmt.Sprintf("failed to send HTTP %s request to Instana API: %s", method, string(resp.Body())),
			err,
		)
	}

	statusCode := resp.StatusCode()
	client.logger.Info("HTTP response received",
		"method", method,
		"status_code", statusCode,
		"status", resp.Status(),
	)

	// Handle specific status codes
	if statusCode == 404 {
		return emptyResponse, ErrEntityNotFound
	}

	if statusCode == 401 || statusCode == 403 {
		return emptyResponse, config.AuthenticationError(
			fmt.Sprintf("authentication failed: %s", resp.Status()),
			nil,
		)
	}

	if statusCode == 429 {
		return emptyResponse, config.RateLimitError(
			"rate limit exceeded",
			0, // No retry after header in current implementation
		)
	}

	if statusCode < 200 || statusCode >= 300 {
		return emptyResponse, config.APIError(
			statusCode,
			fmt.Sprintf("HTTP %s request failed: %s", method, string(resp.Body())),
			nil,
		)
	}

	// Handle 204 No Content - return empty JSON object instead of empty body
	if statusCode == 204 {
		client.logger.Debug("Received 204 No Content, returning empty JSON object")
		return []byte("{}"), nil
	}

	return resp.Body(), nil
}

// appendQueryParameters appends query parameters to the request
func (client *restClientImpl) appendQueryParameters(req *resty.Request, queryParams map[string]string) {
	for k, v := range queryParams {
		req.QueryParam.Add(k, v)
	}
}

// buildResourceURL builds the full URL for a resource with ID
func (client *restClientImpl) buildResourceURL(resourceBasePath string, id string) string {
	pattern := "%s/%s"
	if strings.HasSuffix(resourceBasePath, "/") {
		pattern = "%s%s"
	}
	resourcePath := fmt.Sprintf(pattern, resourceBasePath, id)
	return client.buildURL(resourcePath)
}

// buildURL builds the full URL from base URL and resource path
func (client *restClientImpl) buildURL(resourcePath string) string {
	baseURL := strings.TrimSuffix(client.config.BaseURL, "/")
	if !strings.HasPrefix(resourcePath, "/") {
		resourcePath = "/" + resourcePath
	}
	return fmt.Sprintf("%s%s", baseURL, resourcePath)
}

// redactURL redacts sensitive information from URL for logging
func (client *restClientImpl) redactURL(url string) string {
	// Redact API token if present in URL (shouldn't be, but be safe)
	if client.config.APIToken != "" {
		url = strings.ReplaceAll(url, client.config.APIToken, "***REDACTED***")
	}
	return url
}
