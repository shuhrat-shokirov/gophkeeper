package server

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	g, err := New(Params{})
	require.NoError(t, err)
	require.NotNil(t, g)
}
