package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// TestLoadFromEnv_Empty tests loading with no environment variables set
func TestLoadFromEnv_Empty(t *testing.T) {
	// Clear all relevant env vars
	clearEnvVars(t)

	config, err := LoadFromEnv()
	if err != nil {
		t.Fatalf("LoadFromEnv() failed: %v", err)
	}

	// Should return default config
	defaults := DefaultClientConfig()
	if config.Timeout.Connection != defaults.Timeout.Connection {
		t.Errorf("Expected default connection timeout, got %v", config.Timeout.Connection)
	}
}

// TestLoadFromEnv_BaseURL tests loading base URL from environment
func TestLoadFromEnv_BaseURL(t *testing.T) {
	clearEnvVars(t)
	os.Setenv(EnvBaseURL, "https://test.instana.io")
	defer os.Unsetenv(EnvBaseURL)

	config, err := LoadFromEnv()
	if err != nil {
		t.Fatalf("LoadFromEnv() failed: %v", err)
	}

	if config.BaseURL != "https://test.instana.io" {
		t.Errorf("Expected BaseURL=https://test.instana.io, got %s", config.BaseURL)
	}
}

// TestLoadFromEnv_APIToken tests loading API token from environment
func TestLoadFromEnv_APIToken(t *testing.T) {
	clearEnvVars(t)
	os.Setenv(EnvAPIToken, "test-token-123")
	defer os.Unsetenv(EnvAPIToken)

	config, err := LoadFromEnv()
	if err != nil {
		t.Fatalf("LoadFromEnv() failed: %v", err)
	}

	if config.APIToken != "test-token-123" {
		t.Errorf("Expected APIToken=test-token-123, got %s", config.APIToken)
	}
}

// TestLoadFromEnv_Debug tests loading debug flag from environment
func TestLoadFromEnv_Debug(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected bool
		wantErr  bool
	}{
		{"true", "true", true, false},
		{"false", "false", false, false},
		{"1", "1", true, false},
		{"0", "0", false, false},
		{"invalid", "invalid", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearEnvVars(t)
			os.Setenv(EnvDebug, tt.value)
			defer os.Unsetenv(EnvDebug)

			config, err := LoadFromEnv()
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("LoadFromEnv() failed: %v", err)
			}

			if config.Debug != tt.expected {
				t.Errorf("Expected Debug=%v, got %v", tt.expected, config.Debug)
			}
		})
	}
}

// TestLoadFromEnv_Timeouts tests loading timeout values from environment
func TestLoadFromEnv_Timeouts(t *testing.T) {
	tests := []struct {
		name    string
		envVar  string
		value   string
		check   func(*ClientConfig) time.Duration
		wantErr bool
	}{
		{
			name:   "connection timeout with duration",
			envVar: EnvConnectionTimeout,
			value:  "45s",
			check:  func(c *ClientConfig) time.Duration { return c.Timeout.Connection },
		},
		{
			name:   "connection timeout with seconds",
			envVar: EnvConnectionTimeout,
			value:  "60",
			check:  func(c *ClientConfig) time.Duration { return c.Timeout.Connection },
		},
		{
			name:   "request timeout",
			envVar: EnvRequestTimeout,
			value:  "2m",
			check:  func(c *ClientConfig) time.Duration { return c.Timeout.Request },
		},
		{
			name:   "idle connection timeout",
			envVar: EnvIdleConnectionTimeout,
			value:  "120s",
			check:  func(c *ClientConfig) time.Duration { return c.Timeout.IdleConnection },
		},
		{
			name:    "invalid timeout",
			envVar:  EnvConnectionTimeout,
			value:   "invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearEnvVars(t)
			os.Setenv(tt.envVar, tt.value)
			defer os.Unsetenv(tt.envVar)

			config, err := LoadFromEnv()
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("LoadFromEnv() failed: %v", err)
			}

			// Parse expected duration
			var expected time.Duration
			if d, err := time.ParseDuration(tt.value); err == nil {
				expected = d
			} else if seconds, err := time.ParseDuration(tt.value + "s"); err == nil {
				expected = seconds
			}

			actual := tt.check(config)
			if actual != expected {
				t.Errorf("Expected %v, got %v", expected, actual)
			}
		})
	}
}

