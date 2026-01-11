package email_svc

import (
	"dxxproject/internal/module/email/email_svc/email_provider"
)

type EmailSvc struct {
	emailIsp *email_provider.SmtpIMPL
}

func (r *EmailSvc) SendEmail(toEmail string, subject string, body string) error {
	return r.emailIsp.SendEmail(toEmail, subject, body)
}

func NewEmailSvc(emailIsp *email_provider.SmtpIMPL) (*EmailSvc, error) {
	email := &EmailSvc{
		emailIsp: emailIsp,
	}
	return email, nil
}
