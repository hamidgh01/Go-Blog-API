package services

import (
	"context"
	"database/sql"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	d "github.com/hamidgh01/Go-Blog-API/internal/domain"
	e "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
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

func (l *ListService) Create(
	ctx context.Context, data *dto.CreateListRequest,
) (*dto.ListDetails, *service_errors.ServiceError) {
	userID := ctx.Value("currentUserID").(uint64)
	list := &e.List{Title: data.Title, UserID: userID}

	if data.Description == "" {
		list.Description = sql.NullString{Valid: false}
	} else {
		list.Description = sql.NullString{String: data.Description, Valid: true}
	}

	if data.IsPrivate == nil {
		list.IsPrivate = false
	} else {
		list.IsPrivate = *data.IsPrivate
	}

	username := ctx.Value("currentUserUsername").(string)
	list.User = &e.User{ID: userID, Username: username}

	return create(ctx, "list", list, l.repo.Create, dto.ToListDetails)
}

func (l *ListService) Update(
	ctx context.Context, pk uint64, data *dto.UpdateListRequest,
) *service_errors.ServiceError {
	list := &e.List{Title: data.Title}

	if data.Description == "" {
		list.Description = sql.NullString{Valid: false}
	} else {
		list.Description = sql.NullString{String: data.Description, Valid: true}
	}

	if data.IsPrivate == nil {
		list.IsPrivate = false
	} else {
		list.IsPrivate = *data.IsPrivate
	}

	return update(ctx, pk, "list", list, l.repo.Update)
}

func (l *ListService) UpdatePrivacy(
	ctx context.Context, pk uint64, data *dto.UpdateListPrivacyRequest,
) *service_errors.ServiceError {
	err := l.repo.UpdatePrivacy(ctx, pk, *data.IsPrivate)
	if err != nil {
		return service_errors.MapDBErrToServiceErr(err, "update list privacy")
	}

	return nil
}

func (l *ListService) Delete(ctx context.Context, pk uint64) *service_errors.ServiceError {
	return delete(ctx, pk, "list", l.repo.Delete)
}

func (l *ListService) GetByID(
	ctx context.Context, pk uint64,
) (*dto.ListDetails, *service_errors.ServiceError) {
	return getByID(ctx, pk, "list", l.repo.GetByID, dto.ToListDetails)
}

// -----------------------------------------------------------------------------
// other sources that has FK to `List`

func (l *ListService) GetSavedPosts(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.PostsList], *service_errors.ServiceError) {
	return getListOfOuterResourceByFK(
		ctx, fk, page, "get saved posts of this list", l.repo.GetSavedPosts, dto.ToPostsList,
	)
}

func (l *ListService) GetUsersWhoSaved(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.UsersList], *service_errors.ServiceError) {
	return getListOfOuterResourceByFK(
		ctx, fk, page, "get users who saved this list", l.repo.GetUsersWhoSaved, dto.ToUsersList,
	)
}

// -----------------------------------------------------------------------------

func (l *ListService) GetOwnerID(ctx context.Context, pk uint64) (uint64, error) {
	return getOwnerID(ctx, pk, "list", l.repo.GetOwnerID)
}
