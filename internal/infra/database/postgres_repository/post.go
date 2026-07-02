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
	post := &e.Post{}
	user := &e.User{}
	post.User = user
	err := getPostByIDStmt.QueryRowContext(ctx, pk).Scan(
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

// -----------------------------------------------------------------------------
// other sources that has FK to `Post`

func (r *postRepository) GetComments(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.Comment], error) {
	rows, totalRows, pageNum, pageSize, totalPages, err := getListOfOuterResourceByFK(
		ctx, r.DB, fk, page,
		countPostCommentsQuery,
		getPostCommentsQuery,
		"GetPostComments",
		"there is not any comment for this post",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*e.Comment
	for rows.Next() {
		comment := &e.Comment{User: &e.User{}}
		err := rows.Scan(
			&comment.ID,
			&comment.Content,
			&comment.Status,
			&comment.PostParentID,
			&comment.UserID,
			&comment.CreatedAt,
			&comment.ModifiedAt,
			&comment.User.ID,
			&comment.User.Username,
		)
		if err != nil {
			return nil, dbErrors.GetDBError(err)
		}

		comments = append(comments, comment)
	}

	pagedComments := d.Paginate(comments, totalRows, pageNum, pageSize, totalPages)

	return pagedComments, nil
}

func (r *postRepository) GetLikes(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.User], error) {
	rows, totalRows, pageNum, pageSize, totalPages, err := getListOfOuterResourceByFK(
		ctx, r.DB, fk, page,
		countPostLikesQuery,
		getPostLikesQuery,
		"GetPostLikes",
		"this post isn't liked by any user.",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usersList []*e.User
	for rows.Next() {
		user := &e.User{}
		err := rows.Scan(&user.ID, &user.Username)
		if err != nil {
			return nil, dbErrors.GetDBError(err)
		}

		usersList = append(usersList, user)
	}

	pagedUsers := d.Paginate(usersList, totalRows, pageNum, pageSize, totalPages)

	return pagedUsers, nil
}

func (r *postRepository) GetListsThatSavedThisPost(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.List], error) {
	rows, totalRows, pageNum, pageSize, totalPages, err := getListOfOuterResourceByFK(
		ctx, r.DB, fk, page,
		countListsThatSavedThisPostQuery,
		getListsThatSavedThisPostQuery,
		"GetListsThatSavedThisPost",
		"this post is not saved in any list.",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lists []*e.List
	for rows.Next() {
		list := &e.List{User: &e.User{}}
		err := rows.Scan(
			&list.ID,
			&list.Title,
			&list.IsPrivate,
			&list.UserID,
			&list.CreatedAt,
			&list.ModifiedAt,
			&list.User.ID,
			&list.User.Username,
		)
		if err != nil {
			return nil, dbErrors.GetDBError(err)
		}

		lists = append(lists, list)
	}

	pagedLists := d.Paginate(lists, totalRows, pageNum, pageSize, totalPages)

	return pagedLists, nil
}

func (r *postRepository) GetTags(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.Tag], error) {
	rows, totalRows, pageNum, pageSize, totalPages, err := getListOfOuterResourceByFK(
		ctx, r.DB, fk, page,
		countPostTagsQuery,
		getPostTagsQuery,
		"GetPostTags",
		"this post is not tagged by any tag.",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []*e.Tag
	for rows.Next() {
		tag := &e.Tag{}
		err := rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			return nil, dbErrors.GetDBError(err)
		}

		tags = append(tags, tag)
	}

	pagedTagsList := d.Paginate(tags, totalRows, pageNum, pageSize, totalPages)

	return pagedTagsList, nil
}

// -----------------------------------------------------------------------------

func (r *postRepository) GetOwnerID(ctx context.Context, pk uint64) (uint64, error) {
	return getOwnerID(ctx, getPostOwnerIDStmt, "Post", pk)
}
