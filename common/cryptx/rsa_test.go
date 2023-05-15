package cryptx

import (
	"encoding/base64"
	"fmt"
	"testing"
)

// 私钥生成
//openssl genrsa -out rsa_private_key.pem 1024
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
 
-----END RSA PRIVATE KEY-----
`)

// 公钥: 根据私钥生成
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
 
-----END PUBLIC KEY-----
`)

func TestRSA(t *testing.T) {
	data, _ := RsaEncrypt([]byte("hello world"), publicKey)
	fmt.Println(base64.StdEncoding.EncodeToString(data))
	origData, _ := RsaDecrypt(data, privateKey)
	fmt.Println(string(origData))
}
