package services

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
)

type FollowService struct {
	repo repository.FollowRepository
}

func NewFollowService(r repository.FollowRepository) *FollowService {
	return &FollowService{repo: r}
}

func (f *FollowService) Follow(ctx context.Context, targetUserID uint64) *service_errors.ServiceError {
	return nil
}

func (f *FollowService) UnFollow(ctx context.Context, targetUserID uint64) *service_errors.ServiceError {
	return nil
}

func (f *FollowService) RemoveFollower(ctx context.Context, targetUserID uint64) *service_errors.ServiceError {
	return nil
}
