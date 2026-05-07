package respx

import (
	"context"

	"github.com/SpectatorNan/goutils/middleware/viewerx"
	"github.com/SpectatorNan/goutils/privacy"
)

func desensitizeResp(ctx context.Context, resp any) any {
	if resp == nil {
		return nil
	}

	viewer, ok := viewerx.ViewerContextFromContext(ctx)
	if !ok {
		return resp
	}

	if d, ok := resp.(privacy.Desensitize); ok {
		return d.MakeDesensitize(viewer)
	}

	return resp
}
