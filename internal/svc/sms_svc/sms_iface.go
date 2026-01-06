package sms_svc

import (
	"context"
	"dxxproject/internal/svc/sms_svc/sms_provider"
)

type SvcSmsIface interface {
	SendSms(ctx context.Context, sms sms_provider.SmsSendInfo) (err error)
}
