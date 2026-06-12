package services

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	d "github.com/hamidgh01/Go-Blog-API/internal/domain"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
	"github.com/hamidgh01/Go-Blog-API/internal/http/generics"
)

type CommentService struct {
	repo repository.CommentRepository
}

func NewCommentService(r repository.CommentRepository) *CommentService {
	return &CommentService{repo: r}
}

func (c *CommentService) Create(
	ctx context.Context, data *dto.CreateCommentRequest,
) (*dto.CommentDetails, *service_errors.ServiceError) {
	return nil, nil
}

func (c *CommentService) Update(
	ctx context.Context, pk uint64, data *dto.UpdateCommentRequest,
) *service_errors.ServiceError {
	return nil
}

func (c *CommentService) UpdateStatus(
	ctx context.Context, pk uint64, data *dto.UpdateCommentStatusRequest,
) *service_errors.ServiceError {
	return nil
}

func (c *CommentService) Delete(ctx context.Context, pk uint64) *service_errors.ServiceError {
	return nil
}

func (c *CommentService) GetByID(
	ctx context.Context, pk uint64,
) (*dto.CommentDetails, *service_errors.ServiceError) {
	return nil, nil
}

// -----------------------------------------------------------------------------
// other sources that has FK to `Comment`

func (c *CommentService) GetReplies(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.CommentList], *service_errors.ServiceError) {
	return nil, nil
}
