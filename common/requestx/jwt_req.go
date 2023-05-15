package requestx

import (
	"context"
	"errors"
	"github.com/SpectatorNan/goutils/common/jwtx"
	"github.com/zeromicro/go-zero/core/logx"
	"unsafe"
)

var (
	ErrLoginExpire = errors.New("login expired")
)

const (
	JwtCustomClaimKey = "spectatornan:goutils:jwt:customClaim"
	JwtBaseClaimKey   = "spectatornan:goutils:jwt:baseClaim"
)

func FetchUserIdByJwt[T jwtx.BaseClaim](ctx context.Context) (int64, error) {
	u, ok := ctx.Value(JwtBaseClaimKey).(T)
	if !ok {
		logx.WithContext(ctx).Infof("【valid auth token】: User JWT token illegal")
		return 0, ErrLoginExpire //errorc.LoginExpireErrCode //errorx.NewMsgCodeError(errorx.UnLoginCode, "请重新登录")
	}
	return u.GetUserId(), nil
}

func FetchUserByJwtClaims[T jwtx.BaseClaim](ctx context.Context) (*T, error) {
	//fmt.Println(trace.SpanIdFromContext(ctx))
	//fmt.Println(trace.TraceIdFromContext(ctx))
	u, ok := ctx.Value(JwtCustomClaimKey).(jwtx.CustomClaims[T])
	if !ok {
		logx.WithContext(ctx).Infof("【valid auth token】: User JWT token illegal")
		return nil, ErrLoginExpire //errorx.NewMsgCodeError(errorx.UnLoginCode, "请重新登录")
	}
	return &u.BaseClaims, nil
}

// forget feature ...
func GetKeyValues(ctx context.Context) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	getKeyValue(ctx, m)
	return m
}

type iface struct {
	itab, data uintptr
}

type valueCtx struct {
	context.Context
	key, val interface{}
}

func getKeyValue(ctx context.Context, m map[interface{}]interface{}) {
	ictx := *(*iface)(unsafe.Pointer(&ctx))
	if ictx.data == 0 {
		return
	}

	valCtx := (*valueCtx)(unsafe.Pointer(ictx.data))
	if valCtx != nil && valCtx.key != nil && valCtx.val != nil {
		m[valCtx.key] = valCtx.val
	}
	getKeyValue(valCtx.Context, m)
}
