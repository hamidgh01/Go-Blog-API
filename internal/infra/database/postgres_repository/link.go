package postgres_repository

import (
	"context"
	"database/sql"

	e "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	dbErrors "github.com/hamidgh01/Go-Blog-API/internal/infra/database/errors"
)

type linkRepository struct {
	DB *sql.DB
}

var _ repository.LinkRepository = (*linkRepository)(nil)

func NewLinkRepository(db *sql.DB) *linkRepository {
	prepareAllLinkStatements(db)
	return &linkRepository{DB: db}
}

func (r *linkRepository) Create(ctx context.Context, entity *e.Link) (*e.Link, error) {
	err := createLinkStmt.QueryRowContext(ctx, entity.Title, entity.Url, entity.UserID).Scan(
		&entity.ID,
		&entity.Title,
		&entity.Url,
		&entity.UserID,
	)
	if err != nil {
		return nil, dbErrors.GetDBError(err)
	}

	return entity, nil
}

func (r *linkRepository) Update(ctx context.Context, pk uint64, data *e.Link) error {
	return update(ctx, updateLinkStmt, "Link", pk, data.Title, data.Url)
}

func (r *linkRepository) Delete(ctx context.Context, pk uint64) error {
	return delete(ctx, deleteLinkStmt, "Link", pk)
}

func (r *linkRepository) GetByID(ctx context.Context, pk uint64) (*e.Link, error) {
	link := &e.Link{}
	user := &e.User{}
	link.User = user
	err := getLinkByIDStmt.QueryRowContext(ctx, pk).Scan(
		&link.ID,
		&link.Title,
		&link.Url,
		&link.UserID,
		&user.ID,
		&user.Username,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.NewRecordNotFoundError("Link not found")
		}
		return nil, dbErrors.GetDBError(err)
	}

	return link, nil
}

// -----------------------------------------------------------------------------

func (r *linkRepository) GetOwnerID(ctx context.Context, pk uint64) (uint64, error) {
	return getOwnerID(ctx, getLinkOwnerIDStmt, "Link", pk)
}
