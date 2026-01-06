package ioc

import (
	"dxxproject/internal/svc/email_svc"
	"dxxproject/internal/svc/email_svc/email_provider"
	"dxxproject/internal/svc/rate_limit_svc"
	"dxxproject/internal/svc/sms_svc"
	"dxxproject/internal/svc/sms_svc/sms_provider"
	"dxxproject/internal/svc/user_svc"
	"dxxproject/internal/svc/verify_code_svc"
	"dxxproject/pkg/rate_limiter"
	"github.com/samber/do/v2"
)

func userSvc(injector do.Injector) error {
	var err error

	do.Provide(injector, user_svc.NewUserCache)
	do.Provide(injector, user_svc.NewUserDao)
	do.Provide(injector, user_svc.NewUserRepo)
	do.Provide(injector, user_svc.NewUserSvc)
	return err
}

func verifyCodeSvc(injector do.Injector) error {
	/*
		验证码服务
	*/
	var err error

	do.Provide(injector, verify_code_svc.NewVerifyCodeCache)
	do.Provide(injector, verify_code_svc.NewVerifyCodeSvc)
	return err
}

func rateSvc(injector do.Injector) error {
	/*
		rate limit
		限速服务
	*/
	var err error
	do.Provide(injector, rate_limiter.NewLimiter) //初始化限速器
	do.Provide(injector, rate_limit_svc.NewRateLimit)
	return err
}

func smsSvc(injector do.Injector) error {
	/*
		sms
		短信服务
	*/

	//注入单个sms服务商
	do.Provide(injector, sms_provider.NewFakeIsp)
	err := do.As[*sms_provider.FakeIspIMPL, sms_provider.SmsProviderIF](injector)
	if err != nil {
		return err
	}

	//注入多个sms服务商
	fakeIsp, err := sms_provider.NewFakeIsp(injector)
	if err != nil {
		return err
	}

	aliSms, err := sms_provider.NewAliSmsImpl(injector)
	if err != nil {
		return err
	}

	do.Provide(injector, func(injector do.Injector) ([]sms_provider.SmsProviderIF, error) {
		smsIspMap := []sms_provider.SmsProviderIF{
			aliSms,
			fakeIsp,
		}
		return smsIspMap, nil
	})

	/*
		注入 sms 服务
	*/
	//_ = do.As[*ali_sms.AliSmsImpl, sms_provider.SvcSmsIface](injector)

	//do.Provide(injector, svc_sms.NewSms)
	//_ = do.As[svc_sms.SmsSvc, svc_sms.SvcSmsIface](injector)

	do.Provide(injector, sms_svc.NewSmsFailOverSvc)
	err = do.As[*sms_svc.SvcSmsFailOver, sms_svc.SvcSmsIface](injector)
	if err != nil {
		return err
	}

	return nil

}

func emailSvc(injector do.Injector) (err error) {
	/*
		邮件发送服务
	*/

	/*
		服务商
	*/
	do.Provide(injector, email_provider.NewSmtpIMPL)
	err = do.As[*email_provider.SmtpIMPL, email_provider.EmailProviderIF](injector)
	if err != nil {
		return
	}

	/*
		邮件服务
	*/
	do.Provide(injector, email_svc.NewEmailSvc)
	return
}
