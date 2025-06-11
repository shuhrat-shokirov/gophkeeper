package redis

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/fx/fxtest"

	"gophkeeper/pkg/config"
	"gophkeeper/pkg/logger"
)

func TestNew(t *testing.T) {
	t.Run("ping failed", func(t *testing.T) {

		err := os.Setenv("REDIS_PORT", "6333")
		require.NoError(t, err)

		newConfig, err := config.NewConfig()
		require.NoError(t, err)
		require.NotNil(t, newConfig)

		l, err := logger.New(logger.Params{
			Lifecycle: fxtest.NewLifecycle(t),
			Config:    newConfig,
		})
		require.NoError(t, err)
		require.NotNil(t, l)

		c, err := New(Params{
			Config: newConfig,
			Logger: l,
		})
		require.Error(t, err)
		require.Nil(t, c, "Expected Redis client to be nil when ping fails")
	})
	t.Run("success", func(t *testing.T) {
		testInit(t)
	})
}

func testInit(t *testing.T) Cache {
	t.Helper()

	err := os.Setenv("REDIS_PORT", "6379")
	require.NoError(t, err)

	err = os.Setenv("REDIS_URL", "localhost")
	require.NoError(t, err)

	newConfig, err := config.NewConfig()
	require.NoError(t, err)
	require.NotNil(t, newConfig)

	l, err := logger.New(logger.Params{
		Lifecycle: fxtest.NewLifecycle(t),
		Config:    newConfig,
	})
	require.NoError(t, err)
	require.NotNil(t, l)

	c, err := New(Params{
		Config: newConfig,
		Logger: l,
	})
	require.NoError(t, err)
	require.NotNil(t, c, "Expected Redis client to be non-nil on successful connection")

	return c
}

func Test_cache_Save(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		c := testInit(t)

		err := c.Save(t.Context(), "test_key", "test_value", 0)
		require.NoError(t, err, "Expected Save to succeed")

		var value string
		err = c.Find(t.Context(), "test_key", &value)
		require.NoError(t, err, "Expected Find to succeed")
		require.Equal(t, "test_value", value, "Expected value to match saved value")
	})
	t.Run("not found", func(t *testing.T) {
		c := testInit(t)

		var value string
		err := c.Find(t.Context(), "non_existing_key", &value)
		require.Error(t, err, "Expected Find to fail for non-existing key")
		require.Equal(t, "redis: nil", err.Error(), "Expected error to be 'redis: nil'")
	})
}

func Test_cache_Delete(t *testing.T) {
	t.Run("delete existing key", func(t *testing.T) {
		c := testInit(t)

		err := c.Save(t.Context(), "test_key", "test_value", 0)
		require.NoError(t, err, "Expected Save to succeed")

		err = c.Delete(t.Context(), "test_key")
		require.NoError(t, err, "Expected Delete to succeed")

		var value string
		err = c.Find(t.Context(), "test_key", &value)
		require.Error(t, err, "Expected Find to fail after Delete")
		require.Equal(t, "redis: nil", err.Error(), "Expected error to be 'redis: nil'")
	})

	t.Run("delete non existing key", func(t *testing.T) {
		c := testInit(t)

		err := c.Delete(t.Context(), "non_existing_key")
		require.NoError(t, err, "Expected Delete to succeed even for non-existing key")
	})
}

func Test_cache_Find(t *testing.T) {
	t.Run("find existing key", func(t *testing.T) {
		c := testInit(t)

		err := c.Save(t.Context(), "test_key", "test_value", 0)
		require.NoError(t, err, "Expected Save to succeed")

		var value string
		err = c.Find(t.Context(), "test_key", &value)
		require.NoError(t, err, "Expected Find to succeed")
		require.Equal(t, "test_value", value, "Expected value to match saved value")
	})

	t.Run("find non existing key", func(t *testing.T) {
		c := testInit(t)

		var value string
		err := c.Find(t.Context(), "non_existing_key", &value)
		require.Error(t, err, "Expected Find to fail for non-existing key")
		require.Equal(t, "redis: nil", err.Error(), "Expected error to be 'redis: nil'")
	})
}
