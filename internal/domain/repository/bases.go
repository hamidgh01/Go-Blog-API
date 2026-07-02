package repository

import (
	"context"

	e "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
)

type creator[TEntity e.TDBEntities] interface {
	Create(ctx context.Context, entity *TEntity) (*TEntity, error)
}

type updater[TEntity e.TDBEntities] interface {
	Update(ctx context.Context, pk uint64, data *TEntity) error
}

type deleter interface {
	Delete(ctx context.Context, pk uint64) error
}

type getterByID[TEntity e.TDBEntities] interface {
	GetByID(ctx context.Context, pk uint64) (*TEntity, error)
}

type ownerIDGetter interface {
	GetOwnerID(ctx context.Context, pk uint64) (uint64, error)
}

// type listGetterByM2MKey[TEntity e.TDBEntities] interface {
// 	GetListByM2MKey()
// }

// type listGetterByFilter[TEntity TDBEntities] interface {
// 	GetListByFilter()
// }

type M2MEntityRepository[TEntity e.TM2MDBEntities] interface {
	Create(ctx context.Context, entity *TEntity) error
	Delete(ctx context.Context, entity *TEntity) error
}
