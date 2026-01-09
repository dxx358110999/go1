package ioc

import (
	"dxxproject/config_prepare/app_config"
	"dxxproject/internal/sms"
	"dxxproject/internal/sms/sms_provider"
	"github.com/samber/do/v2"
)

func Sms(injector do.Injector) {
	/*
		sms
		短信服务
	*/

	//注入单个sms服务商
	do.Provide(injector, func(injector do.Injector) (*sms_provider.FakeIspIMPL, error) {
		return sms_provider.NewFakeIsp()
	})

	//注入多个sms服务商
	do.Provide(injector, func(injector do.Injector) ([]sms_provider.SmsProviderIF, error) {
		appCfg := do.MustInvoke[*app_config.AppConfig](injector)

		fakeIsp, err := sms_provider.NewFakeIsp()
		if err != nil {
			panic(err)

		}

		aliSms, err := sms_provider.NewAliSmsImpl(appCfg)
		if err != nil {
			panic(err)
		}

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

	//do.Sms(injector, svc_sms.NewSms)
	//_ = do.As[svc_sms.SmsSvc, svc_sms.SvcSmsIface](injector)

	do.Provide(injector, func(injector do.Injector) (*sms.SvcSmsFailOver, error) {
		smsProviders := do.MustInvoke[[]sms_provider.SmsProviderIF](injector)

		return sms.NewSmsFailOverSvc(smsProviders)
	})

}
