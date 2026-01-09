package ioc

import (
	"dxxproject/internal/email"
	"dxxproject/internal/rate_limit"
	"dxxproject/internal/sms"
	"dxxproject/internal/verify_code/cache"
	"dxxproject/internal/verify_code/hdl"
	"dxxproject/internal/verify_code/svc"
	"github.com/redis/go-redis/v9"
	"github.com/samber/do/v2"
)

func VerifyCode(injector do.Injector) {
	do.Provide(injector, func(injector do.Injector) (*cache.VerifyCodeCache, error) {
		redisClient := do.MustInvoke[*redis.Client](injector) //redis client

		return cache.NewVerifyCodeCache(redisClient)
	})

	do.Provide(injector, func(injector do.Injector) (*svc.VerifyCodeSvc, error) {
		rateLimiter := do.MustInvoke[*rate_limit.RateLimitSvc](injector)
		vcCache := do.MustInvoke[*cache.VerifyCodeCache](injector)
		emailSvc := do.MustInvoke[*email.EmailSvc](injector)
		smsSvc := do.MustInvoke[*sms.SvcSmsFailOver](injector)

		return svc.NewVerifyCodeSvc(
			rateLimiter,
			emailSvc,
			smsSvc,
			vcCache,
		)
	})
	do.Provide(injector, hdl.NewVerifyCodeHdl)
}
