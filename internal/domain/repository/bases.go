package repository

import (
	"context"

	e "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
)

type creator[TEntity e.TDBEntities] interface {
	Create(ctx context.Context, entity *TEntity, userID uint64) (*TEntity, error)
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

// get an object and the count all referenced objects that is related to the object.
// e.g. get a `User` and the count of its related `Posts`, `Lists`, `Followers`, etc.
// e.g. get a `Post` and the count of its related `Likes`, `Comments`, `Tags`, etc.
type getterByIDWithCountOfAllReferencedObjects[TEntity e.TDBEntities] interface {
	GetByIDWithCountOfAllReferencedObjects(
		ctx context.Context, pk uint64,
	) (*e.DBEntityWithCountOfReferencedObjects[TEntity], error)
}

// type listGetterByM2MKey[TEntity e.TDBEntities] interface {
// 	GetListByM2MKey()
// }

// type listGetterByFilter[TEntity TDBEntities] interface {
// 	GetListByFilter()
// }
