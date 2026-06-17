package postgres_repository

import (
	"context"
	"database/sql"

	e "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
)

const (
	savePostQuery string = `
		INSERT INTO saved_posts_m2m (list_id, post_id) VALUES ($1, $2)
		ON CONFLICT (list_id, post_id) DO NOTHING
	`

	unsavePostQuery = `
		DELETE FROM saved_posts_m2m WHERE list_id = $1 AND post_id = $2
	`
)

var (
	savePostStmt   *sql.Stmt
	unsavePostStmt *sql.Stmt
)

func prepareAllSavePostStatements(db *sql.DB) {
	savePostStmt = prepareStatement(db, "savePost", savePostQuery)
	unsavePostStmt = prepareStatement(db, "unsavePost", unsavePostQuery)
}

func closeAllSavePostStatements() {
	savePostStmt.Close()
	unsavePostStmt.Close()
}

// -----------------------------------------------------------------------------

type savePostRepository struct {
	DB *sql.DB
}

var _ repository.SavePostRepository = (*savePostRepository)(nil)

func NewSavePostRepository(db *sql.DB) *savePostRepository {
	prepareAllSavePostStatements(db)
	return &savePostRepository{DB: db}
}

func (r *savePostRepository) Create(ctx context.Context, entity *e.SavedPostsM2M) error {
	return createOrDeleteM2MRelationship(
		ctx,
		savePostStmt,
		entity.ListID, entity.PostID,
		"already saved this post (nothing changed)",
	)
}

func (r *savePostRepository) Delete(ctx context.Context, entity *e.SavedPostsM2M) error {
	return createOrDeleteM2MRelationship(
		ctx,
		unsavePostStmt,
		entity.ListID, entity.PostID,
		"already not saved this post (nothing changed)",
	)
}
