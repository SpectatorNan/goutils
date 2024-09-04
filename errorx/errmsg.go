package errorx

//var errMessage map[uint32]string

//func init() {
//initMessage()
//}

//func initMessage() {
//errMessage = make(map[uint32]string)
//errMessage[OK] = "RequestSuccess"
//errMessage[DEFAULT_ERROR] = "DefaultErr"
//errMessage[REUQEST_PARAM_ERROR] = "RequestParamErr"
//errMessage[NOT_FOUND_RESOURCE_ERROR] = "NotFoundResourceErr"
//errMessage[EXISTS_RESOURCE_ERROR] = "ExistsResourceErr"
//errMessage[NOT_FOUND_USER_ERROR] = "Users.NotExists"
//errMessage[EXISTS_USER_ERROR] = "Users.Exists"
//errMessage[UnLoginCode] = "LoginExpire"
//errMessage[NOT_PERMISSION] = "NotPermission"
//}

//const DEFAULT_ERR_MSG string = "The server is out of service, try again later"
//const PARAMETER_ERR_MSG string = "Parameter error"
//const NOT_FOUND_RESOURCE_ERR_MSG = "Resource does not exist"
//const NOT_PERMISSION_ERR_MSG = "Not permission"
//
//func MapErrMsgKey(errCode uint32) string {
//	if msgKey, ok := errMessage[errCode]; ok {
//		return msgKey
//	} else {
//		return errMessage[DEFAULT_ERROR]
//	}
//
//	return errMessage[DEFAULT_ERROR]
//}

//func IsCodeErr(errcode uint32) bool {
//	if _, ok := errMessage[errcode]; ok {
//		return true
//	} else {
//		return false
//	}
//}

var ErrMsgI18nKey = "DefaultErr"
var ErrMsgDefault = "The server is out of service, try again later"

func SetErrMsgI18nKey(key string) {
	ErrMsgI18nKey = key
}

func SetErrMsgDefault(msg string) {
	ErrMsgDefault = msg
}
