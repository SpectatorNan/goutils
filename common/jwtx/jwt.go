package jwtx

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWT struct {
	//Secret     string
	bufferTime int64
	issuer     string
	effectTime int64
}

func NewJWT(secret string, bufferTime int64, issuer string) *JWT {
	return &JWT{
		//Secret: secret,
		bufferTime: bufferTime, issuer: issuer}
}

func NewJWTWithConfig(c Config) *JWT {
	return &JWT{
		//Secret: c.AccessSecret,
		bufferTime: c.BufferTime, issuer: c.Issuer, effectTime: c.AccessExpire}
}

func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	now := time.Now()
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: j.bufferTime,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(j.effectTime) * time.Second)),
			Issuer:    j.issuer,
		},
	}
	return claims
}

func (j *JWT) CreateToken(claims CustomClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
