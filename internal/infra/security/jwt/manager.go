package jwt

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/hamidgh01/Go-Blog-API/config"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	signingAlgorithm jwt.SigningMethodHMAC
	accessSecret     []byte
	refreshSecret    []byte
	accessTTL        time.Duration
	refreshTTL       time.Duration
}

func NewJWTManager(cfg *config.JwtConf) *JWTManager {
	return &JWTManager{
		signingAlgorithm: *jwt.SigningMethodHS256,
		accessSecret:     []byte(cfg.AccessSecret),
		refreshSecret:    []byte(cfg.RefreshSecret),
		accessTTL:        time.Minute * time.Duration(cfg.AccessTokenExpirationMinutes),
		refreshTTL:       time.Hour * 24 * time.Duration(cfg.RefreshTokenExpirationDays),
	}
}

func generateJTI() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func (m *JWTManager) GenerateAccessToken(userID uint64) (accessToken string, err error) {
	claims := newClaims(userID, "", ACCESS, m.accessTTL)
	return createJWT(claims, m.accessSecret, &m.signingAlgorithm)
}

func (m *JWTManager) GenerateRefreshToken(userID uint64) (refreshToken string, err error) {
	jti, err := generateJTI()
	if err != nil {
		return "", fmt.Errorf("failed to generate jti: %w", err)
	}

	claims := newClaims(userID, jti, REFRESH, m.refreshTTL)
	return createJWT(claims, m.refreshSecret, &m.signingAlgorithm)
}

func (m *JWTManager) ParseAccessToken(tokenString string) (*Claims, error) {
	return parseJWT(tokenString, m.accessSecret, &m.signingAlgorithm)
}

func (m *JWTManager) ParseRefreshToken(tokenString string) (*Claims, error) {
	return parseJWT(tokenString, m.refreshSecret, &m.signingAlgorithm)
}
