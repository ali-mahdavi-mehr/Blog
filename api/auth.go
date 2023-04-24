package api

import (
	"github.com/alima12/Blog-Go/database"
	"github.com/alima12/Blog-Go/models"
	"github.com/alima12/Blog-Go/utils"
	"github.com/alima12/Blog-Go/validations"
	"github.com/labstack/echo"
	"net/http"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username == "" || password == "" {
		return echo.ErrUnauthorized
	}
	accessToken := utils.CreateToken("access", username)
	refreshToken := utils.CreateToken("refresh", username)
	return c.JSON(http.StatusOK, echo.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func SignUp(c echo.Context) error {
	//var MyValidator *validator.Validate
	var data validations.UserSignUpValidation
	if err := c.Bind(&data); err != nil {
		return echo.ErrBadRequest
	}
	if err := c.Validate(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var user models.User
	_ = c.Bind(&user)

	db := database.GetDB()
	result := db.Create(&user)
	if result.Error != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusCreated, user)

}
