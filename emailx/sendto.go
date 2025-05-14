package emailx

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type Email struct {
	host   string
	port   string
	tls    bool
	auth   smtp.Auth
	sender SenderInfo
}

func NewEmail(host, port, username, password string, sender SenderInfo, ssl bool) Email {
	return Email{
		host:   host,
		port:   port,
		tls:    ssl,
		auth:   smtp.PlainAuth("", username, password, host),
		sender: sender,
	}
}

func (e *Email) Send(msg SendInfo) error {
	return e.SendMulti(SendMultiInfo{
		Receivers: []string{msg.Receiver},
		Message:   msg.Message,
	})
}

func (e *Email) SendMulti(msg SendMultiInfo) error {
	cli := e.getCli()
	cli.To = msg.Receivers
	cli.Subject = msg.Message.Subject
	if msg.Message.Type == MessageTypeHTML {
		cli.HTML = []byte(msg.Message.Content)
	} else {
		cli.Text = []byte(msg.Message.Content)
	}
	return e.sendBy(cli)
}

func (e *Email) getCli() *email.Email {
	mail := email.NewEmail()
	mail.From = e.sender.getFrom()
	return mail
}

func (e *Email) sendBy(cli *email.Email) error {
	if e.tls {
		return cli.SendWithTLS(e.Addr(), e.auth, &tls.Config{InsecureSkipVerify: true, ServerName: e.host})
	} else {
		return cli.Send(e.Addr(), e.auth)
	}
}
