package jwtx

import "github.com/golang-jwt/jwt/v4"

// Custom claims structure
type CustomClaims[Base BaseClaim] struct {
	BaseClaims Base
	BufferTime int64
	jwt.RegisteredClaims
}

func (c *CustomClaims[Base]) GetUserId() int64 {
	return c.BaseClaims.GetUserId()
}

type BaseClaim interface {
	GetUserId() int64
}
