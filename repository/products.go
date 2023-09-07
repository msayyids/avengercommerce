package repository

import (
	"echo/models"
	"fmt"
)

func (r *Repo) GetAll() ([]models.Products, error) {
	var products []models.Products
	query := r.DB.Find(&products)
	if query.Error != nil {
		return []models.Products{}, fmt.Errorf("error get all")
	}

	return products, nil
}
