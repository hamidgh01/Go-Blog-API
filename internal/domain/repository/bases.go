package repository

import (
	"context"

	"Go-Blog-API/internal/domain"
	e "Go-Blog-API/internal/domain/entity"
)

type creator[TEntity e.TDBEntities] interface {
	Create(ctx context.Context, entity *TEntity) (*TEntity, error)
}

type updater[TEntity e.TDBEntities] interface {
	Update(ctx context.Context, pk uint64, data *TEntity) (*TEntity, error)
}

type deleter interface {
	Delete(ctx context.Context, pk uint64) error
}

type getterByID[TEntity e.TDBEntities] interface {
	GetByID(ctx context.Context, pk uint64) (*TEntity, error)
}

type listGetterByFK[TEntity e.TDBEntities] interface {
	GetListByFK(ctx context.Context, FK uint64, page *domain.PaginationQueryParams) (*domain.PagedList[TEntity], error)
}

// type listGetterByM2MKey[TEntity e.TDBEntities] interface {
// 	GetListByM2MKey()
// }

// type listGetterByFilter[TEntity TDBEntities] interface {
// 	GetListByFilter()
// }
