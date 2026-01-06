package ioc

import (
	"dxxproject/pkg/sms_provider"
	"github.com/samber/do/v2"
)

func smsProviderIoc(injector do.Injector) (err error) {
	//注入单个sms服务商
	do.Provide(injector, sms_provider.NewFakeIsp)
	_ = do.As[*sms_provider.FakeIspIMPL, sms_provider.SmsProviderIF](injector)

	//注入多个sms服务商
	fakeIsp, err := sms_provider.NewFakeIsp(injector)
	if err != nil {
		return
	}

	aliSms, err := sms_provider.NewAliSmsImpl(injector)
	if err != nil {
		return
	}

	do.Provide(injector, func(injector do.Injector) ([]sms_provider.SmsProviderIF, error) {
		smsIspMap := []sms_provider.SmsProviderIF{
			aliSms,
			fakeIsp,
		}
		return smsIspMap, nil
	})
	return
}
