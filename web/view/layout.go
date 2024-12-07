package view

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

func Page(title string, body g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:    "Template - " + title,
		Language: "en",
		Head: []g.Node{
			Link(g.Attr("rel", "stylesheet"), g.Attr("href", "/assets/css/output.css")),
		},
		Body: []g.Node{
			Div(Class("flex min-h-screen"),
				Sidebar(),
				// Main content area with transition
				Div(Class("flex-1 ml-64 transition-all duration-300 main-content"),
					// Regular content
					Div(Class("p-8"),
						body,
					),
				),
			),
		},
	})
}

func ErrorPage(title string, body g.Node, path, btnText string) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:    "Template - " + title,
		Language: "en",
		Head: []g.Node{
			Link(g.Attr("rel", "stylesheet"), g.Attr("href", "/assets/css/output.css")),
			Link(g.Attr("rel", "icon"), g.Attr("href", "/assets/favicon.svg")),
		},
		Body: []g.Node{
			Div(Class("flex min-h-screen"),
				Sidebar(),
				// Main content area with transition
				Div(Class("flex-1 ml-64 transition-all duration-300 main-content"),
					Div(Class("min-h-screen flex items-center justify-center p-8"),
						Div(Class("max-w-md text-center"),
							SVG(g.Attr("xmlns", "http://www.w3.org/2000/svg"), Width("24"), Height("24"), g.Attr("viewBox", "0 0 24 24"), g.Attr("fill", "none"), g.Attr("stroke", "currentColor"), g.Attr("stroke-width", "2"), g.Attr("stroke-linecap", "round"), g.Attr("stroke-linejoin", "round"), Class("mx-auto h-12 w-12 text-primary"),
								g.El("path", g.Attr("d", "m21.73 18-8-14a2 2 0 0 0-3.48 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3")),
								g.El("path", g.Attr("d", "M12 9v4")),
								g.El("path", g.Attr("d", "M12 17h.01")),
							),
							H1(Class("mt-4 text-6xl font-bold tracking-tight text-foreground"), g.Text(title)),
							P(Class("mt-4 text-muted-foreground"), body),
							Div(Class("mt-6"),
								A(Class("btn-primary"), Href(path), g.Text(btnText)),
							),
						),
					),
				),
			),
		},
	})
}
