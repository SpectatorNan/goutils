package validator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtrans "github.com/go-playground/validator/v10/translations/zh"
	"golang.org/x/text/language"
	"reflect"
	"strings"
	"testing"
)

type Student struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"email"`
	Age   int    `json:"age" validate:"max=30,min=12"`
}

func TestValidateErr(t *testing.T) {
	en := en.New() //英文翻译器
	zh := zh.New() //中文翻译器

	// 第一个参数是必填，如果没有其他的语言设置，就用这第一个
	// 后面的参数是支持多语言环境（
	// uni := ut.New(en, en) 也是可以的
	// uni := ut.New(en, zh, tw)
	uni := ut.New(en, zh)
	locale := language.Chinese.String()
	trans, ok := uni.GetTranslator(locale) //获取需要的语言
	if !ok {
		t.Errorf("uni.GetTranslator(%s) failed", locale)
		return
	}
	student := Student{
		Name:  "tom",
		Email: "testemal",
		Age:   40,
	}
	validate := validator.New()
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	zhtrans.RegisterDefaultTranslations(validate, trans)
	//entrans.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(student)
	if err != nil {
		// fmt.Println(err)

		errs := err.(validator.ValidationErrors)
		t.Log(removeStructName(errs.Translate(trans)))
	}
}

func TestGetLanguage1(t *testing.T) {
	accept := "zh-CN,zh;q=0.9"
	tag, q, err := language.ParseAcceptLanguage(accept)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(q)
	t.Log(tag)
	t.Log(language.Chinese.String())
	t.Log(language.SimplifiedChinese.String())
	t.Log(language.TraditionalChinese.String())
	for _, tag := range tag {
		t.Log(tag.String())
		switch tag {
		case language.Chinese:
			t.Log("China")
		case language.SimplifiedChinese:
			t.Log("China Simplified")
		case language.TraditionalChinese:
			t.Log("traditional china")
		default:
			t.Log("Other")
		}
	}

}

func TestGetLanguage2(t *testing.T) {
	accept := "zh-CN,zh;q=0.9"
	lang := language.Make(accept)
	var matcher = language.NewMatcher([]language.Tag{
		language.English,
		language.Spanish,
		language.Chinese,
	})
	tag, idx := language.MatchStrings(matcher, lang.String())
	t.Log(tag.String())
	switch tag {
	case language.Chinese:
		t.Log("China")
	case language.SimplifiedChinese:
		t.Log("China Simplified")
	case language.TraditionalChinese:
		t.Log("traditional china")
	default:
		t.Log("Other")
	}
	t.Log(idx)
}
