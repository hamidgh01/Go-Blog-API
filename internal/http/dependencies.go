package http

import (
	"context"

	"github.com/hamidgh01/Go-Blog-API/config"
	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	"github.com/hamidgh01/Go-Blog-API/internal/application/services"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/handlers"
	"github.com/hamidgh01/Go-Blog-API/internal/http/middlewares"
	redisInfra "github.com/hamidgh01/Go-Blog-API/internal/infra/redis"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/security/hashing"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/security/jwt"
	"github.com/hamidgh01/Go-Blog-API/pkg/constants"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type accessControlMiddlewareSignature func(
	accessibility constants.EndpointAccessibility,
	getResourceOwnerIdService func(ctx context.Context, pk uint64) (uint64, *service_errors.ServiceError),
) gin.HandlerFunc

// DependencyContainer holds all application dependencies
type DependencyContainer struct {
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
	AuthMiddleware          *middlewares.AuthenticationMiddleware
	AccessControlMiddleware accessControlMiddlewareSignature
}

// NewDependencyContainer creates and wires all dependencies
func NewDependencyContainer(
	cfg *config.Config, repoInjector repository.RepositoryInjector, redis *redis.Client,
) *DependencyContainer {

	// initialize repositories
	userRepo := repoInjector.GetUserRepository()
	postRepo := repoInjector.GetPostRepository()
	commentRepo := repoInjector.GetCommentRepository()
	listRepo := repoInjector.GetListRepository()
	linkRepo := repoInjector.GetLinkRepository()
	tagRepo := repoInjector.GetTagRepository()
	followRepo := repoInjector.GetFollowRepository()
	likeRepo := repoInjector.GetLikeRepository()
	savePostRepo := repoInjector.GetSavePostRepository()
	saveListRepo := repoInjector.GetSaveListRepository()
	postTagsRepo := repoInjector.GetPostTagsRepository()

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

	return &DependencyContainer{
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

		AuthMiddleware:          authMiddleware,
		AccessControlMiddleware: middlewares.AccessControlMiddleware,
	}
}
