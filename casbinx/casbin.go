package casbinx

import (
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

type (
	// Authorizer stores the casbin handler
	Authorizer struct {
		enforcer *casbin.Enforcer
		uidField string
	}
	// AuthorizerOption represents an option.
	AuthorizerOption func(opt *Authorizer)
)

// WithUidField returns a custom user unique identity option.
func WithUidField(uidField string) AuthorizerOption {
	return func(opt *Authorizer) {
		opt.uidField = uidField
	}
}

func NewAuthorizer(e *casbin.Enforcer, opts ...AuthorizerOption) *Authorizer {
	a := &Authorizer{enforcer: e}
	// init an Authorizer
	a.init(opts...)
	return a
}

func (a *Authorizer) Handle(fn func() []byte) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			if !a.CheckPermission(request) {
				writer.WriteHeader(http.StatusOK)
				writer.Write(fn())
				return
			}
			next.ServeHTTP(writer, request)
		}
	}
}

// NewAuthorizer returns the authorizer, uses a Casbin enforcer as input
func NewAuthorizerHandler(e *casbin.Enforcer, opts ...AuthorizerOption) func(http.Handler) http.Handler {
	a := &Authorizer{enforcer: e}
	// init an Authorizer
	a.init(opts...)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if !a.CheckPermission(request) {
				a.RequirePermission(writer)
				return
			}
			next.ServeHTTP(writer, request)
		})
	}
}

func (a *Authorizer) init(opts ...AuthorizerOption) {
	a.uidField = "username"

	for _, opt := range opts {
		opt(a)
	}
}

// GetUid gets the uid from the JWT Claims.
func (a *Authorizer) GetUid(r *http.Request) (string, bool) {
	uid, ok := r.Context().Value(a.uidField).(string)

	return uid, ok
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *Authorizer) CheckPermission(r *http.Request) bool {
	uid, ok := a.GetUid(r)
	if !ok {
		return false
	}
	method := r.Method
	path := r.URL.Path

	allowed, err := a.enforcer.Enforce(uid, path, method)
	if err != nil {
		logx.WithContext(r.Context()).Errorf("[CASBIN] enforce err %s", err.Error())
	}

	return allowed
}

// RequirePermission returns the 403 Forbidden to the client.
func (a *Authorizer) RequirePermission(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusForbidden)
	writer.Write([]byte("not permission1"))
}
