package handlers

import (
	"net/http"

	"github.com/hamidgh01/Go-Blog-API/internal/application/services"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
	"github.com/hamidgh01/Go-Blog-API/internal/http/helpers"
	"github.com/hamidgh01/Go-Blog-API/internal/http/validations"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(s *services.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Register(c *gin.Context) {
	data := new(dto.CreateUserRequest)
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

	successfulRegisteredData, serviceErr := h.service.Register(c, data, c.Writer)
	if serviceErr != nil {
		c.AbortWithStatusJSON(
			serviceErr.Code(),
			helpers.GenerateErrorResponse(serviceErr.Message(), nil),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		helpers.GenerateSuccessfulResponse("user registered successfully.", successfulRegisteredData),
	)
}

func (h *AuthHandler) Login(c *gin.Context) {
	data := new(dto.LoginRequest)
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

	successfulLoginData, serviceErr := h.service.Login(c, data, c.Writer)
	if serviceErr != nil {
		c.AbortWithStatusJSON(
			serviceErr.Code(),
			helpers.GenerateErrorResponse(serviceErr.Message(), nil),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		helpers.GenerateSuccessfulResponse("successfully logged in.", successfulLoginData),
	)
}

func (h *AuthHandler) RenewTokens(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			helpers.GenerateErrorResponse("'refresh_token' cookie Not provided", nil),
		)
		return
	}

	accessTokenDetails, serviceErr := h.service.RenewTokens(c, refreshToken, c.Writer)
	if serviceErr != nil {
		c.AbortWithStatusJSON(
			serviceErr.Code(),
			helpers.GenerateErrorResponse(serviceErr.Message(), nil),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		helpers.GenerateSuccessfulResponse("tokens refreshed successfully.", accessTokenDetails),
	)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			helpers.GenerateErrorResponse("'refresh_token' cookie Not provided", nil),
		)
		return
	}

	if serviceErr := h.service.Logout(c, refreshToken, c.Writer); serviceErr != nil {
		c.AbortWithStatusJSON(
			serviceErr.Code(),
			helpers.GenerateErrorResponse(serviceErr.Message(), nil),
		)
		return
	}

	c.JSON(http.StatusOK, helpers.GenerateSuccessfulResponse("successfully logged out.", nil))
}
