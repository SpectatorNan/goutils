package vipsx

import (
	"github.com/davidbyttow/govips/v2/vips"
	"runtime"
)

var (
	boolFalse   vips.BoolParameter
	intMinusOne vips.IntParameter
)

func init() {
	vips.LoggingSettings(nil, vips.LogLevelError)
	vips.Startup(&vips.Config{
		ConcurrencyLevel: runtime.NumCPU(),
	})
	boolFalse.Set(false)
	intMinusOne.Set(-1)
}