// TestLoadFromEnv_RetryConfig tests loading retry configuration from environment
func TestLoadFromEnv_RetryConfig(t *testing.T) {
	clearEnvVars(t)
	os.Setenv(EnvMaxRetryAttempts, "5")
	os.Setenv(EnvRetryInitialDelay, "2s")
	os.Setenv(EnvRetryMaxDelay, "60s")
	os.Setenv(EnvRetryBackoffMultiplier, "3.0")
	defer func() {
		os.Unsetenv(EnvMaxRetryAttempts)
		os.Unsetenv(EnvRetryInitialDelay)
		os.Unsetenv(EnvRetryMaxDelay)
		os.Unsetenv(EnvRetryBackoffMultiplier)
	}()

	config, err := LoadFromEnv()
	if err != nil {
		t.Fatalf("LoadFromEnv() failed: %v", err)
	}

	if config.Retry.MaxAttempts != 5 {
		t.Errorf("Expected MaxAttempts=5, got %d", config.Retry.MaxAttempts)
	}
	if config.Retry.InitialDelay != 2*time.Second {
		t.Errorf("Expected InitialDelay=2s, got %v", config.Retry.InitialDelay)
	}
	if config.Retry.MaxDelay != 60*time.Second {
		t.Errorf("Expected MaxDelay=60s, got %v", config.Retry.MaxDelay)
	}
	if config.Retry.BackoffMultiplier != 3.0 {
		t.Errorf("Expected BackoffMultiplier=3.0, got %v", config.Retry.BackoffMultiplier)
	}
}

// TestLoadFromEnv_BatchConfig tests loading batch configuration from environment
func TestLoadFromEnv_BatchConfig(t *testing.T) {
	clearEnvVars(t)
	os.Setenv(EnvBatchSize, "200")
	os.Setenv(EnvBatchConcurrentRequests, "20")
	defer func() {
		os.Unsetenv(EnvBatchSize)
		os.Unsetenv(EnvBatchConcurrentRequests)
	}()

	config, err := LoadFromEnv()
	if err != nil {
		t.Fatalf("LoadFromEnv() failed: %v", err)
	}

	if config.Batch.Size != 200 {
		t.Errorf("Expected Batch.Size=200, got %d", config.Batch.Size)
	}
	if config.Batch.ConcurrentRequests != 20 {
		t.Errorf("Expected Batch.ConcurrentRequests=20, got %d", config.Batch.ConcurrentRequests)
	}
}

// TestLoadFromEnv_RateLimitConfig tests loading rate limit configuration from environment
func TestLoadFromEnv_RateLimitConfig(t *testing.T) {
	clearEnvVars(t)
	os.Setenv(EnvRateLimitEnabled, "true")
	os.Setenv(EnvRateLimitRequestsPerSecond, "50")
	os.Setenv(EnvRateLimitBurstCapacity, "100")
	defer func() {
		os.Unsetenv(EnvRateLimitEnabled)
		os.Unsetenv(EnvRateLimitRequestsPerSecond)
		os.Unsetenv(EnvRateLimitBurstCapacity)
	}()

	config, err := LoadFromEnv()
	if err != nil {
		t.Fatalf("LoadFromEnv() failed: %v", err)
	}

	if !config.RateLimit.Enabled {
		t.Error("Expected RateLimit.Enabled=true")
	}
	if config.RateLimit.RequestsPerSecond != 50 {
		t.Errorf("Expected RateLimit.RequestsPerSecond=50, got %d", config.RateLimit.RequestsPerSecond)
	}
	if config.RateLimit.BurstCapacity != 100 {
		t.Errorf("Expected RateLimit.BurstCapacity=100, got %d", config.RateLimit.BurstCapacity)
	}
}

