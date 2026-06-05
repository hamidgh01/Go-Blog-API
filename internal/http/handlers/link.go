package handlers

import (
	"github.com/hamidgh01/Go-Blog-API/internal/application/services"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"

	"github.com/gin-gonic/gin"
)

type LinkHandler struct {
	service *services.LinkService
}

func NewLinkHandler(s *services.LinkService) *LinkHandler {
	return &LinkHandler{service: s}
}

func (h *LinkHandler) Create(c *gin.Context) {
	create(c, dto.NewCreateLinkRequest, h.service.Create)
}

func (h *LinkHandler) Update(c *gin.Context) {
	update(c, dto.NewUpdateLinkRequest, h.service.Update)
}

func (h *LinkHandler) Delete(c *gin.Context) {
	delete(c, h.service.Delete)
}

func (h *LinkHandler) GetByID(c *gin.Context) {
	getByID(c, h.service.GetByID)
}
