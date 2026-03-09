package utils_test

import (
	"testing"

	. "github.com/instana/instana-go-client/utils"
	"github.com/stretchr/testify/require"
)

func TestShouldCreateBoolPointerFromBool(t *testing.T) {
	value := true

	require.Equal(t, &value, BoolPtr(value))
}
