package components

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func AddProductForm() g.Node {
	return FormEl(ID("product-form"), Class("product-form"),
		g.Attr("hx-post", "/products"),
		g.Attr("hx-target", "#product-item-list"),
		g.Attr("hx-swap", "outerHTML"),
		g.Attr("hx-on::after-request", "if(event.detail.successful) this.reset()"),
		H3(g.Text("Tambah Produk")),
		ProductFormFields(nil),
		Button(Class("button"),
			Type("submit"),
			g.Text("Add"),
		),
	)
}
