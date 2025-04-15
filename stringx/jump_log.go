package stringx

import "strings"

func JumpLogVisitPath(visitDomain, visitUri string) string {
	split := strings.Split(visitUri, "?")
	visitPath := strings.Replace(split[0], visitDomain, "", 1)
	return visitPath
}
