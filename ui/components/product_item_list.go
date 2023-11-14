package components

import (
	"github.com/ahmadyogi543/pijarcamp-crud/internal/models"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func ProductItemList(products []*models.Product) g.Node {
	return Ul(ID("product-item-list"),
		g.Group(g.Map[*models.Product](products, func(product *models.Product) g.Node {
			return Li(
				ProductItem(product),
			)
		})),
	)
}
