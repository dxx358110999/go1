package email_svc

import (
	"dxxproject/internal/svc/email_svc/email_provider"
	"github.com/samber/do/v2"
)

func Provide(injector do.Injector) {
	/*
		邮件发送服务
	*/

	/*
		服务商
	*/
	do.Provide(injector, email_provider.NewSmtpIMPL)
	err := do.As[*email_provider.SmtpIMPL, email_provider.EmailProviderIF](injector)
	if err != nil {
		panic(err)
	}

	/*
		邮件服务
	*/
	do.Provide(injector, NewEmailSvc)

}
