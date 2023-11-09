package cryptx

import "encoding/base64"

func Base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func Base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func Base64EncodeString(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

func Base64DecodeString(src string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(src)
	return string(bytes), err
}