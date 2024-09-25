package auth

import (
	"fmt"
	"go-template/web/view"

	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthRoute struct {
	e *echo.Echo
}

func NewAuthRoute(e *echo.Echo) *AuthRoute {
	return &AuthRoute{e: e}
}

func (ar *AuthRoute) Register() {
	ar.e.GET("/login", func(c echo.Context) error {
		html := fmt.Sprint(view.Login())
		return c.HTML(http.StatusOK, html)
	})
	ar.e.POST("/login", func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, "/")
	})
}
