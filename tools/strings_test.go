package tools

import (
	"strings"
	"testing"
)

func TestTrim(t *testing.T) {
	a := " 123 "
	t.Logf("[%s]", strings.Trim(a, ""))
	t.Logf("[%s]", strings.Trim(a, " "))
}

func TestSnowId(t *testing.T) {
	id := Int64ToString(1575884008468779008)
	t.Log(id)
}
