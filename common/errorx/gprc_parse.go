package errorx

import (
	"github.com/zeromicro/go-zero/core/jsonx"
	"google.golang.org/grpc/status"
)

func GrpcErrorParse[E error](err error, custom func(pe E) error) error {
	ge, ok := status.FromError(err)
	if !ok {
		return err
	}
	var ice E
	if err := jsonx.Unmarshal([]byte(ge.Message()), &ice); err == nil {
		return custom(ice)
	}
	return err
}
