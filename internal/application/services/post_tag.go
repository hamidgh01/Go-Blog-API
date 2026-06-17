package services

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	e "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
)

type PostTagsService struct {
	repo     repository.PostTagsRepository
	postRepo repository.PostRepository
}

func NewPostTagsService(
	r repository.PostTagsRepository, pr repository.PostRepository,
) *PostTagsService {
	return &PostTagsService{repo: r, postRepo: pr}
}

func (pt *PostTagsService) Associate(
	ctx context.Context, postID uint64, tagIDs *dto.AssociatePostWithTagsRequest,
) *service_errors.ServiceError {
	postOwnerID, err := pt.postRepo.GetOwnerID(ctx, postID)
	if err != nil {
		return service_errors.MapDBErrToServiceErr(err, "get post owner id")
	}

	currentUserID := ctx.Value("currentUserID").(uint64)
	if currentUserID != postOwnerID {
		return service_errors.PermissionDenied
	}

	entities := make([]*e.PostTagsM2M, 0, len(tagIDs.TagIDs))
	for _, tagID := range tagIDs.TagIDs {
		entity := &e.PostTagsM2M{TagID: tagID, PostID: postID}
		entities = append(entities, entity)
	}

	// ToDo: condition to check:
	// check the target post (to tag) is 'published' or 'draft'

	if err := pt.repo.Create(ctx, entities); err != nil {
		return service_errors.MapDBErrToServiceErr(err, "tagging a post (associate)")
	}

	return nil
}

func (pt *PostTagsService) Dissociate(
	ctx context.Context, postID uint64, tagIDs *dto.DissociatePostWithTagsRequest,
) *service_errors.ServiceError {
	postOwnerID, err := pt.postRepo.GetOwnerID(ctx, postID)
	if err != nil {
		return service_errors.MapDBErrToServiceErr(err, "get post owner id")
	}

	currentUserID := ctx.Value("currentUserID").(uint64)
	if currentUserID != postOwnerID {
		return service_errors.PermissionDenied
	}

	entities := make([]*e.PostTagsM2M, 0, len(tagIDs.TagIDs))
	for _, tagID := range tagIDs.TagIDs {
		entity := &e.PostTagsM2M{TagID: tagID, PostID: postID}
		entities = append(entities, entity)
	}

	// ToDo: condition to check:
	// check the target post (to untag) is 'published' or 'draft'

	if err := pt.repo.Delete(ctx, entities); err != nil {
		return service_errors.MapDBErrToServiceErr(err, "untag a post (dissociate)")
	}

	return nil
}
