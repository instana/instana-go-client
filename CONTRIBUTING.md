# Contributing to Instana Go Client

Thank you for your interest in contributing to the Instana Go Client! This document provides guidelines and instructions for contributing to the project.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Making Changes](#making-changes)
- [Testing](#testing)
- [Code Style](#code-style)
- [Commit Guidelines](#commit-guidelines)
- [Pull Request Process](#pull-request-process)
- [Reporting Issues](#reporting-issues)

---

## Code of Conduct

This project adheres to a code of conduct that all contributors are expected to follow:

- **Be respectful** - Treat everyone with respect and consideration
- **Be collaborative** - Work together to achieve common goals
- **Be inclusive** - Welcome diverse perspectives and experiences
- **Be professional** - Maintain a professional and constructive tone

---

## Getting Started

### Prerequisites

- **Go 1.20 or higher** - [Download Go](https://golang.org/dl/)
- **Git** - [Install Git](https://git-scm.com/downloads)
- **golangci-lint** - [Install golangci-lint](https://golangci-lint.run/usage/install/)

### Fork and Clone

1. Fork the repository on GitHub
2. Clone your fork locally:

```bash
git clone https://github.com/YOUR-USERNAME/instana-go-client.git
cd instana-go-client
```

3. Add the upstream repository:

```bash
git remote add upstream https://github.com/instana/instana-go-client.git
```

4. Verify your remotes:

```bash
git remote -v
```

---

## Development Setup

### Install Dependencies

```bash
# Download Go module dependencies
go mod download

# Verify dependencies
go mod verify
```

### Install Development Tools

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Install mockgen for generating mocks
go install go.uber.org/mock/mockgen@latest
```

### Verify Setup

```bash
# Build the project
go build ./...

# Run tests
go test ./...

# Run linter
golangci-lint run
```

---

## Making Changes

### Create a Branch

Always create a new branch for your changes:

```bash
# Update your local main branch
git checkout main
git pull upstream main

# Create a new feature branch
git checkout -b feature/your-feature-name

# Or for bug fixes
git checkout -b fix/issue-description
```

### Branch Naming Convention

- `feature/` - New features or enhancements
- `fix/` - Bug fixes
- `docs/` - Documentation changes
- `refactor/` - Code refactoring
- `test/` - Test additions or improvements
- `chore/` - Maintenance tasks

Examples:
- `feature/add-webhook-support`
- `fix/rate-limiter-deadlock`
- `docs/update-api-reference`

### Making Your Changes

1. **Write clear, focused code** - Each change should have a single purpose
2. **Follow Go conventions** - Use idiomatic Go patterns
3. **Add tests** - All new code should include tests
4. **Update documentation** - Keep docs in sync with code changes
5. **Run tests locally** - Ensure all tests pass before committing

---

## Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run tests for a specific package
go test ./config/...

# Run tests with verbose output
go test -v ./...

# Run specific test
go test -run TestConfigBuilder ./config/
```

### Writing Tests

#### Unit Tests

Place unit tests in the same package as the code:

```go
// config/config_test.go
package config

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestDefaultClientConfig(t *testing.T) {
    config := DefaultClientConfig()
    
    assert.NotNil(t, config)
    assert.Equal(t, 30*time.Second, config.Timeout.Connection)
    assert.Equal(t, 3, config.Retry.MaxAttempts)
}
```

#### Table-Driven Tests

Use table-driven tests for multiple scenarios:

```go
func TestConfigValidation(t *testing.T) {
    tests := []struct {
        name    string
        config  *ClientConfig
        wantErr bool
    }{
        {
            name:    "valid config",
            config:  validConfig(),
            wantErr: false,
        },
        {
            name:    "missing base URL",
            config:  configWithoutURL(),
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := ValidateConfig(tt.config)
            if tt.wantErr {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
            }
        })
    }
}
```

#### Integration Tests

Use build tags for integration tests:

```go
//go:build integration
// +build integration

package integration_test

import "testing"

func TestAPIIntegration(t *testing.T) {
    // Integration test code
}
```

Run integration tests:

```bash
go test -tags=integration ./...
```

### Test Coverage Goals

- **Minimum coverage**: 80% for new code
- **Critical paths**: 100% coverage for error handling and retry logic
- **Edge cases**: Test boundary conditions and error scenarios

---

## Code Style

### Go Style Guidelines

Follow the official [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments):

1. **Use `gofmt`** - Format all code with `gofmt`
2. **Use `golint`** - Address all linter warnings
3. **Follow naming conventions**:
   - Exported names: `PascalCase`
   - Unexported names: `camelCase`
   - Acronyms: `HTTPClient`, `APIToken`
4. **Write clear comments**:
   - Package comments for all packages
   - Function comments for exported functions
   - Inline comments for complex logic

### Code Organization

```go
// 1. Package declaration
package config

// 2. Imports (grouped: stdlib, external, internal)
import (
    "context"
    "time"
    
    "github.com/external/package"
    
    "github.com/instana/instana-go-client/shared"
)

// 3. Constants
const (
    DefaultTimeout = 30 * time.Second
)

// 4. Types
type ClientConfig struct {
    // ...
}

// 5. Functions
func NewClientConfig() *ClientConfig {
    // ...
}
```

### Documentation Comments

```go
// ClientConfig holds all configuration options for the Instana API client.
// It provides a flexible way to customize client behavior including timeouts,
// retry logic, rate limiting, and connection pooling.
//
// Example:
//
//	config := DefaultClientConfig()
//	config.BaseURL = "https://tenant.instana.io"
//	config.APIToken = "your-token"
//
type ClientConfig struct {
    // BaseURL is the base URL of the Instana API
    BaseURL string
    
    // APIToken is the API token used for authentication
    APIToken string
}
```

### Error Handling

```go
// Good: Return errors, don't panic
func LoadConfig(path string) (*ClientConfig, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("failed to read config: %w", err)
    }
    // ...
}

// Good: Wrap errors with context
if err != nil {
    return nil, fmt.Errorf("failed to validate config: %w", err)
}

// Good: Use typed errors
return &InstanaError{
    Type:    ErrorTypeValidation,
    Message: "invalid configuration",
    Err:     err,
}
```

### Linting

Run the linter before committing:

```bash
# Run all linters
golangci-lint run

# Run specific linters
golangci-lint run --enable-only=errcheck,gosimple,govet

# Auto-fix issues where possible
golangci-lint run --fix
```

---

## Commit Guidelines

### Commit Message Format

Follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

```
<type>(<scope>): <subject>

<body>

<footer>
```

#### Types

- `feat` - New feature
- `fix` - Bug fix
- `docs` - Documentation changes
- `style` - Code style changes (formatting, etc.)
- `refactor` - Code refactoring
- `test` - Test additions or improvements
- `chore` - Maintenance tasks
- `perf` - Performance improvements

#### Examples

```
feat(config): add support for custom headers

Add WithCustomHeader and WithCustomHeaders methods to ConfigBuilder
to allow users to set custom HTTP headers for all requests.

Closes #123
```

```
fix(retry): prevent infinite retry loop

Add maximum retry attempts check to prevent infinite loops when
retryable errors occur continuously.

Fixes #456
```

```
docs(api): update API reference with new methods

- Add documentation for new rate limiting methods
- Update examples with current API
- Fix typos in configuration section
```

### Commit Best Practices

1. **Make atomic commits** - Each commit should represent a single logical change
2. **Write clear messages** - Explain what and why, not how
3. **Reference issues** - Link to related issues or PRs
4. **Keep commits small** - Easier to review and revert if needed
5. **Test before committing** - Ensure tests pass

---

## Pull Request Process

### Before Submitting

1. **Update your branch** with the latest upstream changes:

```bash
git fetch upstream
git rebase upstream/main
```

2. **Run all checks**:

```bash
# Run tests
go test ./...

# Run linter
golangci-lint run

# Format code
go fmt ./...
```

3. **Update documentation** if needed:
   - Update README.md
   - Update API_REFERENCE.md
   - Update CHANGELOG.md
   - Add/update examples

### Creating a Pull Request

1. **Push your branch** to your fork:

```bash
git push origin feature/your-feature-name
```

2. **Create a pull request** on GitHub:
   - Use a clear, descriptive title
   - Fill out the PR template completely
   - Link related issues
   - Add screenshots for UI changes
   - Request reviews from maintainers

### Pull Request Template

```markdown
## Description
Brief description of the changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing
- [ ] Unit tests added/updated
- [ ] Integration tests added/updated
- [ ] Manual testing performed

## Checklist
- [ ] Code follows project style guidelines
- [ ] Self-review completed
- [ ] Comments added for complex code
- [ ] Documentation updated
- [ ] No new warnings generated
- [ ] Tests pass locally
- [ ] Dependent changes merged

## Related Issues
Closes #(issue number)
```

### Review Process

1. **Automated checks** must pass:
   - Tests
   - Linting
   - Code coverage

2. **Code review** by maintainers:
   - At least one approval required
   - Address all review comments
   - Make requested changes

3. **Final approval** and merge:
   - Squash commits if needed
   - Update CHANGELOG.md
   - Merge to main branch

### After Merge

1. **Delete your branch**:

```bash
git branch -d feature/your-feature-name
git push origin --delete feature/your-feature-name
```

2. **Update your local repository**:

```bash
git checkout main
git pull upstream main
```

---

## Reporting Issues

### Before Creating an Issue

1. **Search existing issues** - Check if the issue already exists
2. **Check documentation** - Ensure it's not a usage question
3. **Verify the bug** - Reproduce the issue consistently
4. **Gather information** - Collect relevant details

### Creating an Issue

Use the appropriate issue template:

#### Bug Report

```markdown
**Describe the bug**
A clear description of what the bug is.

**To Reproduce**
Steps to reproduce the behavior:
1. Create client with '...'
2. Call method '...'
3. See error

**Expected behavior**
What you expected to happen.

**Actual behavior**
What actually happened.

**Environment**
- Go version: [e.g., 1.20]
- Library version: [e.g., v0.9.0]
- OS: [e.g., macOS 13.0]

**Additional context**
Any other relevant information.
```

#### Feature Request

```markdown
**Is your feature request related to a problem?**
A clear description of the problem.

**Describe the solution you'd like**
What you want to happen.

**Describe alternatives you've considered**
Other solutions you've thought about.

**Additional context**
Any other relevant information.
```

---

## Development Workflow

### Typical Workflow

```bash
# 1. Update your fork
git checkout main
git pull upstream main

# 2. Create feature branch
git checkout -b feature/my-feature

# 3. Make changes
# ... edit files ...

# 4. Run tests
go test ./...

# 5. Run linter
golangci-lint run

# 6. Commit changes
git add .
git commit -m "feat: add my feature"

# 7. Push to fork
git push origin feature/my-feature

# 8. Create pull request on GitHub
```

### Keeping Your Fork Updated

```bash
# Fetch upstream changes
git fetch upstream

# Update main branch
git checkout main
git merge upstream/main

# Push to your fork
git push origin main
```

---

## Project Structure

Understanding the project structure helps with contributions:

```
instana-go-client/
├── instana/          # Main package
├── client/           # API client interface
├── config/           # Configuration system
├── api/              # API resource models
├── shared/           # Shared utilities
├── mocks/            # Generated mocks
├── testutils/        # Test utilities
├── examples/         # Usage examples
└── docs/             # Additional documentation
```

---

## Getting Help

### Resources

- **Documentation**: [README.md](README.md), [API_REFERENCE.md](API_REFERENCE.md)
- **Examples**: [examples/](examples/)
- **Issues**: [GitHub Issues](https://github.com/instana/instana-go-client/issues)
- **Discussions**: [GitHub Discussions](https://github.com/instana/instana-go-client/discussions)

### Questions

- **Usage questions**: Create a discussion
- **Bug reports**: Create an issue
- **Feature requests**: Create an issue
- **Security issues**: Email security@instana.com

---

## Recognition

Contributors will be recognized in:

- **CHANGELOG.md** - Listed in release notes
- **README.md** - Contributors section
- **GitHub** - Contributor graph

---

## License

By contributing to this project, you agree that your contributions will be licensed under the [Apache License 2.0](LICENSE).

---

Thank you for contributing to the Instana Go Client! 🎉