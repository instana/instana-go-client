package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Error message templates
const (
	errInvalidValue = "invalid value for %s: %w"
)

// Environment variable names for configuration
const (
	EnvBaseURL                    = "INSTANA_BASE_URL"
	EnvAPIToken                   = "INSTANA_API_TOKEN" //nolint:gosec // This is an environment variable name, not a credential
	EnvDebug                      = "INSTANA_DEBUG"
	EnvConnectionTimeout          = "INSTANA_CONNECTION_TIMEOUT"
	EnvRequestTimeout             = "INSTANA_REQUEST_TIMEOUT"
	EnvIdleConnectionTimeout      = "INSTANA_IDLE_CONNECTION_TIMEOUT"
	EnvMaxRetryAttempts           = "INSTANA_MAX_RETRY_ATTEMPTS"
	EnvRetryInitialDelay          = "INSTANA_RETRY_INITIAL_DELAY"
	EnvRetryMaxDelay              = "INSTANA_RETRY_MAX_DELAY"
	EnvRetryBackoffMultiplier     = "INSTANA_RETRY_BACKOFF_MULTIPLIER"
	EnvBatchSize                  = "INSTANA_BATCH_SIZE"
	EnvBatchConcurrentRequests    = "INSTANA_BATCH_CONCURRENT_REQUESTS"
	EnvRateLimitEnabled           = "INSTANA_RATE_LIMIT_ENABLED"
	EnvRateLimitRequestsPerSecond = "INSTANA_RATE_LIMIT_RPS"
	EnvRateLimitBurstCapacity     = "INSTANA_RATE_LIMIT_BURST"
	EnvMaxIdleConnections         = "INSTANA_MAX_IDLE_CONNECTIONS"
	EnvMaxConnectionsPerHost      = "INSTANA_MAX_CONNECTIONS_PER_HOST"
	EnvKeepAliveDuration          = "INSTANA_KEEP_ALIVE_DURATION"
)

