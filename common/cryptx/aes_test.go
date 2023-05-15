package cryptx

import "testing"

func TestAES(t *testing.T) {

	orig := "hello world"
	key := "123456781234567812345678"
	t.Logf("origin: %s", orig)

	encryptCode := AesEncrypt(orig, key)
	t.Logf("encrypt code: %s", encryptCode)

	decryptCode := AesDecrypt(encryptCode, key)
	t.Logf("decrypt code: %s", decryptCode)
}