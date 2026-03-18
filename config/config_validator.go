package config

import (
	"fmt"
	"time"
)

// ValidationError represents a configuration validation error
type ValidationError struct {
	Field   string
	Message string
}

// Error implements the error interface
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error for field '%s': %s", e.Field, e.Message)
}

// ValidationErrors is a collection of validation errors
type ValidationErrors []ValidationError

// Error implements the error interface
func (e ValidationErrors) Error() string {
	if len(e) == 0 {
		return "no validation errors"
	}
	if len(e) == 1 {
		return e[0].Error()
	}
	msg := fmt.Sprintf("%d validation errors:", len(e))
	for _, err := range e {
		msg += fmt.Sprintf("\n  - %s", err.Error())
	}
	return msg
}

// HasErrors returns true if there are any validation errors
func (e ValidationErrors) HasErrors() bool {
	return len(e) > 0
}

// Validate validates the ClientConfig and returns any validation errors
func (c *ClientConfig) Validate() error {
	var errors ValidationErrors

	// Validate BaseURL
	if c.BaseURL == "" {
		errors = append(errors, ValidationError{
			Field:   "BaseURL",
			Message: "base URL is required",
		})
	}

	// Validate APIToken
	if c.APIToken == "" {
		errors = append(errors, ValidationError{
			Field:   "APIToken",
			Message: "API token is required",
		})
	}

	// Validate Timeout configuration
	errors = append(errors, c.validateTimeout()...)

	// Validate Retry configuration
	errors = append(errors, c.validateRetry()...)

	// Validate Batch configuration
	errors = append(errors, c.validateBatch()...)

	// Validate RateLimit configuration
	errors = append(errors, c.validateRateLimit()...)

	// Validate ConnectionPool configuration
	errors = append(errors, c.validateConnectionPool()...)

	if errors.HasErrors() {
		return errors
	}

	return nil
}

// validateTimeout validates timeout configuration
func (c *ClientConfig) validateTimeout() ValidationErrors {
	var errors ValidationErrors

	if c.Timeout.Connection <= 0 {
		errors = append(errors, ValidationError{
			Field:   "Timeout.Connection",
			Message: "connection timeout must be positive",
		})
	}

	if c.Timeout.Connection > 5*time.Minute {
		errors = append(errors, ValidationError{
			Field:   "Timeout.Connection",
			Message: "connection timeout should not exceed 5 minutes",
		})
	}

	if c.Timeout.Request <= 0 {
		errors = append(errors, ValidationError{
			Field:   "Timeout.Request",
			Message: "request timeout must be positive",
		})
	}

	if c.Timeout.Request > 10*time.Minute {
		errors = append(errors, ValidationError{
			Field:   "Timeout.Request",
			Message: "request timeout should not exceed 10 minutes",
		})
	}

	if c.Timeout.IdleConnection <= 0 {
		errors = append(errors, ValidationError{
			Field:   "Timeout.IdleConnection",
			Message: "idle connection timeout must be positive",
		})
	}

	if c.Timeout.ResponseHeader <= 0 {
		errors = append(errors, ValidationError{
			Field:   "Timeout.ResponseHeader",
			Message: "response header timeout must be positive",
		})
	}

	if c.Timeout.TLSHandshake <= 0 {
		errors = append(errors, ValidationError{
			Field:   "Timeout.TLSHandshake",
			Message: "TLS handshake timeout must be positive",
		})
	}

	// Logical validation: Request timeout should be >= Connection timeout
	if c.Timeout.Request < c.Timeout.Connection {
		errors = append(errors, ValidationError{
			Field:   "Timeout.Request",
			Message: "request timeout should be greater than or equal to connection timeout",
		})
	}

	return errors
}

// validateRetry validates retry configuration
func (c *ClientConfig) validateRetry() ValidationErrors {
	var errors ValidationErrors

	if c.Retry.MaxAttempts < 0 {
		errors = append(errors, ValidationError{
			Field:   "Retry.MaxAttempts",
			Message: "max attempts cannot be negative",
		})
	}

	if c.Retry.MaxAttempts > 10 {
		errors = append(errors, ValidationError{
			Field:   "Retry.MaxAttempts",
			Message: "max attempts should not exceed 10",
		})
	}

	if c.Retry.InitialDelay < 0 {
		errors = append(errors, ValidationError{
			Field:   "Retry.InitialDelay",
			Message: "initial delay cannot be negative",
		})
	}

	if c.Retry.InitialDelay > 1*time.Minute {
		errors = append(errors, ValidationError{
			Field:   "Retry.InitialDelay",
			Message: "initial delay should not exceed 1 minute",
		})
	}

	if c.Retry.MaxDelay < 0 {
		errors = append(errors, ValidationError{
			Field:   "Retry.MaxDelay",
			Message: "max delay cannot be negative",
		})
	}

	if c.Retry.MaxDelay > 5*time.Minute {
		errors = append(errors, ValidationError{
			Field:   "Retry.MaxDelay",
			Message: "max delay should not exceed 5 minutes",
		})
	}

	if c.Retry.BackoffMultiplier < 1.0 {
		errors = append(errors, ValidationError{
			Field:   "Retry.BackoffMultiplier",
			Message: "backoff multiplier must be at least 1.0",
		})
	}

	if c.Retry.BackoffMultiplier > 10.0 {
		errors = append(errors, ValidationError{
			Field:   "Retry.BackoffMultiplier",
			Message: "backoff multiplier should not exceed 10.0",
		})
	}

	// Logical validation: MaxDelay should be >= InitialDelay
	if c.Retry.MaxDelay > 0 && c.Retry.InitialDelay > 0 && c.Retry.MaxDelay < c.Retry.InitialDelay {
		errors = append(errors, ValidationError{
			Field:   "Retry.MaxDelay",
			Message: "max delay should be greater than or equal to initial delay",
		})
	}

	// Validate retryable status codes
	for _, code := range c.Retry.RetryableStatusCodes {
		if code < 100 || code > 599 {
			errors = append(errors, ValidationError{
				Field:   "Retry.RetryableStatusCodes",
				Message: fmt.Sprintf("invalid HTTP status code: %d", code),
			})
		}
	}

	return errors
}

