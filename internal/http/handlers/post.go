package handlers

import (
	"fmt"
	"net/http"

	"Go-Blog-API/internal/application/services"
	"Go-Blog-API/internal/http/dto"
	"Go-Blog-API/internal/http/helpers"

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

func (h *PostHandler) updatePostStatus(c *gin.Context, status string) {
	pk, err := extractIDPathParam(c)
	if err != nil {
		return
	}

	data := &dto.UpdatePostStatusRequest{Status: status}

	postResponse, serviceErr := h.service.UpdateStatus(c, pk, data)
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
			fmt.Sprintf("post %s successfully.", status), postResponse),
	)
}

func (h *PostHandler) Publish(c *gin.Context) {
	h.updatePostStatus(c, "published")
}

func (h *PostHandler) Reject(c *gin.Context) {
	h.updatePostStatus(c, "rejected")
}

func (h *PostHandler) Republish(c *gin.Context) {
	h.updatePostStatus(c, "published")
}

func (h *PostHandler) DeleteAtUserRequest(c *gin.Context) {
	h.updatePostStatus(c, "deleted")
}

func (h *PostHandler) Delete(c *gin.Context) {
	delete(c, h.service.Delete)
}

func (h *PostHandler) GetByID(c *gin.Context) {
	getByID(c, h.service.GetByID)
}
