package memorycache

import (
	"time"

	"github.com/patrickmn/go-cache"
	"go.uber.org/fx"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In
}

type c struct {
	cache *cache.Cache
}

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}, duration time.Duration)
	Delete(key string)
}

func New(p Params) Cache {
	return &c{
		cache: cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (c *c) Get(key string) (interface{}, bool) {
	value, found := c.cache.Get(key)
	if !found {
		return nil, false
	}
	return value, true
}

func (c *c) Set(key string, value interface{}, duration time.Duration) {
	c.cache.Set(key, value, duration)
}

func (c *c) Delete(key string) {
	c.cache.Delete(key)
}
