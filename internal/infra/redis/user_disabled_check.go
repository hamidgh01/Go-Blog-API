package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type UserDisabledChecker struct {
	redis *redis.Client
}

func NewUserDisabledChecker(c *redis.Client) *UserDisabledChecker {
	return &UserDisabledChecker{redis: c}
}

const disabledUsersSetKey = "disabledUsersSet"

func (checker *UserDisabledChecker) SaveInDisabledUsers(ctx context.Context, userID uint64) error {
	err := checker.redis.SAdd(ctx, disabledUsersSetKey, userID).Err()
	if err != nil {
		return fmt.Errorf("failed to save (redis.SAdd) 'userID=%d' in disabledUsersSet: %w", userID, err)
	}

	return nil
}

func (checker *UserDisabledChecker) RemoveFromDisabledUsers(ctx context.Context, userID uint64) error {
	err := checker.redis.SRem(ctx, disabledUsersSetKey, userID).Err()
	if err != nil {
		return fmt.Errorf("failed to remove (redis.SRem) 'userID=%d' from disabledUsersSet: %w", userID, err)
	}

	return nil
}

func (checker *UserDisabledChecker) IsDisabled(ctx context.Context, userID uint64) (bool, error) {
	exists, err := checker.redis.SIsMember(ctx, disabledUsersSetKey, userID).Result()
	if err != nil {
		return false, fmt.Errorf("failed to check (redis.SIsMember) 'userID=%d' IsDisabled: %w", userID, err)
	}

	return exists, nil
}
