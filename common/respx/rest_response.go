package respx

type Response struct {
	Code    uint32      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Reason  string      `json:"reason,omitempty"`
}

func NewSuccessEmptyResponse() *Response {
	return &Response{
		Code:    200,
		Message: "success",
	}
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

func NewResponse(code uint32, msg string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

func NewErrorResponse(errCode uint32, errMsg string) *Response {
	return &Response{
		Code:    errCode,
		Message: errMsg,
	}
}
func NewErrorReasonResponse(errCode uint32, errMsg string, reason string) *Response {
	return &Response{
		Code:    errCode,
		Message: errMsg,
		Reason:  reason,
	}
}
