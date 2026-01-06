package ioc

import (
	"dxxproject/internal/svc/email"
	"dxxproject/internal/svc/rate_limit"
	"dxxproject/internal/svc/sms"
	"dxxproject/internal/svc/user"
	"dxxproject/internal/svc/verify_code"
	"dxxproject/pkg/email_provider"
	"dxxproject/pkg/rate_limiter"
	"github.com/samber/do/v2"
)

func svcIoc(injector do.Injector) (err error) {
	// rate limit
	do.Provide(injector, rate_limiter.NewLimiter) //初始化限速器
	do.Provide(injector, rate_limit.NewRateLimit)

	//sms
	//_ = do.As[*ali_sms.AliSmsImpl, sms_provider.SvcSmsIface](injector)

	//do.Provide(injector, svc_sms.NewSms)
	//_ = do.As[svc_sms.SmsSvc, svc_sms.SvcSmsIface](injector)

	do.Provide(injector, sms.NewSmsFailOverSvc)
	err = do.As[*sms.SvcSmsFailOver, sms.SvcSmsIface](injector)
	if err != nil {
		return
	}

	do.Provide(injector, email_provider.NewSmtpIMPL)
	err = do.As[*email_provider.SmtpIMPL, email_provider.EmailProviderIF](injector)
	if err != nil {
		return
	}

	do.Provide(injector, verify_code.NewVerifyCodeSvc)

	do.Provide(injector, email.NewEmailSvc)
	do.Provide(injector, user.NewUserSvc)
	return
}
