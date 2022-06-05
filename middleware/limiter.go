package middleware

import (
	"net/http"
	"time"

	"github.com/izzxx/Go-Restful-Api/utility"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Limiter() echo.MiddlewareFunc {
	limit := middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{
				Rate:      10,
				Burst:     20,
				ExpiresIn: time.Minute,
			},
		),
		IdentifierExtractor: func(context echo.Context) (string, error) {
			ip := context.RealIP()
			return ip, nil
		},
		DenyHandler: func(c echo.Context, identifier string, err error) error {
			return c.JSON(http.StatusTooManyRequests, utility.ErrorResponse(http.StatusTooManyRequests, "to many request"))
		},
	})

	return limit
}
