package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	"github.com/hamidgh01/Go-Blog-API/internal/http/helpers"
	"github.com/hamidgh01/Go-Blog-API/pkg/constants"
)

func AccessControlMiddleware(
	accessibility constants.EndpointAccessibility,
	getResourceOwnerIdService func(ctx context.Context, pk uint64) (uint64, *service_errors.ServiceError),
) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint64("currentUserID")
		isSuperuser := c.GetString("currentUserIsSuperuser") // it would be 'f' (as false) or 't' (as true)

		switch accessibility {

		case constants.ADMIN_ONLY:
			if isSuperuser == "t" {
				c.Next()
				return
			} else {
				c.AbortWithStatusJSON(
					http.StatusForbidden,
					helpers.GenerateErrorResponse("permission denied", nil),
				)
				return
			}
		case constants.ADMIN_AND_OWNER:
		}

		// superuser can do anything
		if isSuperuser == "t" {
			c.Next()
			return
		}

		// get resource ID from URL
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil || id == 0 {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				helpers.GenerateErrorResponse(
					fmt.Sprintf("invalid path parameter: '%s'", c.Params.ByName("id")),
					gin.H{"description": "input must be a valid non-zero integer"},
				),
			)
			return
		}

		// get resource ownerID
		ownerID, serviceErr := getResourceOwnerIdService(c, uint64(id))
		if serviceErr != nil {
			c.AbortWithStatusJSON(
				serviceErr.Code(),
				helpers.GenerateErrorResponse(serviceErr.Message(), nil),
			)
			return
		}

		// check if current user is owner
		if ownerID != userID {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				helpers.GenerateErrorResponse("permission denied", nil),
			)
			return
		}

		c.Next()
	}
}
