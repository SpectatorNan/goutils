package viewerx

import (
	"context"
	"testing"

	"github.com/SpectatorNan/goutils/privacy"
)

func TestViewerContextFromContext(t *testing.T) {
	ctx := WithViewerContext(context.Background(), privacy.ViewerContext{Level: privacy.MaskPublic})

	viewer, ok := ViewerContextFromContext(ctx)
	if !ok {
		t.Fatalf("expected viewer context to exist")
	}
	if viewer.Level != privacy.MaskPublic {
		t.Fatalf("unexpected viewer context: %+v", viewer)
	}
}

func TestViewerContextFromContext_Miss(t *testing.T) {
	_, ok := ViewerContextFromContext(context.Background())
	if ok {
		t.Fatalf("expected missing viewer context")
	}
}
