package share

import (
	"dxxproject/internal/module/email/email_svc/email_provider"
	"dxxproject/internal/module/sms/sms_svc/sms_provider"
	"dxxproject/my/id_generator"
	"dxxproject/my/pwd_util"
	"github.com/go-redis/redis_rate/v10"
	"github.com/samber/do/v2"
)

type Share struct {
	rateLimiter *redis_rate.Limiter
	pwUtil      pwd_util.PwdIface
	snow        id_generator.IdGenIface
	emailSmtp   *email_provider.SmtpIMPL
	smsIsp      sms_provider.SmsProviderIF
}

func NewShare(injector do.Injector) *Share {
	rateLimiter := do.MustInvoke[*redis_rate.Limiter](injector)
	pwUtil := do.MustInvoke[pwd_util.PwdIface](injector)
	snow := do.MustInvoke[id_generator.IdGenIface](injector)
	emailSmtp := do.MustInvoke[*email_provider.SmtpIMPL](injector)
	smsIsp := do.MustInvoke[sms_provider.SmsProviderIF](injector)

	share := &Share{
		rateLimiter: rateLimiter,
		pwUtil:      pwUtil,
		snow:        snow,
		emailSmtp:   emailSmtp,
		smsIsp:      smsIsp,
	}
	return share
}
