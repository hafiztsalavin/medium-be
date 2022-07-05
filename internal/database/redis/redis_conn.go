package redis

import (
	"medium-be/internal/config"

	"github.com/go-redis/redis/v8"
)

func NewRedisClientFromConfig(cfg *config.RedisOption) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return client
}
