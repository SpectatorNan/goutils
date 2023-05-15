package middleware

import (
	"context"
	"github.com/thinkeridea/go-extend/exnet"
	"net/http"
)

type IPMiddleware struct {
}

func NewIPMiddleware() *IPMiddleware {
	return &IPMiddleware{}
}

const ClientIPCtxKey = "client-ip"
const ClientAgentCtxKey = "client-user-agent"

func (m *IPMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		agent := r.Header.Get("User-Agent")
		 
		ctx := context.WithValue(r.Context(), ClientIPCtxKey, exnet.ClientIP(r))
		ctx = context.WithValue(ctx, ClientAgentCtxKey, agent)
		next(w, r.WithContext(ctx))
	}
}

func FetchClientIP(ctx context.Context) string {
	ip, ok := ctx.Value(ClientIPCtxKey).(string)
	if !ok {
		return ""
	}
	return ip
}
func FetchAgent(ctx context.Context) string {
	agent, ok := ctx.Value(ClientAgentCtxKey).(string)
	if !ok {
		return ""
	}
	return agent
}
