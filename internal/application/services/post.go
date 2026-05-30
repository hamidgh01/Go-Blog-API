package services

import (
	"context"

	"Go-Blog-API/internal/application/service_errors"
	"Go-Blog-API/internal/domain/repository"
	"Go-Blog-API/internal/http/dto"
)

type PostService struct {
	repo repository.PostRepository
}

func NewPostService(r repository.PostRepository) *PostService {
	return &PostService{repo: r}
}

func (p *PostService) Create(
	ctx context.Context, data *dto.CreatePostRequest,
) (*dto.PostDetailsResponse, *service_errors.ServiceError) {
	return nil, nil
}

func (p *PostService) Update(
	ctx context.Context, pk uint64, data *dto.UpdatePostRequest,
) (*dto.PostDetailsResponse, *service_errors.ServiceError) {
	return nil, nil
}

func (p *PostService) UpdateStatus(
	ctx context.Context, pk uint64, data *dto.UpdatePostStatusRequest,
) (*dto.PostDetailsResponse, *service_errors.ServiceError) {
	return nil, nil
}

func (p *PostService) UpdatePrivacy(
	ctx context.Context, pk uint64, data *dto.UpdatePostPrivacyRequest,
) (*dto.PostDetailsResponse, *service_errors.ServiceError) {
	return nil, nil
}

func (p *PostService) Delete(ctx context.Context, pk uint64) *service_errors.ServiceError {
	return nil
}

func (p *PostService) GetByID(
	ctx context.Context, pk uint64,
) (*dto.PostDetailsResponse, *service_errors.ServiceError) {
	return nil, nil
}
