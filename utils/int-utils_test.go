package utils_test

import (
	"testing"

	. "github.com/instana/instana-go-client/utils"
	"github.com/stretchr/testify/require"
)

func TestShouldCreateInt64PointerFromInt64(t *testing.T) {
	value := int64(123)

	require.Equal(t, &value, Int64Ptr(value))
}
