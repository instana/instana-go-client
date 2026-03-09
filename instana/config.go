// Package instana provides a Go client library for the Instana API.
package instana

import (
	"net/http"
	"time"
)

// ClientConfig holds all configuration options for the Instana API client.
// It provides a flexible way to customize client behavior including timeouts,
// retry logic, rate limiting, and connection pooling.
type ClientConfig struct {
	// BaseURL is the base URL of the Instana API (e.g., "https://tenant-unit.instana.io")
	BaseURL string

	// APIToken is the API token used for authentication
	APIToken string

	// Timeout configuration for various timeout scenarios
	Timeout TimeoutConfig

	// Retry configuration for automatic retry behavior
	Retry RetryConfig

	// Headers configuration for custom HTTP headers
	Headers HeaderConfig

	// Batch configuration for batch operations
	Batch BatchConfig

	// RateLimit configuration for rate limiting
	RateLimit RateLimitConfig

	// ConnectionPool configuration for HTTP connection pooling
	ConnectionPool ConnectionPoolConfig

	// Logger is the logger instance for client operations
	// If nil, a default logger will be used
	Logger Logger

	// HTTPClient is a custom HTTP client to use for requests
	// If nil, a default client will be created based on other config options
	HTTPClient *http.Client

	// UserAgent is the User-Agent header value
	// Default: "instana-go-client/v{version}"
	UserAgent string

	// Debug enables debug logging when true
	Debug bool
}

// TimeoutConfig holds timeout-related configuration.
type TimeoutConfig struct {
	// Connection is the timeout for establishing a connection
	// Default: 30 seconds
	Connection time.Duration

	// Request is the timeout for the entire request (including retries)
	// Default: 60 seconds
	Request time.Duration

	// IdleConnection is the timeout for idle connections in the pool
	// Default: 90 seconds
	IdleConnection time.Duration

	// ResponseHeader is the timeout for reading response headers
	// Default: 10 seconds
	ResponseHeader time.Duration

	// TLSHandshake is the timeout for TLS handshake
	// Default: 10 seconds
	TLSHandshake time.Duration
}

// RetryConfig holds retry-related configuration.
type RetryConfig struct {
	// MaxAttempts is the maximum number of retry attempts
	// Default: 3
	MaxAttempts int

	// InitialDelay is the initial delay before the first retry
	// Default: 1 second
	InitialDelay time.Duration

	// MaxDelay is the maximum delay between retries
	// Default: 30 seconds
	MaxDelay time.Duration

	// BackoffMultiplier is the multiplier for exponential backoff
	// Default: 2.0
	BackoffMultiplier float64

	// RetryableStatusCodes is a list of HTTP status codes that should trigger a retry
	// Default: [408, 429, 500, 502, 503, 504]
	RetryableStatusCodes []int

	// RetryOnTimeout enables retry on timeout errors
	// Default: true
	RetryOnTimeout bool

	// RetryOnConnectionError enables retry on connection errors
	// Default: true
	RetryOnConnectionError bool

	// Jitter enables adding random jitter to retry delays
	// Default: true
	Jitter bool
}

// HeaderConfig holds custom HTTP header configuration.
type HeaderConfig struct {
	// Custom is a map of custom headers to include in all requests
	Custom map[string]string

	// DisableDefaultHeaders disables default headers (User-Agent, etc.)
	// Default: false
	DisableDefaultHeaders bool
}

// BatchConfig holds batch operation configuration.
type BatchConfig struct {
	// Size is the maximum number of items in a single batch
	// Default: 100
	Size int

	// ConcurrentRequests is the maximum number of concurrent batch requests
	// Default: 5
	ConcurrentRequests int

	// StopOnError stops batch processing on first error when true
	// Default: false
	StopOnError bool

	// RetryFailedItems enables retry of failed items in a batch
	// Default: true
	RetryFailedItems bool
}

