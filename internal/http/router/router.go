package router

import (
	"net/http"
	"strings"

	"Go-Blog-API/internal/http/deps_container"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router       *gin.RouterGroup
	dependencies *deps_container.Container
}

func NewRouter(r *gin.RouterGroup, deps *deps_container.Container) *Router {
	return &Router{router: r, dependencies: deps}
}

func (r *Router) RegisterRoutes() {
	r.router.GET("/", index)
	r.router.GET("/ping", ping)

	users := r.router.Group("/users")
	{
		users.POST("", r.dependencies.UserHandler.Create)
		users.PATCH("/username/:id", r.dependencies.UserHandler.UpdateUsername)
		users.PATCH("/email/:id", r.dependencies.UserHandler.UpdateEmail)
		users.PATCH("/bio/:id", r.dependencies.UserHandler.UpdateBio)
		users.PATCH("/password/:id", r.dependencies.UserHandler.UpdatePassword)
		// reset password
		users.PATCH("/enabled/:id", r.dependencies.UserHandler.UpdateEnabled)
		users.DELETE("/:id", r.dependencies.UserHandler.Delete)
		// users.GET("") // list (needs filter for pagination)
		users.GET("/:id", r.dependencies.UserHandler.GetByID)
		users.GET("/username=:username", r.dependencies.UserHandler.GetByUsername) // ToDo: check and standardize this
		users.GET("/email=:email", r.dependencies.UserHandler.GetByEmail)          // ToDo: check and standardize this
	}
}

func index(c *gin.Context) {
	userAgent := c.Request.Header.Get("User-Agent")
	if strings.Contains(userAgent, "curl") {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "hello from gin server."})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "hello from gin server."})
	}

}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
