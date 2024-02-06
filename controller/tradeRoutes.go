package controller

import (
	"TradingBot/middleware"

	"github.com/labstack/echo/v4"
)

func MapTradeRoutes(echo *echo.Echo, controller *TradeController) {
	v1 := echo.Group("/api/v1")

	protected := v1.Group("", middleware.AuthMiddleware)

	protected.POST("/create/trade", controller.CreateTrade())

	protected.GET("/trade/:market/:direction", controller.GetTrade())

	protected.GET("/all/trades", controller.GetAllTrade())

	protected.DELETE("/delete/trade/:market/:direction", controller.DeleteTrade())
}
