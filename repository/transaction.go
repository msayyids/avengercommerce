package repository

import (
	"echo/models"
	"fmt"
)

func (r *Repo) Transactions(loggedinid, productID, quantity int) (models.Transaction, error) {

	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var transaction models.Transaction

	var product models.Products
	if err := tx.Where("id = ?", productID).First(&product).Error; err != nil {
		tx.Rollback()
		return transaction, fmt.Errorf("product not found")
	}

	userID := loggedinid
	var user models.Users
	if err := tx.Where("id = ?", userID).First(&user).Error; err != nil {
		tx.Rollback()
		return transaction, fmt.Errorf("user not found")
	}

	if user.DepositAmount < float32(quantity)*product.Price {
		tx.Rollback()
		return transaction, fmt.Errorf("balance is not enough -- gagal terooos")
	}

	if product.Stock < quantity {
		tx.Rollback()
		return transaction, fmt.Errorf("stock is not enough/empty -- gagal lagiii")
	}

	user.DepositAmount -= float32(quantity) * product.Price
	product.Stock -= quantity

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return transaction, err
	}

	if err := tx.Save(&product).Error; err != nil {
		tx.Rollback()
		return transaction, err
	}

	inputTransaction := models.Transaction{User_id: userID, Product_id: productID, Quantity: quantity, Total_amount: float32(quantity) * product.Price}
	if err := tx.Create(&inputTransaction).Error; err != nil {
		tx.Rollback()
		return transaction, err
	}

	tx.Commit()

	return inputTransaction, nil
}
