package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
)

var Client *redis.Client

func InitRedis() *redis.Client {

	Client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	return Client
}

func GetRedisContext() context.Context {
	return context.Background()
}
