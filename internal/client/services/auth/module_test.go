package auth

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/fx/fxtest"
)

func TestNew(t *testing.T) {
	t.Run("public key not base64", func(t *testing.T) {
		err := os.Setenv("GOPH_KEEPER_PUBLIC_KEY", "invalid_base64")
		require.NoError(t, err)

		s, err := New(Params{})
		require.Error(t, err)
		require.Nil(t, s)
	})

	t.Run("success", func(t *testing.T) {
		err := os.Setenv("GOPH_KEEPER_PUBLIC_KEY", "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLQ==")
		require.NoError(t, err)

		s, err := New(Params{
			Lifecycle: fxtest.NewLifecycle(t),
		})
		require.NoError(t, err)
		require.NotNil(t, s)
	})

	t.Run("lifecycle start and stop", func(t *testing.T) {
		err := os.Setenv("GOPH_KEEPER_PUBLIC_KEY", "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLQ==")
		require.NoError(t, err)

		lc := fxtest.NewLifecycle(t)
		s, err := New(Params{
			Lifecycle: lc,
		})
		require.NoError(t, err)
		require.NotNil(t, s)

		err = lc.Start(t.Context())
		require.NoError(t, err)

		err = lc.Stop(t.Context())
		require.NoError(t, err)
	})
}
