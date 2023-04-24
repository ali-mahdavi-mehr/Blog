package api

import (
	"github.com/alima12/Blog-Go/database"
	"github.com/alima12/Blog-Go/models"
	"github.com/labstack/echo"
	"net/http"
)

func GetAllPost(c echo.Context) error {
	db := database.GetDB()

	var posts []models.Post
	err := db.Model(&models.Post{}).Find(&posts).Error
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, posts)

}

func GetOnePost(c echo.Context) error {
	db := database.GetDB()
	id := c.Param("id")
	var post models.Post
	result := db.First(&post, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "post not found")
	}
	return c.JSON(http.StatusOK, post)
}

func CreatePost(c echo.Context) error {
	return c.String(http.StatusOK, "create Post")
}

func DeletePost(c echo.Context) error {
	return c.String(http.StatusNoContent, "ok")
}
