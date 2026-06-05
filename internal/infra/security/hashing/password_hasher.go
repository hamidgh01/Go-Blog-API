package hashing

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher struct {
	cost int
}

func NewPasswordHasher(cost int) *PasswordHasher {
	return &PasswordHasher{cost: bcrypt.DefaultCost}
}

func (b *PasswordHasher) Hash(plain string) (string, error) {
	if plain == "" {
		return "", errors.New("password cannot be empty")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), b.cost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashed), nil
}

func (b *PasswordHasher) Verify(hashed, plain string) error {
	if hashed == "" || plain == "" {
		return errors.New("invalid password comparison")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	if err != nil {
		return errors.New("invalid credentials")
	}

	return nil
}
