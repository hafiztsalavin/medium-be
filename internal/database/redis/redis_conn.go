package database

import (
	"medium-be/internal/config"

	"github.com/go-redis/redis"
)

func NewRedisClientFromConfig(cfg *config.RedisOption) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if _, err := client.Ping().Result(); err != nil {
		return client, err
	}

	return client, nil
}
