package sms_provider

import "context"

type SmsProviderIF interface {
	SendSms(ctx context.Context, info SmsSendInfo) (err error)
}

type SmsSendInfo struct {
	PhoneNumbers []string
	SignedName   string
	TemplateId   string
	Params       map[string]string
}
