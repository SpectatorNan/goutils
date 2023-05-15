package emailx

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type Email struct {
	host string
	port string
	tls  bool
	auth smtp.Auth
	from string
}

func (e *Email) Addr() string {
	return fmt.Sprintf("%s:%s", e.host, e.port)
}

func NewEmail(host, port, username, password, from string, ssl bool) Email {
	return Email{
		host: host,
		port: port,
		tls:  ssl,
		auth: smtp.PlainAuth("", username, password, host),
		from: from,
	}
}

func (e *Email) Send(to string, subject string, text string) error {

	mail := e.getClient(to, subject, text)
	if e.tls {
		return mail.SendWithTLS(e.Addr(), e.auth, &tls.Config{InsecureSkipVerify: true, ServerName: e.host})
	} else {
		return mail.Send(e.Addr(), e.auth)
	}
}

func (e *Email) getClient(to string, subject string, text string) *email.Email {

	mail := email.NewEmail()
	mail.From = e.from
	mail.To = []string{to}
	mail.Subject = subject
	//mail.Text = []byte(text)
	mail.HTML = []byte(text)

	return mail
}
