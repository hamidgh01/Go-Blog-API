package handlers

import (
	"fmt"
	"net/http"

	"github.com/hamidgh01/Go-Blog-API/internal/application/services"
	"github.com/hamidgh01/Go-Blog-API/internal/domain"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
	_ "github.com/hamidgh01/Go-Blog-API/internal/http/generics"
	"github.com/hamidgh01/Go-Blog-API/internal/http/helpers"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// Create User godoc
// @Summary 	Create User
// @Description Create a new user (**only admin access**)
// @Tags 		users
// @Accept  	json
// @Produce 	json
// @Param 		request body dto.CreateUserRequest true "Data to create new user"
// @Success		201 {object} helpers.standardResponse{data=dto.UserDetails,details=nil} "Success"
// @Failure		400 "validation error"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		409 "username/email already exists"
// @Failure		500 "internal server error"
// @Router 		/v1/users [POST]
// @Security BearerAuth
func (h *UserHandler) Create(c *gin.Context) {
	create(c, dto.NewCreateUserRequest, h.service.Create)
}

// Updata User.Username godoc
// @Summary 	Updata User.Username
// @Description Update user's username
// @Tags 		users
// @Accept  	json
// @Produce 	json
// @Param		id		path int                      	true "Unique User ID" minimum(1)
// @Param 		request	body dto.UpdateUsernameRequest 	true "new username value to update User.Username"
// @Success 	202 "Success"
// @Failure		400 "validation error | invalid path parameter"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		404 "user not found"
// @Failure		409 "username already exists"
// @Failure		500 "internal server error"
// @Router 		/v1/users/{id}/username [PATCH]
// @Security 	BearerAuth
func (h *UserHandler) UpdateUsername(c *gin.Context) {
	update(c, dto.NewUpdateUsernameRequest, h.service.UpdateUsername)
}

// Updata User.Email godoc
// @Summary 	Updata User.Email
// @Description Update user's email
// @Tags 		users
// @Accept  	json
// @Produce 	json
// @Param		id		path int					true "Unique User ID" minimum(1)
// @Param 		request	body dto.UpdateEmailRequest	true "new email address to update User.Email"
// @Success 	202 "Success"
// @Failure		400 "validation error | invalid path parameter"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		404 "user not found"
// @Failure		409 "email already exists"
// @Failure		500 "internal server error"
// @Router 		/v1/users/{id}/email [PATCH]
// @Security 	BearerAuth
func (h *UserHandler) UpdateEmail(c *gin.Context) {
	update(c, dto.NewUpdateEmailRequest, h.service.UpdateEmail)
}

// Updata User.Bio godoc
// @Summary 	Updata User.Bio
// @Description Update user's bio
// @Tags 		users
// @Accept  	json
// @Produce 	json
// @Param		id		path int					true "Unique User ID" minimum(1)
// @Param 		request	body dto.UpdateBioRequest	true "new bio to update User.Bio"
// @Success 	202 "Success"
// @Failure		400 "validation error | invalid path parameter"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		404 "user not found"
// @Failure		500 "internal server error"
// @Router 		/v1/users/{id}/bio [PATCH]
// @Security 	BearerAuth
func (h *UserHandler) UpdateBio(c *gin.Context) {
	update(c, dto.NewUpdateBioRequest, h.service.UpdateBio)
}

// Updata User.Password godoc
// @Summary 	Updata User.Password
// @Description Update user's password
// @Tags 		users
// @Accept  	json
// @Produce 	json
// @Param		id		path int						true "Unique User ID" minimum(1)
// @Param 		request	body dto.UpdatePasswordRequest	true "old, new, and confirm password to update User.Password"
// @Success 	202 "Success"
// @Failure		400 "validation error | invalid path parameter"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		404 "user not found"
// @Failure		500 "internal server error"
// @Router 		/v1/users/{id}/password [PATCH]
// @Security 	BearerAuth
func (h *UserHandler) UpdatePassword(c *gin.Context) {
	update(c, dto.NewUpdatePasswordRequest, h.service.UpdatePassword)
}

// Updata User.Enabled godoc
// @Summary 	Updata User.Enabled
// @Description Update user's enablement (*activation* or *deactivation*) (**only admin access**)
// @Tags 		users
// @Accept  	json
// @Produce 	json
// @Param		id		path int						true "Unique User ID" minimum(1)
// @Param 		request	body dto.UpdateEnabledRequest	true "new enabled value to update User.Enabled"
// @Success 	202 "Success"
// @Failure		400 "validation error | invalid path parameter"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		404 "user not found"
// @Failure		500 "internal server error"
// @Router 		/v1/users/{id}/enabled [PATCH]
// @Security 	BearerAuth
func (h *UserHandler) UpdateEnabled(c *gin.Context) {
	update(c, dto.NewUpdateEnabledRequest, h.service.UpdateEnabled)
}

// CheckUsernameExists godoc
// @Summary 	Check Username Exists
// @Description Check a username is taken before or not
// @Tags 		users
// @Produce 	json
// @Param		q query string true "a valid username"
// @Success 	200 "Success"
// @Failure		400 "invalid query parameter"
// @Failure		500 "internal server error"
// @Router 		/v1/users/exists/username [GET]
// @Security 	BearerAuth
func (h *UserHandler) CheckUsernameExists(c *gin.Context) {
	username := c.Query("q")
	if username == "" || !domain.CheckUsernamePattern(username) {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			helpers.GenerateErrorResponse(
				fmt.Sprintf("invalid query parameter: '%s'", username),
				gin.H{"description": fmt.Sprintf("input must be a valid username. %s", domain.UsernamePatternDescription)},
			),
		)
		return
	}

	exists, serviceErr := h.service.CheckUsernameExists(c, username)
	if serviceErr != nil {
		c.AbortWithStatusJSON(
			serviceErr.Code(),
			helpers.GenerateErrorResponse(serviceErr.Message(), nil),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		helpers.GenerateSuccessfulResponse("username existence checked successfully.", gin.H{"exists": exists}),
	)
}

// CheckEmailExists godoc
// @Summary 	Check Email Exists
// @Description Check an email is registered before or not
// @Tags 		users
// @Produce 	json
// @Param		q query string true "a valid email address"
// @Success 	200 "Success"
// @Failure		400 "invalid query parameter"
// @Failure		500 "internal server error"
// @Router 		/v1/users/exists/email [GET]
// @Security 	BearerAuth
func (h *UserHandler) CheckEmailExists(c *gin.Context) {
	email := c.Query("q")
	if email == "" || !domain.CheckEmailPattern(email) {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			helpers.GenerateErrorResponse(
				fmt.Sprintf("invalid query parameter: '%s'", email),
				gin.H{"description": "input must be a valid email address"},
			),
		)
		return
	}

	exists, serviceErr := h.service.CheckEmailExists(c, email)
	if serviceErr != nil {
		c.AbortWithStatusJSON(
			serviceErr.Code(),
			helpers.GenerateErrorResponse(serviceErr.Message(), nil),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		helpers.GenerateSuccessfulResponse("email existence checked successfully.", gin.H{"exists": exists}),
	)
}

// Delete User godoc
// @Summary 	Delete User
// @Description Delete a user (**only admin access**)
// @Tags 		users
// @Produce 	json
// @Param		id 	path int true "Unique User ID" minimum(1)
// @Success 	202 "Success"
// @Failure		400 "invalid path parameter"
// @Failure		404 "user not found"
// @Failure		500 "internal server error"
// @Router 		/v1/users/{id} [DELETE]
// @Security 	BearerAuth
func (h *UserHandler) Delete(c *gin.Context) {
	delete(c, h.service.Delete)
}

// func (h *UserHandler) GetList(c *gin.Context) {}

// GetUserByID godoc
// @Summary 	Get User By ID
// @Description Get a user by unique id
// @Tags 		users
// @Produce 	json
// @Param		id path int true "Unique User ID" minimum(1)
// @Success 	200 {object} helpers.standardResponse{data=dto.UserDetails,details=nil} "Success"
// @Failure		400 "invalid path parameter"
// @Failure		404 "user not found"
// @Failure		500 "internal server error"
// @Router 		/v1/users/{id} [GET]
// @Security 	BearerAuth
func (h *UserHandler) GetByID(c *gin.Context) {
	getByID(c, h.service.GetByID)
}

// GetUserByUsername godoc
// @Summary 	Get User By Username
// @Description Get a user by unique username
// @Tags 		users
// @Produce 	json
// @Param		username path string true "Unique User Username"
// @Success 	200 {object} helpers.standardResponse{data=dto.UserDetails,details=nil} "Success"
// @Failure		400 "invalid path parameter"
// @Failure		404 "user not found"
// @Failure		500 "internal server error"
// @Router 		/v1/users/username/{username} [GET]
// @Security 	BearerAuth
func (h *UserHandler) GetByUsername(c *gin.Context) {
	username := c.Param("username")
	if username == "" || !domain.CheckUsernamePattern(username) {
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

// GetUserByEmail godoc
// @Summary 	Get User By unique Email
// @Description Get a user by email
// @Tags 		users
// @Produce 	json
// @Param		email path string true "Unique User Email"
// @Success 	200 {object} helpers.standardResponse{data=dto.UserDetails,details=nil} "Success"
// @Failure		400 "invalid path parameter"
// @Failure		404 "user not found"
// @Failure		500 "internal server error"
// @Router 		/v1/users/email/{email} [GET]
// @Security 	BearerAuth
func (h *UserHandler) GetByEmail(c *gin.Context) {
	email := c.Param("email")
	if email == "" || !domain.CheckEmailPattern(email) {
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

// GetUserPosts godoc
// @Summary     Get User's Posts
// @Description	Get a list of user's posts with pagination
// @Tags		users
// @Produce		json
// @Param		id 		path	int	true 	"Unique User ID to use as FK"	minimum(1)
// @Param		page	query	int	false	"Page Number (default: 1)"		minimum(1)	default(1)
// @Param		size	query	int	false	"Items per Page (default: 10)"	minimum(1)	maximum(100)	default(10)
// @Success		200	{object} helpers.standardResponse{data=generics.PagedList[dto.PostsList]{items=dto.PostsList},details=nil} "Success"
// @Failure		400 "invalid path parameter | invalid query parameters (for 'size' or 'page')"
// @Failure		404 "there's no any post for this user"
// @Failure		500 "internal server error"
// @Router		/v1/users/{id}/posts [GET]
// @Security	BearerAuth
func (h *UserHandler) GetPosts(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetPosts)
}

// GetUserOwnedLists godoc
// @Summary     Get User's OwnedLists
// @Description	Get a list of user's owned-lists with pagination
// @Tags		users
// @Produce		json
// @Param		id		path	int	true 	"Unique User ID to use as FK" 	minimum(1)
// @Param		page	query	int	false	"Page Number (default: 1)"		minimum(1)	default(1)
// @Param		size	query	int	false	"Items per Page (default: 10)"	minimum(1)	maximum(100)	default(10)
// @Success		200	{object} helpers.standardResponse{data=generics.PagedList[dto.ListsList]{items=dto.ListsList},details=nil} "Success"
// @Failure		400 "invalid path parameter | invalid query parameters (for 'size' or 'page')"
// @Failure		404 "there's no any created-list for this user"
// @Failure		500 "internal server error"
// @Router		/v1/users/{id}/owned-lists [GET]
// @Security	BearerAuth
func (h *UserHandler) GetOwnedLists(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetOwnedLists)
}

// GetUserSavedLists godoc
// @Summary     Get User's SavedLists
// @Description	Get a list of user's saved-lists with pagination
// @Tags		users
// @Produce		json
// @Param		id		path	int	true 	"Unique User ID to use as FK" 	minimum(1)
// @Param		page	query	int	false	"Page Number (default: 1)"		minimum(1)	default(1)
// @Param		size	query	int	false	"Items per Page (default: 10)"	minimum(1)	maximum(100)	default(10)
// @Success		200	{object} helpers.standardResponse{data=generics.PagedList[dto.ListsList]{items=dto.ListsList},details=nil} "Success"
// @Failure		400 "invalid path parameter | invalid query parameters (for 'size' or 'page')"
// @Failure		404 "there's no any saved-list for this user"
// @Failure		500 "internal server error"
// @Router		/v1/users/{id}/saved-lists [GET]
// @Security	BearerAuth
func (h *UserHandler) GetSavedLists(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetSavedLists)
}

// GetUserAllLists godoc
// @Summary     Get User's AllLists
// @Description	Get a list of user's all-lists with pagination
// @Tags		users
// @Produce		json
// @Param		id		path	int	true 	"Unique User ID to use as FK" 	minimum(1)
// @Param		page	query	int	false	"Page Number (default: 1)"		minimum(1)	default(1)
// @Param		size	query	int	false	"Items per Page (default: 10)"	minimum(1)	maximum(100)	default(10)
// @Success		200	{object} helpers.standardResponse{data=generics.PagedList[dto.ListsList]{items=dto.ListsList},details=nil} "Success"
// @Failure		400 "invalid path parameter | invalid query parameters (for 'size' or 'page')"
// @Failure		404 "there's no any saved or created list for this user"
// @Failure		500 "internal server error"
// @Router		/v1/users/{id}/all-lists [GET]
// @Security	BearerAuth
func (h *UserHandler) GetAllLists(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetAllLists)
}

// GetUserComments godoc
// @Summary     Get User's Comments
// @Description	Get a list of user's comments with pagination
// @Tags		users
// @Produce		json
// @Param		id		path	int	true 	"Unique User ID to use as FK" 	minimum(1)
// @Param		page	query	int	false	"Page Number (default: 1)"		minimum(1)	default(1)
// @Param		size	query	int	false	"Items per Page (default: 10)"	minimum(1)	maximum(100)	default(10)
// @Success		200	{object} helpers.standardResponse{data=generics.PagedList[dto.CommentList]{items=dto.CommentList},details=nil} "Success"
// @Failure		400 "invalid path parameter | invalid query parameters (for 'size' or 'page')"
// @Failure		404 "there's no any comment for this user"
// @Failure		500 "internal server error"
// @Router		/v1/users/{id}/comments [GET]
// @Security	BearerAuth
func (h *UserHandler) GetComments(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetComments)
}

