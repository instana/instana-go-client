// Package types provides common type definitions used across multiple Instana API packages.
//
// This package contains shared types such as:
//   - Severity levels (Warning, Critical)
//   - Granularity values for time-based operations
//   - Operators for comparisons and expressions
//   - Aggregation types for metrics
//   - Threshold definitions and operators
//
// These types are used by multiple API endpoint packages and are centralized here
// to avoid duplication and ensure consistency across the client library.
//
// Example usage:
//
//	import "github.com/instana/instana-go-client/shared/types"
//
//	severity := types.SeverityCritical
//	granularity := types.Granularity600000 // 10 minutes
//	operator := types.GreaterThanOperator
//	aggregation := types.MeanAggregation
package types
