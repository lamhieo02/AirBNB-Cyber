package cache

import (
	"context"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"time"
)

type Cache interface {
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Get(ctx context.Context, key string, value any) error
	Delete(ctx context.Context, key string)
}

type redisCache struct {
	store *cache.Cache
}

func NewRedisCache(redisClient *redis.Client) *redisCache {
	store := cache.New(&cache.Options{
		Redis:      redisClient,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	return &redisCache{store: store}
}

func (r *redisCache) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	return r.store.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   ttl,
	})
}

func (r *redisCache) Get(ctx context.Context, key string, value any) error {
	return r.store.Get(ctx, key, value)
}

func (r *redisCache) Delete(ctx context.Context, key string) {
	r.store.Delete(ctx, key)
}
