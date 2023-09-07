package middleware

import (
	"echo/helper"
	"echo/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func FindById(id int, db *gorm.DB) (models.Users, error) {
	var user models.Users
	query := db.Table("users").Where("id=?", id).Find(&user)
	if query.Error != nil {
		return models.Users{}, query.Error
	}

	return user, nil
}

func AuthMiddleware(next echo.HandlerFunc, db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("token")

		claims, err := helper.ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "token invalid")
		}

		id := int(claims["id"].(float64))
		LoggedInId, err := FindById(id, db)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "isi token invalid")
		}
		c.Set("LoggedInId", LoggedInId.Id)
		return next(c)
	}
}
