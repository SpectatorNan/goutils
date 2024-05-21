package wechatpayx

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
)

type PayClient interface {
	PrePay(ctx context.Context, desc, tradeNo string, price int64) (*jsapi.PrepayWithRequestPaymentResponse, error)
}
