package email

import (
	"dxxproject/pkg/email_provider"
	"github.com/samber/do/v2"
)

type EmailSvc struct {
	emailIsp *email_provider.SmtpIMPL
}

func (r *EmailSvc) SendEmail(toEmail string, subject string, body string) error {
	return r.emailIsp.SendEmail(toEmail, subject, body)
}

func NewEmailSvc(injector do.Injector) (*EmailSvc, error) {
	emailIsp, err := do.Invoke[*email_provider.SmtpIMPL](injector)
	if err != nil {
		return nil, err
	}
	email := &EmailSvc{
		emailIsp: emailIsp,
	}
	return email, nil
}
