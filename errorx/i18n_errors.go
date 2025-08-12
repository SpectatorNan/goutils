package errorx

import (
	"context"
	"github.com/SpectatorNan/go-zero-i18n/goi18nx"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type I18nCodeError struct {
	Code       uint32
	MsgKey     string
	DefaultMsg string
	ErrType    ErrorType
}

func (e *I18nCodeError) Error() string {
	str, err := jsonx.Marshal(e)
	if err != nil {
		return NewCodeErrWithCodeMsg(e.Code, e.MsgKey).Error()
	}
	return string(str)
}

func NewErrWithI18nCodeErr(ctx context.Context, e I18nCodeError) error {
	return NewCodeErrWithCodeMsg(
		e.Code,
		goi18nx.FormatText(ctx, e.MsgKey, e.DefaultMsg))
}

func (e *I18nCodeError) ErrorType() ErrorType {
	return e.ErrType
}

func NewI18nError(code uint32, msgKey string, defaultMsg string) *I18nCodeError {
	return NewI18nErrorWithType(code, msgKey, defaultMsg, ErrTypeDefault)
}

func NewI18nErrorWithType(code uint32, msgKey string, defaultMsg string, errType ErrorType) *I18nCodeError {
	return &I18nCodeError{
		Code:       code,
		MsgKey:     msgKey,
		DefaultMsg: defaultMsg,
		ErrType:    errType,
	}
}
