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

	if m.isHasI18n() {
		lang := r.Header.Get(defaultLangHeaderKey)
		langTag := m.fetchCurrentLanguageTag(lang)

		// uni only init once, has bug, only first translator is valid
		//trans, _ := m.uniTranslator.GetTranslator(langTag.String())

		//uni := ut.New(zh.New(), zh.New(), en.New())
		//trans, _ := uni.GetTranslator(langTag.String())
		deTrans, _ = ut.New(m.defaultTranslator, m.localTranslator...).GetTranslator(langTag.String())
		if m.translateFunc != nil {
			validate = m.translateFunc(r, validate, deTrans, langTag, m.bundle)
		} else {
			logx.Errorf("translateFunc is nil, must be set <RegisterDefaultTranslations> by language tag")
			return r
		}

		if tempMap := m.bundle.GetTemplateByLanguageTag(langTag); tempMap != nil {
			for tag, temp := range tempMap {
				_ = validate.RegisterTranslation(tag, deTrans, temp.RegisterTranslation, TranslationFunc(r, m.bundle, temp))
			}
		}
		if fn := m.registerTagFunc; fn != nil {
			fn(r, validate)
		}
	}
	// default validator
	v := Validator{
		validate: validate,
		trans:    deTrans,
	}
	return r.WithContext(context.WithValue(r.Context(), I18nKey, &v))
}

// FetchCurrentLanguageTag fetch current language tag
// lang: Accept-Language
// if lang is not support, return default language tag. is first language tag from localizationFiles
func (m *Middleware) fetchCurrentLanguageTag(lang string) language.Tag {
	return m.bundle.matchCurrentLanguageTag(lang)
}
