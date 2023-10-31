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
/*
func generateValidate(r *http.Request) (*validator.Validate, ut.Translator) {

	tag := language2.GetLanguageTag(r)
	en := en.New() //EnglishTrans
	zh := zh.New() //ChineseTrans
	es := es.New() //SpanishTrans
	uni := ut.New(en, zh, es)
	trans, _ := uni.GetTranslator(tag.String())
	validate := validator.New()
	switch tag {
	case language.Chinese, language.SimplifiedChinese, language.TraditionalChinese:
		zhtrans.RegisterDefaultTranslations(validate, trans)
		validate.RegisterTranslation("eqfield", trans, func(ut ut.Translator) error {
			return ut.Add("eqfield", "{0}必须等于{1}", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			a1 := goi18nx.FormatText(r.Context(), fmt.Sprintf("Parameters.%s", fe.StructField()), fe.StructField())
			a2 := goi18nx.FormatText(r.Context(), fmt.Sprintf("Parameters.%s", fe.Param()), fe.Param())
			t, err := ut.T(fe.Tag(), a1, a2)
			if err != nil {
				return fe.(error).Error()
			}
			return t
		})
	case language.Spanish, language.EuropeanSpanish:
		estrans.RegisterDefaultTranslations(validate, trans)
		validate.RegisterTranslation("eqfield", trans, func(ut ut.Translator) error {
			return ut.Add("eqfield", "{0} must be equal to {1}", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			a1 := goi18nx.FormatText(r.Context(), fmt.Sprintf("Parameters.%s", fe.StructField()), fe.StructField())
			a2 := goi18nx.FormatText(r.Context(), fmt.Sprintf("Parameters.%s", fe.Param()), fe.Param())
			t, err := ut.T(fe.Tag(), a1, a2)
			if err != nil {
				return fe.(error).Error()
			}
			return t
		})
	default:
		entrans.RegisterDefaultTranslations(validate, trans)
		validate.RegisterTranslation("eqfield", trans, func(ut ut.Translator) error {
			return ut.Add("eqfield", "{0} must be equal to {1}", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			a1 := goi18nx.FormatText(r.Context(), fmt.Sprintf("Parameters.%s", fe.StructField()), fe.StructField())
			a2 := goi18nx.FormatText(r.Context(), fmt.Sprintf("Parameters.%s", fe.Param()), fe.Param())
			t, err := ut.T(fe.Tag(), a1, a2)
			if err != nil {
				return fe.(error).Error()
			}
			return t
		})
	}
	return validate, trans
}
*/