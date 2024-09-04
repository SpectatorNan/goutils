package oss

type OssJsCredential struct {
	Policy         string `json:"policy"`
	OSSAccessKeyId string `json:"ossAccessKeyId"`
	Signature      string `json:"signature"`
}
