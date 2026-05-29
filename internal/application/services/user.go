package services

import (
	"context"

	"Go-Blog-API/internal/domain/repository"
	"Go-Blog-API/internal/http/dto"
)

type UserService struct {
	repo repository.UserRepository
	// hasher
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (u *UserService) Create(ctx context.Context, data *dto.CreateUserRequest) (*dto.UserResponse, error) {
	return nil, nil
}

func (u *UserService) UpdateUsername(ctx context.Context, pk uint64, data *dto.UpdateUsernameRequest) (*dto.UserResponse, error) {
	return nil, nil
}

func (u *UserService) UpdateEmail(ctx context.Context, pk uint64, data *dto.UpdateEmailRequest) (*dto.UserResponse, error) {
	return nil, nil
}

func (u *UserService) UpdateBio(ctx context.Context, pk uint64, data *dto.UpdateBioRequest) (*dto.UserResponse, error) {
	return nil, nil
}

func (u *UserService) UpdatePassword(ctx context.Context, pk uint64, data *dto.UpdatePasswordRequest) (*dto.UserResponse, error) {
	return nil, nil
}

func (u *UserService) UpdateEnabled(ctx context.Context, pk uint64, data *dto.UpdatePasswordRequest) (*dto.UserResponse, error) {
	return nil, nil
}

func (u *UserService) Delete(ctx context.Context, pk uint64) error {
	return nil
}

// func (u *UserService) GetList()

func (u *UserService) GetByID(ctx context.Context, pk uint64) (*dto.UserResponse, error) {
	return nil, nil
}

func (u *UserService) GetByUsername(ctx context.Context, username string) (*dto.UserResponse, error) {
	return nil, nil
}

func (u *UserService) GetByEmail(ctx context.Context, email string) (*dto.UserResponse, error) {
	return nil, nil
}
