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
}

func (e *I18nCodeError) Error() string {
	str, err := jsonx.Marshal(e)
	if err != nil {
		return NewErrCodeMsg(e.Code, e.MsgKey).Error()
	}
	return string(str)
}

func NewErrWithI18nCodeErr(ctx context.Context, e I18nCodeError) error {
	return NewErrCodeMsg(
		e.Code,
		goi18nx.FormatText(ctx, e.MsgKey, e.DefaultMsg))
}
