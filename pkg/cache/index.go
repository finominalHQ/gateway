package cache

import (
	"context"
	"math"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/gobuffalo/envy"
)

var (
	client *redis.Client
	ctx    context.Context
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     envy.Get("REDIS_HOST", ""),
		Password: envy.Get("REDIS_PASSWORD", ""),
		DB:       0,
	})

	ctx = context.Background()
}

var Cache = &client

func Get(key string) any {
	value, err := client.Get(ctx, key).Result()
	if err != nil {
		return redis.Nil
	}

	return value
}

func Set(key string, value any, sec int) bool {
	if sec < 0 {
		sec = int(math.Inf(1))
	}

	err := client.Set(ctx, key, value, time.Second*time.Duration(sec)).Err()
	if err != nil {
		return false
	}

	return true
}