// LoadFromEnv loads configuration from environment variables
// It starts with default configuration and overrides with environment variables
func LoadFromEnv() (*ClientConfig, error) {
	config := DefaultClientConfig()

	// Load base configuration
	if v := os.Getenv(EnvBaseURL); v != "" {
		config.BaseURL = v
	}

	if v := os.Getenv(EnvAPIToken); v != "" {
		config.APIToken = v
	}

	if v := os.Getenv(EnvDebug); v != "" {
		debug, err := strconv.ParseBool(v)
		if err != nil {
			return nil, fmt.Errorf(errInvalidValue, EnvDebug, err)
		}
		config.Debug = debug
	}

	// Load timeout configuration
	if v := os.Getenv(EnvConnectionTimeout); v != "" {
		timeout, err := parseDuration(v)
		if err != nil {
			return nil, fmt.Errorf(errInvalidValue, EnvConnectionTimeout, err)
		}
		config.Timeout.Connection = timeout
	}

	if v := os.Getenv(EnvRequestTimeout); v != "" {
		timeout, err := parseDuration(v)
		if err != nil {
			return nil, fmt.Errorf(errInvalidValue, EnvRequestTimeout, err)
		}
		config.Timeout.Request = timeout
	}

	if v := os.Getenv(EnvIdleConnectionTimeout); v != "" {
		timeout, err := parseDuration(v)
		if err != nil {
			return nil, fmt.Errorf(errInvalidValue, EnvIdleConnectionTimeout, err)
		}
		config.Timeout.IdleConnection = timeout
	}

	// Load retry configuration
	if v := os.Getenv(EnvMaxRetryAttempts); v != "" {
		attempts, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf(errInvalidValue, EnvMaxRetryAttempts, err)
		}
		config.Retry.MaxAttempts = attempts
	}

	if v := os.Getenv(EnvRetryInitialDelay); v != "" {
		delay, err := parseDuration(v)
		if err != nil {
			return nil, fmt.Errorf(errInvalidValue, EnvRetryInitialDelay, err)
		}
		config.Retry.InitialDelay = delay
	}

	if v := os.Getenv(EnvRetryMaxDelay); v != "" {
		delay, err := parseDuration(v)
		if err != nil {
			return nil, fmt.Errorf(errInvalidValue, EnvRetryMaxDelay, err)
		}
		config.Retry.MaxDelay = delay
	}

	if v := os.Getenv(EnvRetryBackoffMultiplier); v != "" {
		multiplier, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, fmt.Errorf(errInvalidValue, EnvRetryBackoffMultiplier, err)
		}
		config.Retry.BackoffMultiplier = multiplier
	}

	// Load batch configuration
	if v := os.Getenv(EnvBatchSize); v != "" {
		size, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf(errInvalidValue, EnvBatchSize, err)
		}
		config.Batch.Size = size
	}

	if v := os.Getenv(EnvBatchConcurrentRequests); v != "" {
		concurrent, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf(errInvalidValue, EnvBatchConcurrentRequests, err)
		}
		config.Batch.ConcurrentRequests = concurrent
	}

	// Load rate limit configuration
	if v := os.Getenv(EnvRateLimitEnabled); v != "" {
		enabled, err := strconv.ParseBool(v)
		if err != nil {
			return nil, fmt.Errorf(errInvalidValue, EnvRateLimitEnabled, err)
		}
		config.RateLimit.Enabled = enabled
	}

	if v := os.Getenv(EnvRateLimitRequestsPerSecond); v != "" {
		rps, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf(errInvalidValue, EnvRateLimitRequestsPerSecond, err)
		}
		config.RateLimit.RequestsPerSecond = rps
	}

	if v := os.Getenv(EnvRateLimitBurstCapacity); v != "" {
		burst, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf(errInvalidValue, EnvRateLimitBurstCapacity, err)
		}
		config.RateLimit.BurstCapacity = burst
	}

	// Load connection pool configuration
	if v := os.Getenv(EnvMaxIdleConnections); v != "" {
		max, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf(errInvalidValue, EnvMaxIdleConnections, err)
		}
		config.ConnectionPool.MaxIdleConnections = max
	}

	if v := os.Getenv(EnvMaxConnectionsPerHost); v != "" {
		max, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf(errInvalidValue, EnvMaxConnectionsPerHost, err)
		}
		config.ConnectionPool.MaxConnectionsPerHost = max
	}

	if v := os.Getenv(EnvKeepAliveDuration); v != "" {
		duration, err := parseDuration(v)
		if err != nil {
			return nil, fmt.Errorf(errInvalidValue, EnvKeepAliveDuration, err)
		}
		config.ConnectionPool.KeepAliveDuration = duration
	}

	return config, nil
}

// LoadFromJSON loads configuration from a JSON file
func LoadFromJSON(filename string) (*ClientConfig, error) {
	data, err := os.ReadFile(filename) //nolint:gosec // filename is provided by user, not from external input
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config ClientConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse JSON config: %w", err)
	}

	// Set defaults for any missing fields
	applyDefaults(&config)

	return &config, nil
}

// LoadFromJSONWithEnvOverride loads configuration from JSON and overrides with environment variables
func LoadFromJSONWithEnvOverride(filename string) (*ClientConfig, error) {
	// Load from JSON first
	config, err := LoadFromJSON(filename)
	if err != nil {
		return nil, err
	}

	// Override with environment variables
	envConfig, err := LoadFromEnv()
	if err != nil {
		return nil, err
	}

	// Merge configurations (env takes precedence)
	mergeConfigs(config, envConfig)

	return config, nil
}

// parseDuration parses a duration string
// Supports formats like "30s", "1m", "1h", or plain seconds as integer
func parseDuration(s string) (time.Duration, error) {
	// Try parsing as duration string first
	if d, err := time.ParseDuration(s); err == nil {
		return d, nil
	}

	// Try parsing as seconds (integer)
	if seconds, err := strconv.Atoi(s); err == nil {
		return time.Duration(seconds) * time.Second, nil
	}

	return 0, fmt.Errorf("invalid duration format: %s", s)
}

