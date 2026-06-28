package http

import (
	"net/http"
	"strings"

	"github.com/hamidgh01/Go-Blog-API/pkg/constants"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.RouterGroup
	deps   *DependencyContainer
}

func NewRouter(r *gin.RouterGroup, depsContainer *DependencyContainer) *Router {
	return &Router{router: r, deps: depsContainer}
}

func (r *Router) RegisterRoutes() {
	v1 := r.router.Group("/v1")

	v1.GET("/", index)
	v1.GET("/ping", ping)

	auth := v1.Group("")
	{
		auth.POST("/register", r.deps.AuthHandler.Register)
		auth.POST("/login", r.deps.AuthHandler.Login)
		auth.GET("/logout", r.deps.AuthMiddleware.Authenticate(), r.deps.AuthHandler.Logout)
		auth.GET("/renew-tokens", r.deps.AuthMiddleware.Authenticate(), r.deps.AuthHandler.RenewTokens)
	}

	users := v1.Group("/users")
	users.Use(r.deps.AuthMiddleware.Authenticate())
	{
		users.POST("", r.deps.UserHandler.Create)
		users.PATCH("/:id/username", r.deps.UserHandler.UpdateUsername)
		users.PATCH("/:id/email", r.deps.UserHandler.UpdateEmail)
		users.PATCH("/:id/bio", r.deps.UserHandler.UpdateBio)
		users.PATCH("/:id/password", r.deps.UserHandler.UpdatePassword)
		// reset password
		users.PATCH("/:id/enabled", r.deps.UserHandler.UpdateEnabled)
		users.DELETE("/:id", r.deps.UserHandler.Delete)
		// users.GET("") // list (needs filter for pagination)
		users.GET("/:id", r.deps.UserHandler.GetByID)
		users.GET("/username/:username", r.deps.UserHandler.GetByUsername)
		users.GET("/email/:email", r.deps.UserHandler.GetByEmail)
		users.GET("/exists/username", r.deps.UserHandler.CheckUsernameExists)
		users.GET("/exists/email", r.deps.UserHandler.CheckEmailExists)

		// manipulate outer sources (related to User):
		users.POST("/:id/follow", r.deps.FollowHandler.Follow)
		users.DELETE("/:id/unfollow", r.deps.FollowHandler.Unfollow)
		users.DELETE("/:id/remove-follower", r.deps.FollowHandler.RemoveFollower)

		// outer sources (related to User):
		users.GET("/:id/posts", r.deps.UserHandler.GetPosts)
		users.GET("/:id/owned-lists", r.deps.UserHandler.GetOwnedLists)
		users.GET("/:id/saved-lists", r.deps.UserHandler.GetSavedLists)
		users.GET("/:id/all-lists", r.deps.UserHandler.GetAllLists)
		users.GET("/:id/comments", r.deps.UserHandler.GetComments)
		users.GET("/:id/likes", r.deps.UserHandler.GetLikes)
		users.GET("/:id/followers", r.deps.UserHandler.GetFollowers)
		users.GET("/:id/followings", r.deps.UserHandler.GetFollowings)
		users.GET("/:id/links", r.deps.UserHandler.GetLinks)
	}

	links := v1.Group("/links")
	links.Use(r.deps.AuthMiddleware.Authenticate())
	{
		links.POST("", r.deps.LinkHandler.Create)
		links.PUT(
			"/:id",
			r.deps.AccessControlMiddleware(constants.ADMIN_AND_OWNER, r.deps.LinkService.GetOwnerID),
			r.deps.LinkHandler.Update,
		)
		links.DELETE(
			"/:id",
			r.deps.AccessControlMiddleware(constants.ADMIN_AND_OWNER, r.deps.LinkService.GetOwnerID),
			r.deps.LinkHandler.Delete,
		)
		links.GET("/:id", r.deps.LinkHandler.GetByID)
	}

	posts := v1.Group("/posts")
	posts.Use(r.deps.AuthMiddleware.Authenticate())
	{
		posts.POST("", r.deps.PostHandler.Create)
		posts.PUT(
			"/:id",
			r.deps.AccessControlMiddleware(constants.ADMIN_AND_OWNER, r.deps.PostService.GetOwnerID),
			r.deps.PostHandler.Update,
		)
		posts.PATCH(
			"/:id/privacy",
			r.deps.AccessControlMiddleware(constants.ADMIN_AND_OWNER, r.deps.PostService.GetOwnerID),
			r.deps.PostHandler.UpdatePrivacy,
		)
		posts.PATCH(
			"/:id/publish",
			r.deps.AccessControlMiddleware(constants.ADMIN_AND_OWNER, r.deps.PostService.GetOwnerID),
			r.deps.PostHandler.Publish,
		)
		posts.PATCH(
			"/:id/reject",
			r.deps.AccessControlMiddleware(constants.ADMIN_ONLY, r.deps.PostService.GetOwnerID),
			r.deps.PostHandler.Reject,
		)
		posts.PATCH(
			"/:id/republish",
			r.deps.AccessControlMiddleware(constants.ADMIN_ONLY, r.deps.PostService.GetOwnerID),
			r.deps.PostHandler.Republish,
		)
		posts.PATCH(
			"/:id/delete",
			r.deps.AccessControlMiddleware(constants.ADMIN_AND_OWNER, r.deps.PostService.GetOwnerID),
			r.deps.PostHandler.DeleteAtUserRequest,
		)
		posts.DELETE(
			"/:id",
			r.deps.AccessControlMiddleware(constants.ADMIN_ONLY, r.deps.PostService.GetOwnerID),
			r.deps.PostHandler.Delete,
		)
		posts.GET("/:id", r.deps.PostHandler.GetByID)
		// posts.GET("", r.deps.PostHandler.GetList)

		// manipulate outer sources (related to Post):
		posts.POST("/:id/like", r.deps.LikeHandler.Like)
		posts.DELETE("/:id/unlike", r.deps.LikeHandler.Unlike)

		posts.POST("/:id/save-post", r.deps.SavePostHandler.Save)
		posts.DELETE("/:id/unsave-post", r.deps.SavePostHandler.Unsave)

		posts.POST("/:id/associate-tags", r.deps.PostTagsHandler.Associate)
		posts.DELETE("/:id/dissociate-tags", r.deps.PostTagsHandler.Dissociate)

		// outer sources (related to Post):
		posts.GET("/:id/comments", r.deps.PostHandler.GetComments)
		posts.GET("/:id/likes", r.deps.PostHandler.GetLikes)
		posts.GET("/:id/tags", r.deps.PostHandler.GetTags)
		posts.GET("/:id/lists", r.deps.PostHandler.GetLists)
	}

	tags := v1.Group("/tags")
	tags.Use(r.deps.AuthMiddleware.Authenticate())
	{
		tags.POST("", r.deps.TagHandler.Create)
		tags.GET("/:id", r.deps.TagHandler.GetByID)
		tags.GET("/name/:name", r.deps.TagHandler.GetByName)
		// tags.GET("", r.deps.TagHandler.GetList)

		// outer sources (related to Tag):
		tags.GET("/:id/posts", r.deps.TagHandler.GetPosts)
	}

	comments := v1.Group("/comments")
	comments.Use(r.deps.AuthMiddleware.Authenticate())
	{
		comments.POST("", r.deps.CommentHandler.Create)
		comments.PUT(
			"/:id",
			r.deps.AccessControlMiddleware(constants.ADMIN_AND_OWNER, r.deps.CommentService.GetOwnerID),
			r.deps.CommentHandler.Update,
		)
		comments.PATCH(
			"/:id/hide",
			r.deps.AccessControlMiddleware(constants.ADMIN_ONLY, r.deps.CommentService.GetOwnerID),
			r.deps.CommentHandler.Hide,
		)
		comments.PATCH(
			"/:id/republish",
			r.deps.AccessControlMiddleware(constants.ADMIN_ONLY, r.deps.CommentService.GetOwnerID),
			r.deps.CommentHandler.Republish,
		)
		comments.PATCH(
			"/:id/delete",
			r.deps.AccessControlMiddleware(constants.ADMIN_AND_OWNER, r.deps.CommentService.GetOwnerID),
			r.deps.CommentHandler.DeleteAtUserRequest,
		)
		comments.DELETE(
			"/:id",
			r.deps.AccessControlMiddleware(constants.ADMIN_ONLY, r.deps.CommentService.GetOwnerID),
			r.deps.CommentHandler.Delete,
		)
		comments.GET("/:id", r.deps.CommentHandler.GetByID)
		// comments.GET("", r.deps.CommentHandler.GetList)

		// outer sources (related to Comment):
		comments.GET("/:id/replies", r.deps.CommentHandler.GetReplies)
	}

	lists := v1.Group("/lists")
	lists.Use(r.deps.AuthMiddleware.Authenticate())
	{
		lists.POST("", r.deps.ListHandler.Create)
		lists.PUT(
			"/:id",
			r.deps.AccessControlMiddleware(constants.ADMIN_AND_OWNER, r.deps.ListService.GetOwnerID),
			r.deps.ListHandler.Update,
		)
		lists.PATCH(
			"/:id",
			r.deps.AccessControlMiddleware(constants.ADMIN_AND_OWNER, r.deps.ListService.GetOwnerID),
			r.deps.ListHandler.UpdatePrivacy,
		)
		lists.DELETE(
			"/:id",
			r.deps.AccessControlMiddleware(constants.ADMIN_AND_OWNER, r.deps.ListService.GetOwnerID),
			r.deps.ListHandler.Delete,
		)
		lists.GET("/:id", r.deps.ListHandler.GetByID)
		// lists.GET("", r.deps.ListHandler.GetList)

		// outer sources (related to List):
		lists.GET("/:id/saved-posts", r.deps.ListHandler.GetSavedPosts)
		lists.GET("/:id/users-who-saved", r.deps.ListHandler.GetUsersWhoSaved)

		// manipulate outer sources (related to List):
		lists.POST("/:id/save-list", r.deps.SaveListHandler.Save)
		lists.DELETE("/:id/unsave-list", r.deps.SaveListHandler.Unsave)
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
