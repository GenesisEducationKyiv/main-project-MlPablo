package mail

import (
	"context"
	"encoding/json"
	"fmt"
	"net/smtp"
)

// This is the implementation of logic that can send emails.
// So service doesn't need to know about how we do this, and we can implement any mail interfaces we want
// I'm not sure about putting this into infrastructure folder.
type EmailSender struct {
	cfg     *Config
	auth    smtp.Auth
	address string
}

func NewMailService(cfg *Config) *EmailSender {
	return &EmailSender{
		cfg:     cfg,
		auth:    smtp.PlainAuth("", cfg.user, cfg.password, cfg.smtpHost),
		address: fmt.Sprintf("%s:%s", cfg.smtpHost, cfg.smtpPort),
	}
}

func (e *EmailSender) SendEmail(_ context.Context, data any, receivers ...string) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return e.sendEmail(bytes, receivers...)
}

func (e *EmailSender) sendEmail(message []byte, receiversEmail ...string) error {
	return smtp.SendMail(
		e.address,
		e.auth,
		e.cfg.user,
		receiversEmail,
		message,
	)
}
