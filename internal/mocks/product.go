package models

import "github.com/ahmadyogi543/pijarcamp-crud/internal/models"

type MockProductModel struct {
	ID       int
	products map[int]*models.Product
}

func (m *MockProductModel) Create(name, description string, price, qty int) (int, error) {
	m.incrementID()
	product := &models.Product{
		ID:          m.ID,
		Name:        name,
		Description: description,
		Price:       price,
		Qty:         qty,
	}
	m.products[m.ID] = product
	return 0, nil
}

func (m *MockProductModel) Get(id int) (*models.Product, error) {
	product, ok := m.products[id]
	if !ok {
		return nil, models.ErrNoRecord
	}
	return product, nil
}

func (m *MockProductModel) GetAll() ([]*models.Product, error) {
	products := make([]*models.Product, 0, len(m.products))
	for _, v := range m.products {
		products = append(products, v)
	}
	return products, nil
}

func (m *MockProductModel) Update(product *models.Product) error {
	m.products[product.ID] = product
	return nil
}

func (m *MockProductModel) Delete(id int) error {
	delete(m.products, id)
	return nil
}

func (m *MockProductModel) incrementID() {
	m.ID++
}
