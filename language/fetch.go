package language

//func GetLanguageTag(r *http.Request) language.Tag {
//	accept := r.Header.Get("Accept-Language")
//	langTags, _, err := language.ParseAcceptLanguage(accept)
//	if err != nil {
//		langTags = []language.Tag{language.English}
//	}
//	tags := []language.Tag{
//		language.English,
//		language.Spanish,
//		language.Chinese,
//	}
//	var matcher = language.NewMatcher(tags)
//	_, i, _ := matcher.Match(langTags...)
//	//_, i := language.MatchStrings(matcher, langTag.String())
//	tag := tags[i]
//	return tag
//}
