package mailer

import "gopkg.in/gomail.v2"

type Config struct {
	Host     string `validate:"required"`
	Port     int    `validate:"required"`
	Username string `validate:"required"`
	Password string `validate:"required"`
	From     string `validate:"required"`
}

type mailer struct {
	cfg    *Config
	mailer *gomail.Dialer
}

type Mailer interface {
	SendMail(to []string, subject, body string) error
}

func (m *mailer) SendMail(to []string, subject, body string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", m.cfg.From)
	message.SetHeader("To", to...)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)

	return m.mailer.DialAndSend(message)
}

func NewMailer(cfg *Config) Mailer {
	return &mailer{
		cfg:    cfg,
		mailer: gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password),
	}
}
