package repository

import (
	"database/sql"
	"final-project/models"
)

func GetAllStatus(db *sql.DB) (results []models.Status, err error) {
	query := "SELECT * FROM status ORDER by id ASC"
	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var status = models.Status{}
		err = rows.Scan(
			&status.ID,
			&status.Status,
		)
		if err != nil {
			panic(err)
		}

		results = append(results, status)
	}

	return
}

func InsertStatus(db *sql.DB, status models.Status) (err error) {
	query := "INSERT INTO status (status) VALUES ($1)"
	errs := db.QueryRow(query, status.Status)
	return errs.Err()
}
