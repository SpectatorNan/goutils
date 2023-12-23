package validator

import (
	"encoding/json"
	"fmt"
	"golang.org/x/text/language"
	"os"
)

type FileInfo struct {
	Path      string
	Tag       language.Tag
	Format    string
	Templates []*Translate
}

func parsePath(path string) (langTag, format string) {
	formatStartIdx := -1
	for i := len(path) - 1; i >= 0; i-- {
		c := path[i]
		if os.IsPathSeparator(c) {
			if formatStartIdx != -1 {
				langTag = path[i+1 : formatStartIdx]
			}
			return
		}
		if path[i] == '.' {
			if formatStartIdx != -1 {
				langTag = path[i+1 : formatStartIdx]
				return
			}
			if formatStartIdx == -1 {
				format = path[i+1:]
				formatStartIdx = i
			}
		}
	}
	if formatStartIdx != -1 {
		langTag = path[:formatStartIdx]
	}
	return
}

func parseFileBytes(buf []byte, path string, unmarshalFuncs map[string]UnmarshalFunc) (*FileInfo, error) {
	lang, format := parsePath(path)
	tag := language.Make(lang)
	fileInfo := &FileInfo{
		Path:   path,
		Tag:    tag,
		Format: format,
	}
	if len(buf) == 0 {
		return fileInfo, nil
	}
	unmarshalFunc := unmarshalFuncs[fileInfo.Format]
	if unmarshalFunc == nil {
		if fileInfo.Format == "json" {
			unmarshalFunc = json.Unmarshal
		} else {
			return nil, fmt.Errorf("no unmarshaler registered for %s", fileInfo.Format)
		}
	}
	var err error
	var raw translateFile
	if err = unmarshalFunc(buf, &raw); err != nil {
		return nil, err
	}

	fileInfo.Templates = raw.TagTrans

	return fileInfo, nil
}
func (b *Bundle) matchCurrentLanguageTag(lang string) language.Tag {
	tags, _, err := language.ParseAcceptLanguage(lang)
	if err != nil {
		tags = []language.Tag{language.English}
	}
	_, i, _ := b.matcher.Match(tags...)
	return b.tags[i]
}
