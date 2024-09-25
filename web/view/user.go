package view

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func User() g.Node {
	return Page("User", Div(g.Text("User")))
}
