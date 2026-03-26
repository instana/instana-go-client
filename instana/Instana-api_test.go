package instana

import (
	"testing"

	"github.com/instana/instana-go-client/config"
)

// TestNewInstanaAPI tests the basic API client creation
func TestNewInstanaAPI(t *testing.T) {
	api := NewInstanaAPI("test-token", "test.instana.io", false)

	if api == nil {
		t.Fatal("Expected API client to be created, got nil")
	}

	// Verify API methods are accessible
	if api.APITokens() == nil {
		t.Error("Expected APITokens() to return non-nil")
	}

	if api.Teams() == nil {
		t.Error("Expected Teams() to return non-nil")
	}
}

// TestNewInstanaAPI_WithTLSSkip tests API client creation with TLS verification skipped
func TestNewInstanaAPI_WithTLSSkip(t *testing.T) {
	api := NewInstanaAPI("test-token", "test.instana.io", true)

	if api == nil {
		t.Fatal("Expected API client to be created, got nil")
	}

	// Verify API is functional
	if api.CustomEventSpecifications() == nil {
		t.Error("Expected CustomEventSpecifications() to return non-nil")
	}
}

// TestNewInstanaAPIWithUserAgent tests API client creation with custom user agent
func TestNewInstanaAPIWithUserAgent(t *testing.T) {
	api := NewInstanaAPIWithUserAgent(
		"test-token",
		"test.instana.io",
		false,
		"TestApp/1.0.0",
	)

	if api == nil {
		t.Fatal("Expected API client to be created, got nil")
	}

	// Verify all API methods are accessible
	if api.ApplicationConfigs() == nil {
		t.Error("Expected ApplicationConfigs() to return non-nil")
	}

	if api.AlertingChannels() == nil {
		t.Error("Expected AlertingChannels() to return non-nil")
	}
}

// TestNewInstanaAPIWithUserAgent_WithTLSSkip tests with TLS skip and custom user agent
func TestNewInstanaAPIWithUserAgent_WithTLSSkip(t *testing.T) {
	api := NewInstanaAPIWithUserAgent(
		"test-token",
		"test.instana.io",
		true,
		"TestApp/2.0.0",
	)

	if api == nil {
		t.Fatal("Expected API client to be created, got nil")
	}

	// Verify API methods work
	if api.SliConfigs() == nil {
		t.Error("Expected SliConfigs() to return non-nil")
	}
}

// TestNewInstanaAPIWithConfig tests API client creation with full config
func TestNewInstanaAPIWithConfig(t *testing.T) {
	cfg := config.DefaultClientConfig()
	cfg.APIToken = "test-token"
	cfg.BaseURL = "https://test.instana.io"

	api, err := NewInstanaAPIWithConfig(cfg)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if api == nil {
		t.Fatal("Expected API client to be created, got nil")
	}

	// Verify API methods are accessible
	if api.MaintenanceWindowConfigs() == nil {
		t.Error("Expected MaintenanceWindowConfigs() to return non-nil")
	}

	if api.CustomDashboards() == nil {
		t.Error("Expected CustomDashboards() to return non-nil")
	}
}

// TestNewInstanaAPIWithConfig_WithRetryConfig tests with retry configuration
func TestNewInstanaAPIWithConfig_WithRetryConfig(t *testing.T) {
	cfg := config.DefaultClientConfig()
	cfg.APIToken = "test-token"
	cfg.BaseURL = "https://test.instana.io"
	cfg.Retry.MaxAttempts = 5
	cfg.Retry.InitialDelay = 100

	api, err := NewInstanaAPIWithConfig(cfg)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if api == nil {
		t.Fatal("Expected API client to be created, got nil")
	}
}

// TestNewInstanaAPIWithConfig_WithRateLimitConfig tests with rate limit configuration
func TestNewInstanaAPIWithConfig_WithRateLimitConfig(t *testing.T) {
	cfg := config.DefaultClientConfig()
	cfg.APIToken = "test-token"
	cfg.BaseURL = "https://test.instana.io"
	cfg.RateLimit.Enabled = true
	cfg.RateLimit.RequestsPerSecond = 10

	api, err := NewInstanaAPIWithConfig(cfg)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if api == nil {
		t.Fatal("Expected API client to be created, got nil")
	}
}

// TestNewInstanaAPIWithConfig_InvalidConfig tests with invalid configuration
func TestNewInstanaAPIWithConfig_InvalidConfig(t *testing.T) {
	cfg := config.DefaultClientConfig()
	// Missing required fields
	cfg.APIToken = ""
	cfg.BaseURL = ""

	_, err := NewInstanaAPIWithConfig(cfg)

	if err == nil {
		t.Error("Expected error for invalid config, got nil")
	}
}

// TestNewInstanaAPIWithConfig_EmptyAPIToken tests with empty API token
func TestNewInstanaAPIWithConfig_EmptyAPIToken(t *testing.T) {
	cfg := config.DefaultClientConfig()
	cfg.APIToken = ""
	cfg.BaseURL = "https://test.instana.io"

	_, err := NewInstanaAPIWithConfig(cfg)

	if err == nil {
		t.Error("Expected error for empty API token, got nil")
	}
}

