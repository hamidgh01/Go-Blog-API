package postgres_repository

import (
	"context"
	"database/sql"

	e "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
)

const (
	saveListQuery string = `
		INSERT INTO users_saved_lists_m2m (user_id, list_id) VALUES ($1, $2)
		ON CONFLICT (user_id, list_id) DO NOTHING
	`

	unsaveListQuery = `
		DELETE FROM users_saved_lists_m2m WHERE user_id = $1 AND list_id = $2
	`
)

var (
	saveListStmt   *sql.Stmt
	unsaveListStmt *sql.Stmt
)

func prepareAllSaveListStatements(db *sql.DB) {
	saveListStmt = prepareStatement(db, "saveList", saveListQuery)
	unsaveListStmt = prepareStatement(db, "unsaveList", unsaveListQuery)
}

func closeAllSaveListStatements() {
	saveListStmt.Close()
	unsaveListStmt.Close()
}

// -----------------------------------------------------------------------------

type saveListRepository struct {
	DB *sql.DB
}

var _ repository.SaveListRepository = (*saveListRepository)(nil)

func NewSaveListRepository(db *sql.DB) *saveListRepository {
	prepareAllSaveListStatements(db)
	return &saveListRepository{DB: db}
}

func (r *saveListRepository) Create(ctx context.Context, entity *e.UsersSavedListsM2M) error {
	return createOrDeleteM2MRelationship(
		ctx,
		saveListStmt,
		entity.UserID, entity.ListID,
		"already saved this list (nothing changed)",
	)
}

func (r *saveListRepository) Delete(ctx context.Context, entity *e.UsersSavedListsM2M) error {
	return createOrDeleteM2MRelationship(
		ctx,
		unsaveListStmt,
		entity.UserID, entity.ListID,
		"already not saved this list (nothing changed)",
	)
}
