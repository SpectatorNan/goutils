package errorx

import (
	"context"
	"github.com/SpectatorNan/go-zero-i18n/goi18nx"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/jsonx"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(ErrResourceNotFound, err.Error())
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

type GrpcErrorInfoDomain string

const (
	GrpcErrorInfoDomain_I18n      = "i18n"
	GrpcErrorInfoDomain_Code      = "code"
	GrpcErrorInfoDomain_Forbidden = "forbidden"
)

func GrpcErrorWithDetails(ctx context.Context, err error) error {
	cause := errors.Cause(err)

	var i18nErr *I18nCodeError
	if errors.As(cause, &i18nErr) {
		code := i18nErr.Code
		msg := i18nErr.DefaultMsg
		if goi18nx.IsHasI18n(context.Background()) {
			msg = goi18nx.FormatText(ctx, i18nErr.MsgKey, i18nErr.DefaultMsg)
		}
		st := status.New(codes.Code(code), msg)
		detailsProto := &errdetails.ErrorInfo{
			Reason: i18nErr.MsgKey,
			Domain: "i18n",
		}
		st, _ = st.WithDetails(detailsProto)
		return st.Err()
	}

	var codeErr *CodeError
	if errors.As(cause, &codeErr) {
		code := codeErr.Code
		msg := codeErr.Message
		st := status.New(codes.Code(code), msg)
		detailsProto := &errdetails.ErrorInfo{
			Reason: codeErr.Reason,
			Domain: "code",
		}
		st, _ = st.WithDetails(detailsProto)
		return st.Err()
	}

	var forbiddenErr *ForbiddenError
	if errors.As(cause, &forbiddenErr) {
		//code := codes.PermissionDenied
		code := forbiddenErr.Code
		msg := forbiddenErr.Message
		if goi18nx.IsHasI18n(context.Background()) {
			msg = goi18nx.FormatText(ctx, forbiddenErr.MsgKey, forbiddenErr.Message)
		}
		st := status.New(codes.Code(code), msg)
		detailsProto := &errdetails.ErrorInfo{
			Reason: forbiddenErr.Reason,
			Domain: "forbidden",
		}
		st, _ = st.WithDetails(detailsProto)
		return st.Err()
	}

	return status.Error(codes.Code(ErrCodeDefault), err.Error())
}

// Function for extracting error details from gRPC status
func ErrorFromGrpcStatus(err error) error {
	if err == nil {
		return nil
	}

	st, ok := status.FromError(err)
	if !ok {
		return err
	}

	// Extract details
	for _, detail := range st.Details() {
		if info, ok := detail.(*errdetails.ErrorInfo); ok {
			domain := info.Domain
			reason := info.Reason

			switch domain {
			case "i18n":
				return &I18nCodeError{
					Code:       uint32(st.Code()),
					MsgKey:     reason,
					DefaultMsg: st.Message(),
				}
			case "code":
				return &CodeError{
					Code:    uint32(st.Code()),
					Message: st.Message(),
					Reason:  reason,
				}
			case "forbidden":
				return &ForbiddenError{
					Code:    uint32(st.Code()),
					Message: st.Message(),
					Reason:  reason,
				}
			}

		}
	}

	return err
}
