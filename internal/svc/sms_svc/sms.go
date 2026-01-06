package sms_svc

import (
	"context"
	"dxxproject/internal/svc/sms_svc/sms_provider"
	"github.com/samber/do/v2"
)

type SmsSvc struct {
	smsProvider sms_provider.SmsProviderIF
}

var _ SvcSmsIface = new(SmsSvc)

func (r *SmsSvc) SendSms(ctx context.Context, info sms_provider.SmsSendInfo) (err error) {
	err = r.smsProvider.SendSms(ctx, info) //用第一个供应商发送验证码
	if err != nil {
		return
	}
	return
}

func NewSms(injector do.Injector) (*SmsSvc, error) {
	pv, err := do.Invoke[sms_provider.SmsProviderIF](injector)
	if err != nil {
		return nil, err
	}
	sms := &SmsSvc{
		smsProvider: pv,
	}
	return sms, err
}
