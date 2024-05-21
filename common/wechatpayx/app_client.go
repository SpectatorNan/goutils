package wechatpayx

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/app"
	wxpayv3_utils "github.com/wechatpay-apiv3/wechatpay-go/utils"
	"time"
)

type appClient struct {
	conf Conf
}

func NewAppClient(conf Conf) APPPayClient {
	return newAppClient(conf)
}

func newAppClient(conf Conf) APPPayClient {
	return &appClient{conf: conf}
}

func (c *appClient) PrePay(ctx context.Context, desc, tradeNo string, price int64, validDuration time.Duration) (*AppPrepayResponse, error) {

	mchPrivateKey, err := wxpayv3_utils.LoadPrivateKeyWithPath(c.conf.CertPath)
	if err != nil {
		return nil, err
	}

	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(c.conf.MchId, c.conf.SerialNum, mchPrivateKey, c.conf.ApiKey),
	}

	cli, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	svc := app.AppApiService{Client: cli}
	resp, _, err := svc.PrepayWithRequestPayment(ctx,
		app.PrepayRequest{
			Appid:       core.String(c.conf.AppId),
			Mchid:       core.String(c.conf.MchId),
			Description: core.String(desc),
			OutTradeNo:  core.String(tradeNo),
			TimeExpire:  core.Time(time.Now().Add(validDuration)),
			NotifyUrl:   core.String(c.conf.NotifyUrl),
			Amount: &app.Amount{
				Currency: core.String("CNY"),
				Total:    core.Int64(price),
			},
		},
	)

	if err != nil {
		return nil, err
	}
	return &AppPrepayResponse{
		PrepayId:  resp.PrepayId,
		PartnerId: resp.PartnerId,
		TimeStamp: resp.TimeStamp,
		NonceStr:  resp.NonceStr,
		Package:   resp.Package,
		Sign:      resp.Sign,
	}, nil
}
