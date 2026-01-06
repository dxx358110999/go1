package rate_limiter

import (
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
	"github.com/samber/do/v2"
)

func NewLimiter(injector do.Injector) (limiter *redis_rate.Limiter, err error) {
	redisClient := do.MustInvoke[*redis.Client](injector)
	limiter = redis_rate.NewLimiter(redisClient)
	return
}
