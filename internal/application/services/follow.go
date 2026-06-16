package services

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
)

type FollowService struct {
	// repo repository.FollowRepository
}

func NewFollowService() *FollowService {
	return &FollowService{}
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
