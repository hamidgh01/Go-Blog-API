package dto

import "Go-Blog-API/internal/domain/entity"

// -----------------------------------------------------------------------------
// Request DTOs

type CreateTagsRequest struct {
	Tags []string `json:"tags" binding:"required,unique,min=1,max=10,dive,tag_pattern"`
}

func NewCreateTagsRequest() *CreateTagsRequest {
	return new(CreateTagsRequest)
}

// -----------------------------------------------------------------------------
// Response DTOs

type TagDetails struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func ToTagDetails(t *entity.Tag) *TagDetails {
	return &TagDetails{
		ID:   t.ID,
		Name: t.Name,
	}
}

type TagsList []*TagDetails

func ToTagsList(tags []*entity.Tag) TagsList {
	tagsList := make(TagsList, len(tags))
	for _, tag := range tags {
		tagsList = append(tagsList, ToTagDetails(tag))
	}
	return tagsList
}
