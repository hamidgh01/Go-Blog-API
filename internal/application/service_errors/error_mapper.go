package service_errors

import (
	"errors"
	"fmt"
	"net/http"

	dbErrors "github.com/hamidgh01/Go-Blog-API/internal/infra/database/errors"
)

func MapDBErrToServiceErr(err error, serviceName string) *ServiceError {
	switch {
	case errors.As(err, &dbErrors.UniqueViolationError{}):
		return newServiceError(http.StatusConflict, err.Error())
	case errors.As(err, &dbErrors.ForeignKeyViolationError{}):
		return newServiceError(http.StatusUnprocessableEntity, err.Error())
	case errors.As(err, &dbErrors.RecordNotFoundError{}):
		return newServiceError(http.StatusNotFound, err.Error())
	// case errors.As(err, &dbErrors.CheckViolationError{}):
	// 	return newServiceError(...)
	case errors.As(err, &dbErrors.UnexpectedDBError{}):
		fmt.Printf("failed to %s. reason: %s \n", serviceName, err.Error()) // log.Error()
		return InternalServerError
	}

	return InternalServerError
}
