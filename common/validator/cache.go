package validator

import (
	"context"
	"github.com/SpectatorNan/go-zero-i18n/goi18nx"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/text/language"
	"net/http"
)

func (m *Middleware) withRequest(r *http.Request) *http.Request {
	validate := validator.New()
	deTrans, _ := ut.New(en.New()).GetTranslator(language.English.String())
	// default validator
	v := Validator{
		validate: validate,
		trans:    deTrans,
	}
	if m.isHasI18n() {
		// localization validator and translator
		if m.defaultTranslator == nil {
			logx.Errorf("defaultTranslator is nil")
			return r
		}
		if m.translateFunc == nil {
			logx.Errorf("translateFunc is nil, must be set <RegisterDefaultTranslations> by language tag")
			return r
		}

		lang := r.Header.Get(defaultLangHeaderKey)
		langTag := goi18nx.FetchCurrentLanguageTag(lang, m.supportTags)
		uni := ut.New(m.defaultTranslator, m.localTranslator...)
		trans, _ := uni.GetTranslator(langTag.String())
		v.trans = trans
		if m.translateFunc != nil {
			validate = m.translateFunc(r, validate, trans, langTag)
			v.validate = validate
		}
		for _, fn := range m.registerTagFunc {
			fn(r, validate)
		}
	}

	return r.WithContext(context.WithValue(r.Context(), I18nKey, &v))
}

/*
func withRequest(r *http.Request) *http.Request {

	validate, trans := generateValidate(r)
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		jsonKey := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if jsonKey == "-" {
			return ""
		}
		name := goi18nx.FormatText(r.Context(), fmt.Sprintf("Parameters.%s", field.Name), jsonKey)
		return name
	})

	v := Validator{
		validate: validate,
		trans:    trans,
	}
	return r.WithContext(context.WithValue(r.Context(), I18nKey, &v))
}
*/
