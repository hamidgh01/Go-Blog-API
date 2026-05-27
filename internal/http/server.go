package server

import (
	"fmt"

	"Go-Blog-API/config"
	"Go-Blog-API/internal/http/router"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	cfg    *config.Config
	// dependencies
}

func InitAndRun(cfg *config.Config) error {
	server := &Server{
		engine: gin.Default(),
		cfg:    cfg,
	}

	// register routes
	v1 := server.engine.Group("/v1")
	router := router.NewRouter(v1)
	router.RegisterRoutes()

	// run server
	address := fmt.Sprintf("%s:%d", server.cfg.Server.Host, server.cfg.Server.Port)
	return server.engine.Run(address)
}
