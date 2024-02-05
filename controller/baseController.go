package controller

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func NewCustomValidator(v *validator.Validate) echo.Validator {
	return &CustomValidator{Validator: v}
}

type BaseController struct {
	Validator echo.Validator
}

func NewBaseController(validator echo.Validator) *BaseController {
	return &BaseController{
		Validator: validator,
	}
}
