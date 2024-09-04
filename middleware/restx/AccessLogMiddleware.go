package restx

import (
	"bytes"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/timex"
	"io"
	"net/http"
	"strings"
	"time"
)

type AccessLogMiddleware struct {
	timeOut time.Duration
}

func NewAccessLogMiddleware(timeOut int64) *AccessLogMiddleware {
	return &AccessLogMiddleware{timeOut: time.Duration(timeOut) * time.Millisecond}
}

func (m *AccessLogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := timex.Now()

		// Extract query parameters for GET requests
		queryParams := r.URL.Query()

		// Extract body for non-GET requests
		var body []byte
		if r.Method != http.MethodGet {
			var err error
			body, err = io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Failed to read request body", http.StatusInternalServerError)
				return
			}
			// Restore the body for further use
			r.Body = io.NopCloser(bytes.NewBuffer(body))

			str, err := minifyJSON(string(body))
			if err == nil {
				body = []byte(str)
			}
		}

		next(w, r)

		duration := timex.Since(start)
		logger := logx.WithContext(r.Context()).WithDuration(duration)
		builder := strings.Builder{}
		builder.WriteString("Request [%s]")
		args := []interface{}{r.Method}
		builder.WriteString(" - %s")
		args = append(args, r.RequestURI)
		if queryParams != nil {
			builder.WriteString(" - %s")
			args = append(args, queryParams)
		}
		if body != nil {
			builder.WriteString(" - %s")
			args = append(args, body)
		}
		if m.timeOut > duration {
			builder.WriteString(" - Timeout context deadline exceeded")
			logger.Errorf(builder.String(), args...)
		} else {

			logger.Infof(builder.String(), args...)
		}

		//logger.Infof("Access - [%s] - %s - %s - %s - %s", r.Method, r.RequestURI, pathParams, queryParams, body)
	}
}

// MinifyJSON minifies a JSON string by removing all unnecessary whitespace.
func minifyJSON(data string) (string, error) {
	var buf bytes.Buffer
	if err := json.Compact(&buf, []byte(data)); err != nil {
		return "", err
	}
	return buf.String(), nil
}
