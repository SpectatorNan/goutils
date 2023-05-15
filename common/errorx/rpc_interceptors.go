package errorx

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/jsonx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

var ErrResourceNotFound = errors.New("record not found")

func SetResourceNotFound(err error) {
	ErrResourceNotFound = err
}

func ResourceNotFoundErrInterceptors(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	resp, err = handler(ctx, req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrResourceNotFound
		}
		return nil, err
	}
	return resp, nil
}

func ParseGrpcError(err error) error {
	if gstatus, ok := status.FromError(err); ok {
		var ice I18nCodeError
		if err := jsonx.Unmarshal([]byte(gstatus.Message()), &ice); err == nil {
			return &ice
		}
		var ce CodeError
		if err := jsonx.Unmarshal([]byte(gstatus.Message()), &ce); err == nil {
			return &ce
		}
	}
	return err
}
