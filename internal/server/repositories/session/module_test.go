package session

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		repo := New(Params{})
		require.NotNil(t, repo)
	})
}
