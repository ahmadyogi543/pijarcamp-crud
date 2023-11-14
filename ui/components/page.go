package components

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

func Page(children ...g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:    "PijarCamp CRUD App",
		Language: "id",
		Head: []g.Node{
			Link(
				Rel("stylesheet"),
				Href("/static/css/style.css"),
			),
			Script(Src("/static/js/htmx.min.js")),
		},
		Body: []g.Node{
			Header(Navbar()),
			Main(children...),
		},
	})
}
