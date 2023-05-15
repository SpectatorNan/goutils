package tools

import (
	"testing"
	"time"
)

func TestUnixMill(t *testing.T) {

	tt := time.UnixMilli(1655038017807)
	t.Log(tt)
}

func TestFormatISO(t *testing.T) {

	t.Log(time.Now().Format(time.RFC3339))
}