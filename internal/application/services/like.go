package services

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
)

type LikeService struct {
	// repo repository.LikeRepository
}

func NewLikeService() *LikeService {
	return &LikeService{}
}

func (l *LikeService) Like(ctx context.Context, targetPostID uint64) *service_errors.ServiceError {
	return nil
}

func (l *LikeService) Unlike(ctx context.Context, targetPostID uint64) *service_errors.ServiceError {
	return nil
}
