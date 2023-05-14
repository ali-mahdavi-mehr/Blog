package main

import (
	"github.com/alima12/Blog-Go/api"
	"github.com/alima12/Blog-Go/cmd/rpc"
	"github.com/alima12/Blog-Go/models"
	"github.com/alima12/Blog-Go/validations"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	//Load Env files
	_ = godotenv.Load(".env")

	// if you don't want to start two server, comment below code and change start point
	go rpc.StartRpcServer()
	//Register Validators
	e.Validator = &validations.CustomValidator{Validator: validations.RegisterCustomValidations()}

	//CustomErrorHandler
	//e.HTTPErrorHandler = custumerror.CustomHTTPErrorHandler

	// Database Migrations
	models.MigrateModels()

	e.GET("/", api.Home)
	// add Api endpoints
	api.RegisterAPIs(e.Group("/api"))

	//Start server
	e.Logger.Fatal(e.Start(":8000"))
}
