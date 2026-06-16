package dto

import (
	"time"

	"github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
)

// -----------------------------------------------------------------------------
// Request DTOs

type CreateListRequest struct {
	Title       string `json:"title" binding:"required,max=100"`
	Description string `json:"description,omitempty" binding:"max=1000"`
	IsPrivate   *bool  `json:"is_private"`
}

func NewCreateListRequest() *CreateListRequest {
	return new(CreateListRequest)
}

type UpdateListRequest struct {
	Title       string `json:"title" binding:"required,max=100"`
	Description string `json:"description,omitempty" binding:"max=1000"`
	IsPrivate   *bool  `json:"is_private"`
}

func NewUpdateListRequest() *UpdateListRequest {
	return new(UpdateListRequest)
}

type UpdateListPrivacyRequest struct {
	IsPrivate *bool `json:"is_private" binding:"required"`
}

func NewUpdateListPrivacyRequest() *UpdateListPrivacyRequest {
	return new(UpdateListPrivacyRequest)
}

// -----------------------------------------------------------------------------
// Response DTOs

type ListBrief struct {
	ID         uint64     `json:"id"`
	Title      string     `json:"title"`
	IsPrivate  bool       `json:"is_private,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	ModifiedAt time.Time  `json:"modified_at,omitzero"`
	User       *UserBrief `json:"user"`
}

func ToListBrief(l *entity.List) *ListBrief {
	listBrief := &ListBrief{
		ID:        l.ID,
		Title:     l.Title,
		IsPrivate: l.IsPrivate,
		CreatedAt: l.CreatedAt,
		User:      ToUserBrief(l.User),
	}

	if l.ModifiedAt.Valid {
		listBrief.ModifiedAt = l.ModifiedAt.Time
	}

	return listBrief
}

type ListDetails struct {
	*ListBrief
	Description string `json:"description,omitempty"`
}

func ToListDetails(l *entity.List) *ListDetails {
	listDetails := &ListDetails{ListBrief: ToListBrief(l)}

	if l.Description.Valid {
		listDetails.Description = l.Description.String
	}

	return listDetails
}

type ListDetailsWithCountOfReferencedObjects struct {
	*ListDetails
	RefObjCounts  map[entity.CountKey]int `json:"referenced_objects_count"`
	SavedByViewer bool                    `json:"saved_by_viewer"`
}

func ToListDetailsWithCountOfReferencedObjects(
	l *entity.List, refObjCounts map[entity.CountKey]int, savedByViewer bool,
) *ListDetailsWithCountOfReferencedObjects {
	return &ListDetailsWithCountOfReferencedObjects{
		ListDetails:   ToListDetails(l),
		RefObjCounts:  refObjCounts,
		SavedByViewer: savedByViewer,
	}
}

type ListsList []*ListBrief

func ToListsList(lists []*entity.List) ListsList {
	listsList := make(ListsList, 0, len(lists))
	for _, list := range lists {
		listsList = append(listsList, ToListBrief(list))
	}
	return listsList
}
