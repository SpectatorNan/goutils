package validator

import (
	"context"
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
		lang := r.Header.Get(defaultLangHeaderKey)
		langTag := m.fetchCurrentLanguageTag(lang)
		//uni := ut.New(m.defaultTranslator, m.localTranslator...)
		trans, _ := m.uniTranslator.GetTranslator(langTag.String())
		v.trans = trans
		if m.translateFunc != nil {
			validate = m.translateFunc(r, validate, trans, langTag, m.bundle)
			v.validate = validate
		} else {
			logx.Errorf("translateFunc is nil, must be set <RegisterDefaultTranslations> by language tag")
			return r
		}
		for _, fn := range m.registerTagFunc {
			fn(r, validate)
		}
	}

	return r.WithContext(context.WithValue(r.Context(), I18nKey, &v))
}

// FetchCurrentLanguageTag fetch current language tag
// lang: Accept-Language
// if lang is not support, return default language tag. is first language tag from localizationFiles
func (m *Middleware) fetchCurrentLanguageTag(lang string) language.Tag {
	return m.bundle.matchCurrentLanguageTag(lang)
}
