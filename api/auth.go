package api

import (
	"github.com/alima12/Blog-Go/database"
	"github.com/alima12/Blog-Go/models"
	"github.com/alima12/Blog-Go/utils"
	"github.com/alima12/Blog-Go/validations"
	"github.com/jackc/pgx/v5/pgconn"
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
	accessToken := utils.CreateToken("access", userID)
	refreshToken := utils.CreateToken("refresh", userID)
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
		err := result.Error.(*pgconn.PgError)
		switch {
		case err.Code == "23505":
			return echo.NewHTTPError(http.StatusConflict, "user already exists!")
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, result.Error.Error())

		}
	}

	return c.JSON(http.StatusCreated, user)

}
