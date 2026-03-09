package tagfilter

import (
	"testing"

	"github.com/instana/instana-go-client/shared/types"
)

// Test NewLogicalOrTagFilter
func TestNewLogicalOrTagFilter(t *testing.T) {
	filter1 := NewStringTagFilter(TagFilterEntitySource, "tag1", types.EqualsOperator, "value1")
	filter2 := NewStringTagFilter(TagFilterEntitySource, "tag2", types.EqualsOperator, "value2")
	elements := []*TagFilter{filter1, filter2}

	result := NewLogicalOrTagFilter(elements)

	if result == nil {
		t.Fatal("Expected non-nil result")
	}
	if result.Type != TagFilterExpressionType {
		t.Errorf("Expected type EXPRESSION, got %s", result.Type)
	}
	if result.LogicalOperator == nil {
		t.Fatal("Expected non-nil LogicalOperator")
	}
	if *result.LogicalOperator != types.LogicalOr {
		t.Errorf("Expected LogicalOr operator, got %s", *result.LogicalOperator)
	}
	if len(result.Elements) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(result.Elements))
	}
}

// Test NewLogicalAndTagFilter
func TestNewLogicalAndTagFilter(t *testing.T) {
	filter1 := NewStringTagFilter(TagFilterEntitySource, "tag1", types.EqualsOperator, "value1")
	filter2 := NewStringTagFilter(TagFilterEntitySource, "tag2", types.EqualsOperator, "value2")
	elements := []*TagFilter{filter1, filter2}

	result := NewLogicalAndTagFilter(elements)

	if result == nil {
		t.Fatal("Expected non-nil result")
	}
	if result.Type != TagFilterExpressionType {
		t.Errorf("Expected type EXPRESSION, got %s", result.Type)
	}
	if result.LogicalOperator == nil {
		t.Fatal("Expected non-nil LogicalOperator")
	}
	if *result.LogicalOperator != types.LogicalAnd {
		t.Errorf("Expected LogicalAnd operator, got %s", *result.LogicalOperator)
	}
	if len(result.Elements) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(result.Elements))
	}
}

// Test TagFilterEntities.ToStringSlice
func TestTagFilterEntities_ToStringSlice(t *testing.T) {
	entities := TagFilterEntities{
		TagFilterEntitySource,
		TagFilterEntityDestination,
		TagFilterEntityNotApplicable,
	}

	result := entities.ToStringSlice()

	if len(result) != 3 {
		t.Fatalf("Expected 3 elements, got %d", len(result))
	}
	if result[0] != "SOURCE" {
		t.Errorf("Expected SOURCE, got %s", result[0])
	}
	if result[1] != "DESTINATION" {
		t.Errorf("Expected DESTINATION, got %s", result[1])
	}
	if result[2] != "NOT_APPLICABLE" {
		t.Errorf("Expected NOT_APPLICABLE, got %s", result[2])
	}
}

// Test TagFilterEntities.ToStringSlice with empty slice
func TestTagFilterEntities_ToStringSlice_Empty(t *testing.T) {
	entities := TagFilterEntities{}

	result := entities.ToStringSlice()

	if len(result) != 0 {
		t.Errorf("Expected empty slice, got %d elements", len(result))
	}
}

// Test SupportedTagFilterEntities
func TestSupportedTagFilterEntities(t *testing.T) {
	if len(SupportedTagFilterEntities) != 3 {
		t.Errorf("Expected 3 supported entities, got %d", len(SupportedTagFilterEntities))
	}

	expected := map[TagFilterEntity]bool{
		TagFilterEntitySource:        true,
		TagFilterEntityDestination:   true,
		TagFilterEntityNotApplicable: true,
	}

	for _, entity := range SupportedTagFilterEntities {
		if !expected[entity] {
			t.Errorf("Unexpected entity: %s", entity)
		}
	}
}

