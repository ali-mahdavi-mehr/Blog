package validations

import (
	"github.com/alima12/Blog-Go/constants"
	"github.com/go-playground/validator/v10"
	"regexp"
)

type (
	UserSignUpValidation struct {
		Username string `json:"username" validate:"required,min=3,max=40"`
		Email    string `json:"e-mail" validate:"required,email"`
		Password string `json:"password" validate:"password"`
	}
)

func PasswordValidation(f validator.FieldLevel) bool {
	password := f.Field().String()
	passwordRegex := regexp.MustCompile(constants.PasswordRegex)
	result := passwordRegex.MatchString(password)
	if result {
		return true
	}
	return false
}
