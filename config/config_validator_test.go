package config

import (
	"strings"
	"testing"
	"time"
)

func TestValidateValidConfig(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"

	err := config.Validate()
	if err != nil {
		t.Errorf("Expected valid config to pass validation, got error: %v", err)
	}
}

func TestValidateEmptyBaseURL(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = ""
	config.APIToken = "test-token"

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for empty BaseURL")
	}

	if !strings.Contains(err.Error(), "BaseURL") {
		t.Errorf("Expected error message to mention BaseURL, got: %v", err)
	}
}

func TestValidateEmptyAPIToken(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = ""

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for empty APIToken")
	}

	if !strings.Contains(err.Error(), "APIToken") {
		t.Errorf("Expected error message to mention APIToken, got: %v", err)
	}
}

func TestValidateInvalidConnectionTimeout(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.Timeout.Connection = 0

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for zero Connection timeout")
	}
}

func TestValidateConnectionTimeoutTooLarge(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.Timeout.Connection = 6 * time.Minute

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for Connection timeout > 5 minutes")
	}
}

func TestValidateInvalidRequestTimeout(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.Timeout.Request = -1 * time.Second

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for negative Request timeout")
	}
}

func TestValidateRequestTimeoutTooLarge(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.Timeout.Request = 11 * time.Minute

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for Request timeout > 10 minutes")
	}
}

func TestValidateInvalidRetryAttempts(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.Retry.MaxAttempts = -1

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for negative MaxAttempts")
	}
}

func TestValidateRetryAttemptsTooLarge(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.Retry.MaxAttempts = 11

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for MaxAttempts > 10")
	}
}

func TestValidateInvalidRetryInitialDelay(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.Retry.InitialDelay = -1 * time.Second

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for negative InitialDelay")
	}
}

func TestValidateRetryMaxDelayLessThanInitial(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.Retry.InitialDelay = 5 * time.Second
	config.Retry.MaxDelay = 2 * time.Second

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for MaxDelay < InitialDelay")
	}
}

func TestValidateInvalidBackoffMultiplier(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.Retry.BackoffMultiplier = 0.5

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for BackoffMultiplier < 1.0")
	}
}

func TestValidateBackoffMultiplierTooLarge(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.Retry.BackoffMultiplier = 11.0

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for BackoffMultiplier > 10.0")
	}
}

func TestValidateInvalidRetryableStatusCode(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.Retry.RetryableStatusCodes = []int{99, 500}

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for status code < 100")
	}
}

func TestValidateRetryableStatusCodeTooLarge(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.Retry.RetryableStatusCodes = []int{500, 600}

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for status code > 599")
	}
}

func TestValidateInvalidRateLimitRPS(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.RateLimit.RequestsPerSecond = 0

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for zero RequestsPerSecond")
	}
}

func TestValidateRateLimitRPSTooLarge(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.RateLimit.RequestsPerSecond = 10001

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for RequestsPerSecond > 10000")
	}
}

func TestValidateInvalidBurstCapacity(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.RateLimit.BurstCapacity = 0

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for zero BurstCapacity")
	}
}

func TestValidateBurstCapacityLessThanRPS(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.RateLimit.RequestsPerSecond = 100
	config.RateLimit.BurstCapacity = 50

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for BurstCapacity < RequestsPerSecond")
	}
}

func TestValidateInvalidMaxIdleConnections(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.ConnectionPool.MaxIdleConnections = -1

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for negative MaxIdleConnections")
	}
}

func TestValidateMaxIdleConnectionsTooLarge(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.ConnectionPool.MaxIdleConnections = 1001

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for MaxIdleConnections > 1000")
	}
}

func TestValidateInvalidMaxConnectionsPerHost(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.ConnectionPool.MaxConnectionsPerHost = 0

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for zero MaxConnectionsPerHost")
	}
}

func TestValidateInvalidBatchSize(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.Batch.Size = 0

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for zero Batch.Size")
	}
}

func TestValidateBatchSizeTooLarge(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.Batch.Size = 1001

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for Batch.Size > 1000")
	}
}

func TestValidateInvalidConcurrentRequests(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"
	config.Batch.ConcurrentRequests = 0

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation error for zero ConcurrentRequests")
	}
}

func TestValidationErrorsMultiple(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = ""           // Error 1
	config.APIToken = ""          // Error 2
	config.Timeout.Connection = 0 // Error 3

	err := config.Validate()
	if err == nil {
		t.Error("Expected validation errors")
	}

	// Check that error message contains multiple errors
	errMsg := err.Error()
	if !strings.Contains(errMsg, "BaseURL") {
		t.Error("Expected error message to contain BaseURL error")
	}
	if !strings.Contains(errMsg, "APIToken") {
		t.Error("Expected error message to contain APIToken error")
	}
	if !strings.Contains(errMsg, "Connection") {
		t.Error("Expected error message to contain Connection timeout error")
	}
}

func TestValidationErrorsType(t *testing.T) {
	config := DefaultClientConfig()
	config.BaseURL = ""
	config.APIToken = ""

	err := config.Validate()
	if err == nil {
		t.Fatal("Expected validation error")
	}

	// Check if error is of type ValidationErrors
	_, ok := err.(ValidationErrors)
	if !ok {
		t.Errorf("Expected error to be of type ValidationErrors, got %T", err)
	}
}

func TestValidateEdgeCases(t *testing.T) {
	tests := []struct {
		name      string
		setupFunc func(*ClientConfig)
		wantError bool
	}{
		{
			name: "Minimum valid connection timeout",
			setupFunc: func(c *ClientConfig) {
				c.Timeout.Connection = 1 * time.Millisecond
			},
			wantError: false,
		},
		{
			name: "Maximum valid connection timeout",
			setupFunc: func(c *ClientConfig) {
				c.Timeout.Connection = 5 * time.Minute
				c.Timeout.Request = 5 * time.Minute // Must be >= Connection timeout
			},
			wantError: false,
		},
		{
			name: "Zero retry attempts (valid)",
			setupFunc: func(c *ClientConfig) {
				c.Retry.MaxAttempts = 0
			},
			wantError: false,
		},
		{
			name: "Maximum retry attempts",
			setupFunc: func(c *ClientConfig) {
				c.Retry.MaxAttempts = 10
			},
			wantError: false,
		},
		{
			name: "Minimum backoff multiplier",
			setupFunc: func(c *ClientConfig) {
				c.Retry.BackoffMultiplier = 1.0
			},
			wantError: false,
		},
		{
			name: "Maximum backoff multiplier",
			setupFunc: func(c *ClientConfig) {
				c.Retry.BackoffMultiplier = 10.0
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := DefaultClientConfig()
			config.BaseURL = "https://tenant.instana.io"
			config.APIToken = "test-token"
			tt.setupFunc(config)

			err := config.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("Validate() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func BenchmarkValidate(b *testing.B) {
	config := DefaultClientConfig()
	config.BaseURL = "https://tenant.instana.io"
	config.APIToken = "test-token"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = config.Validate()
	}
}

func BenchmarkValidateWithErrors(b *testing.B) {
	config := DefaultClientConfig()
	config.BaseURL = ""
	config.APIToken = ""
	config.Timeout.Connection = 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = config.Validate()
	}
}

// Made with Bob
