package components

import (
	"github.com/ahmadyogi543/pijarcamp-crud/internal/models"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func HomePage(products []*models.Product) g.Node {
	return Page(
		Div(Class("wrapper"),
			AddProductForm(),
			Div(Class("product-item-list"),
				H3(g.Text("Daftar Produk")),
				ProductItemList(products),
			),
		),
	)
}
