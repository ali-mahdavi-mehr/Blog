package api

import (
	"github.com/alima12/Blog-Go/database"
	"github.com/alima12/Blog-Go/models"
	"github.com/alima12/Blog-Go/validations"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
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
	var post models.Post
	err := post.GetOne(c.Param("id"))
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, post)
}

func CreatePost(c echo.Context) error {
	// validate data
	data := new(validations.PostValidation)
	if err := c.Bind(data); err != nil {
		return echo.ErrBadRequest
	}
	if err := validator.New().Struct(data); err != nil {
		return echo.ErrBadRequest
	}

	// binding data
	var post models.Post
	_ = c.Bind(&post)

	// Find Author
	userID := c.Request().Header.Get("user_id")
	id, _ := strconv.ParseInt(userID, 10, 32)
	post.UserID = uint(id)

	// Save Post
	post.Create()
	return c.JSON(http.StatusCreated, post)
}

func DeletePost(c echo.Context) error {
	var post models.Post
	err := post.Delete(c.Param("id"))
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	} else {
		return c.NoContent(http.StatusNoContent)
	}
}
