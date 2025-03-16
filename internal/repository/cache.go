package repository

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	Client *redis.Client
}

func NewCache(redisURL string) (*Cache, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)

	ctx := context.Background()
	_, err = client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &Cache{Client: client}, nil
}

func (c *Cache) Set(ctx context.Context, shortCode, longURL string, ttl time.Duration) error {
	return c.Client.Set(ctx, shortCode, longURL, ttl).Err()
}

func (c *Cache) Get(ctx context.Context, shortCode string) (string, error) {
	return c.Client.Get(ctx, shortCode).Result()
}
