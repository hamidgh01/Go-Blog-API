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

// SaveList godoc
// @Summary 	Save List
// @Description Save a list
// @Tags 		lists
// @Produce 	json
// @Param 		id	path int true "Unique List ID to Save" minimum(1)
// @Success 	201 "Success"
// @Failure		400 "invalid path parameter"
// @Failure		406 "no such list to save (trying to save a list that not exists) | list is already saved by this user"
// @Failure		500 "internal server error"
// @Router 		/v1/lists/{id}/save-list [POST]
// @Security BearerAuth
func (h *SaveListHandler) Save(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithUser(
		c, h.service.Save, http.StatusCreated, "list saved successfully",
	)
}

// UnsaveList godoc
// @Summary 	Unsave List
// @Description Unsave a list
// @Tags 		lists
// @Produce 	json
// @Param 		id	path int true "Unique List ID to Unsave" minimum(1)
// @Success 	202 "Success"
// @Failure		400 "invalid path parameter"
// @Failure		406 "no such list to unsave (trying to unsave a list that not exists) | list already isn't saved by this user"
// @Failure		500 "internal server error"
// @Router 		/v1/lists/{id}/unsave-list [DELETE]
// @Security BearerAuth
func (h *SaveListHandler) Unsave(c *gin.Context) {
	AddOrDeleteM2MRelationShipWithUser(
		c, h.service.Unsave, http.StatusAccepted, "list unsaved successfully",
	)
}
