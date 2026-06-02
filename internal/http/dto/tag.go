package dto

import "Go-Blog-API/internal/domain/entity"

// -----------------------------------------------------------------------------
// Request DTOs

type CreateTagRequest struct {
	Name string `json:"name" binding:"required,tag_pattern"`
}

func NewCreateTagRequest() *CreateTagRequest {
	return new(CreateTagRequest)
}

type BulkCreateTagRequest struct {
	Name []string `json:"name" binding:"required,unique,min=1,max=10,dive,tag_pattern"`
}

func NewBulkCreateTagRequest() *BulkCreateTagRequest {
	return new(BulkCreateTagRequest)
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
