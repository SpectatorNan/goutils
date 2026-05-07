package respx

import "github.com/SpectatorNan/goutils/privacy"

// PageRespT is a generic page response that can participate in response desensitization.
type PageRespT[T any] struct {
	PageData
	Data            T `json:"list"`
	desensitizeData func(ctx privacy.ViewerContext, data T) T
}

func (p PageRespT[T]) Response() *Response {
	return NewSuccessResponse(p)
}

func (p PageRespT[T]) MakeDesensitize(ctx privacy.ViewerContext) any {
	out := p
	if out.desensitizeData == nil {
		out.desensitizeData = defaultPageRespDataDesensitizer[T]
	}
	out.Data = out.desensitizeData(ctx, p.Data)
	return out
}

func (p PageRespT[T]) DesensitizeType() privacy.DesensitizeType {
	return privacy.DesTypeObject
}

func NewPageRespT[T any](page, pageSize int, total int64, data T) *PageRespT[T] {
	pageData := createPaging(page, pageSize, total)
	return &PageRespT[T]{
		PageData:         pageData,
		Data:             data,
		desensitizeData: defaultPageRespDataDesensitizer[T],
	}
}

func NewSensitivePageRespT[T privacy.Desensitize](page, pageSize int, total int64, data T) *PageRespT[T] {
	resp := NewPageRespT(page, pageSize, total, data)
	resp.desensitizeData = MakeDesensitizeValue[T]
	return resp
}

func NewSensitivePageRespSlice[T privacy.Desensitize](page, pageSize int, total int64, data []T) *PageRespT[[]T] {
	resp := NewPageRespT(page, pageSize, total, data)
	resp.desensitizeData = MakeDesensitizeSlice[T]
	return resp
}

func NewSensitivePageRespPtrSlice[T privacy.Desensitize](page, pageSize int, total int64, data []*T) *PageRespT[[]*T] {
	resp := NewPageRespT(page, pageSize, total, data)
	resp.desensitizeData = MakeDesensitizePtrSlice[T]
	return resp
}

func defaultPageRespDataDesensitizer[T any](ctx privacy.ViewerContext, data T) T {
	if d, ok := any(data).(privacy.Desensitize); ok {
		if masked, ok := d.MakeDesensitize(ctx).(T); ok {
			return masked
		}
	}
	return data
}

func MakeDesensitizeValue[T privacy.Desensitize](ctx privacy.ViewerContext, data T) T {
	if masked, ok := data.MakeDesensitize(ctx).(T); ok {
		return masked
	}
	return data
}

func MakeDesensitizeSlice[T privacy.Desensitize](ctx privacy.ViewerContext, data []T) []T {
	if data == nil {
		return nil
	}
	out := make([]T, len(data))
	for i := range data {
		out[i] = MakeDesensitizeValue(ctx, data[i])
	}
	return out
}

func MakeDesensitizePtrSlice[T privacy.Desensitize](ctx privacy.ViewerContext, data []*T) []*T {
	if data == nil {
		return nil
	}
	out := make([]*T, len(data))
	for i, item := range data {
		if item == nil {
			continue
		}
		if masked, ok := (*item).MakeDesensitize(ctx).(T); ok {
			maskedCopy := masked
			out[i] = &maskedCopy
			continue
		}
		copied := *item
		out[i] = &copied
	}
	return out
}

