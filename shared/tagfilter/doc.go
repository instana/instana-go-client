// Package tagfilter provides tag filter expression functionality for Instana API queries.
//
// Tag filters are used across many Instana APIs to filter resources based on tags,
// entities, and various comparison operators. This package provides a flexible
// way to construct complex filter expressions using logical operators (AND/OR)
// and comparison operators.
//
// Example usage:
//
//	import (
//	    "github.com/instana/instana-go-client/shared/tagfilter"
//	    "github.com/instana/instana-go-client/shared/types"
//	)
//
//	// Create a simple string filter
//	filter := tagfilter.NewStringTagFilter(
//	    tagfilter.TagFilterEntitySource,
//	    "service.name",
//	    types.EqualsOperator,
//	    "my-service",
//	)
//
//	// Create a complex AND expression
//	andFilter := tagfilter.NewLogicalAndTagFilter([]*tagfilter.TagFilter{
//	    tagfilter.NewStringTagFilter(
//	        tagfilter.TagFilterEntitySource,
//	        "service.name",
//	        types.EqualsOperator,
//	        "my-service",
//	    ),
//	    tagfilter.NewStringTagFilter(
//	        tagfilter.TagFilterEntitySource,
//	        "environment",
//	        types.EqualsOperator,
//	        "production",
//	    ),
//	})
package tagfilter

// Made with Bob
