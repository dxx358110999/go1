package ioc

import (
	"dxxproject/internal/rate_limit"
	"dxxproject/pkg/rate_limiter"
	"github.com/samber/do/v2"
)

func RateLimit(injector do.Injector) {
	/*
		rate limit
		限速服务
	*/
	do.Provide(injector, rate_limiter.NewLimiter) //初始化限速器
	do.Provide(injector, rate_limit.NewRateLimit)
}
