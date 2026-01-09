package rate_limit

import (
	"dxxproject/pkg/rate_limiter"
	"github.com/samber/do/v2"
)

func Provide(injector do.Injector) {
	/*
		rate limit
		限速服务
	*/
	do.Provide(injector, rate_limiter.NewLimiter) //初始化限速器
	do.Provide(injector, NewRateLimit)
}
