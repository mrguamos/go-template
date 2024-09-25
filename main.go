package main

import (
	"embed"
	"fmt"
	"go-template/auth"
	"go-template/web/view"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

//go:embed web/assets/*
var assets embed.FS

func main() {
	e := echo.New()
	// assetsFS := echo.MustSubFS(assets, "web/assets")
	// e.StaticFS("/assets", assetsFS)
	e.Static("/assets", "web/assets")
	e.Use(auth.AuthMiddlewareWithConfig(auth.AuthMiddlewareConfig{
		Skipper: func(c echo.Context) bool {
			if strings.HasPrefix(c.Path(), "/assets") ||
				strings.HasPrefix(c.Path(), "/error") ||
				strings.HasPrefix(c.Path(), "/login") {
				return true
			}
			return false
		},
		Key: "s3cr3t",
	}))
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		httpError, ok := err.(*echo.HTTPError)
		if ok {
			errorCode := httpError.Code
			fmt.Println(fmt.Sprintf("/error?code=%d", errorCode))
			c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("/error?code=%d", errorCode))
		}
	}
	e.GET("/error", func(c echo.Context) error {
		code := c.QueryParam("code")
		var html string
		switch code {
		case "404":
			html = fmt.Sprint(view.NotFound())
		case "401":
			html = fmt.Sprint(view.Unauthorized())
		case "403":
			html = fmt.Sprint(view.Forbidden())
		default:
			html = fmt.Sprint(view.InternalServerError())
		}
		return c.HTML(http.StatusOK, html)
	})
	auth.NewAuthRoute(e).Register()
	e.Logger.Fatal(e.Start(":8080"))
}
