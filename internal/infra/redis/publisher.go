package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type EventPublisher struct {
	client *redis.Client
}

func NewEventPublisher(client *redis.Client) *EventPublisher {
	return &EventPublisher{client: client}
}

func (p *EventPublisher) PublishEmailVerified(ctx context.Context, email string) error {
	_, err := p.client.XAdd(ctx, &redis.XAddArgs{
		Stream: "email_verified",
		Values: map[string]interface{}{
			"email": email,
		},
	}).Result()
	return err
}
