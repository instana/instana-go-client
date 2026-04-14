# Instana Go Client Library

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://github.com/instana/instana-go-client)

A comprehensive Go client library for the [Instana](https://www.ibm.com/products/instana) API. This library provides a clean, idiomatic Go interface for interacting with Instana's monitoring and observability platform.


## Installation

```bash
go get github.com/instana/instana-go-client
```

## Quick Start

### Configuration Methods

### 1. Simple Client Creation 

```go

    // Create client with basic configuration
    api := instana.NewInstanaAPIWithUserAgent(
        "your-api-token",
        "your-tenant.instana.io",
        false, // skipTlsVerification
        "terraform/6.0.0",
    )
    
    // Use the client...
    channel, err := api.AlertingChannels().GetOne("id")

    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Channel Name: %s\n", channel.Name)

```

#### 2. Builder Pattern 

```go
    config, err := config.NewConfigBuilder().
        WithBaseURL("https://tenant-unit.instana.io").
        WithAPIToken("your-api-token").
        WithConnectionTimeout(45 * time.Second).
        WithRequestTimeout(90 * time.Second).
        WithMaxRetryAttempts(5).
        WithRetryInitialDelay(2 * time.Second).
        WithRateLimitRequestsPerSecond(50).
        WithCustomHeader("X-Custom-Header", "value").
        WithDebug(true).
        Build()

    if err != nil {
        log.Fatal(err)
    }

    api, err := instana.NewInstanaAPIWithConfig(config)
    if err != nil {
        log.Fatal(err)
    }

    channel, err := api.AlertingChannels().GetOne("id")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Channel Name: %s\n", channel.Name)
```


## Examples

See the [examples](examples/) directory for complete examples:


## Support

- **Issues**: [GitHub Issues](https://github.com/instana/instana-go-client/issues)


## Changelog

See [CHANGELOG.md](CHANGELOG.md) for version history and changes.

---
