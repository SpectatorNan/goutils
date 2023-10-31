package validator

import (
	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"golang.org/x/text/language"
	"net/http"
)

type TranslateFunc func(r *http.Request, validate *validator.Validate, trans ut.Translator, lang language.Tag) (*validator.Validate)
type RegisterTagFunc func(r *http.Request, validate *validator.Validate)

type Middleware struct {
	supportTags       []language.Tag
	localizationFiles []string
	translateFunc     TranslateFunc
	registerTagFunc   []RegisterTagFunc
	defaultTranslator locales.Translator
	localTranslator   []locales.Translator
}

func NewDefaultMiddleware() *Middleware {
	return &Middleware{
	}
}

func NewMiddlewareWithLocalization(supportTags []language.Tag, localizationFiles []string, translateFunc TranslateFunc,
	 defaultTranslator locales.Translator, localTranslator   []locales.Translator, registerTagFunc []RegisterTagFunc) *Middleware {
	return &Middleware{
		supportTags:       supportTags,
		localizationFiles: localizationFiles,
		translateFunc:     translateFunc,
		registerTagFunc:   registerTagFunc,
		defaultTranslator: defaultTranslator,
		localTranslator:   localTranslator,
	}
}

func (m *Middleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
			r2 := m.withRequest(r)
			next(w, r2)
	}
}

func (m *Middleware) isHasI18n() bool {
	return len(m.supportTags) > 0 && len(m.localizationFiles) > 0
}

func WithTranslateFunc(m *Middleware, translateFunc TranslateFunc) *Middleware {
		m.translateFunc = translateFunc
		return m
}