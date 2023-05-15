package validator

import (
	"context"
	"fmt"
	"github.com/SpectatorNan/go-zero-i18n/goi18nx"
	"net/http"
	"reflect"
	"strings"
)

func WithRequest(r *http.Request) *http.Request {

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
