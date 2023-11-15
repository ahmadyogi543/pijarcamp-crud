package components

import (
	"fmt"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func ProductFormField(id, name, title, placeholder, inputType string, value any) g.Node {
	return Div(Class("product-form-field"),
		Label(
			For(id),
			g.Text(title),
		),
		Input(
			ID(id),
			Name(name),
			Placeholder(placeholder),
			Type(inputType),
			Value(fmt.Sprint(value)),
		),
	)
}
