package api_test

import (
	"encoding/json"
	"testing"

	. "github.com/instana/instana-go-client/api"
	"github.com/instana/instana-go-client/shared/tagfilter"
	"github.com/instana/instana-go-client/shared/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBusinessPerspectiveResourcePath(t *testing.T) {
	expected := "/api/business-monitoring/business-perspectives"
	assert.Equal(t, expected, BusinessPerspectiveResourcePath)
}

func TestBusinessPerspectiveGetIDForResourcePath(t *testing.T) {
	testID := "test-biz-perspective-123"
	bp := &BusinessPerspectiveConfig{
		ID:          testID,
		Name:        "Test BP",
		Description: "Test Description",
	}

	assert.Equal(t, testID, bp.GetIDForResourcePath())
}

func TestBusinessPerspectiveStructure(t *testing.T) {
	tagFilter := tagfilter.NewStringTagFilter(
		tagfilter.TagFilterEntityDestination,
		"service.name",
		types.EqualsOperator,
		"payment-service",
	)

	bp := BusinessPerspectiveConfig{
		ID:                  "bp-id-1",
		Name:                "Payment Flow",
		Description:         "Monitors payment transactions",
		TagFilterExpression: tagFilter,
		RbacTags: []RbacTag{
			{ID: "team-1-id", DisplayName: "Team One"},
		},
	}

	assert.Equal(t, "bp-id-1", bp.ID)
	assert.Equal(t, "Payment Flow", bp.Name)
	assert.Equal(t, "Monitors payment transactions", bp.Description)
	require.NotNil(t, bp.TagFilterExpression)
	assert.Equal(t, tagfilter.TagFilterType, bp.TagFilterExpression.Type)
	require.Len(t, bp.RbacTags, 1)
	assert.Equal(t, "team-1-id", bp.RbacTags[0].ID)
	assert.Equal(t, "Team One", bp.RbacTags[0].DisplayName)
}

func TestBusinessPerspectiveSerialization(t *testing.T) {
	tagFilter := tagfilter.NewStringTagFilter(
		tagfilter.TagFilterEntityDestination,
		"service.name",
		types.EqualsOperator,
		"payment-service",
	)

	bp := BusinessPerspectiveConfig{
		ID:                  "bp-id-1",
		Name:                "Payment Flow",
		Description:         "Monitors payment transactions",
		TagFilterExpression: tagFilter,
		RbacTags: []RbacTag{
			{ID: "team-1", DisplayName: "Team 1"},
		},
	}

	data, err := json.Marshal(bp)
	require.NoError(t, err)

	var unmarshaled BusinessPerspectiveConfig
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err)

	assert.Equal(t, bp.ID, unmarshaled.ID)
	assert.Equal(t, bp.Name, unmarshaled.Name)
	assert.Equal(t, bp.Description, unmarshaled.Description)
	require.NotNil(t, unmarshaled.TagFilterExpression)
	require.Len(t, unmarshaled.RbacTags, 1)
	assert.Equal(t, "team-1", unmarshaled.RbacTags[0].ID)
	assert.Equal(t, "Team 1", unmarshaled.RbacTags[0].DisplayName)
}

func TestBusinessPerspectiveDeserialization_Minimal(t *testing.T) {
	rawJSON := `{"id": "bp-min", "name": "Minimal BP"}`

	var bp BusinessPerspectiveConfig
	err := json.Unmarshal([]byte(rawJSON), &bp)
	require.NoError(t, err)

	assert.Equal(t, "bp-min", bp.ID)
	assert.Equal(t, "Minimal BP", bp.Name)
	assert.Empty(t, bp.Description)
	assert.Nil(t, bp.TagFilterExpression)
}
