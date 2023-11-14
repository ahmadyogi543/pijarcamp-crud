package components

import (
	"fmt"

	"github.com/ahmadyogi543/pijarcamp-crud/internal/models"
	"github.com/ahmadyogi543/pijarcamp-crud/internal/utils"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func ProductItem(product *models.Product) g.Node {
	return Div(Class("product-item"),
		H4(g.Text(product.Name)),
		P(g.Text(product.Description)),
		Div(Class("product-price-qty-tag"),
			Small(g.Text(fmt.Sprintf("Rp. %s", utils.FormatCurrency(product.Price)))),
			Small(g.Text(fmt.Sprintf("Jumlah: %d item", product.Qty))),
		),
		Div(Class("button-wrapper"),
			Button(
				g.Attr("hx-get", fmt.Sprintf("/edit/%d", product.ID)),
				g.Attr("hx-target", "#product-form"),
				g.Attr("hx-swap", "outerHTML"),
				g.Text("Edit"),
			),
			Button(
				g.Attr("hx-delete", fmt.Sprintf("/delete/%d", product.ID)),
				g.Attr("hx-target", "#product-item-list"),
				g.Attr("hx-confirm", "Are you sure to delete this item?"),
				g.Attr("hx-swap", "outerHTML"),
				g.Text("Delete"),
			),
		),
	)
}
