package deps_container

import (
	"database/sql"

	"Go-Blog-API/config"
	"Go-Blog-API/internal/application/services"
	"Go-Blog-API/internal/domain/repository"
	"Go-Blog-API/internal/http/handlers"
)

// Container holds all application dependencies
type Container struct {
	// Repositories
	UserRepository repository.UserRepository

	// Services
	UserService *services.UserService

	// Handlers
	UserHandler *handlers.UserHandler

	// Middlewares
}

// NewContainer creates and wires all dependencies
func NewContainer(cfg *config.Config, db *sql.DB) *Container {
	return &Container{}
}
