package middleware

import (
	"net/http"

	"github.com/izzxx/Go-Restful-Api/utility"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		token, err := utility.ExtractToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, utility.ErrorResponse(http.StatusUnauthorized, err.Error()))
		}

		_, err = utility.ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, utility.ErrorResponse(http.StatusUnauthorized, err.Error()))
		}

		return next(c)
	}
}

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		token, err := utility.ExtractToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, utility.ErrorResponse(http.StatusUnauthorized, err.Error()))
		}

		claims, err := utility.ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, utility.ErrorResponse(http.StatusUnauthorized, err.Error()))
		}

		if claims.IsAdmin {
			return next(c)
		}

		return c.JSON(http.StatusBadRequest, utility.ErrorResponse(http.StatusBadRequest, "only admin can access this page"))
	}
}
