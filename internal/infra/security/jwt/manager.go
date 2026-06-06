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

func (m *JWTManager) GenerateAccessToken(userID uint64) (string, time.Time, error) {
	claims, expirationTime := newClaims(userID, "", ACCESS, m.accessTTL)
	accessToken, err := createJWT(claims, m.accessSecret, &m.signingAlgorithm)

	return accessToken, expirationTime, err
}

func (m *JWTManager) GenerateRefreshToken(userID uint64) (string, time.Time, error) {
	jti, err := generateJTI()
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to generate jti: %w", err)
	}

	claims, expirationTime := newClaims(userID, jti, REFRESH, m.refreshTTL)
	refreshToken, err := createJWT(claims, m.refreshSecret, &m.signingAlgorithm)

	return refreshToken, expirationTime, err
}

func (m *JWTManager) ParseAccessToken(tokenString string) (*Claims, error) {
	return parseJWT(tokenString, m.accessSecret, &m.signingAlgorithm)
}

func (m *JWTManager) ParseRefreshToken(tokenString string) (*Claims, error) {
	return parseJWT(tokenString, m.refreshSecret, &m.signingAlgorithm)
}
