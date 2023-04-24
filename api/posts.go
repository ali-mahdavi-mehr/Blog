package api

import (
	"github.com/alima12/Blog-Go/database"
	"github.com/alima12/Blog-Go/models"
	"github.com/alima12/Blog-Go/validation"
	"github.com/go-playground/validator/v10"
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
	// validate data
	data := new(validation.PostValidation)
	if err := c.Bind(data); err != nil {
		return echo.ErrBadRequest
	}
	if err := validator.New().Struct(data); err != nil {
		return echo.ErrBadRequest
	}

	var post models.Post
	if err := c.Bind(&post); err != nil {
		return echo.ErrBadRequest
	}

	// Find Author
	db := database.GetDB()
	username := c.Request().Header.Get("user_id")
	var user models.User
	db.Model(&models.User{}).Find(&user, "username = ?", username)
	post.UserID = user.ID

	// Save Post
	db.Create(&post)
	return c.JSON(http.StatusCreated, post)
}

func DeletePost(c echo.Context) error {
	db := database.GetDB()
	id := c.Param("id")
	result := db.Delete(&models.Post{}, "id = ?", id)
	if result.RowsAffected == 1 {
		return c.NoContent(http.StatusNoContent)
	}
	return c.String(http.StatusNotFound, "post not found!")
}
