package emailx

import (
	"fmt"
	"net/smtp"
	"strings"
	"testing"
)

func TestSend(t *testing.T) {
	user := "user@user.com"
	password := "user@123"
	//host := "smtp.mxhichina.com:465"
	host := "smtp.mxhichina.com:25"
	//host := "smtp.qiye.aliyun.com:465"
	to := "user@user.com"

	subject := "使用Golang发送邮件"

	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="iso-8859-15">
			<title>MMOGA POWER</title>
		</head>
		<body>
			GO 发送邮件，官方连包都帮我们写好了，真是贴心啊！！！
		</body>
		</html>`

	sendUserName := "GOLANG SEND MAIL" //发送邮件的人名称
	fmt.Println("send email")
	err := SendToMail(user, sendUserName, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}

}

func SendToMail(user, sendUserName, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + sendUserName + "<" + user + ">" + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func TestSend2(t *testing.T) {
	user := "user@user.com"
	password := "user@123"
	//host := "smtp.mxhichina.com:465"
	//host := "smtp.mxhichina.com:25"
	//host := "smtp.qiye.aliyun.com:465"
	//host := "smtp.qiye.aliyun.com:25"
	host := "smtp.qiye.aliyun.com"
	to := "user@user.com"

	subject := "使用Golang发送邮件"

	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="utf-8">
			<title>user Service</title>
		</head>
		<body>
			Semd msg!!!
		</body>
		</html>`

	sendUserName := "GOLANG SEND MAIL <" + user + ">" //发送邮件的人名称
	fmt.Println("send email")

	email := NewEmail(host, "465", user, password, sendUserName, true)

	err := email.Send(to, subject, body)
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}
