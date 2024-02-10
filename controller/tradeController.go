package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"

	domain "TradingBot/domain/dto"
	createTrade "TradingBot/domain/trade/createTrade"
	deleteTrade "TradingBot/domain/trade/deleteTrade"
	getAllTrade "TradingBot/domain/trade/getAllTrade"
	getTrade "TradingBot/domain/trade/getTrade"
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

// @Summary Create Trade
// @Description Creates a new trade with market, price, amount, and direction
// @Tags Trade
// @Accept json
// @Produce json
// @Param createTradeCommand body createTrade.CreateTradeCommand true "Create Trade Command" example({"market": "ETH/USDT", "price": 100, "amount": 1, "direction": "Ask"})
// @Success 200 {object} createTrade.CreateTradeResponse "Trade created successfully"
// @Failure 400 {string} string "Bad request - invalid input"
// @Failure 401 {string} string "Unauthorized - invalid credentials"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/create/trade [post]
// @Security ApiKeyAuth
func (uc *TradeController) CreateTrade() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := &createTrade.CreateTradeCommand{}
		if err := ctx.Bind(request); err != nil {
			return ctx.JSON(http.StatusBadRequest, "Bad request")
		}

		if err := ctx.Validate(request); err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		result, err := mediatr.Send[*createTrade.CreateTradeCommand, *createTrade.CreateTradeResponse](ctx.Request().Context(), request)
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, err.Error())
		}

		return ctx.JSON(http.StatusOK, result)
	}
}

// GetTrade gets trade details based on market and direction
// @Summary Get trade details
// @Description Retrieve trades based on market and direction
// @Tags Trade
// @Accept json
// @Produce json
// @Param market path string true "Market Identifier" Enums(Eth, Btc, Flow, Sol)
// @Param direction path string true "Trade Direction" Enums(Ask, Bid)
// @Success 200 {array} getTrade.GetTradeResponse "List of trades fetched successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/trade/{market}/{direction} [get]
// @Security ApiKeyAuth
func (uc *TradeController) GetTrade() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		marketStr := ctx.Param("market")
		directionStr := ctx.Param("direction")

		market := domain.TradeMarket(marketStr)
		direction := domain.TradeDirection(directionStr)

		query := getTrade.GetTradeQuery{Market: market, Direction: direction}

		if err := ctx.Validate(&query); err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		response, err := mediatr.Send[*getTrade.GetTradeQuery, *getTrade.GetTradeResponse](ctx.Request().Context(), &query)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, response)
	}
}

// GetAllTrade retrieves all trades
// @Summary Retrieve all trades
// @Description Retrieves a map of all trades grouped by their market and direction from the database
// @Tags Trade
// @Accept json
// @Produce json
// @Success 200 {object} map[string][]domain.Trade "Map of all trades grouped by market and direction"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/all/trades [get]
// @Security ApiKeyAuth
func (uc *TradeController) GetAllTrade() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		query := &getAllTrade.GetAllTradeQuery{}

		response, err := mediatr.Send[*getAllTrade.GetAllTradeQuery, *getAllTrade.GetAllTradeResponse](ctx.Request().Context(), query)

		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, response)
	}
}

// DeleteTrade deletes a trade by market and direction
// @Summary Delete trade
// @Description Deletes a trade with the specified market and direction from the database
// @Tags Trade
// @Accept json
// @Produce json
// @Param market path string true "Market Identifier"
// @Param direction path string true "Trade Direction"
// @Success 200 {object} deleteTrade.DeleteTradeResponse "Trade deleted successfully"
// @Failure 400 {string} string "Bad request - invalid input"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/delete/trade/{market}/{direction} [delete]
// @Security ApiKeyAuth
func (uc *TradeController) DeleteTrade() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		marketStr := ctx.Param("market")
		directionStr := ctx.Param("direction")

		market := domain.TradeMarket(marketStr)
		direction := domain.TradeDirection(directionStr)

		command := deleteTrade.DeleteTradeCommand{Market: market, Direction: direction}

		if err := ctx.Validate(&command); err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		response, err := mediatr.Send[*deleteTrade.DeleteTradeCommand, *deleteTrade.DeleteTradeResponse](ctx.Request().Context(), &command)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, response)
	}
}
