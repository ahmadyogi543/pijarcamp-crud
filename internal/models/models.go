package models

import "errors"

var (
	ErrNoRecord       = errors.New("models: no matching record found")
	ErrRecordNotFound = errors.New("models: record not found")
)

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Qty         int    `json:"qty"`
}

type ProductModeler interface {
	Create(name, description string, price, qty int) (int, error)
	Get(id int) (*Product, error)
	GetAll() ([]*Product, error)
	Update(product *Product) error
	Delete(id int) error
}
