package routes

import (
	"crud/controllers"
	"crud/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	protectedRoutes := e.Group("/")
	protectedRoutes.Use(middleware.JwtMiddleware)

	e.POST("/login", controllers.LoginUser)
	e.POST("/create", controllers.CreateUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
	protectedRoutes.GET("/users/all", controllers.GetUsers)
	protectedRoutes.GET("/users/:id", controllers.GetUserByID)
}
