package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"goutils/common/errorx"
	"net/http"
	"time"
)

func JwtUnAuthorizedHandle() rest.RunOption {
	return rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		logx.Info("===========jwt=WithUnauthorizedCallback=====================")
		httpx.Error(w, errorx.NewErrCodeMsg(errCodeLoginExpire, "请登录"))
	})
}

func RouteMethodNotAllow() rest.RunOption {
	return rest.WithNotAllowedHandler(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write(notAllowBytes(r))
		}),
	)
}

func notAllowBytes(r *http.Request) []byte {
	m := map[string]interface{}{
		"code":      http.StatusMethodNotAllowed,
		"timestamp": time.Now(),
		"status":    http.StatusMethodNotAllowed,
		"error":     "Method Not Allow",
		"message":   fmt.Sprintf("Method Not Allow, [%s]", r.Method),
		"path":      r.URL.Path,
	}
	result, _ := json.Marshal(m)
	return result
}
