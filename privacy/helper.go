package privacy

func MakeDesensitizeValue[T Desensitize](ctx ViewerContext, data T) T {
	if masked, ok := data.MakeDesensitize(ctx).(T); ok {
		return masked
	}
	return data
}

func MakeDesensitizeSlice[T Desensitize](ctx ViewerContext, data []T) []T {
	if data == nil {
		return nil
	}
	out := make([]T, len(data))
	for i := range data {
		out[i] = MakeDesensitizeValue(ctx, data[i])
	}
	return out
}

func MakeDesensitizePtrSlice[T Desensitize](ctx ViewerContext, data []*T) []*T {
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
