package config

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewConfig)

type Config interface {
	GetInt(key string) int
	GetString(key string) string
}

type config struct {
	cfg *viper.Viper
}

func NewConfig() (Config, error) {

	cfg := viper.New()
	cfg.SetConfigName("config")
	cfg.SetConfigType("json")
	cfg.AddConfigPath("./configs")

	cfg.AddConfigPath(rootDir() + "/configs")

	if err := cfg.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	cfg.AutomaticEnv()

	cfg.WatchConfig()

	return &config{cfg: cfg}, nil
}

func (c *config) GetInt(key string) int {
	return c.cfg.GetInt(key)
}

func (c *config) GetString(key string) string {
	return c.cfg.GetString(key)
}

func rootDir() string {
	_, currFilePath, _, _ := runtime.Caller(0) // get current file full path
	d := path.Dir(path.Dir(currFilePath))      // this equals: `cd ../../config.go` command
	return filepath.Dir(d)
}
