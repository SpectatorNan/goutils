package restx

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"time"
)

func RouteMethodNotAllow() rest.RunOption {
	return rest.WithNotAllowedHandler(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.WriteHeader(http.StatusMethodNotAllowed)
			_, _ = w.Write(notAllowBytes(r))
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
