package services

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
)

type SavePostService struct {
	// repo repository.SavePostRepository
}

func NewSavePostService() *SavePostService {
	return &SavePostService{}
}

func (sp *SavePostService) Save(
	ctx context.Context, postID uint64, listID *dto.SavePostRequest,
) *service_errors.ServiceError {
	return nil
}

func (sp *SavePostService) Unsave(
	ctx context.Context, postID uint64, listID *dto.UnsavePostRequest,
) *service_errors.ServiceError {
	return nil
}
