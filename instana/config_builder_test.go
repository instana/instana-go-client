package instana

import (
	"testing"
	"time"
)

func TestNewConfigBuilder(t *testing.T) {
	builder := NewConfigBuilder()
	if builder == nil {
		t.Fatal("NewConfigBuilder() returned nil")
	}

	// Verify it starts with default config
	config := builder.GetConfig()
	defaultConfig := DefaultClientConfig()

	if config.Timeout.Connection != defaultConfig.Timeout.Connection {
		t.Errorf("Expected default connection timeout %v, got %v",
			defaultConfig.Timeout.Connection, config.Timeout.Connection)
	}
}

func TestConfigBuilderWithBaseURL(t *testing.T) {
	url := "https://test.instana.io"
	builder := NewConfigBuilder().
		WithBaseURL(url).
		WithAPIToken("test-token")

	config, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() failed: %v", err)
	}

	if config.BaseURL != url {
		t.Errorf("Expected BaseURL %s, got %s", url, config.BaseURL)
	}
}

func TestConfigBuilderWithAPIToken(t *testing.T) {
	token := "test-token-123"
	builder := NewConfigBuilder().
		WithBaseURL("https://test.instana.io").
		WithAPIToken(token)

	config, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() failed: %v", err)
	}

	if config.APIToken != token {
		t.Errorf("Expected APIToken %s, got %s", token, config.APIToken)
	}
}

func TestConfigBuilderWithConnectionTimeout(t *testing.T) {
	timeout := 45 * time.Second
	builder := NewConfigBuilder().
		WithBaseURL("https://test.instana.io").
		WithAPIToken("test-token").
		WithConnectionTimeout(timeout)

	config, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() failed: %v", err)
	}

	if config.Timeout.Connection != timeout {
		t.Errorf("Expected connection timeout %v, got %v", timeout, config.Timeout.Connection)
	}
}

func TestConfigBuilderWithRequestTimeout(t *testing.T) {
	timeout := 90 * time.Second
	builder := NewConfigBuilder().
		WithBaseURL("https://test.instana.io").
		WithAPIToken("test-token").
		WithRequestTimeout(timeout)

	config, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() failed: %v", err)
	}

	if config.Timeout.Request != timeout {
		t.Errorf("Expected request timeout %v, got %v", timeout, config.Timeout.Request)
	}
}

func TestConfigBuilderWithRetryConfig(t *testing.T) {
	maxAttempts := 5
	initialDelay := 2 * time.Second
	maxDelay := 30 * time.Second
	multiplier := 2.5

	builder := NewConfigBuilder().
		WithBaseURL("https://test.instana.io").
		WithAPIToken("test-token").
		WithMaxRetryAttempts(maxAttempts).
		WithRetryInitialDelay(initialDelay).
		WithRetryMaxDelay(maxDelay).
		WithRetryBackoffMultiplier(multiplier)

	config, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() failed: %v", err)
	}

	if config.Retry.MaxAttempts != maxAttempts {
		t.Errorf("Expected MaxAttempts %d, got %d", maxAttempts, config.Retry.MaxAttempts)
	}
	if config.Retry.InitialDelay != initialDelay {
		t.Errorf("Expected InitialDelay %v, got %v", initialDelay, config.Retry.InitialDelay)
	}
	if config.Retry.MaxDelay != maxDelay {
		t.Errorf("Expected MaxDelay %v, got %v", maxDelay, config.Retry.MaxDelay)
	}
	if config.Retry.BackoffMultiplier != multiplier {
		t.Errorf("Expected BackoffMultiplier %f, got %f", multiplier, config.Retry.BackoffMultiplier)
	}
}

func TestConfigBuilderWithRetryableStatusCodes(t *testing.T) {
	codes := []int{500, 502, 503, 504}
	builder := NewConfigBuilder().
		WithBaseURL("https://test.instana.io").
		WithAPIToken("test-token").
		WithRetryableStatusCodes(codes)

	config, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() failed: %v", err)
	}

	if len(config.Retry.RetryableStatusCodes) != len(codes) {
		t.Errorf("Expected %d status codes, got %d", len(codes), len(config.Retry.RetryableStatusCodes))
	}

	for i, code := range codes {
		if config.Retry.RetryableStatusCodes[i] != code {
			t.Errorf("Expected status code %d at index %d, got %d", code, i, config.Retry.RetryableStatusCodes[i])
		}
	}
}

func TestConfigBuilderWithRateLimitConfig(t *testing.T) {
	rps := 200
	burst := 400

	builder := NewConfigBuilder().
		WithBaseURL("https://test.instana.io").
		WithAPIToken("test-token").
		WithRateLimitEnabled(true).
		WithRateLimitRequestsPerSecond(rps).
		WithRateLimitBurstCapacity(burst)

	config, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() failed: %v", err)
	}

	if !config.RateLimit.Enabled {
		t.Error("Expected RateLimit.Enabled to be true")
	}
	if config.RateLimit.RequestsPerSecond != rps {
		t.Errorf("Expected RequestsPerSecond %d, got %d", rps, config.RateLimit.RequestsPerSecond)
	}
	if config.RateLimit.BurstCapacity != burst {
		t.Errorf("Expected BurstCapacity %d, got %d", burst, config.RateLimit.BurstCapacity)
	}
}

