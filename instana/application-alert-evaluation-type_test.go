package instana_test

import (
	"testing"

	. "github.com/instana/instana-go-client/instana"
	"github.com/stretchr/testify/require"
)

func TestShouldReturnSupportedApplicationAlertEvaluationTypesAsStringSlice(t *testing.T) {
	expected := []string{"PER_AP", "PER_AP_SERVICE", "PER_AP_ENDPOINT"}
	require.Equal(t, expected, SupportedApplicationAlertEvaluationTypes.ToStringSlice())
}
