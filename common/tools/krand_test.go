package tools

import (
	"goutils/common/cryptx"
	"testing"
)

func TestMd5ByString(t *testing.T) {
	s := cryptx.Md5ByString("AAA")
	t.Log(s)
}
