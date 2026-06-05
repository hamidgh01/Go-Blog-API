package jwt

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func createJWT(claims *Claims, secretKey []byte, signingAlg *jwt.SigningMethodHMAC) (string, error) {
	token := jwt.NewWithClaims(signingAlg, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}

func parseJWT(tokenStr string, secretKey []byte, expectedSigningAlg *jwt.SigningMethodHMAC) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&Claims{},
		func(token *jwt.Token) (any, error) {
			value, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok || (value.Name != expectedSigningAlg.Name) || (value.Hash != expectedSigningAlg.Hash) {
				return nil, errors.New("unexpected signing method")
			}
			return secretKey, nil
		},
	)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
