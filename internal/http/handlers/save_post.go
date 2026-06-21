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

// SavePost godoc
// @Summary 	Save Post
// @Description Save a post
// @Tags 		posts
// @Accept  	json
// @Produce 	json
// @Param 		id		path int 					true "Unique Post ID to Save" minimum(1)
// @Param 		request body dto.SavePostRequest	true "Unique List ID to Save in"
// @Success 	201 "Success"
// @Failure		400 "validation error | invalid path parameter"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		403 "trying to save to other's lists (each user only can save posts in its own lists)"
// @Failure		404 "no such list to save post in (trying to save post in a list that not exists)"
// @Failure		406 "no such post to save (trying to save a post that not exists) | post is already saved in this list"
// @Failure		500 "internal server error"
// @Router 		/v1/posts/{id}/save-post [POST]
// @Security BearerAuth
func (h *SavePostHandler) Save(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithoutUser(
		c, dto.NewSavePostRequest, h.service.Save, http.StatusOK, "post saved successfully",
	)
}

// UnsavePost godoc
// @Summary 	Unsave Post
// @Description Unsave a post
// @Tags 		posts
// @Accept  	json
// @Produce 	json
// @Param 		id		path int 					true "Unique Post ID to Unsave" minimum(1)
// @Param 		request body dto.UnsavePostRequest	true "Unique List ID to Unsave From"
// @Success 	202 "Success"
// @Failure		400 "validation error | invalid path parameter"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		403 "trying to unsave a post from other's lists (each user can only unsave posts from its own lists)"
// @Failure		404 "no such list to unsave post from (trying to unsave a post from a list that not exists)"
// @Failure		406 "the post (with entered id) already isn't saved in this list"
// @Failure		500 "internal server error"
// @Router 		/v1/posts/{id}/unsave-post [DELETE]
// @Security BearerAuth
func (h *SavePostHandler) Unsave(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithoutUser(
		c, dto.NewUnsavePostRequest, h.service.Unsave, http.StatusAccepted, "post unsaved successfully",
	)
}
