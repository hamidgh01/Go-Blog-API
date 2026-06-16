package dto

import (
	"time"

	"github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
)

// -----------------------------------------------------------------------------
// Request DTOs

// NOTE: this struct isn't used directly; it's embedded into others.
type setPasswordOperation struct {
	Password        string `json:"password" binding:"required,min=8,max=64,strong_password"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type CreateUserRequest struct {
	Username    string `json:"username" binding:"required,min=3,max=64,username_pattern"`
	Email       string `json:"email" binding:"required,email"`
	AcceptTerms *bool  `json:"accept_terms" binding:"required,eq=true"`
	setPasswordOperation
}

func NewCreateUserRequest() *CreateUserRequest { return new(CreateUserRequest) }

type UpdateUsernameRequest struct {
	Username string `json:"username" binding:"required,min=3,max=64,username_pattern"`
}

func NewUpdateUsernameRequest() *UpdateUsernameRequest { return new(UpdateUsernameRequest) }

type UpdateEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func NewUpdateEmailRequest() *UpdateEmailRequest { return new(UpdateEmailRequest) }

type UpdateBioRequest struct {
	Bio string `json:"bio" binding:"max=500"`
}

func NewUpdateBioRequest() *UpdateBioRequest { return new(UpdateBioRequest) }

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required,min=8,max=60,nefield=Password"`
	setPasswordOperation
}

func NewUpdatePasswordRequest() *UpdatePasswordRequest { return new(UpdatePasswordRequest) }

type ResetPasswordRequest struct {
	setPasswordOperation
}

func NewResetPasswordRequest() *ResetPasswordRequest { return new(ResetPasswordRequest) }

type UpdateEnabledRequest struct {
	Enabled *bool `json:"enabled" binding:"required"`
}

func NewUpdateEnabledRequest() *UpdateEnabledRequest { return new(UpdateEnabledRequest) }

// -----------------------------------------------------------------------------
// Response DTOs

type UserBrief struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

func ToUserBrief(u *entity.User) *UserBrief {
	return &UserBrief{ID: u.ID, Username: u.Username}
}

type UserDetails struct {
	*UserBrief
	Email      string    `json:"email"`
	Bio        string    `json:"bio,omitempty"`
	Enabled    bool      `json:"enabled"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at,omitzero"`
}

func ToUserDetails(u *entity.User) *UserDetails {
	userDetails := &UserDetails{
		UserBrief: ToUserBrief(u),
		Email:     u.Email,
		Enabled:   u.Enabled,
		CreatedAt: u.CreatedAt,
	}

	if u.Bio.Valid {
		userDetails.Bio = u.Bio.String
	}

	if u.ModifiedAt.Valid {
		userDetails.ModifiedAt = u.ModifiedAt.Time
	}

	return userDetails
}

type UserDetailsWithCountOfReferencedObjects struct {
	*UserDetails
	RefObjCounts     map[entity.CountKey]int `json:"referenced_objects_count"`
	FollowedByViewer bool                    `json:"followed_by_viewer"`
	// (later)
	// [option: recent] all (5) pined & 5 recent posts (5+5=10)
	// [option: popular] 10 popular posts (most likes)
}

func ToUserDetailsWithCountOfReferencedObjects(
	u *entity.User, refObjCounts map[entity.CountKey]int, followedByViewer bool,
) *UserDetailsWithCountOfReferencedObjects {
	return &UserDetailsWithCountOfReferencedObjects{
		UserDetails:      ToUserDetails(u),
		RefObjCounts:     refObjCounts,
		FollowedByViewer: followedByViewer,
	}
}

type UsersList []*UserBrief

func ToUsersList(users []*entity.User) UsersList {
	usersList := make(UsersList, 0, len(users))
	for _, user := range users {
		usersList = append(usersList, ToUserBrief(user))
	}
	return usersList
}