// validateBatch validates batch configuration
func (c *ClientConfig) validateBatch() ValidationErrors {
	var errors ValidationErrors

	if c.Batch.Size <= 0 {
		errors = append(errors, ValidationError{
			Field:   "Batch.Size",
			Message: "batch size must be positive",
		})
	}

	if c.Batch.Size > 1000 {
		errors = append(errors, ValidationError{
			Field:   "Batch.Size",
			Message: "batch size should not exceed 1000",
		})
	}

	if c.Batch.ConcurrentRequests <= 0 {
		errors = append(errors, ValidationError{
			Field:   "Batch.ConcurrentRequests",
			Message: "concurrent requests must be positive",
		})
	}

	if c.Batch.ConcurrentRequests > 100 {
		errors = append(errors, ValidationError{
			Field:   "Batch.ConcurrentRequests",
			Message: "concurrent requests should not exceed 100",
		})
	}

	return errors
}

// validateRateLimit validates rate limit configuration
func (c *ClientConfig) validateRateLimit() ValidationErrors {
	var errors ValidationErrors

	if c.RateLimit.Enabled {
		if c.RateLimit.RequestsPerSecond <= 0 {
			errors = append(errors, ValidationError{
				Field:   "RateLimit.RequestsPerSecond",
				Message: "requests per second must be positive when rate limiting is enabled",
			})
		}

		if c.RateLimit.RequestsPerSecond > 10000 {
			errors = append(errors, ValidationError{
				Field:   "RateLimit.RequestsPerSecond",
				Message: "requests per second should not exceed 10000",
			})
		}

		if c.RateLimit.BurstCapacity <= 0 {
			errors = append(errors, ValidationError{
				Field:   "RateLimit.BurstCapacity",
				Message: "burst capacity must be positive when rate limiting is enabled",
			})
		}

		if c.RateLimit.BurstCapacity > 20000 {
			errors = append(errors, ValidationError{
				Field:   "RateLimit.BurstCapacity",
				Message: "burst capacity should not exceed 20000",
			})
		}

		// Logical validation: BurstCapacity should be >= RequestsPerSecond
		if c.RateLimit.BurstCapacity < c.RateLimit.RequestsPerSecond {
			errors = append(errors, ValidationError{
				Field:   "RateLimit.BurstCapacity",
				Message: "burst capacity should be greater than or equal to requests per second",
			})
		}
	}

	return errors
}

// validateConnectionPool validates connection pool configuration
func (c *ClientConfig) validateConnectionPool() ValidationErrors {
	var errors ValidationErrors

	if c.ConnectionPool.MaxIdleConnections < 0 {
		errors = append(errors, ValidationError{
			Field:   "ConnectionPool.MaxIdleConnections",
			Message: "max idle connections cannot be negative",
		})
	}

	if c.ConnectionPool.MaxIdleConnections > 1000 {
		errors = append(errors, ValidationError{
			Field:   "ConnectionPool.MaxIdleConnections",
			Message: "max idle connections should not exceed 1000",
		})
	}

	if c.ConnectionPool.MaxConnectionsPerHost <= 0 {
		errors = append(errors, ValidationError{
			Field:   "ConnectionPool.MaxConnectionsPerHost",
			Message: "max connections per host must be positive",
		})
	}

	if c.ConnectionPool.MaxConnectionsPerHost > 100 {
		errors = append(errors, ValidationError{
			Field:   "ConnectionPool.MaxConnectionsPerHost",
			Message: "max connections per host should not exceed 100",
		})
	}

	if c.ConnectionPool.MaxIdleConnectionsPerHost < 0 {
		errors = append(errors, ValidationError{
			Field:   "ConnectionPool.MaxIdleConnectionsPerHost",
			Message: "max idle connections per host cannot be negative",
		})
	}

	if c.ConnectionPool.MaxIdleConnectionsPerHost > c.ConnectionPool.MaxConnectionsPerHost {
		errors = append(errors, ValidationError{
			Field:   "ConnectionPool.MaxIdleConnectionsPerHost",
			Message: "max idle connections per host should not exceed max connections per host",
		})
	}

	if c.ConnectionPool.KeepAliveDuration < 0 {
		errors = append(errors, ValidationError{
			Field:   "ConnectionPool.KeepAliveDuration",
			Message: "keep-alive duration cannot be negative",
		})
	}

	if c.ConnectionPool.KeepAliveDuration > 10*time.Minute {
		errors = append(errors, ValidationError{
			Field:   "ConnectionPool.KeepAliveDuration",
			Message: "keep-alive duration should not exceed 10 minutes",
		})
	}

	return errors
}
