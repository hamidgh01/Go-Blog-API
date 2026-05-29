package server

import (
	"fmt"

	"Go-Blog-API/config"
	"Go-Blog-API/internal/http/deps_container"
	"Go-Blog-API/internal/http/router"
	"Go-Blog-API/internal/http/validations"

	"github.com/gin-gonic/gin"
)

type Server struct {
	dependencyContainer *deps_container.Container
	engine              *gin.Engine
	cfg                 *config.Config
}

func InitAndRun(cfg *config.Config, deps *deps_container.Container) error {
	server := &Server{
		engine:              gin.Default(),
		cfg:                 cfg,
		dependencyContainer: deps,
	}

	// register custom validators
	if err := validations.RegisterCustomValidators(); err != nil {
		return fmt.Errorf("failed to register custom validators. origin: %w", err)
	}

	// register routes
	v1 := server.engine.Group("/v1")
	router := router.NewRouter(v1, server.dependencyContainer)
	router.RegisterRoutes()

	// run server
	address := fmt.Sprintf("%s:%d", server.cfg.Server.Host, server.cfg.Server.Port)
	return server.engine.Run(address)
}
