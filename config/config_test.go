package config

import (
	"testing"
	"time"
)

func TestDefaultClientConfig(t *testing.T) {
	config := DefaultClientConfig()

	// Test Timeout defaults
	if config.Timeout.Connection != 30*time.Second {
		t.Errorf("Expected Connection timeout 30s, got %v", config.Timeout.Connection)
	}
	if config.Timeout.Request != 60*time.Second {
		t.Errorf("Expected Request timeout 60s, got %v", config.Timeout.Request)
	}
	if config.Timeout.IdleConnection != 90*time.Second {
		t.Errorf("Expected IdleConnection timeout 90s, got %v", config.Timeout.IdleConnection)
	}
	if config.Timeout.ResponseHeader != 10*time.Second {
		t.Errorf("Expected ResponseHeader timeout 10s, got %v", config.Timeout.ResponseHeader)
	}
	if config.Timeout.TLSHandshake != 10*time.Second {
		t.Errorf("Expected TLSHandshake timeout 10s, got %v", config.Timeout.TLSHandshake)
	}

	// Test Retry defaults
	if config.Retry.MaxAttempts != 3 {
		t.Errorf("Expected MaxAttempts 3, got %d", config.Retry.MaxAttempts)
	}
	if config.Retry.InitialDelay != 1*time.Second {
		t.Errorf("Expected InitialDelay 1s, got %v", config.Retry.InitialDelay)
	}
	if config.Retry.MaxDelay != 30*time.Second {
		t.Errorf("Expected MaxDelay 30s, got %v", config.Retry.MaxDelay)
	}
	if config.Retry.BackoffMultiplier != 2.0 {
		t.Errorf("Expected BackoffMultiplier 2.0, got %f", config.Retry.BackoffMultiplier)
	}
	if !config.Retry.RetryOnTimeout {
		t.Error("Expected RetryOnTimeout to be true")
	}
	if !config.Retry.RetryOnConnectionError {
		t.Error("Expected RetryOnConnectionError to be true")
	}
	if !config.Retry.Jitter {
		t.Error("Expected Jitter to be true")
	}

	// Test RetryableStatusCodes
	expectedCodes := []int{408, 429, 500, 502, 503, 504}
	if len(config.Retry.RetryableStatusCodes) != len(expectedCodes) {
		t.Errorf("Expected %d retryable status codes, got %d", len(expectedCodes), len(config.Retry.RetryableStatusCodes))
	}
	for i, code := range expectedCodes {
		if config.Retry.RetryableStatusCodes[i] != code {
			t.Errorf("Expected status code %d at index %d, got %d", code, i, config.Retry.RetryableStatusCodes[i])
		}
	}

	// Test RateLimit defaults
	if config.RateLimit.RequestsPerSecond != 100 {
		t.Errorf("Expected RequestsPerSecond 100, got %d", config.RateLimit.RequestsPerSecond)
	}
	if config.RateLimit.BurstCapacity != 200 {
		t.Errorf("Expected BurstCapacity 200, got %d", config.RateLimit.BurstCapacity)
	}
	if !config.RateLimit.Enabled {
		t.Error("Expected RateLimit.Enabled to be true")
	}
	if !config.RateLimit.WaitForToken {
		t.Error("Expected WaitForToken to be true")
	}

	// Test ConnectionPool defaults
	if config.ConnectionPool.MaxIdleConnections != 100 {
		t.Errorf("Expected MaxIdleConnections 100, got %d", config.ConnectionPool.MaxIdleConnections)
	}
	if config.ConnectionPool.MaxConnectionsPerHost != 10 {
		t.Errorf("Expected MaxConnectionsPerHost 10, got %d", config.ConnectionPool.MaxConnectionsPerHost)
	}
	if config.ConnectionPool.MaxIdleConnectionsPerHost != 10 {
		t.Errorf("Expected MaxIdleConnectionsPerHost 10, got %d", config.ConnectionPool.MaxIdleConnectionsPerHost)
	}
	if config.ConnectionPool.KeepAliveDuration != 30*time.Second {
		t.Errorf("Expected KeepAliveDuration 30s, got %v", config.ConnectionPool.KeepAliveDuration)
	}
	if config.ConnectionPool.DisableKeepAlives {
		t.Error("Expected DisableKeepAlives to be false")
	}
	if config.ConnectionPool.DisableCompression {
		t.Error("Expected DisableCompression to be false")
	}

	// Test Headers defaults
	if config.Headers.Custom == nil {
		t.Error("Expected Custom headers map to be initialized")
	}
	if len(config.Headers.Custom) != 0 {
		t.Errorf("Expected empty Custom headers map, got %d entries", len(config.Headers.Custom))
	}
	if config.Headers.DisableDefaultHeaders {
		t.Error("Expected DisableDefaultHeaders to be false")
	}

	// Test Batch defaults
	if config.Batch.Size != 100 {
		t.Errorf("Expected Batch.Size 100, got %d", config.Batch.Size)
	}
	if config.Batch.ConcurrentRequests != 5 {
		t.Errorf("Expected ConcurrentRequests 5, got %d", config.Batch.ConcurrentRequests)
	}
	if config.Batch.StopOnError {
		t.Error("Expected StopOnError to be false")
	}
	if !config.Batch.RetryFailedItems {
		t.Error("Expected RetryFailedItems to be true")
	}

	// Test other defaults
	if config.UserAgent != "instana-go-client/v1.0.0" {
		t.Errorf("Expected UserAgent 'instana-go-client/v1.0.0', got '%s'", config.UserAgent)
	}
	if config.Debug {
		t.Error("Expected Debug to be false")
	}
	if config.Logger != nil {
		t.Error("Expected Logger to be nil")
	}
	if config.HTTPClient != nil {
		t.Error("Expected HTTPClient to be nil")
	}
}

