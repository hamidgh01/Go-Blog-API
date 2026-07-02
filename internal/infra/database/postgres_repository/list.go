package postgres_repository

import (
	"context"
	"database/sql"

	d "github.com/hamidgh01/Go-Blog-API/internal/domain"
	e "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	dbErrors "github.com/hamidgh01/Go-Blog-API/internal/infra/database/errors"
)

type listRepository struct {
	DB *sql.DB
}

var _ repository.ListRepository = (*listRepository)(nil)

func NewListRepository(db *sql.DB) *listRepository {
	prepareAllListStatements(db)
	return &listRepository{DB: db}
}

func (r *listRepository) Create(ctx context.Context, entity *e.List) (*e.List, error) {
	err := createListStmt.QueryRowContext(
		ctx, entity.Title, entity.Description, entity.IsPrivate, entity.UserID,
	).Scan(
		&entity.ID,
		&entity.Title,
		&entity.Description,
		&entity.IsPrivate,
		&entity.UserID,
		&entity.CreatedAt,
		&entity.ModifiedAt,
	)
	if err != nil {
		return nil, dbErrors.GetDBError(err)
	}

	return entity, nil
}

func (r *listRepository) Update(ctx context.Context, pk uint64, data *e.List) error {
	return update(ctx, updateListStmt, "List", pk, data.Title, data.Description, data.IsPrivate)
}

func (r *listRepository) UpdatePrivacy(ctx context.Context, pk uint64, isPrivate bool) error {
	return update(ctx, updateListPrivacyStmt, "List", pk, isPrivate)
}

func (r *listRepository) Delete(ctx context.Context, pk uint64) error {
	return delete(ctx, deleteListStmt, "List", pk)
}

func (r *listRepository) GetByID(ctx context.Context, pk uint64) (*e.List, error) {
	list := &e.List{}
	user := &e.User{}
	list.User = user
	err := getListByIDStmt.QueryRowContext(ctx, pk).Scan(
		&list.ID,
		&list.Title,
		&list.Description,
		&list.IsPrivate,
		&list.UserID,
		&list.CreatedAt,
		&list.ModifiedAt,
		&user.ID,
		&user.Username,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.NewRecordNotFoundError("List not found")
		}
		return nil, dbErrors.GetDBError(err)
	}

	return list, nil
}

// -----------------------------------------------------------------------------
// other sources that has FK to `List`

func (r *listRepository) GetSavedPosts(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.Post], error) {
	// implement later
	return nil, nil
}

func (r *listRepository) GetUsersWhoSaved(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.User], error) {
	// implement later
	return nil, nil
}

// -----------------------------------------------------------------------------

func (r *listRepository) GetOwnerID(ctx context.Context, pk uint64) (uint64, error) {
	return getOwnerID(ctx, getListOwnerIDStmt, "List", pk)
}
