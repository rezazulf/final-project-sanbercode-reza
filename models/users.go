package models

type Users struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Balance    int64  `json:"balance"`
	Role       string `json:"role"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
