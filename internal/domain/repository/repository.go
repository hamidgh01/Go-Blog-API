package repository

import (
	"context"

	d "github.com/hamidgh01/Go-Blog-API/internal/domain"
	e "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
)

type UserRepository interface {
	creator[e.User]

	UpdateUsername(ctx context.Context, pk uint64, username string) error
	UpdateEmail(ctx context.Context, pk uint64, email string) error
	UpdateBio(ctx context.Context, pk uint64, bio string) error
	UpdatePassword(ctx context.Context, pk uint64, password string) error
	UpdateEnabled(ctx context.Context, pk uint64, enabled bool) error

	deleter

	CheckUsernameExists(ctx context.Context, username string) (bool, error)
	CheckEmailExists(ctx context.Context, email string) (bool, error)
	CheckIsEnabled(ctx context.Context, pk uint64) (bool, error)
	CheckIsSuperuser(ctx context.Context, pk uint64) (bool, error)
	GetHashedPassword(ctx context.Context, pk uint64) (string, error)
	GetUserForLoginVerification(ctx context.Context, identifier string) (*e.User, error)

	getterByID[e.User]
	GetByUsername(ctx context.Context, username string) (*e.User, error)
	GetByEmail(ctx context.Context, email string) (*e.User, error)
	getterByIDWithCountOfAllReferencedObjects[e.User]

	// get list of other sources (have FK to a `User`)
	GetPosts(ctx context.Context, fk uint64, page *d.PaginationQueryParams) (*d.PagedList[e.Post], error)
	GetLists(ctx context.Context, fk uint64, page *d.PaginationQueryParams) (*d.PagedList[e.List], error)
	GetSavedLists(ctx context.Context, fk uint64, page *d.PaginationQueryParams) (*d.PagedList[e.List], error)
	GetComments(ctx context.Context, fk uint64, page *d.PaginationQueryParams) (*d.PagedList[e.Comment], error)
	GetLikes(ctx context.Context, fk uint64, page *d.PaginationQueryParams) (*d.PagedList[e.Post], error)
	GetFollowers(ctx context.Context, fk uint64, page *d.PaginationQueryParams) (*d.PagedList[e.User], error)
	GetFollowings(ctx context.Context, fk uint64, page *d.PaginationQueryParams) (*d.PagedList[e.User], error)
	GetLinks(ctx context.Context, fk uint64, page *d.PaginationQueryParams) (*d.PagedList[e.Link], error)
}

type LinkRepository interface {
	creator[e.Link]
	updater[e.Link]
	deleter
	getterByID[e.Link]

	ownerIDGetter
}

type PostRepository interface {
	creator[e.Post]

	updater[e.Post]
	PublishDraftPost(ctx context.Context, pk uint64) error
	UpdateStatus(ctx context.Context, pk uint64, status e.PostStatus) error
	UpdatePrivacy(ctx context.Context, pk uint64, isPrivate bool) error

	deleter

	getterByID[e.Post]
	getterByIDWithCountOfAllReferencedObjects[e.Post]

	// get list of other sources (have FK to a `Post`)
	GetComments(ctx context.Context, fk uint64, page *d.PaginationQueryParams) (*d.PagedList[e.Comment], error)
	GetLikes(ctx context.Context, fk uint64, page *d.PaginationQueryParams) (*d.PagedList[e.User], error)
	GetListsThatSavedThisPost(ctx context.Context, fk uint64, page *d.PaginationQueryParams) (*d.PagedList[e.List], error)
	GetTags(ctx context.Context, fk uint64, page *d.PaginationQueryParams) (*d.PagedList[e.Tag], error)

	ownerIDGetter
}

type TagRepository interface {
	Create(ctx context.Context, entities []*e.Tag) ([]*e.Tag, error)

	getterByID[e.Tag]
	GetByName(ctx context.Context, name string) (*e.Tag, error)
	getterByIDWithCountOfAllReferencedObjects[e.Tag]
	GetByNameWithCountOfAllReferencedObjects(ctx context.Context, name string) (*e.DBEntityWithCountOfReferencedObjects[e.Tag], error)

	// `Tag` is only referenced by `Post` (via `PostTagsM2M`)
	GetPosts(ctx context.Context, fk uint64, page *d.PaginationQueryParams) (*d.PagedList[e.Post], error)
}

type CommentRepository interface {
	creator[e.Comment]

	updater[e.Comment]
	UpdateStatus(ctx context.Context, pk uint64, status e.CommentStatus) error

	deleter

	getterByID[e.Comment]
	getterByIDWithCountOfAllReferencedObjects[e.Comment]

	// `Comment` is only referenced by itself (its Replies)
	GetReplies(ctx context.Context, fk uint64, page *d.PaginationQueryParams) (*d.PagedList[e.Comment], error)

	ownerIDGetter
}

type ListRepository interface {
	creator[e.List]

	updater[e.List]
	UpdatePrivacy(ctx context.Context, pk uint64, isPrivate bool) error

	deleter

	getterByID[e.List]
	getterByIDWithCountOfAllReferencedObjects[e.List]

	// get list of other sources (have FK to a `List`) (`Post` via `SavedPostsM2M`, `User` via `UsersSavedListsM2M`)
	GetSavedPosts(ctx context.Context, fk uint64, page *d.PaginationQueryParams) (*d.PagedList[e.Post], error)
	GetUsersWhoSaved(ctx context.Context, fk uint64, page *d.PaginationQueryParams) (*d.PagedList[e.User], error)

	ownerIDGetter
}
