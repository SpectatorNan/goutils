package requestx

import (
	"github.com/SpectatorNan/goutils/casbinx"
)

func DefaultCasbin() []casbinx.Info {
	return []casbinx.Info{
		{
			Path:   "v1/login",
			Method: "POST",
		},
		{
			Path:   "v1/captcha",
			Method: "POST",
		},
	}
}
