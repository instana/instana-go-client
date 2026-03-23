# Local Testing Guide for GitHub Actions Workflows

This guide explains how to test the CI and release workflows locally before pushing to GitHub.

## Prerequisites

### 1. Install act

**act** is a tool that runs GitHub Actions locally using Docker.

**macOS**:
```bash
brew install act
```

**Linux**:
```bash
curl https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash
```

**Windows**:
```bash
choco install act-cli
```

Or download from: https://github.com/nektos/act/releases

### 2. Install Docker

act requires Docker to run containers:
- **macOS/Windows**: Install Docker Desktop
- **Linux**: Install Docker Engine

Verify installation:
```bash
docker --version
```

### 3. Verify act Installation

```bash
act --version
```

---

## Testing CI Workflow Locally

### Quick Test (All Jobs)

```bash
# Test the entire CI workflow
act -W .github/workflows/ci.yml
```

### Test Specific Jobs

**Test Linting Only**:
```bash
act -W .github/workflows/ci.yml -j lint
```

**Test Security Scan**:
```bash
act -W .github/workflows/ci.yml -j security
```

**Test Coverage Report**:
```bash
act -W .github/workflows/ci.yml -j coverage-report
```

**Test on Specific Event**:
```bash
# Test as if it's a pull request
act pull_request -W .github/workflows/ci.yml

# Test as if it's a push to main
act push -W .github/workflows/ci.yml
```

### Test Multiple Jobs

```bash
# Test lint and test jobs
act -W .github/workflows/ci.yml -j lint -j test
```

---

## Testing Release Workflow Locally

### Simulate Tag Push

```bash
# Simulate pushing a tag
act -W .github/workflows/release.yml -e .github/workflows/test-events/tag-push.json
```

First, create the test event file:

```bash
mkdir -p .github/workflows/test-events
cat > .github/workflows/test-events/tag-push.json << 'EOF'
{
  "ref": "refs/tags/v1.0.0",
  "repository": {
    "name": "instana-go-client",
    "owner": {
      "login": "instana"
    }
  }
}
EOF
```

### Test Specific Release Jobs

**Test Version Validation**:
```bash
act -W .github/workflows/release.yml -j version-validation -e .github/workflows/test-events/tag-push.json
```

**Test Quality Checks**:
```bash
act -W .github/workflows/release.yml -j quality-checks
```

**Test Security Scan**:
```bash
act -W .github/workflows/release.yml -j security-scan
```

---

## Manual Testing (Without act)

If you prefer not to use act, you can run the commands manually:

### 1. Test Linting

```bash
# Run golangci-lint
golangci-lint run --timeout=5m

# Check formatting
gofmt -l .
```

### 2. Test Security

```bash
# Install govulncheck
go install golang.org/x/vuln/cmd/govulncheck@latest

# Run vulnerability scan
govulncheck ./...

# Run gosec (if installed)
gosec ./...
```

### 3. Test Coverage

```bash
# Run tests with coverage
go test -coverprofile=coverage.out -covermode=atomic ./...

# Check coverage percentage
go tool cover -func=coverage.out | grep total

# Generate HTML report
go tool cover -html=coverage.out -o coverage.html
open coverage.html  # macOS
xdg-open coverage.html  # Linux
```

### 4. Test go.mod Tidiness

```bash
# Tidy go.mod
go mod tidy

# Check for changes
git diff go.mod go.sum
```

### 5. Test Cross-Platform Builds

```bash
# Linux
GOOS=linux GOARCH=amd64 go build ./...
GOOS=linux GOARCH=arm64 go build ./...

# macOS
GOOS=darwin GOARCH=amd64 go build ./...
GOOS=darwin GOARCH=arm64 go build ./...

# Windows
GOOS=windows GOARCH=amd64 go build ./...
```

### 6. Test with Race Detector

```bash
go test -v -race ./...
```

---

## Pre-Commit Testing Script

Create a script to run all checks before committing:

```bash
cat > scripts/pre-commit-check.sh << 'EOF'
#!/bin/bash

set -e

echo "🔍 Running pre-commit checks..."

echo ""
echo "1️⃣ Checking code formatting..."
if [ -n "$(gofmt -l .)" ]; then
  echo "❌ Code is not formatted. Run: gofmt -w ."
  gofmt -l .
  exit 1
fi
echo "✅ Code formatting OK"

echo ""
echo "2️⃣ Running golangci-lint..."
golangci-lint run --timeout=5m
echo "✅ Linting OK"

echo ""
echo "3️⃣ Running tests with race detector..."
go test -v -race ./...
echo "✅ Tests OK"

echo ""
echo "4️⃣ Checking coverage..."
go test -coverprofile=coverage.out -covermode=atomic ./...
COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')
THRESHOLD=80
echo "Coverage: ${COVERAGE}%"
if (( $(echo "$COVERAGE < $THRESHOLD" | bc -l) )); then
  echo "❌ Coverage ${COVERAGE}% is below threshold ${THRESHOLD}%"
  exit 1
fi
echo "✅ Coverage OK"

echo ""
echo "5️⃣ Running security scan..."
govulncheck ./...
echo "✅ Security scan OK"

echo ""
echo "6️⃣ Verifying go.mod and go.sum..."
go mod tidy
if [ -n "$(git diff go.mod go.sum)" ]; then
  echo "❌ go.mod or go.sum is not tidy"
  git diff go.mod go.sum
  exit 1
fi
echo "✅ go.mod and go.sum OK"

echo ""
echo "7️⃣ Testing cross-platform builds..."
GOOS=linux GOARCH=amd64 go build ./...
GOOS=darwin GOARCH=amd64 go build ./...
GOOS=windows GOARCH=amd64 go build ./...
echo "✅ Cross-platform builds OK"

echo ""
echo "✅ All pre-commit checks passed!"
EOF

chmod +x scripts/pre-commit-check.sh
```

