package data

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	s := New(Params{})
	require.NotNil(t, s)
}