// applyDefaults applies default values to any zero-value fields in the config
func applyDefaults(config *ClientConfig) {
	defaults := DefaultClientConfig()

	if config.Timeout.Connection == 0 {
		config.Timeout.Connection = defaults.Timeout.Connection
	}
	if config.Timeout.Request == 0 {
		config.Timeout.Request = defaults.Timeout.Request
	}
	if config.Timeout.IdleConnection == 0 {
		config.Timeout.IdleConnection = defaults.Timeout.IdleConnection
	}
	if config.Timeout.ResponseHeader == 0 {
		config.Timeout.ResponseHeader = defaults.Timeout.ResponseHeader
	}
	if config.Timeout.TLSHandshake == 0 {
		config.Timeout.TLSHandshake = defaults.Timeout.TLSHandshake
	}

	if config.Retry.MaxAttempts == 0 {
		config.Retry.MaxAttempts = defaults.Retry.MaxAttempts
	}
	if config.Retry.InitialDelay == 0 {
		config.Retry.InitialDelay = defaults.Retry.InitialDelay
	}
	if config.Retry.MaxDelay == 0 {
		config.Retry.MaxDelay = defaults.Retry.MaxDelay
	}
	if config.Retry.BackoffMultiplier == 0 {
		config.Retry.BackoffMultiplier = defaults.Retry.BackoffMultiplier
	}
	if config.Retry.RetryableStatusCodes == nil {
		config.Retry.RetryableStatusCodes = defaults.Retry.RetryableStatusCodes
	}

	if config.Headers.Custom == nil {
		config.Headers.Custom = make(map[string]string)
	}

	if config.Batch.Size == 0 {
		config.Batch.Size = defaults.Batch.Size
	}
	if config.Batch.ConcurrentRequests == 0 {
		config.Batch.ConcurrentRequests = defaults.Batch.ConcurrentRequests
	}

	if config.RateLimit.RequestsPerSecond == 0 {
		config.RateLimit.RequestsPerSecond = defaults.RateLimit.RequestsPerSecond
	}
	if config.RateLimit.BurstCapacity == 0 {
		config.RateLimit.BurstCapacity = defaults.RateLimit.BurstCapacity
	}

	if config.ConnectionPool.MaxIdleConnections == 0 {
		config.ConnectionPool.MaxIdleConnections = defaults.ConnectionPool.MaxIdleConnections
	}
	if config.ConnectionPool.MaxConnectionsPerHost == 0 {
		config.ConnectionPool.MaxConnectionsPerHost = defaults.ConnectionPool.MaxConnectionsPerHost
	}
	if config.ConnectionPool.MaxIdleConnectionsPerHost == 0 {
		config.ConnectionPool.MaxIdleConnectionsPerHost = defaults.ConnectionPool.MaxIdleConnectionsPerHost
	}
	if config.ConnectionPool.KeepAliveDuration == 0 {
		config.ConnectionPool.KeepAliveDuration = defaults.ConnectionPool.KeepAliveDuration
	}

	if config.UserAgent == "" {
		config.UserAgent = defaults.UserAgent
	}
}

// mergeConfigs merges source config into target config
// Only non-zero values from source are copied to target
func mergeConfigs(target, source *ClientConfig) {
	if source.BaseURL != "" {
		target.BaseURL = source.BaseURL
	}
	if source.APIToken != "" {
		target.APIToken = source.APIToken
	}
	if source.UserAgent != "" {
		target.UserAgent = source.UserAgent
	}

	target.Debug = source.Debug

	// Merge timeout config
	if source.Timeout.Connection > 0 {
		target.Timeout.Connection = source.Timeout.Connection
	}
	if source.Timeout.Request > 0 {
		target.Timeout.Request = source.Timeout.Request
	}
	if source.Timeout.IdleConnection > 0 {
		target.Timeout.IdleConnection = source.Timeout.IdleConnection
	}
	if source.Timeout.ResponseHeader > 0 {
		target.Timeout.ResponseHeader = source.Timeout.ResponseHeader
	}
	if source.Timeout.TLSHandshake > 0 {
		target.Timeout.TLSHandshake = source.Timeout.TLSHandshake
	}

	// Merge retry config
	if source.Retry.MaxAttempts > 0 {
		target.Retry.MaxAttempts = source.Retry.MaxAttempts
	}
	if source.Retry.InitialDelay > 0 {
		target.Retry.InitialDelay = source.Retry.InitialDelay
	}
	if source.Retry.MaxDelay > 0 {
		target.Retry.MaxDelay = source.Retry.MaxDelay
	}
	if source.Retry.BackoffMultiplier > 0 {
		target.Retry.BackoffMultiplier = source.Retry.BackoffMultiplier
	}

	// Merge custom headers
	for k, v := range source.Headers.Custom {
		target.Headers.Custom[k] = v
	}

	// Merge batch config
	if source.Batch.Size > 0 {
		target.Batch.Size = source.Batch.Size
	}
	if source.Batch.ConcurrentRequests > 0 {
		target.Batch.ConcurrentRequests = source.Batch.ConcurrentRequests
	}

	// Merge rate limit config
	target.RateLimit.Enabled = source.RateLimit.Enabled
	if source.RateLimit.RequestsPerSecond > 0 {
		target.RateLimit.RequestsPerSecond = source.RateLimit.RequestsPerSecond
	}
	if source.RateLimit.BurstCapacity > 0 {
		target.RateLimit.BurstCapacity = source.RateLimit.BurstCapacity
	}

	// Merge connection pool config
	if source.ConnectionPool.MaxIdleConnections > 0 {
		target.ConnectionPool.MaxIdleConnections = source.ConnectionPool.MaxIdleConnections
	}
	if source.ConnectionPool.MaxConnectionsPerHost > 0 {
		target.ConnectionPool.MaxConnectionsPerHost = source.ConnectionPool.MaxConnectionsPerHost
	}
	if source.ConnectionPool.KeepAliveDuration > 0 {
		target.ConnectionPool.KeepAliveDuration = source.ConnectionPool.KeepAliveDuration
	}
}

