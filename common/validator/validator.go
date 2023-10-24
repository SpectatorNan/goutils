package validator

import (
	"github.com/SpectatorNan/goutils/common/errorx"
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

func (v Validator) Validate(r *http.Request, data any) error {
	//val, ok := data.(*UserLoginReq)
	//if !ok {
	//	return errors.New("data is not correct type")
	//}
	//
	//if *val == req {
	//	return nil
	//}
	validate := r.Context().Value(I18nKey).(*Validator)
	msg, valid := validate.Valid(data)
	if !valid {
		//respx.HttpResult(r, w, nil, errorx.NewErrMsg(msg))
		return errorx.NewErrMsg(msg)
	}
	return nil
}