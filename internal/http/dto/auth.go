package dto

// auth operations:

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

func NewLoginRequest() *LoginRequest { return new(LoginRequest) }

type Token struct {
	AccessToken  string `json:"access_token"`
	ExpirationTS int64  `json:"expiration_timestamp"`
}

type LoginSuccessfulData struct {
	User  *UserBrief `json:"user"`
	Token *Token     `json:"token"`
}

func ToLoginSuccessfulData(user *UserBrief, token *Token) *LoginSuccessfulData {
	return &LoginSuccessfulData{User: user, Token: token}
}
