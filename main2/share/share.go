package share

import (
	"dxxproject/pkg/email_provider"
	"dxxproject/pkg/password_utils"
	"dxxproject/pkg/sf_utils"
	"dxxproject/pkg/sms_provider"
	"github.com/go-redis/redis_rate/v10"
	"github.com/samber/do/v2"
)

type Share struct {
	rateLimiter *redis_rate.Limiter
	pwUtil      password_utils.PasswordUtilIF
	snow        sf_utils.SnowflakeIF
	emailSmtp   *email_provider.SmtpIMPL
	smsIsp      sms_provider.SmsProviderIF
}

func NewShare(injector do.Injector) *Share {
	rateLimiter := do.MustInvoke[*redis_rate.Limiter](injector)
	pwUtil := do.MustInvoke[password_utils.PasswordUtilIF](injector)
	snow := do.MustInvoke[sf_utils.SnowflakeIF](injector)
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
