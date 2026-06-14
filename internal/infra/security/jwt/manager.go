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

type TokenPairDetails struct {
	AccessToken    string
	AccessExpTime  time.Time
	RefreshToken   string
	RefreshExpTime time.Time
}

func (m *JWTManager) GenerateTokenPair(userID uint64) (*TokenPairDetails, error) {
	accessClaims, accessExpTime := newClaims(userID, "", ACCESS, m.accessTTL)
	accessToken, err := createJWT(accessClaims, m.accessSecret, &m.signingAlgorithm)
	if err != nil {
		return nil, err
	}

	jti, err := generateJTI()
	if err != nil {
		return nil, fmt.Errorf("failed to generate jti: %w", err)
	}

	refreshClaims, refreshExpTime := newClaims(userID, jti, REFRESH, m.refreshTTL)
	refreshToken, err := createJWT(refreshClaims, m.refreshSecret, &m.signingAlgorithm)
	if err != nil {
		return nil, err
	}

	return &TokenPairDetails{
		AccessToken:    accessToken,
		AccessExpTime:  accessExpTime,
		RefreshToken:   refreshToken,
		RefreshExpTime: refreshExpTime,
	}, nil
}

func (m *JWTManager) ParseAccessToken(tokenString string) (*Claims, error) {
	return parseJWT(tokenString, m.accessSecret, &m.signingAlgorithm)
}

func (m *JWTManager) ParseRefreshToken(tokenString string) (*Claims, error) {
	return parseJWT(tokenString, m.refreshSecret, &m.signingAlgorithm)
}
