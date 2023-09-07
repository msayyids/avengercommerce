package repository

import (
	"echo/models"
	"fmt"

	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func (r *Repo) AddUser(username, password string, depositeaomut float32) (models.Users, error) {
	input := models.Users{UserName: username, Password: password, DepositAmount: depositeaomut}
	query := r.DB.Select("user_name", "password", "deposit_amount").Create(&input)
	if query.Error != nil {
		return models.Users{}, fmt.Errorf("erorr nih")
	}
	return input, nil
}

func (r *Repo) FindByUserName(username string) (models.Users, error) {
	var user models.Users
	query := r.DB.First(&user, "user_name=?", username)
	if query.Error != nil {
		return models.Users{}, query.Error
	}

	return user, nil
}
