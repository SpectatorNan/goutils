package validator

import (
	"github.com/go-playground/locales"
	//"github.com/go-playground/locales/en"
	//"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"golang.org/x/text/language"
	"net/http"
)

type TranslateFunc func(r *http.Request, validate *validator.Validate, trans ut.Translator, lang language.Tag, bundle *Bundle) *validator.Validate
type RegisterTagFunc func(r *http.Request, validate *validator.Validate)

type ILocales interface {
	GetDefaultTranslator() locales.Translator
	GetSupportTranslators() []locales.Translator
	GetTranslateFunc() TranslateFunc
	GetRegisterTagFunc() RegisterTagFunc
	GetUnmarshal() (string, UnmarshalFunc)
	//VariableNameHandler(r *http.Request, msgId, fieldName string) string
	VariableNameHandler() VariableNameHandlerFunc
}

type Middleware struct {
	//supportTags       []language.Tag
	localizationFiles []string
	translateFunc     TranslateFunc
	registerTagFunc   RegisterTagFunc
	//uniTranslator     *ut.UniversalTranslator // uni only init once, has bug, only first translator is valid
	bundle            *Bundle
	defaultTranslator locales.Translator
	localTranslator   []locales.Translator
}

func NewDefaultMiddleware() *Middleware {
	return &Middleware{}
}

func NewMiddlewareWithILocales(localizationFiles []string, iLocales ILocales) *Middleware {
	m := &Middleware{
		localizationFiles: localizationFiles,
		translateFunc:     iLocales.GetTranslateFunc(),
		registerTagFunc:   iLocales.GetRegisterTagFunc(),
		defaultTranslator: iLocales.GetDefaultTranslator(),
		localTranslator:   iLocales.GetSupportTranslators(),
	}
	if len(m.localizationFiles) > 0 {
		format, unmarshalFunc := iLocales.GetUnmarshal()
		bundle := NewBundleWithTemplatePaths(format, unmarshalFunc, m.localizationFiles...)
		bundle.SetVariableNameHandlerFunc(iLocales.VariableNameHandler())
		m.bundle = bundle 
	}
	return m
}

func NewMiddlewareWithLocalization(localizationFiles []string, translateFunc TranslateFunc,
	defaultTranslator locales.Translator, localTranslator []locales.Translator, registerTagFunc RegisterTagFunc,
	unmarshalFormat string, unmarshalFunc UnmarshalFunc) *Middleware {
	//uni := ut.New(defaultTranslator, localTranslator...)
	m := &Middleware{
		//supportTags:       supportTags,
		localizationFiles: localizationFiles,
		translateFunc:     translateFunc,
		registerTagFunc:   registerTagFunc,
		//uniTranslator:     uni,
		defaultTranslator: defaultTranslator,
		localTranslator:   localTranslator,
	}
	if len(m.localizationFiles) > 0 {
		//defLangTag := m.supportTags[0]
		bundle := NewBundleWithTemplatePaths(unmarshalFormat, unmarshalFunc, m.localizationFiles...)
		m.bundle = bundle
	}
	return m
}

func (m *Middleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r2 := m.withRequest(r)
		next(w, r2)
	}
}

func (m *Middleware) isHasI18n() bool {
	return len(m.localizationFiles) > 0
}

func WithTranslateFunc(m *Middleware, translateFunc TranslateFunc) *Middleware {
	m.translateFunc = translateFunc
	return m
}

func WithBundleVariableNameHandler(m *Middleware, f VariableNameHandlerFunc) {
	m.bundle.SetVariableNameHandlerFunc(f)
}
