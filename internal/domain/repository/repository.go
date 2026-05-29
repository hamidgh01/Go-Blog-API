package repository

import (
	"context"

	"Go-Blog-API/internal/domain/entity"
)

type UserRepository interface {
	creator[entity.User]
	UpdateUsername(ctx context.Context, pk uint64, username string) (*entity.User, error)
	UpdateEmail(ctx context.Context, pk uint64, email string) (*entity.User, error)
	UpdateBio(ctx context.Context, pk uint64, bio string) (*entity.User, error)
	UpdatePassword(ctx context.Context, pk uint64, password string) (*entity.User, error)
	UpdateEnabled(ctx context.Context, pk uint64, enabled bool) (*entity.User, error)
	deleter
	getterByID[entity.User]
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
}
