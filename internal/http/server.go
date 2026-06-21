package http

import (
	"fmt"

	"github.com/hamidgh01/Go-Blog-API/config"
	"github.com/hamidgh01/Go-Blog-API/docs"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/validations"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	dependencyContainer *DependencyContainer // container with all needed dependencies
	engine              *gin.Engine
	cfg                 *config.Config
}

func InitServerAndRun(
	cfg *config.Config, repoInjector repository.RepositoryInjector, redis *redis.Client,
) error {
	server := &Server{
		engine:              gin.Default(),
		dependencyContainer: NewDependencyContainer(cfg, repoInjector, redis),
		cfg:                 cfg,
	}

	// register custom validators
	if err := validations.RegisterCustomValidators(); err != nil {
		return fmt.Errorf("failed to register custom validators. origin: %w", err)
	}

	baseRoute := server.engine.Group("/api")

	// register routes
	router := NewRouter(baseRoute, server.dependencyContainer)
	router.RegisterRoutes()

	// register swagger
	RegisterSwagger(baseRoute, cfg)

	// run server
	address := fmt.Sprintf("%s:%d", server.cfg.Server.Host, server.cfg.Server.Port)
	return server.engine.Run(address)
}

func RegisterSwagger(r *gin.RouterGroup, cfg *config.Config) {
	docs.SwaggerInfo.BasePath = r.BasePath()
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
