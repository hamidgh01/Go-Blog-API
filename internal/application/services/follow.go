package services

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/pkg/constants"
)

type FollowService struct {
	repo repository.FollowRepository
}

func NewFollowService(r repository.FollowRepository) *FollowService {
	return &FollowService{repo: r}
}

func (f *FollowService) Follow(ctx context.Context, targetUserID uint64) *service_errors.ServiceError {
	currentUserID := ctx.Value(constants.CurrentUserID).(uint64)
	entity := &entity.FollowsM2M{FollowedBy: currentUserID, Followed: targetUserID}

	// ToDo: condition to check:
	// check the target user (to follow) is enabled

	return createOrDeleteM2MRelationship(ctx, "follow a user", entity, f.repo.Create)
}

func (f *FollowService) UnFollow(ctx context.Context, targetUserID uint64) *service_errors.ServiceError {
	currentUserID := ctx.Value(constants.CurrentUserID).(uint64)
	entity := &entity.FollowsM2M{FollowedBy: currentUserID, Followed: targetUserID}

	return createOrDeleteM2MRelationship(ctx, "unfollow a user", entity, f.repo.Delete)
}

func (f *FollowService) RemoveFollower(ctx context.Context, targetUserID uint64) *service_errors.ServiceError {
	currentUserID := ctx.Value(constants.CurrentUserID).(uint64)
	entity := &entity.FollowsM2M{FollowedBy: targetUserID, Followed: currentUserID}

	return createOrDeleteM2MRelationship(ctx, "remove a follower", entity, f.repo.Delete)
}
