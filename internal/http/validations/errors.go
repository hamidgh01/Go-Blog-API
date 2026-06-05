package validations

import (
	"errors"
	"fmt"

	"github.com/hamidgh01/Go-Blog-API/internal/domain"

	"github.com/go-playground/validator/v10"
)

type TranslatedValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func translateValidationErrorToReadableMessage(err validator.FieldError) TranslatedValidationError {
	field := err.Field()
	var msg string

	switch err.Tag() {

	case "required":
		msg = fmt.Sprintf("%s is required.", field)

	case "eqfield":
		msg = fmt.Sprintf("%s and %s don't match.", field, err.Param())
	case "nefield":
		msg = fmt.Sprintf("%s and %s should not be same.", field, err.Param())

	case "eq":
		msg = fmt.Sprintf("%s should be equal to: %s.", field, err.Param())

	// use these for length
	case "min":
		msg = fmt.Sprintf("%s is too short (min=%s).", field, err.Param())
	case "max": // use this for length
		msg = fmt.Sprintf("'%s' is too long (max=%s).", field, err.Param())

	// use these for numbers
	case "gt":
		msg = fmt.Sprintf("'%s' must be greater than %s (%s > %s).", field, err.Param(), field, err.Param())
	case "gte":
		msg = fmt.Sprintf("'%s' must be greater than or equal %s (%s >= %s).", field, err.Param(), field, err.Param())
	case "lt":
		msg = fmt.Sprintf("'%s' must be less than %s (%s < %s).", field, err.Param(), field, err.Param())
	case "lte":
		msg = fmt.Sprintf("'%s' must be less than or equal %s (%s <= %s).", field, err.Param(), field, err.Param())

	case "email":
		msg = fmt.Sprintf("'%s' is not a valid email address.", err.Value())
	case "url":
		msg = fmt.Sprintf("'%s' is not a valid url.", err.Value())
	case "strong_password":
		msg = fmt.Sprintf("'%s' is not a valid password. %s", err.Value(), domain.PasswordPatternDescription)
	case "username_pattern":
		msg = fmt.Sprintf("'%s' is not a valid username. %s", err.Value(), domain.UsernamePatternDescription)
	case "tag_pattern":
		msg = fmt.Sprintf("'%s' is not a valid tag. %s", err.Value(), domain.TagPatternDescription)

	case "oneof":
		msg = fmt.Sprintf("'%s' is not a valid input for %s. input should be one of: %s", err.Value(), field, err.Param())

	default:
		msg = fmt.Sprintf("there is error at field '%s'.", field)
	}

	return TranslatedValidationError{Field: field, Message: msg}
}

func GetValidationErrors(err error) *[]TranslatedValidationError {
	var result []TranslatedValidationError
	if errors.As(err, &validator.ValidationErrors{}) {
		for _, err := range err.(validator.ValidationErrors) {
			translatedVldErr := translateValidationErrorToReadableMessage(err)
			result = append(result, translatedVldErr)
		}

		return &result
	}

	return nil
}
