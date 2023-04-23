package api

import (
	"github.com/alima12/Blog-Go/utils"
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
