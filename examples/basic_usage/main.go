//go:build ignore
// +build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/instana/instana-go-client/config"
	"github.com/instana/instana-go-client/instana"
)

func main() {
	// Example 1: Using builder pattern with custom configuration
	fmt.Println("=== Example 1: Builder Pattern ===")
	conf, err := config.NewConfigBuilder().
		WithBaseURL("https://tenant-unit.instana.io").
		WithAPIToken("your-api-token-here").
		WithConnectionTimeout(45*time.Second).
		WithRequestTimeout(90*time.Second).
		WithMaxRetryAttempts(5).
		WithRetryInitialDelay(2*time.Second).
		WithRateLimitRequestsPerSecond(50).
		WithCustomHeader("X-Custom-Header", "custom-value").
		WithDebug(true).
		Build()

	if err != nil {
		log.Fatalf("Failed to build config: %v", err)
	}

	fmt.Printf("Configuration created successfully!\n")
	fmt.Printf("  Base URL: %s\n", conf.BaseURL)
	fmt.Printf("  Connection Timeout: %s\n", conf.Timeout.Connection)
	fmt.Printf("  Max Retry Attempts: %d\n", conf.Retry.MaxAttempts)
	fmt.Printf("  Rate Limit: %d req/s\n", conf.RateLimit.RequestsPerSecond)

	// Create API client
	api, err := instana.NewInstanaAPIWithConfig(conf)
	if err != nil {
		log.Fatalf("Failed to create API: %v", err)
	}
	fmt.Printf("✓ API client created successfully\n")

	// Example 2: Using default configuration
	fmt.Println("\n=== Example 2: Default Configuration ===")
	defaultConfig := config.DefaultClientConfig()
	defaultConfig.BaseURL = "https://tenant-unit.instana.io"
	defaultConfig.APIToken = "your-api-token-here"

	if err := defaultConfig.Validate(); err != nil {
		log.Fatalf("Validation failed: %v", err)
	}

	fmt.Printf("Default configuration:\n")
	fmt.Printf("  Connection Timeout: %s\n", defaultConfig.Timeout.Connection)
	fmt.Printf("  Request Timeout: %s\n", defaultConfig.Timeout.Request)
	fmt.Printf("  Max Retry Attempts: %d\n", defaultConfig.Retry.MaxAttempts)
	fmt.Printf("  Retry Initial Delay: %s\n", defaultConfig.Retry.InitialDelay)
	fmt.Printf("  Retry Max Delay: %s\n", defaultConfig.Retry.MaxDelay)
	fmt.Printf("  Batch Size: %d\n", defaultConfig.Batch.Size)
	fmt.Printf("  Rate Limit: %d req/s\n", defaultConfig.RateLimit.RequestsPerSecond)

	// Example 3: Demonstrating retry mechanism
	fmt.Println("\n=== Example 3: Retry Mechanism ===")
	retryConfig := config.DefaultRetryConfig()
	retryer := config.NewRetryer(retryConfig, config.NewDefaultLogger(config.ClientLogLevelInfo))

	attemptCount := 0
	err = retryer.Do(context.Background(), func() error {
		attemptCount++
		fmt.Printf("  Attempt %d...\n", attemptCount)
		if attemptCount < 3 {
			return config.NetworkError("simulated network error", nil)
		}
		return nil
	})

	if err != nil {
		log.Printf("Operation failed after retries: %v", err)
	} else {
		fmt.Printf("  Operation succeeded after %d attempts!\n", attemptCount)
	}

	// Example 4: Demonstrating rate limiter
	fmt.Println("\n=== Example 4: Rate Limiter ===")
	rateLimitConfig := config.RateLimitConfig{
		Enabled:           true,
		RequestsPerSecond: 5,
		BurstCapacity:     10,
		WaitForToken:      true,
	}
	rateLimiter := config.NewRateLimiter(rateLimitConfig, config.NewDefaultLogger(config.ClientLogLevelInfo))
	defer rateLimiter.Stop()

	fmt.Printf("Making 10 requests with rate limit of 5 req/s...\n")
	start := time.Now()
	for i := 1; i <= 10; i++ {
		if err := rateLimiter.Wait(context.Background()); err != nil {
			log.Printf("Rate limiter error: %v", err)
			break
		}
		fmt.Printf("  Request %d sent (elapsed: %s)\n", i, time.Since(start).Round(time.Millisecond))
	}

	// Example 5: Error handling
	fmt.Println("\n=== Example 5: Error Handling ===")

	// Network error (retryable)
	netErr := config.NetworkError("connection refused", nil)
	fmt.Printf("Network Error:\n")
	fmt.Printf("  Message: %s\n", netErr.Error())
	fmt.Printf("  Retryable: %v\n", netErr.IsRetryable())
	fmt.Printf("  Temporary: %v\n", netErr.IsTemporary())

	// API error (may be retryable based on status code)
	apiErr := config.APIError(503, "service unavailable", nil)
	fmt.Printf("\nAPI Error (503):\n")
	fmt.Printf("  Message: %s\n", apiErr.Error())
	fmt.Printf("  Retryable: %v\n", apiErr.IsRetryable())
	fmt.Printf("  Status Code: %d\n", config.ExtractStatusCode(apiErr))

	// Validation error (not retryable)
	valErr := config.NewValidationError("invalid input", nil)
	fmt.Printf("\nValidation Error:\n")
	fmt.Printf("  Message: %s\n", valErr.Error())
	fmt.Printf("  Retryable: %v\n", valErr.IsRetryable())

	// Example 6: Using the API client
	fmt.Println("\n=== Example 6: Using API Client ===")

	// Note: This will fail without valid credentials
	tokens, err := api.APITokens().GetAll()
	if err != nil {
		// Proper error handling
		if instanaErr, ok := err.(*config.InstanaError); ok {
			fmt.Printf("Error Type: %s\n", instanaErr.Type)
			fmt.Printf("Retryable: %v\n", instanaErr.IsRetryable())
			fmt.Printf("Status Code: %d\n", instanaErr.StatusCode)
		}
		fmt.Printf("Note: Expected error without valid credentials: %v\n", err)
	} else {
		fmt.Printf("✓ Retrieved %d tokens\n", len(*tokens))
	}

	fmt.Println("\n=== Examples Complete ===")
}
