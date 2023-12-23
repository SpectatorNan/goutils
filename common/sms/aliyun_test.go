package sms

import (
	"github.com/alibabacloud-go/tea/tea"
	"testing"
)

func TestSendAliSmsCode(t *testing.T) {

	sms := NewAliyunSms("", "", TempConfig{
		SignName: "",
		TplCode:  "",
	}, TempConfig{
		SignName: "阿里云短信测试",
		TplCode:  "SMS_154950909",
	})
	res, err := sms.Send("", "{\"code\":\"123456\"}")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}
	if res.Body.Code == tea.String("OK") {
		t.Log("Success")
	} else {
		t.Log("Failed")
		t.Log(res.Body.String())
	}
}

type Config struct {
	SignName string
	TplCode  string
}

var international = Config{
	SignName: "",
	TplCode:  "",
}
