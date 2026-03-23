// Package instana provides a comprehensive Go client library for the Instana API.
//
// This package serves as the main entry point for creating and configuring Instana API clients.
// It provides both simple and advanced client creation methods, along with a flexible configuration
// system that supports multiple configuration sources.
//
// # Quick Start
//
// The simplest way to create a client:
//
//	api := instana.NewInstanaAPI("your-api-token", "tenant.instana.io", false)
//	tokens, err := api.APITokens().GetAll()
//
// # Configuration
//
// The package supports multiple configuration methods:
//
// 1. Builder Pattern (Recommended):
//
//	config, err := instana.NewConfigBuilder().
//	    WithBaseURL("https://tenant.instana.io").
//	    WithAPIToken("your-token").
//	    WithMaxRetryAttempts(5).
//	    WithRateLimitRequestsPerSecond(100).
//	    Build()
//
//	api, err := instana.NewInstanaAPIWithConfig(config)
//
// 2. Environment Variables:
//
//	// Set: INSTANA_BASE_URL, INSTANA_API_TOKEN, etc.
//	config, err := instana.LoadFromEnv()
//	api, err := instana.NewInstanaAPIWithConfig(config)
//
// 3. JSON Configuration File:
//
//	config, err := instana.LoadFromJSON("config.json")
//	api, err := instana.NewInstanaAPIWithConfig(config)
//
// # Features
//
// The client provides production-ready features including:
//
//   - Automatic retry with exponential backoff
//   - Rate limiting with token bucket algorithm
//   - Connection pooling and keep-alive
//   - Typed error handling
//   - Structured logging with sensitive data redaction
//   - Context support for cancellation
//   - Thread-safe operations
//
// # API Resources
//
// The client provides access to all Instana API resources through a unified interface:
//
//	// Application Monitoring
//	apps := api.ApplicationConfigs()
//	alerts := api.ApplicationAlertConfigs()
//
//	// Infrastructure Monitoring
//	infraAlerts := api.InfraAlertConfigs()
//	agents := api.HostAgents()
//
//	// Synthetic Monitoring
//	tests := api.SyntheticTests()
//	locations := api.SyntheticLocations()
//
//	// SLO/SLI Management
//	slos := api.SloConfigs()
//	slis := api.SliConfigs()
//
//	// Access Control
//	tokens := api.APITokens()
//	users := api.Users()
//	groups := api.Groups()
//	roles := api.Roles()
//	teams := api.Teams()
//
// # Error Handling
//
// The package provides typed errors for better error handling:
//
//	data, err := api.APITokens().GetAll()
//	if err != nil {
//	    if instanaErr, ok := err.(*config.InstanaError); ok {
//	        switch instanaErr.Type {
//	        case config.ErrorTypeAuthentication:
//	            // Handle authentication error
//	        case config.ErrorTypeRateLimit:
//	            // Handle rate limit error
//	        case config.ErrorTypeNetwork:
//	            // Handle network error
//	        }
//	    }
//	}
//
// # Advanced Usage
//
// For advanced use cases, you can customize various aspects:
//
//	config, err := instana.NewConfigBuilder().
//	    // Timeouts
//	    WithConnectionTimeout(30 * time.Second).
//	    WithRequestTimeout(60 * time.Second).
//	    // Retry
//	    WithMaxRetryAttempts(5).
//	    WithRetryInitialDelay(1 * time.Second).
//	    WithRetryMaxDelay(30 * time.Second).
//	    // Rate Limiting
//	    WithRateLimitEnabled(true).
//	    WithRateLimitRequestsPerSecond(100).
//	    WithRateLimitBurstCapacity(200).
//	    // Connection Pooling
//	    WithMaxIdleConnections(100).
//	    WithMaxConnectionsPerHost(10).
//	    WithKeepAliveDuration(30 * time.Second).
//	    // Custom Headers
//	    WithCustomHeader("X-Request-ID", "unique-id").
//	    // Logging
//	    WithDebug(true).
//	    Build()
//
// # Thread Safety
//
// All client operations are thread-safe and can be used concurrently:
//
//	var wg sync.WaitGroup
//	for i := 0; i < 10; i++ {
//	    wg.Add(1)
//	    go func() {
//	        defer wg.Done()
//	        tokens, _ := api.APITokens().GetAll()
//	        // Process tokens
//	    }()
//	}
//	wg.Wait()
//
// # REST Client
//
// For low-level HTTP operations, you can use the REST client directly:
//
//	restClient := instana.NewClient("api-token", "tenant.instana.io", false)
//	data, err := restClient.Get("/api/application-monitoring/applications")
//
// # See Also
//
// For more information, see:
//   - Package config: Configuration management and validation
//   - Package client: High-level API interface and resource clients
//   - Package api: Data models for all API resources
//   - GitHub: https://github.com/instana/instana-go-client
//   - Documentation: https://pkg.go.dev/github.com/instana/instana-go-client
package instana
