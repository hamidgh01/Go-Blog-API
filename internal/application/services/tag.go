package services

import (
	"context"
	"slices"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	d "github.com/hamidgh01/Go-Blog-API/internal/domain"
	e "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
	"github.com/hamidgh01/Go-Blog-API/internal/http/generics"
)

type TagService struct {
	repo repository.TagRepository
}

func NewTagService(r repository.TagRepository) *TagService {
	return &TagService{repo: r}
}

func containTagName(tagList []*e.Tag, tag *e.Tag) bool {
	for _, t := range tagList {
		if tag.Name == t.Name {
			return true
		}
	}

	return false
}

func (t *TagService) Create(
	ctx context.Context, data *dto.CreateTagsRequest,
) (dto.TagsList, *service_errors.ServiceError) {
	enteredTags := make([]*e.Tag, 0, len(data.Tags))
	for _, tag := range data.Tags {
		enteredTags = append(enteredTags, &e.Tag{Name: tag})
	}

	createdTags, err := t.repo.Create(ctx, enteredTags)
	if err != nil {
		return nil, service_errors.MapDBErrToServiceErr(err, "create tags")
	}

	previouslyExistentTags := make([]*e.Tag, 0, len(enteredTags))
	for _, entTag := range enteredTags {
		if containTagName(createdTags, entTag) {
			continue
		}

		previouslyExistentTags = append(previouslyExistentTags, entTag)
	}

	var fetchedExistentTags []*e.Tag
	if len(previouslyExistentTags) > 0 {
		fetchedExistentTags, err = t.repo.GetListOfTagsByNames(ctx, previouslyExistentTags)
		if err != nil {
			return nil, service_errors.MapDBErrToServiceErr(err, "get list of tags by names")

		}
	}

	allTags := slices.Concat(createdTags, fetchedExistentTags)
	return dto.ToTagsList(allTags), nil
}

func (t *TagService) GetByID(
	ctx context.Context, pk uint64,
) (*dto.TagDetails, *service_errors.ServiceError) {
	return getByID(ctx, pk, "tag", t.repo.GetByID, dto.ToTagDetails)
}

func (t *TagService) GetByName(
	ctx context.Context, name string,
) (*dto.TagDetails, *service_errors.ServiceError) {
	tag, err := t.repo.GetByName(ctx, name)
	if err != nil {
		return nil, service_errors.MapDBErrToServiceErr(err, "get tag by name")
	}

	return dto.ToTagDetails(tag), nil
}

// -----------------------------------------------------------------------------
// other sources that has FK to `Tag`

func (t *TagService) GetPosts(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.PostsList], *service_errors.ServiceError) {
	// implement later
	return nil, nil
}
