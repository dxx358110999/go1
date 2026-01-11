package sms_provider

import (
	"context"
	"fmt"
)

type FakeIspIMPL struct{}

func (rec *FakeIspIMPL) SendSms(ctx context.Context, info SmsSendInfo) (err error) {
	fmt.Println("fake:", info.PhoneNumbers, info.Params)
	return
}

var _ SmsProviderIF = new(FakeIspIMPL)

func NewFakeIsp() (isp *FakeIspIMPL, err error) {
	isp = &FakeIspIMPL{}
	return
}
