package main

import (
	"fmt"
	"log"
	"net/http"

	"Go-Blog-API/config"
	"Go-Blog-API/internal/infra/database"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(rootCmd.Execute())
}

func serve() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to init configurations. reason:", err)
	}

	db, err := database.InitDB(&cfg.Postgres)
	if err != nil {
		log.Fatal("failed to establish database connection. origin: ", err)
	}
	defer db.Close()

	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world!"})
	})

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	app.Run(address)
}
