package redis

import (
	"context"
	"fmt"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/config"
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(config config.RedisConfig) *RedisClient {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &RedisClient{client: redisClient}
}

func (client *RedisClient) CreateSubscription(channel string) *redis.PubSub {
	fmt.Println("Creating subscription")
	return client.client.Subscribe(context.Background(), channel)
}

func (client *RedisClient) Publish(channel string, message []byte) error {
	return client.client.Publish(context.Background(), channel, message).Err()
}
