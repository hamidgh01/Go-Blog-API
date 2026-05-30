package dto

import (
	"time"

	"Go-Blog-API/internal/domain/entity"
)

// -----------------------------------------------------------------------------
// Request DTOs

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,min=3,max=250"`
	Content string `json:"content,omitempty"`
	Status  string `json:"status,omitempty" binding:"oneof=draft published"`
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

type PostOut struct {
	ID            uint64    `json:"id"`
	Title         string    `json:"title"`
	CreatedAt     time.Time `json:"created_at"`
	ModifiedAt    time.Time `json:"modified_at"`
	UserID        uint64    `json:"user_id"`
	User          *UserOut  `json:"user"`
	SavedByViewer bool      `json:"saved_by_viewer"`
}

func ToPostOut(p *entity.Post, userOut *UserOut, savedByViewer bool) *PostOut {
	return &PostOut{
		ID:            p.ID,
		Title:         p.Title,
		CreatedAt:     p.CreatedAt,
		ModifiedAt:    p.ModifiedAt.Time,
		UserID:        p.UserID,
		User:          userOut,
		SavedByViewer: savedByViewer,
	}
}

type PostDetailsResponse struct {
	*PostOut
	Content       string `json:"content,omitempty"`
	LikeCount     int    `json:"like_count"`
	LikedByViewer bool   `json:"like_by_viewer"`
}

func ToPostDetailsResponse(
	p *entity.Post, userOut *UserOut, savedByViewer bool, likeCount int, likedByViewer bool,
) *PostDetailsResponse {
	return &PostDetailsResponse{
		PostOut:       ToPostOut(p, userOut, savedByViewer),
		Content:       p.Content,
		LikeCount:     likeCount,
		LikedByViewer: likedByViewer,
	}
}

type PostsListResponse []*PostOut
