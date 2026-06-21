package handlers

import (
	"github.com/hamidgh01/Go-Blog-API/internal/application/services"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
	_ "github.com/hamidgh01/Go-Blog-API/internal/http/helpers"

	"github.com/gin-gonic/gin"
)

type LinkHandler struct {
	service *services.LinkService
}

func NewLinkHandler(s *services.LinkService) *LinkHandler {
	return &LinkHandler{service: s}
}

// Create Link godoc
// @Summary 	Create Link
// @Description Create a new link
// @Tags 		links
// @Accept  	json
// @Produce 	json
// @Param 		request body dto.CreateLinkRequest true "Data to create new link"
// @Success		201 {object} helpers.standardResponse{data=dto.LinkDetails,details=nil} "Success"
// @Failure		400 "validation error"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		500 "internal server error"
// @Router 		/v1/links [POST]
// @Security BearerAuth
func (h *LinkHandler) Create(c *gin.Context) {
	create(c, dto.NewCreateLinkRequest, h.service.Create)
}

// Updata Link godoc
// @Summary 	Update Link
// @Description Update a link
// @Tags 		links
// @Accept  	json
// @Produce 	json
// @Param		id		path int					true "Unique Link ID" minimum(1)
// @Param 		request body dto.UpdateLinkRequest	true "Data to update a link"
// @Success 	202 "Success"
// @Failure		400 "validation error | invalid path parameter"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		404 "link not found"
// @Failure		500 "internal server error"
// @Router 		/v1/links/{id} [PUT]
// @Security 	BearerAuth
func (h *LinkHandler) Update(c *gin.Context) {
	update(c, dto.NewUpdateLinkRequest, h.service.Update)
}

// Delete Link godoc
// @Summary 	Delete Link
// @Description Delete a link
// @Tags 		links
// @Produce 	json
// @Param		id 	path int true "Unique Link ID" minimum(1)
// @Success 	202 "Success"
// @Failure		400 "invalid path parameter"
// @Failure		404 "link not found"
// @Failure		500 "internal server error"
// @Router 		/v1/links/{id} [DELETE]
// @Security 	BearerAuth
func (h *LinkHandler) Delete(c *gin.Context) {
	delete(c, h.service.Delete)
}

// GetLinkByID godoc
// @Summary 	Get Link By ID
// @Description Get a link by unique id
// @Tags 		links
// @Produce 	json
// @Param		id path int true "Unique Link ID" minimum(1)
// @Success 	200 {object} helpers.standardResponse{data=dto.LinkDetails,details=nil} "Success"
// @Failure		400 "invalid path parameter"
// @Failure		404 "link not found"
// @Failure		500 "internal server error"
// @Router 		/v1/links/{id} [GET]
// @Security 	BearerAuth
func (h *LinkHandler) GetByID(c *gin.Context) {
	getByID(c, h.service.GetByID)
}
