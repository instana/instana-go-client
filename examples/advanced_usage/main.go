package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/instana/instana-go-client/config"
	"github.com/instana/instana-go-client/instana"
)

// This example demonstrates advanced usage patterns of the Instana Go Client,
// including configuration, error handling, retry logic, and rate limiting.

func main() {
	// Example 1: Advanced Configuration
	advancedConfiguration()

	// Example 2: Error Handling
	errorHandling()

	// Example 3: Retry Mechanism
	retryMechanism()

	// Example 4: Rate Limiting
	rateLimiting()

	// Example 5: Custom Logger
	customLogger()
}

// advancedConfiguration demonstrates comprehensive configuration options
func advancedConfiguration() {
	fmt.Println("\n=== Advanced Configuration ===")

	// Create configuration with all options
	cfg, err := config.NewConfigBuilder().
		// Core settings
		WithBaseURL("https://tenant-unit.instana.io").
		WithAPIToken("your-api-token").
		WithUserAgent("MyApp/1.0.0").
		WithDebug(true).
		// Timeout configuration
		WithConnectionTimeout(30*time.Second).
		WithRequestTimeout(60*time.Second).
		WithIdleConnectionTimeout(90*time.Second).
		WithResponseHeaderTimeout(10*time.Second).
		WithTLSHandshakeTimeout(10*time.Second).
		// Retry configuration
		WithMaxRetryAttempts(5).
		WithRetryInitialDelay(1*time.Second).
		WithRetryMaxDelay(30*time.Second).
		WithRetryBackoffMultiplier(2.0).
		WithRetryOnTimeout(true).
		WithRetryOnConnectionError(true).
		WithRetryJitter(true).
		// Rate limiting
		WithRateLimitEnabled(true).
		WithRateLimitRequestsPerSecond(100).
		WithRateLimitBurstCapacity(200).
		WithRateLimitWaitForToken(true).
		// Connection pooling
		WithMaxIdleConnections(100).
		WithMaxConnectionsPerHost(10).
		WithMaxIdleConnectionsPerHost(10).
		WithKeepAliveDuration(30*time.Second).
		WithDisableKeepAlives(false).
		WithDisableCompression(false).
		// Custom headers
		WithCustomHeader("X-Request-ID", "unique-request-id").
		WithCustomHeader("X-Trace-ID", "trace-123").
		// Batch configuration
		WithBatchSize(100).
		WithBatchConcurrentRequests(5).
		WithBatchStopOnError(false).
		WithBatchRetryFailedItems(true).
		Build()

	if err != nil {
		log.Fatalf("Failed to build config: %v", err)
	}

	// Create client with configuration
	api, err := instana.NewInstanaAPIWithConfig(cfg)
	if err != nil {
		log.Fatalf("Failed to create API client: %v", err)
	}

	fmt.Printf("✓ Client configured successfully\n")
	fmt.Printf("  Base URL: %s\n", cfg.BaseURL)
	fmt.Printf("  Rate Limit: %d req/s\n", cfg.RateLimit.RequestsPerSecond)
	fmt.Printf("  Max Retries: %d\n", cfg.Retry.MaxAttempts)

	// Use the client
	_ = api
}

// errorHandling demonstrates comprehensive error handling
func errorHandling() {
	fmt.Println("\n=== Error Handling ===")

	// Create a simple client for demonstration
	api := instana.NewInstanaAPI("invalid-token", "tenant.instana.io", false)

	// Attempt an operation that will fail
	_, err := api.APITokens().GetAll()
	if err != nil {
		// Check if it's an Instana error
		if instanaErr, ok := err.(*config.InstanaError); ok {
			fmt.Printf("✓ Caught InstanaError:\n")
			fmt.Printf("  Type: %s\n", instanaErr.Type)
			fmt.Printf("  Message: %s\n", instanaErr.Message)
			fmt.Printf("  Status Code: %d\n", instanaErr.StatusCode)
			fmt.Printf("  Retryable: %v\n", instanaErr.IsRetryable())

			// Handle different error types
			switch instanaErr.Type {
			case config.ErrorTypeAuthentication:
				fmt.Println("  → Authentication failed, check API token")
			case config.ErrorTypeRateLimit:
				fmt.Println("  → Rate limit exceeded, wait before retrying")
			case config.ErrorTypeNetwork:
				fmt.Println("  → Network error, check connectivity")
			case config.ErrorTypeTimeout:
				fmt.Println("  → Request timeout, increase timeout or retry")
			default:
				fmt.Printf("  → Unhandled error type: %s\n", instanaErr.Type)
			}
		}

		// Check if error is retryable
		if config.IsRetryableError(err) {
			fmt.Println("✓ Error is retryable")
		}

		// Check if error is temporary
		if config.IsTemporaryError(err) {
			fmt.Println("✓ Error is temporary")
		}

		// Extract status code
		statusCode := config.ExtractStatusCode(err)
		fmt.Printf("✓ HTTP Status Code: %d\n", statusCode)
	}
}

