package deps_container

import (
	"database/sql"

	"github.com/hamidgh01/Go-Blog-API/config"
	"github.com/hamidgh01/Go-Blog-API/internal/application/services"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/handlers"
	"github.com/hamidgh01/Go-Blog-API/internal/http/middlewares"
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
	JwtManager     *jwt.JWTManager
	PasswordHasher *hashing.PasswordHasher
	TokenRevoker   *redisInfra.TokenRevoker
	UserInfoCache  *redisInfra.UserInfoCache

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
	AuthMiddleware *middlewares.AuthenticationMiddleware
}

// NewContainer creates and wires all dependencies
func NewContainer(cfg *config.Config, db *sql.DB, redis *redis.Client) (*Container, func()) {
	// initialize repositories
	userRepo := postgres_repository.NewUserRepository(db)
	postRepo := postgres_repository.NewPostRepository(db)

	// initialize infrastructure services
	jwtManager := jwt.NewJWTManager(&cfg.Jwt)
	passwordHasher := hashing.NewPasswordHasher()
	tokenRevoker := redisInfra.NewTokenRevoker(redis)
	userInfoCache := redisInfra.NewUserInfoCache(redis)

	// initialize services
	authService := services.NewAuthService(userRepo, passwordHasher, jwtManager, tokenRevoker, userInfoCache, &cfg.Server)
	userService := services.NewUserService(userRepo, passwordHasher, userInfoCache)
	postService := services.NewPostService(postRepo)

	// initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	postHandler := handlers.NewPostHandler(postService)

	// middlewares
	authMiddleware := middlewares.NewAuthenticationMiddleware(jwtManager, userInfoCache)

	return &Container{
		UserRepository: userRepo,
		PostRepository: postRepo,

		JwtManager:     jwtManager,
		PasswordHasher: passwordHasher,
		TokenRevoker:   tokenRevoker,
		UserInfoCache:  userInfoCache,

		AuthService: authService,
		UserService: userService,
		PostService: postService,

		AuthHandler: authHandler,
		UserHandler: userHandler,
		PostHandler: postHandler,

		AuthMiddleware: authMiddleware,
	}, postgres_repository.CloseAllPreparedStatements
}
