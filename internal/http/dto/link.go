package dto

import "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"

// -----------------------------------------------------------------------------
// Request DTOs

type CreateLinkRequest struct {
	Title string `json:"title" binding:"required,max=32"`
	Url   string `json:"url" binding:"required,url"`
}

func NewCreateLinkRequest() *CreateLinkRequest {
	return new(CreateLinkRequest)
}

type UpdateLinkRequest struct {
	Title string `json:"title" binding:"required,max=32"`
	Url   string `json:"url" binding:"required,url"`
}

func NewUpdateLinkRequest() *UpdateLinkRequest {
	return new(UpdateLinkRequest)
}

// -----------------------------------------------------------------------------
// Response DTOs

type LinkDetails struct {
	ID    uint64     `json:"id"`
	Title string     `json:"title"`
	Url   string     `json:"url"`
	User  *UserBrief `json:"user"`
}

func ToLinkDetails(l *entity.Link) *LinkDetails {
	return &LinkDetails{
		ID:    l.ID,
		Title: l.Title,
		Url:   l.Url,
		User:  ToUserBrief(l.User),
	}
}

type LinksList []*LinkDetails

func ToLinksList(links []*entity.Link) LinksList {
	linksList := make(LinksList, len(links))
	for _, link := range links {
		linksList = append(linksList, ToLinkDetails(link))
	}
	return linksList
}
