package interceptor

import (
	"context"
	"github.com/SpectatorNan/go-zero-i18n/goi18nx"
	"github.com/SpectatorNan/goutils/common/errorx"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	resp, err = handler(ctx, req)
	if err != nil {
		causeErr := errors.Cause(err)
		switch causeErr.(type) {
		case *errorx.I18nCodeError:
			serr := causeErr.(*errorx.I18nCodeError)
			msg := serr.DefaultMsg
			if goi18nx.IsHasI18n(ctx) {
				msg = goi18nx.FormatText(ctx, serr.MsgKey, serr.DefaultMsg)
			}
			err = status.Error(codes.Code(serr.Code), msg)
		case *errorx.CodeError:
			serr := causeErr.(*errorx.CodeError)
			err = status.Error(codes.Code(serr.Code), serr.Message)
		default:
			err = status.Error(codes.Code(errorx.ErrCodeDefault), causeErr.Error())
		}
		logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
	}

	return resp, err
}
