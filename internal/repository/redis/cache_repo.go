package redis

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisRepository interface {
	CreateCache(key string, tag interface{}, second time.Duration) error
	GetCache(key string) (string, error)
	DeleteCache(entity string) error
}

type redisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) RedisRepository {
	return &redisRepository{
		client: client,
	}
}

func (r *redisRepository) CreateCache(key string, tag interface{}, second time.Duration) error {

	err := r.client.Set(key, tag, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *redisRepository) GetCache(key string) (string, error) {

	data, err := r.client.Get(key).Result()
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *redisRepository) DeleteCache(entity string) error {
	iter := r.client.Scan(0, entity+"*", 0).Iterator()

	for iter.Next() {
		if err := r.client.Del(iter.Val()).Err(); err != nil {
			return err
		}
	}

	if err := iter.Err(); err != nil {
		return err
	}

	return nil
}
