package handlers

import (
	"fmt"
	"net/http"

	"github.com/hamidgh01/Go-Blog-API/internal/application/services"
	"github.com/hamidgh01/Go-Blog-API/internal/domain"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
	"github.com/hamidgh01/Go-Blog-API/internal/http/helpers"

	"github.com/gin-gonic/gin"
)

type TagHandler struct {
	service *services.TagService
}

func NewTagHandler(s *services.TagService) *TagHandler {
	return &TagHandler{service: s}
}

func (h *TagHandler) Create(c *gin.Context) {
	create(c, dto.NewCreateTagsRequest, h.service.Create)
}

func (h *TagHandler) GetByID(c *gin.Context) {
	getByID(c, h.service.GetByID)
}

func (h *TagHandler) GetByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" || domain.CheckTagPattern(name) {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			helpers.GenerateErrorResponse(
				fmt.Sprintf("invalid path parameter: '%s'", name),
				gin.H{"description": fmt.Sprintf("input must be a valid tag-name. %s", domain.TagPatternDescription)},
			),
		)
		return
	}

	tagDetails, serviceErr := h.service.GetByName(c, name)
	if serviceErr != nil {
		c.AbortWithStatusJSON(
			serviceErr.Code(),
			helpers.GenerateErrorResponse(serviceErr.Message(), nil),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		helpers.GenerateSuccessfulResponse("object fetched successfully.", tagDetails), // typeName
	)
}

// -----------------------------------------------------------------------------
// other sources that has FK to `Tag`

func (h *TagHandler) GetPosts(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetPosts)
}
