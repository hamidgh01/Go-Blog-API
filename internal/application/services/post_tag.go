package services

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
)

type PostTagsService struct {
	repo repository.PostTagsRepository
}

func NewPostTagsService(r repository.PostTagsRepository) *PostTagsService {
	return &PostTagsService{repo: r}
}

func (pt *PostTagsService) Associate(
	ctx context.Context, postID uint64, tagIDs *dto.AssociatePostWithTagsRequest,
) *service_errors.ServiceError {
	return nil
}

func (pt *PostTagsService) Dissociate(
	ctx context.Context, postID uint64, tagIDs *dto.DissociatePostWithTagsRequest,
) *service_errors.ServiceError {
	return nil
}
