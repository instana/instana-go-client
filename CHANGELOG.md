# Changelog

All notable changes to the Instana Go Client will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Comprehensive project documentation
  - API Reference documentation
  - Architecture documentation
  - Contributing guidelines
  - Changelog template

## [0.9.0] - 2026-03-09

### Added
- Configuration system with builder pattern
- Environment variable configuration support
- JSON configuration file support
- Retry mechanism with exponential backoff
- Rate limiting with token bucket algorithm
- Typed error handling system
- Structured logging with sensitive data redaction
- Connection pooling configuration
- Custom headers support
- Comprehensive configuration validation

### Changed
- Migrated from Terraform provider to standalone library
- Improved error messages with detailed context
- Enhanced timeout configuration options

### Fixed
- Configuration validation edge cases
- Rate limiter thread safety issues

## [0.8.0] - 2026-02-15

### Added
- Initial package migration from Terraform provider
- Core REST client implementation
- Support for 28 Instana API resources
- Basic authentication and TLS support

### Changed
- Package structure reorganization
- Improved type safety with generics

## Types of Changes

- `Added` - New features
- `Changed` - Changes in existing functionality
- `Deprecated` - Soon-to-be removed features
- `Removed` - Removed features
- `Fixed` - Bug fixes
- `Security` - Security vulnerability fixes

## Version Format

This project uses [Semantic Versioning](https://semver.org/):

- **MAJOR** version for incompatible API changes
- **MINOR** version for backwards-compatible functionality additions
- **PATCH** version for backwards-compatible bug fixes

## Links

- [Unreleased]: https://github.com/instana/instana-go-client/compare/v0.9.0...HEAD
- [0.9.0]: https://github.com/instana/instana-go-client/compare/v0.8.0...v0.9.0
- [0.8.0]: https://github.com/instana/instana-go-client/releases/tag/v0.8.0