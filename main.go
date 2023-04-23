package main

import (
	"github.com/alima12/Blog-Go/api"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", api.Home)
	api.Register(e.Router())
	e.Logger.Fatal(e.Start(":8000"))
}
