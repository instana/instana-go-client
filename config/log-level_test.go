package config_test

import (
	"testing"

	"github.com/instana/instana-go-client/config"
	"github.com/stretchr/testify/require"
)

func TestShouldReturnSupportedLogLevelsAsStringSlice(t *testing.T) {
	expected := []string{"WARN", "ERROR", "ANY"}
	require.Equal(t, expected, config.SupportedLogLevels.ToStringSlice())
}