// TestLoadFromEnv_ConnectionPoolConfig tests loading connection pool configuration from environment
func TestLoadFromEnv_ConnectionPoolConfig(t *testing.T) {
	clearEnvVars(t)
	os.Setenv(EnvMaxIdleConnections, "200")
	os.Setenv(EnvMaxConnectionsPerHost, "50")
	os.Setenv(EnvKeepAliveDuration, "45s")
	defer func() {
		os.Unsetenv(EnvMaxIdleConnections)
		os.Unsetenv(EnvMaxConnectionsPerHost)
		os.Unsetenv(EnvKeepAliveDuration)
	}()

	config, err := LoadFromEnv()
	if err != nil {
		t.Fatalf("LoadFromEnv() failed: %v", err)
	}

	if config.ConnectionPool.MaxIdleConnections != 200 {
		t.Errorf("Expected MaxIdleConnections=200, got %d", config.ConnectionPool.MaxIdleConnections)
	}
	if config.ConnectionPool.MaxConnectionsPerHost != 50 {
		t.Errorf("Expected MaxConnectionsPerHost=50, got %d", config.ConnectionPool.MaxConnectionsPerHost)
	}
	if config.ConnectionPool.KeepAliveDuration != 45*time.Second {
		t.Errorf("Expected KeepAliveDuration=45s, got %v", config.ConnectionPool.KeepAliveDuration)
	}
}

// TestLoadFromEnv_InvalidValues tests error handling for invalid environment variable values
func TestLoadFromEnv_InvalidValues(t *testing.T) {
	tests := []struct {
		name   string
		envVar string
		value  string
	}{
		{"invalid int", EnvMaxRetryAttempts, "not-a-number"},
		{"invalid float", EnvRetryBackoffMultiplier, "not-a-float"},
		{"invalid bool", EnvRateLimitEnabled, "not-a-bool"},
		{"invalid duration", EnvConnectionTimeout, "not-a-duration"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearEnvVars(t)
			os.Setenv(tt.envVar, tt.value)
			defer os.Unsetenv(tt.envVar)

			_, err := LoadFromEnv()
			if err == nil {
				t.Error("Expected error for invalid value, got nil")
			}
		})
	}
}

// TestLoadFromJSON tests loading configuration from JSON file
func TestLoadFromJSON(t *testing.T) {
	// Create temporary JSON file
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "config.json")

	// Use nanoseconds for durations in JSON (Go's time.Duration JSON format)
	config := map[string]interface{}{
		"baseURL":  "https://json.instana.io",
		"apiToken": "json-token",
		"debug":    true,
		"timeout": map[string]interface{}{
			"connection": 45000000000, // 45 seconds in nanoseconds
			"request":    90000000000, // 90 seconds in nanoseconds
		},
		"retry": map[string]interface{}{
			"maxAttempts":       5,
			"initialDelay":      2000000000,  // 2 seconds in nanoseconds
			"maxDelay":          60000000000, // 60 seconds in nanoseconds
			"backoffMultiplier": 3.0,
		},
	}

	data, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	if err := os.WriteFile(configFile, data, 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Load from JSON
	loaded, err := LoadFromJSON(configFile)
	if err != nil {
		t.Fatalf("LoadFromJSON() failed: %v", err)
	}

	if loaded.BaseURL != "https://json.instana.io" {
		t.Errorf("Expected BaseURL from JSON, got %s", loaded.BaseURL)
	}
	if loaded.APIToken != "json-token" {
		t.Errorf("Expected APIToken from JSON, got %s", loaded.APIToken)
	}
	if !loaded.Debug {
		t.Error("Expected Debug=true from JSON")
	}
}

// TestLoadFromJSON_NonExistentFile tests error handling for missing file
func TestLoadFromJSON_NonExistentFile(t *testing.T) {
	_, err := LoadFromJSON("/nonexistent/config.json")
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}