// GetUserLikes godoc
// @Summary     Get User's Likes
// @Description	Get a list of user's likes with pagination
// @Tags		users
// @Produce		json
// @Param		id		path	int	true 	"Unique User ID to use as FK" 	minimum(1)
// @Param		page	query	int	false	"Page Number (default: 1)"		minimum(1)	default(1)
// @Param		size	query	int	false	"Items per Page (default: 10)"	minimum(1)	maximum(100)	default(10)
// @Success		200	{object} helpers.standardResponse{data=generics.PagedList[dto.PostsList]{items=dto.PostsList},details=nil} "Success"
// @Failure		400 "invalid path parameter | invalid query parameters (for 'size' or 'page')"
// @Failure		404 "there's no any liked post for this user"
// @Failure		500 "internal server error"
// @Router		/v1/users/{id}/likes [GET]
// @Security	BearerAuth
func (h *UserHandler) GetLikes(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetLikes)
}

// GetUserFollowers godoc
// @Summary     Get User's Followers
// @Description	Get a list of user's followers with pagination
// @Tags		users
// @Produce		json
// @Param		id		path	int	true 	"Unique User ID to use as FK" 	minimum(1)
// @Param		page	query	int	false	"Page Number (default: 1)"		minimum(1)	default(1)
// @Param		size	query	int	false	"Items per Page (default: 10)"	minimum(1)	maximum(100)	default(10)
// @Success		200	{object} helpers.standardResponse{data=generics.PagedList[dto.UsersList]{items=dto.UsersList},details=nil} "Success"
// @Failure		400 "invalid path parameter | invalid query parameters (for 'size' or 'page')"
// @Failure		404 "there's no any follower for this user"
// @Failure		500 "internal server error"
// @Router		/v1/users/{id}/followers [GET]
// @Security	BearerAuth
func (h *UserHandler) GetFollowers(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetFollowers)
}

