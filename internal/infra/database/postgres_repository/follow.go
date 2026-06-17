package postgres_repository

import (
	"context"
	"database/sql"

	e "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
)

const (
	createFollowRelationshipQuery string = `
		INSERT INTO follows_m2m (followed_by, followed) VALUES ($1, $2)
		ON CONFLICT (followed_by, followed) DO NOTHING
	`

	deleteFollowRelationshipQuery = `
		DELETE FROM follows_m2m WHERE followed_by = $1 AND followed = $2
	`
)

var (
	createFollowRelationshipStmt *sql.Stmt
	deleteFollowRelationshipStmt *sql.Stmt
)

func prepareAllFollowStatements(db *sql.DB) {
	createFollowRelationshipStmt = prepareStatement(db, "createFollowRelationship", createFollowRelationshipQuery)
	deleteFollowRelationshipStmt = prepareStatement(db, "deleteFollowRelationship", deleteFollowRelationshipQuery)
}

func closeAllFollowStatements() {
	createLikeStmt.Close()
	deleteLikeStmt.Close()
}

// -----------------------------------------------------------------------------

type followRepository struct {
	DB *sql.DB
}

var _ repository.FollowRepository = (*followRepository)(nil)

func NewFollowRepository(db *sql.DB) *followRepository {
	prepareAllFollowStatements(db)
	return &followRepository{DB: db}
}

func (r *followRepository) Create(ctx context.Context, entity *e.FollowsM2M) error {
	return createOrDeleteM2MRelationship(
		ctx,
		createFollowRelationshipStmt,
		entity.FollowedBy, entity.Followed,
		"already followed (nothing changed)",
	)
}

func (r *followRepository) Delete(ctx context.Context, entity *e.FollowsM2M) error {
	return createOrDeleteM2MRelationship(
		ctx,
		deleteFollowRelationshipStmt,
		entity.FollowedBy, entity.Followed,
		"there's no follow relationship to delete (nothing changed)",
	)
}
