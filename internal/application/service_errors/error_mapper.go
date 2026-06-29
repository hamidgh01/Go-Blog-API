package service_errors

import (
	"errors"
	"net/http"

	dbErrors "github.com/hamidgh01/Go-Blog-API/internal/infra/database/errors"
	"github.com/hamidgh01/Go-Blog-API/pkg/logging"

	"github.com/golang-jwt/jwt/v5"
)

func MapDBErrToServiceErr(err error, serviceName string) *ServiceError {
	logger := logging.GetLogger()

	switch {

	case errors.As(err, &dbErrors.UniqueViolationError{}):
		return newServiceError(http.StatusConflict, err.Error())

	case errors.As(err, &dbErrors.ForeignKeyViolationError{}):
		return newServiceError(http.StatusNotAcceptable, err.Error())

	case errors.As(err, &dbErrors.RecordNotFoundError{}):
		return newServiceError(http.StatusNotFound, err.Error())

	case errors.As(err, &dbErrors.CheckViolationError{}):
		return newServiceError(http.StatusNotAcceptable, err.Error())

	case errors.As(err, &dbErrors.BadInputError{}):
		return newServiceError(http.StatusBadRequest, err.Error())

	case errors.As(err, &dbErrors.NoRowsAffectedOnM2MEntity{}):
		return newServiceError(http.StatusNotAcceptable, err.Error())

	case errors.As(err, &dbErrors.UnexpectedDBError{}):
		logger.Errorf("failed to %s. origin: %s \n", serviceName, err.Error())
		return InternalServerError
	}

	logger.Errorf("failed to recognize error for %s. origin: %s \n", serviceName, err.Error())
	return InternalServerError
}

func MapJwtErrToServiceErr(err error) *ServiceError {
	switch err {
	case jwt.ErrTokenExpired:
		return TokenExpired
	default:
		return InvalidToken
	}
}
