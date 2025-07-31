package errorx

// ForbiddenError represents a 403 Forbidden error
//type ForbiddenError struct {
//	Code    uint32
//	MsgKey  string
//	Message string
//	Reason  string
//}
//
//func (e *ForbiddenError) Error() string {
//	return e.Message
//}
//
//// NewForbiddenError creates a new forbidden error
//func NewForbiddenError(message string, reason string) error {
//	return &ForbiddenError{
//		Message: message,
//		Reason:  reason,
//	}
//}
//
//func NewForbiddenErrorWithCode(code uint32, message string, reason string) error {
//	return &ForbiddenError{
//		Code:    code,
//		Message: message,
//		Reason:  reason,
//	}
//}
//
//func NewForbiddenErrorWithMsgKey(code uint32, msgKey string, message string, reason string) error {
//	return &ForbiddenError{
//		Code:    code,
//		MsgKey:  msgKey,
//		Message: message,
//		Reason:  reason,
//	}
//}
