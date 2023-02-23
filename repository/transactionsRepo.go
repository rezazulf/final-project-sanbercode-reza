package repository

import (
	"database/sql"
	"final-project/models"
)

func GetAllTransactions(db *sql.DB) (err error, results []models.Transactions) {
	query := "SELECT * FROM transactions ORDER BY ID ASC"

	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var order = models.Transactions{}

		err = rows.Scan(
			&order.ID,
			&order.Sum_item,
			&order.Payment_bills,
			&order.Product_id,
			&order.Customer_id,
			&order.Created_at,
			&order.Updated_at,
		)

		results = append(results, order)
	}

	return
}

func PostTransactions(db *sql.DB, transactions *models.Users, dataProduct *models.Product, dataCheck *models.Transactions) (error, string, *models.Transactions, bool) {
	query := "SELECT name, stock, price FROM product WHERE id = $1"
	err := db.QueryRow(query, dataProduct.ID).Scan(&dataProduct.Name, &dataProduct.Stock, &dataProduct.Price)
	if err != nil {
		panic(err)
	}

	query = "SELECT username, balance FROM users WHERE id = $1"
	err = db.QueryRow(query, transactions.ID).Scan(&transactions.Username, &transactions.Balance)
	if err != nil {
		panic(err)
	}

	if dataCheck.Sum_item <= 0 {
		result := "jumlah barang dibeli tidak boleh < 0"
		return err, result, dataCheck, false
	}

	dataCheck.Payment_bills = dataCheck.Sum_item * dataProduct.Price

	if transactions.Balance < dataCheck.Payment_bills {
		result := "Maaf saldo anda tidak cukup"
		return err, result, dataCheck, false
	}

	if dataProduct.Stock <= 0 || dataProduct.Stock < dataCheck.Sum_item {
		result := "Maaf, stock habis atau dana tidak cukup"
		return err, result, dataCheck, false
	}

	transactions.Balance -= dataCheck.Payment_bills

	dataProduct.Stock -= dataCheck.Sum_item

	query = "UPDATE product SET stock = $1 WHERE id = $2"
	_, err = db.Exec(query, dataProduct.Stock, dataProduct.ID)
	if err != nil {
		panic(err)
	}

	query = "UPDATE users SET balance = $1 WHERE id = $2"
	_, err = db.Exec(query, transactions.Balance, transactions.ID)
	if err != nil {
		panic(err)
	}

	query = "INSERT INTO transactions (product_id, customer_id, sum_item, payment_bills) VALUES ($1, $2, $3, $4)"
	_, err = db.Exec(query, dataProduct.ID, transactions.ID, dataCheck.Sum_item, dataCheck.Payment_bills)
	if err != nil {
		panic(err)
	}

	result := "Purchase Success"
	return err, result, dataCheck, true

}
