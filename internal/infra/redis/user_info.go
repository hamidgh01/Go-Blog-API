package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type UserInfo struct {
	Username    string `redis:"u"`
	IsSuperuser string `redis:"s"`
	Enabled     string `redis:"e"`
}

type UserInfoCache struct {
	redis *redis.Client
}

func NewUserInfoCache(c *redis.Client) *UserInfoCache {
	return &UserInfoCache{redis: c}
}

const userInfoPrefix = "user-info"

func userInfoKey(userID uint64) string {
	return fmt.Sprintf("%s:%d", userInfoPrefix, userID)
}

func (u *UserInfoCache) SetAllInfo(
	ctx context.Context, userID uint64, username string, isSuperuser bool, enabled bool,
) error {
	userInfo := UserInfo{Username: username}
	if isSuperuser {
		userInfo.IsSuperuser = "t"
	} else {
		userInfo.IsSuperuser = "f"
	}

	if enabled {
		userInfo.Enabled = "t"
	} else {
		userInfo.Enabled = "f"
	}

	if err := u.redis.HSet(ctx, userInfoKey(userID), userInfo).Err(); err != nil {
		return fmt.Errorf("failed to set (redis.HSet) user info (userID=%d). origin: %w", userID, err)
	}

	return nil
}

func (u *UserInfoCache) UpdateUsername(ctx context.Context, userID uint64, username string) error {
	if err := u.redis.HSet(ctx, userInfoKey(userID), "u", username).Err(); err != nil {
		return fmt.Errorf("failed to update (redis.HSet) username (userID=%d). origin: %w", userID, err)
	}

	return nil
}

func (u *UserInfoCache) UpdateEnabled(ctx context.Context, userID uint64, enabled bool) error {
	var err error
	if enabled {
		err = u.redis.HSet(ctx, userInfoKey(userID), "e", "t").Err()
	} else {
		err = u.redis.HSet(ctx, userInfoKey(userID), "e", "f").Err()
	}

	if err != nil {
		return fmt.Errorf("failed to update (redis.HSet) enabled (userID=%d). origin: %w", userID, err)
	}

	return nil
}

func (u *UserInfoCache) UpdateIsSuperuser(ctx context.Context, userID uint64, isSuperuser bool) error {
	var err error
	if isSuperuser {
		err = u.redis.HSet(ctx, userInfoKey(userID), "s", "t").Err()
	} else {
		err = u.redis.HSet(ctx, userInfoKey(userID), "s", "f").Err()
	}

	if err != nil {
		return fmt.Errorf("failed to update (redis.HSet) isSuperuser (userID=%d). origin: %w", userID, err)
	}

	return nil
}

func (u *UserInfoCache) GetAllInfo(ctx context.Context, userID uint64) (UserInfo, error) {
	var userInfo UserInfo
	if err := u.redis.HGetAll(ctx, userInfoKey(userID)).Scan(&userInfo); err != nil {
		return UserInfo{}, fmt.Errorf("failed to get (redis.HGetAll) user info (userID=%d). origin: %w", userID, err)
	}

	return userInfo, nil
}

func (u *UserInfoCache) DeleteUserInfo(ctx context.Context, userID uint64) error {
	if err := u.redis.Del(ctx, userInfoKey(userID)).Err(); err != nil {
		return fmt.Errorf("failed to delete (redis.Del) user info from cache (userID=%d). origin: %w", userID, err)
	}

	return nil
}
