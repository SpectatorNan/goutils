package sms

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/ttacon/libphonenumber"
	"goutils/common/errorx"
	"strings"
)

type TempConfig struct {
	SignName string
	TplCode  string
	Endpoint string
}

type AliyunSms struct {
	accessKeyId     string
	accessKeySecret string
	domesticTemp    TempConfig
	usdTemp         TempConfig
}

func NewAliyunSms(accessKeyId string, accessKeySecret string, usdTemp TempConfig, domesticTemp TempConfig) *AliyunSms {

	return &AliyunSms{
		accessKeyId:     accessKeyId,
		accessKeySecret: accessKeySecret,
		domesticTemp:    domesticTemp,
		usdTemp:         usdTemp,
	}
}
func (s *AliyunSms) createClient(endpoint string) *dysmsapi20170525.Client {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: tea.String(s.accessKeyId),
		// 您的 AccessKey Secret
		AccessKeySecret: tea.String(s.accessKeySecret),
	}
	config.Endpoint = tea.String(endpoint)
	cli, _ := dysmsapi20170525.NewClient(config)
	return cli
}
func (s *AliyunSms) Send(phone, code string) (*dysmsapi20170525.SendSmsResponse, error) {
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		TemplateParam: tea.String(code),
	}
	var client *dysmsapi20170525.Client
	if s.checkPhoneIsDomestic(phone) {
		sendSmsRequest.TemplateCode = tea.String(s.domesticTemp.TplCode)
		sendSmsRequest.SignName = tea.String(s.domesticTemp.SignName)
		client = s.createClient(s.domesticTemp.Endpoint)
	} else {
		sendSmsRequest.TemplateCode = tea.String(s.usdTemp.TplCode)
		sendSmsRequest.SignName = tea.String(s.usdTemp.SignName)
		client = s.createClient(s.usdTemp.Endpoint)
	}
	if client == nil {
		return nil, errorx.NewErrMsg("init sms client failed")
	}

	runtime := &util.RuntimeOptions{}
	_resp, err := client.SendSmsWithOptions(sendSmsRequest, runtime)
	if err != nil {
		//var error = &tea.SDKError{}
		//if _t, ok := err.(*tea.SDKError); ok {
		//	error = _t
		//} else {
		//	error.Message = tea.String(err.Error())
		//}
		//// 如有需要，请打印 error
		//util.AssertAsString(error.Message)
		return nil, err
	}

	return _resp, nil
}
func (s *AliyunSms) HandResponse(res *dysmsapi20170525.SendSmsResponse) error {
	//compare := strings.Compare(*res.Body.Code, "OK")
	//fmt.Println(compare)
	//fmt.Println(*res.Body.Code == "OK")
	//fmt.Println(strings.EqualFold(*res.Body.Code, "OK"))
	//result := *res.Body.Code
	//fmt.Println(result == "OK")
	//if *res.Body.Code == "OK" {
	if strings.EqualFold(*res.Body.Code, "OK") {
		return nil
	} else {
		return errorx.NewErrMsg(res.Body.String())
	}
}

func (s *AliyunSms) checkPhoneIsDomestic(phone string) bool {
	p, _ := libphonenumber.Parse(phone, "")
	if p == nil || p.CountryCode == nil {
		return true
	}
	if *p.CountryCode == 86 {
		return true
	}
	return false
}
