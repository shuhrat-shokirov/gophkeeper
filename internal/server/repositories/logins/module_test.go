package logins

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	r := New(Params{})
	require.NotNil(t, r)
}
