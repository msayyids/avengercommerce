package models

type Transaction struct {
	Id           int     `json:"id"`
	User_id      int     `json:"user_id"`
	Product_id   int     `json:"product_id"`
	Quantity     int     `json:"quantity" binding:"required"`
	Total_amount float32 `json:"total_amount"`
}
