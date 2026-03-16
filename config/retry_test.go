package config

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestNewRetryer(t *testing.T) {
	config := DefaultRetryConfig()
	retryer := NewRetryer(config, nil)

	if retryer == nil {
		t.Fatal("NewRetryer() returned nil")
	}

	if retryer.config.MaxAttempts != config.MaxAttempts {
		t.Error("Config not properly set")
	}

	if retryer.logger == nil {
		t.Error("Logger should be set to NoOpLogger when nil is passed")
	}
}

func TestRetryerDoSuccess(t *testing.T) {
	config := DefaultRetryConfig()
	retryer := NewRetryer(config, nil)

	callCount := 0
	fn := func() error {
		callCount++
		return nil
	}

	ctx := context.Background()
	err := retryer.Do(ctx, fn)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if callCount != 1 {
		t.Errorf("Expected function to be called once, got %d calls", callCount)
	}
}

func TestRetryerDoSuccessAfterRetries(t *testing.T) {
	config := DefaultRetryConfig()
	config.InitialDelay = 10 * time.Millisecond
	config.MaxDelay = 50 * time.Millisecond
	retryer := NewRetryer(config, nil)

	callCount := 0
	fn := func() error {
		callCount++
		if callCount < 3 {
			return NetworkError("temporary error", nil)
		}
		return nil
	}

	ctx := context.Background()
	err := retryer.Do(ctx, fn)

	if err != nil {
		t.Errorf("Expected no error after retries, got: %v", err)
	}

	if callCount != 3 {
		t.Errorf("Expected function to be called 3 times, got %d calls", callCount)
	}
}

func TestRetryerDoFailureNonRetryable(t *testing.T) {
	config := DefaultRetryConfig()
	retryer := NewRetryer(config, nil)

	callCount := 0
	expectedErr := NewValidationError("validation error", nil)
	fn := func() error {
		callCount++
		return expectedErr
	}

	ctx := context.Background()
	err := retryer.Do(ctx, fn)

	if err != expectedErr {
		t.Errorf("Expected validation error, got: %v", err)
	}

	if callCount != 1 {
		t.Errorf("Expected function to be called once (no retries), got %d calls", callCount)
	}
}

func TestRetryerDoFailureAfterMaxAttempts(t *testing.T) {
	config := DefaultRetryConfig()
	config.MaxAttempts = 2
	config.InitialDelay = 10 * time.Millisecond
	retryer := NewRetryer(config, nil)

	callCount := 0
	expectedErr := NetworkError("persistent error", nil)
	fn := func() error {
		callCount++
		return expectedErr
	}

	ctx := context.Background()
	err := retryer.Do(ctx, fn)

	if err == nil {
		t.Error("Expected error after max attempts")
	}

	// MaxAttempts + 1 (initial attempt)
	expectedCalls := config.MaxAttempts + 1
	if callCount != expectedCalls {
		t.Errorf("Expected function to be called %d times, got %d calls", expectedCalls, callCount)
	}
}

func TestRetryerDoContextCancellation(t *testing.T) {
	config := DefaultRetryConfig()
	config.InitialDelay = 100 * time.Millisecond
	retryer := NewRetryer(config, nil)

	ctx, cancel := context.WithCancel(context.Background())

	callCount := 0
	fn := func() error {
		callCount++
		if callCount == 1 {
			// Cancel context after first attempt
			cancel()
		}
		return NetworkError("error", nil)
	}

	err := retryer.Do(ctx, fn)

	if err != context.Canceled {
		t.Errorf("Expected context.Canceled error, got: %v", err)
	}

	if callCount > 2 {
		t.Errorf("Expected at most 2 calls before cancellation, got %d calls", callCount)
	}
}

func TestRetryerDoWithValue(t *testing.T) {
	config := DefaultRetryConfig()
	config.InitialDelay = 10 * time.Millisecond
	retryer := NewRetryer(config, nil)

	callCount := 0
	expectedValue := "success"
	fn := func() (interface{}, error) {
		callCount++
		if callCount < 2 {
			return nil, NetworkError("temporary error", nil)
		}
		return expectedValue, nil
	}

	ctx := context.Background()
	result, err := retryer.DoWithValue(ctx, fn)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if result != expectedValue {
		t.Errorf("Expected result %v, got %v", expectedValue, result)
	}

	if callCount != 2 {
		t.Errorf("Expected 2 calls, got %d", callCount)
	}
}

func TestRetryerDoWithValueFailure(t *testing.T) {
	config := DefaultRetryConfig()
	config.MaxAttempts = 1
	config.InitialDelay = 10 * time.Millisecond
	retryer := NewRetryer(config, nil)

	expectedErr := NetworkError("persistent error", nil)
	fn := func() (interface{}, error) {
		return nil, expectedErr
	}

	ctx := context.Background()
	result, err := retryer.DoWithValue(ctx, fn)

	if err == nil {
		t.Error("Expected error")
	}

	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}
}

