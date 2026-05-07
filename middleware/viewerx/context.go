package viewerx

import (
	"context"

	"github.com/SpectatorNan/goutils/privacy"
)

type contextKey struct{}

// WithViewerContext stores privacy.ViewerContext into request context.
func WithViewerContext(ctx context.Context, viewer privacy.ViewerContext) context.Context {
	return context.WithValue(ctx, contextKey{}, viewer)
}

// ViewerContextFromContext fetches privacy.ViewerContext from context.
func ViewerContextFromContext(ctx context.Context) (privacy.ViewerContext, bool) {
	viewer, ok := ctx.Value(contextKey{}).(privacy.ViewerContext)
	return viewer, ok
}

