package handlers

import (
	"github.com/hamidgh01/Go-Blog-API/internal/application/services"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
	_ "github.com/hamidgh01/Go-Blog-API/internal/http/generics"
	_ "github.com/hamidgh01/Go-Blog-API/internal/http/helpers"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service *services.PostService
}

func NewPostHandler(s *services.PostService) *PostHandler {
	return &PostHandler{service: s}
}

// Create Post godoc
// @Summary 	Create Post
// @Description Create a new post
// @Tags 		posts
// @Accept  	json
// @Produce 	json
// @Param 		request body dto.CreatePostRequest true "Data to create new post"
// @Success		201 {object} helpers.standardResponse{data=dto.PostDetails,details=nil} "Success"
// @Failure		400 "validation error"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		500 "internal server error"
// @Router 		/v1/posts [POST]
// @Security BearerAuth
func (h *PostHandler) Create(c *gin.Context) {
	create(c, dto.NewCreatePostRequest, h.service.Create)
}

// Updata Post godoc
// @Summary 	Update Post
// @Description Update a post
// @Tags 		posts
// @Accept  	json
// @Produce 	json
// @Param		id		path int					true "Unique Post ID" minimum(1)
// @Param 		request body dto.UpdatePostRequest	true "Data to update a post"
// @Success 	202 "Success"
// @Failure		400 "validation error | invalid path parameter"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		404 "post not found"
// @Failure		500 "internal server error"
// @Router 		/v1/posts/{id} [PUT]
// @Security 	BearerAuth
func (h *PostHandler) Update(c *gin.Context) {
	update(c, dto.NewUpdatePostRequest, h.service.Update)
}

// Updata Post.IsPrivate godoc
// @Summary 	Updata Post.IsPrivate
// @Description Update post's privacy
// @Tags 		posts
// @Accept  	json
// @Produce 	json
// @Param		id		path int                      		true "Unique Post ID" minimum(1)
// @Param 		request	body dto.UpdatePostPrivacyRequest 	true "new privacy value to update Post.IsPrivate"
// @Success 	202 "Success"
// @Failure		400 "validation error | invalid path parameter"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		404 "post not found"
// @Failure		500 "internal server error"
// @Router 		/v1/posts/{id}/privacy [PATCH]
// @Security 	BearerAuth
func (h *PostHandler) UpdatePrivacy(c *gin.Context) {
	update(c, dto.NewUpdatePostPrivacyRequest, h.service.UpdatePrivacy)
}

// Publish Post godoc
// @Summary 	Publish Post
// @Description Publish a draft post (change status to "published")
// @Description (**NOTE :** only draft posts (status="draft") can be published)
// @Tags 		posts
// @Produce 	json
// @Param		id 	path int true "Unique Post ID" minimum(1)
// @Success 	202 "Success"
// @Failure		400 "invalid path parameter"
// @Failure		404 "post not found"
// @Failure		500 "internal server error"
// @Router 		/v1/posts/{id}/publish [PATCH]
// @Security 	BearerAuth
func (h *PostHandler) Publish(c *gin.Context) {
	requestDTO := &dto.UpdatePostStatusRequest{Status: "published"}
	updateStatusEnum(c, requestDTO, h.service.PublishDraftPost)
}

// Reject Post godoc
// @Summary 	Reject Post
// @Description Reject a post (change status to "rejected") (**only admin access**)
// @Tags 		posts
// @Produce 	json
// @Param		id 	path int true "Unique Post ID" minimum(1)
// @Success 	202 "Success"
// @Failure		400 "invalid path parameter"
// @Failure		404 "post not found"
// @Failure		500 "internal server error"
// @Router 		/v1/posts/{id}/reject [PATCH]
// @Security 	BearerAuth
func (h *PostHandler) Reject(c *gin.Context) {
	requestDTO := &dto.UpdatePostStatusRequest{Status: "rejected"}
	updateStatusEnum(c, requestDTO, h.service.UpdateStatus)
}

// Republish Post godoc
// @Summary 	Republish Post
// @Description Republish a post (change status to "published") (**only admin access**)
// @Tags 		posts
// @Produce 	json
// @Param		id 	path int true "Unique Post ID" minimum(1)
// @Success 	202 "Success"
// @Failure		400 "invalid path parameter"
// @Failure		404 "post not found"
// @Failure		500 "internal server error"
// @Router 		/v1/posts/{id}/republish [PATCH]
// @Security 	BearerAuth
func (h *PostHandler) Republish(c *gin.Context) {
	requestDTO := &dto.UpdatePostStatusRequest{Status: "published"}
	updateStatusEnum(c, requestDTO, h.service.UpdateStatus)
}

// DeletePostAtUserRequest godoc
// @Summary 	Delete Post At User Request
// @Description Delete a post at user request (change status to "deleted")
// @Tags 		posts
// @Produce 	json
// @Param		id 	path int true "Unique Post ID" minimum(1)
// @Success 	202 "Success"
// @Failure		400 "invalid path parameter"
// @Failure		404 "post not found"
// @Failure		500 "internal server error"
// @Router 		/v1/posts/{id}/delete [PATCH]
// @Security 	BearerAuth
func (h *PostHandler) DeleteAtUserRequest(c *gin.Context) {
	requestDTO := &dto.UpdatePostStatusRequest{Status: "deleted-by-author"}
	updateStatusEnum(c, requestDTO, h.service.UpdateStatus)
}

