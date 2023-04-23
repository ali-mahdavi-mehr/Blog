package api

import "github.com/labstack/echo"

func Register(router *echo.Router) {
	router.Add(echo.GET, "/api/post/", GetAllPost)
	router.Add(echo.GET, "/api/post/:id/", GetOnePost)
	router.Add(echo.POST, "/api/post/", CreatePost)
	router.Add(echo.DELETE, "/api/post/:id/", DeletePost)

}