// Test NewStringTagFilter
func TestNewStringTagFilter(t *testing.T) {
	entity := TagFilterEntitySource
	name := "tag.name"
	operator := types.EqualsOperator
	value := "test-value"

	result := NewStringTagFilter(entity, name, operator, value)

	if result == nil {
		t.Fatal("Expected non-nil result")
	}
	if result.Type != TagFilterType {
		t.Errorf("Expected type TAG_FILTER, got %s", result.Type)
	}
	if result.Entity == nil || *result.Entity != entity {
		t.Error("Entity not set correctly")
	}
	if result.Name == nil || *result.Name != name {
		t.Error("Name not set correctly")
	}
	if result.Operator == nil || *result.Operator != operator {
		t.Error("Operator not set correctly")
	}
	if result.StringValue == nil || *result.StringValue != value {
		t.Error("StringValue not set correctly")
	}
	if result.Value != value {
		t.Error("Value not set correctly")
	}
}

// Test NewNumberTagFilter
func TestNewNumberTagFilter(t *testing.T) {
	entity := TagFilterEntityDestination
	name := "metric.value"
	operator := types.GreaterThanOperator
	value := int64(100)

	result := NewNumberTagFilter(entity, name, operator, value)

	if result == nil {
		t.Fatal("Expected non-nil result")
	}
	if result.Type != TagFilterType {
		t.Errorf("Expected type TAG_FILTER, got %s", result.Type)
	}
	if result.Entity == nil || *result.Entity != entity {
		t.Error("Entity not set correctly")
	}
	if result.Name == nil || *result.Name != name {
		t.Error("Name not set correctly")
	}
	if result.Operator == nil || *result.Operator != operator {
		t.Error("Operator not set correctly")
	}
	if result.NumberValue == nil || *result.NumberValue != value {
		t.Error("NumberValue not set correctly")
	}
	if result.Value != value {
		t.Error("Value not set correctly")
	}
}

// Test NewTagTagFilter
func TestNewTagTagFilter(t *testing.T) {
	entity := TagFilterEntitySource
	name := "tag.name"
	operator := types.EqualsOperator
	key := "environment"
	value := "production"

	result := NewTagTagFilter(entity, name, operator, key, value)

	if result == nil {
		t.Fatal("Expected non-nil result")
	}
	if result.Type != TagFilterType {
		t.Errorf("Expected type TAG_FILTER, got %s", result.Type)
	}
	if result.Entity == nil || *result.Entity != entity {
		t.Error("Entity not set correctly")
	}
	if result.Name == nil || *result.Name != name {
		t.Error("Name not set correctly")
	}
	if result.Operator == nil || *result.Operator != operator {
		t.Error("Operator not set correctly")
	}
	if result.Key == nil || *result.Key != key {
		t.Error("Key not set correctly")
	}
	if result.Value != value {
		t.Error("Value not set correctly")
	}
	expectedString := "environment=production"
	if result.StringValue == nil || *result.StringValue != expectedString {
		t.Errorf("StringValue not set correctly, expected %s, got %v", expectedString, result.StringValue)
	}
}

// Test NewBooleanTagFilter
func TestNewBooleanTagFilter(t *testing.T) {
	entity := TagFilterEntityNotApplicable
	name := "flag.enabled"
	operator := types.EqualsOperator
	value := true

	result := NewBooleanTagFilter(entity, name, operator, value)

	if result == nil {
		t.Fatal("Expected non-nil result")
	}
	if result.Type != TagFilterType {
		t.Errorf("Expected type TAG_FILTER, got %s", result.Type)
	}
	if result.Entity == nil || *result.Entity != entity {
		t.Error("Entity not set correctly")
	}
	if result.Name == nil || *result.Name != name {
		t.Error("Name not set correctly")
	}
	if result.Operator == nil || *result.Operator != operator {
		t.Error("Operator not set correctly")
	}
	if result.BooleanValue == nil || *result.BooleanValue != value {
		t.Error("BooleanValue not set correctly")
	}
	if result.Value != value {
		t.Error("Value not set correctly")
	}
}

