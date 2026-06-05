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

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) Create(c *gin.Context) {
	create(c, dto.NewCreateUserRequest, h.service.Create)
}

func (h *UserHandler) UpdateUsername(c *gin.Context) {
	update(c, dto.NewUpdateUsernameRequest, h.service.UpdateUsername)
}

func (h *UserHandler) UpdateEmail(c *gin.Context) {
	update(c, dto.NewUpdateEmailRequest, h.service.UpdateEmail)
}

func (h *UserHandler) UpdateBio(c *gin.Context) {
	update(c, dto.NewUpdateBioRequest, h.service.UpdateBio)
}

func (h *UserHandler) UpdatePassword(c *gin.Context) {
	update(c, dto.NewUpdatePasswordRequest, h.service.UpdatePassword)
}

func (h *UserHandler) UpdateEnabled(c *gin.Context) {
	update(c, dto.NewUpdateEnabledRequest, h.service.UpdateEnabled)
}

func (h *UserHandler) Delete(c *gin.Context) {
	delete(c, h.service.Delete)
}

// func (h *UserHandler) GetList(c *gin.Context) {}

func (h *UserHandler) GetByID(c *gin.Context) {
	getByID(c, h.service.GetByID)
}

func (h *UserHandler) GetByUsername(c *gin.Context) {
	username := c.Param("username")
	if username == "" || domain.CheckUsernamePattern(username) {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			helpers.GenerateErrorResponse(
				fmt.Sprintf("invalid path parameter: '%s'", username),
				gin.H{"description": fmt.Sprintf("input must be a valid username. %s", domain.UsernamePatternDescription)},
			),
		)
		return
	}

	userDetails, serviceErr := h.service.GetByUsername(c, username)
	if serviceErr != nil {
		c.AbortWithStatusJSON(
			serviceErr.Code(),
			helpers.GenerateErrorResponse(serviceErr.Message(), nil),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		helpers.GenerateSuccessfulResponse("object fetched successfully.", userDetails), // typeName
	)
}

func (h *UserHandler) GetByEmail(c *gin.Context) {
	email := c.Param("email")
	if email == "" { // check email
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			helpers.GenerateErrorResponse(
				fmt.Sprintf("invalid path parameter: '%s'", email),
				gin.H{"description": "input must be a valid email address"},
			),
		)
		return
	}

	userDetails, serviceErr := h.service.GetByEmail(c, email)
	if serviceErr != nil {
		c.AbortWithStatusJSON(
			serviceErr.Code(),
			helpers.GenerateErrorResponse(serviceErr.Message(), nil),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		helpers.GenerateSuccessfulResponse("object fetched successfully.", userDetails), // typeName
	)
}

// -----------------------------------------------------------------------------
// other sources that has FK to `User`

func (h *UserHandler) GetPosts(c *gin.Context) {
	// to implement later
}

func (h *UserHandler) GetLists(c *gin.Context) {
	// to implement later
}

func (h *UserHandler) GetSavedLists(c *gin.Context) {
	// to implement later
}

func (h *UserHandler) GetComments(c *gin.Context) {
	// to implement later
}

func (h *UserHandler) GetLikes(c *gin.Context) {
	// to implement later
}

func (h *UserHandler) GetFollowers(c *gin.Context) {
	// to implement later
}

func (h *UserHandler) GetFollowings(c *gin.Context) {
	// to implement later
}

func (h *UserHandler) GetLinks(c *gin.Context) {
	// to implement later
}
