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
	cli := e.getCli()
	cli.To = []string{msg.Receiver}
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

//func (e *Email) Send(to string, subject string, text string) error {
//
//	mail := e.getClient(to, subject, text)
//	if e.tls {
//		return mail.SendWithTLS(e.Addr(), e.auth, &tls.Config{InsecureSkipVerify: true, ServerName: e.host})
//	} else {
//		return mail.Send(e.Addr(), e.auth)
//	}
//}
//func (e *Email) getClient(to string, subject string, text string) *email.Email {
//
//	mail := email.NewEmail()
//	mail.From = e.sender.getFrom()
//	mail.To = []string{to}
//	mail.Subject = subject
//	//mail.Text = []byte(text)
//	mail.HTML = []byte(text)
//
//	return mail
//}
