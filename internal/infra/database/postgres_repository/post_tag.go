package postgres_repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	e "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	dbErrors "github.com/hamidgh01/Go-Blog-API/internal/infra/database/errors"
)

const (
	associateQuery string = `
		INSERT INTO posts_tags_m2m (tag_id, post_id) VALUES %s
		ON CONFLICT (tag_id, post_id) DO NOTHING
	`
	// constructed query would be something like this:
	// 	INSERT INTO posts_tags_m2m (tag_id, post_id) VALUES ($1, $2), ($3, $4) ...
	// 	ON CONFLICT (tag_id, post_id) DO NOTHING

	dissociateQuery = `
		DELETE FROM posts_tags_m2m
		WHERE %s
	`
	// constructed query would be something like this:
	// 	DELETE FROM posts_tags_m2m
	// 	WHERE (tag_id = $1 AND post_id = $2) OR (tag_id = $3 AND post_id = $4) ...
)

type postTagsRepository struct {
	DB *sql.DB
}

var _ repository.PostTagsRepository = (*postTagsRepository)(nil)

func NewPostTagsRepository(db *sql.DB) *postTagsRepository {
	return &postTagsRepository{DB: db}
}

func associateOrDissociatePostWithTags(
	ctx context.Context,
	db *sql.DB,
	entities []*e.PostTagsM2M,
	associateOrDissociateQuery string,
	placeholder string,
	joiner string,
	noRowsAffectedMessage string,
) error {
	var (
		placeholders []string
		values       []any
	)

	for i, entity := range entities {
		n := i * 2
		placeholders = append(placeholders, fmt.Sprintf(placeholder, n+1, n+2))
		values = append(values, entity.TagID, entity.PostID)
	}

	query := fmt.Sprintf(associateOrDissociateQuery, strings.Join(placeholders, joiner))

	result, err := db.ExecContext(ctx, query, values...)
	if err != nil {
		return dbErrors.GetDBError(err)
	}

	n, e := result.RowsAffected()
	if n == 0 || e == sql.ErrNoRows {
		return dbErrors.NewNoRowsAffectedOnM2MEntity(noRowsAffectedMessage)
	} else if e != nil {
		return dbErrors.GetDBError(e)
	}

	return nil
}

func (r *postTagsRepository) Create(ctx context.Context, entities []*e.PostTagsM2M) error {
	return associateOrDissociatePostWithTags(
		ctx,
		r.DB,
		entities,
		associateQuery,
		"($%d, $%d)",
		", ",
		"the post already associated with all of these tags (nothing changed)",
	)
}

func (r *postTagsRepository) Delete(ctx context.Context, entities []*e.PostTagsM2M) error {
	return associateOrDissociatePostWithTags(
		ctx,
		r.DB,
		entities,
		dissociateQuery,
		"(tag_id = $%d AND post_id = $%d)",
		" OR ",
		"the post didn't tagged with any of these tags (nothing changed)",
	)
}