Run it before committing:
```bash
./scripts/pre-commit-check.sh
```

---

## Pre-Release Testing Script

Create a script to test release readiness:

```bash
cat > scripts/pre-release-check.sh << 'EOF'
#!/bin/bash

set -e

VERSION=$1

if [ -z "$VERSION" ]; then
  echo "Usage: $0 <version>"
  echo "Example: $0 v1.0.0"
  exit 1
fi

echo "🚀 Running pre-release checks for $VERSION..."

echo ""
echo "1️⃣ Validating version format..."
if ! [[ "$VERSION" =~ ^v[0-9]+\.[0-9]+\.[0-9]+(-[a-zA-Z0-9.]+)?(\+[a-zA-Z0-9.]+)?$ ]]; then
  echo "❌ Invalid version format: $VERSION"
  echo "Expected format: v1.2.3, v1.2.3-beta.1, v1.2.3+build.1"
  exit 1
fi
echo "✅ Version format OK"

echo ""
echo "2️⃣ Checking if tag already exists..."
if git rev-parse "$VERSION" >/dev/null 2>&1; then
  echo "❌ Tag $VERSION already exists"
  exit 1
fi
echo "✅ Tag doesn't exist"

echo ""
echo "3️⃣ Checking working directory..."
if [ -n "$(git status --porcelain)" ]; then
  echo "❌ Working directory is not clean"
  git status
  exit 1
fi
echo "✅ Working directory is clean"

echo ""
echo "4️⃣ Running all quality checks..."
./scripts/pre-commit-check.sh

echo ""
echo "✅ All pre-release checks passed!"
echo ""
echo "Ready to release $VERSION"
echo "Run: git tag $VERSION && git push origin $VERSION"
EOF

chmod +x scripts/pre-release-check.sh
```

Run it before releasing:
```bash
./scripts/pre-release-check.sh v1.0.0
```

---

## act Configuration

Create `.actrc` file for act configuration:

```bash
cat > .actrc << 'EOF'
# Use medium-sized runner image
-P ubuntu-latest=catthehacker/ubuntu:act-latest

# Bind workspace
--bind

# Use GitHub token from environment
--secret GITHUB_TOKEN

# Verbose output
--verbose
EOF
```

---

## Common act Commands

### List Available Jobs

```bash
# List all jobs in CI workflow
act -W .github/workflows/ci.yml -l

# List all jobs in release workflow
act -W .github/workflows/release.yml -l
```

### Dry Run (Don't Execute)

```bash
# See what would run without executing
act -W .github/workflows/ci.yml -n
```

### Run with Secrets

```bash
# Create secrets file
cat > .secrets << 'EOF'
GITHUB_TOKEN=your_token_here
INSTANA_BASE_URL=https://your-instana.com
INSTANA_API_TOKEN=your_api_token
EOF

# Run with secrets
act -W .github/workflows/ci.yml --secret-file .secrets
```

**Important**: Add `.secrets` to `.gitignore`!

### Debug Mode

```bash
# Run with verbose output
act -W .github/workflows/ci.yml -v

# Run with debug logging
act -W .github/workflows/ci.yml --verbose
```

---

## Limitations of Local Testing

### What Works Locally
- ✅ Running tests
- ✅ Linting
- ✅ Building
- ✅ Security scans
- ✅ Coverage checks
- ✅ Most workflow logic

### What Doesn't Work Locally
- ❌ GitHub-specific actions (creating releases, uploading artifacts to GitHub)
- ❌ Codecov uploads (requires GitHub context)
- ❌ Dependency review (requires GitHub API)
- ❌ SARIF uploads (requires GitHub Security tab)
- ❌ Go proxy notifications (requires actual release)

### Workarounds
- Use `--dry-run` or `-n` flag to see what would happen
- Mock GitHub-specific actions
- Test GitHub-specific features in a test branch

---

## Recommended Testing Workflow

### Before Every Commit
```bash
# Quick checks
golangci-lint run
go test -v -race ./...
```

### Before Creating PR
```bash
# Full pre-commit checks
./scripts/pre-commit-check.sh
```

### Before Release
```bash
# Full pre-release checks
./scripts/pre-release-check.sh v1.0.0

# Test release workflow locally (partial)
act -W .github/workflows/release.yml -j version-validation -j quality-checks -j security-scan
```

---

## Troubleshooting

### act Fails to Pull Docker Images

```bash
# Pull images manually
docker pull catthehacker/ubuntu:act-latest
```

### Permission Denied Errors

```bash
# Run with sudo (Linux)
sudo act -W .github/workflows/ci.yml
```

### Out of Memory

```bash
# Increase Docker memory limit in Docker Desktop settings
# Or run fewer jobs at once
act -W .github/workflows/ci.yml -j lint
```

### Slow Performance

```bash
# Use smaller runner image
act -W .github/workflows/ci.yml -P ubuntu-latest=node:16-buster-slim
```

---

## Summary

**Quick Local Testing**:
```bash
# Install act
brew install act  # macOS

# Test CI workflow
act -W .github/workflows/ci.yml

# Test specific job
act -W .github/workflows/ci.yml -j lint
```

**Manual Testing**:
```bash
# Run pre-commit checks
./scripts/pre-commit-check.sh

# Run pre-release checks
./scripts/pre-release-check.sh v1.0.0
```

**Best Practice**:
1. Run manual checks before every commit
2. Use act to test workflow changes
3. Run full pre-release checks before tagging
4. Test in a feature branch first for major changes

This ensures your workflows work correctly before pushing to GitHub!