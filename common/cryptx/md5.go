package cryptx

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

/** 加密方式 **/

func Md5ByString(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		panic(err)
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}

func Md5ByBytes(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}

// 返回一个32位md5加密后的字符串
func GetMD5Encode(data string, salt *string) string {
	h := md5.New()
	h.Write([]byte(data))
	if salt != nil {
		h.Write([]byte(*salt))
	}
	return hex.EncodeToString(h.Sum(nil))
}

// 返回一个16位md5加密后的字符串
func Get16MD5Encode(data string) string {
	return GetMD5Encode(data, nil)[8:24]
}
