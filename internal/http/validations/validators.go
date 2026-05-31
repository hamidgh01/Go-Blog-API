package validations

import (
	"errors"

	"Go-Blog-API/internal/domain"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidators() error {
	validator, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return errors.New("validator engine is not of type `*validator.Validate`")
	}

	if err := validator.RegisterValidation("username_pattern", usernameValidator); err != nil {
		return err
	}
	if err := validator.RegisterValidation("strong_password", passwordValidator); err != nil {
		return err
	}
	if err := validator.RegisterValidation("tag_pattern", tagValidator); err != nil {
		return err
	}

	return nil
}

func usernameValidator(f validator.FieldLevel) bool {
	value, ok := f.Field().Interface().(string)
	if !ok {
		return false
	}

	return domain.CheckUsernamePattern(value)
}

func passwordValidator(f validator.FieldLevel) bool {
	value, ok := f.Field().Interface().(string)
	if !ok {
		return false
	}

	return domain.CheckPasswordPattern(value)
}

func tagValidator(f validator.FieldLevel) bool {
	value, ok := f.Field().Interface().(string)
	if !ok {
		return false
	}

	return domain.CheckTagPattern(value)
}
