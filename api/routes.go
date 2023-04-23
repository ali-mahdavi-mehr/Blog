package api

import "github.com/labstack/echo"

func RegisterAPIs(router *echo.Group) {

	router.GET("/post/", GetAllPost)
	router.GET("/post/:id/", GetOnePost)
	router.POST("/post/", CreatePost)
	router.DELETE("/post/:id/", DeletePost)

}