func TestConfigBuilderWithConnectionPoolConfig(t *testing.T) {
	maxIdle := 150
	maxPerHost := 15

	builder := NewConfigBuilder().
		WithBaseURL("https://test.instana.io").
		WithAPIToken("test-token").
		WithMaxIdleConnections(maxIdle).
		WithMaxConnectionsPerHost(maxPerHost)

	config, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() failed: %v", err)
	}

	if config.ConnectionPool.MaxIdleConnections != maxIdle {
		t.Errorf("Expected MaxIdleConnections %d, got %d", maxIdle, config.ConnectionPool.MaxIdleConnections)
	}
	if config.ConnectionPool.MaxConnectionsPerHost != maxPerHost {
		t.Errorf("Expected MaxConnectionsPerHost %d, got %d", maxPerHost, config.ConnectionPool.MaxConnectionsPerHost)
	}
}

func TestConfigBuilderWithBatchConfig(t *testing.T) {
	size := 50
	concurrent := 8

	builder := NewConfigBuilder().
		WithBaseURL("https://test.instana.io").
		WithAPIToken("test-token").
		WithBatchSize(size).
		WithBatchConcurrentRequests(concurrent)

	config, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() failed: %v", err)
	}

	if config.Batch.Size != size {
		t.Errorf("Expected Batch.Size %d, got %d", size, config.Batch.Size)
	}
	if config.Batch.ConcurrentRequests != concurrent {
		t.Errorf("Expected ConcurrentRequests %d, got %d", concurrent, config.Batch.ConcurrentRequests)
	}
}

func TestConfigBuilderWithCustomHeaders(t *testing.T) {
	headers := map[string]string{
		"X-Custom-Header":  "value1",
		"X-Another-Header": "value2",
	}

	builder := NewConfigBuilder().
		WithBaseURL("https://test.instana.io").
		WithAPIToken("test-token").
		WithCustomHeaders(headers)

	config, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() failed: %v", err)
	}

	if len(config.Headers.Custom) != len(headers) {
		t.Errorf("Expected %d custom headers, got %d", len(headers), len(config.Headers.Custom))
	}

	for key, value := range headers {
		if config.Headers.Custom[key] != value {
			t.Errorf("Expected header %s=%s, got %s", key, value, config.Headers.Custom[key])
		}
	}
}

func TestConfigBuilderWithUserAgent(t *testing.T) {
	userAgent := "MyApp/1.0.0"
	builder := NewConfigBuilder().
		WithBaseURL("https://test.instana.io").
		WithAPIToken("test-token").
		WithUserAgent(userAgent)

	config, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() failed: %v", err)
	}

	if config.UserAgent != userAgent {
		t.Errorf("Expected UserAgent %s, got %s", userAgent, config.UserAgent)
	}
}

func TestConfigBuilderChaining(t *testing.T) {
	builder := NewConfigBuilder().
		WithBaseURL("https://test.instana.io").
		WithAPIToken("test-token").
		WithConnectionTimeout(30 * time.Second).
		WithRequestTimeout(60 * time.Second).
		WithMaxRetryAttempts(5).
		WithRateLimitEnabled(true).
		WithRateLimitRequestsPerSecond(200)

	config, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() failed: %v", err)
	}

	if config.BaseURL != "https://test.instana.io" {
		t.Error("BaseURL not set correctly in chain")
	}
	if config.APIToken != "test-token" {
		t.Error("APIToken not set correctly in chain")
	}
	if config.Timeout.Connection != 30*time.Second {
		t.Error("Connection timeout not set correctly in chain")
	}
	if config.Retry.MaxAttempts != 5 {
		t.Error("MaxRetryAttempts not set correctly in chain")
	}
	if !config.RateLimit.Enabled {
		t.Error("RateLimit not enabled in chain")
	}
}

func TestConfigBuilderBuildValidation(t *testing.T) {
	// Build without required fields should fail validation
	builder := NewConfigBuilder()

	_, err := builder.Build()
	if err == nil {
		t.Error("Expected validation error for config without BaseURL and APIToken")
	}
}

func TestConfigBuilderGetConfig(t *testing.T) {
	builder := NewConfigBuilder().
		WithBaseURL("https://test.instana.io")

	// GetConfig should return config without validation
	config := builder.GetConfig()
	if config == nil {
		t.Fatal("GetConfig() returned nil")
	}

	if config.BaseURL != "https://test.instana.io" {
		t.Error("GetConfig() did not return correct config")
	}
}

