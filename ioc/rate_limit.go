package ioc

import (
	"dxxproject/internal/rate_limit"
	"dxxproject/pkg/rate_limiter"
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
	"github.com/samber/do/v2"
)

func RateLimit(injector do.Injector) {
	/*
		rate limit
		限速服务
	*/
	//初始化限速器
	do.Provide(injector, func(injector do.Injector) (*redis_rate.Limiter, error) {
		redisClient := do.MustInvoke[*redis.Client](injector)
		return rate_limiter.NewLimiter(redisClient)
	})

	//限速服务
	do.Provide(injector, func(injector do.Injector) (*rate_limit.RateLimitSvc, error) {
		limiter := do.MustInvoke[*redis_rate.Limiter](injector)
		return rate_limit.NewRateLimit(limiter)
	})
}
