package services

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/pkg/constants"
)

type LikeService struct {
	repo repository.LikeRepository
}

func NewLikeService(r repository.LikeRepository) *LikeService {
	return &LikeService{repo: r}
}

func (l *LikeService) Like(ctx context.Context, targetPostID uint64) *service_errors.ServiceError {
	currentUserID := ctx.Value(constants.CurrentUserID).(uint64)
	entity := &entity.PostLikesM2M{PostID: targetPostID, UserID: currentUserID}

	// ToDo: condition to check:
	// check the target post (to like) is 'published' & is not private

	return createOrDeleteM2MRelationship(ctx, "like a post", entity, l.repo.Create)
}

func (l *LikeService) Unlike(ctx context.Context, targetPostID uint64) *service_errors.ServiceError {
	currentUserID := ctx.Value(constants.CurrentUserID).(uint64)
	entity := &entity.PostLikesM2M{PostID: targetPostID, UserID: currentUserID}

	return createOrDeleteM2MRelationship(ctx, "unlike a post", entity, l.repo.Delete)
}
