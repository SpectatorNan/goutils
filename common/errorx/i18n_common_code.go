package errorx

var (
	DefaultErr                = &I18nCodeError{Code: ErrCodeDefault, MsgKey: "DefaultErr", DefaultMsg: "The server is out of service, try again later"}
	RequestParamsErr          = &I18nCodeError{Code: ErrCodeRequestParams, MsgKey: "RequestParamErr", DefaultMsg: "Parameter error"}
	NotFoundResourceErr       = &I18nCodeError{Code: ErrCodeNotFoundResource, MsgKey: "NotFoundResourceErr", DefaultMsg: "Resource does not exist"}
	TypeMismatchForConvertErr = &I18nCodeError{Code: ErrCodeTypeMismatchForConvert, MsgKey: "DefaultErr", DefaultMsg: "type mismatch for convert"}
)