// PrintEnvVarHelp prints help text for all supported environment variables
func PrintEnvVarHelp() string {
	var sb strings.Builder
	sb.WriteString("Instana Go Client - Environment Variables:\n\n")
	fmt.Fprintf(&sb, "  %-40s  %s\n", EnvBaseURL, "Base URL of the Instana API")
	fmt.Fprintf(&sb, "  %-40s  %s\n", EnvAPIToken, "API token for authentication")
	fmt.Fprintf(&sb, "  %-40s  %s\n", EnvDebug, "Enable debug logging (true/false)")
	fmt.Fprintf(&sb, "  %-40s  %s\n", EnvConnectionTimeout, "Connection timeout (e.g., 30s, 1m)")
	fmt.Fprintf(&sb, "  %-40s  %s\n", EnvRequestTimeout, "Request timeout (e.g., 60s, 2m)")
	fmt.Fprintf(&sb, "  %-40s  %s\n", EnvIdleConnectionTimeout, "Idle connection timeout (e.g., 90s)")
	fmt.Fprintf(&sb, "  %-40s  %s\n", EnvMaxRetryAttempts, "Maximum retry attempts (integer)")
	fmt.Fprintf(&sb, "  %-40s  %s\n", EnvRetryInitialDelay, "Initial retry delay (e.g., 1s)")
	fmt.Fprintf(&sb, "  %-40s  %s\n", EnvRetryMaxDelay, "Maximum retry delay (e.g., 30s)")
	fmt.Fprintf(&sb, "  %-40s  %s\n", EnvRetryBackoffMultiplier, "Retry backoff multiplier (float)")
	fmt.Fprintf(&sb, "  %-40s  %s\n", EnvBatchSize, "Batch operation size (integer)")
	fmt.Fprintf(&sb, "  %-40s  %s\n", EnvBatchConcurrentRequests, "Concurrent batch requests (integer)")
	fmt.Fprintf(&sb, "  %-40s  %s\n", EnvRateLimitEnabled, "Enable rate limiting (true/false)")
	fmt.Fprintf(&sb, "  %-40s  %s\n", EnvRateLimitRequestsPerSecond, "Requests per second limit (integer)")
	fmt.Fprintf(&sb, "  %-40s  %s\n", EnvRateLimitBurstCapacity, "Rate limit burst capacity (integer)")
	fmt.Fprintf(&sb, "  %-40s  %s\n", EnvMaxIdleConnections, "Maximum idle connections (integer)")
	fmt.Fprintf(&sb, "  %-40s  %s\n", EnvMaxConnectionsPerHost, "Maximum connections per host (integer)")
	fmt.Fprintf(&sb, "  %-40s  %s\n", EnvKeepAliveDuration, "Keep-alive duration (e.g., 30s)")
	return sb.String()
}