// TestNewInstanaAPIWithConfig_EmptyBaseURL tests with empty base URL
func TestNewInstanaAPIWithConfig_EmptyBaseURL(t *testing.T) {
	cfg := config.DefaultClientConfig()
	cfg.APIToken = "test-token"
	cfg.BaseURL = ""

	_, err := NewInstanaAPIWithConfig(cfg)

	if err == nil {
		t.Error("Expected error for empty base URL, got nil")
	}
}

// TestCreateHTTPClient tests the HTTP client creation
func TestCreateHTTPClient(t *testing.T) {
	cfg := config.DefaultClientConfig()
	cfg.ConnectionPool.MaxIdleConnections = 100
	cfg.ConnectionPool.MaxConnectionsPerHost = 10

	httpClient := createHTTPClient(cfg)

	if httpClient == nil {
		t.Fatal("Expected HTTP client to be created, got nil")
	}

	if httpClient.Timeout != cfg.Timeout.Request {
		t.Errorf("Expected timeout %v, got %v", cfg.Timeout.Request, httpClient.Timeout)
	}

	// Verify transport is configured
	if httpClient.Transport == nil {
		t.Error("Expected transport to be configured")
	}
}

// TestCreateHTTPClient_WithCustomTimeout tests HTTP client with custom timeout
func TestCreateHTTPClient_WithCustomTimeout(t *testing.T) {
	cfg := config.DefaultClientConfig()
	cfg.Timeout.Request = 30000 // 30 seconds

	httpClient := createHTTPClient(cfg)

	if httpClient == nil {
		t.Fatal("Expected HTTP client to be created, got nil")
	}

	if httpClient.Timeout != cfg.Timeout.Request {
		t.Errorf("Expected timeout %v, got %v", cfg.Timeout.Request, httpClient.Timeout)
	}
}

// TestNewInstanaAPI_AllAPIMethods tests that all API methods are accessible
func TestNewInstanaAPI_AllAPIMethods(t *testing.T) {
	api := NewInstanaAPI("test-token", "test.instana.io", false)

	tests := []struct {
		name   string
		method func() interface{}
	}{
		{"CustomEventSpecifications", func() interface{} { return api.CustomEventSpecifications() }},
		{"BuiltinEventSpecifications", func() interface{} { return api.BuiltinEventSpecifications() }},
		{"APITokens", func() interface{} { return api.APITokens() }},
		{"ApplicationConfigs", func() interface{} { return api.ApplicationConfigs() }},
		{"ApplicationAlertConfigs", func() interface{} { return api.ApplicationAlertConfigs() }},
		{"GlobalApplicationAlertConfigs", func() interface{} { return api.GlobalApplicationAlertConfigs() }},
		{"AlertingChannels", func() interface{} { return api.AlertingChannels() }},
		{"AlertingConfigurations", func() interface{} { return api.AlertingConfigurations() }},
		{"SliConfigs", func() interface{} { return api.SliConfigs() }},
		{"SloConfigs", func() interface{} { return api.SloConfigs() }},
		{"SloAlertConfigs", func() interface{} { return api.SloAlertConfigs() }},
		{"SloCorrectionConfigs", func() interface{} { return api.SloCorrectionConfigs() }},
		{"WebsiteMonitoringConfigs", func() interface{} { return api.WebsiteMonitoringConfigs() }},
		{"WebsiteAlertConfigs", func() interface{} { return api.WebsiteAlertConfigs() }},
		{"InfraAlertConfigs", func() interface{} { return api.InfraAlertConfigs() }},
		{"MobileAlertConfigs", func() interface{} { return api.MobileAlertConfigs() }},
		{"MaintenanceWindowConfigs", func() interface{} { return api.MaintenanceWindowConfigs() }},
		{"Teams", func() interface{} { return api.Teams() }},
		{"Groups", func() interface{} { return api.Groups() }},
		{"Roles", func() interface{} { return api.Roles() }},
		{"CustomDashboards", func() interface{} { return api.CustomDashboards() }},
		{"SyntheticTests", func() interface{} { return api.SyntheticTests() }},
		{"SyntheticLocations", func() interface{} { return api.SyntheticLocations() }},
		{"SyntheticAlertConfigs", func() interface{} { return api.SyntheticAlertConfigs() }},
		{"AutomationActions", func() interface{} { return api.AutomationActions() }},
		{"AutomationPolicies", func() interface{} { return api.AutomationPolicies() }},
		{"HostAgents", func() interface{} { return api.HostAgents() }},
		{"Users", func() interface{} { return api.Users() }},
		{"LogAlertConfigs", func() interface{} { return api.LogAlertConfigs() }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.method()
			if result == nil {
				t.Errorf("%s() returned nil", tt.name)
			}
		})
	}
}

// TestNewInstanaAPIWithUserAgent_EmptyUserAgent tests with empty user agent
func TestNewInstanaAPIWithUserAgent_EmptyUserAgent(t *testing.T) {
	api := NewInstanaAPIWithUserAgent(
		"test-token",
		"test.instana.io",
		false,
		"",
	)

	if api == nil {
		t.Fatal("Expected API client to be created even with empty user agent")
	}
}

// Made with Bob
