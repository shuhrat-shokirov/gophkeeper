package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"

	"gophkeeper/internal/server/exceptions"
	"gophkeeper/pkg/config"
	"gophkeeper/pkg/logger"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In
	Config config.Config
	Logger logger.Logger
}

type Cache interface {
	Save(ctx context.Context, key string, value any, dur time.Duration) error
	Find(ctx context.Context, key string, value any) error
	Delete(ctx context.Context, key string) (err error)
}

type cache struct {
	client *redis.Client
	config config.Config
	logger logger.Logger
	prefix string
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
		prefix: p.Config.GetString("redis.prefix"),
	}, nil
}

func (c *cache) Save(ctx context.Context, key string, value any, dur time.Duration) error {
	marshal, err := json.Marshal(value)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return exceptions.ErrNotFound
		}
		return fmt.Errorf("marshal json: %w", err)
	}

	return c.client.Set(ctx, c.prefix+key, marshal, dur).Err()
}

func (c *cache) Find(ctx context.Context, key string, value any) error {
	val, err := c.client.Get(ctx, c.prefix+key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(val), value)
}

func (c *cache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, c.prefix+key).Err()
}
