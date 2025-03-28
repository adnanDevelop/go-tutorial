package routes

import (
	"crud/controllers"
	"crud/middleware"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {

	e.POST("/login", controllers.LoginUser)
	e.POST("/create", controllers.CreateUser)

	protectedRoutes := e.Group("/users")
	protectedRoutes.Use(middleware.JwtMiddleware)

	protectedRoutes.GET("/all", controllers.GetUsers)
	protectedRoutes.GET("/:id", controllers.GetUserByID)
	protectedRoutes.PUT("/:id", controllers.UpdateUser)
	protectedRoutes.DELETE("/:id", controllers.DeleteUser)
}
