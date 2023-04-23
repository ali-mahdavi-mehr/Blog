package api

import (
	"github.com/labstack/echo"
	"net/http"
)

func GetAllPost(c echo.Context) error {
	return c.String(http.StatusOK, "ok")

}

func GetOnePost(c echo.Context) error {
	return c.String(http.StatusOK, "one post")
}

func CreatePost(c echo.Context) error {
	return c.String(http.StatusOK, "create Post")
}

func DeletePost(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
