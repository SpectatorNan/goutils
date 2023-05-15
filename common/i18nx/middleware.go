package i18nx

import (
	"context"
	"github.com/SpectatorNan/go-zero-i18n/goi18nx"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	language2 "goutils/common/language"
	"net/http"
)

type Middleware struct {
	configs []string
}

func NewMiddleware(configs ...string) *Middleware {
	return &Middleware{
		configs: configs,
	}
}

func (m *Middleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tag := language2.GetLanguageTag(r)
		bundle := goi18nx.NewBundle(tag, m.configs...)
		//langMiddleware := goi18nx.NewI18nMiddleware(bundle)

		langTag := language2.GetLanguageTag(r)

		lang := r.FormValue("lang")
		accept := r.Header.Get("Accept-Language")
		localizer := i18n.NewLocalizer(bundle, lang, accept)
		ctx := context.WithValue(r.Context(), goi18nx.I18nKey, localizer)
		ctx = context.WithValue(ctx, AcceptLanguageTypeCtxKey, langTag.String())
		r2 := r.WithContext(ctx)

		next(w, r2)
	}
}
