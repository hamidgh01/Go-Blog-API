package handlers

import (
	"github.com/hamidgh01/Go-Blog-API/internal/application/services"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"

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

func (h *CommentHandler) Hide(c *gin.Context) {
	requestDTO := &dto.UpdateCommentStatusRequest{Status: "hidden-by-Admin"}
	updateStatusEnum(c, requestDTO, h.service.UpdateStatus)
}

func (h *CommentHandler) Republish(c *gin.Context) {
	requestDTO := &dto.UpdateCommentStatusRequest{Status: "published"}
	updateStatusEnum(c, requestDTO, h.service.UpdateStatus)
}

func (h *CommentHandler) DeleteAtUserRequest(c *gin.Context) {
	requestDTO := &dto.UpdateCommentStatusRequest{Status: "deleted-by-commenter"}
	updateStatusEnum(c, requestDTO, h.service.UpdateStatus)
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
	getListOfOuterResourceByFK(c, h.service.GetReplies)
}
