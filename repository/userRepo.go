package repository

import (
	"database/sql"
	"final-project/models"
	"time"
)

func GetAdmins(db *sql.DB) (results []models.Users, err error) {
	query := "SELECT * FROM users WHERE role = 'Admin' ORDER BY id ASC"

	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var admin = models.Users{}

		err = rows.Scan(
			&admin.ID,
			&admin.Username,
			&admin.Password,
			&admin.Balance,
			&admin.Role,
			&admin.Created_at,
			&admin.Updated_at,
		)
		if err != nil {
			panic(err)
		}

		results = append(results, admin)
	}

	return
}

func GetUser(db *sql.DB) (results []models.Users, err error) {
	sql := "SELECT * FROM users WHERE role = 'Customer' ORDER BY id ASC"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user = models.Users{}

		err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Balance,
			&user.Role,
			&user.Created_at,
			&user.Updated_at,
		)
		if err != nil {
			panic(err)
		}

		results = append(results, user)
	}

	return
}

func EditBalance(db *sql.DB, users *models.Users) (err error) {
	query := `UPDATE users SET balance = $1, updated_at = $2 WHERE id = $3`
	if users.Balance <= 0 {
		panic(err)
	}

	users.Updated_at = time.Now().Format(time.RFC3339)

	errs := db.QueryRow(
		query,
		users.Balance,
		users.Updated_at,
		users.ID)
	return errs.Err()
}
