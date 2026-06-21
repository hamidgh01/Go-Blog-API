package handlers

import (
	"github.com/hamidgh01/Go-Blog-API/internal/application/services"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
	_ "github.com/hamidgh01/Go-Blog-API/internal/http/generics"
	_ "github.com/hamidgh01/Go-Blog-API/internal/http/helpers"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	service *services.CommentService
}

func NewCommentHandler(s *services.CommentService) *CommentHandler {
	return &CommentHandler{service: s}
}

// Create Comment godoc
// @Summary 	Create Comment
// @Description Create a new comment
// @Tags 		comments
// @Accept  	json
// @Produce 	json
// @Param 		request body dto.CreateCommentRequest true "Data to create new comment"
// @Success		201 {object} helpers.standardResponse{data=dto.CommentDetails,details=nil} "Success"
// @Failure		400 "validation error"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		500 "internal server error"
// @Router 		/v1/comments [POST]
// @Security BearerAuth
func (h *CommentHandler) Create(c *gin.Context) {
	create(c, dto.NewCreateCommentRequest, h.service.Create)
}

// Updata Comment godoc
// @Summary 	Update Comment
// @Description Update a comment
// @Tags 		comments
// @Accept  	json
// @Produce 	json
// @Param		id		path int                      true "Unique Comment ID" minimum(1)
// @Param 		request body dto.UpdateCommentRequest true "Data to update a comment"
// @Success 	202 "Success"
// @Failure		400 "validation error | invalid path parameter"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		404 "comment not found"
// @Failure		500 "internal server error"
// @Router 		/v1/comments/{id} [PUT]
// @Security 	BearerAuth
func (h *CommentHandler) Update(c *gin.Context) {
	update(c, dto.NewUpdateCommentRequest, h.service.Update)
}

// Hide Comment godoc
// @Summary 	Hide Comment
// @Description Hide a comment (change status to "hidden") (**only admin access**)
// @Tags 		comments
// @Produce 	json
// @Param		id 	path int true "Unique Comment ID" minimum(1)
// @Success 	202 "Success"
// @Failure		400 "invalid path parameter"
// @Failure		404 "comment not found"
// @Failure		500 "internal server error"
// @Router 		/v1/comments/{id}/hide [PATCH]
// @Security 	BearerAuth
func (h *CommentHandler) Hide(c *gin.Context) {
	requestDTO := &dto.UpdateCommentStatusRequest{Status: "hidden-by-Admin"}
	updateStatusEnum(c, requestDTO, h.service.UpdateStatus)
}

// Republish Comment godoc
// @Summary 	Republish Comment
// @Description Republish a comment (change status to "published") (**only admin access**)
// @Tags 		comments
// @Produce 	json
// @Param		id 	path int true "Unique Comment ID" minimum(1)
// @Success 	202 "Success"
// @Failure		400 "invalid path parameter"
// @Failure		404 "comment not found"
// @Failure		500 "internal server error"
// @Router 		/v1/comments/{id}/republish [PATCH]
// @Security 	BearerAuth
func (h *CommentHandler) Republish(c *gin.Context) {
	requestDTO := &dto.UpdateCommentStatusRequest{Status: "published"}
	updateStatusEnum(c, requestDTO, h.service.UpdateStatus)
}

// DeleteCommentAtUserRequest godoc
// @Summary 	Delete Comment At User Request
// @Description Delete a comment at user request (change status to "deleted")
// @Tags 		comments
// @Produce 	json
// @Param		id 	path int true "Unique Comment ID" minimum(1)
// @Success 	202 "Success"
// @Failure		400 "invalid path parameter"
// @Failure		404 "comment not found"
// @Failure		500 "internal server error"
// @Router 		/v1/comments/{id}/delete [PATCH]
// @Security 	BearerAuth
func (h *CommentHandler) DeleteAtUserRequest(c *gin.Context) {
	requestDTO := &dto.UpdateCommentStatusRequest{Status: "deleted-by-commenter"}
	updateStatusEnum(c, requestDTO, h.service.UpdateStatus)
}

// Delete Comment godoc
// @Summary 	Delete Comment
// @Description Delete a comment (**only admin access**)
// @Tags 		comments
// @Produce 	json
// @Param		id 	path int true "Unique Comment ID" minimum(1)
// @Success 	202 "Success"
// @Failure		400 "invalid path parameter"
// @Failure		404 "comment not found"
// @Failure		500 "internal server error"
// @Router 		/v1/comments/{id} [DELETE]
// @Security 	BearerAuth
func (h *CommentHandler) Delete(c *gin.Context) {
	delete(c, h.service.Delete)
}

// GetCommentByID godoc
// @Summary 	Get Comment By ID
// @Description Get a comment by unique id
// @Tags 		comments
// @Produce 	json
// @Param		id path int true "Unique Comment ID" minimum(1)
// @Success 	200 {object} helpers.standardResponse{data=dto.CommentDetails,details=nil} "Success"
// @Failure		400 "invalid path parameter"
// @Failure		404 "comment not found"
// @Failure		500 "internal server error"
// @Router 		/v1/comments/{id} [GET]
// @Security 	BearerAuth
func (h *CommentHandler) GetByID(c *gin.Context) {
	getByID(c, h.service.GetByID)
}

// -----------------------------------------------------------------------------
// other sources that has FK to `Comment`

// GetCommentReplies godoc
// @Summary     Get Comment's Replies
// @Description	Get a list of comment's replies with pagination
// @Tags		comments
// @Produce		json
// @Param		id		path	int	true 	"Unique Comment ID to use as FK"	minimum(1)
// @Param		page	query	int	false	"Page Number (default: 1)"			minimum(1)	default(1)
// @Param		size	query	int	false	"Items per Page (default: 10)"		minimum(1)	maximum(100)	default(10)
// @Success		200	{object} helpers.standardResponse{data=generics.PagedList[dto.CommentList]{items=dto.CommentList},details=nil} "Success"
// @Failure		400 "invalid path parameter | invalid query parameters (for 'size' or 'page')"
// @Failure		404 "there's no any reply for this comment"
// @Failure		500 "internal server error"
// @Router		/v1/comments/{id}/replies [GET]
// @Security	BearerAuth
func (h *CommentHandler) GetReplies(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetReplies)
}
