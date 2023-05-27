package api

import (
	"github.com/alima12/Blog-Go/middlewares"
	"github.com/labstack/echo"
)

func RegisterAPIs(router *echo.Group) {
	//Auth
	authGroup := router.Group("/auth")
	authGroup.POST("/login/", Login)
	authGroup.POST("/refresh-token/", RefreshToken)
	authGroup.POST("/sign-up/", SignUp)
	authGroup.GET("/logout/", Logout, middlewares.LoginRequired)
	authGroup.POST("/change-password/", ChangePassword, middlewares.LoginRequired)
	authGroup.POST("/delete-account/", DeleteAccount, middlewares.LoginRequired)

	//Posts
	postRouter := router.Group("/post")
	postRouter.GET("/", GetAllPost)
	postRouter.GET("/:slug/", GetOnePost, middlewares.LoginRequired)
	postRouter.POST("/", CreatePost, middlewares.LoginRequired)
	postRouter.PUT("/:slug/", UpdatePost, middlewares.LoginRequired)
	postRouter.DELETE("/:slug/", DeletePost, middlewares.LoginRequired)

	// Admin Routes
	adminRouter := router.Group("/admin")
	adminRouter.GET("/users/list/", GetAllUsers, middlewares.LoginRequired)

}
