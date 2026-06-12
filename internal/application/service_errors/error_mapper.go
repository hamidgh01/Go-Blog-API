package service_errors

import (
	"errors"
	"net/http"

	dbErrors "github.com/hamidgh01/Go-Blog-API/internal/infra/database/errors"
)

func MapDBErrToServiceErr(err error) *ServiceError {
	switch {
	case errors.Is(err, &dbErrors.UniqueViolationError{}):
		return newServiceError(http.StatusConflict, err.Error())
	case errors.Is(err, &dbErrors.ForeignKeyViolationError{}):
		return newServiceError(http.StatusUnprocessableEntity, err.Error())
	case errors.Is(err, &dbErrors.RecordNotFoundError{}):
		return newServiceError(http.StatusNotFound, err.Error())
	// case errors.As(err, &dbErrors.CheckViolationError{}):
	// 	return newServiceError(...)
	case errors.Is(err, &dbErrors.UnexpectedDBError{}):
		// log.Error(err.Error())
		return InternalServerError
	}

	return InternalServerError
}
