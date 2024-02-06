package controller

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	dto "TradingBot/domain/dto"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func NewCustomValidator(v *validator.Validate) echo.Validator {
	cv := &CustomValidator{Validator: v}

	cv.Validator.RegisterValidation("validMarket", validMarket)
	cv.Validator.RegisterValidation("validDirection", validDirection)

	return cv
}

type BaseController struct {
	Validator echo.Validator
}

func NewBaseController(validator echo.Validator) *BaseController {
	return &BaseController{
		Validator: validator,
	}
}

func validMarket(fl validator.FieldLevel) bool {
	market, ok := fl.Field().Interface().(dto.TradeMarket)
	if !ok {
		return false
	}

	return market.IsValid()
}

func validDirection(fl validator.FieldLevel) bool {
	direction, ok := fl.Field().Interface().(dto.TradeDirection)
	if !ok {
		return false
	}

	return direction.IsValid()
}
