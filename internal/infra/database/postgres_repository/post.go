package postgres_repository

import (
	"context"
	"database/sql"

	d "github.com/hamidgh01/Go-Blog-API/internal/domain"
	e "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	dbErrors "github.com/hamidgh01/Go-Blog-API/internal/infra/database/errors"
)

type postRepository struct {
	DB *sql.DB
}

var _ repository.PostRepository = (*postRepository)(nil)

func NewPostRepository(db *sql.DB) *postRepository {
	prepareAllPostStatements(db)
	return &postRepository{DB: db}
}

func (r *postRepository) Create(ctx context.Context, entity *e.Post) (*e.Post, error) {
	var row *sql.Row
	switch entity.Status {
	case e.DR:
		row = createDraftPostStmt.QueryRowContext(
			ctx, entity.Title, entity.Content, entity.IsPrivate, entity.UserID,
		)
	case e.PB:
		row = createPublishedPostStmt.QueryRowContext(
			ctx, entity.Title, entity.Content, entity.IsPrivate, entity.UserID,
		)
	default:
		return nil, dbErrors.NewBadInputError(
			"invalid status input. for creating a post, status can be 'draft' or 'published'",
		)
	}

	err := row.Scan(
		&entity.ID,
		&entity.Title,
		&entity.Content,
		&entity.Status,
		&entity.IsPrivate,
		&entity.UserID,
		&entity.CreatedAt,
		&entity.ModifiedAt,
		&entity.FirstPublishedAt,
	)
	if err != nil {
		return nil, dbErrors.GetDBError(err)
	}

	return entity, nil
}

func (r *postRepository) Update(ctx context.Context, pk uint64, data *e.Post) error {
	return update(ctx, updatePostStmt, "Post", pk, data.Title, data.Content)
}

func (r *postRepository) UpdateStatus(ctx context.Context, pk uint64, status e.PostStatus) error {
	if status == e.DR {
		return dbErrors.NewBadInputError(
			"invalid status input. can't 'draft' a post that already is published",
		)
	}

	return update(ctx, updatePostStatusStmt, "Post", pk, status)
}

func (r *postRepository) PublishDraftPost(ctx context.Context, pk uint64) error {
	return update(ctx, publishDraftPostStmt, "Post", pk)
}

func (r *postRepository) UpdatePrivacy(ctx context.Context, pk uint64, isPrivate bool) error {
	return update(ctx, updatePostPrivacyStmt, "Post", pk, isPrivate)
}

func (r *postRepository) Delete(ctx context.Context, pk uint64) error {
	return delete(ctx, deletePostStmt, "Post", pk)
}

func (r *postRepository) GetByID(ctx context.Context, pk uint64) (*e.Post, error) {
	row := getPostByIDStmt.QueryRowContext(ctx, pk)

	post := &e.Post{}
	user := &e.User{}
	post.User = user
	err := row.Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Status,
		&post.IsPrivate,
		&post.UserID,
		&post.CreatedAt,
		&post.ModifiedAt,
		&post.FirstPublishedAt,
		&user.ID,
		&user.Username,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.NewRecordNotFoundError("Post not found")
		}
		return nil, dbErrors.GetDBError(err)
	}

	return post, nil
}

func (r *postRepository) GetByIDWithCountOfAllReferencedObjects(
	ctx context.Context, pk uint64,
) (*e.DBEntityWithCountOfReferencedObjects[e.Post], error) {
	// implement later
	return nil, nil
}

// -----------------------------------------------------------------------------
// other sources that has FK to `Post`

func (r *postRepository) GetComments(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.Comment], error) {
	// implement later
	return nil, nil
}

func (r *postRepository) GetLikes(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.User], error) {
	// implement later
	return nil, nil
}

func (r *postRepository) GetListsThatSavedThisPost(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.List], error) {
	// implement later
	return nil, nil
}

func (r *postRepository) GetTags(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.Tag], error) {
	// implement later
	return nil, nil
}

// -----------------------------------------------------------------------------

func (r *postRepository) GetOwnerID(ctx context.Context, pk uint64) (uint64, error) {
	return getOwnerID(ctx, getPostOwnerIDStmt, "Post", pk)
}
