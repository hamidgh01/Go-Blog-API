package dto

import (
	"time"

	"Go-Blog-API/internal/domain/entity"
)

// -----------------------------------------------------------------------------
// Request DTOs

type CreatePostRequest struct {
	Title     string `json:"title" binding:"required,max=200"`
	Content   string `json:"content,omitempty"`
	Status    string `json:"status,omitempty" binding:"oneof=draft published"`
	IsPrivate bool   `json:"is_private,omitempty"`
}

func NewCreatePostRequest() *CreatePostRequest {
	return new(CreatePostRequest)
}

type UpdatePostRequest struct {
	Title   string `json:"title,omitempty" binding:"min=3,max=250"`
	Content string `json:"content,omitempty"`
}

func NewUpdatePostRequest() *UpdatePostRequest {
	return new(UpdatePostRequest)
}

type UpdatePostStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=published rejected deleted"`
}

func NewUpdatePostStatusRequest() *UpdatePostStatusRequest {
	return new(UpdatePostStatusRequest)
}

type UpdatePostPrivacyRequest struct {
	IsPrivate bool `json:"is_private" binding:"required"`
}

func NewUpdatePostPrivacyRequest() *UpdatePostPrivacyRequest {
	return new(UpdatePostPrivacyRequest)
}

// -----------------------------------------------------------------------------
// Response DTOs

type PostBrief struct {
	ID               uint64     `json:"id"`
	Title            string     `json:"title"`
	CreatedAt        time.Time  `json:"created_at"`
	ModifiedAt       time.Time  `json:"modified_at"`
	FirstPublishedAt time.Time  `json:"first_published_at"`
	User             *UserBrief `json:"user"`
}

func ToPostBrief(p *entity.Post) *PostBrief {
	return &PostBrief{
		ID:               p.ID,
		Title:            p.Title,
		CreatedAt:        p.CreatedAt,
		ModifiedAt:       p.ModifiedAt.Time,
		FirstPublishedAt: p.FirstPublishedAt.Time,
		User:             ToUserBrief(p.User),
	}
}

type PostDetails struct {
	*PostBrief
	Content   string `json:"content,omitempty"`
	IsPrivate bool   `json:"is_private"`
	Status    string `json:"status"`
}

func ToPostDetails(p *entity.Post) *PostDetails {
	return &PostDetails{
		PostBrief: ToPostBrief(p),
		Content:   p.Content,
		IsPrivate: p.IsPrivate,
		Status:    string(p.Status),
	}
}

type PostDetailsWithCountOfReferencedObjects struct {
	*PostDetails
	RefObjCounts  map[entity.CountKey]int `json:"referenced_objects_count"`
	LikeCount     int                     `json:"like_count"`
	LikedByViewer bool                    `json:"like_by_viewer"`
	SavedByViewer bool                    `json:"saved_by_viewer"`
}

func ToPostDetailsWithCountOfReferencedObjects(
	p *entity.Post, refObjCounts map[entity.CountKey]int, likeCount int, likedByViewer bool, savedByViewer bool,
) *PostDetailsWithCountOfReferencedObjects {
	return &PostDetailsWithCountOfReferencedObjects{
		PostDetails:   ToPostDetails(p),
		RefObjCounts:  refObjCounts,
		LikeCount:     likeCount,
		LikedByViewer: likedByViewer,
		SavedByViewer: savedByViewer,
	}
}

type PostsList []*PostBrief

func ToPostsList(posts []*entity.Post) PostsList {
	postsList := make(PostsList, len(posts))
	for _, post := range posts {
		postsList = append(postsList, ToPostBrief(post))
	}
	return postsList
}
