package cryptx

import "testing"

func TestGenPassword(t *testing.T) {
	t.Logf(PasswordEncrypt("HWVOFkGgPTryzICwd7qnJaZR9KQ2i8xe", "123456"))
}
