package handlers

import (
	"github.com/hamidgh01/Go-Blog-API/internal/application/services"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"

	"github.com/gin-gonic/gin"
)

type ListHandler struct {
	service *services.ListService
}

func NewListHandler(s *services.ListService) *ListHandler {
	return &ListHandler{service: s}
}

func (h *ListHandler) Create(c *gin.Context) {
	create(c, dto.NewCreateListRequest, h.service.Create)
}

func (h *ListHandler) Update(c *gin.Context) {
	update(c, dto.NewUpdateListRequest, h.service.Update)
}

func (h *ListHandler) UpdatePrivacy(c *gin.Context) {
	update(c, dto.NewUpdateListPrivacyRequest, h.service.UpdatePrivacy)
}

func (h *ListHandler) Delete(c *gin.Context) {
	delete(c, h.service.Delete)
}

func (h *ListHandler) GetByID(c *gin.Context) {
	getByID(c, h.service.GetByID)
}

// -----------------------------------------------------------------------------
// other sources that has FK to `List`

func (h *ListHandler) GetSavedPosts(c *gin.Context) {
	// to implement later
}

func (h *ListHandler) GetUsersWhoSaved(c *gin.Context) {
	// to implement later
}
