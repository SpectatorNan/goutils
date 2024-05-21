package wechatpayx

// PrepayWithRequestPaymentResponse 预下单ID，并包含了调起支付的请求参数
type JsPrepayResponse struct {
	// 预支付交易会话标识
	PrepayId *string `json:"prepay_id"` // revive:disable-line:var-naming
	// 应用ID
	Appid *string `json:"appId"`
	// 时间戳
	TimeStamp *string `json:"timeStamp"`
	// 随机字符串
	NonceStr *string `json:"nonceStr"`
	// 订单详情扩展字符串
	Package *string `json:"package"`
	// 签名方式
	SignType *string `json:"signType"`
	// 签名
	PaySign *string `json:"paySign"`
}

// PrepayWithRequestPaymentResponse 预下单ID，并包含了调起支付的请求参数
type AppPrepayResponse struct {
	// 预支付交易会话标识
	PrepayId *string `json:"prepayId"` // revive:disable-line:var-naming
	// 商户号
	PartnerId *string `json:"partnerId"` // revive:disable-line:var-naming
	// 时间戳
	TimeStamp *string `json:"timeStamp"`
	// 随机字符串
	NonceStr *string `json:"nonceStr"`
	// 订单详情扩展字符串
	Package *string `json:"package"`
	// 签名
	Sign *string `json:"sign"`
}
