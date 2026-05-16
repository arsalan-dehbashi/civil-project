package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/arsalan-dehbashi/civil-project.git/internal/config"
	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

type Client struct {
	Rdb *redis.Client
}

func NewRedisClient(cfg *config.Config) (*Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: "",
		DB:       0,
	})

	if err := rdb.Ping(Ctx).Err(); err != nil {
		return nil, err
	}

	fmt.Println("Redis connected successfully")
	return &Client{Rdb: rdb}, nil
}

func (c *Client) Set(key, value string, ttl time.Duration) error {
	return c.Rdb.Set(Ctx, key, value, ttl).Err()
}

func (c *Client) Get(key string) (string, error) {
	return c.Rdb.Get(Ctx, key).Result()
}

func (c *Client) Delete(key string) error {
	return c.Rdb.Del(Ctx, key).Err()
}
