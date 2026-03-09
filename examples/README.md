# Instana Go Client Examples

This directory contains practical examples demonstrating how to use the Instana Go Client library with the proper API methods.

## 📁 Example File

### Usage Examples
**File:** [`usage_examples.go`](usage_examples.go)

Demonstrates 5 comprehensive examples of using the Instana Go Client:

1. **Basic Usage with User Agent** - Using `NewInstanaAPIWithUserAgent()` for simple client creation
2. **Advanced Configuration** - Using `NewInstanaAPIWithConfig()` with comprehensive configuration options
3. **Managing Alerting Channels** - CRUD operations on alerting channels
4. **Application Configurations** - Working with application configs and alerts
5. **API Tokens and Users** - Managing API tokens, users, groups, roles, and teams



## 🚀 Quick Start

### Prerequisites
```bash
go get github.com/instana/instana-go-client
```

## 📚 Available API Methods

The `InstanaAPI` interface provides access to all Instana resources:

- **APITokens()** - Manage API authentication tokens
- **AlertingChannels()** - Manage notification channels
- **AlertingConfigurations()** - Manage alert rules
- **ApplicationAlertConfigs()** - Application alert configurations
- **ApplicationConfigs()** - Application configurations
- **AutomationActions()** - Automation actions
- **AutomationPolicies()** - Automation policies
- **CustomDashboards()** - Custom dashboards
- **CustomEventSpecifications()** - Custom event specifications
- **Groups()** - RBAC groups
- **InfraAlertConfigs()** - Infrastructure alert configurations
- **LogAlertConfigs()** - Log alert configurations
- **MaintenanceWindowConfigs()** - Maintenance windows
- **MobileAlertConfigs()** - Mobile app alert configurations
- **Roles()** - RBAC roles
- **SliConfigs()** - SLI configurations
- **SloAlertConfigs()** - SLO alert configurations
- **SloConfigs()** - SLO configurations
- **SyntheticAlertConfigs()** - Synthetic alert configurations
- **SyntheticTests()** - Synthetic tests
- **Teams()** - RBAC teams
- **Users()** - Users (read-only)
- **WebsiteAlertConfigs()** - Website alert configurations
- **WebsiteMonitoringConfigs()** - Website monitoring configurations

## 💡 Best Practices

1. **Use User Agent**: Always set a meaningful user agent to track client usage
2. **Configure Timeouts**: Set appropriate connection and request timeouts
3. **Enable Retry**: Configure retry mechanism for better reliability
4. **Rate Limiting**: Enable rate limiting to avoid hitting API limits
5. **Connection Pooling**: Configure connection pooling for better performance

## 🆘 Need Help?

- [GitHub Issues](https://github.com/instana/instana-go-client/issues)
- [Main Documentation](../README.md)
