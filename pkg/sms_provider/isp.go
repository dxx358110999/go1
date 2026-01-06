package sms_provider

import (
	"context"
	"fmt"
	"github.com/samber/do/v2"
)

type FakeIspIMPL struct{}

func (rec *FakeIspIMPL) SendSms(ctx context.Context, info SmsSendInfo) (err error) {
	fmt.Println("fake:", info.PhoneNumbers, info.Params)
	return
}

var _ SmsProviderIF = new(FakeIspIMPL)

func NewFakeIsp(injector do.Injector) (isp *FakeIspIMPL, err error) {
	isp = &FakeIspIMPL{}
	return
}
