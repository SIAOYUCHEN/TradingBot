package controller

import (
	"TradingBot/middleware"

	"github.com/labstack/echo/v4"
)

func MapUserRoutes(echo *echo.Echo, controller *UserController) {
	v1 := echo.Group("/api/v1")

	v1.POST("/login", controller.Login())

	protected := v1.Group("", middleware.AuthMiddleware)

	protected.GET("/users", controller.GetAllUsers())

	protected.GET("/user:id", controller.GetUser())

	protected.POST("/create/user", controller.CreateUser())

	protected.DELETE("/delete/user:id", controller.DeleteUser())

	protected.PUT("/users/:id/email", controller.UpdateUserEmail())
}
