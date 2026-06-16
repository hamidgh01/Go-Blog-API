package handlers

import (
	"net/http"

	"github.com/hamidgh01/Go-Blog-API/internal/application/services"

	"github.com/gin-gonic/gin"
)

type SaveListHandler struct {
	service *services.SaveListService
}

func NewSaveListHandler(s *services.SaveListService) *SaveListHandler {
	return &SaveListHandler{service: s}
}

func (h *SaveListHandler) Save(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithUser(
		c, h.service.Save, http.StatusCreated, "list saved successfully",
	)
}

func (h *SaveListHandler) Unsave(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithUser(
		c, h.service.Unsave, http.StatusAccepted, "list unsaved successfully",
	)
}
