package middlewares

import (
	"github.com/alima12/Blog-Go/utils"
	"github.com/labstack/echo"
	"net/http"
)

func LoginRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ok, userID, err := utils.IsValidToken(c.Request().Header.Get("Authorization"))
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		if ok {
			c.Request().Header.Set("user_id", userID)
			return next(c)
		}
		return echo.ErrUnauthorized

	}
}
