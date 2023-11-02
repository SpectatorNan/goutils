package validator

import (
	"github.com/go-playground/validator/v10"
)

const I18nKey = "SpectatorNan/validate/i18n"

func (v Validator) valid(s interface{}) (string, bool) {

	e := v.validate.Struct(s)
	if e != nil {
		err := e.(validator.ValidationErrors)
		result := removeStructName(err.Translate(v.trans))
		if len(result) > 0 {
			return result[0], false
		}
		return "Parameters valid failed", false
	}
	return "", true
}

func removeStructName(fields map[string]string) []string {
	//result := map[string]string{}
	errs := make([]string, 0)
	for _, err := range fields {
		//result[field[strings.Index(field, ".")+1:]] = err
		errs = append(errs, err)
	}

	return errs //strings.Join(errs, ", ")
}