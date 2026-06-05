package services

import (
	"context"
	"net/http"

	"github.com/hamidgh01/Go-Blog-API/config"
	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/redis"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/security/hashing"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/security/jwt"
)

type AuthService struct {
	userRepo     repository.UserRepository
	pwHasher     *hashing.PasswordHasher
	jwtMgr       *jwt.JWTManager
	tokenRevoker *redis.TokenRevoker
	cfg          *config.Config
}

func NewAuthService(
	u repository.UserRepository,
	p *hashing.PasswordHasher,
	j *jwt.JWTManager,
	r *redis.TokenRevoker,
	c *config.Config,
) *AuthService {
	return &AuthService{userRepo: u, pwHasher: p, jwtMgr: j, tokenRevoker: r, cfg: c}
}

func (a *AuthService) Register(
	ctx context.Context, data *dto.CreateUserRequest, w http.ResponseWriter,
) (*dto.LoginSuccessfulData, *service_errors.ServiceError) {

	refreshToken := "should be generated"
	a.addRefreshTokenCookie(w, refreshToken, a.cfg.Jwt.RefreshTokenExpirationDays*86400) // days*24*60*60 = days*86400 seconds
	return nil, nil
}

func (a *AuthService) Login(
	ctx context.Context, data *dto.LoginRequest, w http.ResponseWriter,
) (*dto.LoginSuccessfulData, *service_errors.ServiceError) {

	refreshToken := "should be generated"
	a.addRefreshTokenCookie(w, refreshToken, a.cfg.Jwt.RefreshTokenExpirationDays*86400)
	return nil, nil
}

func (a *AuthService) RenewTokens(
	ctx context.Context, refreshToken string, w http.ResponseWriter,
) (*dto.Token, *service_errors.ServiceError) {

	newRefreshToken := "should be generated"
	a.addRefreshTokenCookie(w, newRefreshToken, a.cfg.Jwt.RefreshTokenExpirationDays*86400)
	return nil, nil
}

func (a *AuthService) Logout(
	ctx context.Context, refreshToken string, w http.ResponseWriter,
) *service_errors.ServiceError {

	a.addRefreshTokenCookie(w, "", -1) // MaxAge = -1 --> delete cookie
	return nil
}

func (a *AuthService) addRefreshTokenCookie(
	w http.ResponseWriter, refreshToken string, expirationMaxAge int,
) {
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		Domain:   a.cfg.Server.Host,
		MaxAge:   expirationMaxAge,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
}
