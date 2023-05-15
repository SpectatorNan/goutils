package requestx

import (
	"context"
	"errors"
	"fmt"
	"github.com/SpectatorNan/goutils/common/jwtx"
	"github.com/SpectatorNan/goutils/common/trace"
	"github.com/zeromicro/go-zero/core/logx"
	"unsafe"
)

var (
	ErrLoginExpire = errors.New("login expired")
)

func FetchUserIdByJwt(ctx context.Context) (int64, error) {
	u, ok := ctx.Value("UserInfo").(jwtx.BaseClaims)
	if !ok {
		logx.WithContext(ctx).Infof("【valid auth token】: User JWT token illegal")
		return 0, ErrLoginExpire //errorc.LoginExpireErrCode //errorx.NewMsgCodeError(errorx.UnLoginCode, "请重新登录")
	}
	return u.ID, nil
}

func FetchUserByJwtClaims(ctx context.Context) (*jwtx.BaseClaims, error) {
	fmt.Println(trace.SpanIdFromContext(ctx))
	fmt.Println(trace.TraceIdFromContext(ctx))
	u, ok := ctx.Value("userInfo").(jwtx.CustomClaims)
	if !ok {
		logx.WithContext(ctx).Infof("【valid auth token】: User JWT token illegal")
		return nil, ErrLoginExpire //errorx.NewMsgCodeError(errorx.UnLoginCode, "请重新登录")
	}
	return &u.BaseClaims, nil
}

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
