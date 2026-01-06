package sms

import (
	"context"
	"dxxproject/pkg/sms_provider"
)

type SvcSmsIface interface {
	SendSms(ctx context.Context, sms sms_provider.SmsSendInfo) (err error)
}
