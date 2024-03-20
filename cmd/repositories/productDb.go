package repositories

import (
	"database/sql"
	"marketplace/cmd/models"
	"marketplace/cmd/storage"
)

func CreateProduct(product models.Product) (models.Product, error) {
	db := storage.GetDB()
	sqlStatement := `INSERT INTO products (name, description, stock, price) VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.QueryRow(sqlStatement, product.Name, product.Description, product.Stock, product.Price).Scan(&product.Id)
	if err != nil {
		return product, err
	}
	return product, nil
}

func UpdateProduct(product models.Product, id int) (models.Product, error) {
	db := storage.GetDB()
	sqlStatement := `
    UPDATE products
    SET name = $2, description = $3, stock = $4, price = $5
    WHERE id = $1
    RETURNING id`
	err := db.QueryRow(sqlStatement, id, product.Name, product.Description, product.Stock, product.Price).Scan(&id)
	if err != nil {
		return models.Product{}, err
	}
	product.Id = id
	return product, nil
}

func GetProduct(id int) (models.Product, error) {
	db := storage.GetDB()
	var product models.Product
	sqlStatement := `
    SELECT id, name, description, stock, price
    FROM products
    WHERE id = $1`
	err := db.QueryRow(sqlStatement, id).Scan(&product.Id, &product.Name, &product.Description, &product.Stock, &product.Price)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func GetAllProducts() ([]models.Product, error) {
	db := storage.GetDB()

	rows, err := db.Query("SELECT id, name, description, stock, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Stock, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func DeleteProduct(id int) error {
	db := storage.GetDB()
	sqlStatement := `
    DELETE FROM products
    WHERE id = $1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	//
	rowCount, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowCount == 0 {
		return sql.ErrNoRows
	}

	//

	return nil
}
