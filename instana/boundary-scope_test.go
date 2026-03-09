package instana_test

import (
	"testing"

	. "github.com/instana/instana-go-client/instana"
	"github.com/stretchr/testify/require"
)

func TestShouldReturnStringRepresentationOfSupportedApplicationConfigBoundaryScopes(t *testing.T) {
	require.Equal(t, []string{"ALL", "INBOUND", "DEFAULT"}, SupportedApplicationConfigBoundaryScopes.ToStringSlice())
}

func TestShouldReturnStringRepresentationOfSupportedApplicationAlertConfigBoundaryScopes(t *testing.T) {
	require.Equal(t, []string{"ALL", "INBOUND"}, SupportedApplicationAlertConfigBoundaryScopes.ToStringSlice())
}
