package models

type ProductModeler interface {
	Create(name, description string, price, qty int) (int, error)
	Get(id int) (*Product, error)
	GetAll() ([]*Product, error)
	Update(product *Product) error
	Delete(id int) error
}
