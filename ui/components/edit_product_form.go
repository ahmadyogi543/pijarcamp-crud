package components

import (
	"fmt"

	"github.com/ahmadyogi543/pijarcamp-crud/internal/models"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func EditProductForm(product *models.Product) g.Node {
	return FormEl(ID("product-form"), Class("product-form"),
		H3(g.Text("Edit Produk")),
		ProductFormFields(product),
		Div(Class("button-wrapper"),
			Button(Class("button"),
				g.Attr("hx-put", fmt.Sprintf("/products/edit/%d", product.ID)),
				g.Attr("hx-target", "#product-item-list"),
				g.Attr("hx-swap", "outerHTML"),
				g.Attr("hx-confirm", "Are you sure want to update the item?"),
				Type("submit"),
				g.Text("Update"),
			),
			Button(Class("button"),
				g.Attr("hx-get", "/products/create"),
				g.Attr("hx-target", "#product-form"),
				g.Attr("hx-swap", "outerHTML"),
				g.Text("Back"),
			),
		),
	)
}
