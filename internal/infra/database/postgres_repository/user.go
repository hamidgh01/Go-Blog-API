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

func (r *userRepository) GetByIDWithCountOfAllReferencedObjects(
	ctx context.Context, pk uint64,
) (*e.DBEntityWithCountOfReferencedObjects[e.User], error) {
	return nil, nil
}

// -----------------------------------------------------------------------------
// other sources that has FK to `User`

func (r *userRepository) GetPosts(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.Post], error) {
	// implement later
	return nil, nil
}

func (r *userRepository) GetOwnedLists(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.List], error) {
	// implement later
	return nil, nil
}

func (r *userRepository) GetSavedLists(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.List], error) {
	// implement later
	return nil, nil
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
	// implement later
	return nil, nil
}

func (r *userRepository) GetLikes(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.Post], error) {
	// implement later
	return nil, nil
}

func (r *userRepository) GetFollowers(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.User], error) {
	// implement later
	return nil, nil
}

func (r *userRepository) GetFollowings(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.User], error) {
	// implement later
	return nil, nil
}

func (r *userRepository) GetLinks(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*d.PagedList[e.Link], error) {
	// implement later
	return nil, nil
}
