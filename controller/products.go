package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (cn *Controller) GetAllProduct(c echo.Context) error {

	AllProduct, err := cn.Repo.GetAll()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "product empty",
		})
	}

	return c.JSON(http.StatusOK, AllProduct)
}
