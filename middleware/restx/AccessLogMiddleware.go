package restx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/timex"
	//"github.com/zeromicro/go-zero/rest/httpx"
	"io"
	"mime"
	"mime/multipart"
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

var defaultLogHeaderKeys = map[string]struct{}{
	"User-Agent": struct{}{},
}

func NewAccessLogMiddleware(timeOut int64, headerKeys []string, opts ...AccessLogOption) *AccessLogMiddleware {
	keys := defaultLogHeaderKeys
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
		body := dealBody(r)

		done := make(chan struct{})

		// Use custom response writer to capture response data
		crw := &CustomResponseWriter{ResponseWriter: w, Body: &bytes.Buffer{}, StatusCode: http.StatusOK}
		//next(crw, r)
		//next(w, r)

		go func() {
			defer close(done)
			next(crw, r)
		}()

		// Wait for completion or timeout
		var timedOut bool
		select {
		case <-done:
			// Handler completed normally
		case <-r.Context().Done():
			// Timeout occurred
			timedOut = true
			crw.StatusCode = http.StatusGatewayTimeout
			//http.Error(w, "Request timed out", http.StatusGatewayTimeout)
			//httpx.ErrorCtx(r.Context(), crw, r.Context().Err())
		}

		duration := timex.Since(start)
		logger := logx.WithContext(r.Context()).WithDuration(duration)
		builder := strings.Builder{}

		builder.WriteString(" %s ")
		args := []interface{}{r.Host}

		builder.WriteString("- %s %d")
		args = append(args, r.Method, crw.StatusCode)
		builder.WriteString(" %s")
		args = append(args, r.RequestURI)
		if len(m.logHeaderKeys) > 0 {
			builder.WriteString(" - Headers: %s")
			hdMap := make(map[string]string)
			for key := range m.logHeaderKeys {
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
		builder.WriteString(" - Response: %s")
		if m.limitResponseSize > 0 && len(crw.Body.Bytes()) > m.limitResponseSize {
			args = append(args, crw.Body.String()[:m.limitResponseSize])
		} else {
			args = append(args, crw.Body.String())
		}
		//v := gjson.Get(crw.Body.String(), "data.page")

		//if v.Type == gjson.Null {
		//	builder.WriteString(" - Response: %s")
		//	args = append(args, crw.Body.String())
		//} else if m.limitResponseSize > 0 && len(crw.Body.Bytes()) > m.limitResponseSize {
		//	builder.WriteString(" - Response: %s")
		//	args = append(args, crw.Body.String()[:m.limitResponseSize])
		//} else {
		//	builder.WriteString(" - Response: %s")
		//	args = append(args, crw.Body.String())
		//}

		// 单独超时时间会可能会有错误的超时日志。比如上传文件
		//if m.timeOut < duration {
		if timedOut {
			logger.Errorf("[Access] - Timeout context deadline exceeded "+builder.String(), args...)
		} else {
			logger.Infof("[Access] "+builder.String(), args...)
		}

	}
}

func dealBody(r *http.Request) []byte {
	var body []byte
	if r.Method != http.MethodGet {
		var err error
		body, err = io.ReadAll(r.Body)
		if err != nil {
			m := map[string]string{
				"dumpErr": "Failed to read request body",
			}
			b, _ := json.Marshal(m)
			return b
		}
		// Restore the body for further use
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		body = dealUploadBody(r, body)

		str, err := minifyJSON(string(body))
		if err == nil {
			body = []byte(str)
		}
	}
	return body
}

func dealUploadBody(r *http.Request, body []byte) []byte {
	if strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
		bodyMap := make(map[string]any)

		// Extract boundary from Content-Type header
		contentType := r.Header.Get("Content-Type")
		_, params, err := mime.ParseMediaType(contentType)
		if err != nil {
			bodyMap["dumpErr"] = fmt.Sprintf("Failed to parse Content-Type header, val: %s", contentType)
			b, _ := json.Marshal(bodyMap)
			return b
		}
		boundary, ok := params["boundary"]
		if !ok {
			bodyMap["dumpErr"] = fmt.Sprintf("Boundary not found in Content-Type header, content-type: %s", contentType)
			b, _ := json.Marshal(bodyMap)
			return b
		}
		// Parse multipart form
		reader := multipart.NewReader(bytes.NewReader(body), boundary)
		for {
			part, err := reader.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				bodyMap["parseErr"] = fmt.Sprintf("Failed to parse multipart form")
				b, _ := json.Marshal(bodyMap)
				return b
			}
			if part.FileName() != "" {
				file := map[string]string{
					"field":     part.FormName(),
					"file_name": part.FileName(),
					"file_type": part.Header.Get("Content-Type"),
				}

				if files, exists := bodyMap["files"].([]map[string]string); exists {
					files = append(files, file)
					bodyMap["files"] = files
				} else {
					files := []map[string]string{file}
					bodyMap["files"] = files
				}
			} else {
				fieldValue, err := io.ReadAll(part)
				if err != nil {
					bodyMap["parseErr"] = fmt.Sprintf("Failed to read form field: %s", part.FormName())
					b, _ := json.Marshal(bodyMap)
					return b
				}
				bodyMap[part.FormName()] = string(fieldValue)
			}

		}
		b, _ := json.Marshal(bodyMap)
		return b
	}
	return body
}

// MinifyJSON minifies a JSON string by removing all unnecessary whitespace.
func minifyJSON(data string) (string, error) {
	var buf bytes.Buffer
	if err := json.Compact(&buf, []byte(data)); err != nil {
		return "", err
	}
	return buf.String(), nil
}