// TestLoadFromJSON_InvalidJSON tests error handling for invalid JSON
func TestLoadFromJSON_InvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "invalid.json")

	if err := os.WriteFile(configFile, []byte("not valid json"), 0644); err != nil {
		t.Fatalf("Failed to write invalid JSON file: %v", err)
	}

	_, err := LoadFromJSON(configFile)
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

// TestLoadFromJSON_AppliesDefaults tests that defaults are applied to missing fields
func TestLoadFromJSON_AppliesDefaults(t *testing.T) {
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "minimal.json")

	// Minimal config with only required fields
	config := map[string]interface{}{
		"baseURL":  "https://minimal.instana.io",
		"apiToken": "minimal-token",
	}

	data, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	if err := os.WriteFile(configFile, data, 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	loaded, err := LoadFromJSON(configFile)
	if err != nil {
		t.Fatalf("LoadFromJSON() failed: %v", err)
	}

	// Check that defaults were applied
	defaults := DefaultClientConfig()
	if loaded.Timeout.Connection != defaults.Timeout.Connection {
		t.Errorf("Expected default connection timeout, got %v", loaded.Timeout.Connection)
	}
	if loaded.Retry.MaxAttempts != defaults.Retry.MaxAttempts {
		t.Errorf("Expected default max attempts, got %d", loaded.Retry.MaxAttempts)
	}
}

// TestLoadFromJSONWithEnvOverride tests JSON loading with environment variable override
func TestLoadFromJSONWithEnvOverride(t *testing.T) {
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "config.json")

	// Create JSON config
	config := map[string]interface{}{
		"baseURL":  "https://json.instana.io",
		"apiToken": "json-token",
		"retry": map[string]interface{}{
			"maxAttempts": 3,
		},
	}

	data, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	if err := os.WriteFile(configFile, data, 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Set environment variable to override
	clearEnvVars(t)
	os.Setenv(EnvMaxRetryAttempts, "7")
	os.Setenv(EnvAPIToken, "env-token")
	defer func() {
		os.Unsetenv(EnvMaxRetryAttempts)
		os.Unsetenv(EnvAPIToken)
	}()

	// Load with override
	loaded, err := LoadFromJSONWithEnvOverride(configFile)
	if err != nil {
		t.Fatalf("LoadFromJSONWithEnvOverride() failed: %v", err)
	}

	// Environment should override JSON
	if loaded.APIToken != "env-token" {
		t.Errorf("Expected env token to override JSON, got %s", loaded.APIToken)
	}
	if loaded.Retry.MaxAttempts != 7 {
		t.Errorf("Expected env max attempts to override JSON, got %d", loaded.Retry.MaxAttempts)
	}
	// JSON value should remain for non-overridden fields
	if loaded.BaseURL != "https://json.instana.io" {
		t.Errorf("Expected JSON BaseURL to remain, got %s", loaded.BaseURL)
	}
}

