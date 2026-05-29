package generics

import "Go-Blog-API/internal/http/dto"

type CreateRequestTypes interface {
	dto.CreateUserRequest
}

type UpdateRequestTypes interface {
	dto.UpdateUsernameRequest | dto.UpdateEmailRequest | dto.UpdateBioRequest |
		dto.UpdatePasswordRequest | dto.UpdateEnabledRequest // | dto.ResetPasswordRequest
}

type OutputTypes interface {
	dto.UserResponse
}

type OutputListTypes interface {
	dto.UsersList
}
