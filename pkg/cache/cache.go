package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"go.uber.org/fx"

	"gophkeeper/pkg/config"
	"gophkeeper/pkg/logger"
)

type Cache interface {
	Save(ctx context.Context, key string, value interface{}, dur time.Duration) error
	Find(ctx context.Context, key string) (value []byte, err error)
	Delete(ctx context.Context, key string) (err error)
}

var Module = fx.Provide(New)

const prefix = "server:"

type Params struct {
	fx.In
	Config config.Config
	Logger logger.Logger
}

type cache struct {
	client *redis.Client
	config config.Config
	logger logger.Logger
}

func New(p Params) (Cache, error) {
	var (
		url      = p.Config.GetString("redis.url")
		port     = p.Config.GetString("redis.port")
		password = p.Config.GetString("redis.password")
		db       = p.Config.GetInt("redis.db")
	)

	client := redis.NewClient(&redis.Options{
		Addr:     url + ":" + port,
		Password: password,
		DB:       db,
	})

	ping := client.Ping(context.Background())
	if ping.Err() != nil {
		p.Logger.Error("Failed to connect to Redis", "error", ping.Err())
		return nil, fmt.Errorf("failed to connect to redis: %w", ping.Err())
	}

	return &cache{
		client: client,
		logger: p.Logger,
		config: p.Config,
	}, nil

}

func (c *cache) Save(ctx context.Context, key string, value interface{}, dur time.Duration) error {
	return c.client.Set(ctx, prefix+key, value, dur).Err()
}

func (c *cache) Find(ctx context.Context, key string) (value []byte, err error) {
	value, err = c.client.Get(ctx, prefix+key).Bytes()
	return
}

func (c *cache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, prefix+key).Err()
}