// TestParseDuration tests duration parsing
func TestParseDuration(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Duration
		wantErr  bool
	}{
		{"duration string", "30s", 30 * time.Second, false},
		{"minutes", "2m", 2 * time.Minute, false},
		{"hours", "1h", 1 * time.Hour, false},
		{"seconds as int", "45", 45 * time.Second, false},
		{"zero", "0", 0, false},
		{"invalid", "invalid", 0, true},
		{"empty", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseDuration(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("parseDuration() failed: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// TestApplyDefaults tests default value application
func TestApplyDefaults(t *testing.T) {
	config := &ClientConfig{
		BaseURL:  "https://test.instana.io",
		APIToken: "test-token",
		// Leave other fields as zero values
	}

	applyDefaults(config)

	defaults := DefaultClientConfig()

	// Check that defaults were applied
	if config.Timeout.Connection != defaults.Timeout.Connection {
		t.Errorf("Expected default connection timeout, got %v", config.Timeout.Connection)
	}
	if config.Retry.MaxAttempts != defaults.Retry.MaxAttempts {
		t.Errorf("Expected default max attempts, got %d", config.Retry.MaxAttempts)
	}
	if config.UserAgent != defaults.UserAgent {
		t.Errorf("Expected default user agent, got %s", config.UserAgent)
	}

	// Check that non-zero values were preserved
	if config.BaseURL != "https://test.instana.io" {
		t.Errorf("BaseURL should not be overridden, got %s", config.BaseURL)
	}
}

// TestMergeConfigs tests configuration merging
func TestMergeConfigs(t *testing.T) {
	target := &ClientConfig{
		BaseURL:  "https://target.instana.io",
		APIToken: "target-token",
		Timeout: TimeoutConfig{
			Connection: 30 * time.Second,
		},
		Retry: RetryConfig{
			MaxAttempts: 3,
		},
	}

	source := &ClientConfig{
		BaseURL: "https://source.instana.io",
		// APIToken not set (should not override)
		Timeout: TimeoutConfig{
			Connection: 45 * time.Second,
			Request:    90 * time.Second,
		},
		Retry: RetryConfig{
			MaxAttempts: 5,
		},
	}

	mergeConfigs(target, source)

	// Source should override target for non-zero values
	if target.BaseURL != "https://source.instana.io" {
		t.Errorf("Expected source BaseURL, got %s", target.BaseURL)
	}
	if target.Timeout.Connection != 45*time.Second {
		t.Errorf("Expected source connection timeout, got %v", target.Timeout.Connection)
	}
	if target.Timeout.Request != 90*time.Second {
		t.Errorf("Expected source request timeout, got %v", target.Timeout.Request)
	}
	if target.Retry.MaxAttempts != 5 {
		t.Errorf("Expected source max attempts, got %d", target.Retry.MaxAttempts)
	}

	// Target should keep original value when source is zero/empty
	if target.APIToken != "target-token" {
		t.Errorf("Expected target APIToken to remain, got %s", target.APIToken)
	}
}

// TestPrintEnvVarHelp tests help text generation
func TestPrintEnvVarHelp(t *testing.T) {
	help := PrintEnvVarHelp()

	if help == "" {
		t.Error("Expected non-empty help text")
	}

	// Check that all environment variables are documented
	envVars := []string{
		EnvBaseURL,
		EnvAPIToken,
		EnvDebug,
		EnvConnectionTimeout,
		EnvRequestTimeout,
		EnvMaxRetryAttempts,
		EnvRateLimitEnabled,
	}

	for _, envVar := range envVars {
		if !contains(help, envVar) {
			t.Errorf("Help text should contain %s", envVar)
		}
	}
}

// Helper functions

func clearEnvVars(t *testing.T) {
	t.Helper()
	envVars := []string{
		EnvBaseURL,
		EnvAPIToken,
		EnvDebug,
		EnvConnectionTimeout,
		EnvRequestTimeout,
		EnvIdleConnectionTimeout,
		EnvMaxRetryAttempts,
		EnvRetryInitialDelay,
		EnvRetryMaxDelay,
		EnvRetryBackoffMultiplier,
		EnvBatchSize,
		EnvBatchConcurrentRequests,
		EnvRateLimitEnabled,
		EnvRateLimitRequestsPerSecond,
		EnvRateLimitBurstCapacity,
		EnvMaxIdleConnections,
		EnvMaxConnectionsPerHost,
		EnvKeepAliveDuration,
	}

	for _, envVar := range envVars {
		os.Unsetenv(envVar)
	}
}

func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && (s == substr || len(s) >= len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || containsMiddle(s, substr)))
}

func containsMiddle(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// Benchmark tests

// BenchmarkLoadFromEnv benchmarks environment variable loading
func BenchmarkLoadFromEnv(b *testing.B) {
	os.Setenv(EnvBaseURL, "https://bench.instana.io")
	os.Setenv(EnvAPIToken, "bench-token")
	defer func() {
		os.Unsetenv(EnvBaseURL)
		os.Unsetenv(EnvAPIToken)
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = LoadFromEnv()
	}
}

// BenchmarkParseDuration benchmarks duration parsing
func BenchmarkParseDuration(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = parseDuration("30s")
	}
}

// Made with Bob
