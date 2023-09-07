package models

type Products struct {
	Id    int     `json:"id" `
	Name  string  `json:"name" binding:"required"`
	Stock int     `json:"stock" binding:"required"`
	Price float32 `json:"price" binding:"required"`
}
