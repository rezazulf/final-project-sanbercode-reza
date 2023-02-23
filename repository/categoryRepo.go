package repository

import (
	"database/sql"
	"final-project/models"
	"time"
)

func GetAllCategory(db *sql.DB) (results []models.Category, err error) {
	query := "SELECT * FROM category ORDER by id ASC"
	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var category = models.Category{}
		err = rows.Scan(
			&category.ID,
			&category.Name,
			&category.Created_at,
			&category.Updated_at,
		)
		if err != nil {
			panic(err)
		}

		results = append(results, category)
	}

	return
}

func InsertCategory(db *sql.DB, category models.Category) (err error) {
	query := "INSERT INTO category (name) VALUES ($1)"
	errs := db.QueryRow(query, category.Name)
	return errs.Err()
}

func UpdateCategory(db *sql.DB, category models.Category) (err error) {
	time := time.Now().Format(time.RFC3339)
	query := "UPDATE category SET name = $1, updated_at = $2 WHERE id = $3"
	category.Updated_at = time
	errs := db.QueryRow(
		query,
		category.Name,
		category.Updated_at,
		category.ID)
	return errs.Err()
}

func DeleteCategory(db *sql.DB, category models.Category) (err error) {
	query := "DELETE FROM category WHERE id = $1"
	errs := db.QueryRow(
		query,
		category.ID)
	return errs.Err()
}
