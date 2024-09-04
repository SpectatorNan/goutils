package requestx

import (
	"fmt"
	"net/http"
	"strconv"
)

func SetHeaderUserId(r *http.Request, userId *int64) {
	if userId != nil {
		r.Header.Set("userId", fmt.Sprintf("%d", *userId))
	}
}
func GetHeaderUserId(r *http.Request) *int64 {
	idstr := r.Header.Get("userId")
	r.Header.Del("userId")
	if idstr == "" {
		return nil
	}
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return nil
	}
	return &id
}

func SetHeaderLogBody(r *http.Request, body *string) {
	if body != nil {
		r.Header.Set("logBody", *body)
	}
}
func GetHeaderLogBody(r *http.Request) string {
	body := r.Header.Get("logBody")
	r.Header.Del("logBody")
	return body
}
