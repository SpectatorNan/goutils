package interceptor

import (
	"context"

	"github.com/SpectatorNan/goutils/errors"
	"github.com/SpectatorNan/goutils/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	resp, err = handler(ctx, req)
	if err != nil {
		logErr := err
		err = errorx.GrpcErrorWithDetails(ctx, err)
		if shouldUseInfoLevel(err, logErr) {
			logx.WithContext(ctx).Infof("【RPC-SRV-ERR】 %v", logErr)
		} else {
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", logErr)
		}
	}

	return resp, err
}
func shouldUseInfoLevel(err error, logErr error) bool {
	if et, hasErrorType := err.(errorx.IErrorType); hasErrorType {
		return et.ErrorType().LogLevel() == errorx.ErrLogLevelInfo
	}
	return errors.Is(logErr, errorx.ErrResourceNotFound)
}
