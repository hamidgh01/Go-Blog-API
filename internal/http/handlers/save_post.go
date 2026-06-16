package handlers

import (
	"net/http"

	"github.com/hamidgh01/Go-Blog-API/internal/application/services"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"

	"github.com/gin-gonic/gin"
)

type SavePostHandler struct {
	service *services.SavePostService
}

func NewSavePostHandler(s *services.SavePostService) *SavePostHandler {
	return &SavePostHandler{service: s}
}

func (h *SavePostHandler) Save(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithoutUser(
		c, dto.NewSavePostRequest, h.service.Save, http.StatusOK, "post saved successfully",
	)
}

func (h *SavePostHandler) Unsave(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithoutUser(
		c, dto.NewUnsavePostRequest, h.service.Unsave, http.StatusAccepted, "post unsaved successfully",
	)
}
