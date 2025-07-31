package errorx

import (
	"context"
	"fmt"
	"github.com/SpectatorNan/goutils/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithStackResult(t *testing.T) {
	ctx := context.Background()
	err := errors.New("test error")
	newErr := errors.WithStackFrom(ErrResourceNotFound, err)
	errWithMsg := errors.WithMessage(newErr, fmt.Sprintf("original error: %v", err))
	derr := GrpcErrorWithDetails(ctx, err)

	assert.Equal(t, true, errors.Is(newErr, ErrResourceNotFound))
	assert.Equal(t, true, errors.Is(errWithMsg, ErrResourceNotFound))
	assert.Equal(t, true, errors.Is(errWithMsg, newErr))
	assert.Equal(t, false, errors.Is(derr, ErrResourceNotFound))

	errWithMsg1 := errors.WithMessage(ErrResourceNotFound, fmt.Sprintf("original error: %v", err))
	derr1 := GrpcErrorWithDetails(ctx, errWithMsg1)

	t.Logf("errWithMsg: %v", errWithMsg)
	t.Logf("derr: %+v", derr)
	t.Logf("errWithMsg1: %q", errWithMsg1)
	t.Logf("derr1: %v", derr1)
}

func TestGrpcErrorWithDetails_ErrorFromGrpcStatus(t *testing.T) {
	ctx := context.Background()
	origErr := errors.New("test grpc error")
	grpcErr := GrpcErrorWithDetails(ctx, origErr)
	extractedErr := ErrorFromGrpcStatus(grpcErr)

	assert.NotNil(t, grpcErr)
	assert.NotNil(t, extractedErr)
	assert.Contains(t, extractedErr.Error(), origErr.Error())
}

// ErrorFromGrpcStatus
// GrpcErrorWithDetails
// httpResult
