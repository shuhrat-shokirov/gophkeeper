package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRootDir(t *testing.T) {
	// Get the current file's directory
	root := rootDir()

	assert.NotContains(t, root, "configs")
	assert.NotContains(t, root, "pkg")
}

func TestNew(t *testing.T) {
	t.Run("config", func(t *testing.T) {
		cf, err := NewConfig()
		require.NoError(t, err)

		assert.NotNil(t, cf, "Expected config to be non-nil")
		assert.NotEmpty(t, cf.GetString("database.dsn"), "Expected db.dsn to be set")
		assert.NotEmpty(t, cf.GetString("logger.level"), "Expected db.dbName to be set")
	})

}

func Test_config_Get(t *testing.T) {
	t.Run("config", func(t *testing.T) {
		cf, err := NewConfig()
		require.NoError(t, err)

		assert.NotNil(t, cf, "Expected config to be non-nil")
		assert.NotEmpty(t, cf.GetString("database.dsn"), "Expected db.dsn to be set")
		assert.NotEmpty(t, cf.GetString("logger.level"), "Expected db.dbName to be set")

		assert.Equal(t, "debug", cf.GetString("logger.level"), "Expected logger level to be 'debug'")
		assert.Equal(t, 6379, cf.GetInt("redis.port"), "Expected redis port to be 6379")
	})
}
