package instana_test

import (
	"testing"

	. "github.com/instana/instana-go-client/instana"
	"github.com/stretchr/testify/require"
)

func TestShouldReturnSupportedAccessTypesAsStringSlice(t *testing.T) {
	expected := []string{"READ", "READ_WRITE"}
	require.Equal(t, expected, SupportedAccessTypes.ToStringSlice())
}
