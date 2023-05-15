package jwtx

import "github.com/golang-jwt/jwt/v4"

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

func (c *CustomClaims) GetUserId() int64 {
	return c.BaseClaims.ID
}

type BaseClaims struct {
	ID          int64
	NickName    string
	UserName    string
	AuthorityId string
}
