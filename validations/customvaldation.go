package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
)

type (
	CustomValidator struct {
		Validator *validator.Validate
	}
)

func RegisterCustomValidations() *validator.Validate {
	var Validator = validator.New()
	_ = Validator.RegisterValidation("password", PasswordValidation)
	return Validator
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
