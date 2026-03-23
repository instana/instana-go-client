package config

import (
	"context"
	"math"
	"math/rand"
	"time"
)

// RetryFunc is a function that can be retried
type RetryFunc func() error

// RetryableFuncWithValue is a function that returns a value and can be retried
type RetryableFuncWithValue func() (interface{}, error)

// Retryer handles retry logic with exponential backoff
type Retryer struct {
	config RetryConfig
	logger Logger
}

// NewRetryer creates a new Retryer with the given configuration
func NewRetryer(config RetryConfig, logger Logger) *Retryer {
	if logger == nil {
		logger = NewNoOpLogger()
	}
	return &Retryer{
		config: config,
		logger: logger,
	}
}

// Do executes the function with retry logic
func (r *Retryer) Do(ctx context.Context, fn RetryFunc) error {
	var lastErr error

	for attempt := 0; attempt <= r.config.MaxAttempts; attempt++ {
		// Check context cancellation
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// Execute the function
		err := fn()
		if err == nil {
			if attempt > 0 {
				r.logger.Info("Operation succeeded after retry",
					"attempt", attempt+1,
					"total_attempts", r.config.MaxAttempts+1)
			}
			return nil
		}

		lastErr = err

		// Check if we should retry
		if !r.shouldRetry(err, attempt) {
			r.logger.Debug("Not retrying",
				"error", err.Error(),
				"attempt", attempt+1,
				"retryable", IsRetryableError(err))
			return err
		}

		// Calculate delay for next attempt
		if attempt < r.config.MaxAttempts {
			delay := r.calculateDelay(attempt)
			r.logger.Warn("Operation failed, retrying",
				"error", err.Error(),
				"attempt", attempt+1,
				"max_attempts", r.config.MaxAttempts+1,
				"retry_delay", delay.String())

			// Wait before retry
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(delay):
				// Continue to next attempt
			}
		}
	}

	r.logger.Error("Operation failed after all retry attempts",
		"error", lastErr.Error(),
		"total_attempts", r.config.MaxAttempts+1)

	return lastErr
}

// DoWithValue executes the function with retry logic and returns a value
func (r *Retryer) DoWithValue(ctx context.Context, fn RetryableFuncWithValue) (interface{}, error) {
	var lastErr error

	for attempt := 0; attempt <= r.config.MaxAttempts; attempt++ {
		// Check context cancellation
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// Execute the function
		result, err := fn()
		if err == nil {
			if attempt > 0 {
				r.logger.Info("Operation succeeded after retry",
					"attempt", attempt+1,
					"total_attempts", r.config.MaxAttempts+1)
			}
			return result, nil
		}

		lastErr = err

		// Check if we should retry
		if !r.shouldRetry(err, attempt) {
			r.logger.Debug("Not retrying",
				"error", err.Error(),
				"attempt", attempt+1,
				"retryable", IsRetryableError(err))
			return nil, err
		}

		// Calculate delay for next attempt
		if attempt < r.config.MaxAttempts {
			delay := r.calculateDelay(attempt)
			r.logger.Warn("Operation failed, retrying",
				"error", err.Error(),
				"attempt", attempt+1,
				"max_attempts", r.config.MaxAttempts+1,
				"retry_delay", delay.String())

			// Wait before retry
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(delay):
				// Continue to next attempt
			}
		}
	}

	r.logger.Error("Operation failed after all retry attempts",
		"error", lastErr.Error(),
		"total_attempts", r.config.MaxAttempts+1)

	return nil, lastErr
}

// shouldRetry determines if an error should trigger a retry
func (r *Retryer) shouldRetry(err error, attempt int) bool {
	// Don't retry if we've exhausted attempts
	if attempt >= r.config.MaxAttempts {
		return false
	}

	// Check if error is retryable
	if !IsRetryableError(err) {
		return false
	}

	// Check specific error types based on configuration
	if instanaErr, ok := err.(*InstanaError); ok {
		switch instanaErr.Type {
		case ErrorTypeTimeout:
			return r.config.RetryOnTimeout
		case ErrorTypeNetwork:
			return r.config.RetryOnConnectionError
		case ErrorTypeRateLimit, ErrorTypeAPI:
			// Check if status code is in retryable list
			if instanaErr.StatusCode > 0 {
				return r.isRetryableStatusCode(instanaErr.StatusCode)
			}
			return true
		default:
			return instanaErr.IsRetryable()
		}
	}

	return true
}

// isRetryableStatusCode checks if a status code is in the retryable list
func (r *Retryer) isRetryableStatusCode(statusCode int) bool {
	for _, code := range r.config.RetryableStatusCodes {
		if code == statusCode {
			return true
		}
	}
	return false
}

// calculateDelay calculates the delay before the next retry attempt
func (r *Retryer) calculateDelay(attempt int) time.Duration {
	// Calculate exponential backoff
	delay := float64(r.config.InitialDelay) * math.Pow(r.config.BackoffMultiplier, float64(attempt))

	// Apply jitter if enabled
	if r.config.Jitter {
		jitter := rand.Float64() * 0.3 * delay //nolint:gosec // Weak random is acceptable for jitter
		delay = delay + jitter
	}

	// Cap at max delay
	if delay > float64(r.config.MaxDelay) {
		delay = float64(r.config.MaxDelay)
	}

	return time.Duration(delay)
}

// RetryWithBackoff is a convenience function for simple retry scenarios
func RetryWithBackoff(ctx context.Context, config RetryConfig, fn RetryFunc) error {
	retryer := NewRetryer(config, nil)
	return retryer.Do(ctx, fn)
}

// RetryWithBackoffAndValue is a convenience function for retry scenarios with return values
func RetryWithBackoffAndValue(ctx context.Context, config RetryConfig, fn RetryableFuncWithValue) (interface{}, error) {
	retryer := NewRetryer(config, nil)
	return retryer.DoWithValue(ctx, fn)
}

// DefaultRetryConfig returns a default retry configuration
func DefaultRetryConfig() RetryConfig {
	return RetryConfig{
		MaxAttempts:            3,
		InitialDelay:           1 * time.Second,
		MaxDelay:               30 * time.Second,
		BackoffMultiplier:      2.0,
		RetryableStatusCodes:   []int{408, 429, 500, 502, 503, 504},
		RetryOnTimeout:         true,
		RetryOnConnectionError: true,
		Jitter:                 true,
	}
}
