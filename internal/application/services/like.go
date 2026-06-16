package services

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
)

type LikeService struct {
	repo repository.LikeRepository
}

func NewLikeService(r repository.LikeRepository) *LikeService {
	return &LikeService{repo: r}
}

func (l *LikeService) Like(ctx context.Context, targetPostID uint64) *service_errors.ServiceError {
	return nil
}

func (l *LikeService) Unlike(ctx context.Context, targetPostID uint64) *service_errors.ServiceError {
	return nil
}
