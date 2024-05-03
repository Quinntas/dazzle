package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func GetRedisClientFromContext(ctx context.Context) RedisClient {
	return ctx.Value(REDIS_CTX_KEY).(RedisClient)
}

func NewRedisClient(redisUrl string) RedisClient {
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		panic(err)
	}
	return RedisClient{
		client: redis.NewClient(opt),
	}
}

func (rc *RedisClient) Set(key string, value string, expiration time.Duration) error {
	return rc.client.Set(context.Background(), key, value, expiration).Err()
}

func (rc *RedisClient) Get(key string) (string, error) {
	return rc.client.Get(context.Background(), key).Result()
}
