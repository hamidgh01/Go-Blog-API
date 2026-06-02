package repository

import (
	"context"

	e "Go-Blog-API/internal/domain/entity"
)

type UserRepository[TEntity e.TDBEntities] interface {
	creator[e.User]

	UpdateUsername(ctx context.Context, pk uint64, username string) (*e.User, error)
	UpdateEmail(ctx context.Context, pk uint64, email string) (*e.User, error)
	UpdateBio(ctx context.Context, pk uint64, bio string) (*e.User, error)
	UpdatePassword(ctx context.Context, pk uint64, password string) (*e.User, error)
	UpdateEnabled(ctx context.Context, pk uint64, enabled bool) (*e.User, error)

	// ExistEmail(...)
	// ExistUsername(...)
	// VerifyLoginRequest(...)

	deleter

	getterByID[e.User]
	GetByUsername(ctx context.Context, username string) (*e.User, error)
	GetByEmail(ctx context.Context, email string) (*e.User, error)
	getterByIDWithCountOfAllReferencedObjects[e.User]
	referredObjectListGetterByFK[TEntity] // tables referred User: all of other tables
}

type LinkRepository interface {
	creator[e.Link]
	updater[e.Link]
	deleter
	getterByID[e.Link]
}

type PostRepository[TEntity e.TDBEntities] interface {
	creator[e.Post]

	updater[e.Post]
	UpdateStatus(ctx context.Context, pk uint64, data *e.Post) (*e.Post, error)
	UpdatePrivacy(ctx context.Context, pk uint64, data *e.Post) (*e.Post, error)

	deleter

	getterByID[e.Post]
	getterByIDWithCountOfAllReferencedObjects[e.Post]
	referredObjectListGetterByFK[TEntity] // tables referred Post: comments, likes, tags, lists
}

type TagRepository interface {
	creator[e.Tag]
	BalkCreate(ctx context.Context, entity []*e.Tag) ([]*e.Tag, error)

	getterByID[e.Tag]
	GetByName(ctx context.Context, name string) (*e.User, error)
	getterByIDWithCountOfAllReferencedObjects[e.Tag]
	GetByNameWithCountOfAllReferencedObjects(ctx context.Context, name string) (*e.DBEntityWithCountOfReferencedObjects[e.Tag], error)
	referredObjectListGetterByFK[e.Post] // tags table is only referenced by posts table (via PostTagsM2M)
}

type CommentRepository interface {
	creator[e.Comment]

	updater[e.Comment]
	UpdateStatus(ctx context.Context, pk uint64, data *e.Comment) (*e.Comment, error)

	deleter

	getterByID[e.Comment]
	getterByIDWithCountOfAllReferencedObjects[e.Comment]
	referredObjectListGetterByFK[e.Comment] // comments table is only referenced by itself (its replies)
}

type listRepository[TEntity e.TDBEntities] interface {
	creator[e.List]

	updater[e.List]
	UpdatePrivacy(ctx context.Context, pk uint64, data *e.Link) (*e.Link, error)

	deleter

	getterByID[e.List]
	getterByIDWithCountOfAllReferencedObjects[e.List]
	referredObjectListGetterByFK[TEntity] // tables referred List: posts (SavedPostsM2M), users (UsersSavedListsM2M)
}
