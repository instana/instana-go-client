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
	fmt.Println("=== Instana REST Client Usage Examples ===")
	fmt.Println()

	// Example 1: Using the legacy NewClient (backward compatible)
	fmt.Println("1. Legacy Client (Backward Compatible)")
	fmt.Println("---------------------------------------")
	legacyClient := instana.NewClient(
		"your-api-token",
		"tenant-unit.instana.io",
		false, // skipTlsVerification
	)
	fmt.Printf("Legacy client created: %T\n\n", legacyClient)

	// Example 2: Using NewClientWithConfig with builder pattern
	fmt.Println("2. New Client with Builder Pattern")
	fmt.Println("-----------------------------------")
	conf, err := config.NewConfigBuilder().
		WithBaseURL("https://tenant-unit.instana.io").
		WithAPIToken("your-api-token").
		WithConnectionTimeout(10*time.Second).
		WithRequestTimeout(30*time.Second).
		WithMaxRetryAttempts(5).
		WithRetryInitialDelay(1*time.Second).
		WithRetryMaxDelay(30*time.Second).
		WithRateLimitEnabled(true).
		WithRateLimitRequestsPerSecond(10).
		WithRateLimitBurstCapacity(20).
		WithMaxIdleConnections(100).
		WithMaxConnectionsPerHost(10).
		WithCustomHeader("X-Custom-Header", "custom-value").
		WithDebug(true).
		Build()

	if err != nil {
		log.Fatalf("Failed to build config: %v", err)
	}

	client, err := instana.NewClientWithConfig(conf)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	fmt.Printf("Client created with custom configuration: %T\n\n", client)

	// Example 3: Using environment variables
	fmt.Println("3. Client from Environment Variables")
	fmt.Println("-------------------------------------")
	fmt.Println("Set these environment variables:")
	fmt.Println("  export INSTANA_BASE_URL=https://tenant-unit.instana.io")
	fmt.Println("  export INSTANA_API_TOKEN=your-api-token")
	fmt.Println("  export INSTANA_MAX_RETRY_ATTEMPTS=5")
	fmt.Println("  export INSTANA_RATE_LIMIT_ENABLED=true")
	fmt.Println("  export INSTANA_RATE_LIMIT_RPS=10")

	envConfig, err := config.LoadFromEnv()
	if err != nil {
		fmt.Printf("Note: Environment variables not set, using defaults: %v\n", err)
		envConfig = config.DefaultClientConfig()
		envConfig.BaseURL = "https://tenant-unit.instana.io"
		envConfig.APIToken = "your-api-token"
	}

	envClient, err := instana.NewClientWithConfig(envConfig)
	if err != nil {
		log.Fatalf("Failed to create client from env: %v", err)
	}
	fmt.Printf("Client created from environment: %T\n\n", envClient)

	// Example 4: Making API calls with the new client
	fmt.Println("4. Making API Calls")
	fmt.Println("-------------------")

	// Note: These are example calls - they will fail without valid credentials
	_ = context.Background() // Context for future use

	// GET request
	fmt.Println("GET /api/application-monitoring/applications")
	data, err := client.Get("/api/application-monitoring/applications")
	if err != nil {
		fmt.Printf("Error (expected without valid token): %v\n", err)

		// Check error type
		if config.IsRetryableError(err) {
			fmt.Println("  → This error is retryable")
		}
		if config.IsTemporaryError(err) {
			fmt.Println("  → This error is temporary")
		}
		if statusCode := config.ExtractStatusCode(err); statusCode > 0 {
			fmt.Printf("  → HTTP Status Code: %d\n", statusCode)
		}
	} else {
		fmt.Printf("Success! Received %d bytes\n", len(data))
	}
	fmt.Println()

	// Example 5: Demonstrating retry mechanism
	fmt.Println("5. Retry Mechanism")
	fmt.Println("------------------")
	retryConfig := config.DefaultClientConfig()
	retryConfig.BaseURL = "https://tenant-unit.instana.io"
	retryConfig.APIToken = "your-api-token"
	retryConfig.Retry.MaxAttempts = 3
	retryConfig.Retry.InitialDelay = 500 * time.Millisecond
	retryConfig.Retry.MaxDelay = 5 * time.Second
	retryConfig.Retry.Jitter = true

	retryClient, _ := instana.NewClientWithConfig(retryConfig)

	fmt.Println("Attempting request with retry (will retry on failure)...")
	_, err = retryClient.Get("/api/application-monitoring/applications")
	if err != nil {
		fmt.Printf("Request failed after retries: %v\n", err)
	}
	fmt.Println()

	// Example 6: Rate limiting demonstration
	fmt.Println("6. Rate Limiting")
	fmt.Println("----------------")
	rateLimitConfig := config.DefaultClientConfig()
	rateLimitConfig.BaseURL = "https://tenant-unit.instana.io"
	rateLimitConfig.APIToken = "your-api-token"
	rateLimitConfig.RateLimit.Enabled = true
	rateLimitConfig.RateLimit.RequestsPerSecond = 2 // Only 2 requests per second
	rateLimitConfig.RateLimit.BurstCapacity = 5

	rateLimitClient, _ := instana.NewClientWithConfig(rateLimitConfig)

	fmt.Println("Making 5 rapid requests (rate limited to 2 req/s)...")
	start := time.Now()
	for i := 1; i <= 5; i++ {
		fmt.Printf("  Request %d at %v\n", i, time.Since(start).Round(time.Millisecond))
		_, _ = rateLimitClient.Get("/api/application-monitoring/applications")
	}
	fmt.Printf("Total time: %v (should be ~2 seconds due to rate limiting)\n\n", time.Since(start).Round(time.Millisecond))

	// Example 7: Custom headers
	fmt.Println("7. Custom Headers")
	fmt.Println("-----------------")
	headerConfig := config.DefaultClientConfig()
	headerConfig.BaseURL = "https://tenant-unit.instana.io"
	headerConfig.APIToken = "your-api-token"
	headerConfig.Headers.Custom = map[string]string{
		"X-Request-ID":  "custom-request-id-123",
		"X-Client-Name": "my-application",
		"X-Version":     "1.0.0",
	}

	headerClient, _ := instana.NewClientWithConfig(headerConfig)
	fmt.Println("Client created with custom headers:")
	for key, value := range headerConfig.Headers.Custom {
		fmt.Printf("  %s: %s\n", key, value)
	}
	fmt.Printf("Client: %T\n\n", headerClient)

	// Example 8: Connection pooling
	fmt.Println("8. Connection Pooling")
	fmt.Println("---------------------")
	poolConfig := config.DefaultClientConfig()
	poolConfig.BaseURL = "https://tenant-unit.instana.io"
	poolConfig.APIToken = "your-api-token"
	poolConfig.ConnectionPool.MaxIdleConnections = 200
	poolConfig.ConnectionPool.MaxConnectionsPerHost = 20
	poolConfig.ConnectionPool.KeepAliveDuration = 90 * time.Second

	poolClient, _ := instana.NewClientWithConfig(poolConfig)
	fmt.Printf("Client created with connection pool:\n")
	fmt.Printf("  Max Idle Connections: %d\n", poolConfig.ConnectionPool.MaxIdleConnections)
	fmt.Printf("  Max Connections Per Host: %d\n", poolConfig.ConnectionPool.MaxConnectionsPerHost)
	fmt.Printf("  Keep-Alive Duration: %v\n", poolConfig.ConnectionPool.KeepAliveDuration)
	fmt.Printf("Client: %T\n\n", poolClient)

	// Example 9: Complete production-ready configuration
	fmt.Println("9. Production-Ready Configuration")
	fmt.Println("----------------------------------")
	prodConfig, err := config.NewConfigBuilder().
		WithBaseURL("https://tenant-unit.instana.io").
		WithAPIToken("your-api-token").
		// Timeouts
		WithConnectionTimeout(10*time.Second).
		WithRequestTimeout(30*time.Second).
		WithIdleConnectionTimeout(90*time.Second).
		// Retry
		WithMaxRetryAttempts(5).
		WithRetryInitialDelay(1*time.Second).
		WithRetryMaxDelay(30*time.Second).
		WithRetryBackoffMultiplier(2.0).
		WithRetryJitter(true).
		WithRetryOnTimeout(true).
		WithRetryOnConnectionError(true).
		// Rate Limiting
		WithRateLimitEnabled(true).
		WithRateLimitRequestsPerSecond(10).
		WithRateLimitBurstCapacity(20).
		// Connection Pool
		WithMaxIdleConnections(100).
		WithMaxConnectionsPerHost(10).
		WithKeepAliveDuration(90*time.Second).
		// Headers
		WithCustomHeader("X-Application", "my-app").
		WithCustomHeader("X-Environment", "production").
		// Logging
		WithDebug(false).
		Build()

	if err != nil {
		log.Fatalf("Failed to build production config: %v", err)
	}

	prodClient, err := instana.NewClientWithConfig(prodConfig)
	if err != nil {
		log.Fatalf("Failed to create production client: %v", err)
	}

	fmt.Println("Production client created with:")
	fmt.Println("  ✓ Automatic retries with exponential backoff")
	fmt.Println("  ✓ Rate limiting (10 req/s with burst of 20)")
	fmt.Println("  ✓ Connection pooling (100 idle, 10 per host)")
	fmt.Println("  ✓ Custom headers for tracking")
	fmt.Println("  ✓ Comprehensive timeout configuration")
	fmt.Println("  ✓ Typed error handling")
	fmt.Printf("Client: %T\n\n", prodClient)

	fmt.Println("=== Examples Complete ===")
	fmt.Println("\nKey Features Demonstrated:")
	fmt.Println("1. ✓ Backward compatibility with legacy NewClient()")
	fmt.Println("2. ✓ Flexible configuration with builder pattern")
	fmt.Println("3. ✓ Environment variable loading")
	fmt.Println("4. ✓ Automatic retry with exponential backoff")
	fmt.Println("5. ✓ Rate limiting with token bucket algorithm")
	fmt.Println("6. ✓ Custom headers support")
	fmt.Println("7. ✓ Connection pooling")
	fmt.Println("8. ✓ Typed error handling")
	fmt.Println("9. ✓ Production-ready configuration")
}
