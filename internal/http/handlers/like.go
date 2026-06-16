package handlers

import (
	"net/http"

	"github.com/hamidgh01/Go-Blog-API/internal/application/services"

	"github.com/gin-gonic/gin"
)

type LikeHandler struct {
	service *services.LikeService
}

func NewLikeHandler(s *services.LikeService) *LikeHandler {
	return &LikeHandler{service: s}
}

func (h *LikeHandler) Like(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithUser(
		c, h.service.Like, http.StatusCreated, "post liked successfully",
	)
}

func (h *LikeHandler) Unlike(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithUser(
		c, h.service.Unlike, http.StatusAccepted, "post unliked successfully",
	)
}
