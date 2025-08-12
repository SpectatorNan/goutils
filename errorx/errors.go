package errorx

import (
	"fmt"
)

type CodeError struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
	Reason  string `json:"reason,omitempty"`
	TraceId string `json:"trace_id,omitempty"`
	ErrType ErrorType
}

var CodeErrorDebug bool = false

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.Code, e.Message)
}
func (e *CodeError) ErrorType() ErrorType {
	return e.ErrType
}

func NewCodeErrWithMsg(errMsg string) *CodeError {
	return NewCodeErrWithCodeMsg(ErrCodeDefault, errMsg)
}
func NewCodeErrWithCodeMsg(errCode uint32, errMsg string) *CodeError {
	return NewErrCodeMsgReason(errCode, errMsg, "", "")
}
func NewErrCodeMsgReason(errCode uint32, errMsg string, reason string, traceId string) *CodeError {
	return NewCodeErrWithType(errCode, errMsg, reason, traceId, ErrTypeDefault)
}

func NewCodeErrWithType(code uint32, errMsg string, reason string, traceId string, errType ErrorType) *CodeError {
	return &CodeError{Code: code, Message: errMsg, Reason: reason, TraceId: traceId, ErrType: errType}
}
