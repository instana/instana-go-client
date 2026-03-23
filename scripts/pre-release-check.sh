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
echo "4️⃣ Checking current branch..."
CURRENT_BRANCH=$(git branch --show-current)
if [ "$CURRENT_BRANCH" != "main" ]; then
  echo "⚠️  Warning: You are on branch '$CURRENT_BRANCH', not 'main'"
  read -p "Continue anyway? (y/N) " -n 1 -r
  echo
  if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    exit 1
  fi
fi
echo "✅ Branch check OK"

echo ""
echo "5️⃣ Running all quality checks..."
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
if [ -f "$SCRIPT_DIR/pre-commit-check.sh" ]; then
  "$SCRIPT_DIR/pre-commit-check.sh"
else
  echo "⚠️  pre-commit-check.sh not found, running basic checks..."
  go test -v -race ./...
  go build -v ./...
fi

echo ""
echo "✅ All pre-release checks passed!"
echo ""
echo "Ready to release $VERSION"
echo ""
echo "Next steps:"
echo "  1. Create tag:  git tag $VERSION"
echo "  2. Push tag:    git push origin $VERSION"
echo "  3. Monitor:     https://github.com/instana/instana-go-client/actions"

# Made with Bob