// Test NewBooleanTagFilter with false value
func TestNewBooleanTagFilter_False(t *testing.T) {
	entity := TagFilterEntitySource
	name := "flag.disabled"
	operator := types.NotEqualOperator
	value := false

	result := NewBooleanTagFilter(entity, name, operator, value)

	if result == nil {
		t.Fatal("Expected non-nil result")
	}
	if result.BooleanValue == nil || *result.BooleanValue != false {
		t.Error("BooleanValue should be false")
	}
}

// Test NewUnaryTagFilter
func TestNewUnaryTagFilter(t *testing.T) {
	entity := TagFilterEntitySource
	name := "tag.exists"
	operator := types.IsEmptyOperator

	result := NewUnaryTagFilter(entity, name, operator)

	if result == nil {
		t.Fatal("Expected non-nil result")
	}
	if result.Type != TagFilterType {
		t.Errorf("Expected type TAG_FILTER, got %s", result.Type)
	}
	if result.Entity == nil || *result.Entity != entity {
		t.Error("Entity not set correctly")
	}
	if result.Name == nil || *result.Name != name {
		t.Error("Name not set correctly")
	}
	if result.Operator == nil || *result.Operator != operator {
		t.Error("Operator not set correctly")
	}
	// Unary filters should not have values
	if result.StringValue != nil {
		t.Error("StringValue should be nil for unary filter")
	}
	if result.NumberValue != nil {
		t.Error("NumberValue should be nil for unary filter")
	}
	if result.BooleanValue != nil {
		t.Error("BooleanValue should be nil for unary filter")
	}
}

// Test NewUnaryTagFilterWithTagKey
func TestNewUnaryTagFilterWithTagKey(t *testing.T) {
	entity := TagFilterEntityDestination
	name := "tag.name"
	tagKey := "custom-key"
	operator := types.IsEmptyOperator

	result := NewUnaryTagFilterWithTagKey(entity, name, &tagKey, operator)

	if result == nil {
		t.Fatal("Expected non-nil result")
	}
	if result.Type != TagFilterType {
		t.Errorf("Expected type TAG_FILTER, got %s", result.Type)
	}
	if result.Entity == nil || *result.Entity != entity {
		t.Error("Entity not set correctly")
	}
	if result.Name == nil || *result.Name != name {
		t.Error("Name not set correctly")
	}
	if result.Key == nil || *result.Key != tagKey {
		t.Error("Key not set correctly")
	}
	if result.Operator == nil || *result.Operator != operator {
		t.Error("Operator not set correctly")
	}
}

// Test NewUnaryTagFilterWithTagKey with nil key
func TestNewUnaryTagFilterWithTagKey_NilKey(t *testing.T) {
	entity := TagFilterEntitySource
	name := "tag.name"
	operator := types.IsEmptyOperator

	result := NewUnaryTagFilterWithTagKey(entity, name, nil, operator)

	if result == nil {
		t.Fatal("Expected non-nil result")
	}
	if result.Key != nil {
		t.Error("Key should be nil")
	}
}

// Test TagFilter.GetType
func TestTagFilter_GetType(t *testing.T) {
	filter := &TagFilter{Type: TagFilterType}

	result := filter.GetType()

	if result != TagFilterType {
		t.Errorf("Expected TAG_FILTER, got %s", result)
	}
}

// Test TagFilter.GetType for expression
func TestTagFilter_GetType_Expression(t *testing.T) {
	filter := &TagFilter{Type: TagFilterExpressionType}

	result := filter.GetType()

	if result != TagFilterExpressionType {
		t.Errorf("Expected EXPRESSION, got %s", result)
	}
}

