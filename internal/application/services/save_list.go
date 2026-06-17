package services

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
)

type SaveListService struct {
	repo repository.SaveListRepository
}

func NewSaveListService(r repository.SaveListRepository) *SaveListService {
	return &SaveListService{repo: r}
}

func (sl *SaveListService) Save(ctx context.Context, targetListID uint64) *service_errors.ServiceError {
	currentUserID := ctx.Value("currentUserID").(uint64)
	entity := &entity.UsersSavedListsM2M{UserID: currentUserID, ListID: targetListID}

	// ToDo: 2 condition to check:
	// check current user isn't the owner of the list (that is going to be saved)
	// check the list is going to be saved is not private

	return createOrDeleteM2MRelationship(ctx, "save a list", entity, sl.repo.Create)
}

func (sl *SaveListService) Unsave(ctx context.Context, targetListID uint64) *service_errors.ServiceError {
	currentUserID := ctx.Value("currentUserID").(uint64)
	entity := &entity.UsersSavedListsM2M{UserID: currentUserID, ListID: targetListID}

	return createOrDeleteM2MRelationship(ctx, "unsave a list", entity, sl.repo.Delete)
}
