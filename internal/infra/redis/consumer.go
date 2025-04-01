package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type EmailVerifiedHandler func(email string)

type EventConsumer struct {
	client  *redis.Client
	handler EmailVerifiedHandler
}

func NewEventConsumer(client *redis.Client, handler EmailVerifiedHandler) *EventConsumer {
	return &EventConsumer{
		client:  client,
		handler: handler,
	}
}

func (c *EventConsumer) Start(ctx context.Context) {
	go func() {
		for {
			streams, err := c.client.XRead(ctx, &redis.XReadArgs{
				Streams: []string{"email_verified", "0"},
				Block:   0,
			}).Result()
			if err != nil {
				fmt.Println("Stream read error:", err)
				time.Sleep(time.Second)
				continue
			}
			for _, stream := range streams {
				for _, message := range stream.Messages {
					email := message.Values["email"].(string)
					c.handler(email)
				}
			}
		}
	}()
}

func StartEmailVerifiedConsumer() {
	consumer := NewEventConsumer(Client, func(email string) {
		//TODO: 인증된 유저 등록 로직
	})
	consumer.Start(context.Background())
}