func TestConfigClone(t *testing.T) {
	original := DefaultClientConfig()
	original.BaseURL = "https://original.instana.io"
	original.APIToken = "original-token"
	original.Retry.MaxAttempts = 5
	original.Headers.Custom["X-Test"] = "test-value"

	clone := original.Clone()

	// Verify clone has same values
	if clone.BaseURL != original.BaseURL {
		t.Errorf("Expected BaseURL '%s', got '%s'", original.BaseURL, clone.BaseURL)
	}
	if clone.APIToken != original.APIToken {
		t.Errorf("Expected APIToken '%s', got '%s'", original.APIToken, clone.APIToken)
	}
	if clone.Retry.MaxAttempts != original.Retry.MaxAttempts {
		t.Errorf("Expected MaxAttempts %d, got %d", original.Retry.MaxAttempts, clone.Retry.MaxAttempts)
	}

	// Verify deep copy - modifying clone shouldn't affect original
	clone.BaseURL = "https://clone.instana.io"
	clone.APIToken = "clone-token"
	clone.Retry.MaxAttempts = 10
	clone.Headers.Custom["X-Test"] = "modified"
	clone.Headers.Custom["X-New"] = "new-value"

	if original.BaseURL == clone.BaseURL {
		t.Error("Modifying clone BaseURL affected original")
	}
	if original.APIToken == clone.APIToken {
		t.Error("Modifying clone APIToken affected original")
	}
	if original.Retry.MaxAttempts == clone.Retry.MaxAttempts {
		t.Error("Modifying clone MaxAttempts affected original")
	}
	if original.Headers.Custom["X-Test"] == clone.Headers.Custom["X-Test"] {
		t.Error("Modifying clone Custom headers affected original")
	}
	if _, exists := original.Headers.Custom["X-New"]; exists {
		t.Error("Adding to clone Custom headers affected original")
	}

	// Verify RetryableStatusCodes deep copy
	clone.Retry.RetryableStatusCodes[0] = 999
	if original.Retry.RetryableStatusCodes[0] == 999 {
		t.Error("Modifying clone RetryableStatusCodes affected original")
	}
}

func TestConfigCloneNil(t *testing.T) {
	var config *ClientConfig
	clone := config.Clone()
	if clone != nil {
		t.Error("Expected nil clone from nil config")
	}
}

