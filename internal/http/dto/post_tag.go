package dto

import "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"

// -----------------------------------------------------------------------------
// Request DTOs

type AssociatePostWithTagsRequest struct {
	TagIDs []uint64 `json:"tag_ids" binding:"required,unique,min=1,max=10,dive,gt=0"`
}

func NewAssociatePostWithTagsRequest() *AssociatePostWithTagsRequest {
	return new(AssociatePostWithTagsRequest)
}

type DissociatePostWithTagsRequest struct {
	TagIDs []uint64 `json:"tag_ids" binding:"required,unique,min=1,max=10,dive,gt=0"`
}

func NewDissociatePostWithTagsRequest() *DissociatePostWithTagsRequest {
	return new(DissociatePostWithTagsRequest)
}

// -----------------------------------------------------------------------------
// Response DTOs

type ListOfPostsThatHaveATag struct {
	Posts PostsList `json:"posts"`
}

func ToListOfPostsThatHaveATag(posts []*entity.Post) *ListOfPostsThatHaveATag {
	return &ListOfPostsThatHaveATag{Posts: ToPostsList(posts)}
}

type TagsOfAPost struct {
	Tags TagsList `json:"tags"`
}

func ToTagsOfAPost(tags []*entity.Tag) *TagsOfAPost {
	return &TagsOfAPost{Tags: ToTagsList(tags)}
}
