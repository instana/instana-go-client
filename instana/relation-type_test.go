package instana_test

import (
	"testing"

	. "github.com/instana/instana-go-client/instana"
	"github.com/stretchr/testify/require"
)

func TestShouldReturnSupportedRelationTypesAsStringSlice(t *testing.T) {
	expected := []string{"USER", "API_TOKEN", "ROLE", "TEAM", "GLOBAL"}
	require.Equal(t, expected, SupportedRelationTypes.ToStringSlice())
}
