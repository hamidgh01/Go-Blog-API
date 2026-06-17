package services

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
)

type SavePostService struct {
	repo     repository.SavePostRepository
	listRepo repository.ListRepository
}

func NewSavePostService(
	r repository.SavePostRepository, lr repository.ListRepository,
) *SavePostService {
	return &SavePostService{repo: r, listRepo: lr}
}

func (sp *SavePostService) Save(
	ctx context.Context, postID uint64, listID *dto.SavePostRequest,
) *service_errors.ServiceError {
	listOwnerID, err := sp.listRepo.GetOwnerID(ctx, listID.ListID)
	if err != nil {
		return service_errors.MapDBErrToServiceErr(err, "get list owner id")
	}

	currentUserID := ctx.Value("currentUserID").(uint64)
	if currentUserID != listOwnerID {
		return service_errors.PermissionDenied
	}

	// another condition to check:
	// check the post is going to be saved is not private

	entity := &entity.SavedPostsM2M{ListID: listID.ListID, PostID: postID}

	return createOrDeleteM2MRelationship(ctx, "save a list", entity, sp.repo.Create)
}

func (sp *SavePostService) Unsave(
	ctx context.Context, postID uint64, listID *dto.UnsavePostRequest,
) *service_errors.ServiceError {
	listOwnerID, err := sp.listRepo.GetOwnerID(ctx, listID.ListID)
	if err != nil {
		return service_errors.MapDBErrToServiceErr(err, "get list owner id")
	}

	currentUserID := ctx.Value("currentUserID").(uint64)
	if currentUserID != listOwnerID {
		return service_errors.PermissionDenied
	}

	entity := &entity.SavedPostsM2M{ListID: listID.ListID, PostID: postID}

	return createOrDeleteM2MRelationship(ctx, "unsave a list", entity, sp.repo.Delete)
}
