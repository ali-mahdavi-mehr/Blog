package api

import (
	"github.com/alima12/Blog-Go/middlewares"
	"github.com/labstack/echo"
)

func RegisterAPIs(router *echo.Group) {
	//Auth
	router.POST("/auth/login/", Login)

	//Posts
	postRouter := router.Group("/post/")
	postRouter.GET("", GetAllPost)
	postRouter.GET(":id/", GetOnePost, middlewares.LoginRequired)
	postRouter.POST("", CreatePost, middlewares.LoginRequired)
	postRouter.DELETE(":id/", DeletePost, middlewares.LoginRequired)

}
