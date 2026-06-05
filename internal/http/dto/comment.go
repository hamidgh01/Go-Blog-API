package dto

import (
	"time"

	"github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
)

// -----------------------------------------------------------------------------
// Request DTOs

type CreateCommentRequest struct {
	Content    string `json:"content" binding:"required,max=1500"`
	ParentType string `json:"parent_type" binding:"required,oneof=post comment"`
	ParentID   uint64 `json:"parent_id" binding:"required,gt=0"`
}

func NewCreateCommentRequest() *CreateCommentRequest {
	return new(CreateCommentRequest)
}

type UpdateCommentRequest struct {
	Content string `json:"content" binding:"required,max=1500"`
}

func NewUpdateCommentRequest() *UpdateCommentRequest {
	return new(UpdateCommentRequest)
}

type UpdateCommentStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=published hidden deleted"`
}

func NewUpdateCommentStatusRequest() *UpdateCommentStatusRequest {
	return new(UpdateCommentStatusRequest)
}

// -----------------------------------------------------------------------------
// Response DTOs

type CommentDetails struct {
	ID              uint64     `json:"id"`
	Content         string     `json:"content"`
	CreatedAt       time.Time  `json:"created_at"`
	ModifiedAt      time.Time  `json:"modified_at"`
	Status          string     `json:"status"`
	User            *UserBrief `json:"user"`
	PostParentID    uint64     `json:"post_parent_id,omitempty"`
	PostParent      *PostBrief `json:"post_parent,omitempty"`
	CommentParentID uint64     `json:"comment_parent_id,omitempty"`
}

func ToCommentDetails(c *entity.Comment) *CommentDetails {
	cd := &CommentDetails{
		ID:         c.ID,
		Content:    c.Content,
		CreatedAt:  c.CreatedAt,
		ModifiedAt: c.ModifiedAt.Time,
		Status:     string(c.Status),
		User:       ToUserBrief(c.User),
	}

	if c.PostParentID == 0 {
		cd.CommentParentID = c.CommentParentID
		return cd
	} else {
		cd.PostParentID = c.PostParentID
		cd.PostParent = ToPostBrief(c.PostParent)
		return cd
	}
}

type CommentDetailsWithRepliesCount struct {
	*CommentDetails
	RepliesCount int `json:"replies_count"`
}

func ToCommentDetailsWithRepliesCount(c *entity.Comment, repliesCount int) *CommentDetailsWithRepliesCount {
	return &CommentDetailsWithRepliesCount{
		CommentDetails: ToCommentDetails(c),
		RepliesCount:   repliesCount,
	}
}

type CommentList []*CommentDetails

func ToCommentList(comments []*entity.Comment) CommentList {
	commentsList := make(CommentList, len(comments))
	for _, comment := range comments {
		commentsList = append(commentsList, ToCommentDetails(comment))
	}
	return commentsList
}
