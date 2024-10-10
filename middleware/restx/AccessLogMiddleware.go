package restx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/timex"
	"io"
	"net/http"
	"strings"
	"time"
)

// CustomResponseWriter is a custom implementation of http.ResponseWriter to capture response data
type CustomResponseWriter struct {
	http.ResponseWriter
	Body       *bytes.Buffer
	StatusCode int
}

func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *CustomResponseWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *CustomResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

type AccessLogMiddleware struct {
	timeOut           time.Duration
	logHeaderKeys     map[string]struct{}
	limitResponseSize int
}

type AccessLogOption func(m *AccessLogMiddleware)

func NewAccessLogMiddleware(timeOut int64, headerKeys []string, opts ...AccessLogOption) *AccessLogMiddleware {
	keys := make(map[string]struct{})
	for _, key := range headerKeys {
		keys[key] = struct{}{}
	}
	m := &AccessLogMiddleware{timeOut: time.Duration(timeOut) * time.Millisecond, logHeaderKeys: keys, limitResponseSize: 512}
	for _, opt := range opts {
		opt(m)
	}
	return m
}
func WithResponseLogSizeLimit(limit int) AccessLogOption {
	return func(m *AccessLogMiddleware) {
		m.limitResponseSize = limit
	}
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

		// Use custom response writer to capture response data
		crw := &CustomResponseWriter{ResponseWriter: w, Body: &bytes.Buffer{}, StatusCode: http.StatusOK}
		next(crw, r)
		//next(w, r)

		duration := timex.Since(start)
		logger := logx.WithContext(r.Context()).WithDuration(duration)
		builder := strings.Builder{}
		builder.WriteString("[Access] - %s")
		args := []interface{}{r.Method}
		builder.WriteString(" %s")
		args = append(args, r.RequestURI)
		if len(m.logHeaderKeys) > 0 {
			builder.WriteString(" - Headers: %s")
			hdMap := make(map[string]string)
			for key := range m.logHeaderKeys {
				//builder.WriteString(" - %s: %s")
				//args = append(args, key, r.Header.Get(key))
				hdMap[key] = r.Header.Get(key)
			}
			js, _ := json.Marshal(hdMap)
			args = append(args, js)
		}
		if queryParams != nil && len(queryParams) > 0 {
			builder.WriteString(" - Query: %s")
			js, _ := json.Marshal(queryParams)
			args = append(args, js)
		}
		if body != nil {
			builder.WriteString(" - Body: %s")
			args = append(args, body)
		}

		v := gjson.Get(crw.Body.String(), "data.page")
		fmt.Println(v)
		if v.Type == gjson.Null {
			builder.WriteString(" - Response: %s")
			args = append(args, crw.Body.String())
		} else if m.limitResponseSize > 0 && len(crw.Body.Bytes()) > m.limitResponseSize {
			builder.WriteString(" - Response: %s")
			args = append(args, crw.Body.String()[:m.limitResponseSize])
		} else {
			builder.WriteString(" - Response: %s")
			args = append(args, crw.Body.String())
		}

		if m.timeOut < duration {
			builder.WriteString(" - Timeout context deadline exceeded")
			logger.Errorf(builder.String(), args...)
		} else {

			logger.Infof(builder.String(), args...)
		}

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
