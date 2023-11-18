package models

type MockProductModel struct {
	products []*Product
}

func (m *MockProductModel) Create(name, description string, price, qty int) (int, error) {
	return 0, nil
}

func (m *MockProductModel) Get(id int) (*Product, error) {
	return nil, nil
}

func (m *MockProductModel) GetAll() ([]*Product, error) {
	return nil, nil
}

func (m *MockProductModel) Update(product *Product) error {
	return nil
}

func (m *MockProductModel) Delete(id int) error {
	return nil
}
