package deps_container

import (
	"database/sql"

	"github.com/hamidgh01/Go-Blog-API/config"
	"github.com/hamidgh01/Go-Blog-API/internal/application/services"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/handlers"
	redisInfra "github.com/hamidgh01/Go-Blog-API/internal/infra/redis"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/security/hashing"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/security/jwt"

	"github.com/redis/go-redis/v9"
)

// Container holds all application dependencies
type Container struct {
	// Repositories
	UserRepository repository.UserRepository
	PostRepository repository.PostRepository

	// other infrastructure services (security, caching, etc.)
	JwtManager          *jwt.JWTManager
	PasswordHasher      *hashing.PasswordHasher
	TokenRevoker        *redisInfra.TokenRevoker
	UserDisabledChecker *redisInfra.UserDisabledChecker

	// Services
	UserService *services.UserService
	PostService *services.PostService

	// Handlers
	UserHandler *handlers.UserHandler
	PostHandler *handlers.PostHandler

	// Middlewares
}

// NewContainer creates and wires all dependencies
func NewContainer(cfg *config.Config, db *sql.DB, redis *redis.Client) *Container {
	return &Container{}
}
