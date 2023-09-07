package controller

import (
	"echo/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (cn *Controller) StoreWeatherbyId(c echo.Context) error {

	id := c.Param("id")
	intId, _ := strconv.Atoi(id)

	storeById, err := cn.Repo.GetById(intId)
	if err != nil {
		panic(err)
	}

	urlAddress := fmt.Sprintf("https://weather-by-api-ninjas.p.rapidapi.com/v1/weather?city=%s", storeById.Address)

	url := urlAddress
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-RapidAPI-Key", "2d2c39d15cmsh403852f286a60f6p1f68c5jsn89ba3ddc4108")
	req.Header.Add("X-RapidAPI-Host", "weather-by-api-ninjas.p.rapidapi.com")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	defer res.Body.Close()

	var weather models.Weather

	if err := json.NewDecoder(res.Body).Decode(&weather); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"store":   storeById,
		"weather": weather,
	})
}

func (cn *Controller) GetAllStore(c echo.Context) error {
	Allstore, err := cn.Repo.FindStore()
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, Allstore)

}
