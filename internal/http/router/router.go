package router

import (
	"net/http"
	"strings"

	"github.com/hamidgh01/Go-Blog-API/internal/http/deps_container"

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

	auth := r.router.Group("")
	{
		auth.POST("/register", r.dependencies.AuthHandler.Register)
		auth.POST("/login", r.dependencies.AuthHandler.Login)
		auth.GET("/logout", r.dependencies.AuthMiddleware.Authenticate(), r.dependencies.AuthHandler.Logout)
		auth.GET("/renew-tokens", r.dependencies.AuthMiddleware.Authenticate(), r.dependencies.AuthHandler.RenewTokens)
	}

	users := r.router.Group("/users")
	users.Use(r.dependencies.AuthMiddleware.Authenticate())
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
		users.GET("/exists/username", r.dependencies.UserHandler.CheckUsernameExists)
		users.GET("/exists/email", r.dependencies.UserHandler.CheckEmailExists)

		// manipulate outer sources (related to User):
		users.POST("/:id/follow", r.dependencies.FollowHandler.Follow)
		users.DELETE("/:id/unfollow", r.dependencies.FollowHandler.Unfollow)
		users.DELETE("/:id/remove-follower", r.dependencies.FollowHandler.RemoveFollower)

		// outer sources (related to User):
		users.GET("/:id/posts", r.dependencies.UserHandler.GetPosts)
		users.GET("/:id/owned-lists", r.dependencies.UserHandler.GetOwnedLists)
		users.GET("/:id/saved-lists", r.dependencies.UserHandler.GetSavedLists)
		users.GET("/:id/all-lists", r.dependencies.UserHandler.GetAllLists)
		users.GET("/:id/comments", r.dependencies.UserHandler.GetComments)
		users.GET("/:id/likes", r.dependencies.UserHandler.GetLikes)
		users.GET("/:id/followers", r.dependencies.UserHandler.GetFollowers)
		users.GET("/:id/followings", r.dependencies.UserHandler.GetFollowings)
		users.GET("/:id/links", r.dependencies.UserHandler.GetLinks)
	}

	links := r.router.Group("/links")
	links.Use(r.dependencies.AuthMiddleware.Authenticate())
	{
		links.POST("", r.dependencies.LinkHandler.Create)
		links.PUT("/:id", r.dependencies.LinkHandler.Update)
		links.DELETE("/:id", r.dependencies.LinkHandler.Delete)
		links.GET("/:id", r.dependencies.LinkHandler.GetByID)
		// links.GET("", r.dependencies.LinkHandler.GetList)
	}

	posts := r.router.Group("/posts")
	posts.Use(r.dependencies.AuthMiddleware.Authenticate())
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

		// manipulate outer sources (related to Post):
		posts.POST("/:id/like", r.dependencies.LikeHandler.Like)
		posts.DELETE("/:id/unlike", r.dependencies.LikeHandler.Unlike)

		posts.POST("/:id/save-post", r.dependencies.SavePostHandler.Save)
		posts.DELETE("/:id/unsave-post", r.dependencies.SavePostHandler.Unsave)

		posts.POST("/:id/associate-tags", r.dependencies.PostTagsHandler.Associate)
		posts.DELETE("/:id/dissociate-tags", r.dependencies.PostTagsHandler.Dissociate)

		// outer sources (related to Post):
		posts.GET("/:id/comments", r.dependencies.PostHandler.GetComments)
		posts.GET("/:id/likes", r.dependencies.PostHandler.GetLikes)
		posts.GET("/:id/tags", r.dependencies.PostHandler.GetTags)
		posts.GET("/:id/lists", r.dependencies.PostHandler.GetLists)
	}

	tags := r.router.Group("/tags")
	tags.Use(r.dependencies.AuthMiddleware.Authenticate())
	{
		tags.POST("", r.dependencies.TagHandler.Create)
		tags.GET("/:id", r.dependencies.TagHandler.GetByID)
		tags.GET("/name=:name", r.dependencies.TagHandler.GetByName) // ToDo: check and standardize this
		// tags.GET("", r.dependencies.TagHandler.GetList)

		// outer sources (related to Tag):
		tags.GET("/:id/posts", r.dependencies.TagHandler.GetPosts)
	}

	comments := r.router.Group("/comments")
	comments.Use(r.dependencies.AuthMiddleware.Authenticate())
	{
		comments.POST("", r.dependencies.CommentHandler.Create)
		comments.PUT("/:id", r.dependencies.CommentHandler.Update)
		comments.PATCH("/:id/hide", r.dependencies.CommentHandler.Hide)
		comments.PATCH("/:id/republish", r.dependencies.CommentHandler.Republish)
		comments.PATCH("/:id/delete", r.dependencies.CommentHandler.DeleteAtUserRequest)
		comments.DELETE("/:id", r.dependencies.CommentHandler.Delete)
		comments.GET("/:id", r.dependencies.CommentHandler.GetByID)
		// comments.GET("", r.dependencies.CommentHandler.GetList)

		// outer sources (related to Comment):
		comments.GET("/:id/replies", r.dependencies.CommentHandler.GetReplies)
	}

	lists := r.router.Group("/lists")
	lists.Use(r.dependencies.AuthMiddleware.Authenticate())
	{
		lists.POST("", r.dependencies.ListHandler.Create)
		lists.PUT("/:id", r.dependencies.ListHandler.Update)
		lists.PATCH("/:id", r.dependencies.ListHandler.UpdatePrivacy)
		lists.DELETE("/:id", r.dependencies.ListHandler.Delete)
		lists.GET("/:id", r.dependencies.ListHandler.GetByID)
		// lists.GET("", r.dependencies.ListHandler.GetList)

		// outer sources (related to List):
		lists.GET("/:id/saved-posts", r.dependencies.ListHandler.GetSavedPosts)
		lists.GET("/:id/users-who-saved", r.dependencies.ListHandler.GetUsersWhoSaved)

		// manipulate outer sources (related to List):
		lists.POST("/:id/save-list", r.dependencies.SaveListHandler.Save)
		lists.DELETE("/:id/unsave-list", r.dependencies.SaveListHandler.Unsave)
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
