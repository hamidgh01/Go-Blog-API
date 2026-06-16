package dto

import "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"

// -----------------------------------------------------------------------------
// Request DTOs

type SavePostRequest struct {
	ListID uint64 `json:"list_id" binding:"required,gt=0"`
}

func NewSavePostRequest() *SavePostRequest {
	return new(SavePostRequest)
}

type UnsavePostRequest struct {
	ListID uint64 `json:"list_id" binding:"required,gt=0"`
}

func NewUnsavePostRequest() *UnsavePostRequest {
	return new(UnsavePostRequest)
}

// -----------------------------------------------------------------------------
// Response DTOs

type ListOfSavedPosts struct {
	Posts PostsList `json:"posts"`
}

func ToListOfSavedPosts(posts []*entity.Post) *ListOfSavedPosts {
	return &ListOfSavedPosts{Posts: ToPostsList(posts)}
}

type ListOfListsAPostSavedIn struct {
	Lists ListsList `json:"lists"`
}

func ToListOfListsAPostSavedIn(lists []*entity.List) *ListOfListsAPostSavedIn {
	return &ListOfListsAPostSavedIn{Lists: ToListsList(lists)}
}
