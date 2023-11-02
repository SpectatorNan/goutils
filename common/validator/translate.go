package validator

import (
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type translateFile struct {
	TagTrans []*Translate `yaml:"TagTrans"`
}

type Translate struct {
	Tag       string      `yaml:"Tag"`
	Template  string      `yaml:"Template"`
	Params []TranslateParams `yaml:"Params"`
}
type TranslateParams struct {
	Prefix  string `yaml:"Prefix"`
	ObjName string `yaml:"ObjName"`
}


func (t *Translate) RegisterTranslation(ut ut.Translator) error {
	return ut.Add(t.Tag, t.Template, true)
}

func TranslationFunc(r *http.Request, b *Bundle, t *Translate) validator.TranslationFunc {
	return func(ut ut.Translator, fe validator.FieldError) string {
		var variables []string
		for _, v := range t.Params {
			msgId, fieldName := fetchMsgId(fe, v)
			if b.variableNameHandlerFunc != nil {
				variables = append(variables, b.variableNameHandlerFunc(r, msgId, fieldName))
			} else {
				variables = append(variables, fieldName)
			}
		}
		result, err := ut.T(t.Tag, variables...)
		if err != nil {
			return fe.(error).Error()
		}
		return result
	}
}

func fetchMsgId(fe validator.FieldError, val TranslateParams) (string, string) {
	fieldName := fetchFieldName(fe, val)
	if len(val.Prefix) > 0 {
		return fmt.Sprintf("%s.%s", val.Prefix, fieldName), fieldName
	} else {
		return fieldName, fieldName
	}
}

func fetchFieldName(fe validator.FieldError, val TranslateParams) string {
	if val.ObjName == "StructField" {
		return fe.StructField()
	} else if val.ObjName == "Param" {
		return fe.Param()
	} else {
		return fe.Field()
	}
}