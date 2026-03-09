package instana_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	. "github.com/instana/instana-go-client/instana"
)

func TestShouldReturnStringRepresentationOfSupporedApplicationConfigScopes(t *testing.T) {
	require.Equal(t, []string{"INCLUDE_NO_DOWNSTREAM", "INCLUDE_IMMEDIATE_DOWNSTREAM_DATABASE_AND_MESSAGING", "INCLUDE_ALL_DOWNSTREAM"}, SupportedApplicationConfigScopes.ToStringSlice())
}
