package commands

import (
	"log"

	"github.com/hamidgh01/Go-Blog-API/config"
	"github.com/hamidgh01/Go-Blog-API/internal/dependencies"
	"github.com/hamidgh01/Go-Blog-API/internal/http"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/database"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/redis"
	"github.com/hamidgh01/Go-Blog-API/pkg/logging"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "init and run the api server",
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

	// init logger
	logging.InitLogger(cfg.Logger)

	// establish database connection
	db, err := database.InitDB(&cfg.Postgres)
	if err != nil {
		log.Fatal("failed to establish database connection. reason: ", err)
	}
	defer db.Close()

	// establish redis connection
	redisClient, err := redis.InitRedis(&cfg.Redis)
	if err != nil {
		log.Println("failed to establish redis connection. reason:", err)
		return
	}
	defer redisClient.Close()

	// initialize dependency injector
	repositoryInjector, repositoryCleanup := dependencies.NewRepositoryInjector(db)
	defer repositoryCleanup()

	// init and run server
	if err := http.InitServerAndRun(cfg, repositoryInjector, redisClient); err != nil {
		log.Println("failed to init and run server. reason:", err)
		return
	}
}
