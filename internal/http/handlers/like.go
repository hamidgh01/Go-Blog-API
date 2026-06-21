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

// Like godoc
// @Summary 	Like
// @Description Like a post
// @Tags 		posts
// @Produce 	json
// @Param 		id	path int true "Unique Post ID to Like" minimum(1)
// @Success 	201 "Success"
// @Failure		400 "invalid path parameter"
// @Failure		406 "no such post to like (trying to like a post that not exists) | the target post is already liked by this user"
// @Failure		500 "internal server error"
// @Router 		/v1/posts/{id}/like [POST]
// @Security BearerAuth
func (h *LikeHandler) Like(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithUser(
		c, h.service.Like, http.StatusCreated, "post liked successfully",
	)
}

// Unlike godoc
// @Summary 	Unlike
// @Description Unlike a post
// @Tags 		posts
// @Produce 	json
// @Param 		id	path int true "Unique Post ID to Unlike" minimum(1)
// @Success 	202 "Success"
// @Failure		400 "invalid path parameter"
// @Failure		406 "no such post to unlike (trying to unlike a post that not exists) | the target post already isn't liked by this user"
// @Failure		500 "internal server error"
// @Router 		/v1/posts/{id}/unlike [DELETE]
// @Security BearerAuth
func (h *LikeHandler) Unlike(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithUser(
		c, h.service.Unlike, http.StatusAccepted, "post unliked successfully",
	)
}
