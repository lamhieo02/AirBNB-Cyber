package redis

import (
	"github.com/go-redis/redis/v9"
	"go01-airbnb/config"
)

type redisClient struct {
	client *redis.Client
}

func NewRedisClient(cfg *config.Config) *redisClient {
	return &redisClient{
		client:
	}
}
