package http

import (
	"fmt"

	"github.com/hamidgh01/Go-Blog-API/config"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/validations"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
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

	// register routes
	v1 := server.engine.Group("/v1")
	router := NewRouter(v1, server.dependencyContainer)
	router.RegisterRoutes()

	// run server
	address := fmt.Sprintf("%s:%d", server.cfg.Server.Host, server.cfg.Server.Port)
	return server.engine.Run(address)
}
