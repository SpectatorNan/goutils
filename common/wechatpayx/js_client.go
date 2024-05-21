package wechatpayx

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	wxpayv3_utils "github.com/wechatpay-apiv3/wechatpay-go/utils"
	"time"
)

type jsClient struct {
	conf Conf
}

func NewJsClient(conf Conf) JSPayClient {
	return newJsClient(conf)
}

func newJsClient(conf Conf) JSPayClient {
	return &jsClient{conf: conf}
}

func (c *jsClient) PrePay(ctx context.Context, openId, desc, tradeNo string, price int64, validDuration time.Duration) (*JsPrepayResponse, error) {
	// Pay logic

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

	svc := jsapi.JsapiApiService{Client: cli}

	resp, _, err := svc.PrepayWithRequestPayment(ctx,
		jsapi.PrepayRequest{
			Appid:       core.String(c.conf.AppId),
			Mchid:       core.String(c.conf.MchId),
			Description: core.String(desc),
			OutTradeNo:  core.String(tradeNo),
			TimeExpire:  core.Time(time.Now().Add(validDuration)),
			NotifyUrl:   core.String(c.conf.NotifyUrl),
			Amount: &jsapi.Amount{
				Currency: core.String("CNY"),
				Total:    core.Int64(price),
			},
			Payer: &jsapi.Payer{
				Openid: core.String(openId),
			},
			//Detail: &native.Detail{
			//	CostPrice: core.Int64(608800),
			//	GoodsDetail: []native.GoodsDetail{native.GoodsDetail{
			//		GoodsName:        core.String("iPhoneX 256G"),
			//		MerchantGoodsId:  core.String("ABC"),
			//		Quantity:         core.Int64(1),
			//		UnitPrice:        core.Int64(828800),
			//		WechatpayGoodsId: core.String("1001"),
			//	}},
			//	InvoiceId: core.String("wx123"),
			//},
			//SettleInfo: &native.SettleInfo{
			//	ProfitSharing: core.Bool(false),
			//},
			//SceneInfo: &native.SceneInfo{
			//	DeviceId:      core.String("013467007045764"),
			//	PayerClientIp: core.String("14.23.150.211"),
			//	StoreInfo: &native.StoreInfo{
			//		Address:  core.String("广东省深圳市南山区科技中一道10000号"),
			//		AreaCode: core.String("440305"),
			//		Id:       core.String("0001"),
			//		Name:     core.String("腾讯大厦分店"),
			//	},
			//},
		},
	)

	if err != nil {
		return nil, err
	}
	return &JsPrepayResponse{
		PrepayId:  resp.PrepayId,
		Appid:     resp.Appid,
		TimeStamp: resp.TimeStamp,
		NonceStr:  resp.NonceStr,
		Package:   resp.Package,
		SignType:  resp.SignType,
		PaySign:   resp.PaySign,
	}, nil
}
