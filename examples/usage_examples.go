//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/instana/instana-go-client/api"
	"github.com/instana/instana-go-client/client"
	"github.com/instana/instana-go-client/config"
)

// This file demonstrates 5 practical examples of using the Instana Go Client
// with the proper API methods from the InstanaAPI interface.

func main() {
	fmt.Println("=== Instana Go Client Usage Examples ===\n")

	// Example 1: Basic API usage with NewInstanaAPIWithUserAgent
	example1_BasicUsageWithUserAgent()

	// Example 2: Advanced configuration with NewInstanaAPIWithConfig
	example2_AdvancedConfiguration()

	// Example 3: Managing Alerting Channels
	example3_ManagingAlertingChannels()

	// Example 4: Working with Application Configurations
	example4_ApplicationConfigurations()

	// Example 5: Managing API Tokens and Users
	example5_APITokensAndUsers()
}

// Example 1: Basic API usage with NewInstanaAPIWithUserAgent
func example1_BasicUsageWithUserAgent() {
	fmt.Println("=== Example 1: Basic Usage with User Agent ===")

	// Create API client with user agent (useful for tracking client usage)
	client := client.NewInstanaAPIWithUserAgent(
		"your-api-token",
		"tenant-unit.instana.io",
		false, // skipTlsVerification
		"MyApp/1.0.0",
	)

	// Get all API tokens
	tokens, err := client.APITokens().GetAll()
	if err != nil {
		log.Printf("Error getting API tokens: %v\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d API tokens\n", len(*tokens))

	// Get all alerting channels
	channels, err := client.AlertingChannels().GetAll()
	if err != nil {
		log.Printf("Error getting alerting channels: %v\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d alerting channels\n", len(*channels))

	// Get all custom dashboards
	dashboards, err := client.CustomDashboards().GetAll()
	if err != nil {
		log.Printf("Error getting dashboards: %v\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d custom dashboards\n\n", len(*dashboards))
}

// Example 2: Advanced configuration with NewInstanaAPIWithConfig
func example2_AdvancedConfiguration() {
	fmt.Println("=== Example 2: Advanced Configuration ===")

	// Build comprehensive configuration
	cfg, err := config.NewConfigBuilder().
		WithBaseURL("https://tenant-unit.instana.io").
		WithAPIToken("your-api-token").
		WithUserAgent("MyApp/2.0.0").
		WithConnectionTimeout(45*time.Second).
		WithRequestTimeout(90*time.Second).
		WithMaxRetryAttempts(5).
		WithRetryInitialDelay(2*time.Second).
		WithRateLimitEnabled(true).
		WithRateLimitRequestsPerSecond(50).
		WithCustomHeader("X-Request-ID", "unique-id-123").
		WithDebug(true).
		Build()

	if err != nil {
		log.Fatalf("Failed to build config: %v", err)
	}

	// Create API client with configuration
	client, err := client.NewInstanaAPIWithConfig(cfg)
	if err != nil {
		log.Fatalf("Failed to create API client: %v", err)
	}

	fmt.Printf("✓ Client configured successfully\n")
	fmt.Printf("  Base URL: %s\n", cfg.BaseURL)
	fmt.Printf("  Connection Timeout: %s\n", cfg.Timeout.Connection)
	fmt.Printf("  Max Retry Attempts: %d\n", cfg.Retry.MaxAttempts)
	fmt.Printf("  Rate Limit: %d req/s\n", cfg.RateLimit.RequestsPerSecond)

	// Use the configured client to get maintenance windows
	windows, err := client.MaintenanceWindowConfigs().GetAll()
	if err != nil {
		log.Printf("Error getting maintenance windows: %v\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d maintenance windows\n\n", len(*windows))
}

// Example 3: Managing Alerting Channels
func example3_ManagingAlertingChannels() {
	fmt.Println("=== Example 3: Managing Alerting Channels ===")

	client := client.NewInstanaAPIWithUserAgent(
		"your-api-token",
		"tenant-unit.instana.io",
		false,
		"AlertingManager/1.0.0",
	)

	// Create a new email alerting channel
	newChannel := &api.AlertingChannel{
		Name:   "DevOps Team Email",
		Kind:   api.EmailChannelType,
		Emails: []string{"devops@example.com", "alerts@example.com"},
	}

	createdChannel, err := client.AlertingChannels().Create(newChannel)
	if err != nil {
		log.Printf("Error creating alerting channel: %v\n", err)
		return
	}
	fmt.Printf("✓ Created alerting channel: %s (ID: %s)\n", createdChannel.Name, createdChannel.ID)

	// Get a specific alerting channel by ID
	channel, err := client.AlertingChannels().GetOne(createdChannel.ID)
	if err != nil {
		log.Printf("Error getting alerting channel: %v\n", err)
		return
	}
	fmt.Printf("✓ Retrieved channel: %s with %d email(s)\n", channel.Name, len(channel.Emails))

	// Update the alerting channel
	channel.Emails = append(channel.Emails, "oncall@example.com")
	updatedChannel, err := client.AlertingChannels().Update(channel)
	if err != nil {
		log.Printf("Error updating alerting channel: %v\n", err)
		return
	}
	fmt.Printf("✓ Updated channel now has %d email(s)\n", len(updatedChannel.Emails))

	// Delete the alerting channel
	err = client.AlertingChannels().DeleteByID(createdChannel.ID)
	if err != nil {
		log.Printf("Error deleting alerting channel: %v\n", err)
		return
	}
	fmt.Printf("✓ Deleted alerting channel: %s\n\n", createdChannel.ID)
}

// Example 4: Working with Application Configurations
func example4_ApplicationConfigurations() {
	fmt.Println("=== Example 4: Application Configurations ===")

	client := client.NewInstanaAPIWithUserAgent(
		"your-api-token",
		"tenant-unit.instana.io",
		false,
		"AppConfigManager/1.0.0",
	)

	// Get all application configurations
	appConfigs, err := client.ApplicationConfigs().GetAll()
	if err != nil {
		log.Printf("Error getting application configs: %v\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d application configurations\n", len(*appConfigs))

	// List application configurations with details
	for i, appConfig := range *appConfigs {
		if i < 3 { // Show first 3 only
			fmt.Printf("  - %s (ID: %s)\n", appConfig.Label, appConfig.ID)
		}
	}

	// Get all application alert configurations
	appAlerts, err := client.ApplicationAlertConfigs().GetAll()
	if err != nil {
		log.Printf("Error getting application alerts: %v\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d application alert configurations\n", len(*appAlerts))

	// Get global application alert configurations
	globalAlerts, err := client.GlobalApplicationAlertConfigs().GetAll()
	if err != nil {
		log.Printf("Error getting global application alerts: %v\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d global application alert configurations\n\n", len(*globalAlerts))
}

// Example 5: Managing API Tokens and Users
func example5_APITokensAndUsers() {
	fmt.Println("=== Example 5: API Tokens and Users ===")

	client := client.NewInstanaAPIWithUserAgent(
		"your-api-token",
		"tenant-unit.instana.io",
		false,
		"TokenManager/1.0.0",
	)

	// Get all API tokens
	tokens, err := client.APITokens().GetAll()
	if err != nil {
		log.Printf("Error getting API tokens: %v\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d API tokens\n", len(*tokens))

	// List tokens with details
	for i, token := range *tokens {
		if i < 3 { // Show first 3 only
			fmt.Printf("  - %s (ID: %s)\n", token.Name, token.ID)
		}
	}

	// Get all users (read-only)
	users, err := client.Users().GetAll()
	if err != nil {
		log.Printf("Error getting users: %v\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d users\n", len(*users))

	// Get all RBAC groups
	groups, err := client.Groups().GetAll()
	if err != nil {
		log.Printf("Error getting groups: %v\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d RBAC groups\n", len(*groups))

	// Get all RBAC roles
	roles, err := client.Roles().GetAll()
	if err != nil {
		log.Printf("Error getting roles: %v\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d RBAC roles\n", len(*roles))

	// Get all teams
	teams, err := client.Teams().GetAll()
	if err != nil {
		log.Printf("Error getting teams: %v\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d teams\n\n", len(*teams))
}

// Additional helper function to demonstrate error handling
func demonstrateErrorHandling() {
	fmt.Println("== Error Handling Example ===")

	// Create client with invalid credentials to demonstrate error handling
	api := client.NewInstanaAPI("invalid-token", "tenant.instana.io", false)

	// Attempt operation that will fail
	_, err := api.APITokens().GetAll()
	if err != nil {
		// Check if it's an Instana error
		if instanaErr, ok := err.(*config.InstanaError); ok {
			fmt.Printf("✓ Caught InstanaError:\n")
			fmt.Printf("  Type: %s\n", instanaErr.Type)
			fmt.Printf("  Message: %s\n", instanaErr.Message)
			fmt.Printf("  Status Code: %d\n", instanaErr.StatusCode)
			fmt.Printf("  Retryable: %v\n", instanaErr.IsRetryable())
			fmt.Printf("  Temporary: %v\n", instanaErr.IsTemporary())

			// Handle specific error types
			switch instanaErr.Type {
			case config.ErrorTypeAuthentication:
				fmt.Println("  → Action: Check API token and permissions")
			case config.ErrorTypeRateLimit:
				fmt.Println("  → Action: Wait before retrying")
			case config.ErrorTypeNetwork:
				fmt.Println("  → Action: Check network connectivity")
			case config.ErrorTypeTimeout:
				fmt.Println("  → Action: Increase timeout or retry")
			}
		}
	}
	fmt.Println()
}
