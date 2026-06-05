package deps_container

import (
	"database/sql"

	"github.com/hamidgh01/Go-Blog-API/config"
	"github.com/hamidgh01/Go-Blog-API/internal/application/services"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/handlers"
)

// Container holds all application dependencies
type Container struct {
	// Repositories
	UserRepository repository.UserRepository
	PostRepository repository.PostRepository

	// Services
	UserService *services.UserService
	PostService *services.PostService

	// Handlers
	UserHandler *handlers.UserHandler
	PostHandler *handlers.PostHandler

	// Middlewares
}

// NewContainer creates and wires all dependencies
func NewContainer(cfg *config.Config, db *sql.DB) *Container {
	return &Container{}
}
