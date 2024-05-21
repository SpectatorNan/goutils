package wechatpayx

type Conf struct {
	AppId     string
	MchId     string // 商户Id
	SerialNum string // 证书序列号
	ApiKey    string // API密钥
	CertPath  string // 证书路径
	NotifyUrl string // 通知地址
}
