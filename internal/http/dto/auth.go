package dto

// auth operations:

type LoginRequest struct {
	Identifier string `identifier:"email" binding:"required,email|username_pattern"`
	Password   string `json:"password" binding:"required,min=8,max=64,strong_password"`
}

func NewLoginRequest() *LoginRequest { return new(LoginRequest) }

type Token struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsI..."`
	ExpirationTS int64  `json:"expiration_timestamp"`
}

type LoginSuccessfulData struct {
	User  *UserBrief `json:"user"`
	Token *Token     `json:"token"`
}

func ToLoginSuccessfulData(user *UserBrief, token *Token) *LoginSuccessfulData {
	return &LoginSuccessfulData{User: user, Token: token}
}
