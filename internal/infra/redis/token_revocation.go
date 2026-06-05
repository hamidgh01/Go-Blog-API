package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type TokenRevoker struct {
	redis *redis.Client
}

func NewTokenRevoker(c *redis.Client) *TokenRevoker {
	return &TokenRevoker{redis: c}
}

const blacklistPrefix = "jwt-bl:"

func blacklistKey(jti string) string {
	return blacklistPrefix + jti
}

func (tr *TokenRevoker) Blacklist(
	ctx context.Context, jti string, userID uint64, expiration time.Time,
) error {
	remainingTTL := time.Until(expiration)
	if remainingTTL <= 0 { // Already expired
		return nil
	}

	key := blacklistKey(jti)

	// store userID for debugging
	err := tr.redis.Set(ctx, key, userID, remainingTTL).Err()
	if err != nil {
		return fmt.Errorf("failed to blacklist (redis.Set) token: %w", err)
	}

	return nil
}

func (tr *TokenRevoker) IsBlacklisted(ctx context.Context, jti string) (bool, error) {
	key := blacklistKey(jti)
	returnedInt, err := tr.redis.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("failed to check (redis.Exists) blacklist: %w", err)
	}

	return returnedInt > 0, nil
}
