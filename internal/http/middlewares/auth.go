package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	"github.com/hamidgh01/Go-Blog-API/internal/http/helpers"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/redis"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/security/jwt"

	"github.com/gin-gonic/gin"
)

type AuthenticationMiddleware struct {
	jwtMgr        *jwt.JWTManager
	userInfoCache *redis.UserInfoCache
}

func NewAuthenticationMiddleware(
	j *jwt.JWTManager, uic *redis.UserInfoCache,
) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{jwtMgr: j, userInfoCache: uic}
}

func (m *AuthenticationMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. get authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			sendUnauthorizedResponse(c, "missing Authorization header")
			return
		}

		// authorization header format: "Bearer <accessToken>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			sendUnauthorizedResponse(c, "invalid Authorization header format")
			return
		}

		// 2. parse access token and get userID form claims
		claims, err := m.jwtMgr.ParseAccessToken(parts[1])
		if err != nil {
			sendUnauthorizedResponse(c, service_errors.MapJwtErrToServiceErr(err).Error())
			return
		}

		userID := claims.GetUserID()
		if userID == 0 {
			sendUnauthorizedResponse(c, "invalid token claims (missing user_id key in claims)")
			return
		}

		// 3. fetch essential user info from cache
		userInfo, redisErr := m.userInfoCache.GetAllInfo(c, claims.GetUserID())
		if redisErr != nil {
			fmt.Println(redisErr) // log.Error

			c.AbortWithStatusJSON(
				service_errors.InternalServerError.Code(),
				helpers.GenerateErrorResponse(
					service_errors.InternalServerError.Message(), nil,
				),
			)
			return
		}

		// 4. check user is suspended or not
		if userInfo.Enabled == "f" {
			c.AbortWithStatusJSON(
				service_errors.PermissionDenied.Code(),
				helpers.GenerateErrorResponse(
					service_errors.PermissionDenied.Message(),
					gin.H{"reason": "user is suspended"},
				),
			)
			return
		}

		// 5. set keys into gin context
		c.Set("currentUserID", userID)
		c.Set("currentUserUsername", userInfo.Username)
		c.Set("currentUserEnabled", true)

		if userInfo.IsSuperuser == "f" {
			c.Set("currentUserIsSuperuser", false)
		} else {
			c.Set("currentUserIsSuperuser", true)
		}

		c.Next()
	}
}

func sendUnauthorizedResponse(c *gin.Context, message string) {
	c.AbortWithStatusJSON(
		http.StatusUnauthorized,
		helpers.GenerateErrorResponse("authentication failed", gin.H{"reason": message}),
	)
}
