package validator

import (
	"github.com/BurntSushi/toml"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
	"net/http"
	"os"
	"testing"
)

func TestLoadToml(t *testing.T) {
	unmarshalFuncMap := map[string]UnmarshalFunc{
		"toml": toml.Unmarshal,
	}
	path := "validate.en.toml"
	buf, err := os.ReadFile(path)
	if err != nil {
		t.Error(err)
		return
	}
	fInfo, err := parseFileBytes(buf, path, unmarshalFuncMap)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(fInfo)
}

func TestLoadYaml(t *testing.T) {
	unmarshalFuncMap := map[string]UnmarshalFunc{
		"yml": yaml.Unmarshal,
	}
	path := "validate.en.yml"
	buf, err := os.ReadFile(path)
	if err != nil {
		t.Error(err)
		return
	}
	fInfo, err := parseFileBytes(buf, path, unmarshalFuncMap)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(fInfo)
}

func TestParsePath(t *testing.T) {
	tag, format := parsePath("./validate.en.toml")
	t.Log(tag, format)
	tag, format = parsePath("./validate.en.yml")
	t.Log(tag, format)
}

func TestLoadYamlByBundle(t *testing.T) {
	bundle := NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("yml", yaml.Unmarshal)
	err := bundle.LoadFile("validate.en.yml")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestRegisterTranslation(t *testing.T) {
	bundle := NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("yml", yaml.Unmarshal)
	err := bundle.LoadFile("validate.en.yml")
	if err != nil {
		t.Error(err)
		return
	}

	validate := validator.New()
	r := &http.Request{}
	trans, _ := ut.New(en.New()).GetTranslator(language.English.String())
	for a, b := range bundle.messageTemplates[language.English] {
		t.Log(a, b)
		_ = validate.RegisterTranslation(a, trans, b.RegisterTranslation, TranslationFunc(r, bundle, b))
	}
}
