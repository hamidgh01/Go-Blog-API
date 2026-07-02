package postgres_repository

import (
	"context"
	"database/sql"

	d "github.com/hamidgh01/Go-Blog-API/internal/domain"
	e "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	dbErrors "github.com/hamidgh01/Go-Blog-API/internal/infra/database/errors"
)

type userRepository struct {
	DB *sql.DB
}

var _ repository.UserRepository = (*userRepository)(nil)

func NewUserRepository(db *sql.DB) *userRepository {
	prepareAllUserStatements(db)
	return &userRepository{DB: db}
}

func fillUserEntityForDetailsResponse(row *sql.Row, entity *e.User) (*e.User, error) {
	err := row.Scan(
		&entity.ID,
		&entity.Username,
		&entity.Email,
		&entity.Bio,
		&entity.Enabled,
		&entity.CreatedAt,
		&entity.ModifiedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.NewRecordNotFoundError("User not found")
		}
		return nil, dbErrors.GetDBError(err)
	}

	return entity, nil
}

func (r *userRepository) Create(ctx context.Context, entity *e.User) (*e.User, error) {
	row := createUserStmt.QueryRowContext(ctx, entity.Username, entity.Email, entity.Password)
	return fillUserEntityForDetailsResponse(row, entity)
}

func (r *userRepository) UpdateUsername(ctx context.Context, pk uint64, username string) error {
	return update(ctx, updateUsernameStmt, "User", pk, username)
}

func (r *userRepository) UpdateEmail(ctx context.Context, pk uint64, email string) error {
	return update(ctx, updateEmailStmt, "User", pk, email)
}

func (r *userRepository) UpdateBio(ctx context.Context, pk uint64, bio string) error {
	return update(ctx, updateBioStmt, "User", pk, bio)
}

func (r *userRepository) UpdatePassword(ctx context.Context, pk uint64, password string) error {
	return update(ctx, updatePasswordStmt, "User", pk, password)
}

func (r *userRepository) UpdateEnabled(ctx context.Context, pk uint64, enabled bool) error {
	return update(ctx, updateEnabledStmt, "User", pk, enabled)
}

func (r *userRepository) Delete(ctx context.Context, pk uint64) error {
	return delete(ctx, deleteUserStmt, "User", pk)
}

func (r *userRepository) CheckUsernameExists(ctx context.Context, username string) (bool, error) {
	return checkUniqueFieldExists(ctx, checkUsernameExistsStmt, username)
}

func (r *userRepository) CheckEmailExists(ctx context.Context, email string) (bool, error) {
	return checkUniqueFieldExists(ctx, checkEmailExistsStmt, email)
}

func (r *userRepository) CheckIsEnabled(ctx context.Context, pk uint64) (bool, error) {
	isEnabled, err := getFieldValue(ctx, checkIsEnabledStmt, "User", pk)
	return isEnabled.(bool), err
}

func (r *userRepository) CheckIsSuperuser(ctx context.Context, pk uint64) (bool, error) {
	isSuperuser, err := getFieldValue(ctx, checkIsSuperuserStmt, "User", pk)
	return isSuperuser.(bool), err
}

func (r *userRepository) GetHashedPassword(ctx context.Context, pk uint64) (string, error) {
	hashedPassword, err := getFieldValue(ctx, getHashedPasswordStmt, "User", pk)
	return hashedPassword.(string), err
}

func (r *userRepository) GetUserForLoginVerification(ctx context.Context, identifier string) (*e.User, error) {
	var user = &e.User{}
	err := getUserForLoginVerificationStmt.QueryRowContext(ctx, identifier).Scan(
		&user.ID, &user.Username, &user.Enabled, &user.Password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.NewRecordNotFoundError("User not found")
		}
		return nil, dbErrors.GetDBError(err)
	}

	return user, nil
}

func (r *userRepository) GetByID(ctx context.Context, pk uint64) (*e.User, error) {
	row := getUserByIDStmt.QueryRowContext(ctx, pk)
	return fillUserEntityForDetailsResponse(row, &e.User{})
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*e.User, error) {
	row := getUserByUsernameStmt.QueryRowContext(ctx, username)
	return fillUserEntityForDetailsResponse(row, &e.User{})
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*e.User, error) {
	row := getUserByEmailStmt.QueryRowContext(ctx, email)
	return fillUserEntityForDetailsResponse(row, &e.User{})
}

// -----------------------------------------------------------------------------
// other sources that has FK to `User`

func (r *userRepository) GetPosts(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.Post], error) {
	rows, totalRows, pageNum, pageSize, totalPages, err := getListOfOuterResourceByFK(
		ctx, r.DB, fk, page,
		countUserPostsQuery,
		getPostsByUserIdFkQuery,
		"GetUserPosts",
		"there is not any post for this user",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*e.Post
	for rows.Next() {
		post := &e.Post{User: &e.User{}}
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.IsPrivate,
			&post.UserID,
			&post.CreatedAt,
			&post.ModifiedAt,
			&post.FirstPublishedAt,
			&post.User.ID,
			&post.User.Username,
		)
		if err != nil {
			return nil, dbErrors.GetDBError(err)
		}

		posts = append(posts, post)
	}

	pagedPostList := d.Paginate(posts, totalRows, pageNum, pageSize, totalPages)

	return pagedPostList, nil
}

