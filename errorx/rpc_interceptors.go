package errorx

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/SpectatorNan/go-zero-i18n/goi18nx"
	"github.com/SpectatorNan/goutils/errors"
	"github.com/SpectatorNan/goutils/tools"
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

			//newErr := errors.WithStackFrom(ErrResourceNotFound, err)
			newErr := ErrResourceNotFound
			return nil, errors.WithMessage(newErr, fmt.Sprintf("original error: %v", err))
		}
		return nil, err
	}
	return resp, nil
}
func framesToString(frames *runtime.Frames) string {
	var sb strings.Builder
	for {
		frame, more := frames.Next()
		sb.WriteString(fmt.Sprintf(
			"%s\n\t%s:%d\n",
			frame.Function,
			frame.File,
			frame.Line,
		))
		if !more {
			break
		}
	}
	return sb.String()
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
	GrpcErrorInfoDomain_I18n = "i18n"
	GrpcErrorInfoDomain_Code = "code"
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
		st := status.New(i18nErr.ErrorType().StatusCode(), msg)
		detailsProto := &errdetails.ErrorInfo{
			Reason: i18nErr.MsgKey,
			Domain: "i18n",
			Metadata: map[string]string{
				"code":        fmt.Sprintf("%d", code),
				"message":     msg,
				"message_key": i18nErr.MsgKey,
				"error_type":  i18nErr.ErrorType().String(),
			},
		}
		st, _ = st.WithDetails(detailsProto)
		return st.Err()
	}

	var codeErr *CodeError
	if errors.As(cause, &codeErr) {
		code := codeErr.Code
		msg := codeErr.Message
		st := status.New(codeErr.ErrorType().StatusCode(), msg)
		detailsProto := &errdetails.ErrorInfo{
			Reason: codeErr.Reason,
			Domain: "code",
			Metadata: map[string]string{
				"code":       fmt.Sprintf("%d", code),
				"message":    msg,
				"error_type": codeErr.ErrorType().String(),
			},
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
				codestr := info.Metadata["code"]
				message := info.Metadata["message"]
				messageKey := info.Metadata["message_key"]
				errorTypestr := info.Metadata["error_type"]
				errorType := ErrorTypeFromString(errorTypestr)
				code := tools.StringToInt64(codestr)
				return &I18nCodeError{
					Code:       uint32(code),
					MsgKey:     messageKey,
					DefaultMsg: message,
					ErrType:    errorType,
				}
			case "code":
				codestr := info.Metadata["code"]
				message := info.Metadata["message"]
				errorTypestr := info.Metadata["error_type"]
				errorType := ErrorTypeFromString(errorTypestr)
				code := tools.StringToInt64(codestr)
				return &CodeError{
					Code:    uint32(code),
					Message: message,
					Reason:  reason,
					ErrType: errorType,
				}
				//case "forbidden":
				//	return &ForbiddenError{
				//		Code:    uint32(st.Code()),
				//		Message: st.Message(),
				//		Reason:  reason,
				//	}
			}

		}
	}

	return err
}
