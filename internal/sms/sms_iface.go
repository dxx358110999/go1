package sms

import (
	"context"
	"dxxproject/internal/sms/sms_provider"
)

type SvcSmsIface interface {
	SendSms(ctx context.Context, sms sms_provider.SmsSendInfo) (err error)
}
