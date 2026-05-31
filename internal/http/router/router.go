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
		users.PATCH("/:id/username", r.dependencies.UserHandler.UpdateUsername)
		users.PATCH("/:id/email", r.dependencies.UserHandler.UpdateEmail)
		users.PATCH("/:id/bio", r.dependencies.UserHandler.UpdateBio)
		users.PATCH("/:id/password", r.dependencies.UserHandler.UpdatePassword)
		// reset password
		users.PATCH("/:id/enabled", r.dependencies.UserHandler.UpdateEnabled)
		users.DELETE("/:id", r.dependencies.UserHandler.Delete)
		// users.GET("") // list (needs filter for pagination)
		users.GET("/:id", r.dependencies.UserHandler.GetByID)
		users.GET("/username=:username", r.dependencies.UserHandler.GetByUsername) // ToDo: check and standardize this
		users.GET("/email=:email", r.dependencies.UserHandler.GetByEmail)          // ToDo: check and standardize this
	}

	posts := r.router.Group("/posts")
	{
		posts.POST("", r.dependencies.PostHandler.Create)
		posts.PUT("/:id", r.dependencies.PostHandler.Update)
		posts.PATCH("/:id/privacy", r.dependencies.PostHandler.UpdatePrivacy)
		posts.PATCH("/:id/publish", r.dependencies.PostHandler.Publish)
		posts.PATCH("/:id/reject", r.dependencies.PostHandler.Reject)
		posts.PATCH("/:id/republish", r.dependencies.PostHandler.Republish)
		posts.PATCH("/:id/delete", r.dependencies.PostHandler.DeleteAtUserRequest)
		posts.DELETE("/:id", r.dependencies.PostHandler.Delete)
		posts.GET("/:id", r.dependencies.PostHandler.GetByID)
		// posts.GET("", r.dependencies.PostHandler.GetList)
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
