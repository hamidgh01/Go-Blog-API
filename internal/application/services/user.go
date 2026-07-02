package services

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	d "github.com/hamidgh01/Go-Blog-API/internal/domain"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
	"github.com/hamidgh01/Go-Blog-API/internal/http/generics"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/redis"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/security/hashing"
	"github.com/hamidgh01/Go-Blog-API/pkg/logging"
)

type UserService struct {
	repo          repository.UserRepository
	pswHasher     *hashing.PasswordHasher
	userInfoCache *redis.UserInfoCache
	logger        *logging.Logger
}

func NewUserService(
	r repository.UserRepository, p *hashing.PasswordHasher, uic *redis.UserInfoCache,
) *UserService {
	return &UserService{repo: r, pswHasher: p, userInfoCache: uic, logger: logging.GetLogger()}
}

func (u *UserService) Create(
	ctx context.Context, data *dto.CreateUserRequest,
) (*dto.UserDetails, *service_errors.ServiceError) {
	hashedPassword, err := u.pswHasher.Hash(data.Password)
	if err != nil {
		u.logger.Errorf("failed to hash password. reason: %s", err.Error())
		return nil, service_errors.InternalServerError
	}

	user := &entity.User{Username: data.Username, Email: data.Email, Password: hashedPassword}
	createdUser, err := u.repo.Create(ctx, user)
	if err != nil {
		return nil, service_errors.MapDBErrToServiceErr(err, "create user")
	}

	// save user info in cache
	// NOTE: when a user first created (here by this service) --> superuser=false, enabled=true
	redisErr := u.userInfoCache.SetAllInfo(ctx, createdUser.ID, createdUser.Username, false, true)
	if redisErr != nil {
		u.logger.Error(redisErr.Error())
	}

	return dto.ToUserDetails(createdUser), nil
}

func (u *UserService) UpdateUsername(
	ctx context.Context, pk uint64, data *dto.UpdateUsernameRequest,
) *service_errors.ServiceError {
	err := u.repo.UpdateUsername(ctx, pk, data.Username)
	if err != nil {
		return service_errors.MapDBErrToServiceErr(err, "update user.username")
	}

	redisErr := u.userInfoCache.UpdateUsername(ctx, pk, data.Username)
	if redisErr != nil {
		u.logger.Error(redisErr.Error())
	}

	return nil
}

func (u *UserService) UpdateEmail(
	ctx context.Context, pk uint64, data *dto.UpdateEmailRequest,
) *service_errors.ServiceError {
	err := u.repo.UpdateEmail(ctx, pk, data.Email)
	if err != nil {
		return service_errors.MapDBErrToServiceErr(err, "update user.email")
	}
	return nil
}

func (u *UserService) UpdateBio(
	ctx context.Context, pk uint64, data *dto.UpdateBioRequest,
) *service_errors.ServiceError {
	err := u.repo.UpdateBio(ctx, pk, data.Bio)
	if err != nil {
		return service_errors.MapDBErrToServiceErr(err, "update user.bio")
	}
	return nil
}

func (u *UserService) UpdatePassword(
	ctx context.Context, pk uint64, data *dto.UpdatePasswordRequest,
) *service_errors.ServiceError {
	usersHashedPsw, err := u.repo.GetHashedPassword(ctx, pk)
	if err != nil {
		return service_errors.MapDBErrToServiceErr(err, "get user's hashed password")
	}

	if err := u.pswHasher.Verify(usersHashedPsw, data.OldPassword); err != nil {
		return service_errors.InvalidOldPassword
	}

	newHashedPsw, err := u.pswHasher.Hash(data.Password)
	if err != nil {
		u.logger.Errorf("failed to hash password. reason: %s", err.Error())
		return service_errors.InternalServerError
	}

	err = u.repo.UpdatePassword(ctx, pk, newHashedPsw)
	if err != nil {
		return service_errors.MapDBErrToServiceErr(err, "update user.password")
	}

	return nil
}

func (u *UserService) UpdateEnabled(
	ctx context.Context, pk uint64, data *dto.UpdateEnabledRequest,
) *service_errors.ServiceError {
	err := u.repo.UpdateEnabled(ctx, pk, *data.Enabled)
	if err != nil {
		return service_errors.MapDBErrToServiceErr(err, "update user.enabled")
	}

	redisErr := u.userInfoCache.UpdateEnabled(ctx, pk, *data.Enabled)
	if redisErr != nil {
		u.logger.Error(redisErr.Error())
	}

	return nil
}

