package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type SuperUserChecker struct {
	redis *redis.Client
}

func NewSuperUserChecker(c *redis.Client) *SuperUserChecker {
	return &SuperUserChecker{redis: c}
}

const superUsersSetKey = "superUsersSet"

func (checker *SuperUserChecker) SaveInSuperUsers(ctx context.Context, userID uint64) error {
	err := checker.redis.SAdd(ctx, superUsersSetKey, userID).Err()
	if err != nil {
		return fmt.Errorf("failed to save (redis.SAdd) 'userID=%d' in superUsersSet: %w", userID, err)
	}

	return nil
}

func (checker *SuperUserChecker) RemoveFromSuperUsers(ctx context.Context, userID uint64) error {
	err := checker.redis.SRem(ctx, superUsersSetKey, userID).Err()
	if err != nil {
		return fmt.Errorf("failed to remove (redis.SRem) 'userID=%d' from superUsersSet: %w", userID, err)
	}

	return nil
}

func (checker *SuperUserChecker) IsSuperUsers(ctx context.Context, userID uint64) (bool, error) {
	exists, err := checker.redis.SIsMember(ctx, superUsersSetKey, userID).Result()
	if err != nil {
		return false, fmt.Errorf("failed to check (redis.SIsMember) 'userID=%d' IsSuperUsers: %w", userID, err)
	}

	return exists, nil
}
