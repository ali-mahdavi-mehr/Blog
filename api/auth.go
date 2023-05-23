package api

import (
	customError "github.com/alima12/Blog-Go/custom_error"
	"github.com/alima12/Blog-Go/database"
	"github.com/alima12/Blog-Go/models"
	"github.com/alima12/Blog-Go/utils"
	"github.com/alima12/Blog-Go/validations"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	db := database.GetDB()
	var user models.User
	db.Model(&models.User{}).Find(&user, "username = ?", username)
	if !utils.CheckPassword(password, user.Password) {
		return echo.ErrUnauthorized
	}
	userID := strconv.FormatUint(uint64(user.ID), 10)
	accessToken, refreshToken, err := utils.CreateTokens(userID)
	if err != nil {
		return echo.ErrInternalServerError
	}
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
	user.Password, _ = utils.HashPassword(user.Password)
	result := db.Create(&user)
	if result.Error != nil {
		return customError.FindDBError(result.Error, "user")
	}

	return c.JSON(http.StatusCreated, user)

}

func ChangePassword(c echo.Context) error {
	//var MyValidator *validator.Validate
	return nil

}

func Logout(c echo.Context) error {
	//var MyValidator *validator.Validate
	return nil

}
