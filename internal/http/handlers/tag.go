package handlers

import (
	"fmt"
	"net/http"

	"github.com/hamidgh01/Go-Blog-API/internal/application/services"
	"github.com/hamidgh01/Go-Blog-API/internal/domain"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
	_ "github.com/hamidgh01/Go-Blog-API/internal/http/generics"
	"github.com/hamidgh01/Go-Blog-API/internal/http/helpers"
	"github.com/hamidgh01/Go-Blog-API/internal/http/validations"

	"github.com/gin-gonic/gin"
)

type TagHandler struct {
	service *services.TagService
}

func NewTagHandler(s *services.TagService) *TagHandler {
	return &TagHandler{service: s}
}

// Create Tags godoc
// @Summary 	Create Tags
// @Description Create some tags
// @Tags 		tags
// @Accept  	json
// @Produce 	json
// @Param 		request body dto.CreateTagsRequest true "Array of valid tag names"
// @Success 	201 {object} helpers.standardResponse{data=dto.TagsList,details=nil} "Success"
// @Failure		400 "validation error"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		500 "internal server error"
// @Router 		/v1/tags [POST]
// @Security BearerAuth
func (h *TagHandler) Create(c *gin.Context) {
	data := dto.NewCreateTagsRequest()
	if err := c.ShouldBindJSON(data); err != nil {
		if translatedVldErrs := validations.GetValidationErrors(err); translatedVldErrs != nil {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				helpers.GenerateErrorResponse("invalid input (there is validation errors)", translatedVldErrs),
			)
			return
		}

		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			helpers.GenerateErrorResponse(
				"invalid request body (json)",
				gin.H{"description": "json format is invalid, or fields have invalid type"},
			),
		)
		return
	}

	tagsList, serviceErr := h.service.Create(c, data)
	if serviceErr != nil {
		c.AbortWithStatusJSON(
			serviceErr.Code(),
			helpers.GenerateErrorResponse(serviceErr.Message(), nil),
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		helpers.GenerateSuccessfulResponse("tags created successfully.", tagsList),
	)
}

// GetTagByID godoc
// @Summary 	Get Tag By ID
// @Description Get a tag by unique id
// @Tags 		tags
// @Produce 	json
// @Param		id path int true "Unique Tag ID" minimum(1)
// @Success 	200 {object} helpers.standardResponse{data=dto.TagDetails,details=nil} "Success"
// @Failure		400 "invalid path parameter"
// @Failure		404 "tag not found"
// @Failure		500 "internal server error"
// @Router 		/v1/tags/{id} [GET]
// @Security 	BearerAuth
func (h *TagHandler) GetByID(c *gin.Context) {
	getByID(c, h.service.GetByID)
}

// GetTagByName godoc
// @Summary 	Get Tag By Name
// @Description Get a tag by unique tag-name
// @Tags 		tags
// @Produce 	json
// @Param		name path string true "Unique Tag Name"
// @Success 	200 {object} helpers.standardResponse{data=dto.TagDetails,details=nil} "Success"
// @Failure		400 "invalid path parameter"
// @Failure		404 "tag not found"
// @Failure		500 "internal server error"
// @Router 		/v1/tags/name/{name} [GET]
// @Security 	BearerAuth
func (h *TagHandler) GetByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" || !domain.CheckTagPattern(name) {
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

// GetTagPosts godoc
// @Summary     Get Tag's Posts
// @Description	Get a list of tag's posts with pagination
// @Tags		tags
// @Produce		json
// @Param		id 		path	int	true 	"Unique Tag ID to use as FK"	minimum(1)
// @Param		page	query	int	false	"Page Number (default: 1)"		minimum(1)	default(1)
// @Param		size	query	int	false	"Items per Page (default: 10)"	minimum(1)	maximum(100)	default(10)
// @Success		200	{object} helpers.standardResponse{data=generics.PagedList[dto.PostsList]{items=dto.PostsList},details=nil} "Success"
// @Failure		400 "invalid path parameter | invalid query parameters (for 'size' or 'page')"
// @Failure		404 "there's no any post for this tag"
// @Failure		500 "internal server error"
// @Router		/v1/tags/{id}/posts [GET]
// @Security	BearerAuth
func (h *TagHandler) GetPosts(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetPosts)
}
