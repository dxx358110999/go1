package sms_svc

import (
	"dxxproject/internal/svc/sms_svc/sms_provider"
	"github.com/samber/do/v2"
)

func Provide(injector do.Injector) {
	/*
		sms
		短信服务
	*/

	//注入单个sms服务商
	do.Provide(injector, sms_provider.NewFakeIsp)
	err := do.As[*sms_provider.FakeIspIMPL, sms_provider.SmsProviderIF](injector)
	if err != nil {
		panic(err)
	}

	//注入多个sms服务商
	fakeIsp, err := sms_provider.NewFakeIsp(injector)
	if err != nil {
		panic(err)

	}

	aliSms, err := sms_provider.NewAliSmsImpl(injector)
	if err != nil {
		panic(err)
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

	do.Provide(injector, NewSmsFailOverSvc)
	err = do.As[*SvcSmsFailOver, SvcSmsIface](injector)
	if err != nil {
		panic(err)

	}

}
