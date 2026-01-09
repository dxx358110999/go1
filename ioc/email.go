package ioc

import (
	"dxxproject/config_prepare/app_config"
	"dxxproject/internal/email"
	"dxxproject/internal/email/email_provider"
	"github.com/samber/do/v2"
)

func Email(injector do.Injector) {
	/*
		邮件发送服务
	*/

	/*
		服务商
	*/
	do.Provide(injector, func(injector do.Injector) (*email_provider.SmtpIMPL, error) {
		appCfg := do.MustInvoke[*app_config.AppConfig](injector)
		return email_provider.NewSmtpIMPL(appCfg)
	})

	/*
		邮件服务
	*/
	do.Provide(injector, func(injector do.Injector) (*email.EmailSvc, error) {
		emailIsp := do.MustInvoke[*email_provider.SmtpIMPL](injector)
		return email.NewEmailSvc(emailIsp)
	})

}
