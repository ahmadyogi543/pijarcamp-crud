package components

import (
	"github.com/ahmadyogi543/pijarcamp-crud/internal/models"
	g "github.com/maragudk/gomponents"
)

func ProductFormFields(product *models.Product) g.Node {
	if product != nil {
		return g.Group(
			[]g.Node{
				ProductFormField("name", "name", "Nama", "Masukan nama produk", "text", product.Name),
				ProductFormField("desc", "desc", "Deskripsi", "Masukan deskripsi produk", "text", product.Description),
				ProductFormField("price", "price", "Harga", "Masukan harga produk", "number", product.Price),
				ProductFormField("qty", "qty", "Jumlah", "Masukan jumlah produk", "number", product.Qty),
			},
		)
	}
	return g.Group(
		[]g.Node{
			ProductFormField("name", "name", "Nama", "Masukan nama produk", "text", ""),
			ProductFormField("desc", "desc", "Deskripsi", "Masukan deskripsi produk", "text", ""),
			ProductFormField("price", "price", "Harga", "Masukan harga produk", "number", ""),
			ProductFormField("qty", "qty", "Jumlah", "Masukan jumlah produk", "number", ""),
		},
	)
}
