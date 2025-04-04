package main

import (
	"crud/config"
	"crud/controllers" // Import controllers
	"crud/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	client := config.ConnectDB()
	controllers.Init(client.Database("testdb"))
	controllers.InitClient(client.Database("testdb"))

	routes.UserRoutes(e)
	routes.ClientRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