func TestTimeoutConfig(t *testing.T) {
	config := DefaultClientConfig()

	// Test that all timeouts are positive
	if config.Timeout.Connection <= 0 {
		t.Error("Connection timeout must be positive")
	}
	if config.Timeout.Request <= 0 {
		t.Error("Request timeout must be positive")
	}
	if config.Timeout.IdleConnection <= 0 {
		t.Error("IdleConnection timeout must be positive")
	}
	if config.Timeout.ResponseHeader <= 0 {
		t.Error("ResponseHeader timeout must be positive")
	}
	if config.Timeout.TLSHandshake <= 0 {
		t.Error("TLSHandshake timeout must be positive")
	}

	// Test logical relationships
	if config.Timeout.Request < config.Timeout.Connection {
		t.Error("Request timeout should be >= Connection timeout")
	}
}

func TestRetryConfig(t *testing.T) {
	config := DefaultClientConfig()

	// Test retry configuration is sensible
	if config.Retry.MaxAttempts < 0 {
		t.Error("MaxAttempts must be non-negative")
	}
	if config.Retry.InitialDelay <= 0 {
		t.Error("InitialDelay must be positive")
	}
	if config.Retry.MaxDelay <= 0 {
		t.Error("MaxDelay must be positive")
	}
	if config.Retry.MaxDelay < config.Retry.InitialDelay {
		t.Error("MaxDelay should be >= InitialDelay")
	}
	if config.Retry.BackoffMultiplier < 1.0 {
		t.Error("BackoffMultiplier should be >= 1.0")
	}
}

func TestRateLimitConfig(t *testing.T) {
	config := DefaultClientConfig()

	// Test rate limit configuration
	if config.RateLimit.RequestsPerSecond <= 0 {
		t.Error("RequestsPerSecond must be positive")
	}
	if config.RateLimit.BurstCapacity <= 0 {
		t.Error("BurstCapacity must be positive")
	}
	if config.RateLimit.BurstCapacity < config.RateLimit.RequestsPerSecond {
		t.Error("BurstCapacity should be >= RequestsPerSecond")
	}
}

func TestConnectionPoolConfig(t *testing.T) {
	config := DefaultClientConfig()

	// Test connection pool configuration
	if config.ConnectionPool.MaxIdleConnections <= 0 {
		t.Error("MaxIdleConnections must be positive")
	}
	if config.ConnectionPool.MaxConnectionsPerHost <= 0 {
		t.Error("MaxConnectionsPerHost must be positive")
	}
	if config.ConnectionPool.MaxIdleConnectionsPerHost <= 0 {
		t.Error("MaxIdleConnectionsPerHost must be positive")
	}
	if config.ConnectionPool.KeepAliveDuration <= 0 {
		t.Error("KeepAliveDuration must be positive")
	}

	// Test logical relationships
	if config.ConnectionPool.MaxIdleConnectionsPerHost > config.ConnectionPool.MaxConnectionsPerHost {
		t.Error("MaxIdleConnectionsPerHost should be <= MaxConnectionsPerHost")
	}
}

func TestBatchConfig(t *testing.T) {
	config := DefaultClientConfig()

	// Test batch configuration
	if config.Batch.Size <= 0 {
		t.Error("Batch.Size must be positive")
	}
	if config.Batch.ConcurrentRequests <= 0 {
		t.Error("ConcurrentRequests must be positive")
	}
}

func TestHeadersConfig(t *testing.T) {
	config := DefaultClientConfig()

	// Test headers configuration
	if config.Headers.Custom == nil {
		t.Error("Custom headers map should be initialized")
	}

	// Test that we can add custom headers
	config.Headers.Custom["X-Test"] = "test-value"
	if config.Headers.Custom["X-Test"] != "test-value" {
		t.Error("Failed to add custom header")
	}
}

func TestConfigImmutability(t *testing.T) {
	// Test that DefaultClientConfig returns a new instance each time
	config1 := DefaultClientConfig()
	config2 := DefaultClientConfig()

	// Modify config1
	config1.BaseURL = "https://test1.instana.io"
	config1.Retry.MaxAttempts = 10

	// Verify config2 is not affected
	if config2.BaseURL == config1.BaseURL {
		t.Error("DefaultClientConfig should return independent instances")
	}
	if config2.Retry.MaxAttempts == config1.Retry.MaxAttempts {
		t.Error("DefaultClientConfig should return independent instances")
	}
}

func BenchmarkDefaultClientConfig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = DefaultClientConfig()
	}
}

func BenchmarkConfigClone(b *testing.B) {
	config := DefaultClientConfig()
	config.Headers.Custom["X-Test"] = "test"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = config.Clone()
	}
}
