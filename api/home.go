package api

import (
	"github.com/alima12/Blog-Go/models"
	"github.com/labstack/echo"
	"net/http"
)

func Home(c echo.Context) error {
	p := models.Post{
		Title: "First Post",
		Author: models.User{
			FirstName: "Ali",
			LastName:  "Mahdavi",
			Username:  "alima12",
		},
		Body: "this is my first post",
	}
	return c.JSON(http.StatusOK, p)
}
