package sms

import (
	"context"
	"errors"
	"github.com/SpectatorNan/goutils/common/errorx"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/ttacon/libphonenumber"
	"strings"
)

const (
	Code_Success   = "OK"
	Code_SendLimit = "isv.BUSINESS_LIMIT_CONTROL"
)

var notSetTempConfigErr = errors.New("not set temp config")
var sendLimitErr = errors.New("The number of SMS sent has reached the limit, please try again later")

func SetErrWithNotSetTempConfig(err error) {
	notSetTempConfigErr = err
}
func SetErrWithSendLimitErr(err error) {
	sendLimitErr = err
}

type ErrorHandler func(ctx context.Context) error

var errHandlerMaps = map[string]ErrorHandler{}

func RegisterErrorHandler(code string, handler ErrorHandler) {
	errHandlerMaps[code] = handler
}

type TempConfig struct {
	SignName string
	TplCode  string
	Endpoint string
}

type Client struct {
	accessKeyId     string
	accessKeySecret string
	domesticTemp    *TempConfig
	usdTemp         *TempConfig
}

// NewSmsClient
// accessKeyId: 阿里云短信服务的accessKeyId
// accessKeySecret: 阿里云短信服务的accessKeySecret
// usdTemp: 国际短信配置
// domesticTemp: 国内短信配置
func NewSmsClient(accessKeyId string, accessKeySecret string, usdTemp *TempConfig, domesticTemp *TempConfig) *Client {

	return &Client{
		accessKeyId:     accessKeyId,
		accessKeySecret: accessKeySecret,
		domesticTemp:    domesticTemp,
		usdTemp:         usdTemp,
	}
}
func (s *Client) createClient(endpoint string) *dysmsapi20170525.Client {
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
func (s *Client) Send(phone, code string) (*dysmsapi20170525.SendSmsResponse, error) {
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		TemplateParam: tea.String(code),
	}
	template := s.getTemplate(phone)
	if template == nil {
		return nil, notSetTempConfigErr
	}
	sendSmsRequest.TemplateCode = tea.String(template.TplCode)
	sendSmsRequest.SignName = tea.String(template.SignName)
	client := s.createClient(template.Endpoint)
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

func (s *Client) getTemplate(phone string) *TempConfig {
	if s.usdTemp != nil && s.domesticTemp != nil {
		if s.checkPhoneIsDomestic(phone) {
			return s.domesticTemp
		} else {
			return s.usdTemp
		}
	} else if s.usdTemp != nil {
		return s.usdTemp
	} else if s.domesticTemp != nil {
		return s.domesticTemp
	}
	return nil
}

func (s *Client) getSendClient1(phone string) (*dysmsapi20170525.Client, *string, error) {
	if s.usdTemp != nil && s.domesticTemp != nil {
		if s.checkPhoneIsDomestic(phone) {
			return s.createClient(s.domesticTemp.Endpoint), tea.String(s.domesticTemp.SignName), nil
		} else {
			return s.createClient(s.usdTemp.Endpoint), tea.String(s.usdTemp.SignName), nil
		}
	} else if s.usdTemp != nil {
		return s.createClient(s.usdTemp.Endpoint), tea.String(s.usdTemp.SignName), nil
	} else if s.domesticTemp != nil {
		return s.createClient(s.domesticTemp.Endpoint), tea.String(s.domesticTemp.SignName), nil
	}
	return nil, nil, notSetTempConfigErr
}
func (s *Client) HandResponse(ctx context.Context, res *dysmsapi20170525.SendSmsResponse) error {
	//compare := strings.Compare(*res.Body.Code, "OK")
	//fmt.Println(compare)
	//fmt.Println(*res.Body.Code == "OK")
	//fmt.Println(strings.EqualFold(*res.Body.Code, "OK"))
	//result := *res.Body.Code
	//fmt.Println(result == "OK")
	//if *res.Body.Code == "OK" {
	if strings.EqualFold(*res.Body.Code, Code_Success) {
		return nil
	} else {

		if handler, ok := errHandlerMaps[*res.Body.Code]; ok {
			return handler(ctx)
		}

		if strings.EqualFold(*res.Body.Code, Code_SendLimit) {
			return sendLimitErr
			//return errorx.NewErrMsg(goi18nx.FormatText(ctx, "Code.SMSLimitControl", "The number of SMS sent has reached the limit, please try again later"))
		}
		msg := ""
		if res.Body.Message != nil {
			msg = *res.Body.Message
		}
		return errorx.NewErrMsg(msg)
	}
}

func (s *Client) checkPhoneIsDomestic(phone string) bool {
	p, _ := libphonenumber.Parse(phone, "")
	if p == nil || p.CountryCode == nil {
		return true
	}
	if *p.CountryCode == 86 {
		return true
	}
	return false
}
