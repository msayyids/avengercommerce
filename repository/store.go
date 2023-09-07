package repository

import (
	"echo/models"
	"fmt"
)

func (r *Repo) FindStore() ([]models.Store, error) {
	var store []models.Store
	query := r.DB.Table("store").Find(&store)
	if query.Error != nil {
		return []models.Store{}, fmt.Errorf("erorr lgi error lgi")
	}

	return store, nil
}

func (r *Repo) GetById(id int) (models.Store, error) {
	var store models.Store
	query := r.DB.Table("store").Where("id=?", id).Find(&store)
	if query.Error != nil {
		return models.Store{}, fmt.Errorf("errooor")
	}

	return store, nil
}
