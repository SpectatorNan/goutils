package errorx

import (
	"fmt"
)

type CodeError struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
	Reason  string `json:"reason,omitempty"`
	TraceId string `json:"trace_id,omitempty"`
}

var CodeErrorDebug bool = false

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.Code, e.Message)
}

func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{Code: errCode, Message: errMsg}
}

//func NewErrCode(errCode uint32) *CodeError {
//	return &CodeError{Code: errCode, Message: MapErrMsgKey(errCode)}
//}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{Code: ErrCodeDefault, Message: errMsg}
}

func NewErrCodeMsgReason(errCode uint32, errMsg string, reason string, traceId string) *CodeError {
	return &CodeError{Code: errCode, Message: errMsg, Reason: reason, TraceId: traceId}
}

//func NewErrCodeReason(errCode uint32, reason string, traceId string) *CodeError {
//	return &CodeError{Code: errCode, Message: MapErrMsgKey(errCode), Reason: reason, TraceId: traceId}
//}

//func NewErrMsgReason(errMsg string, reason string, traceId string) *CodeError {
//	return &CodeError{Code: DEFAULT_ERROR, Message: errMsg, Reason: reason, TraceId: traceId}
//}
