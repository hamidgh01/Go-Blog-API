package dto

import (
	"time"

	"Go-Blog-API/internal/domain/entity"
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
	AcceptTerms bool   `json:"accept_terms" binding:"required,eq=true"`
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
	Enabled bool `json:"enabled" binding:"required"`
}

func NewUpdateEnabledRequest() *UpdateEnabledRequest { return new(UpdateEnabledRequest) }

// -----------------------------------------------------------------------------
// Response DTOs

type UserOut struct {
	ID       uint64 `json:"id"`
	Username string `json:"name,omitempty"`
}

func ToUserOut(u *entity.User) *UserOut {
	return &UserOut{ID: u.ID, Username: u.Username}
}

type UserResponse struct {
	*UserOut
	Email      string    `json:"email"`
	Bio        string    `json:"bio"`
	Enabled    bool      `json:"enabled"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

func ToUserResponse(u *entity.User) *UserResponse {
	return &UserResponse{
		UserOut:    ToUserOut(u),
		Email:      u.Email,
		Bio:        u.Bio,
		Enabled:    u.Enabled,
		CreatedAt:  u.CreatedAt,
		ModifiedAt: u.ModifiedAt.Time,
	}
}

type UsersList []UserOut

func ToUserList(users []*entity.User) *UsersList {
	var usersList UsersList
	for _, u := range users {
		usersList = append(usersList, *ToUserOut(u))
	}

	return &usersList
}
