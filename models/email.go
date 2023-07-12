package models

import "github.com/go-mail/mail/v2"

const (
	DefaultSender = "support@lenslocked.dummyTLD"
)

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewEmailService(config SMTPConfig) *EmailService {
	es := EmailService{
		dialer: mail.NewDialer(config.Host, config.Port, config.Username, config.Password),
	}
	return &es
}

type EmailService struct {
	// DefaultSender to be used when one isn't provided. Used also when sender
	// is predetermined like "forgot password" emails
	DefaultSender string

	// unexported fields
	dialer *mail.Dialer
}
