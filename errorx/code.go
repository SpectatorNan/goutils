package errorx

import "google.golang.org/grpc/codes"

const OK uint32 = 200

/*
| 1 | 00 | 02 |
| :------ | :------ | :------ |
| 服务级错误（1为系统级错误） | 服务模块代码 | 具体错误代码 |

- 服务级别错误：1 为系统级错误；2 为普通错误，通常是由用户非法操作引起的
- 服务模块为两位数：一个大型系统的服务模块通常不超过两位数，如果超过，说明这个系统该拆分了
- 错误码为两位数：防止一个模块定制过多的错误码，后期不好维护
- `code = 200` 说明是正确返回，`code > 10000` 说明是错误返回
- 错误通常包括系统级错误码和服务级错误码
- 建议代码中按服务模块将错误分类
- 错误码均为 >= 0 的数
- 在本项目中 HTTP Code 固定为 http.StatusOK，错误码通过 code 来表示。
*/

/* Common */
const (
	ErrCodeDefault                uint32 = 10001
	ErrCodeRequestParams                 = 10002
	ErrCodeNotFoundResource              = 10003
	ErrCodeTypeMismatchForConvert        = 10004
)

type ErrorType uint32

const (
	ErrTypeDefault ErrorType = iota
	ErrTypeNotFoundResource
	ErrTypeForbidden
)

type IErrorType interface {
	ErrorType() ErrorType
}

func (e ErrorType) StatusCode() codes.Code {
	switch e {
	case ErrTypeDefault:
		return codes.Code(ErrCodeDefault)
	case ErrTypeNotFoundResource:
		return codes.NotFound
	case ErrTypeForbidden:
		return codes.PermissionDenied
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
	default:
		return "unknown"
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
	default:
		return ErrTypeDefault
	}
}
