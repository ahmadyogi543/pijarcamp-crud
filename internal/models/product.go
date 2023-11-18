package models

import (
	"database/sql"
	"errors"
)

type ProductModel struct {
	DB *sql.DB
}

func (pm *ProductModel) Create(name, description string, price, qty int) (int, error) {
	query := `
		INSERT INTO products (name, description, price, qty)
		VALUES (?, ?, ?, ?)
	`
	result, err := pm.DB.Exec(query, name, description, price, qty)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (pm *ProductModel) Get(id int) (*Product, error) {
	query := `
		SELECT id, name, description, price, qty
		FROM products
		WHERE id = ?
	`
	product := &Product{}
	row := pm.DB.QueryRow(query, id)
	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Qty,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return product, nil
}

func (pm *ProductModel) GetAll() ([]*Product, error) {
	query := `
		SELECT id, name, description, price, qty
		FROM products
		ORDER BY id DESC
	`
	rows, err := pm.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*Product{}
	for rows.Next() {
		product := &Product{}
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Qty,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (pm *ProductModel) Update(product *Product) error {
	query := `
		UPDATE products
		SET name = ?, description = ?, price = ?, qty = ?
		WHERE id = ?
	`
	result, err := pm.DB.Exec(
		query,
		product.Name,
		product.Description,
		product.Price,
		product.Qty,
		product.ID,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (pm *ProductModel) Delete(id int) error {
	query := `
		DELETE FROM products
		WHERE id = ?	
	`
	result, err := pm.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
