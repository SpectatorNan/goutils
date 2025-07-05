package cryptx

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

func RSASignWithByte(data []byte, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return RSASignWithPriKey(data, priv)
}

func RSASignWithFileName(data []byte, privateKeyFileName string) ([]byte, error) {
	privateKey, err := LoadKeyBytes(privateKeyFileName)
	if err != nil {
		return nil, err
	}
	return RSASignWithByte(data, privateKey)
}

func RSASignWithPriKey(data []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	if privateKey == nil {
		return nil, errors.New("private key is nil")
	}
	hashed := sha256.Sum256(data)
	return rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
}

// 加密
func RsaEncrypt(origData []byte, publicKey []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

func RsaEncryptWithFileName(origData []byte, publicKeyFileName string) ([]byte, error) {
	publicKey, err := LoadKeyBytes(publicKeyFileName)
	if err != nil {
		return nil, err
	}
	return RsaEncrypt(origData, publicKey)
}

// 解密
func RsaDecrypt(ciphertext []byte, privateKey []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func RsaDecryptWithFileName(ciphertext []byte, privateKeyFileName string) ([]byte, error) {
	privateKey, err := LoadKeyBytes(privateKeyFileName)
	if err != nil {
		return nil, err
	}
	return RsaDecrypt(ciphertext, privateKey)
}

func RSAVerify(pubKey *rsa.PublicKey, data, sign []byte) error {
	return rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, data, sign)
}

func GenerateRSAKeyPair(keySize int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	// Generate a new RSA private key
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return nil, nil, err
	}

	// Get the public key from the private key
	publicKey := &privateKey.PublicKey

	return privateKey, publicKey, nil
}

// GenerateRsaKey create rsa private and public pem file.
// Play: https://go.dev/play/p/zutRHrDqs0X
func GenerateRsaKey(keySize int, priKeyFile, pubKeyFile string) error {
	// private key
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return err
	}

	derText := x509.MarshalPKCS1PrivateKey(privateKey)

	block := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derText,
	}

	file, err := os.Create(priKeyFile)
	if err != nil {
		panic(err)
	}
	err = pem.Encode(file, &block)
	if err != nil {
		return err
	}

	file.Close()

	// public key
	publicKey := privateKey.PublicKey

	derpText, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}

	block = pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: derpText,
	}

	file, err = os.Create(pubKeyFile)
	if err != nil {
		return err
	}

	err = pem.Encode(file, &block)
	if err != nil {
		return err
	}

	file.Close()

	return nil
}

func LoadKeyBytes(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	_, err = file.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func LoadPublicKey(fileName string) (*rsa.PublicKey, error) {

	buf, err := LoadKeyBytes(fileName)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(buf)
	if block == nil {
		return nil, err
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pubKey.(*rsa.PublicKey), nil
}

func LoadPrivateKey(fileName string) (*rsa.PrivateKey, error) {
	buf, err := LoadKeyBytes(fileName)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(buf)
	if block == nil {
		return nil, err
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return priv, nil
}
