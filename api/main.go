package main

import (
	"crud/config"
	"crud/controllers"
	"crud/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	} else {
		log.Println("✅ .env file loaded")
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	client := config.ConnectDB()
	dbName := os.Getenv("MONGO_DB")

	controllers.Init(client.Database(dbName))
	controllers.InitClient(client.Database(dbName))

	routes.UserRoutes(e)
	routes.ClientRoutes(e)

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
