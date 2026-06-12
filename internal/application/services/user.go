package services

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	d "github.com/hamidgh01/Go-Blog-API/internal/domain"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
	"github.com/hamidgh01/Go-Blog-API/internal/http/generics"
)

type UserService struct {
	repo repository.UserRepository
	// hasher
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (u *UserService) Create(
	ctx context.Context, data *dto.CreateUserRequest,
) (*dto.UserDetails, *service_errors.ServiceError) {
	return nil, nil
}

func (u *UserService) UpdateUsername(
	ctx context.Context, pk uint64, data *dto.UpdateUsernameRequest,
) *service_errors.ServiceError {
	return nil
}

func (u *UserService) UpdateEmail(
	ctx context.Context, pk uint64, data *dto.UpdateEmailRequest,
) *service_errors.ServiceError {
	return nil
}

func (u *UserService) UpdateBio(
	ctx context.Context, pk uint64, data *dto.UpdateBioRequest,
) *service_errors.ServiceError {
	return nil
}

func (u *UserService) UpdatePassword(
	ctx context.Context, pk uint64, data *dto.UpdatePasswordRequest,
) *service_errors.ServiceError {
	return nil
}

func (u *UserService) UpdateEnabled(
	ctx context.Context, pk uint64, data *dto.UpdateEnabledRequest,
) *service_errors.ServiceError {
	return nil
}

func (u *UserService) Delete(ctx context.Context, pk uint64) *service_errors.ServiceError {
	return nil
}

// func (u *UserService) GetList()

func (u *UserService) GetByID(
	ctx context.Context, pk uint64,
) (*dto.UserDetails, *service_errors.ServiceError) {
	return nil, nil
}

func (u *UserService) GetByUsername(
	ctx context.Context, username string,
) (*dto.UserDetails, *service_errors.ServiceError) {
	return nil, nil
}

func (u *UserService) GetByEmail(
	ctx context.Context, email string,
) (*dto.UserDetails, *service_errors.ServiceError) {
	return nil, nil
}

// -----------------------------------------------------------------------------
// other sources that has FK to `User`

func (u *UserService) GetPosts(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.PostsList], *service_errors.ServiceError) {
	return nil, nil
}

func (u *UserService) GetLists(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.ListsList], *service_errors.ServiceError) {
	return nil, nil
}

func (u *UserService) GetSavedLists(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.ListsList], *service_errors.ServiceError) {
	return nil, nil
}

func (u *UserService) GetComments(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.CommentList], *service_errors.ServiceError) {
	return nil, nil
}

func (u *UserService) GetLikes(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.PostsList], *service_errors.ServiceError) {
	return nil, nil
}

func (u *UserService) GetFollowers(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.UsersList], *service_errors.ServiceError) {
	return nil, nil
}

func (u *UserService) GetFollowings(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.UsersList], *service_errors.ServiceError) {
	return nil, nil
}

func (u *UserService) GetLinks(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.LinksList], *service_errors.ServiceError) {
	return nil, nil
}
