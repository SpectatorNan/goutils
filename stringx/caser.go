package stringx

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	titleCaser = cases.Title(language.English)
	lowerCaser = cases.Lower(language.English)
	upperCaser = cases.Upper(language.English)
)

func Title(s string) string {
	return titleCaser.String(s)
}

func Lower(s string) string {
	return lowerCaser.String(s)
}

func Upper(s string) string {
	return upperCaser.String(s)
}