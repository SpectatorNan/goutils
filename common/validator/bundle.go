package validator

import (
	"golang.org/x/text/language"
	"io/ioutil"
	"net/http"
)

// UnmarshalFunc unmarshals data into v.
type UnmarshalFunc func(data []byte, v interface{}) error
type VariableNameHandlerFunc func(r *http.Request, msgId, fieldName string) string

type Bundle struct {
	defaultLanguage         language.Tag
	unmarshalFuncs          map[string]UnmarshalFunc
	messageTemplates        map[language.Tag]map[string]*Translate
	tags                    []language.Tag
	matcher                 language.Matcher
	variableNameHandlerFunc VariableNameHandlerFunc
}

// NewBundle returns a bundle with a default language.
func NewBundle(defaultLanguage language.Tag) *Bundle {
	return &Bundle{
		defaultLanguage: defaultLanguage,
	}
}

func NewBundleWithTemplatePaths(defaultLanguage language.Tag, unmarshalFormat string, unmarshalFunc UnmarshalFunc, templatePaths ...string) *Bundle {
	bundle := NewBundle(defaultLanguage)
	bundle.RegisterUnmarshalFunc(unmarshalFormat, unmarshalFunc)
	for _, path := range templatePaths {
		if e := bundle.LoadFile(path); e != nil {
			panic(e)
		}
	}
	return bundle
}

// RegisterUnmarshalFunc registers an UnmarshalFunc for format.
func (b *Bundle) RegisterUnmarshalFunc(format string, unmarshalFunc UnmarshalFunc) {
	if b.unmarshalFuncs == nil {
		b.unmarshalFuncs = make(map[string]UnmarshalFunc)
	}
	b.unmarshalFuncs[format] = unmarshalFunc
}

// SetVariableNameHandlerFunc set the variable name handler function.
func (b *Bundle) SetVariableNameHandlerFunc(f VariableNameHandlerFunc) {
	b.variableNameHandlerFunc = f
}

func (b *Bundle) LoadFile(path string) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return b.ParseFileBytes(buf, path)
}

func (b *Bundle) ParseFileBytes(buf []byte, path string) error {
	fileInfo, err := parseFileBytes(buf, path, b.unmarshalFuncs)
	if err != nil {
		return err
	}
	if err = b.AddTemplate(fileInfo.Tag, fileInfo.Templates...); err != nil {
		return err
	}
	return nil
}

func (b *Bundle) AddTemplate(tag language.Tag, templates ...*Translate) error {
	if b.messageTemplates == nil {
		b.messageTemplates = make(map[language.Tag]map[string]*Translate)
	}
	if _, ok := b.messageTemplates[tag]; !ok {
		b.messageTemplates[tag] = make(map[string]*Translate)
	}
	for _, template := range templates {
		b.messageTemplates[tag][template.Tag] = template
	}
	b.addTag(tag)
	return nil
}

func (b *Bundle) addTag(tag language.Tag) {
	for _, t := range b.tags {
		if t == tag {
			// Tag already exists
			return
		}
	}
	b.tags = append(b.tags, tag)
	b.matcher = language.NewMatcher(b.tags)
}

func (b *Bundle) GetTemplateByLanguageTag(tag language.Tag) map[string]*Translate {
	return b.messageTemplates[tag]
}