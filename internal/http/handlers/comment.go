package handlers

import (
	"fmt"
	"net/http"

	"github.com/hamidgh01/Go-Blog-API/internal/application/services"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
	"github.com/hamidgh01/Go-Blog-API/internal/http/helpers"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	service *services.CommentService
}

func NewCommentHandler(s *services.CommentService) *CommentHandler {
	return &CommentHandler{service: s}
}

func (h *CommentHandler) Create(c *gin.Context) {
	create(c, dto.NewCreateCommentRequest, h.service.Create)
}

func (h *CommentHandler) Update(c *gin.Context) {
	update(c, dto.NewUpdateCommentRequest, h.service.Update)
}

func (h *CommentHandler) updateCommentStatus(c *gin.Context, status string) {
	pk, err := extractIDPathParam(c)
	if err != nil {
		return
	}

	data := &dto.UpdateCommentStatusRequest{Status: status}
	commentDetails, serviceErr := h.service.UpdateStatus(c, pk, data)
	if serviceErr != nil {
		c.AbortWithStatusJSON(
			serviceErr.Code(),
			helpers.GenerateErrorResponse(serviceErr.Message(), nil),
		)
		return
	}

	c.JSON(
		http.StatusAccepted,
		helpers.GenerateSuccessfulResponse(
			fmt.Sprintf("comment %s successfully.", status), commentDetails),
	)
}

func (h *CommentHandler) Hide(c *gin.Context) {
	h.updateCommentStatus(c, "rejected")
}

func (h *CommentHandler) Republish(c *gin.Context) {
	h.updateCommentStatus(c, "published")
}

func (h *CommentHandler) DeleteAtUserRequest(c *gin.Context) {
	h.updateCommentStatus(c, "deleted")
}

func (h *CommentHandler) Delete(c *gin.Context) {
	delete(c, h.service.Delete)
}

func (h *CommentHandler) GetByID(c *gin.Context) {
	getByID(c, h.service.GetByID)
}

// -----------------------------------------------------------------------------
// other sources that has FK to `Comment`

func (h *CommentHandler) GetReplies(c *gin.Context) {
	// to implement later
}
