package ioc

import (
	"dxxproject/internal/module/email/email_svc"
	"dxxproject/internal/module/rate_limit/rate_limit_svc"
	"dxxproject/internal/module/sms/sms_svc"
	"dxxproject/internal/module/verify_code/vc_cache"
	"dxxproject/internal/module/verify_code/vc_handler"
	"dxxproject/internal/module/verify_code/vc_svc"
	"github.com/redis/go-redis/v9"
	"github.com/samber/do/v2"
)

func VerifyCode(injector do.Injector) {
	do.Provide(injector, func(injector do.Injector) (*vc_cache.VerifyCode, error) {
		redisClient := do.MustInvoke[*redis.Client](injector) //redis client

		return vc_cache.NewVerifyCodeCache(redisClient)
	})

	do.Provide(injector, func(injector do.Injector) (*vc_svc.VerifyCodeSvc, error) {
		rateLimiter := do.MustInvoke[*rate_limit_svc.RateLimitSvc](injector)
		vcCache := do.MustInvoke[*vc_cache.VerifyCode](injector)
		emailSvc := do.MustInvoke[*email_svc.EmailSvc](injector)
		smsSvc := do.MustInvoke[*sms_svc.SvcSmsFailOver](injector)

		return vc_svc.NewVerifyCodeSvc(
			rateLimiter,
			emailSvc,
			smsSvc,
			vcCache,
		)
	})
	do.Provide(injector, vc_handler.NewVerifyCodeHdl)
}
