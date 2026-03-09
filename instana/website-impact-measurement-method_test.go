package instana_test

import (
	"testing"

	. "github.com/instana/instana-go-client/instana"
	"github.com/stretchr/testify/require"
)

func TestShouldReturnSupportedWebsiteImpactMeasurementMethodsAsStringSlice(t *testing.T) {
	expected := []string{"AGGREGATED", "PER_WINDOW"}
	require.Equal(t, expected, SupportedWebsiteImpactMeasurementMethods.ToStringSlice())
}
