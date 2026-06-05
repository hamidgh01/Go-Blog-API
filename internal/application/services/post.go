package services

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
)

type PostService struct {
	repo repository.PostRepository
}

func NewPostService(r repository.PostRepository) *PostService {
	return &PostService{repo: r}
}

func (p *PostService) Create(
	ctx context.Context, data *dto.CreatePostRequest,
) (*dto.PostDetails, *service_errors.ServiceError) {
	return nil, nil
}

func (p *PostService) Update(
	ctx context.Context, pk uint64, data *dto.UpdatePostRequest,
) (*dto.PostDetails, *service_errors.ServiceError) {
	return nil, nil
}

func (p *PostService) UpdateStatus(
	ctx context.Context, pk uint64, data *dto.UpdatePostStatusRequest,
) (*dto.PostDetails, *service_errors.ServiceError) {
	return nil, nil
}

func (p *PostService) UpdatePrivacy(
	ctx context.Context, pk uint64, data *dto.UpdatePostPrivacyRequest,
) (*dto.PostDetails, *service_errors.ServiceError) {
	return nil, nil
}

func (p *PostService) Delete(ctx context.Context, pk uint64) *service_errors.ServiceError {
	return nil
}

func (p *PostService) GetByID(
	ctx context.Context, pk uint64,
) (*dto.PostDetails, *service_errors.ServiceError) {
	return nil, nil
}
