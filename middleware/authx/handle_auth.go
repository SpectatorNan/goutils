package authx

import (
	"context"
	"errors"
	"github.com/SpectatorNan/go-zero-i18n/goi18nx"
	errorx2 "github.com/SpectatorNan/goutils/errorx"
	"github.com/SpectatorNan/goutils/jwtx"
	"github.com/SpectatorNan/goutils/respx"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"net/http/httputil"
)

var (
	errCodeLoginExpire uint32 = 401
	loginExpireErr     error  = &errorx2.I18nCodeError{Code: errCodeLoginExpire, MsgKey: "Users.LoginExpire", DefaultMsg: "Please login"}
	debugMode                 = false
)

func SetDebugMode(debug bool) {
	debugMode = debug
}

func SetLoginExpireCode(code uint32) {
	errCodeLoginExpire = code
}

func SetLoginExpireWithCodeErr(err *errorx2.CodeError) {
	SetLoginExpireCode(err.Code)
	loginExpireErr = err
}

func SetLoginExpireWithI18nErr(err *errorx2.I18nCodeError) {
	SetLoginExpireCode(err.Code)
	loginExpireErr = err
}

func Unauthorized(w http.ResponseWriter, r *http.Request, err error) {

	if debugMode {
		if err != nil {
			DetailAuthLog(r, err.Error())
		} else {
			DetailAuthLog(r, noDetailReason)
		}
	}

	// if user not setting HTTP header, we set header with 401
	//w.WriteHeader(http.StatusUnauthorized)

	httpx.WriteJson(w, http.StatusUnauthorized,
		respx.NewErrorResponse(errCodeLoginExpire,
			goi18nx.FormatText(r.Context(), "User.LoginExpire", "Please login")))

}

const (
	jwtAudience    = "aud"
	jwtExpire      = "exp"
	jwtId          = "jti"
	jwtIssueAt     = "iat"
	jwtIssuer      = "iss"
	jwtNotBefore   = "nbf"
	jwtSubject     = "sub"
	noDetailReason = "no detail reason"
)

var (
	ErrInvalidToken = errors.New("invalid auth token")
	ErrNoClaims     = errors.New("no auth params")
)

func DetailAuthLog(r *http.Request, reason string) {
	// discard dump error, only for debug purpose
	details, _ := httputil.DumpRequest(r, true)
	logx.Errorf("authorize failed: %s\n=> %+v", reason, string(details))
}

func CheckLogin[T jwtx.BaseClaim](w http.ResponseWriter, r *http.Request,
	checkBlackFn func(authorization string) bool, fetchSaltFn func(bClaim T) ([]byte, error)) (*jwtx.CustomClaims[T], context.Context, bool) {
	authorization := r.Header.Get("Authorization")
	if len(authorization) < 1 {
		Unauthorized(w, r, nil)
		return nil, r.Context(), false
	}
	if checkBlackFn(authorization) {
		Unauthorized(w, r, loginExpireErr)
		return nil, r.Context(), false
	}
	var claim jwtx.CustomClaims[T]
	tok, err := jwtv4.ParseWithClaims(authorization, &claim, func(t *jwtv4.Token) (interface{}, error) {
		return fetchSaltFn(claim.BaseClaims)
	})
	//tok, err := parser.ParseToken(r, m.secret, "")
	if err != nil {
		Unauthorized(w, r, err)
		return nil, r.Context(), false
	}

	if !tok.Valid {
		Unauthorized(w, r, ErrInvalidToken)
		return nil, r.Context(), false
	}
	ctx := r.Context()
	ctx = context.WithValue(ctx, JwtBaseClaimKey, claim)

	return &claim, ctx, true
}
