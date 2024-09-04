package tools

import (
	"fmt"
	"strings"
)

func SpliceUrlDomainWithPath(domain, path string) string {
	if strings.HasPrefix(path, "http") {
		return path
	} else if strings.HasPrefix(path, "/") {
		return fmt.Sprintf("%s%s", domain, path)
	} else {
		return fmt.Sprintf("%s/%s", domain, path)
	}
}
