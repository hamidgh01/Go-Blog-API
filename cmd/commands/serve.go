package commands

import (
	"log"

	"Go-Blog-API/config"
	server "Go-Blog-API/internal/http"
	"Go-Blog-API/internal/http/deps_container"
	"Go-Blog-API/internal/infra/database"

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

	// initialize container with all dependencies
	dependencyContainer := deps_container.NewContainer(cfg, db)

	// init and run server
	if err := server.InitAndRun(cfg, dependencyContainer); err != nil {
		log.Fatalf("failed to init and run server. reason: %v", err)
	}
}
