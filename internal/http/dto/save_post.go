package dto

import "Go-Blog-API/internal/domain/entity"

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
