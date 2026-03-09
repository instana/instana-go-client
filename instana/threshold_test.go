package instana_test

import (
	"testing"

	. "github.com/instana/instana-go-client/instana"
	"github.com/stretchr/testify/require"
)

func TestShouldReturnSupportedThresholdOperatorsAsStringSlice(t *testing.T) {
	expected := []string{">", ">=", "<", "<="}
	require.Equal(t, expected, SupportedThresholdOperators.ToStringSlice())
}

func TestShouldReturnSupportedThresholdSeasonalitiesAsStringSlice(t *testing.T) {
	expected := []string{"WEEKLY", "DAILY"}
	require.Equal(t, expected, SupportedThresholdSeasonalities.ToStringSlice())
}