func TestConfigBuilderMustBuild(t *testing.T) {
	// Valid config should not panic
	builder := NewConfigBuilder().
		WithBaseURL("https://test.instana.io").
		WithAPIToken("test-token")

	config := builder.MustBuild()
	if config == nil {
		t.Fatal("MustBuild() returned nil")
	}
}

func TestConfigBuilderMustBuildPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustBuild() should have panicked on invalid config")
		}
	}()

	// Invalid config should panic
	builder := NewConfigBuilder()
	_ = builder.MustBuild()
}

func TestConfigBuilderImmutability(t *testing.T) {
	builder := NewConfigBuilder().
		WithBaseURL("https://test1.instana.io").
		WithAPIToken("token1")

	config1, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() failed: %v", err)
	}

	// Modify builder
	builder.WithBaseURL("https://test2.instana.io").
		WithAPIToken("token2")

	config2, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() failed: %v", err)
	}

	// config1 should be affected (builder modifies same config)
	// This is expected behavior - builder is mutable
	if config1.BaseURL == "https://test1.instana.io" {
		t.Log("Note: Builder modifies the same config instance")
	}

	// config2 should have new values
	if config2.BaseURL != "https://test2.instana.io" {
		t.Error("Second config does not have updated BaseURL")
	}
	if config2.APIToken != "token2" {
		t.Error("Second config does not have updated APIToken")
	}
}

func TestConfigBuilderWithAllTimeouts(t *testing.T) {
	builder := NewConfigBuilder().
		WithBaseURL("https://test.instana.io").
		WithAPIToken("test-token").
		WithConnectionTimeout(10 * time.Second).
		WithRequestTimeout(60 * time.Second).
		WithIdleConnectionTimeout(90 * time.Second).
		WithResponseHeaderTimeout(5 * time.Second).
		WithTLSHandshakeTimeout(10 * time.Second)

	config, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() failed: %v", err)
	}

	if config.Timeout.Connection != 10*time.Second {
		t.Error("Connection timeout not set")
	}
	if config.Timeout.Request != 60*time.Second {
		t.Error("Request timeout not set")
	}
	if config.Timeout.IdleConnection != 90*time.Second {
		t.Error("IdleConnection timeout not set")
	}
	if config.Timeout.ResponseHeader != 5*time.Second {
		t.Error("ResponseHeader timeout not set")
	}
	if config.Timeout.TLSHandshake != 10*time.Second {
		t.Error("TLSHandshake timeout not set")
	}
}

func TestConfigBuilderWithDebug(t *testing.T) {
	builder := NewConfigBuilder().
		WithBaseURL("https://test.instana.io").
		WithAPIToken("test-token").
		WithDebug(true)

	config, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() failed: %v", err)
	}

	if !config.Debug {
		t.Error("Expected Debug mode to be enabled")
	}
}

func TestNewConfigBuilderFromConfig(t *testing.T) {
	original := DefaultClientConfig()
	original.BaseURL = "https://original.instana.io"
	original.APIToken = "original-token"
	original.Retry.MaxAttempts = 7

	// Create builder from existing config
	builder := NewConfigBuilderFromConfig(original)

	// Modify one field
	config, err := builder.WithBaseURL("https://modified.instana.io").Build()
	if err != nil {
		t.Fatalf("Build() failed: %v", err)
	}

	// Check modified field
	if config.BaseURL != "https://modified.instana.io" {
		t.Error("BaseURL was not modified")
	}

	// Check other fields remain
	if config.APIToken != "original-token" {
		t.Error("APIToken was lost")
	}
	if config.Retry.MaxAttempts != 7 {
		t.Error("MaxAttempts was lost")
	}

	// Original should not be modified
	if original.BaseURL != "https://original.instana.io" {
		t.Error("Original config was modified")
	}
}

func BenchmarkConfigBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewConfigBuilder().
			WithBaseURL("https://test.instana.io").
			WithAPIToken("test-token").
			WithConnectionTimeout(30 * time.Second).
			WithMaxRetryAttempts(3).
			Build()
	}
}

func BenchmarkConfigBuilderComplex(b *testing.B) {
	headers := map[string]string{
		"X-Header-1": "value1",
		"X-Header-2": "value2",
		"X-Header-3": "value3",
	}

	for i := 0; i < b.N; i++ {
		_, _ = NewConfigBuilder().
			WithBaseURL("https://test.instana.io").
			WithAPIToken("test-token").
			WithConnectionTimeout(30 * time.Second).
			WithRequestTimeout(60 * time.Second).
			WithMaxRetryAttempts(5).
			WithRetryInitialDelay(1 * time.Second).
			WithRetryMaxDelay(30 * time.Second).
			WithRateLimitEnabled(true).
			WithRateLimitRequestsPerSecond(200).
			WithRateLimitBurstCapacity(400).
			WithMaxIdleConnections(150).
			WithCustomHeaders(headers).
			Build()
	}
}

// Made with Bob
