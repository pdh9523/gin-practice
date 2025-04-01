package cache

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisTokenStore struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisTokenStore(addr string, password string, db int) GlobalCacheStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisTokenStore{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (r *RedisTokenStore) Save(key, value string, ttl time.Duration) error {
	return r.client.Set(r.ctx, key, value, ttl).Err()
}

func (r *RedisTokenStore) Find(key string) (string, error) {
	val, err := r.client.Get(r.ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", errors.New("token not found")
	}
	return val, err
}

func (r *RedisTokenStore) Delete(key string) error {
	return r.client.Del(r.ctx, key).Err()
}
