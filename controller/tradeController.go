package controller

import (
	"github.com/labstack/echo/v4"
)

type TradeController struct {
	*BaseController
	echo *echo.Echo
}

func NewTradeController(echo *echo.Echo, baseController *BaseController) *TradeController {
	return &TradeController{
		BaseController: baseController,
		echo:           echo,
	}
}
