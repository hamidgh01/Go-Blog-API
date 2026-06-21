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

// Register godoc
// @Summary 	Register
// @Description Register (sign up) new user
// @Tags 		auth
// @Accept  	json
// @Produce 	json
// @Param 		request body dto.CreateUserRequest true "User registration data"
// @Success		201 {object} helpers.standardResponse{data=dto.LoginSuccessfulData,details=nil} "Success"
// @Failure		400 "validation error"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		409 "username/email already exists"
// @Failure		500 "internal server error"
// @Router 		/v1/register [POST]
func (h *AuthHandler) Register(c *gin.Context) {
	data := dto.NewCreateUserRequest()
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
		http.StatusCreated,
		helpers.GenerateSuccessfulResponse("user registered successfully.", successfulRegisteredData),
	)
}

// Login godoc
// @Summary 	Login
// @Description Login by identifier (**username** or **email**) and **password**
// @Tags 		auth
// @Accept  	json
// @Produce 	json
// @Param 		request body dto.LoginRequest true "Login credentials"
// @Success		200 {object} helpers.standardResponse{data=dto.LoginSuccessfulData,details=nil} "Success"
// @Failure		400 "validation error"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		401 "invalid credentials"
// @Failure		403 "suspended user"
// @Failure		500 "internal server error"
// @Router 		/v1/login [POST]
func (h *AuthHandler) Login(c *gin.Context) {
	data := dto.NewLoginRequest()
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

// RenewTokens godoc
// @Summary 	RenewTokens
// @Description Get new access token (and a new refresh_token cookie will be set)
// @Tags 		auth
// @Produce 	json
// @Success		200 {object} helpers.standardResponse{data=dto.Token,details=nil} "Success"
// @Failure		401 "invalid token | expired token | blacklisted token | not providing 'refresh_token' cookie"
// @Failure		500 "internal server error"
// @Router 		/v1/renew-tokens [GET]
// @Security 	BearerAuth
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

// Logout godoc
// @Summary 	Logout
// @Description Logout a user
// @Tags 		auth
// @Produce 	json
// @Success 	200 "Success"
// @Failure		401 "invalid token | expired token | blacklisted token | not providing 'refresh_token' cookie"
// @Failure		500 "internal server error"
// @Router 		/v1/logout [GET]
// @Security 	BearerAuth
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
