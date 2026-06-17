package postgres_repository

import (
	"context"
	"database/sql"

	e "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
)

const (
	createLikeQuery string = `
		INSERT INTO post_likes_m2m (post_id, user_id) VALUES ($1, $2)
		ON CONFLICT (post_id, user_id) DO NOTHING
	`

	deleteLikeQuery = `
		DELETE FROM post_likes_m2m WHERE post_id = $1 AND user_id = $2
	`
)

var (
	createLikeStmt *sql.Stmt
	deleteLikeStmt *sql.Stmt
)

func prepareAllLikeStatements(db *sql.DB) {
	createLikeStmt = prepareStatement(db, "createLike", createLikeQuery)
	deleteLikeStmt = prepareStatement(db, "deleteLike", deleteLikeQuery)
}

func closeAllLikeStatements() {
	createLikeStmt.Close()
	deleteLikeStmt.Close()
}

// -----------------------------------------------------------------------------

type likeRepository struct {
	DB *sql.DB
}

var _ repository.LikeRepository = (*likeRepository)(nil)

func NewLikeRepository(db *sql.DB) *likeRepository {
	prepareAllLikeStatements(db)
	return &likeRepository{DB: db}
}

func (r *likeRepository) Create(ctx context.Context, entity *e.PostLikesM2M) error {
	return createOrDeleteM2MRelationship(
		ctx,
		createLikeStmt,
		entity.PostID, entity.UserID,
		"already liked (nothing changed)",
	)
}

func (r *likeRepository) Delete(ctx context.Context, entity *e.PostLikesM2M) error {
	return createOrDeleteM2MRelationship(
		ctx,
		deleteLikeStmt,
		entity.PostID, entity.UserID,
		"no like to delete (nothing changed)",
	)
}
