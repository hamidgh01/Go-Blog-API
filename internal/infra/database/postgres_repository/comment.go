package postgres_repository

import (
	"context"
	"database/sql"
	"fmt"

	d "github.com/hamidgh01/Go-Blog-API/internal/domain"
	e "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	dbErrors "github.com/hamidgh01/Go-Blog-API/internal/infra/database/errors"
)

type commentRepository struct {
	DB *sql.DB
}

var _ repository.CommentRepository = (*commentRepository)(nil)

func NewCommentRepository(db *sql.DB) *commentRepository {
	prepareAllCommentStatements(db)
	return &commentRepository{DB: db}
}

func (r *commentRepository) Create(ctx context.Context, entity *e.Comment) (*e.Comment, error) {
	var row *sql.Row
	if entity.PostParentID.Valid && !entity.CommentParentID.Valid {
		row = createCommentStmt.QueryRowContext(ctx, entity.Content, entity.PostParentID.V, entity.UserID)
	} else if !entity.PostParentID.Valid && entity.CommentParentID.Valid {
		row = createReplyStmt.QueryRowContext(ctx, entity.Content, entity.CommentParentID.V, entity.UserID)
	} else {
		return nil, fmt.Errorf("internal server Error. comment parent conflict.")
	}

	err := row.Scan(
		&entity.ID,
		&entity.Content,
		&entity.Status,
		&entity.PostParentID,
		&entity.CommentParentID,
		&entity.UserID,
		&entity.CreatedAt,
		&entity.ModifiedAt,
	)

	if err != nil {
		return nil, dbErrors.GetDBError(err)
	}

	return entity, nil
}

func (r *commentRepository) Update(ctx context.Context, pk uint64, data *e.Comment) error {
	return update(ctx, updateCommentStmt, "Comment", pk, data.Content)
}

func (r *commentRepository) UpdateStatus(ctx context.Context, pk uint64, status e.CommentStatus) error {
	return update(ctx, updateCommentStatusStmt, "Comment", pk, status)
}

func (r *commentRepository) Delete(ctx context.Context, pk uint64) error {
	return delete(ctx, deleteCommentStmt, "Comment", pk)
}

func (r *commentRepository) GetByID(ctx context.Context, pk uint64) (*e.Comment, error) {
	comment := &e.Comment{}
	user := &e.User{}
	comment.User = user
	err := getCommentByIDStmt.QueryRowContext(ctx, pk).Scan(
		&comment.ID,
		&comment.Content,
		&comment.Status,
		&comment.PostParentID,
		&comment.CommentParentID,
		&comment.UserID,
		&comment.CreatedAt,
		&comment.ModifiedAt,
		&user.ID,
		&user.Username,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.NewRecordNotFoundError("Comment not found")
		}
		return nil, dbErrors.GetDBError(err)
	}

	return comment, nil
}

// -----------------------------------------------------------------------------
// other sources that has FK to `Comment`

func (r *commentRepository) GetReplies(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.Comment], error) {
	// implement later
	return nil, nil
}

// -----------------------------------------------------------------------------

func (r *commentRepository) GetOwnerID(ctx context.Context, pk uint64) (uint64, error) {
	return getOwnerID(ctx, getCommentOwnerIDStmt, "Comment", pk)
}
