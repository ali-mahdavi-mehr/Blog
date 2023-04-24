package api

import (
	"github.com/labstack/echo"
	"net/http"
)

func Home(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
