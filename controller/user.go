package controller

import (
	"echo/helper"
	"echo/models"
	"echo/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	Repo repository.Repo
}

func (cn *Controller) Register(c echo.Context) error {
	var reqBody models.Users

	err := c.Bind(&reqBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	result, err := cn.Repo.AddUser(reqBody.UserName, helper.Hashpw(reqBody.Password), reqBody.DepositAmount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	_ = result

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "succes create",
	})

}

func (cn *Controller) Login(c echo.Context) error {
	requestBody := models.Users{}
	// decode
	err := c.Bind(&requestBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if requestBody.UserName == "" || requestBody.Password == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid input, value cannot be empty",
		})
	}

	loginUser, err := cn.Repo.FindByUserName(requestBody.UserName)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "invalid email/password",
		})
	}

	validPasswd := helper.ValidatePw(requestBody.Password, loginUser.Password)

	if !validPasswd {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "invalid email/password",
		})
	}

	token, _ := helper.GenerateToken(uint(loginUser.Id))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes login",
		"token":   token,
	})

}
