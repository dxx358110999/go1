package rate_limiter

import (
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

func NewLimiter(redisClient *redis.Client) (*redis_rate.Limiter, error) {
	return redis_rate.NewLimiter(redisClient), nil
}
