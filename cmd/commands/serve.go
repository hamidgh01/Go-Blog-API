package commands

import (
	"log"

	"Go-Blog-API/config"
	server "Go-Blog-API/internal/http"
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

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to init configurations. reason:", err)
	}

	db, err := database.InitDB(&cfg.Postgres)
	if err != nil {
		log.Fatal("failed to establish database connection. reason: ", err)
	}
	defer db.Close()

	if err := server.InitAndRun(cfg); err != nil {
		log.Fatalf("failed to run server. reason: %v", err)
	}
}
