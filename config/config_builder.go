package config

import (
	"net/http"
	"time"
)

// ConfigBuilder provides a fluent interface for building ClientConfig
type ConfigBuilder struct {
	config *ClientConfig
}

// NewConfigBuilder creates a new ConfigBuilder with default configuration
func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{
		config: DefaultClientConfig(),
	}
}

// NewConfigBuilderFromConfig creates a new ConfigBuilder from an existing config
func NewConfigBuilderFromConfig(config *ClientConfig) *ConfigBuilder {
	return &ConfigBuilder{
		config: config.Clone(),
	}
}

// WithBaseURL sets the base URL
func (b *ConfigBuilder) WithBaseURL(baseURL string) *ConfigBuilder {
	b.config.BaseURL = baseURL
	return b
}

// WithAPIToken sets the API token
func (b *ConfigBuilder) WithAPIToken(token string) *ConfigBuilder {
	b.config.APIToken = token
	return b
}

// WithUserAgent sets the user agent
func (b *ConfigBuilder) WithUserAgent(userAgent string) *ConfigBuilder {
	b.config.UserAgent = userAgent
	return b
}

// WithDebug enables or disables debug mode
func (b *ConfigBuilder) WithDebug(debug bool) *ConfigBuilder {
	b.config.Debug = debug
	return b
}

// WithLogger sets a custom logger
func (b *ConfigBuilder) WithLogger(logger Logger) *ConfigBuilder {
	b.config.Logger = logger
	return b
}

// WithHTTPClient sets a custom HTTP client
func (b *ConfigBuilder) WithHTTPClient(client *http.Client) *ConfigBuilder {
	b.config.HTTPClient = client
	return b
}

// Timeout configuration methods

// WithConnectionTimeout sets the connection timeout
func (b *ConfigBuilder) WithConnectionTimeout(timeout time.Duration) *ConfigBuilder {
	b.config.Timeout.Connection = timeout
	return b
}

// WithRequestTimeout sets the request timeout
func (b *ConfigBuilder) WithRequestTimeout(timeout time.Duration) *ConfigBuilder {
	b.config.Timeout.Request = timeout
	return b
}

// WithIdleConnectionTimeout sets the idle connection timeout
func (b *ConfigBuilder) WithIdleConnectionTimeout(timeout time.Duration) *ConfigBuilder {
	b.config.Timeout.IdleConnection = timeout
	return b
}

// WithResponseHeaderTimeout sets the response header timeout
func (b *ConfigBuilder) WithResponseHeaderTimeout(timeout time.Duration) *ConfigBuilder {
	b.config.Timeout.ResponseHeader = timeout
	return b
}

// WithTLSHandshakeTimeout sets the TLS handshake timeout
func (b *ConfigBuilder) WithTLSHandshakeTimeout(timeout time.Duration) *ConfigBuilder {
	b.config.Timeout.TLSHandshake = timeout
	return b
}

// WithTimeoutConfig sets the entire timeout configuration
func (b *ConfigBuilder) WithTimeoutConfig(config TimeoutConfig) *ConfigBuilder {
	b.config.Timeout = config
	return b
}

// Retry configuration methods

// WithMaxRetryAttempts sets the maximum number of retry attempts
func (b *ConfigBuilder) WithMaxRetryAttempts(attempts int) *ConfigBuilder {
	b.config.Retry.MaxAttempts = attempts
	return b
}

// WithRetryInitialDelay sets the initial retry delay
func (b *ConfigBuilder) WithRetryInitialDelay(delay time.Duration) *ConfigBuilder {
	b.config.Retry.InitialDelay = delay
	return b
}

// WithRetryMaxDelay sets the maximum retry delay
func (b *ConfigBuilder) WithRetryMaxDelay(delay time.Duration) *ConfigBuilder {
	b.config.Retry.MaxDelay = delay
	return b
}

// WithRetryBackoffMultiplier sets the backoff multiplier
func (b *ConfigBuilder) WithRetryBackoffMultiplier(multiplier float64) *ConfigBuilder {
	b.config.Retry.BackoffMultiplier = multiplier
	return b
}

// WithRetryableStatusCodes sets the retryable HTTP status codes
func (b *ConfigBuilder) WithRetryableStatusCodes(codes []int) *ConfigBuilder {
	b.config.Retry.RetryableStatusCodes = codes
	return b
}

// WithRetryOnTimeout enables or disables retry on timeout
func (b *ConfigBuilder) WithRetryOnTimeout(retry bool) *ConfigBuilder {
	b.config.Retry.RetryOnTimeout = retry
	return b
}

// WithRetryOnConnectionError enables or disables retry on connection errors
func (b *ConfigBuilder) WithRetryOnConnectionError(retry bool) *ConfigBuilder {
	b.config.Retry.RetryOnConnectionError = retry
	return b
}

// WithRetryJitter enables or disables jitter in retry delays
func (b *ConfigBuilder) WithRetryJitter(jitter bool) *ConfigBuilder {
	b.config.Retry.Jitter = jitter
	return b
}

// WithRetryConfig sets the entire retry configuration
func (b *ConfigBuilder) WithRetryConfig(config RetryConfig) *ConfigBuilder {
	b.config.Retry = config
	return b
}

// Header configuration methods

// WithCustomHeader adds a custom header
func (b *ConfigBuilder) WithCustomHeader(key, value string) *ConfigBuilder {
	if b.config.Headers.Custom == nil {
		b.config.Headers.Custom = make(map[string]string)
	}
	b.config.Headers.Custom[key] = value
	return b
}

