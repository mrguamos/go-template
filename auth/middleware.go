package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Id       uuid.UUID
	MobileNo string
	Role     string
}

type Claims struct {
	User *User `json:"user"`
	jwt.RegisteredClaims
}

var DefaultBasicAuthConfig = AuthMiddlewareConfig{
	Skipper: middleware.DefaultSkipper,
}

type AuthMiddlewareConfig struct {
	Skipper middleware.Skipper
	Key     string
}

func AuthMiddlewareWithConfig(config AuthMiddlewareConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}
			cookie, err := c.Cookie("act")
			if err != nil || cookie.Value == "" {
				return echo.ErrUnauthorized
			}
			var claims Claims
			token, err := jwt.ParseWithClaims(cookie.Value, &claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(config.Key), nil
			})
			if err != nil {
				return echo.ErrUnauthorized
			}

			if !token.Valid {
				return echo.ErrUnauthorized
			}

			return next(c)
		}
	}
}

func AuthMiddleware() echo.MiddlewareFunc {
	return AuthMiddlewareWithConfig(DefaultBasicAuthConfig)
}
