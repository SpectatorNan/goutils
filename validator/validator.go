package validator

import (
	"github.com/SpectatorNan/goutils/errorx"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Validator struct {
	validate *validator.Validate
	trans    ut.Translator
}

func (v Validator) Valid(s interface{}) (string, bool) {
	return v.valid(s)
}

type GoZeroValidator struct{}

func (v GoZeroValidator) Validate(r *http.Request, data any) error {
	validate := r.Context().Value(I18nKey).(*Validator)
	msg, valid := validate.Valid(data)
	if !valid {
		return errorx.NewErrMsg(msg)
	}
	return nil
}
