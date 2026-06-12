package services

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
)

type LinkService struct {
	repo repository.LinkRepository
}

func NewLinkService(r repository.LinkRepository) *LinkService {
	return &LinkService{repo: r}
}

func (l *LinkService) Create(
	ctx context.Context, data *dto.CreateLinkRequest,
) (*dto.LinkDetails, *service_errors.ServiceError) {
	return nil, nil
}

func (l *LinkService) Update(
	ctx context.Context, pk uint64, data *dto.UpdateLinkRequest,
) *service_errors.ServiceError {
	return nil
}

func (l *LinkService) Delete(ctx context.Context, pk uint64) *service_errors.ServiceError {
	return nil
}

// func (u *LinkService) GetList() {}

func (l *LinkService) GetByID(
	ctx context.Context, pk uint64,
) (*dto.LinkDetails, *service_errors.ServiceError) {
	return nil, nil
}
