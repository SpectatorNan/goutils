package cookiex

import "strings"


func parseCookie(cookieStr string) map[string]string {
	cookieMap := make(map[string]string)
	cookies := strings.Split(cookieStr, ";")
	for _, cookie := range cookies {
		parts := strings.Split(strings.TrimSpace(cookie), "=")
		if len(parts) == 2 {
			cookieMap[parts[0]] = parts[1]
		}
	}
	return cookieMap
}

func ParseCookie(cookieStr string) map[string]string {
	return parseCookie(cookieStr)
}

func GetCookieValue(cookieStr, key string) (string, bool) {
	cookieMap := parseCookie(cookieStr)
	val, ok := cookieMap[key]
	return val, ok
}
