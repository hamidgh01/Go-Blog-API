package repository

import (
	"context"

	e "Go-Blog-API/internal/domain/entity"
)

type TDBEntities interface {
	e.User | e.FollowsM2M | e.Link | e.Post | e.PostLikesM2M | e.Comment |
		e.Tag | e.PostTagsM2M | e.List | e.SavedPostsM2M | e.UsersSavedListsM2M
}

type creator[TEntity TDBEntities] interface {
	Create(ctx context.Context, entity *TEntity) (*TEntity, error)
}

type updater[TEntity TDBEntities] interface {
	Update(ctx context.Context, pk uint64, data *TEntity) (*TEntity, error)
}

type deleter interface {
	Delete(ctx context.Context, pk uint64) error
}

type getterByID[TEntity TDBEntities] interface {
	GetByID(ctx context.Context, pk uint64) (*TEntity, error)
}