// RateLimitConfig holds rate limiting configuration.
type RateLimitConfig struct {
	// RequestsPerSecond is the maximum number of requests per second
	// Default: 100
	RequestsPerSecond int

	// BurstCapacity is the maximum burst capacity for rate limiting
	// Default: 200
	BurstCapacity int

	// Enabled enables rate limiting when true
	// Default: true
	Enabled bool

	// WaitForToken waits for rate limit token when true, otherwise returns error
	// Default: true
	WaitForToken bool
}

// ConnectionPoolConfig holds HTTP connection pool configuration.
type ConnectionPoolConfig struct {
	// MaxIdleConnections is the maximum number of idle connections
	// Default: 100
	MaxIdleConnections int

	// MaxConnectionsPerHost is the maximum number of connections per host
	// Default: 10
	MaxConnectionsPerHost int

	// MaxIdleConnectionsPerHost is the maximum number of idle connections per host
	// Default: 10
	MaxIdleConnectionsPerHost int

	// KeepAliveDuration is the keep-alive duration for connections
	// Default: 30 seconds
	KeepAliveDuration time.Duration

	// DisableKeepAlives disables HTTP keep-alives
	// Default: false
	DisableKeepAlives bool

	// DisableCompression disables HTTP compression
	// Default: false
	DisableCompression bool
}

// DefaultClientConfig returns a ClientConfig with sensible default values.
func DefaultClientConfig() *ClientConfig {
	return &ClientConfig{
		Timeout: TimeoutConfig{
			Connection:     30 * time.Second,
			Request:        60 * time.Second,
			IdleConnection: 90 * time.Second,
			ResponseHeader: 10 * time.Second,
			TLSHandshake:   10 * time.Second,
		},
		Retry: RetryConfig{
			MaxAttempts:            3,
			InitialDelay:           1 * time.Second,
			MaxDelay:               30 * time.Second,
			BackoffMultiplier:      2.0,
			RetryableStatusCodes:   []int{408, 429, 500, 502, 503, 504},
			RetryOnTimeout:         true,
			RetryOnConnectionError: true,
			Jitter:                 true,
		},
		Headers: HeaderConfig{
			Custom:                make(map[string]string),
			DisableDefaultHeaders: false,
		},
		Batch: BatchConfig{
			Size:               100,
			ConcurrentRequests: 5,
			StopOnError:        false,
			RetryFailedItems:   true,
		},
		RateLimit: RateLimitConfig{
			RequestsPerSecond: 100,
			BurstCapacity:     200,
			Enabled:           true,
			WaitForToken:      true,
		},
		ConnectionPool: ConnectionPoolConfig{
			MaxIdleConnections:        100,
			MaxConnectionsPerHost:     10,
			MaxIdleConnectionsPerHost: 10,
			KeepAliveDuration:         30 * time.Second,
			DisableKeepAlives:         false,
			DisableCompression:        false,
		},
		UserAgent: "instana-go-client/v1.0.0",
		Debug:     false,
	}
}

// Clone creates a deep copy of the ClientConfig.
func (c *ClientConfig) Clone() *ClientConfig {
	if c == nil {
		return nil
	}

	clone := &ClientConfig{
		BaseURL:        c.BaseURL,
		APIToken:       c.APIToken,
		Timeout:        c.Timeout,
		Retry:          c.Retry,
		Batch:          c.Batch,
		RateLimit:      c.RateLimit,
		ConnectionPool: c.ConnectionPool,
		Logger:         c.Logger,
		HTTPClient:     c.HTTPClient,
		UserAgent:      c.UserAgent,
		Debug:          c.Debug,
	}

	// Deep copy Headers.Custom map
	clone.Headers.Custom = make(map[string]string, len(c.Headers.Custom))
	for k, v := range c.Headers.Custom {
		clone.Headers.Custom[k] = v
	}
	clone.Headers.DisableDefaultHeaders = c.Headers.DisableDefaultHeaders

	// Deep copy RetryableStatusCodes slice
	if c.Retry.RetryableStatusCodes != nil {
		clone.Retry.RetryableStatusCodes = make([]int, len(c.Retry.RetryableStatusCodes))
		copy(clone.Retry.RetryableStatusCodes, c.Retry.RetryableStatusCodes)
	}

	return clone
}

// Made with Bob