// retryMechanism demonstrates the retry mechanism
func retryMechanism() {
	fmt.Println("\n=== Retry Mechanism ===")

	// Create retry configuration
	retryConfig := config.RetryConfig{
		MaxAttempts:       5,
		InitialDelay:      1 * time.Second,
		MaxDelay:          30 * time.Second,
		BackoffMultiplier: 2.0,
		RetryOnTimeout:    true,
		Jitter:            true,
	}

	// Create logger
	logger := config.NewDefaultLogger(config.ClientLogLevelInfo)

	// Create retryer
	retryer := config.NewRetryer(retryConfig, logger)

	// Simulate an operation with retries
	attemptCount := 0
	err := retryer.Do(context.Background(), func() error {
		attemptCount++
		fmt.Printf("  Attempt %d...\n", attemptCount)

		// Simulate failure for first 2 attempts
		if attemptCount < 3 {
			return &config.InstanaError{
				Type:    config.ErrorTypeNetwork,
				Message: "simulated network error",
			}
		}

		// Success on 3rd attempt
		return nil
	})

	if err != nil {
		fmt.Printf("✗ Operation failed after %d attempts: %v\n", attemptCount, err)
	} else {
		fmt.Printf("✓ Operation succeeded after %d attempts\n", attemptCount)
	}
}

// rateLimiting demonstrates rate limiting
func rateLimiting() {
	fmt.Println("\n=== Rate Limiting ===")

	// Create rate limiter configuration
	rateLimitConfig := config.RateLimitConfig{
		Enabled:           true,
		RequestsPerSecond: 10,
		BurstCapacity:     20,
		WaitForToken:      true,
	}

	// Create logger
	logger := config.NewDefaultLogger(config.ClientLogLevelInfo)

	// Create rate limiter
	rateLimiter := config.NewRateLimiter(rateLimitConfig, logger)
	defer rateLimiter.Stop()

	fmt.Printf("✓ Rate limiter created: %d req/s, burst: %d\n",
		rateLimitConfig.RequestsPerSecond,
		rateLimitConfig.BurstCapacity)

	// Simulate multiple requests
	start := time.Now()
	for i := 0; i < 15; i++ {
		// Wait for rate limit token
		if err := rateLimiter.Wait(context.Background()); err != nil {
			fmt.Printf("✗ Rate limit error: %v\n", err)
			continue
		}

		fmt.Printf("  Request %d sent at %v\n", i+1, time.Since(start).Round(time.Millisecond))
	}

	fmt.Printf("✓ All requests completed in %v\n", time.Since(start).Round(time.Millisecond))
}

// customLogger demonstrates using a custom logger
func customLogger() {
	fmt.Println("\n=== Custom Logger ===")

	// Create custom logger
	logger := &CustomLogger{prefix: "[MyApp]"}

	// Create configuration with custom logger
	cfg, err := config.NewConfigBuilder().
		WithBaseURL("https://tenant.instana.io").
		WithAPIToken("your-token").
		WithLogger(logger).
		WithDebug(true).
		Build()

	if err != nil {
		log.Fatalf("Failed to build config: %v", err)
	}

	fmt.Printf("✓ Custom logger configured\n")

	// Test logger
	cfg.Logger.Info("This is an info message", "key", "value")
	cfg.Logger.Debug("This is a debug message", "count", 42)
	cfg.Logger.Error("This is an error message", "error", "something went wrong")
}

// CustomLogger implements the config.Logger interface
type CustomLogger struct {
	prefix string
}

func (l *CustomLogger) Debug(msg string, keysAndValues ...interface{}) {
	fmt.Printf("%s [DEBUG] %s %v\n", l.prefix, msg, keysAndValues)
}

func (l *CustomLogger) Info(msg string, keysAndValues ...interface{}) {
	fmt.Printf("%s [INFO] %s %v\n", l.prefix, msg, keysAndValues)
}

func (l *CustomLogger) Warn(msg string, keysAndValues ...interface{}) {
	fmt.Printf("%s [WARN] %s %v\n", l.prefix, msg, keysAndValues)
}

func (l *CustomLogger) Error(msg string, keysAndValues ...interface{}) {
	fmt.Printf("%s [ERROR] %s %v\n", l.prefix, msg, keysAndValues)
}
