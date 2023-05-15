package cryptx

import (
	"log"
	"testing"
)

func TestDES(t *testing.T) {
	key := []byte("2fa6c1e")
	str := "I love this beautiful world!"
	strEncrypted, err := DESEncrypt(str, key)
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("Encrypted: %s", strEncrypted)
	strDecrypted, err := DESDecrypt(strEncrypted, key)
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("Decrypted: %s", strDecrypted)
}

func TestDES1(t *testing.T) {
	key := []byte("ASkW3C9o")
	strEncrypted := "b9c6148cd226e7fb919b90fdc82b1998108711d874873849c385e47226e54dcceb05afebec715eba"
	strDecrypted, err := DESDecrypt(strEncrypted, key)
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("Decrypted: %s", strDecrypted)
}
