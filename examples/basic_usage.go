package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/instana/instana-go-client/config"

	"github.com/instana/instana-go-client/instana"
)

// Application represents an Instana application
type Application struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

func main() {
	// Example 1: Basic client creation (legacy compatible)
	fmt.Println("=== Example 1: Basic Client Creation ===")
	basicExample()

	// Example 2: Using builder pattern
	fmt.Println("\n=== Example 2: Builder Pattern ===")
	builderExample()

	// Example 3: Loading from environment
	fmt.Println("\n=== Example 3: Environment Variables ===")
	envExample()

	// Example 4: Error handling
	fmt.Println("\n=== Example 4: Error Handling ===")
	errorHandlingExample()
}

// basicExample demonstrates basic client creation
func basicExample() {
	// Create client with minimal configuration
	client := instana.NewClient(
		"your-api-token",
		"your-tenant.instana.io",
		false, // skipTlsVerification
	)

	// Make a simple GET request
	data, err := client.Get("/api/application-monitoring/applications")
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	// Parse response
	var apps []Application
	if err := json.Unmarshal(data, &apps); err != nil {
		log.Printf("Failed to parse response: %v\n", err)
		return
	}

	fmt.Printf("Found %d applications\n", len(apps))
	for _, app := range apps {
		fmt.Printf("  - %s (ID: %s)\n", app.Label, app.ID)
	}
}

// builderExample demonstrates using the builder pattern
func builderExample() {
	// Create configuration using builder
	config, err := config.NewConfigBuilder().
		WithBaseURL("https://your-tenant.instana.io").
		WithAPIToken("your-api-token").
		WithConnectionTimeout(30 * time.Second).
		WithRequestTimeout(60 * time.Second).
		WithMaxRetryAttempts(3).
		WithDebug(true).
		Build()

	if err != nil {
		log.Printf("Failed to create config: %v\n", err)
		return
	}

	// Create client with configuration
	client, err := instana.NewClientWithConfig(config)
	if err != nil {
		log.Printf("Failed to create client: %v\n", err)
		return
	}

	// Use the client
	data, err := client.Get("/api/application-monitoring/applications")
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Retrieved %d bytes of data\n", len(data))
}

// envExample demonstrates loading configuration from environment variables
func envExample() {
	// Set environment variables first:
	// export INSTANA_BASE_URL="https://your-tenant.instana.io"
	// export INSTANA_API_TOKEN="your-api-token"
	// export INSTANA_DEBUG="true"

	// Load configuration from environment
	config, err := config.LoadFromEnv()
	if err != nil {
		log.Printf("Failed to load config from environment: %v\n", err)
		return
	}

	// Create client
	client, err := instana.NewClientWithConfig(config)
	if err != nil {
		log.Printf("Failed to create client: %v\n", err)
		return
	}

	// Use the client
	data, err := client.Get("/api/application-monitoring/applications")
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Retrieved %d bytes of data\n", len(data))
}

// errorHandlingExample demonstrates enhanced error handling
func errorHandlingExample() {
	conf, _ := config.NewConfigBuilder().
		WithBaseURL("https://your-tenant.instana.io").
		WithAPIToken("invalid-token"). // Intentionally invalid
		Build()

	client, _ := instana.NewClientWithConfig(conf)

	// Make request that will fail
	_, err := client.Get("/api/application-monitoring/applications")
	if err != nil {
		// Check for specific error types
		var instanaErr *config.InstanaError
		if errors.As(err, &instanaErr) {
			switch instanaErr.Type {
			case config.ErrorTypeAuthentication:
				fmt.Printf("Authentication failed: %s\n", instanaErr.Message)
			case config.ErrorTypeRateLimit:
				fmt.Printf("Rate limit exceeded: %s\n", instanaErr.Message)
			case config.ErrorTypeAPI:
				fmt.Printf("API error: %s (status: %d)\n", instanaErr.Message, instanaErr.StatusCode)
			case config.ErrorTypeTimeout:
				fmt.Printf("Request timed out: %s\n", instanaErr.Message)
			case config.ErrorTypeNetwork:
				fmt.Printf("Network error: %s\n", instanaErr.Message)
			default:
				fmt.Printf("Error: %s (status: %d)\n", instanaErr.Message, instanaErr.StatusCode)
			}

			// Check if error is retryable
			if instanaErr.IsRetryable() {
				fmt.Println("This error can be retried")
			}

			// Check if error is temporary
			if instanaErr.IsTemporary() {
				fmt.Println("This is a temporary error")
			}
		} else {
			fmt.Printf("Unexpected error: %v\n", err)
		}
	}
}

// Made with Bob
