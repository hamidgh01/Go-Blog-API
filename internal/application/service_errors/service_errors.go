package service_errors

import "net/http"

type ServiceError struct {
	code    int
	message string
}

var _ error = (*ServiceError)(nil)

func newServiceError(httpCode int, message string) *ServiceError {
	return &ServiceError{code: httpCode, message: message}
}

func (s *ServiceError) Code() int {
	return s.code
}

func (s *ServiceError) Message() string {
	return s.message
}

func (s *ServiceError) Error() string {
	return s.message
}

var (
	// User
	InvalidCredentials = newServiceError(http.StatusUnauthorized, "invalid credentials")
	PermissionDenied   = newServiceError(http.StatusForbidden, "permission denied")

	// server
	InternalServerError = newServiceError(http.StatusInternalServerError, "internal server error")
)
