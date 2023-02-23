package models

type Transactions struct {
	ID            int64  `json:"id"`
	Sum_item      int64  `json:"sum_item"`
	Payment_bills int64  `json:"payment_bills"`
	Created_at    string `json:"created_at"`
	Updated_at    string `json:"updated_at"`
	Product_id    int64  `json:"product_id"`
	Customer_id   int64  `json:"customer_id"`
}
