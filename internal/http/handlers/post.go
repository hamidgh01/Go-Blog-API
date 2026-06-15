package handlers

import (
	"github.com/hamidgh01/Go-Blog-API/internal/application/services"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service *services.PostService
}

func NewPostHandler(s *services.PostService) *PostHandler {
	return &PostHandler{service: s}
}

func (h *PostHandler) Create(c *gin.Context) {
	create(c, dto.NewCreatePostRequest, h.service.Create)
}

func (h *PostHandler) Update(c *gin.Context) {
	update(c, dto.NewUpdatePostRequest, h.service.Update)
}

func (h *PostHandler) UpdatePrivacy(c *gin.Context) {
	update(c, dto.NewUpdatePostPrivacyRequest, h.service.UpdatePrivacy)
}

func (h *PostHandler) Publish(c *gin.Context) {
	requestDTO := &dto.UpdatePostStatusRequest{Status: "published"}
	updateStatusEnum(c, requestDTO, h.service.PublishDraftPost)
}

func (h *PostHandler) Reject(c *gin.Context) {
	requestDTO := &dto.UpdatePostStatusRequest{Status: "rejected"}
	updateStatusEnum(c, requestDTO, h.service.UpdateStatus)
}

func (h *PostHandler) Republish(c *gin.Context) {
	requestDTO := &dto.UpdatePostStatusRequest{Status: "published"}
	updateStatusEnum(c, requestDTO, h.service.UpdateStatus)
}

func (h *PostHandler) DeleteAtUserRequest(c *gin.Context) {
	requestDTO := &dto.UpdatePostStatusRequest{Status: "deleted-by-author"}
	updateStatusEnum(c, requestDTO, h.service.UpdateStatus)
}

func (h *PostHandler) Delete(c *gin.Context) {
	delete(c, h.service.Delete)
}

func (h *PostHandler) GetByID(c *gin.Context) {
	getByID(c, h.service.GetByID)
}

// -----------------------------------------------------------------------------
// other sources that has FK to `Post`

func (h *PostHandler) GetComments(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetComments)
}

func (h *PostHandler) GetLikes(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetLikes)
}

func (h *PostHandler) GetTags(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetTags)
}

func (h *PostHandler) GetLists(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetLists)
}
