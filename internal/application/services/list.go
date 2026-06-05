package services

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
	"github.com/hamidgh01/Go-Blog-API/internal/http/generics"
)

type ListService struct {
	repo repository.ListRepository
}

func NewListService(r repository.ListRepository) *ListService {
	return &ListService{repo: r}
}

func (c *ListService) Create(
	ctx context.Context, data *dto.CreateListRequest,
) (*dto.ListDetails, *service_errors.ServiceError) {
	return nil, nil
}

func (c *ListService) Update(
	ctx context.Context, pk uint64, data *dto.UpdateListRequest,
) (*dto.ListDetails, *service_errors.ServiceError) {
	return nil, nil
}

func (c *ListService) UpdatePrivacy(
	ctx context.Context, pk uint64, data *dto.UpdateListPrivacyRequest,
) (*dto.ListDetails, *service_errors.ServiceError) {
	return nil, nil
}

func (c *ListService) Delete(ctx context.Context, pk uint64) *service_errors.ServiceError {
	return nil
}

func (c *ListService) GetByID(
	ctx context.Context, pk uint64,
) (*dto.ListDetails, *service_errors.ServiceError) {
	return nil, nil
}

// -----------------------------------------------------------------------------
// other sources that has FK to `List`

func (c *ListService) GetSavedPosts(
	ctx context.Context, fk uint64,
) (*generics.PagedList[dto.PostsList], *service_errors.ServiceError) {
	return nil, nil
}

func (c *ListService) GetUsersWhoSaved(
	ctx context.Context, fk uint64,
) (*generics.PagedList[dto.UsersList], *service_errors.ServiceError) {
	return nil, nil
}
