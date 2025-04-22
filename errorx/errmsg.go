package errorx

var ErrMsgI18nKey = "DefaultErr"
var ErrMsgDefault = "The server is out of service, try again later"

func SetErrMsgI18nKey(key string) {
	ErrMsgI18nKey = key
}

func SetErrMsgDefault(msg string) {
	ErrMsgDefault = msg
}
