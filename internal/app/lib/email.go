package lib

import (
	"fmt"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/config"
	"go.uber.org/fx"
	"gopkg.in/gomail.v2"
)

type Email interface {
	SendMail(emails []string, subject string, body string) error
}

type EmailParams struct {
	fx.In

	config.EmailConfig
}

func NewEmail2(params EmailParams) Email {
	return &params
}
func (e *EmailParams) SendMail(emails []string, subject string, body string) error {
	fmt.Println(e.Sender())
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", e.Sender())
	mailer.SetHeader("To", emails...)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)
	dialer := gomail.NewDialer(
		e.Host(),
		e.Port(),
		e.Username(),
		e.Password(),
	)

	err := dialer.DialAndSend(mailer)
	return err
}
