package services

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
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
) (*dto.UserDetails, *service_errors.ServiceError) {
	return nil, nil
}

func (u *UserService) UpdateEmail(
	ctx context.Context, pk uint64, data *dto.UpdateEmailRequest,
) (*dto.UserDetails, *service_errors.ServiceError) {
	return nil, nil
}

func (u *UserService) UpdateBio(
	ctx context.Context, pk uint64, data *dto.UpdateBioRequest,
) (*dto.UserDetails, *service_errors.ServiceError) {
	return nil, nil
}

func (u *UserService) UpdatePassword(
	ctx context.Context, pk uint64, data *dto.UpdatePasswordRequest,
) (*dto.UserDetails, *service_errors.ServiceError) {
	return nil, nil
}

func (u *UserService) UpdateEnabled(
	ctx context.Context, pk uint64, data *dto.UpdateEnabledRequest,
) (*dto.UserDetails, *service_errors.ServiceError) {
	return nil, nil
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