// GetUserFollowings godoc
// @Summary     Get User's Followings
// @Description	Get a list of user's followings with pagination
// @Tags		users
// @Produce		json
// @Param		id		path	int	true 	"Unique User ID to use as FK" 	minimum(1)
// @Param		page	query	int	false	"Page Number (default: 1)"		minimum(1)	default(1)
// @Param		size	query	int	false	"Items per Page (default: 10)"	minimum(1)	maximum(100)	default(10)
// @Success		200	{object} helpers.standardResponse{data=generics.PagedList[dto.UsersList]{items=dto.UsersList},details=nil} "Success"
// @Failure		400 "invalid path parameter | invalid query parameters (for 'size' or 'page')"
// @Failure		404 "there's no any following for this user"
// @Failure		500 "internal server error"
// @Router		/v1/users/{id}/followings [GET]
// @Security	BearerAuth
func (h *UserHandler) GetFollowings(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetFollowings)
}

// GetUserLinks godoc
// @Summary     Get User's Links
// @Description	Get a list of user's links with pagination
// @Tags		users
// @Produce		json
// @Param		id		path	int	true 	"Unique User ID to use as FK" 	minimum(1)
// @Param		page	query	int	false	"Page Number (default: 1)"		minimum(1)	default(1)
// @Param		size	query	int	false	"Items per Page (default: 10)"	minimum(1)	maximum(100)	default(10)
// @Success		200	{object} helpers.standardResponse{data=generics.PagedList[dto.LinksList]{items=dto.LinksList},details=nil} "Success"
// @Failure		400 "invalid path parameter | invalid query parameters (for 'size' or 'page')"
// @Failure		404 "there's no any link for this user"
// @Failure		500 "internal server error"
// @Router		/v1/users/{id}/links [GET]
// @Security	BearerAuth
func (h *UserHandler) GetLinks(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetLinks)
}
