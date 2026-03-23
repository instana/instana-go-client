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
if ! command -v golangci-lint &> /dev/null; then
  echo "⚠️  golangci-lint not found. Install it from: https://golangci-lint.run/usage/install/"
  echo "Skipping linting..."
else
  golangci-lint run --timeout=5m
  echo "✅ Linting OK"
fi

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
if command -v bc &> /dev/null; then
  if (( $(echo "$COVERAGE < $THRESHOLD" | bc -l) )); then
    echo "❌ Coverage ${COVERAGE}% is below threshold ${THRESHOLD}%"
    exit 1
  fi
else
  # Fallback for systems without bc
  COVERAGE_INT=${COVERAGE%.*}
  if [ "$COVERAGE_INT" -lt "$THRESHOLD" ]; then
    echo "❌ Coverage ${COVERAGE}% is below threshold ${THRESHOLD}%"
    exit 1
  fi
fi
echo "✅ Coverage OK"

echo ""
echo "5️⃣ Running security scan..."
if ! command -v govulncheck &> /dev/null; then
  echo "⚠️  govulncheck not found. Installing..."
  go install golang.org/x/vuln/cmd/govulncheck@latest
fi
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

# Made with Bob
