package sts

type STSCredentials struct {
	AccessKeySecret string `json:"accessKeySecret"` // key secret
	AccessKeyId     string `json:"accessKeyId"`     // key id
	Expiration      string `json:"expiration"`      // 过期时间
	SecurityToken   string `json:"securityToken"`   // token str
}