func (u *UserService) Delete(ctx context.Context, pk uint64) *service_errors.ServiceError {
	err := u.repo.Delete(ctx, pk)
	if err != nil {
		return service_errors.MapDBErrToServiceErr(err, "delete user")
	}

	redisErr := u.userInfoCache.DeleteUserInfo(ctx, pk)
	if redisErr != nil {
		u.logger.Error(redisErr.Error())
	}

	return nil
}

func (u *UserService) CheckUsernameExists(ctx context.Context, username string) (bool, *service_errors.ServiceError) {
	exists, err := u.repo.CheckUsernameExists(ctx, username)
	if err != nil {
		return false, service_errors.MapDBErrToServiceErr(err, "check user.username exists")
	}

	return exists, nil
}

func (u *UserService) CheckEmailExists(ctx context.Context, email string) (bool, *service_errors.ServiceError) {
	exists, err := u.repo.CheckEmailExists(ctx, email)
	if err != nil {
		return false, service_errors.MapDBErrToServiceErr(err, "check user.email exists")
	}

	return exists, nil
}

// func (u *UserService) GetList()

func (u *UserService) GetByID(
	ctx context.Context, pk uint64,
) (*dto.UserDetails, *service_errors.ServiceError) {
	// try to fetch from cache first (maybe)
	return getByID(ctx, pk, "user", u.repo.GetByID, dto.ToUserDetails)
}

func (u *UserService) GetByUsername(
	ctx context.Context, username string,
) (*dto.UserDetails, *service_errors.ServiceError) {
	user, err := u.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, service_errors.MapDBErrToServiceErr(err, "get user by username")
	}

	return dto.ToUserDetails(user), nil
}

func (u *UserService) GetByEmail(
	ctx context.Context, email string,
) (*dto.UserDetails, *service_errors.ServiceError) {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, service_errors.MapDBErrToServiceErr(err, "get user by email")
	}

	return dto.ToUserDetails(user), nil
}

// func (u *UserService) GetByIDWithCountOfAllReferencedObjects()

// -----------------------------------------------------------------------------
// other sources that has FK to `User`

func (u *UserService) GetPosts(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.PostsList], *service_errors.ServiceError) {
	return getListOfOuterResourceByFK(
		ctx, fk, page, "get posts of user", u.repo.GetPosts, dto.ToPostsList,
	)
}

func (u *UserService) GetOwnedLists(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.ListsList], *service_errors.ServiceError) {
	return getListOfOuterResourceByFK(
		ctx, fk, page, "get owned-lists of user", u.repo.GetOwnedLists, dto.ToListsList,
	)
}

func (u *UserService) GetSavedLists(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.ListsList], *service_errors.ServiceError) {
	return getListOfOuterResourceByFK(
		ctx, fk, page, "get saved-lists of user", u.repo.GetSavedLists, dto.ToListsList,
	)
}

func (u *UserService) GetAllLists(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.ListsList], *service_errors.ServiceError) {
	// return getListOfOuterResourceByFK(
	// 	ctx, fk, page, "get all-lists of user", u.repo.GetAllLists, dto.ToListsList,
	// )
	return nil, service_errors.NotImplemented
}

func (u *UserService) GetComments(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.CommentList], *service_errors.ServiceError) {
	return getListOfOuterResourceByFK(
		ctx, fk, page, "get comments of user", u.repo.GetComments, dto.ToCommentList,
	)
}

func (u *UserService) GetLikes(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.PostsList], *service_errors.ServiceError) {
	return getListOfOuterResourceByFK(
		ctx, fk, page, "get likes of user", u.repo.GetLikes, dto.ToPostsList,
	)
}

func (u *UserService) GetFollowers(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.UsersList], *service_errors.ServiceError) {
	return getListOfOuterResourceByFK(
		ctx, fk, page, "get followers of user", u.repo.GetFollowers, dto.ToUsersList,
	)
}

func (u *UserService) GetFollowings(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.UsersList], *service_errors.ServiceError) {
	return getListOfOuterResourceByFK(
		ctx, fk, page, "get followings of user", u.repo.GetFollowings, dto.ToUsersList,
	)
}

func (u *UserService) GetLinks(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.LinksList], *service_errors.ServiceError) {
	return getListOfOuterResourceByFK(
		ctx, fk, page, "get links of user", u.repo.GetLinks, dto.ToLinksList,
	)
}
