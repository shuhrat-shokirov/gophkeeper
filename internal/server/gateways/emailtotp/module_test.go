package emailtotp

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"gophkeeper/pkg/config"
)

func TestNew(t *testing.T) {
	newConfig, err := config.NewConfig()
	require.NoError(t, err)
	require.NotNil(t, newConfig)

	g := New(Params{
		Config: newConfig,
	})
	require.NotNil(t, g)
}

func Test_gateway_SendEmail(t *testing.T) {
	t.Run("empty address", func(t *testing.T) {
		err := os.Setenv("EMAIL_MAIL", "")
		require.NoError(t, err)

		newConfig, err := config.NewConfig()
		require.NoError(t, err)
		require.NotNil(t, newConfig)

		g := New(Params{
			Config: newConfig,
		})
		require.NotNil(t, g)

		err = g.SendEmail(t.Context(), &Request{})
		require.Error(t, err)
	})

	t.Run("test send email", func(t *testing.T) {
		err := os.Setenv("EMAIL_MAIL", "123@gmail.com")
		require.NoError(t, err)

		newConfig, err := config.NewConfig()
		require.NoError(t, err)
		require.NotNil(t, newConfig)

		g := New(Params{
			Config: newConfig,
		})
		require.NotNil(t, g)

		err = g.SendEmail(t.Context(), &Request{})
		require.Error(t, err)
	})
}
