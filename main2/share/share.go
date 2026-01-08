package share

import (
	"dxxproject/internal/svc/email_svc/email_provider"
	"dxxproject/internal/svc/sms_svc/sms_provider"
	"dxxproject/my/passwd_util"
	"dxxproject/pkg/snowflake_ok"
	"github.com/go-redis/redis_rate/v10"
	"github.com/samber/do/v2"
)

type Share struct {
	rateLimiter *redis_rate.Limiter
	pwUtil      passwd_util.PasswordUtilIface
	snow        snowflake_ok.SnowflakeIface
	emailSmtp   *email_provider.SmtpIMPL
	smsIsp      sms_provider.SmsProviderIF
}

func NewShare(injector do.Injector) *Share {
	rateLimiter := do.MustInvoke[*redis_rate.Limiter](injector)
	pwUtil := do.MustInvoke[passwd_util.PasswordUtilIface](injector)
	snow := do.MustInvoke[snowflake_ok.SnowflakeIface](injector)
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
