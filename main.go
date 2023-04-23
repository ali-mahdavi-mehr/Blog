package main

import (
	"github.com/alima12/Blog-Go/api"
	"github.com/alima12/Blog-Go/models"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	_ = godotenv.Load(".env")
	models.MigrateModels()
	e := echo.New()
	e.GET("/", api.Home)
	api.RegisterAPIs(e.Group("/api"))
	e.Logger.Fatal(e.Start(":8000"))
}
