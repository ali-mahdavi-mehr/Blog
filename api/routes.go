package api

import (
	"github.com/alima12/Blog-Go/middlewares"
	"github.com/labstack/echo"
)

func RegisterAPIs(router *echo.Group) {
	//Auth
	router.POST("/auth/login/", Login)
	router.POST("/auth/sign-up/", SignUp)

	//Posts
	postRouter := router.Group("/post/")
	postRouter.GET("", GetAllPost)
	postRouter.GET(":slug/", GetOnePost, middlewares.LoginRequired)
	postRouter.POST("", CreatePost, middlewares.LoginRequired)
	postRouter.DELETE(":slug/", DeletePost, middlewares.LoginRequired)

}
