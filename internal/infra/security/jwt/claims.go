package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenType uint8

const (
	ACCESS TokenType = iota
	REFRESH
)

type Claims struct {
	jwt.RegisteredClaims
	UserID uint64    `json:"user_id"`
	Type   TokenType `json:"type"`
}

func newClaims(userID uint64, jti string, tokenType TokenType, tokenTTL time.Duration) (*Claims, time.Time) {
	expirationTime := time.Now().Add(tokenTTL)

	return &Claims{
		UserID: userID,
		Type:   tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			ID:        jti,
		},
	}, expirationTime
}

func (c *Claims) GetJTI() string {
	return c.ID
}

func (c *Claims) GetUserID() uint64 {
	return c.UserID
}

func (c *Claims) GetExpiresAt() time.Time {
	return c.ExpiresAt.Time
}

func (c *Claims) GetTokenType() TokenType {
	return c.Type
}