// WithCustomHeaders sets multiple custom headers
func (b *ConfigBuilder) WithCustomHeaders(headers map[string]string) *ConfigBuilder {
	if b.config.Headers.Custom == nil {
		b.config.Headers.Custom = make(map[string]string)
	}
	for k, v := range headers {
		b.config.Headers.Custom[k] = v
	}
	return b
}

// WithDisableDefaultHeaders disables default headers
func (b *ConfigBuilder) WithDisableDefaultHeaders(disable bool) *ConfigBuilder {
	b.config.Headers.DisableDefaultHeaders = disable
	return b
}

// WithHeaderConfig sets the entire header configuration
func (b *ConfigBuilder) WithHeaderConfig(config HeaderConfig) *ConfigBuilder {
	b.config.Headers = config
	return b
}

// Batch configuration methods

// WithBatchSize sets the batch size
func (b *ConfigBuilder) WithBatchSize(size int) *ConfigBuilder {
	b.config.Batch.Size = size
	return b
}

// WithBatchConcurrentRequests sets the number of concurrent batch requests
func (b *ConfigBuilder) WithBatchConcurrentRequests(concurrent int) *ConfigBuilder {
	b.config.Batch.ConcurrentRequests = concurrent
	return b
}

// WithBatchStopOnError enables or disables stopping on first error
func (b *ConfigBuilder) WithBatchStopOnError(stop bool) *ConfigBuilder {
	b.config.Batch.StopOnError = stop
	return b
}

// WithBatchRetryFailedItems enables or disables retry of failed items
func (b *ConfigBuilder) WithBatchRetryFailedItems(retry bool) *ConfigBuilder {
	b.config.Batch.RetryFailedItems = retry
	return b
}

// WithBatchConfig sets the entire batch configuration
func (b *ConfigBuilder) WithBatchConfig(config BatchConfig) *ConfigBuilder {
	b.config.Batch = config
	return b
}

// Rate limit configuration methods

// WithRateLimitEnabled enables or disables rate limiting
func (b *ConfigBuilder) WithRateLimitEnabled(enabled bool) *ConfigBuilder {
	b.config.RateLimit.Enabled = enabled
	return b
}

// WithRateLimitRequestsPerSecond sets the requests per second limit
func (b *ConfigBuilder) WithRateLimitRequestsPerSecond(rps int) *ConfigBuilder {
	b.config.RateLimit.RequestsPerSecond = rps
	return b
}

// WithRateLimitBurstCapacity sets the burst capacity
func (b *ConfigBuilder) WithRateLimitBurstCapacity(burst int) *ConfigBuilder {
	b.config.RateLimit.BurstCapacity = burst
	return b
}

// WithRateLimitWaitForToken enables or disables waiting for rate limit token
func (b *ConfigBuilder) WithRateLimitWaitForToken(wait bool) *ConfigBuilder {
	b.config.RateLimit.WaitForToken = wait
	return b
}

// WithRateLimitConfig sets the entire rate limit configuration
func (b *ConfigBuilder) WithRateLimitConfig(config RateLimitConfig) *ConfigBuilder {
	b.config.RateLimit = config
	return b
}

// Connection pool configuration methods

// WithMaxIdleConnections sets the maximum number of idle connections
func (b *ConfigBuilder) WithMaxIdleConnections(max int) *ConfigBuilder {
	b.config.ConnectionPool.MaxIdleConnections = max
	return b
}

// WithMaxConnectionsPerHost sets the maximum connections per host
func (b *ConfigBuilder) WithMaxConnectionsPerHost(max int) *ConfigBuilder {
	b.config.ConnectionPool.MaxConnectionsPerHost = max
	return b
}

// WithMaxIdleConnectionsPerHost sets the maximum idle connections per host
func (b *ConfigBuilder) WithMaxIdleConnectionsPerHost(max int) *ConfigBuilder {
	b.config.ConnectionPool.MaxIdleConnectionsPerHost = max
	return b
}

// WithKeepAliveDuration sets the keep-alive duration
func (b *ConfigBuilder) WithKeepAliveDuration(duration time.Duration) *ConfigBuilder {
	b.config.ConnectionPool.KeepAliveDuration = duration
	return b
}

// WithDisableKeepAlives disables HTTP keep-alives
func (b *ConfigBuilder) WithDisableKeepAlives(disable bool) *ConfigBuilder {
	b.config.ConnectionPool.DisableKeepAlives = disable
	return b
}

// WithDisableCompression disables HTTP compression
func (b *ConfigBuilder) WithDisableCompression(disable bool) *ConfigBuilder {
	b.config.ConnectionPool.DisableCompression = disable
	return b
}

// WithConnectionPoolConfig sets the entire connection pool configuration
func (b *ConfigBuilder) WithConnectionPoolConfig(config ConnectionPoolConfig) *ConfigBuilder {
	b.config.ConnectionPool = config
	return b
}

// Build validates and returns the final ClientConfig
func (b *ConfigBuilder) Build() (*ClientConfig, error) {
	if err := b.config.Validate(); err != nil {
		return nil, err
	}
	return b.config, nil
}

// MustBuild validates and returns the final ClientConfig, panics on error
func (b *ConfigBuilder) MustBuild() *ClientConfig {
	config, err := b.Build()
	if err != nil {
		panic(err)
	}
	return config
}

// GetConfig returns the current configuration without validation
func (b *ConfigBuilder) GetConfig() *ClientConfig {
	return b.config
}
