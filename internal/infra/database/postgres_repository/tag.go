package postgres_repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	d "github.com/hamidgh01/Go-Blog-API/internal/domain"
	e "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	dbErrors "github.com/hamidgh01/Go-Blog-API/internal/infra/database/errors"
)

type tagRepository struct {
	DB *sql.DB
}

var _ repository.TagRepository = (*tagRepository)(nil)

func NewTagRepository(db *sql.DB) *tagRepository {
	prepareAllTagStatements(db)
	return &tagRepository{DB: db}
}

func getOrCreateBulkTags(
	ctx context.Context, db *sql.DB, getOrCreateQuery string, tags []*e.Tag,
) ([]*e.Tag, error) {
	var (
		placeholders []string
		values       []any
	)

	for i, tag := range tags {
		placeholders = append(placeholders, fmt.Sprintf("($%d)", i+1))
		values = append(values, tag.Name)
	}

	query := fmt.Sprintf(getOrCreateQuery, strings.Join(placeholders, ", "))

	rows, err := db.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, dbErrors.GetDBError(err)
	}
	defer rows.Close()

	fetchedOrCreatedTags := make([]*e.Tag, 0, len(tags))
	for rows.Next() {
		tag := &e.Tag{}
		if err := rows.Scan(&tag.ID, &tag.Name); err != nil {
			return nil, dbErrors.GetDBError(err)
		}

		fetchedOrCreatedTags = append(fetchedOrCreatedTags, tag)
	}

	return fetchedOrCreatedTags, nil
}

func (r *tagRepository) Create(ctx context.Context, entities []*e.Tag) ([]*e.Tag, error) {
	return getOrCreateBulkTags(ctx, r.DB, bulkCreateTagQuery, entities)
}

func (r *tagRepository) GetByID(ctx context.Context, pk uint64) (*e.Tag, error) {
	tag := &e.Tag{}
	err := getTagByIDStmt.QueryRowContext(ctx, pk).Scan(&tag.ID, &tag.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.NewRecordNotFoundError("Tag not found")
		}
		return nil, dbErrors.GetDBError(err)
	}

	return tag, nil
}

func (r *tagRepository) GetByName(ctx context.Context, name string) (*e.Tag, error) {
	tag := &e.Tag{}
	err := getTagByNameStmt.QueryRowContext(ctx, name).Scan(&tag.ID, &tag.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.NewRecordNotFoundError("Tag not found")
		}
		return nil, dbErrors.GetDBError(err)
	}

	return tag, nil
}

func (r *tagRepository) GetListOfTagsByNames(ctx context.Context, tags []*e.Tag) ([]*e.Tag, error) {
	return getOrCreateBulkTags(ctx, r.DB, getListOfTagsByNamesQuery, tags)
}

// -----------------------------------------------------------------------------
// other sources that has FK to `Tag`

func (r *tagRepository) GetPosts(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.Post], error) {
	// implement later
	return nil, nil
}
