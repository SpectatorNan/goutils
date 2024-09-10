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
	timeOut       time.Duration
	logHeaderKeys map[string]struct{}
}

func NewAccessLogMiddleware(timeOut int64, headerKeys []string) *AccessLogMiddleware {
	keys := make(map[string]struct{})
	for _, key := range headerKeys {
		keys[key] = struct{}{}
	}
	return &AccessLogMiddleware{timeOut: time.Duration(timeOut) * time.Millisecond, logHeaderKeys: keys}
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
			for key := range m.logHeaderKeys {
				builder.WriteString(" - %s: %s")
				args = append(args, key, r.Header.Get(key))
			}
		}
		if queryParams != nil {
			builder.WriteString(" - Query: %s")
			args = append(args, queryParams)
		}
		if body != nil {
			builder.WriteString(" - Body: %s")
			args = append(args, body)
		}

		// Add response data to log
		builder.WriteString(" - Response: %s")
		args = append(args, crw.Body.String())

		if m.timeOut < duration {
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
