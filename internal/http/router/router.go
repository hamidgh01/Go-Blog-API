package router

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.RouterGroup
	// dependencies
}

func NewRouter(r *gin.RouterGroup) *Router {
	return &Router{router: r}
}

func (r *Router) RegisterRoutes() {
	r.router.GET("/", index)
	r.router.GET("/ping", ping)

	users := r.router.Group("/users")
	{
		users.POST("")
		users.PATCH("/username/:id")
		users.PATCH("/email/:id")
		users.PATCH("/bio/:id")
		users.PATCH("/password/:id")
		// reset password
		users.PATCH("/enabled/:id")
		users.DELETE("/:id")
		// users.GET("") // list (needs filter for pagination)
		users.GET("/:id")
		users.GET("/:username")
		users.GET("/:email")
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
