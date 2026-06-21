package services

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	e "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
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
	userID := ctx.Value("currentUserID").(uint64)
	link := &e.Link{Title: data.Title, Url: data.Url, UserID: userID}

	username := ctx.Value("currentUserUsername").(string)
	link.User = &e.User{ID: userID, Username: username}

	return create(ctx, "link", link, l.repo.Create, dto.ToLinkDetails)
}

func (l *LinkService) Update(
	ctx context.Context, pk uint64, data *dto.UpdateLinkRequest,
) *service_errors.ServiceError {
	link := &e.Link{Title: data.Title, Url: data.Url}
	return update(ctx, pk, "link", link, l.repo.Update)
}

func (l *LinkService) Delete(ctx context.Context, pk uint64) *service_errors.ServiceError {
	return delete(ctx, pk, "link", l.repo.Delete)
}

func (l *LinkService) GetByID(
	ctx context.Context, pk uint64,
) (*dto.LinkDetails, *service_errors.ServiceError) {
	return getByID(ctx, pk, "link", l.repo.GetByID, dto.ToLinkDetails)
}

// -----------------------------------------------------------------------------

func (l *LinkService) GetOwnerID(ctx context.Context, pk uint64) (uint64, error) {
	return getOwnerID(ctx, pk, "link", l.repo.GetOwnerID)
}
