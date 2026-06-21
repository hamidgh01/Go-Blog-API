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

// Follow godoc
// @Summary 	Follow
// @Description Follow a user
// @Tags 		users
// @Produce 	json
// @Param 		id	path int true "Unique User ID to Follow" minimum(1)
// @Success 	201 "Success"
// @Failure		400 "invalid path parameter"
// @Failure		406 "no such user to follow (trying to follow a user that not exists) | the target user is already followed"
// @Failure		500 "internal server error"
// @Router 		/v1/users/{id}/follow [POST]
// @Security BearerAuth
func (h *FollowHandler) Follow(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithUser(
		c, h.service.Follow, http.StatusCreated, "followed successfully",
	)
}

// Unfollow godoc
// @Summary 	Unfollow
// @Description Unfollow a following
// @Tags 		users
// @Produce 	json
// @Param 		id	path int true "Unique User ID to Unfollow" minimum(1)
// @Success 	202 "Success"
// @Failure		400 "invalid path parameter"
// @Failure		406 "no such user to unfollow (trying to unfollow a user that not exists) | the target user already is not followed current user"
// @Failure		500 "internal server error"
// @Router 		/v1/users/{id}/unfollow [DELETE]
// @Security BearerAuth
func (h *FollowHandler) Unfollow(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithUser(
		c, h.service.UnFollow, http.StatusAccepted, "unfollowed successfully",
	)
}

// RemoveFollower godoc
// @Summary 	Remove Follower
// @Description Remove a follower
// @Tags 		users
// @Produce 	json
// @Param 		id	path int true "Unique User (Follower) ID to Remove from followers" minimum(1)
// @Success 	202 "Success"
// @Failure		400 "invalid path parameter"
// @Failure		406 "no such user to remove from followers (trying to remove a user that not exists) | the target user already is not followed current user"
// @Failure		500 "internal server error"
// @Router 		/v1/users/{id}/remove-follower [DELETE]
// @Security BearerAuth
func (h *FollowHandler) RemoveFollower(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithUser(
		c, h.service.RemoveFollower, http.StatusAccepted, "follower removed successfully",
	)
}
