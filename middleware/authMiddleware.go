package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	dto "TradingBot/domain/dto"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := &dto.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("TradingBot"), nil
		})

		if err != nil {
			message := "Error processing request"
			if err == jwt.ErrSignatureInvalid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Signature is invalid")
			} else if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					message = "Token is malformed"
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					message = "Token is either expired or not active yet"
				} else {
					message = "Couldn't handle this token"
				}
				return echo.NewHTTPError(http.StatusBadRequest, message)
			}
			return echo.NewHTTPError(http.StatusInternalServerError, message)
		}

		if !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token is invalid")
		}

		c.Set("claims", claims)

		return next(c)
	}
}
