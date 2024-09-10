package restx

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
)

func TestLog(t *testing.T) {

	ctx := context.Background()
	ctx = logx.ContextWithFields(ctx, logx.Field("level1", "access"))
	logx.WithContext(ctx).Info("test")
}