func TestShouldRetryWithDifferentErrors(t *testing.T) {
	config := DefaultRetryConfig()
	retryer := NewRetryer(config, nil)

	tests := []struct {
		name        string
		err         error
		attempt     int
		shouldRetry bool
	}{
		{"network error", NetworkError("test", nil), 0, true},
		{"timeout error", TimeoutError("test", nil), 0, true},
		{"rate limit error", RateLimitError("test", 60), 0, true},
		{"API 500 error", APIError(500, "test", nil), 0, true},
		{"API 400 error", APIError(400, "test", nil), 0, false},
		{"validation error", NewValidationError("test", nil), 0, false},
		{"authentication error", AuthenticationError("test", nil), 0, false},
		{"max attempts reached", NetworkError("test", nil), config.MaxAttempts, false},
		{"standard error", errors.New("test"), 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := retryer.shouldRetry(tt.err, tt.attempt)
			if result != tt.shouldRetry {
				t.Errorf("Expected shouldRetry=%v, got %v", tt.shouldRetry, result)
			}
		})
	}
}

func TestShouldRetryWithConfigFlags(t *testing.T) {
	tests := []struct {
		name                   string
		retryOnTimeout         bool
		retryOnConnectionError bool
		err                    error
		shouldRetry            bool
	}{
		{
			name:                   "timeout with retry enabled",
			retryOnTimeout:         true,
			retryOnConnectionError: true,
			err:                    TimeoutError("test", nil),
			shouldRetry:            true,
		},
		{
			name:                   "timeout with retry disabled",
			retryOnTimeout:         false,
			retryOnConnectionError: true,
			err:                    TimeoutError("test", nil),
			shouldRetry:            false,
		},
		{
			name:                   "network error with retry enabled",
			retryOnTimeout:         true,
			retryOnConnectionError: true,
			err:                    NetworkError("test", nil),
			shouldRetry:            true,
		},
		{
			name:                   "network error with retry disabled",
			retryOnTimeout:         true,
			retryOnConnectionError: false,
			err:                    NetworkError("test", nil),
			shouldRetry:            false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := DefaultRetryConfig()
			config.RetryOnTimeout = tt.retryOnTimeout
			config.RetryOnConnectionError = tt.retryOnConnectionError
			retryer := NewRetryer(config, nil)

			result := retryer.shouldRetry(tt.err, 0)
			if result != tt.shouldRetry {
				t.Errorf("Expected shouldRetry=%v, got %v", tt.shouldRetry, result)
			}
		})
	}
}

func TestIsRetryableStatusCode(t *testing.T) {
	config := DefaultRetryConfig()
	config.RetryableStatusCodes = []int{500, 502, 503}
	retryer := NewRetryer(config, nil)

	tests := []struct {
		statusCode int
		retryable  bool
	}{
		{500, true},
		{502, true},
		{503, true},
		{400, false},
		{404, false},
		{429, false}, // Not in custom list
	}

	for _, tt := range tests {
		t.Run(string(rune(tt.statusCode)), func(t *testing.T) {
			result := retryer.isRetryableStatusCode(tt.statusCode)
			if result != tt.retryable {
				t.Errorf("Status code %d: expected retryable=%v, got %v",
					tt.statusCode, tt.retryable, result)
			}
		})
	}
}

func TestCalculateDelay(t *testing.T) {
	config := DefaultRetryConfig()
	config.InitialDelay = 1 * time.Second
	config.BackoffMultiplier = 2.0
	config.MaxDelay = 10 * time.Second
	config.Jitter = false // Disable jitter for predictable testing
	retryer := NewRetryer(config, nil)

	tests := []struct {
		attempt       int
		expectedDelay time.Duration
	}{
		{0, 1 * time.Second},  // 1 * 2^0 = 1
		{1, 2 * time.Second},  // 1 * 2^1 = 2
		{2, 4 * time.Second},  // 1 * 2^2 = 4
		{3, 8 * time.Second},  // 1 * 2^3 = 8
		{4, 10 * time.Second}, // 1 * 2^4 = 16, capped at 10
		{5, 10 * time.Second}, // 1 * 2^5 = 32, capped at 10
	}

	for _, tt := range tests {
		t.Run(string(rune(tt.attempt)), func(t *testing.T) {
			delay := retryer.calculateDelay(tt.attempt)
			if delay != tt.expectedDelay {
				t.Errorf("Attempt %d: expected delay %v, got %v",
					tt.attempt, tt.expectedDelay, delay)
			}
		})
	}
}

func TestCalculateDelayWithJitter(t *testing.T) {
	config := DefaultRetryConfig()
	config.InitialDelay = 1 * time.Second
	config.BackoffMultiplier = 2.0
	config.MaxDelay = 10 * time.Second
	config.Jitter = true
	retryer := NewRetryer(config, nil)

	// Test that jitter adds variability
	delays := make([]time.Duration, 10)
	for i := 0; i < 10; i++ {
		delays[i] = retryer.calculateDelay(1)
	}

	// Check that not all delays are the same (jitter is working)
	allSame := true
	for i := 1; i < len(delays); i++ {
		if delays[i] != delays[0] {
			allSame = false
			break
		}
	}

	if allSame {
		t.Error("Expected jitter to produce different delays, but all were the same")
	}

	// Check that all delays are within expected range (2s ± 30%)
	baseDelay := 2 * time.Second
	minDelay := baseDelay
	maxDelay := time.Duration(float64(baseDelay) * 1.3)

	for i, delay := range delays {
		if delay < minDelay || delay > maxDelay {
			t.Errorf("Delay %d (%v) is outside expected range [%v, %v]",
				i, delay, minDelay, maxDelay)
		}
	}
}

