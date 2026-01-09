package ioc

import (
	"dxxproject/internal/email"
	email_provider2 "dxxproject/internal/email/email_provider"
	"github.com/samber/do/v2"
)

func Email(injector do.Injector) {
	/*
		邮件发送服务
	*/

	/*
		服务商
	*/
	do.Provide(injector, email_provider2.NewSmtpIMPL)
	err := do.As[*email_provider2.SmtpIMPL, email_provider2.EmailProviderIF](injector)
	if err != nil {
		panic(err)
	}

	/*
		邮件服务
	*/
	do.Provide(injector, email.NewEmailSvc)

}
