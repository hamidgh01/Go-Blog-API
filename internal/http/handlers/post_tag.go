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

func (h *PostTagsHandler) Associate(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithoutUser(
		c,
		dto.NewAssociatePostWithTagsRequest,
		h.service.Associate,
		http.StatusOK,
		"post associated with tags successfully",
	)
}

func (h *PostTagsHandler) Dissociate(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithoutUser(
		c,
		dto.NewDissociatePostWithTagsRequest,
		h.service.Dissociate,
		http.StatusAccepted,
		"post dissociated with tags successfully",
	)
}
