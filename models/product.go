package models

type Product struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Image_url   string `json:"image_url"`
	Stock       int64  `json:"stock"`
	Status_id   int64  `json:"status_id"`
	Category_id int64  `json:"category_id"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}
