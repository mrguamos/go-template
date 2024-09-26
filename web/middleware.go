package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type WebMiddlewareConfig struct {
	Skipper middleware.Skipper
}

var DefaultBasicWebConfig = WebMiddlewareConfig{
	Skipper: middleware.DefaultSkipper,
}

func WebMiddlewareWithConfig(config WebMiddlewareConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}
			err := next(c)
			return err
		}
	}
}

func WebMiddleware() echo.MiddlewareFunc {
	return WebMiddlewareWithConfig(DefaultBasicWebConfig)
}
