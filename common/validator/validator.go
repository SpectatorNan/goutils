package validator

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
	trans    ut.Translator
}

func (v Validator) Valid(s interface{}) (string, bool) {
	return v.valid(s)
}
