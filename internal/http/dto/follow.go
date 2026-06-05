package dto

import "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"

type FollowersList struct {
	Followers UsersList `json:"followers"`
}

func ToFollowersList(users []*entity.User) *FollowersList {
	return &FollowersList{Followers: ToUsersList(users)}
}

type FollowingsList struct {
	Followings UsersList `json:"followings"`
}

func ToFollowingsList(users []*entity.User) *FollowingsList {
	return &FollowingsList{Followings: ToUsersList(users)}
}
