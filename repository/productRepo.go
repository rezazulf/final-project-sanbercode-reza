package repository

import (
	"database/sql"
	"final-project/models"
	"time"
)

func GetAllProduct(db *sql.DB) (err error, results []models.Product) {

	query := "SELECT * FROM product ORDER BY ID ASC"

	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var product = models.Product{}

		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Image_url,
			&product.Stock,
			&product.Status_id,
			&product.Category_id,
			&product.Created_at,
			&product.Updated_at,
		)

		results = append(results, product)
	}

	return
}

func InsertProduct(db *sql.DB, product models.Product) (err error) {
	query := "INSERT INTO product(name, description, price, image_url, stock, status_id, category_id) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	if product.Price <= 0 {
		panic(err)
	}

	errs := db.QueryRow(
		query,
		product.Name,
		product.Description,
		product.Price,
		product.Image_url,
		product.Stock,
		product.Status_id,
		product.Category_id)
	return errs.Err()
}

func UpdateProduct(db *sql.DB, product *models.Product) (err error) {
	query := "UPDATE product SET name = $1, updated_at = $2,description = $3, price = $4, image_url = $5, stock = $6, status_id = $7, category_id = $8 WHERE id = $9"
	if product.Price <= 0 {
		panic(err)
	}
	product.Updated_at = time.Now().Format(time.RFC3339)

	errs := db.QueryRow(
		query,
		product.Name,
		product.Updated_at,
		product.Description,
		product.Price,
		product.Image_url,
		product.Stock,
		product.Status_id,
		product.Category_id,
		product.ID,
	)
	return errs.Err()
}

func DeleteProduct(db *sql.DB, product models.Product) (err error) {
	sql := "DELETE FROM product WHERE id = $1"
	_, errs := db.Exec(sql, product.ID)
	return errs
}
