package instana_test

import (
	"testing"

	. "github.com/instana/instana-go-client/instana"
	"github.com/stretchr/testify/require"
)

func TestShouldReturnSupportedLogLevelsAsStringSlice(t *testing.T) {
	expected := []string{"WARN", "ERROR", "ANY"}
	require.Equal(t, expected, SupportedLogLevels.ToStringSlice())
}
