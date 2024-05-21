package wechatpayx

import (
	"context"
	"time"
)

type JSPayClient interface {
	PrePay(ctx context.Context, openId, desc, tradeNo string, price int64, validDuration time.Duration) (*JsPrepayResponse, error)
}

type APPPayClient interface {
	PrePay(ctx context.Context, desc, tradeNo string, price int64, validDuration time.Duration) (*AppPrepayResponse, error)
}