func (r *userRepository) GetOwnedLists(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.List], error) {
	rows, totalRows, pageNum, pageSize, totalPages, err := getListOfOuterResourceByFK(
		ctx, r.DB, fk, page,
		countOwnedListsQuery,
		getOwnedListsByUserIdFkQuery,
		"GetUserOwnedLists",
		"there is not any owned-list for this user",
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

func (r *userRepository) GetSavedLists(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.List], error) {
	rows, totalRows, pageNum, pageSize, totalPages, err := getListOfOuterResourceByFK(
		ctx, r.DB, fk, page,
		countSavedListsQuery,
		getSavedListsByUserIdFkQuery,
		"GetUserSavedLists",
		"there is not any saved-list for this user",
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

func (r *userRepository) GetAllLists(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.List], error) {
	// implement later
	return nil, nil
}

func (r *userRepository) GetComments(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.Comment], error) {
	rows, totalRows, pageNum, pageSize, totalPages, err := getListOfOuterResourceByFK(
		ctx, r.DB, fk, page,
		countUserCommentsQuery,
		getCommentsByUserIdFkQuery,
		"GetUserComments",
		"there is not any comment for this user",
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
			&comment.CommentParentID,
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

	pagedCommentList := d.Paginate(comments, totalRows, pageNum, pageSize, totalPages)

	return pagedCommentList, nil
}

func (r *userRepository) GetLikes(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.Post], error) {
	rows, totalRows, pageNum, pageSize, totalPages, err := getListOfOuterResourceByFK(
		ctx, r.DB, fk, page,
		countUserLikesQuery,
		getLikesByUserIdFkQuery,
		"GetUserLikes",
		"there is not any liked post for this user",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var likedPosts []*e.Post
	for rows.Next() {
		likedPost := &e.Post{User: &e.User{}}
		err := rows.Scan(
			&likedPost.ID,
			&likedPost.Title,
			&likedPost.IsPrivate,
			&likedPost.UserID,
			&likedPost.CreatedAt,
			&likedPost.ModifiedAt,
			&likedPost.FirstPublishedAt,
			&likedPost.User.ID,
			&likedPost.User.Username,
		)
		if err != nil {
			return nil, dbErrors.GetDBError(err)
		}

		likedPosts = append(likedPosts, likedPost)
	}

	pagedLikedPostsList := d.Paginate(likedPosts, totalRows, pageNum, pageSize, totalPages)

	return pagedLikedPostsList, nil
}

func (r *userRepository) GetFollowers(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.User], error) {
	rows, totalRows, pageNum, pageSize, totalPages, err := getListOfOuterResourceByFK(
		ctx, r.DB, fk, page,
		countFollowersQuery,
		getFollowersByUserIdFkQuery,
		"GetUserFollowers",
		"there is not any follower for this user",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followersList []*e.User
	for rows.Next() {
		follower := &e.User{}
		err := rows.Scan(&follower.ID, &follower.Username)
		if err != nil {
			return nil, dbErrors.GetDBError(err)
		}

		followersList = append(followersList, follower)
	}

	pagedFollowersList := d.Paginate(followersList, totalRows, pageNum, pageSize, totalPages)

	return pagedFollowersList, nil
}

func (r *userRepository) GetFollowings(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.User], error) {
	rows, totalRows, pageNum, pageSize, totalPages, err := getListOfOuterResourceByFK(
		ctx, r.DB, fk, page,
		countFollowingsQuery,
		getFollowingsByUserIdFkQuery,
		"GetUserFollowings",
		"there is not any following for this user",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followingList []*e.User
	for rows.Next() {
		following := &e.User{}
		err := rows.Scan(&following.ID, &following.Username)
		if err != nil {
			return nil, dbErrors.GetDBError(err)
		}

		followingList = append(followingList, following)
	}

	pagedFollowingsList := d.Paginate(followingList, totalRows, pageNum, pageSize, totalPages)

	return pagedFollowingsList, nil
}

func (r *userRepository) GetLinks(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.Link], error) {
	rows, totalRows, pageNum, pageSize, totalPages, err := getListOfOuterResourceByFK(
		ctx, r.DB, fk, page,
		countUserLinksQuery,
		getLinksByUserIdFkQuery,
		"GetUserLinks",
		"there is not any link for this user",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []*e.Link
	for rows.Next() {
		link := &e.Link{User: &e.User{}}
		err := rows.Scan(
			&link.ID, &link.Title, &link.Url, &link.UserID, &link.User.ID, &link.User.Username,
		)
		if err != nil {
			return nil, dbErrors.GetDBError(err)
		}

		links = append(links, link)
	}

	pagedLinksList := d.Paginate(links, totalRows, pageNum, pageSize, totalPages)

	return pagedLinksList, nil
}
