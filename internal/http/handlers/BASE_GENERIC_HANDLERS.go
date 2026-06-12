package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	"github.com/hamidgh01/Go-Blog-API/internal/domain"
	"github.com/hamidgh01/Go-Blog-API/internal/http/generics"
	"github.com/hamidgh01/Go-Blog-API/internal/http/helpers"
	"github.com/hamidgh01/Go-Blog-API/internal/http/validations"

	"github.com/gin-gonic/gin"
)

func create[TRequest generics.CreateRequestTypes, TResponse generics.OutputTypes](
	c *gin.Context,
	NewCreateObjRequestDTOFunc func() *TRequest,
	createObjService func(ctx context.Context, data *TRequest) (*TResponse, *service_errors.ServiceError),
) {
	data := NewCreateObjRequestDTOFunc()
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

	objResponse, serviceErr := createObjService(c, data)
	if serviceErr != nil {
		c.AbortWithStatusJSON(
			serviceErr.Code(),
			helpers.GenerateErrorResponse(serviceErr.Message(), nil),
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		helpers.GenerateSuccessfulResponse("object created successfully.", objResponse),
	)
}

func update[TRequest generics.UpdateRequestTypes](
	c *gin.Context,
	NewUpdateObjRequestDTOFunc func() *TRequest,
	updataObjService func(ctx context.Context, pk uint64, data *TRequest) *service_errors.ServiceError,
) {
	pk, err := extractIDPathParamOrAbortWithStatusBadRequest(c)
	if err != nil {
		return
	}

	data := NewUpdateObjRequestDTOFunc()
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

	serviceErr := updataObjService(c, pk, data)
	if serviceErr != nil {
		c.AbortWithStatusJSON(
			serviceErr.Code(),
			helpers.GenerateErrorResponse(serviceErr.Message(), nil),
		)
		return
	}

	c.JSON(
		http.StatusAccepted,
		helpers.GenerateSuccessfulResponse("object updated successfully.", nil),
	)
}

func updateStatusEnum[TRequest generics.UpdateRequestTypes](
	c *gin.Context,
	updateStatusRequestDTO *TRequest,
	updateStatusService func(ctx context.Context, pk uint64, data *TRequest) *service_errors.ServiceError,
) {
	pk, err := extractIDPathParamOrAbortWithStatusBadRequest(c)
	if err != nil {
		return
	}

	serviceErr := updateStatusService(c, pk, updateStatusRequestDTO)
	if serviceErr != nil {
		c.AbortWithStatusJSON(
			serviceErr.Code(),
			helpers.GenerateErrorResponse(serviceErr.Message(), nil),
		)
		return
	}

	c.JSON(
		http.StatusAccepted,
		helpers.GenerateSuccessfulResponse("object updated successfully.", nil),
	)
}

func delete(c *gin.Context, deleteObjService func(ctx context.Context, pk uint64) *service_errors.ServiceError) {
	pk, err := extractIDPathParamOrAbortWithStatusBadRequest(c)
	if err != nil {
		return
	}

	serviceErr := deleteObjService(c, pk)
	if serviceErr != nil {
		c.AbortWithStatusJSON(
			serviceErr.Code(),
			helpers.GenerateErrorResponse(serviceErr.Message(), nil),
		)
		return
	}

	c.JSON(http.StatusAccepted, helpers.GenerateSuccessfulResponse("object deleted successfully.", nil))
}

func getByID[TResponse generics.OutputTypes](
	c *gin.Context,
	getObjByIDService func(ctx context.Context, pk uint64) (*TResponse, *service_errors.ServiceError),
) {
	pk, err := extractIDPathParamOrAbortWithStatusBadRequest(c)
	if err != nil {
		return
	}

	objResponse, serviceErr := getObjByIDService(c, pk)
	if serviceErr != nil {
		c.AbortWithStatusJSON(
			serviceErr.Code(),
			helpers.GenerateErrorResponse(serviceErr.Message(), nil),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		helpers.GenerateSuccessfulResponse("object fetched successfully.", objResponse),
	)
}

func getListOfOuterResourceByFK[TResponse generics.OutputListTypes](
	c *gin.Context,
	getOfOuterResourceByFKService func(
		ctx context.Context, fk uint64, page *domain.PaginationQueryParams,
	) (*generics.PagedList[TResponse], *service_errors.ServiceError),
) {
	pk, err := extractIDPathParamOrAbortWithStatusBadRequest(c)
	if err != nil {
		return
	}

	pageRequest := new(domain.PaginationQueryParams)
	if err := c.ShouldBindQuery(pageRequest); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			helpers.GenerateErrorResponse(
				fmt.Sprintf(
					"invalid query parameters: size='%s', page='%s'",
					c.Params.ByName("size"), c.Params.ByName("page"),
				),
				gin.H{"description": "both 'size' and 'page' must be a valid integer"},
			),
		)
		return
	}

	data, serviceErr := getOfOuterResourceByFKService(c, pk, pageRequest)
	if serviceErr != nil {
		c.AbortWithStatusJSON(
			serviceErr.Code(),
			helpers.GenerateErrorResponse(serviceErr.Message(), nil),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		helpers.GenerateSuccessfulResponse("objects fetched successfully.", data),
	)
}

func extractIDPathParamOrAbortWithStatusBadRequest(c *gin.Context) (uint64, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			helpers.GenerateErrorResponse(
				fmt.Sprintf("invalid path parameter: '%s'", c.Params.ByName("id")),
				gin.H{"description": "input must be a valid non-zero integer"},
			),
		)
		return 0, err
	}

	return uint64(id), nil
}
