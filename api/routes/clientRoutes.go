package routes

import (
	"crud/controllers"
	"crud/middleware"

	"github.com/labstack/echo/v4"
)

func ClientRoutes(e *echo.Echo) {

	protectedRoutes := e.Group("/client")
	protectedRoutes.Use(middleware.JwtMiddleware)

	// protectedRoutes.GET("/all", controllers.ListClient)
	// protectedRoutes.GET("/:id", controllers.GetClientById)
	protectedRoutes.POST("/create", controllers.CreateClient)
	// protectedRoutes.PUT("/update/:id", controllers.UpdateClient)
	protectedRoutes.DELETE("/:id", controllers.DeleteClient)
}
