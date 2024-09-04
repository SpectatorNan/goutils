package interceptor

import (
	"context"
	"github.com/SpectatorNan/go-zero-i18n/goi18nx"
	errorx2 "github.com/SpectatorNan/goutils/errorx"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	resp, err = handler(ctx, req)
	if err != nil {
		logErr := err
		causeErr := errors.Cause(err)
		switch causeErr.(type) {
		case *errorx2.I18nCodeError:
			serr := causeErr.(*errorx2.I18nCodeError)
			msg := serr.DefaultMsg
			if goi18nx.IsHasI18n(ctx) {
				msg = goi18nx.FormatText(ctx, serr.MsgKey, serr.DefaultMsg)
			}
			err = status.Error(codes.Code(serr.Code), msg)
		case *errorx2.CodeError:
			serr := causeErr.(*errorx2.CodeError)
			err = status.Error(codes.Code(serr.Code), serr.Message)
		default:
			st := status.Convert(causeErr)
			if st.Code() != codes.Unknown {
				// This is a gRPC error
				err = causeErr
			} else {
				err = status.Error(codes.Code(errorx2.ErrCodeDefault), causeErr.Error())
			}
		}
		logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", logErr)
	}

	return resp, err
}
