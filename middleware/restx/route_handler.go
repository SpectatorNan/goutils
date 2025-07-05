package restx

import (
	"fmt"
	"github.com/SpectatorNan/goutils/respx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"time"
)

func AllRouteHandlers() []rest.RunOption {
	return []rest.RunOption{
		RouteNotFound(),
		RouteMethodNotAllow(),
		WithHealthCheckApiHandler(),
	}
}

func RouteNotFoundWithNotAllow() []rest.RunOption {
	return []rest.RunOption{
		RouteNotFound(),
		RouteMethodNotAllow(),
	}
}

func RouteMethodNotAllow() rest.RunOption {
	return rest.WithNotAllowedHandler(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			httpx.WriteJson(w, http.StatusMethodNotAllowed, notAllowData(r))
		}),
	)
}

func notAllowData(r *http.Request) map[string]interface{} {
	m := map[string]interface{}{
		"code":      http.StatusMethodNotAllowed,
		"timestamp": time.Now(),
		"status":    http.StatusMethodNotAllowed,
		"error":     "Method Not Allow",
		"message":   fmt.Sprintf("Method Not Allow, [%s]", r.Method),
		"path":      r.URL.Path,
	}
	return m
}

func RouteNotFound() rest.RunOption {
	return rest.WithNotFoundHandler(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			httpx.WriteJson(w, http.StatusNotFound, notFoundData(r))
		}),
	)
}
func notFoundData(r *http.Request) map[string]interface{} {
	m := map[string]interface{}{
		"code":      http.StatusNotFound,
		"timestamp": time.Now(),
		"status":    http.StatusNotFound,
		"error":     "Route Not Found",
		"message":   "The requested resource was not found",
		"path":      r.URL.Path,
	}
	return m
}

func WithHealthCheckApiHandler() rest.RunOption {
	return func(server *rest.Server) {
		server.AddRoute(rest.Route{
			Method: http.MethodGet,
			Path:   "/healthcheck",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				respx.HttpResult(r, w, map[string]string{"status": "ok"}, nil)
			},
		})
	}
}
