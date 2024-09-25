package view

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

func Login() g.Node {
	return c.HTML5(c.HTML5Props{
		Title:    "Login",
		Language: "en",
		Head: []g.Node{
			Link(g.Attr("rel", "stylesheet"), g.Attr("href", "/assets/css/output.css")),
		},
		Body: []g.Node{Class(""),
			Div(Class("center-layout"),
				Div(Class("max-w-md w-full"),

					Form(Class("card pt-6"), g.Attr("action", "#"), g.Attr("method", "POST"),
						Div(Class("card-header"),
							H3(g.Text("Sign in to your account")),
							P(g.Text("Don't have an account? "), A(Class("text-primary hover:underline"), g.Attr("href", "/register"), g.Text("Register"))),
						),
						Div(Class("card-content"),
							Div(
								Label(g.Text("Email address")),
								Div(Class("mt-1"),
									Input(Type("email"), Name("email"), ID("email"), g.Attr("autocomplete", "email"), g.Attr("required", "true"),
										Class("")),
								),
							),
							Div(
								Label(g.Text("Password")),
								Div(Class("mt-1"),
									Input(Type("password"), Name("password"), ID("password"), g.Attr("autocomplete", "current-password"), g.Attr("required", "true"),
										Class("")),
								),
							),
						),
						Div(Class("card-footer"),
							Button(Type("submit"),
								Class("w-full btn-primary"),
								g.Text("Sign in"),
							),
							Button(Type("submit"),
								Class("w-full btn-ghost"),
								g.Text("Forgot Password?"),
							),
						),
					),
				),
			),
		},
	})
}
