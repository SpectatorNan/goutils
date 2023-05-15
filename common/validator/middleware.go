package validator

import "net/http"

type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r2 := WithRequest(r)
		next(w, r2)
	}
}
