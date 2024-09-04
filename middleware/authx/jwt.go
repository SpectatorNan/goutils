package authx

import (
	"context"
	"github.com/SpectatorNan/goutils/jwtx"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	//JwtCustomClaimKey = "spectatornan:goutils:jwt:customClaim"
	JwtBaseClaimKey = "spectatornan:goutils:jwt:baseClaim"
)

func FetchUserIdByJwt[T jwtx.BaseClaim](ctx context.Context) (int64, error) {
	u, ok := ctx.Value(JwtBaseClaimKey).(jwtx.CustomClaims[T])
	if !ok {
		logx.WithContext(ctx).Infof("【valid auth token】: User JWT token illegal")
		return 0, loginExpireErr //errorc.LoginExpireErrCode //errorx.NewMsgCodeError(errorx.UnLoginCode, "请重新登录")
	}
	return u.GetUserId(), nil
}

func FetchUserByJwtClaims[T jwtx.BaseClaim](ctx context.Context) (*T, error) {
	//fmt.Println(trace.SpanIdFromContext(ctx))
	//fmt.Println(trace.TraceIdFromContext(ctx))
	u, ok := ctx.Value(JwtBaseClaimKey).(jwtx.CustomClaims[T])
	if !ok {
		logx.WithContext(ctx).Infof("【valid auth token】: User JWT token illegal")
		return nil, loginExpireErr //errorx.NewMsgCodeError(errorx.UnLoginCode, "请重新登录")
	}
	return &u.BaseClaims, nil
}
