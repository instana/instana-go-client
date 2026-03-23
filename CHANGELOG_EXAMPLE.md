# CHANGELOG Example

This file shows what the CHANGELOG.md will look like after releases.

## Example: First Release (v1.0.0)

After running the first release with `git tag v1.0.0 && git push origin v1.0.0`, the CHANGELOG.md will be automatically updated to:

```markdown
# Changelog

All notable changes to the Instana Go Client will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [v1.0.0](https://github.com/instana/instana-go-client/tree/v1.0.0) (2026-03-19)

**Changes:**

- Initial release of Instana Go Client
- Add configuration system with builder pattern [\#45](https://github.com/instana/instana-go-client/pull/45)
- Add retry mechanism with exponential backoff [\#44](https://github.com/instana/instana-go-client/pull/44)
- Add rate limiting with token bucket algorithm [\#43](https://github.com/instana/instana-go-client/pull/43)
- Add structured logging with sensitive data redaction [\#42](https://github.com/instana/instana-go-client/pull/42)

**Note**: This CHANGELOG is automatically generated during the release process. Release entries will be added here when versions are tagged and released.
```

## Example: Second Release (v1.1.0)

After the second release, the CHANGELOG.md will have both versions:

```markdown
# Changelog

All notable changes to the Instana Go Client will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [v1.1.0](https://github.com/instana/instana-go-client/tree/v1.1.0) (2026-04-15)

[Full Changelog](https://github.com/instana/instana-go-client/compare/v1.0.0...v1.1.0)

**Changes:**

- Add support for custom dashboards [\#50](https://github.com/instana/instana-go-client/pull/50)
- Add support for SLO configurations [\#49](https://github.com/instana/instana-go-client/pull/49)
- Fix timeout handling in REST client [\#48](https://github.com/instana/instana-go-client/pull/48)

## [v1.0.0](https://github.com/instana/instana-go-client/tree/v1.0.0) (2026-03-19)

**Changes:**

- Initial release of Instana Go Client
- Add configuration system with builder pattern [\#45](https://github.com/instana/instana-go-client/pull/45)
- Add retry mechanism with exponential backoff [\#44](https://github.com/instana/instana-go-client/pull/44)
- Add rate limiting with token bucket algorithm [\#43](https://github.com/instana/instana-go-client/pull/43)
- Add structured logging with sensitive data redaction [\#42](https://github.com/instana/instana-go-client/pull/42)

**Note**: This CHANGELOG is automatically generated during the release process. Release entries will be added here when versions are tagged and released.
```

## Key Points

1. **First Release**: No "Full Changelog" comparison link (no previous version to compare)
2. **Subsequent Releases**: Include comparison link to previous version
3. **Automatic Generation**: All entries are generated from commit messages and PR numbers
4. **No Manual Editing**: Developers never need to edit CHANGELOG.md manually
5. **PR Attribution**: PR numbers are automatically extracted and linked (without user names)

## Format Details

- **Version Header**: Links to the release tag on GitHub
- **Date**: Automatically set to release date
- **Full Changelog**: Comparison link between versions (not shown for first release)
- **Changes Section**: Single section with all changes
- **PR Links**: Format `[\#123](PR_URL)` without user attribution
- **Commits without PRs**: Listed without links

## How It Works

1. Developer commits: `git commit -m "Add new feature (#50)"`
2. Create release: `git tag v1.1.0 && git push origin v1.1.0`
3. Workflow automatically:
   - Extracts all commits since last tag
   - Finds PR numbers in commit messages
   - Generates CHANGELOG entry
   - Prepends to CHANGELOG.md
   - Commits and pushes the update