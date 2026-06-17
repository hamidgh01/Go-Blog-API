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
	UserRepository     repository.UserRepository
	PostRepository     repository.PostRepository
	CommentRepository  repository.CommentRepository
	TagRepository      repository.TagRepository
	LinkRepository     repository.LinkRepository
	ListRepository     repository.ListRepository
	FollowRepository   repository.FollowRepository
	LikeRepository     repository.LikeRepository
	SavePostRepository repository.SavePostRepository
	SaveListRepository repository.SaveListRepository
	PostTagsRepository repository.PostTagsRepository

	// other infrastructure services (security, caching, etc.)
	JwtManager     *jwt.JWTManager
	PasswordHasher *hashing.PasswordHasher
	TokenRevoker   *redisInfra.TokenRevoker
	UserInfoCache  *redisInfra.UserInfoCache

	// Services
	AuthService     *services.AuthService
	UserService     *services.UserService
	PostService     *services.PostService
	CommentService  *services.CommentService
	TagService      *services.TagService
	LinkService     *services.LinkService
	ListService     *services.ListService
	FollowService   *services.FollowService
	LikeService     *services.LikeService
	SavePostService *services.SavePostService
	SaveListService *services.SaveListService
	PostTagsService *services.PostTagsService

	// Handlers
	AuthHandler     *handlers.AuthHandler
	UserHandler     *handlers.UserHandler
	PostHandler     *handlers.PostHandler
	CommentHandler  *handlers.CommentHandler
	TagHandler      *handlers.TagHandler
	LinkHandler     *handlers.LinkHandler
	ListHandler     *handlers.ListHandler
	FollowHandler   *handlers.FollowHandler
	LikeHandler     *handlers.LikeHandler
	SavePostHandler *handlers.SavePostHandler
	SaveListHandler *handlers.SaveListHandler
	PostTagsHandler *handlers.PostTagsHandler

	// Middlewares
	AuthMiddleware *middlewares.AuthenticationMiddleware
}

// NewContainer creates and wires all dependencies
func NewContainer(cfg *config.Config, db *sql.DB, redis *redis.Client) (*Container, func()) {
	// initialize repositories
	userRepo := postgres_repository.NewUserRepository(db)
	postRepo := postgres_repository.NewPostRepository(db)
	commentRepo := postgres_repository.NewCommentRepository(db)
	listRepo := postgres_repository.NewListRepository(db)
	linkRepo := postgres_repository.NewLinkRepository(db)
	tagRepo := postgres_repository.NewTagRepository(db)
	followRepo := postgres_repository.NewFollowRepository(db)
	likeRepo := postgres_repository.NewLikeRepository(db)
	savePostRepo := postgres_repository.NewSavePostRepository(db)
	saveListRepo := postgres_repository.NewSaveListRepository(db)
	postTagsRepo := postgres_repository.NewPostTagsRepository(db)

	// initialize infrastructure services
	jwtManager := jwt.NewJWTManager(&cfg.Jwt)
	passwordHasher := hashing.NewPasswordHasher()
	tokenRevoker := redisInfra.NewTokenRevoker(redis)
	userInfoCache := redisInfra.NewUserInfoCache(redis)

	// initialize services
	authService := services.NewAuthService(userRepo, passwordHasher, jwtManager, tokenRevoker, userInfoCache, &cfg.Server)
	userService := services.NewUserService(userRepo, passwordHasher, userInfoCache)
	postService := services.NewPostService(postRepo)
	commentService := services.NewCommentService(commentRepo)
	listService := services.NewListService(listRepo)
	linkService := services.NewLinkService(linkRepo)
	tagService := services.NewTagService(tagRepo)
	followService := services.NewFollowService(followRepo)
	likeService := services.NewLikeService(likeRepo)
	savePostService := services.NewSavePostService(savePostRepo, listRepo)
	saveListService := services.NewSaveListService(saveListRepo)
	postTagsService := services.NewPostTagsService(postTagsRepo, postRepo)

	// initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	postHandler := handlers.NewPostHandler(postService)
	commentHandler := handlers.NewCommentHandler(commentService)
	listHandler := handlers.NewListHandler(listService)
	linkHandler := handlers.NewLinkHandler(linkService)
	tagHandler := handlers.NewTagHandler(tagService)
	followHandler := handlers.NewFollowHandler(followService)
	likeHandler := handlers.NewLikeHandler(likeService)
	savePostHandler := handlers.NewSavePostHandler(savePostService)
	saveListHandler := handlers.NewSaveListHandler(saveListService)
	postTagsHandler := handlers.NewPostTagsHandler(postTagsService)

	// middlewares
	authMiddleware := middlewares.NewAuthenticationMiddleware(jwtManager, userInfoCache)

	return &Container{
		UserRepository:     userRepo,
		PostRepository:     postRepo,
		CommentRepository:  commentRepo,
		ListRepository:     listRepo,
		LinkRepository:     linkRepo,
		TagRepository:      tagRepo,
		FollowRepository:   followRepo,
		LikeRepository:     likeRepo,
		SavePostRepository: savePostRepo,
		SaveListRepository: saveListRepo,
		PostTagsRepository: postTagsRepo,

		JwtManager:     jwtManager,
		PasswordHasher: passwordHasher,
		TokenRevoker:   tokenRevoker,
		UserInfoCache:  userInfoCache,

		AuthService:     authService,
		UserService:     userService,
		PostService:     postService,
		CommentService:  commentService,
		ListService:     listService,
		LinkService:     linkService,
		TagService:      tagService,
		FollowService:   followService,
		LikeService:     likeService,
		SavePostService: savePostService,
		SaveListService: saveListService,
		PostTagsService: postTagsService,

		AuthHandler:     authHandler,
		UserHandler:     userHandler,
		PostHandler:     postHandler,
		CommentHandler:  commentHandler,
		ListHandler:     listHandler,
		LinkHandler:     linkHandler,
		TagHandler:      tagHandler,
		FollowHandler:   followHandler,
		LikeHandler:     likeHandler,
		SavePostHandler: savePostHandler,
		SaveListHandler: saveListHandler,
		PostTagsHandler: postTagsHandler,

		AuthMiddleware: authMiddleware,
	}, postgres_repository.CloseAllPreparedStatements
}
