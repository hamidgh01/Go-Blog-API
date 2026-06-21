package handlers

import (
	"net/http"

	"github.com/hamidgh01/Go-Blog-API/internal/application/services"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"

	"github.com/gin-gonic/gin"
)

type PostTagsHandler struct {
	service *services.PostTagsService
}

func NewPostTagsHandler(s *services.PostTagsService) *PostTagsHandler {
	return &PostTagsHandler{service: s}
}

// Associate godoc
// @Summary 	Tag Post
// @Description Tag a post with some tags
// @Tags 		posts
// @Accept  	json
// @Produce 	json
// @Param 		id		path int 								true "Unique Post ID to Tag" minimum(1)
// @Param 		request body dto.AssociatePostWithTagsRequest	true "Array of Tag IDs to Associate with Post"
// @Success 	201 "Success"
// @Failure		400 "validation error | invalid path parameter"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		403 "trying to tag other's posts (each user only can tag its own posts)"
// @Failure		404 "no such post to tag (trying to tag a post that not exists)"
// @Failure		406 "no such tag to associate with post (trying to tag a post with non-existent tags) | the post already associated with all of these tags"
// @Failure		500 "internal server error"
// @Router 		/v1/posts/{id}/associate-tags [POST]
// @Security BearerAuth
func (h *PostTagsHandler) Associate(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithoutUser(
		c,
		dto.NewAssociatePostWithTagsRequest,
		h.service.Associate,
		http.StatusOK,
		"post associated with tags successfully",
	)
}

// Dissociate godoc
// @Summary 	Untag Post
// @Description Untag a post with some tags
// @Tags 		posts
// @Accept  	json
// @Produce 	json
// @Param 		id		path int 								true "Unique Post ID to Untag" minimum(1)
// @Param 		request body dto.DissociatePostWithTagsRequest	true "Array of Tag IDs to dissociate with Post"
// @Success 	202 "Success"
// @Failure		400 "validation error | invalid path parameter"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		403 "trying to untag other's posts (each user only can untag its own posts)"
// @Failure		404 "no such post to untag (trying to untag a post that not exists)"
// @Failure		406 "no such tag to dissociate with post (trying to untag a post with non-existent tags) | the post didn't tagged with any of these tags"
// @Failure		500 "internal server error"
// @Router 		/v1/posts/{id}/dissociate-tags [DELETE]
// @Security BearerAuth
func (h *PostTagsHandler) Dissociate(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithoutUser(
		c,
		dto.NewDissociatePostWithTagsRequest,
		h.service.Dissociate,
		http.StatusAccepted,
		"post dissociated with tags successfully",
	)
}
