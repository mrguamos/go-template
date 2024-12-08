package main

import (
	"embed"
	"fmt"
	"go-template/auth"
	"go-template/web/view"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed web/assets/*
var assets embed.FS

func main() {
	e := echo.New()
	// assetsFS := echo.MustSubFS(assets, "web/assets")
	// e.StaticFS("/assets", assetsFS)
	e.Static("/assets", "web/assets")
	e.Use(middleware.Gzip())
	e.Use(auth.AuthMiddlewareWithConfig(auth.AuthMiddlewareConfig{
		Skipper: func(c echo.Context) bool {
			fmt.Println(c.Path())
			if strings.HasPrefix(c.Path(), "/assets") ||
				strings.HasPrefix(c.Path(), "/error") ||
				strings.HasPrefix(c.Path(), "/login") ||
				c.Path() == "favicon.ico" {
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
			c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("/error?code=%d", errorCode))
			return
		}
	}
	e.GET("/error", func(c echo.Context) error {
		code := c.QueryParam("code")
		var html templ.Component
		switch code {
		case "404":
			html = view.NotFound()
		case "401":
			html = view.Unauthorized()
		case "403":
			html = view.Forbidden()
		default:
			html = view.InternalServerError()
		}
		return Render(c, http.StatusOK, html)
	})
	auth.NewAuthRoute(e).Register()
	e.Logger.Fatal(e.Start(":8080"))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
