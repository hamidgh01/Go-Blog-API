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

// get an object and the count all referenced objects that is related to the object.
// e.g. get a `User` and the count of its related `Posts`, `Lists`, `Followers`, etc.
// e.g. get a `Post` and the count of its related `Likes`, `Comments`, `Tags`, etc.
type getterByIDWithCountOfAllReferencedObjects[TEntity e.TDBEntities] interface {
	GetByIDWithCountOfAllReferencedObjects(
		ctx context.Context, pk uint64,
	) (*e.DBEntityWithCountOfReferencedObjects[TEntity], error)
}

// get list of referenced objects related to a specific object.
// e.g. get list of `Posts` related to a `User`
// e.g. get list of `Comments` related to a `Post`
type referredObjectListGetterByFK[TEntity e.TDBEntities] interface {
	GetListOfReferencedObject(
		ctx context.Context, fk uint64, referencedEntity e.ReferencedObjectKey, page *domain.PaginationQueryParams,
	) (*domain.PagedList[TEntity], error)
}

// type listGetterByM2MKey[TEntity e.TDBEntities] interface {
// 	GetListByM2MKey()
// }

// type listGetterByFilter[TEntity TDBEntities] interface {
// 	GetListByFilter()
// }
