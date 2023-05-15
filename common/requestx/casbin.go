package requestx

import "goutils/common/casbinx"

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
