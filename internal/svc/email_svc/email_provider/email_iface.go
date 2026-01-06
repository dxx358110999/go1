package email_provider

type EmailProviderIF interface {
	SendEmail(toEmail string, subject string, body string) error
}