// Test TagFilter.PrependElement
func TestTagFilter_PrependElement(t *testing.T) {
	parent := NewLogicalAndTagFilter([]*TagFilter{})
	child := NewStringTagFilter(TagFilterEntitySource, "tag1", types.EqualsOperator, "value1")

	if len(parent.Elements) != 0 {
		t.Error("Expected empty elements initially")
	}

	parent.PrependElement(child)

	if len(parent.Elements) != 1 {
		t.Errorf("Expected 1 element, got %d", len(parent.Elements))
	}
	if parent.Elements[0] != child {
		t.Error("Element not added correctly")
	}
}

// Test TagFilter.PrependElement multiple times
func TestTagFilter_PrependElement_Multiple(t *testing.T) {
	parent := NewLogicalOrTagFilter([]*TagFilter{})
	child1 := NewStringTagFilter(TagFilterEntitySource, "tag1", types.EqualsOperator, "value1")
	child2 := NewNumberTagFilter(TagFilterEntityDestination, "metric", types.GreaterThanOperator, 50)
	child3 := NewBooleanTagFilter(TagFilterEntityNotApplicable, "flag", types.EqualsOperator, true)

	parent.PrependElement(child1)
	parent.PrependElement(child2)
	parent.PrependElement(child3)

	if len(parent.Elements) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(parent.Elements))
	}
}

// Test complex nested filter structure
func TestComplexNestedFilter(t *testing.T) {
	// Create leaf filters
	filter1 := NewStringTagFilter(TagFilterEntitySource, "service", types.EqualsOperator, "api")
	filter2 := NewStringTagFilter(TagFilterEntitySource, "environment", types.EqualsOperator, "prod")
	filter3 := NewNumberTagFilter(TagFilterEntityDestination, "response_time", types.GreaterThanOperator, 1000)

	// Create AND expression
	andFilter := NewLogicalAndTagFilter([]*TagFilter{filter1, filter2})

	// Create OR expression with AND as child
	orFilter := NewLogicalOrTagFilter([]*TagFilter{andFilter, filter3})

	if orFilter.Type != TagFilterExpressionType {
		t.Error("Root should be expression type")
	}
	if len(orFilter.Elements) != 2 {
		t.Errorf("Expected 2 elements in OR, got %d", len(orFilter.Elements))
	}
	if orFilter.Elements[0].Type != TagFilterExpressionType {
		t.Error("First element should be expression type")
	}
	if len(orFilter.Elements[0].Elements) != 2 {
		t.Errorf("Expected 2 elements in AND, got %d", len(orFilter.Elements[0].Elements))
	}
}

// Test all operators
func TestAllOperators(t *testing.T) {
	operators := []types.ExpressionOperator{
		types.EqualsOperator,
		types.NotEqualOperator,
		types.ContainsOperator,
		types.NotContainOperator,
		types.StartsWithOperator,
		types.EndsWithOperator,
		types.NotStartsWithOperator,
		types.NotEndsWithOperator,
		types.GreaterOrEqualThanOperator,
		types.LessOrEqualThanOperator,
		types.GreaterThanOperator,
		types.LessThanOperator,
		types.IsEmptyOperator,
		types.NotEmptyOperator,
		types.IsBlankOperator,
		types.NotBlankOperator,
	}

	for _, op := range operators {
		t.Run(string(op), func(t *testing.T) {
			filter := NewStringTagFilter(TagFilterEntitySource, "test", op, "value")
			if filter.Operator == nil || *filter.Operator != op {
				t.Errorf("Operator not set correctly for %s", op)
			}
		})
	}
}

// Test all entities
func TestAllEntities(t *testing.T) {
	entities := []TagFilterEntity{
		TagFilterEntitySource,
		TagFilterEntityDestination,
		TagFilterEntityNotApplicable,
	}

	for _, entity := range entities {
		t.Run(string(entity), func(t *testing.T) {
			filter := NewStringTagFilter(entity, "test", types.EqualsOperator, "value")
			if filter.Entity == nil || *filter.Entity != entity {
				t.Errorf("Entity not set correctly for %s", entity)
			}
		})
	}
}
