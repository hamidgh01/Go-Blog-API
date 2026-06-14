package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/hamidgh01/Go-Blog-API/config"
	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
	dbErrors "github.com/hamidgh01/Go-Blog-API/internal/infra/database/errors"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/redis"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/security/hashing"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/security/jwt"
)

type AuthService struct {
	userRepo      repository.UserRepository
	pswHasher     *hashing.PasswordHasher
	jwtMgr        *jwt.JWTManager
	tokenRevoker  *redis.TokenRevoker
	userInfoCache *redis.UserInfoCache
	serverConf    *config.ServerConf
}

func NewAuthService(
	u repository.UserRepository,
	p *hashing.PasswordHasher,
	j *jwt.JWTManager,
	r *redis.TokenRevoker,
	uic *redis.UserInfoCache,
	c *config.ServerConf,
) *AuthService {
	return &AuthService{userRepo: u, pswHasher: p, jwtMgr: j, tokenRevoker: r, userInfoCache: uic, serverConf: c}
}

func (a *AuthService) Register(
	ctx context.Context, data *dto.CreateUserRequest, w http.ResponseWriter,
) (*dto.LoginSuccessfulData, *service_errors.ServiceError) {
	// 1. hash entered password
	hashedPassword, err := a.pswHasher.Hash(data.Password)
	if err != nil {
		fmt.Println("failed to hash password. reason:", err) // log.Error
		return nil, service_errors.InternalServerError
	}

	// 2. created user (save hashed-password in database)
	user := &entity.User{Username: data.Username, Email: data.Email, Password: hashedPassword}
	createdUser, err := a.userRepo.Create(ctx, user)
	if err != nil {
		return nil, service_errors.MapDBErrToServiceErr(err, "create user for register")
	}

	// 3. save user info in cache
	// NOTE: when a user first created (here by this service) --> superuser=false, enabled=true
	redisErr := a.userInfoCache.SetAllInfo(ctx, createdUser.ID, createdUser.Username, false, true)
	if redisErr != nil {
		fmt.Println(redisErr) // log.Error
	}

	// 4. generate token pair (access and refresh tokens + expirations)
	tokenPair, err := a.jwtMgr.GenerateTokenPair(createdUser.ID)
	if err != nil {
		fmt.Printf("failed to generate JWT pair. reason: %s \n", err) // log.Error
		return nil, service_errors.InternalServerError
	}

	// 5. prepare needed data and add refresh_token cookie
	userData := dto.ToUserBrief(createdUser)
	tokenData := &dto.Token{AccessToken: tokenPair.AccessToken, ExpirationTS: tokenPair.AccessExpTime.Unix()}
	a.addRefreshTokenCookie(w, tokenPair.RefreshToken, int(time.Until(tokenPair.RefreshExpTime).Seconds()))

	return dto.ToLoginSuccessfulData(userData, tokenData), nil
}

func (a *AuthService) Login(
	ctx context.Context, requestData *dto.LoginRequest, w http.ResponseWriter,
) (*dto.LoginSuccessfulData, *service_errors.ServiceError) {
	// 1. get needed user data from db
	user, err := a.userRepo.GetUserForLoginVerification(ctx, requestData.Identifier)
	if err != nil {
		if errors.As(err, &dbErrors.UnexpectedDBError{}) {
			fmt.Printf("failed to get user for login verification. reason: %s \n", err.Error()) // log.Error()
			return nil, service_errors.InternalServerError
		}

		return nil, service_errors.InvalidCredentials
	}

	// 2. verify entered password
	if err := a.pswHasher.Verify(user.Password, requestData.Password); err != nil {
		return nil, service_errors.InvalidCredentials
	}

	// 3. generate token pair (access and refresh tokens + expirations)
	tokenPair, err := a.jwtMgr.GenerateTokenPair(user.ID)
	if err != nil {
		fmt.Printf("failed to generate JWT pair. reason: %s \n", err) // log.Error
		return nil, service_errors.InternalServerError
	}

	// 4. prepare needed data and add refresh_token cookie
	userData := dto.ToUserBrief(user)
	tokenData := &dto.Token{AccessToken: tokenPair.AccessToken, ExpirationTS: tokenPair.AccessExpTime.Unix()}
	a.addRefreshTokenCookie(w, tokenPair.RefreshToken, int(time.Until(tokenPair.RefreshExpTime).Seconds()))

	return dto.ToLoginSuccessfulData(userData, tokenData), nil
}

func (a *AuthService) RenewTokens(
	ctx context.Context, refreshToken string, w http.ResponseWriter,
) (*dto.Token, *service_errors.ServiceError) {
	// 1. parse refresh token and get claims
	claims, err := a.jwtMgr.ParseRefreshToken(refreshToken)
	if err != nil {
		return nil, service_errors.MapJwtErrToServiceErr(err)
	}

	// 2. check token is blacklisted or not
	isBlacklisted, err := a.tokenRevoker.IsBlacklisted(ctx, claims.GetJTI())
	if err != nil {
		fmt.Println(err.Error()) // log.Error()
		return nil, service_errors.InternalServerError
	} else if isBlacklisted {
		return nil, service_errors.TokenBlacklisted
	}

	// 3. generate token pair (access and refresh tokens + expirations)
	tokenPair, err := a.jwtMgr.GenerateTokenPair(claims.GetUserID())
	if err != nil {
		fmt.Printf("failed to generate JWT pair. reason: %s \n", err) // log.Error
		return nil, service_errors.InternalServerError
	}

	// 4. blacklist old refresh token
	err = a.tokenRevoker.Blacklist(ctx, claims.GetJTI(), claims.GetUserID(), claims.GetExpiresAt())
	if err != nil {
		fmt.Println(err.Error()) // log.Error()
	}

	// 5. add refresh_token cookie and send access token
	a.addRefreshTokenCookie(w, tokenPair.RefreshToken, int(time.Until(tokenPair.RefreshExpTime).Seconds()))
	return &dto.Token{AccessToken: tokenPair.AccessToken, ExpirationTS: tokenPair.AccessExpTime.Unix()}, nil
}

func (a *AuthService) Logout(
	ctx context.Context, refreshToken string, w http.ResponseWriter,
) *service_errors.ServiceError {
	// 1. parse refresh token and get claims
	claims, err := a.jwtMgr.ParseRefreshToken(refreshToken)
	if err != nil {
		return service_errors.MapJwtErrToServiceErr(err)
	}

	// don't need to check token is blacklisted before or not

	// 2. blacklist refresh token
	err = a.tokenRevoker.Blacklist(ctx, claims.GetJTI(), claims.GetUserID(), claims.GetExpiresAt())
	if err != nil {
		fmt.Println(err.Error()) // log.Error()
	}

	// 3. delete refresh_token cookie
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
		Domain:   a.serverConf.Host,
		MaxAge:   expirationMaxAge,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
}