// Delete Post godoc
// @Summary 	Delete Post
// @Description Delete a post (**only admin access**)
// @Tags 		posts
// @Produce 	json
// @Param		id 	path int true "Unique Post ID" minimum(1)
// @Success 	202 "Success"
// @Failure		400 "invalid path parameter"
// @Failure		404 "post not found"
// @Failure		500 "internal server error"
// @Router 		/v1/posts/{id} [DELETE]
// @Security 	BearerAuth
func (h *PostHandler) Delete(c *gin.Context) {
	delete(c, h.service.Delete)
}

// GetPostByID godoc
// @Summary 	Get Post By ID
// @Description Get a post by unique id
// @Tags 		posts
// @Produce 	json
// @Param		id path int true "Unique Post ID" minimum(1)
// @Success 	200 {object} helpers.standardResponse{data=dto.PostDetails,details=nil} "Success"
// @Failure		400 "invalid path parameter"
// @Failure		404 "post not found"
// @Failure		500 "internal server error"
// @Router 		/v1/posts/{id} [GET]
// @Security 	BearerAuth
func (h *PostHandler) GetByID(c *gin.Context) {
	getByID(c, h.service.GetByID)
}

// -----------------------------------------------------------------------------
// other sources that has FK to `Post`

// GetPostComments godoc
// @Summary     Get Posts's Comments
// @Description	Get a list of post's comments with pagination
// @Tags		posts
// @Produce		json
// @Param		id		path	int	true 	"Unique Post ID to use as FK" 	minimum(1)
// @Param		page	query	int	false	"Page Number (default: 1)"		minimum(1)	default(1)
// @Param		size	query	int	false	"Items per Page (default: 10)"	minimum(1)	maximum(100)	default(10)
// @Success		200	{object} helpers.standardResponse{data=generics.PagedList[dto.CommentList]{items=dto.CommentList},details=nil} "Success"
// @Failure		400 "invalid path parameter | invalid query parameters (for 'size' or 'page')"
// @Failure		404 "there's no any comment for this post"
// @Failure		500 "internal server error"
// @Router		/v1/posts/{id}/comments [GET]
// @Security	BearerAuth
func (h *PostHandler) GetComments(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetComments)
}

// GetPostLikes godoc
// @Summary     Get Posts's Likes
// @Description	Get a list of post's likes (users who liked this post) with pagination
// @Tags		posts
// @Produce		json
// @Param		id		path	int	true 	"Unique Post ID to use as FK" 	minimum(1)
// @Param		page	query	int	false	"Page Number (default: 1)"		minimum(1)	default(1)
// @Param		size	query	int	false	"Items per Page (default: 10)"	minimum(1)	maximum(100)	default(10)
// @Success		200	{object} helpers.standardResponse{data=generics.PagedList[dto.UsersList]{items=dto.UsersList},details=nil} "Success"
// @Failure		400 "invalid path parameter | invalid query parameters (for 'size' or 'page')"
// @Failure		404 "there's no any like for this post"
// @Failure		500 "internal server error"
// @Router		/v1/posts/{id}/likes [GET]
// @Security	BearerAuth
func (h *PostHandler) GetLikes(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetLikes)
}

// GetPostTags godoc
// @Summary     Get Posts's Tags
// @Description	Get a list of post's tags with pagination
// @Tags		posts
// @Produce		json
// @Param		id		path	int	true 	"Unique Post ID to use as FK" 	minimum(1)
// @Param		page	query	int	false	"Page Number (default: 1)"		minimum(1)	default(1)
// @Param		size	query	int	false	"Items per Page (default: 10)"	minimum(1)	maximum(100)	default(10)
// @Success		200	{object} helpers.standardResponse{data=generics.PagedList[dto.TagsList]{items=dto.TagsList},details=nil} "Success"
// @Failure		400 "invalid path parameter | invalid query parameters (for 'size' or 'page')"
// @Failure		404 "there's no any tag for this post"
// @Failure		500 "internal server error"
// @Router		/v1/posts/{id}/tags [GET]
// @Security	BearerAuth
func (h *PostHandler) GetTags(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetTags)
}

// GetPostLists godoc
// @Summary     Get Posts's Lists
// @Description	Get a list of post's lists (lists that saved this post) with pagination
// @Tags		posts
// @Produce		json
// @Param		id		path	int	true 	"Unique Post ID to use as FK" 	minimum(1)
// @Param		page	query	int	false	"Page Number (default: 1)"		minimum(1)	default(1)
// @Param		size	query	int	false	"Items per Page (default: 10)"	minimum(1)	maximum(100)	default(10)
// @Success		200	{object} helpers.standardResponse{data=generics.PagedList[dto.ListsList]{items=dto.ListsList},details=nil} "Success"
// @Failure		400 "invalid path parameter | invalid query parameters (for 'size' or 'page')"
// @Failure		404 "there's no any list that saves this post"
// @Failure		500 "internal server error"
// @Router		/v1/posts/{id}/lists [GET]
// @Security	BearerAuth
func (h *PostHandler) GetLists(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetLists)
}
