package controller

import (
	"echo/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (cn *Controller) Transaction(c echo.Context) error {
	userId := c.Get("LoggedInId").(int)
	var reqTransacton models.Transaction

	err := c.Bind(&reqTransacton)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	transaction, err := cn.Repo.Transactions(userId, reqTransacton.Product_id, reqTransacton.Quantity)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "transaction failed",
		})
	}

	return c.JSON(http.StatusOK, transaction)
}
