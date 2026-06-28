package services

import (
	"context"
	"database/sql"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	d "github.com/hamidgh01/Go-Blog-API/internal/domain"
	e "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
	"github.com/hamidgh01/Go-Blog-API/internal/http/generics"
)

type PostService struct {
	repo repository.PostRepository
}

func NewPostService(r repository.PostRepository) *PostService {
	return &PostService{repo: r}
}

func (p *PostService) Create(
	ctx context.Context, data *dto.CreatePostRequest,
) (*dto.PostDetails, *service_errors.ServiceError) {
	//
	userID := ctx.Value("currentUserID").(uint64)
	post := &e.Post{Title: data.Title, Status: e.PostStatus(data.Status), UserID: userID}

	if data.Content == "" {
		post.Content = sql.NullString{Valid: false}
	} else {
		post.Content = sql.NullString{String: data.Content, Valid: true}
	}

	if data.IsPrivate == nil {
		post.IsPrivate = false
	} else {
		post.IsPrivate = *data.IsPrivate
	}

	username := ctx.Value("currentUserUsername").(string)
	post.User = &e.User{ID: userID, Username: username}

	return create(ctx, "post", post, p.repo.Create, dto.ToPostDetails)
}

func (p *PostService) Update(
	ctx context.Context, pk uint64, data *dto.UpdatePostRequest,
) *service_errors.ServiceError {
	post := &e.Post{Title: data.Title}

	if data.Content == "" {
		post.Content = sql.NullString{Valid: false}
	} else {
		post.Content = sql.NullString{String: data.Content, Valid: true}
	}

	return update(ctx, pk, "post", post, p.repo.Update)
}

func (p *PostService) PublishDraftPost(
	ctx context.Context, pk uint64, data *dto.UpdatePostStatusRequest,
) *service_errors.ServiceError {
	err := p.repo.PublishDraftPost(ctx, pk)
	if err != nil {
		return service_errors.MapDBErrToServiceErr(err, "publish post")
	}

	return nil
}

func (p *PostService) UpdateStatus(
	ctx context.Context, pk uint64, data *dto.UpdatePostStatusRequest,
) *service_errors.ServiceError {
	err := p.repo.UpdateStatus(ctx, pk, e.PostStatus(data.Status))
	if err != nil {
		return service_errors.MapDBErrToServiceErr(err, "update post status")
	}

	return nil
}

func (p *PostService) UpdatePrivacy(
	ctx context.Context, pk uint64, data *dto.UpdatePostPrivacyRequest,
) *service_errors.ServiceError {
	err := p.repo.UpdatePrivacy(ctx, pk, *data.IsPrivate)
	if err != nil {
		return service_errors.MapDBErrToServiceErr(err, "update post privacy")
	}

	return nil
}

func (p *PostService) Delete(ctx context.Context, pk uint64) *service_errors.ServiceError {
	return delete(ctx, pk, "post", p.repo.Delete)
}

func (p *PostService) GetByID(
	ctx context.Context, pk uint64,
) (*dto.PostDetails, *service_errors.ServiceError) {
	return getByID(ctx, pk, "post", p.repo.GetByID, dto.ToPostDetails)
}

// -----------------------------------------------------------------------------
// other sources that has FK to `Post`

func (p *PostService) GetComments(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.CommentList], *service_errors.ServiceError) {
	return getListOfOuterResourceByFK(
		ctx, fk, page, "get comments of post", p.repo.GetComments, dto.ToCommentList,
	)
}

func (p *PostService) GetLikes(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.UsersList], *service_errors.ServiceError) {
	return getListOfOuterResourceByFK(
		ctx, fk, page, "get likes of post", p.repo.GetLikes, dto.ToUsersList,
	)
}

func (p *PostService) GetTags(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.TagsList], *service_errors.ServiceError) {
	return getListOfOuterResourceByFK(
		ctx, fk, page, "get tags of post", p.repo.GetTags, dto.ToTagsList,
	)
}

func (p *PostService) GetLists(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.ListsList], *service_errors.ServiceError) {
	return getListOfOuterResourceByFK(
		ctx,
		fk,
		page,
		"get lists that saved this post",
		p.repo.GetListsThatSavedThisPost,
		dto.ToListsList,
	)
}

// -----------------------------------------------------------------------------

func (p *PostService) GetOwnerID(
	ctx context.Context, pk uint64,
) (uint64, *service_errors.ServiceError) {
	return getOwnerID(ctx, pk, "post", p.repo.GetOwnerID)
}
