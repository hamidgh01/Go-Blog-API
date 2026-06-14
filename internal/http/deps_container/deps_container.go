package deps_container

import (
	"database/sql"

	"github.com/hamidgh01/Go-Blog-API/config"
	"github.com/hamidgh01/Go-Blog-API/internal/application/services"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/handlers"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/database/postgres_repository"
	redisInfra "github.com/hamidgh01/Go-Blog-API/internal/infra/redis"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/security/hashing"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/security/jwt"

	"github.com/redis/go-redis/v9"
)

// Container holds all application dependencies
type Container struct {
	// Repositories
	UserRepository    repository.UserRepository
	PostRepository    repository.PostRepository
	CommentRepository repository.CommentRepository
	TagRepository     repository.TagRepository
	LinkRepository    repository.LinkRepository
	ListRepository    repository.ListRepository

	// other infrastructure services (security, caching, etc.)
	JwtManager          *jwt.JWTManager
	PasswordHasher      *hashing.PasswordHasher
	TokenRevoker        *redisInfra.TokenRevoker
	UserDisabledChecker *redisInfra.UserDisabledChecker

	// Services
	AuthService    *services.AuthService
	UserService    *services.UserService
	PostService    *services.PostService
	CommentService *services.CommentService
	TagService     *services.TagService
	LinkService    *services.LinkService
	ListService    *services.ListService

	// Handlers
	AuthHandler    *handlers.AuthHandler
	UserHandler    *handlers.UserHandler
	PostHandler    *handlers.PostHandler
	CommentHandler *handlers.CommentHandler
	TagHandler     *handlers.TagHandler
	LinkHandler    *handlers.LinkHandler
	ListHandler    *handlers.ListHandler

	// Middlewares
}

// NewContainer creates and wires all dependencies
func NewContainer(cfg *config.Config, db *sql.DB, redis *redis.Client) (*Container, func()) {
	// initialize repositories
	userRepo := postgres_repository.NewUserRepository(db)

	// initialize infrastructure services
	jwtManager := jwt.NewJWTManager(&cfg.Jwt)
	passwordHasher := hashing.NewPasswordHasher()
	tokenRevoker := redisInfra.NewTokenRevoker(redis)
	UserDisabledChecker := redisInfra.NewUserDisabledChecker(redis)

	// initialize services
	authService := services.NewAuthService(userRepo, passwordHasher, jwtManager, tokenRevoker, &cfg.Server)
	userService := services.NewUserService(userRepo, passwordHasher, UserDisabledChecker)

	// initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)

	// middlewares
	// ...

	return &Container{
		UserRepository: userRepo,

		JwtManager:          jwtManager,
		PasswordHasher:      passwordHasher,
		TokenRevoker:        tokenRevoker,
		UserDisabledChecker: UserDisabledChecker,

		AuthService: authService,
		UserService: userService,

		AuthHandler: authHandler,
		UserHandler: userHandler,
	}, postgres_repository.CloseAllPreparedStatements
}
