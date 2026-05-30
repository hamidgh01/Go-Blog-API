package generics

import "Go-Blog-API/internal/http/dto"

type CreateRequestTypes interface {
	dto.CreateUserRequest | dto.CreatePostRequest
}

type UpdateRequestTypes interface {
	dto.UpdateUsernameRequest | dto.UpdateEmailRequest | dto.UpdateBioRequest |
		dto.UpdatePasswordRequest | dto.UpdateEnabledRequest | // dto.ResetPasswordRequest |
		dto.UpdatePostRequest | dto.UpdatePostStatusRequest | dto.UpdatePostPrivacyRequest
}

type OutputTypes interface {
	dto.UserResponse | dto.PostDetailsResponse
}

type OutputListTypes interface {
	dto.UsersList | dto.PostsListResponse
}
