package errorx

import "google.golang.org/grpc/codes"

type ErrorType uint32

const (
	ErrTypeDefault ErrorType = iota
	ErrTypeNotFoundResource
	ErrTypeForbidden
	ErrTypeInfo
)

const (
	ErrLogLevelInfo = iota
	ErrLogLevelError
)

type IErrorType interface {
	ErrorType() ErrorType
}

func (e ErrorType) StatusCode() codes.Code {
	switch e {
	case ErrTypeDefault:
		return codes.OK
	case ErrTypeNotFoundResource:
		return codes.NotFound
	case ErrTypeForbidden:
		return codes.PermissionDenied
	case ErrTypeInfo:
		return codes.OK
	default:
		return codes.Internal
	}
}
func (e ErrorType) String() string {
	switch e {
	case ErrTypeDefault:
		return "default"
	case ErrTypeNotFoundResource:
		return "not_found_resource"
	case ErrTypeForbidden:
		return "forbidden"
	case ErrTypeInfo:
		return "info"
	default:
		return "unknown"
	}
}
func (e ErrorType) LogLevel() int {
	switch e {
	case ErrTypeNotFoundResource:
		return ErrLogLevelInfo
	case ErrTypeInfo:
		return ErrLogLevelInfo
	default:
		return ErrLogLevelError
	}
}

func ErrorTypeFromString(s string) ErrorType {
	switch s {
	case "default":
		return ErrTypeDefault
	case "not_found_resource":
		return ErrTypeNotFoundResource
	case "forbidden":
		return ErrTypeForbidden
	case "info":
		return ErrTypeInfo
	default:
		return ErrTypeDefault
	}
}
