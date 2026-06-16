package handlers

import (
	"net/http"

	"github.com/hamidgh01/Go-Blog-API/internal/application/services"

	"github.com/gin-gonic/gin"
)

type FollowHandler struct {
	service *services.FollowService
}

func NewFollowHandler(s *services.FollowService) *FollowHandler {
	return &FollowHandler{service: s}
}

func (h *FollowHandler) Follow(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithUser(
		c, h.service.Follow, http.StatusCreated, "followed successfully",
	)
}

func (h *FollowHandler) Unfollow(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithUser(
		c, h.service.UnFollow, http.StatusAccepted, "unfollowed successfully",
	)
}

func (h *FollowHandler) RemoveFollower(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithUser(
		c, h.service.RemoveFollower, http.StatusAccepted, "follower removed successfully",
	)
}
