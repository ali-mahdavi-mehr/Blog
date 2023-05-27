package api

import (
	"github.com/alima12/Blog-Go/database"
	"github.com/alima12/Blog-Go/models"
	"github.com/labstack/echo"
)

func GetAllUsers(c echo.Context) error {
	var users []models.User
	db := database.GetDB()
	err := db.Find(&users).Error
	if err != nil {
		return err
	}
	return c.JSON(200, users)
}
