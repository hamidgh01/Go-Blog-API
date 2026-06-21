package handlers

import (
	"github.com/hamidgh01/Go-Blog-API/internal/application/services"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
	_ "github.com/hamidgh01/Go-Blog-API/internal/http/generics"
	_ "github.com/hamidgh01/Go-Blog-API/internal/http/helpers"

	"github.com/gin-gonic/gin"
)

type ListHandler struct {
	service *services.ListService
}

func NewListHandler(s *services.ListService) *ListHandler {
	return &ListHandler{service: s}
}

// Create List godoc
// @Summary 	Create List
// @Description Create a new list
// @Tags 		lists
// @Accept  	json
// @Produce 	json
// @Param 		request body dto.CreateListRequest true "Data to create new list"
// @Success		201 {object} helpers.standardResponse{data=dto.ListDetails,details=nil} "Success"
// @Failure		400 "validation error"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		500 "internal server error"
// @Router 		/v1/lists [POST]
// @Security BearerAuth
func (h *ListHandler) Create(c *gin.Context) {
	create(c, dto.NewCreateListRequest, h.service.Create)
}

// Updata List godoc
// @Summary 	Update List
// @Description Update a list
// @Tags 		lists
// @Accept  	json
// @Produce 	json
// @Param		id		path int					true "Unique List ID" minimum(1)
// @Param 		request body dto.UpdateListRequest	true "Data to update a list"
// @Success 	202 "Success"
// @Failure		400 "validation error | invalid path parameter"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		404 "list not found"
// @Failure		500 "internal server error"
// @Router 		/v1/lists/{id} [PUT]
// @Security 	BearerAuth
func (h *ListHandler) Update(c *gin.Context) {
	update(c, dto.NewUpdateListRequest, h.service.Update)
}

// Updata List.IsPrivate godoc
// @Summary 	Updata List.IsPrivate
// @Description Update list's privacy
// @Tags 		lists
// @Accept  	json
// @Produce 	json
// @Param		id		path int                      		true "Unique List ID" minimum(1)
// @Param 		request	body dto.UpdateListPrivacyRequest 	true "new privacy value to update List.IsPrivate"
// @Success 	202 "Success"
// @Failure		400 "validation error | invalid path parameter"
// @Failure		422 "invalid json format or json fields with invalid value type"
// @Failure		404 "list not found"
// @Failure		500 "internal server error"
// @Router 		/v1/lists/{id} [PATCH]
// @Security 	BearerAuth
func (h *ListHandler) UpdatePrivacy(c *gin.Context) {
	update(c, dto.NewUpdateListPrivacyRequest, h.service.UpdatePrivacy)
}

// Delete List godoc
// @Summary 	Delete List
// @Description Delete a list
// @Tags 		lists
// @Produce 	json
// @Param		id 	path int true "Unique List ID" minimum(1)
// @Success 	202 "Success"
// @Failure		400 "invalid path parameter"
// @Failure		404 "list not found"
// @Failure		500 "internal server error"
// @Router 		/v1/lists/{id} [DELETE]
// @Security 	BearerAuth
func (h *ListHandler) Delete(c *gin.Context) {
	delete(c, h.service.Delete)
}

// GetListByID godoc
// @Summary 	Get List By ID
// @Description Get a list by unique id
// @Tags 		lists
// @Produce 	json
// @Param		id path int true "Unique List ID" minimum(1)
// @Success 	200 {object} helpers.standardResponse{data=dto.ListDetails,details=nil} "Success"
// @Failure		400 "invalid path parameter"
// @Failure		404 "list not found"
// @Failure		500 "internal server error"
// @Router 		/v1/lists/{id} [GET]
// @Security 	BearerAuth
func (h *ListHandler) GetByID(c *gin.Context) {
	getByID(c, h.service.GetByID)
}

// -----------------------------------------------------------------------------
// other sources that has FK to `List`

// GetListSavedPosts godoc
// @Summary     Get List's Saved Posts
// @Description	Get list of saved posts of the list with pagination
// @Tags		lists
// @Produce		json
// @Param		id 		path	int	true 	"Unique List ID to use as FK"	minimum(1)
// @Param		page	query	int	false	"Page Number (default: 1)"		minimum(1)	default(1)
// @Param		size	query	int	false	"Items per Page (default: 10)"	minimum(1)	maximum(100)	default(10)
// @Success		200	{object} helpers.standardResponse{data=generics.PagedList[dto.PostsList]{items=dto.PostsList},details=nil} "Success"
// @Failure		400 "invalid path parameter | invalid query parameters (for 'size' or 'page')"
// @Failure		404 "there's no any saved-post in this list"
// @Failure		500 "internal server error"
// @Router		/v1/lists/{id}/saved-posts [GET]
// @Security	BearerAuth
func (h *ListHandler) GetSavedPosts(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetSavedPosts)
}

// GetUsersWhoSaved godoc
// @Summary     Get Users Who Saved The List
// @Description	Get list of users who saved the list with pagination
// @Tags		lists
// @Produce		json
// @Param		id 		path	int	true 	"Unique List ID to use as FK"	minimum(1)
// @Param		page	query	int	false	"Page Number (default: 1)"		minimum(1)	default(1)
// @Param		size	query	int	false	"Items per Page (default: 10)"	minimum(1)	maximum(100)	default(10)
// @Success		200	{object} helpers.standardResponse{data=generics.PagedList[dto.UsersList]{items=dto.UsersList},details=nil} "Success"
// @Failure		400 "invalid path parameter | invalid query parameters (for 'size' or 'page')"
// @Failure		404 "there's no any user who saves this list"
// @Failure		500 "internal server error"
// @Router		/v1/lists/{id}/users-who-saved [GET]
// @Security	BearerAuth
func (h *ListHandler) GetUsersWhoSaved(c *gin.Context) {
	getListOfOuterResourceByFK(c, h.service.GetUsersWhoSaved)
}
