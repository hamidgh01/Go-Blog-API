package services

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
)

type SaveListService struct {
	// repo repository.SaveListRepository
}

func NewSaveListService() *SaveListService {
	return &SaveListService{}
}

func (sl *SaveListService) Save(ctx context.Context, targetListID uint64) *service_errors.ServiceError {
	return nil
}

func (sl *SaveListService) Unsave(ctx context.Context, targetListID uint64) *service_errors.ServiceError {
	return nil
}
