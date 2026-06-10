package services

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	d "github.com/hamidgh01/Go-Blog-API/internal/domain"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
	"github.com/hamidgh01/Go-Blog-API/internal/http/generics"
)

type TagService struct {
	repo repository.TagRepository
}

func NewTagService(r repository.TagRepository) *TagService {
	return &TagService{repo: r}
}

func (t *TagService) Create(
	ctx context.Context, data *dto.CreateTagsRequest,
) (*dto.TagDetails, *service_errors.ServiceError) {
	return nil, nil
}

func (t *TagService) GetByID(
	ctx context.Context, pk uint64,
) (*dto.TagDetails, *service_errors.ServiceError) {
	return nil, nil
}

func (t *TagService) GetByName(
	ctx context.Context, name string,
) (*dto.TagDetails, *service_errors.ServiceError) {
	return nil, nil
}

// -----------------------------------------------------------------------------
// other sources that has FK to `Tag`

func (t *TagService) GetPosts(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.PostsList], *service_errors.ServiceError) {
	return nil, nil
}