func TestRetryWithBackoff(t *testing.T) {
	config := DefaultRetryConfig()
	config.MaxAttempts = 2
	config.InitialDelay = 10 * time.Millisecond

	callCount := 0
	fn := func() error {
		callCount++
		if callCount < 2 {
			return NetworkError("temporary error", nil)
		}
		return nil
	}

	ctx := context.Background()
	err := RetryWithBackoff(ctx, config, fn)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if callCount != 2 {
		t.Errorf("Expected 2 calls, got %d", callCount)
	}
}

func TestRetryWithBackoffAndValue(t *testing.T) {
	config := DefaultRetryConfig()
	config.MaxAttempts = 2
	config.InitialDelay = 10 * time.Millisecond

	callCount := 0
	expectedValue := 42
	fn := func() (interface{}, error) {
		callCount++
		if callCount < 2 {
			return nil, NetworkError("temporary error", nil)
		}
		return expectedValue, nil
	}

	ctx := context.Background()
	result, err := RetryWithBackoffAndValue(ctx, config, fn)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if result != expectedValue {
		t.Errorf("Expected result %v, got %v", expectedValue, result)
	}
}

func TestDefaultRetryConfig(t *testing.T) {
	config := DefaultRetryConfig()

	if config.MaxAttempts != 3 {
		t.Errorf("Expected MaxAttempts=3, got %d", config.MaxAttempts)
	}

	if config.InitialDelay != 1*time.Second {
		t.Errorf("Expected InitialDelay=1s, got %v", config.InitialDelay)
	}

	if config.MaxDelay != 30*time.Second {
		t.Errorf("Expected MaxDelay=30s, got %v", config.MaxDelay)
	}

	if config.BackoffMultiplier != 2.0 {
		t.Errorf("Expected BackoffMultiplier=2.0, got %f", config.BackoffMultiplier)
	}

	if !config.RetryOnTimeout {
		t.Error("Expected RetryOnTimeout=true")
	}

	if !config.RetryOnConnectionError {
		t.Error("Expected RetryOnConnectionError=true")
	}

	if !config.Jitter {
		t.Error("Expected Jitter=true")
	}

	expectedCodes := []int{408, 429, 500, 502, 503, 504}
	if len(config.RetryableStatusCodes) != len(expectedCodes) {
		t.Errorf("Expected %d retryable status codes, got %d",
			len(expectedCodes), len(config.RetryableStatusCodes))
	}
}

func TestRetryerWithCustomLogger(t *testing.T) {
	config := DefaultRetryConfig()
	config.InitialDelay = 10 * time.Millisecond

	// Create a simple test logger that counts calls
	logCalls := 0
	logger := &testLogger{
		onLog: func() { logCalls++ },
	}

	retryer := NewRetryer(config, logger)

	callCount := 0
	fn := func() error {
		callCount++
		if callCount < 2 {
			return NetworkError("temporary error", nil)
		}
		return nil
	}

	ctx := context.Background()
	_ = retryer.Do(ctx, fn)

	if logCalls == 0 {
		t.Error("Expected logger to be called")
	}
}

// testLogger is a simple logger for testing
type testLogger struct {
	onLog func()
}

func (l *testLogger) Debug(msg string, keysAndValues ...interface{}) {
	if l.onLog != nil {
		l.onLog()
	}
}

func (l *testLogger) Info(msg string, keysAndValues ...interface{}) {
	if l.onLog != nil {
		l.onLog()
	}
}

func (l *testLogger) Warn(msg string, keysAndValues ...interface{}) {
	if l.onLog != nil {
		l.onLog()
	}
}

func (l *testLogger) Error(msg string, keysAndValues ...interface{}) {
	if l.onLog != nil {
		l.onLog()
	}
}

func BenchmarkRetryerDoSuccess(b *testing.B) {
	config := DefaultRetryConfig()
	retryer := NewRetryer(config, nil)
	ctx := context.Background()

	fn := func() error {
		return nil
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = retryer.Do(ctx, fn)
	}
}

func BenchmarkRetryerDoWithRetries(b *testing.B) {
	config := DefaultRetryConfig()
	config.MaxAttempts = 2
	config.InitialDelay = 1 * time.Millisecond
	retryer := NewRetryer(config, nil)
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		callCount := 0
		fn := func() error {
			callCount++
			if callCount < 2 {
				return NetworkError("error", nil)
			}
			return nil
		}
		_ = retryer.Do(ctx, fn)
	}
}

func BenchmarkCalculateDelay(b *testing.B) {
	config := DefaultRetryConfig()
	retryer := NewRetryer(config, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = retryer.calculateDelay(i % 5)
	}
}

// Made with Bob
