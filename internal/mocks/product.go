package models

import "github.com/ahmadyogi543/pijarcamp-crud/internal/models"

type MockProductModel struct {
	products []*models.Product
}

func (m *MockProductModel) Create(name, description string, price, qty int) (int, error) {
	return 0, nil
}

func (m *MockProductModel) Get(id int) (*models.Product, error) {
	return nil, nil
}

func (m *MockProductModel) GetAll() ([]*models.Product, error) {
	return nil, nil
}

func (m *MockProductModel) Update(product *models.Product) error {
	return nil
}

func (m *MockProductModel) Delete(id int) error {
	return nil
}
