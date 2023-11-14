package components

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Navbar() g.Node {
	return Nav(Class("navbar"),
		Img(Class("navbar-icon"),
			Src("/static/images/pijar-camp.png"),
			Alt("pijar-camp-logo"),
		),
		H1(g.Text("PijarCamp CRUD App")),
	)
}
