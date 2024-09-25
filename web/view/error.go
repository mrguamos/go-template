package view

import (
	g "github.com/maragudk/gomponents"
)

func NotFound() g.Node {
	return ErrorPage("404", g.Text("Oops, the page you're looking for doesn't exist."), "/", "Go to Homepage")
}

func Unauthorized() g.Node {
	return ErrorPage("401", g.Text(`Oops, it looks like you're not currently logged in.`), "/login", "Go to Login Page")
}

func Forbidden() g.Node {
	return ErrorPage("403", g.Text("Oops, it looks like you don't have permission to access this page."), "/", "Go to Homepage")
}

func InternalServerError() g.Node {
	return ErrorPage("500", g.Text(`We apologize for the inconvenience, but an unexpected error has occurred on our end. 
		Please try again later or contact our support team if the issue persists.`), "/", "Go to Homepage")
}
