package sms_svc

import (
	"context"
	"dxxproject/internal/module/sms/sms_svc/sms_provider"
)

type SvcSmsIface interface {
	SendSms(ctx context.Context, sms sms_provider.SmsSendInfo) (err error)
}
