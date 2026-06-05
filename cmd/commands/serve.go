package commands

import (
	"log"

	"github.com/hamidgh01/Go-Blog-API/config"
	server "github.com/hamidgh01/Go-Blog-API/internal/http"
	"github.com/hamidgh01/Go-Blog-API/internal/http/deps_container"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/database"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/redis"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start api server",
	// Long:    "...",
	// Example: "...",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	// init configurations
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to init configurations. reason:", err)
	}

	// establish database connection
	db, err := database.InitDB(&cfg.Postgres)
	if err != nil {
		log.Fatal("failed to establish database connection. reason: ", err)
	}
	defer db.Close()

	// establish redis connection
	redisClient, err := redis.InitRedis(&cfg.Redis)
	if err != nil {
		log.Fatal("failed to establish redis connection. reason: ", err)
	}
	defer redisClient.Close()

	// initialize container with all dependencies
	dependencyContainer := deps_container.NewContainer(cfg, db, redisClient)

	// init and run server
	if err := server.InitAndRun(cfg, dependencyContainer); err != nil {
		log.Fatalf("failed to init and run server. reason: %v", err)
	}
}
