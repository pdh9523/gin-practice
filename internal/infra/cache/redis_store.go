package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisTokenStore struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisTokenStore(addr string, password string, db int) RefreshTokenStore {
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

func (r *RedisTokenStore) key(userID uint) string {
	return fmt.Sprintf("refresh_token:%d", userID)
}

func (r *RedisTokenStore) Save(userID uint, token string, ttl time.Duration) error {
	return r.client.Set(r.ctx, r.key(userID), token, ttl).Err()
}

func (r *RedisTokenStore) Find(userID uint) (string, error) {
	val, err := r.client.Get(r.ctx, r.key(userID)).Result()
	if errors.Is(err, redis.Nil) {
		return "", errors.New("token not found")
	}
	return val, err
}

func (r *RedisTokenStore) Delete(userID uint) error {
	return r.client.Del(r.ctx, r.key(userID)).Err()
}
